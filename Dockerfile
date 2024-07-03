FROM docker.m.daocloud.io/library/golang:1.20 AS builder
WORKDIR /app
ARG VERSION=latest
ARG GO_PKG=go.k6.io/k6/lib/consts
ARG PERF_PKG=github.com/goharbor/xk6-harbor
RUN git clone https://github.com/goharbor/xk6-harbor.git \
    && cd xk6-harbor \
    && git rev-parse --short HEAD > commit.txt \
    && CGO_ENABLED=0 go build -ldflags="-extldflags -static -X  ${GO_PKG}.VersionDetails=$(cat commit.txt)" -o k6-harbor ./cmd/k6/main.go
ADD . /app
RUN cd /app \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && CGO_ENABLE=0 go run mage.go -compile ./magef --debug --ldflags="-X ${PERF_PKG}.Version=${VERSION}"

FROM alpine:latest AS modules
WORKDIR /app
ADD package.json /app
ADD package-lock.json /app
RUN cd /app \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update && apk add --no-cache nodejs npm \
    && npm install



FROM mcr.microsoft.com/cbl-mariner/distroless/debug:2.0
WORKDIR /app

ENV HARBOR_SIZE "small"
ENV HARBOR_VUS 100
ENV HARBOR_ITERATIONS 200
ENV HARBOR_REPORT true
ENV K6_ALWAYS_UPDATE false
ENV K6_QUIET false
ENV K6_PROMETHEUS_RW_INSECURE_SKIP_TLS_VERIFY false
ENV K6_PROMETHEUS_RW_TREND_AS_NATIVE_HISTOGRAM true
ENV K6_PROMETHEUS_RW_TREND_STATS=p(95),p(99),min,max

COPY --from=builder /app/magef /app/mage
COPY --from=builder /app/k6-harbor /usr/local/bin/k6-harbor
COPY --from=builder /app/package.json /app/package.json
COPY --from=modules /app/node_modules /app/node_modules
COPY --from=builder /app/scripts /app/scripts




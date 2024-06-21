FROM docker.m.daocloud.io/library/golang:1.20 AS builder
WORKDIR /app
ADD . /app
RUN cd /app \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && CGO_ENABLE=0 go run mage.go -compile ./magef

FROM alpine:latest AS modules
WORKDIR /app
ADD package.json /app
ADD package-lock.json /app
RUN cd /app \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update && apk add --no-cache nodejs npm \
    && npm install



FROM gcr.m.daocloud.io/distroless/base@sha256:5e3fac1733c75e0e879a9770724e3960610a5cfbbfb5366559fbc334fe86c249 AS final
WORKDIR /app

ENV HARBOR_SIZE "small"
ENV HARBOR_VUS 100
ENV HARBOR_ITERATIONS 200
ENV HARBOR_REPORT true
ENV K6_ALWAYS_UPDATE false
ENV K6_QUIET false
ENV K6_JSON_OUTPUT true
ENV K6_PROMETHEUS_RW_INSECURE_SKIP_TLS_VERIFY false
ENV K6_PROMETHEUS_RW_TREND_AS_NATIVE_HISTOGRAM true
ENV K6_PROMETHEUS_RW_TREND_STATS=p(95),p(99),min,max

COPY --from=builder /app/magef /app/mage
COPY --from=builder /app/xk6-harbor/k6 /usr/local/bin/k6-harbor
COPY --from=builder /app/package.json /app/package.json
COPY --from=modules /app/node_modules /app/node_modules
COPY --from=builder /app/scripts /app/scripts




FROM golang:alpine AS builder
WORKDIR /app
ADD . /app
RUN cd /app \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o  mage mage.go \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update && apk add --no-cache nodejs npm \
    && npm install \


FROM golang:alpine
WORKDIR /app
COPY --from=builder /app/mage /app/mage
COPY --from=builder /app/magefile.go /app/magefile.go
COPY --from=builder /app/go.mod /app/go.mod
COPY --from=builder /app/go.sum /app/go.sum
COPY --from=builder /app/node_modules /app/node_modules
COPY --from=builder /app/xk6-harbor/k6 /app/k6

RUN mv /app/k6 /usr/local/bin/k6





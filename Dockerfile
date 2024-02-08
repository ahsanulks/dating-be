FROM golang:1.21 AS builder

COPY . /app
WORKDIR /app

RUN GOPROXY=https://goproxy.cn make build
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /app/bin /app
COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /go/bin/goose /app/goose

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ./goose up && ./app

FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY ./ ./

RUN go mod download
RUN GOOS=linux go build -o /api cmd/api/main.go

FROM debian:12-slim AS release

RUN apt-get update && apt-get install -y --no-install-recommends tzdata && apt-get clean
ENV TZ=Asia/Saigon
RUN ln -snf "/usr/share/zoneinfo/$TZ" /etc/localtime && echo "$TZ" > /etc/timezone
RUN apt-get install -y ca-certificates && apt-get clean

WORKDIR /

COPY --from=builder /api /api
COPY --from=builder /app/config.prod.yaml /config.yaml
COPY --from=builder /app/internal/adapter/repository/postgres/database/migration/seeders/ /internal/adapter/repository/postgres/database/migration/seeders/

ENTRYPOINT [ "/api" ]

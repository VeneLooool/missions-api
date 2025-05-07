FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

# ВАЖНО: явно указываем GOOS и GOARCH
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o missions-api ./cmd/missions-api

FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/missions-api .

COPY swagger-ui ./swagger-ui/
COPY internal ./internal/

EXPOSE 8082 50053

CMD ["./missions-api"]

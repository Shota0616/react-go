# ビルドステージ
FROM golang:1.22.4 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# 本番ステージ
FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
RUN chown nobody:nobody /root/main
USER nobody
CMD ["./main"]

# 開発ステージ
FROM golang:1.22.4 AS development
WORKDIR /app
RUN go install github.com/githubnemo/CompileDaemon@latest
COPY . .
CMD CompileDaemon -build="go build -o main main.go" -command="./main"

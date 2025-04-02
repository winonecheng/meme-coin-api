# Stage 1: Build
FROM golang:1.24 AS builder
WORKDIR /app

# Install SQLite development libraries for CGO
RUN apt-get update && apt-get install -y gcc libc6-dev sqlite3 libsqlite3-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Enable CGO and build a static binary
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-linkmode external -extldflags -static" -o meme-coin-api

# Stage 2: Run
FROM alpine:latest
RUN apk add --no-cache sqlite-libs
WORKDIR /app

COPY --from=builder /app/meme-coin-api .
COPY meme_coins.db .

EXPOSE 8080
CMD ["./meme-coin-api"]

# Stage 1: Build
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o meme-coin-api

# Stage 2: Run
FROM alpine:latest
RUN apk add --no-cache sqlite-libs
WORKDIR /app
COPY --from=builder /app/meme-coin-api .
COPY meme_coins.db .
EXPOSE 8080
CMD ["./meme-coin-api"]

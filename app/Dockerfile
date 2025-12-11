FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o football-service ./cmd/server

FROM alpine:latest AS runner

WORKDIR /

COPY --from=builder /app/football-service .

EXPOSE 8080

CMD ["./football-service"]
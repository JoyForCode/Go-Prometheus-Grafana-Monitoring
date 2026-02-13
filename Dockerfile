FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o exporter ./cmd/exporter

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/exporter .

EXPOSE 9105

CMD ["./exporter"]

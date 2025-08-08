FROM golang:1.24.0-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o gateway .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gateway .
COPY .env .
EXPOSE 8080
CMD ["./gateway"]
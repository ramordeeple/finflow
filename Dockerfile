FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o finflow

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/finflow .
EXPOSE 8080
CMD ["./finflow"]

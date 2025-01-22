FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
WORKDIR /app/services/trip-service
RUN CGO_ENABLED=0 GOOS=linux go build -o trip-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/services/trip-service/trip-service .
EXPOSE 9093
CMD ["./trip-service"] 
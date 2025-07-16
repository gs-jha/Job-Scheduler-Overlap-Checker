# Build stage
FROM golang:1.21.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o job-scheduler-overlap-checker ./cmd/main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/job-scheduler-overlap-checker .
EXPOSE 8080
CMD ["./job-scheduler-overlap-checker"]

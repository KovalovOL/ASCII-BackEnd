FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -v -o /app/main ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8080
CMD ["./main"]
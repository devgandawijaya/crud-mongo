FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /root
COPY --from=builder /app/app .
COPY .env.example .env
EXPOSE 3030
CMD ["./app"]

FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/gateway_bin main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder /app/gateway_bin .


RUN apk add --no-cache curl

RUN chmod +x /app/gateway_bin

CMD ["./gateway_bin"]

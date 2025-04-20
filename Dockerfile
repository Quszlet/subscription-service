FROM golang:1.23.4-alpine AS builder

RUN go install github.com/air-verse/air@latest 

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

FROM alpine:latest

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/main .

RUN mkdir -p logs

CMD ["./main"]
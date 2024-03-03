FROM golang:1.22.0 AS builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -ldflags '-s -w' -o app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 3000

CMD ["./app"]
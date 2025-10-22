FROM golang:1.23.5-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o vehicle-resale-api ./src/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/vehicle-resale-api .

EXPOSE 8080

CMD ["./vehicle-resale-api"]

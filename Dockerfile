FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .
COPY .env .
COPY cloud_credentials.json .
COPY start-script.sql .

EXPOSE 8000

CMD ["./main"]

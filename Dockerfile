# syntax=docker/dockerfile:1

FROM golang:latest AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main
EXPOSE 3500
CMD ["./main"]

#docker build --tag prueba-go .
#docker run --rm -it -p 3500:3500 prueba-go

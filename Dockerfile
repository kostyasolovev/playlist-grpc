# syntax=docker/dockerfile:1

FROM golang:latest
WORKDIR /grpc-server
COPY . .
RUN go mod download
RUN go build -o grpc-server /cmd/main.go
EXPOSE 8081
CMD ./grpc-server

# syntax=docker/dockerfile:1

FROM golang:latest
ENV YTAPIKEY=AIzaSyBtkJgSIgfjM_4sb5vjac5d3j2H7l6gwhw
WORKDIR /grpc-server
COPY . .
RUN go mod download
RUN go build -o ./grpc-server cmd/main.go
EXPOSE 8083
CMD ./grpc-server

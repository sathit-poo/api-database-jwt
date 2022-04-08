FROM golang:1.16.2-stretch AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main .
CMD ["/app/main"]
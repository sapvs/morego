FROM golang:1.15.2-alpine3.12 AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod . 
COPY go.sum .
RUN go mod download

ADD . /app
RUN go test -v */**
RUN go build -o main

FROM alpine:3.12 AS production 
COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]
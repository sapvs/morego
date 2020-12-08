FROM golang:1.15.2-alpine3.12 AS builder
RUN apk add build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main

FROM alpine:3.12 AS production 
COPY --from=builder /app .

RUN apk update \
    && apk add sqlite

ENTRYPOINT [ "./main" ]


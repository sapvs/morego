FROM golang:1.15.2-alpine3.12 AS builder

LABEL author="sudosapan@gmail.com"

RUN apk add --no-cache gcc musl-dev

RUN mkdir /app
WORKDIR /app

COPY go.mod . 
COPY go.sum .
RUN go mod download

ADD . /app
RUN go build -o main

FROM alpine:3.12 AS production 
COPY --from=builder /app .

# For including sqlite in image
RUN apk --no-cache add sqlite

ENTRYPOINT [ "./main" ]

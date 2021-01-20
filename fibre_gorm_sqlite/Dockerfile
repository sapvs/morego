FROM sudosapan/go-alpine-gcc:latest AS builder

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

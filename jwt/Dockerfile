FROM golang:1.15.2-alpine3.12 AS builder

RUN mkdir /app
WORKDIR /app
# Supplied as dockerbuild arg for creating specific image
ARG MAIN_FILE

COPY go.mod . 
COPY go.sum .
RUN go mod download

ADD . /app
RUN go build -o main ${MAIN_FILE}

FROM alpine:3.12 AS production 
COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]
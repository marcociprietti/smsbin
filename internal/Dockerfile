FROM golang:1.17-alpine

WORKDIR /go/src

COPY docker/air_conf.toml ./.air.toml

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
FROM golang:1.14

RUN go get -u github.com/pressly/goose/cmd/goose

RUN apt-get update && apt-get install -y postgresql-client

FROM golang:1.19-alpine

WORKDIR /jack/server

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download

CMD [ "air" ]
# syntax=docker/dockerfile:1

FROM golang:1.23

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN go build -o /Mensa-API

EXPOSE 80

CMD [ "/Mensa-API" ]
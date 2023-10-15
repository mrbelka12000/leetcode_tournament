## Build
FROM golang:1.20.10-alpine3.18 AS buildenv

ADD go.mod go.sum /

RUN go mod download

WORKDIR /app

ADD . .

ADD .env .

ENV GO111MODULE=on

RUN  go build -o main cmd/main.go

## Deploy
FROM alpine

WORKDIR /

COPY --from=buildenv  /app/main /main

EXPOSE 3000

CMD ["/main"]
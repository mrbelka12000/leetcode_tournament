FROM golang:1.16

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8080

ENV GIN_MODE=release

CMD ./main
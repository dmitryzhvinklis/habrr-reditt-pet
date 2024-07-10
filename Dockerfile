FROM golang:latest

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN go build -o ozon-test ./cmd

EXPOSE 8080

CMD ["./ozon-test"]

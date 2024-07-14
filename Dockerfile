FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ozon-test cmd/main.go

CMD ["./ozon-test"]

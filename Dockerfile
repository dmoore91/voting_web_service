FROM golang:1.15.2

WORKDIR /go/src/app

COPY . .

CMD ["go", "run", "cmd/data-access-api/main.go"]

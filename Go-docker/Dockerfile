FROM golang:1.15.2

RUN apt-get update
RUN apt-get install vim -y
RUN go get "github.com/go-sql-driver/mysql"

# Create dir and copy all files
WORKDIR /go/src/app/

COPY go.sum go.mod ./
COPY cmd internal ./

# Run application
CMD ["go", "run", "cmd/data-access-api/main.go"]

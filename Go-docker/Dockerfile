FROM golang:1.15.2

# Create dir and copy all files
WORKDIR /go/src/app/

COPY go.sum go.mod ./
COPY cmd internal ./

# Run application
CMD ["go", "run", "cmd/data-access-api/main.go"]

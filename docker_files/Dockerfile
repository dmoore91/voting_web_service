FROM golang:1.15.2
RUN apt-get update
RUN apt-get install vim -y
RUN go get "github.com/go-sql-driver/mysql"

# Create dir and copy all files
WORKDIR /go/src/app/

RUN apt-get update && apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get update && apt-get install -y nodejs
RUN npm install speakeasy qrcode
COPY package*.json ./

# COPY go.sum go.mod /go/src/app/
# COPY cmd internal /go/src/app/
#COPY ./ ./

EXPOSE 8880

# Run application
#CMD ["go", "run", "cmd/data-access-api/main.go"]
#CMD ["/bin/bash"]

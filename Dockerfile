FROM golang:1.19-alpine

RUN apk update && apk add --no-cache make protobuf-dev

WORKDIR /app/
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
RUN apk add --no-cache protobuf git
RUN go get github.com/golang/protobuf/protoc-gen-go


COPY . .
RUN #go generate
RUN go build -o data-handler main.go
EXPOSE 9009

CMD ["./app/data-handler"]
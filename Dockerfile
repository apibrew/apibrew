FROM golang:1.19-alpine as buildenv

RUN apk update && apk add --no-cache make protobuf-dev
RUN apk add --no-cache protobuf git
RUN wget https://github.com/bufbuild/buf/releases/download/v1.12.0/buf-Linux-x86_64
RUN mv buf-Linux-x86_64 /bin/buf
RUN chmod +x /bin/buf
RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

WORKDIR /app/
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY proto proto
COPY generate.go generate.go

COPY data data
COPY app app
COPY cmd cmd
COPY helper helper
COPY logging logging
COPY server server
COPY service service
COPY util util

FROM buildenv as test
WORKDIR /app/

# setup database
RUN apk add postgresql
RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/

RUN sh -c "cd proto; buf mod update"
RUN go generate

COPY test test
COPY run-tests.sh .

RUN sh /app/run-tests.sh

FROM buildenv as builder

RUN sh -c "cd proto; buf mod update"
RUN go generate

RUN go build -o data-handler cmd/server/main.go

FROM golang:1.19-alpine
WORKDIR /

COPY data /data
COPY --from=builder /app/data-handler /bin/data-handler

EXPOSE 9009

CMD ["/bin/data-handler", "-init", "/data/init.json"]
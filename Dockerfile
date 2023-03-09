FROM golang:1.19-alpine as buildenv

WORKDIR /app/
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY proto proto
COPY cmd cmd
COPY pkg pkg

FROM buildenv as test
WORKDIR /app/

# setup database
RUN apk add postgresql
RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/

RUN sh /app/pkg/test/run-tests.sh

FROM buildenv as builder

RUN go build -o data-handler cmd/server/main.go

FROM golang:1.19-alpine
WORKDIR /

COPY --from=builder /app/data-handler /bin/data-handler

EXPOSE 9009

CMD ["/bin/data-handler", "-init", "/app/config.json"]
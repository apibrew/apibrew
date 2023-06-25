FROM golang:1.19-alpine as buildenv

WORKDIR /app/
RUN apk add build-base

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cmd cmd
COPY pkg pkg

FROM buildenv as apbr

RUN go build -o apbr cmd/apbr/main.go

ENTRYPOINT ["/app/apbr"]

FROM buildenv as builder

RUN go build -o apibrew cmd/apbr-server/main.go

FROM golang:1.19-alpine as app
WORKDIR /

RUN apk --update --no-cache add curl
COPY --from=builder /app/apibrew /bin/apibrew

EXPOSE 9009

CMD ["/bin/apibrew", "-init", "/app/config.json"]

FROM golang:1.19-alpine as app-full
WORKDIR /

RUN apk add postgresql
RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/
RUN apk --update --no-cache add curl
VOLUME /var/lib/postgresql/data

COPY --from=builder /app/apibrew /bin/apibrew
COPY run/run-standalone-postgres.sh /app/run.sh
COPY run/init.sql /app/init.sql
COPY run/config.json /app/config.json

EXPOSE 9009

RUN ls -alsh /app

CMD ["/bin/sh", "/app/run.sh"]

FROM buildenv as test
WORKDIR /app/

# setup database
RUN apk add postgresql
RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/

RUN sh /app/pkg/test/run-tests.sh

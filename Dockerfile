FROM golang:1.19-alpine as buildenv

WORKDIR /app/
RUN apk add build-base

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

FROM buildenv as build-binaries

ARG OS
ARG ARCH
ARG APP
ARG SUFFIX
RUN mkdir /app/dist

RUN sh -c "env GOOS=${OS} GOARCH=${ARCH} go build -o /app/dist/${APP}-${OS}-${ARCH}${SUFFIX} cmd/${APP}/main.go"

FROM scratch as build-binaries-export

ARG OS
ARG ARCH
ARG APP
ARG SUFFIX

COPY --from=build-binaries /app/dist/${APP}-${OS}-${ARCH}${SUFFIX} /app/output/${APP}-${OS}-${ARCH}${SUFFIX}

FROM golang:1.19-alpine as app-full
WORKDIR /

RUN apk add postgresql
RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/
VOLUME /var/lib/postgresql/data

COPY --from=builder /app/data-handler /bin/data-handler
COPY run/run-standalone-postgres.sh /app/run.sh
COPY run/init.sql /app/init.sql
COPY run/config.json /app/config.json

EXPOSE 9009

RUN ls -alsh /app

CMD ["/bin/sh", "/app/run.sh"]

FROM golang:1.19-alpine as app
WORKDIR /

COPY --from=builder /app/data-handler /bin/data-handler

EXPOSE 9009

CMD ["/bin/data-handler", "-init", "/app/config.json"]
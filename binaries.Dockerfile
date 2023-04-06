FROM golang:1.19-alpine as buildenv

WORKDIR /app/
RUN apk add build-base

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY proto proto
COPY cmd cmd
COPY pkg pkg

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

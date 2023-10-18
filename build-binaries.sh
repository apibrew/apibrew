#!/bin/sh

mkdir -p dist

buildBinary() {
  OS=$1
  ARCH=$2
  SUFFIX=$3

  echo "Building apbr for ${OS}/${ARCH}${SUFFIX}"
  GOOS=${OS} GOARCH=${ARCH} go build -o dist/apbr-${OS}-${ARCH}${SUFFIX} cmd/apbr/main.go
  echo "Building apibrew for ${OS}/${ARCH}${SUFFIX}"
  GOOS=${OS} GOARCH=${ARCH} go build -o dist/apibrew-server-${OS}-${ARCH}${SUFFIX} cmd/apbr-server/main.go
}

buildBinary linux amd64 ""
buildBinary linux 386 ""
buildBinary linux arm64 ""
buildBinary windows 386 ".exe"
buildBinary windows amd64 ".exe"
buildBinary darwin amd64 ""
buildBinary darwin arm64 ""

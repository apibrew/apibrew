#!/bin/sh

mkdir -p dist

buildBinary() {
  OS=$1
  ARCH=$2
  SUFFIX=$3

  echo "Building dhctl for ${OS}/${ARCH}${SUFFIX}"
  GOOS=${OS} GOARCH=${ARCH} go build -o dist/dhctl-${OS}-${ARCH}${SUFFIX} cmd/dhctl/main.go
  echo "Building apibrew for ${OS}/${ARCH}${SUFFIX}"
  GOOS=${OS} GOARCH=${ARCH} go build -o dist/apibrew-${OS}-${ARCH}${SUFFIX} cmd/server/main.go
}

buildBinary linux amd64 ""
buildBinary linux 386 ""
buildBinary linux arm64 ""
buildBinary windows 386 ".exe"
buildBinary windows amd64 ".exe"
buildBinary darwin amd64 ""
buildBinary darwin arm64 ""

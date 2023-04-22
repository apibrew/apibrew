#!/usr/bin/env sh

SUFFIX=''
ARCH=$(uname -m)

echo installation for $OSTYPE $(uname -m)

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  PLATFORM="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  PLATFORM="darwin"
elif [[ "$OSTYPE" == "cygwin" ]]; then
  PLATFORM="windows"
  SUFFIX='.exe'
elif [[ "$OSTYPE" == "msys" ]]; then
  PLATFORM="windows"
  SUFFIX='.exe'
elif [[ "$OSTYPE" == "win32" ]]; then
  PLATFORM="windows"
else
  echo "Unknown OS"
  exit 1
fi

if [[ "${ARCH}" == "x86_64"* ]]; then
  PLATFORM="${PLATFORM}-amd64"
elif [[ "${ARCH}" == "i386"* ]]; then
  PLATFORM="${PLATFORM}-386"
elif [[ "${ARCH}" == "i686"* ]]; then
  PLATFORM="${PLATFORM}-386"
elif [[ "${ARCH}" == "arm64"* ]]; then
  PLATFORM="${PLATFORM}-arm64"
else
  echo "Unknown Architecture"
  exit 1
fi

echo "Downloading ${PLATFORM} binary"
curl -L -o "dhctl${SUFFIX}" "https://github.com/tislib/apibrew/releases/download/v1.1.8/dhctl-${PLATFORM}${SUFFIX}"

chmod +x dhctl
sudo mv dhctl /usr/local/bin/dhctl

echo "Dhctl installation is done"
echo "Configuring dhctl"

mkdir -p ~/.dhctl
cp ~/.dhctl/config ~/.dhctl/config.bak
curl -L -o ~/.dhctl/config "https://raw.githubusercontent.com/tislib/apibrew/master/deploy/easy-install/config"

echo "Dhctl configuration is done"

echo "Run apibrew standalone mode"

docker run --name apibrew-standalone -d -p 9009:9009 -v ${PWD}/data:/var/lib/postgresql/data tislib/apibrew:full-latest

echo "Done!"

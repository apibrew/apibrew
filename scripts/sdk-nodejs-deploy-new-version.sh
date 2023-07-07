#!/bin/bash
cd "$(dirname "$0")"/..

pwd
cd sdk/nodejs
npm run deploy

NODEJS_SDK_VERSION=`node -p "require('./package.json').version"`

cd ../..

cd extensions/nodejs-engine
npm install @apibrew/client@${NODEJS_SDK_VERSION}

cd ../..

cd examples/nodejs
npm install @apibrew/client@${NODEJS_SDK_VERSION}

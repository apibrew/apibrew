#!/usr/bin/env bash

echo --- Building my-envoy docker image ---
docker build -t dh-envoy:1.0 .


echo --- Running my-envoy docker image ---
docker run -p 50051:50051 -p 50052:50052 --name dh-envoy -d dh-envoy:1.0


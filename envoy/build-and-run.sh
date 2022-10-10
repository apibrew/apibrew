#!/usr/bin/env bash

echo --- Building my-envoy docker image ---
docker build -t dh-envoy:1.0 .


echo --- Running my-envoy docker image ---
docker run -p 9901:9901 -p 8100:8100 -d dh-envoy:1.0


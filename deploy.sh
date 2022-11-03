#!/bin/bash
cd "$(dirname "$0")"

go generate
docker-compose up --build -d

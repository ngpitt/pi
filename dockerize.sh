#!/bin/bash
set -xe

docker build -t ngpitt/pi:v1 .
docker push ngpitt/pi

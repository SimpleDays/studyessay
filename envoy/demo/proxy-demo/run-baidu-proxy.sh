#!/bin/bash

docker rm -f google-proxy

docker build -t google-proxy:v1 .

docker run -d --name google-proxy -p 8000:8000 -p 10000:10000 google-proxy:v1
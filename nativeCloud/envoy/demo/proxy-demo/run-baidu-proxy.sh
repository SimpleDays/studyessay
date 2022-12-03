#!/bin/bash

docker rm -f baidu-proxy

docker build -t baidu-proxy:v1 .

docker run -d --name baidu-proxy -p 8000:8000 -p 10000:10000 baidu-proxy:v1
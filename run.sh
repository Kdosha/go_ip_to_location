#!/bin/bash
TAG=$(git describe --tags)
docker run -p 8080:8080 "docker.io/library/ip2location:$TAG"
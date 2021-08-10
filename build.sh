#!/bin/bash
TAG=$(git describe --tags)
docker build . --tag "ip2location:$TAG"
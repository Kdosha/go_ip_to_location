$env:TAG=$(git describe --tags)
docker build . --tag "ip2location:$env:TAG"
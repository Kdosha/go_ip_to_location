FROM golang:1.16-alpine

WORKDIR /app

COPY src/ ./src
WORKDIR src

RUN go mod download

ENV GIN_MODE=release

RUN go build -o /ip2location

EXPOSE 8080

CMD [ "/ip2location" ]
FROM golang:1.16

WORKDIR /opt/fakeralic

COPY go.mod go.mod

RUN go get -d -v ./...

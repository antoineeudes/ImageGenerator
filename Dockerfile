FROM golang:1.14.6

RUN mkdir -p /go/src/ImageGenerator/
WORKDIR /go/src/ImageGenerator/
COPY . /go/src/ImageGenerator/

RUN go get -v

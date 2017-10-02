FROM golang:1.9-alpine
RUN mkdir -p /go/src/github.com/murphybytes/goroku
WORKDIR /go/src/github.com/murphybytes/goroku
ADD . .
EXPOSE 8080
RUN go install github.com/murphybytes/goroku
RUN goroku

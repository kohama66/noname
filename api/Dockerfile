FROM golang:1.13.5
RUN mkdir -p /go/src/github.com/myapp/noname/api
WORKDIR /go/src/github.com/myapp/noname/api
ADD . /go/src/github.com/myapp/noname/api
ENV GO111MODULE=on
EXPOSE 3000
CMD ["go", "run", "main.go"]
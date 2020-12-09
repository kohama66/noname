FROM golang:1.13.5
RUN mkdir -p /go/src/github.com/myapp/noname/api
WORKDIR /go/src/github.com/myapp/noname/api
ADD . /go/src/github.com/myapp/noname/api
ENV GO111MODULE=on
RUN go get github.com/pilu/fresh
EXPOSE 8080
CMD ["fresh"]
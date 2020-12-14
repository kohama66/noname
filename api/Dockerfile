FROM golang:1.13.5
RUN mkdir -p /go/src/github.com/myapp/noname/api
WORKDIR /go/src/github.com/myapp/noname/api
ADD . /go/src/github.com/myapp/noname/api
ENV GO111MODULE=on
ARG _CLOUD_SQL_CONNECTION_PASS
ARG _CLOUD_SQL_CONNECTION_NAME
ENV CLOUD_SQL_CONNECTION_PASS ${_CLOUD_SQL_CONNECTION_PASS}
ENV CLOUD_SQL_CONNECTION_NAME ${_CLOUD_SQL_CONNECTION_NAME}
RUN go get github.com/pilu/fresh
EXPOSE 8080
CMD ["fresh"]
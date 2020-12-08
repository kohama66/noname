FROM golang:1.13.5
RUN mkdir -p /go/src/github.com/myapp/noname/api
WORKDIR /go/src/github.com/myapp/noname/api
ADD . /go/src/github.com/myapp/noname/api
ENV GO111MODULE=on
# ARG _GOOGLE_APPLICATION_CREDENTIALS
# ENV GOOGLE_APPLICATION_CREDENTIALS ${_GOOGLE_APPLICATION_CREDENTIALS}
RUN go get github.com/pilu/fresh
# ENV PORT=${PORT}
EXPOSE 8080
CMD ["fresh"]
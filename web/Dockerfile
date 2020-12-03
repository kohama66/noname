FROM node:13.7.0
RUN mkdir -p /go/src/github.com/myapp/noname/web
WORKDIR /go/src/github.com/myapp/noname/web
COPY ["package.json", "yarn.lock", "./"]
COPY . /go/src/github.com/myapp/noname/web
RUN yarn install
CMD ["yarn", "start"]
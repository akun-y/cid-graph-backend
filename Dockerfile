FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/akun-y/cid-graph-backend
COPY . $GOPATH/src/github.com/akun-y/cid-graph-backend
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./cid-graph-backend"]

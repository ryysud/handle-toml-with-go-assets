FROM golang:1.10
LABEL maintainer "ryysud <ryuma.y1117@gmail.com>"

WORKDIR /go/src/github.com/ryysud/handle-toml-with-go-assets

COPY . .

RUN go get github.com/jessevdk/go-assets-builder && \
    go get github.com/BurntSushi/toml && \
    go generate && \
    go build

ENTRYPOINT ["./handle-toml-with-go-assets"]

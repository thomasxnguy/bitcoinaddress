FROM golang:1.15 AS builder

# Add source code
ADD ./ /go/src/github.com/thomasxnguy/bitcoinaddress/

RUN cd /go/src/github.com/thomasxnguy/bitcoinaddress && \
    go build

EXPOSE 3000

# Set the binary as the entrypoint of the container
WORKDIR "/go/src/github.com/thomasxnguy/bitcoinaddress"
CMD ["go","run","main.go", "serve"]
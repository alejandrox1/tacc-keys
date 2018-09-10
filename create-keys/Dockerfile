FROM golang:1.11.0-stretch AS base

WORKDIR /go/src/app
ADD keys.go main.go rsa_keys.go tokens.go /go/src/app
    
RUN go get github.com/spf13/pflag \
    && go get github.com/spf13/viper \
    && go get golang.org/x/crypto/ssh \
    && go build -o create-keys

CMD ["/bin/bash"]

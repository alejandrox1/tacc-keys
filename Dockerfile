FROM golang:1.11.0-stretch AS base

WORKDIR /src
ADD main.go .
RUN go get github.com/gorilla/handlers \
    && go get github.com/gorilla/mux \
    && go build -o keys-service


FROM centos:7

WORKDIR /app
COPY --from=base /src/keys-service /app/keys-service

ENTRYPOINT ["./keys-service"]

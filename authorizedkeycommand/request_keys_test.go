package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)


func TestGetUserPubKeys(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "ssh-rsa some-key-contents")
    }))
    defer ts.Close()

    keysEndpoint = ts.URL
}

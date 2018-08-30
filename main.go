package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)


// statusHandler serves a status update.
func statusHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Keys service is up and running")
}


func main() {
    r := mux.NewRouter()

    r.HandleFunc("/status", statusHandler).Methods("GET")

    err := http.ListenAndServe("0.0.0.0:8000", handlers.LoggingHandler(os.Stdout, r))
    if err != nil {
        log.Fatal(err)
    }
}

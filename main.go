package main

import (
    "fmt"
    "just_health/healthcheck"
    "log"
    "net/http"
)

type PGDatastore struct{}

func (pg *PGDatastore) GetHealthCheck(path string) string {
    if path == "/health" {
        return "ok"
    } else {
        return "BADBADBAD"
    }
}

func main() {
    fmt.Println("ZAPPAROO")
    datastore := &PGDatastore{}
    server := &healthcheck.NcpdpServer{
        Store: datastore,
    }
    
    err := http.ListenAndServe("localhost:5000", server)
    if err != nil {
        log.Println(err)
    }
}

package healthcheck

import (
    "fmt"
    "net/http"
)

type Datastore interface {
    GetHealthCheck(path string) string
}

type NcpdpServer struct {
    Store Datastore
}

func (server *NcpdpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    
    fmt.Fprint(w, server.Store.GetHealthCheck(path))
}

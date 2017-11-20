package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/csepanda/gonetint/server/apiv0"
)

// Key-Value record for http header
type KVRecord struct {
    key   string
    value string
}

var currentApi = &apiv0.ApiController{}
func version(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "{\"version\":\"%s\"}", currentApi.GetVersion())
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.Methods("GET").
           Path("/version").
           HandlerFunc(version)

    currentApi.Init(router)

    handler := responseHeadersPreset(router, KVRecord{
        "Content-Type",
        "application/json;charset=UTF-8",
    })

    log.Fatal(http.ListenAndServe(":9000", handler))
}

// Presets headers via handler wrapper
func responseHeadersPreset(h http.Handler, kvs ...KVRecord) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        header := w.Header()
        for _, kv := range kvs {
            header.Set(kv.key, kv.value)
        }

        h.ServeHTTP(w, r)
    })
}

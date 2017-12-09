/* This Source Code Form is subject to the terms of the Mozilla
 * Public License, v. 2.0. If a copy of the MPL was not distributed
 * with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright © 2017 Andrey Bova                                       */

package main

import (
    "fmt"
    "log"
    "flag"
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
    portPtr := flag.String("port", "8080", "port")
    flag.Parse()
    router := mux.NewRouter().StrictSlash(true)
    router.Methods("GET").
           Path("/version").
           HandlerFunc(version)

    currentApi.Init(router)

    handler := responseHeadersPreset(router, KVRecord{
        "Content-Type",
        "application/json;charset=UTF-8",
    })

    log.Fatal(http.ListenAndServe(":" + *portPtr, handler))
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

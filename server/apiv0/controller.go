/* This Source Code Form is subject to the terms of the Mozilla
 * Public License, v. 2.0. If a copy of the MPL was not distributed
 * with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright © 2017 Andrey Bova                                       */
package apiv0

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/csepanda/gonetint/domain/rv0"
)

const VERSION = "v0.1"

type ApiController struct {
    root *mux.Router
}

func (api *ApiController) GetVersion() string {
    return VERSION
}

func (api *ApiController) Init(router *mux.Router) error {
    api.root = router.PathPrefix("/" + VERSION).Subrouter()
    api.root.Methods("GET").
             Path("/interfaces").
             HandlerFunc(interfaceList)
    api.root.Methods("GET").
             Path("/interface/{name}").
             HandlerFunc(interfaceDetails)
    return nil
}

func interfaceList(w http.ResponseWriter, r *http.Request) {
    list, err := getInterfacesNames()
    if err != nil {
        responseError(err, w)
        return
    }


    jsonStr, e := json.Marshal(list)
    if e != nil {
        responseError(e, w)
    } else {
        w.Write(jsonStr)
    }
}

func interfaceDetails(w http.ResponseWriter, r *http.Request) {
    name := mux.Vars(r)["name"]
    if name == "" {
        responseError(serverError("wrong routing"), w)
        return
    }

    details, err := getInterfaceDetails(name)
    if err != nil {
        responseError(err, w)
        return
    }

    jsonStr, e := json.Marshal(details)
    if e != nil {
        responseError(e, w)
    } else {
        w.Write(jsonStr)
    }
}

func responseError(e error, w http.ResponseWriter) {
    errorResponse := rv0.ErrorResponse{e.Error()}
    jsonStr, err := json.Marshal(errorResponse)

    if err != nil {
        w.WriteHeader(500);
        fmt.Fprintf(w, "{\"Error\":\"%s\"}", err.Error())
        return
    }

    if serviceErr, ok := e.(*serviceError); ok {
        switch serviceErr.errType {
            case SERVER_SIDE_ERROR:    w.WriteHeader(500)
            case CLIENT_REQUEST_ERROR: w.WriteHeader(404)
        }
    } else {
        w.WriteHeader(500);
    }

    w.Write(jsonStr)
}


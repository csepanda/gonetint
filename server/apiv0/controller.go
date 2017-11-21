package apiv0

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

const (
    VERSION = "v0.1"
)

type ApiController struct {
    root *mux.Router
}

func (api *ApiController) GetVersion() string {
    return VERSION
}

func (api *ApiController) Init(router *mux.Router) error {
    api.root = router.PathPrefix("/" + VERSION).Subrouter()
    api.root.Path("/interfaces").HandlerFunc(interfaceList)

    return nil
}

func interfaceList(w http.ResponseWriter, r *http.Request) {
    list, err := getInterfacesNames()
    if err != nil {
        handleServiceError(err, w)
        return
    }


    jsonStr, e := json.Marshal(list)
    if e != nil {
        handleServiceError(serverError(e.Error()), w)
    } else {
        w.Write(jsonStr)
    }
}

func handleServiceError(e *serviceError, w http.ResponseWriter) {
    fmt.Fprintf(w, "{\"error\":\"%s\"}", e.message)
    switch e.errType {
        case SERVER_SIDE_ERROR:    w.WriteHeader(500)
        case CLIENT_REQUEST_ERROR: w.WriteHeader(404)
    }
}

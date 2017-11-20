package apiv0

import (
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

    return nil
}

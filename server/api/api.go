package api

import "github.com/gorilla/mux"

type API interface {
    Init(router *mux.Router) error
    GetVersion() string
}


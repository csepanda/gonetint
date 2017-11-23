package main

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "github.com/csepanda/gonetint/domain/rv0"
)

func fetchList(srvAddress string) (rv0.InterfaceListResponse, error) {
    url := srvAddress + "/" + VERSION + "/interfaces"
    resp, err := http.DefaultClient.Get(url)

    if err != nil {
        return rv0.InterfaceListResponse{}, err
    }

    defer resp.Body.Close()

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)

    if resp.StatusCode != 200 {
        return rv0.InterfaceListResponse{}, errors.New(buf.String())
    }

    var list rv0.InterfaceListResponse
    jsonErr := json.Unmarshal(buf.Bytes(), &list)
    if jsonErr != nil {
        return rv0.InterfaceListResponse{}, jsonErr
    }

    return list, nil
}

func fetchDetails(host string, port int) (rv0.InterfaceResponse, error) {
    return rv0.InterfaceResponse{}, nil
}


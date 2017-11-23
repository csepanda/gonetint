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
        var errResponse rv0.ErrorResponse
        jsonErr := json.Unmarshal(buf.Bytes(), &errResponse)
        if jsonErr != nil {
          return rv0.InterfaceListResponse{}, jsonErr
        } else {
          errMsg := "error: " + errResponse.Error + "\n"
          return rv0.InterfaceListResponse{}, errors.New(errMsg)
        }
    }

    var list rv0.InterfaceListResponse
    jsonErr := json.Unmarshal(buf.Bytes(), &list)
    if jsonErr != nil {
        return rv0.InterfaceListResponse{}, jsonErr
    }

    return list, nil
}

func fetchDetails(srvAddress string, name string) (rv0.InterfaceResponse, error) {
    url := srvAddress + "/" + VERSION + "/interface/" + name
    resp, err := http.DefaultClient.Get(url)

    if err != nil {
        return rv0.InterfaceResponse{}, err
    }

    defer resp.Body.Close()

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)

    if resp.StatusCode != 200 {
        var errResponse rv0.ErrorResponse
        jsonErr := json.Unmarshal(buf.Bytes(), &errResponse)
        if jsonErr != nil {
          return rv0.InterfaceResponse{}, jsonErr
        } else {
          errMsg := "error: " + errResponse.Error + "\n"
          return rv0.InterfaceResponse{}, errors.New(errMsg)
        }
    }

    var info rv0.InterfaceResponse
    jsonErr := json.Unmarshal(buf.Bytes(), &info)
    if jsonErr != nil {
        return rv0.InterfaceResponse{}, jsonErr
    }

    return info, nil
}

package main

import (
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/csepanda/gonetint/server/apiv0"
    "github.com/csepanda/gonetint/domain/rv0"
)

func TestInterfaceListFetch(t *testing.T) {
    currentApi := &apiv0.ApiController{}
    router     := mux.NewRouter().StrictSlash(true)
    currentApi.Init(router)
    srv := httptest.NewServer(router)

    url := srv.URL + "/" + VERSION + "/interfaces"
    expectedResponse, _ := http.DefaultClient.Get(url)
    actual, actualErr := fetchList(srv.URL)

    if expectedResponse.StatusCode != 200 && actualErr != nil {
        t.Fatalf("Expeced status code %d got 200", expectedResponse.StatusCode)
    }

    defer expectedResponse.Body.Close()

    buf := new(bytes.Buffer)
    buf.ReadFrom(expectedResponse.Body)

    var expected rv0.InterfaceListResponse
    jsonErr := json.Unmarshal(buf.Bytes(), &expected)
    if jsonErr != nil {
        t.Fatalf("Couldn't parse json from expected response %#v\n", jsonErr)
    }

    for i := range expected.Interfaces {
        if act, exp := actual.Interfaces[i], expected.Interfaces[i];
           act != exp {
            t.Fatalf("Expected %#v got %#v\n", exp, act)
        }
    }
}

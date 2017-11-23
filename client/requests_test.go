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

    if expectedResponse.StatusCode != 200 && actualErr == nil {
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

func TestInterfaceDetailsFetch(t *testing.T) {
    currentApi := &apiv0.ApiController{}
    router     := mux.NewRouter().StrictSlash(true)
    currentApi.Init(router)
    srv := httptest.NewServer(router)

    list, listErr := fetchList(srv.URL)
    if listErr != nil {
        t.Fatalf("Test cannot be complited due the %#v\n", listErr)
    } else if len(list.Interfaces) == 0 {
        t.Fatal("Test cannot be complited due the lack of interfaces\n")
    }

    for _, name := range list.Interfaces {
        url := srv.URL + "/" + VERSION + "/interface/" + name
        expectedResponse, _ := http.DefaultClient.Get(url)
        actual,   actualErr := fetchDetails(srv.URL, name)

        if expectedResponse.StatusCode != 200 && actualErr == nil {
            t.Fatalf("Expeced status code %d got 200", expectedResponse.StatusCode)
        }

        defer expectedResponse.Body.Close()
        buf := new(bytes.Buffer)
        buf.ReadFrom(expectedResponse.Body)

        var expected rv0.InterfaceResponse
        jsonErr := json.Unmarshal(buf.Bytes(), &expected)
        if jsonErr != nil {
            t.Fatalf("Couldn't parse json '%s' from expected response %#v\n",
                     buf.String(), jsonErr)
        }

        if actual.Name != expected.Name {
            t.Errorf("Expected %s got %s", expected.Name, actual.Name)
        }

        if actual.Hw_address != expected.Hw_address {
            t.Errorf("Expected %s got %s", expected.Hw_address, actual.Hw_address)
        }

        if actual.MTU != expected.MTU {
            t.Errorf("Expected %s got %s", expected.MTU, actual.MTU)
        }

        for i := range expected.Inet_address {
            if act, exp := actual.Inet_address[i], expected.Inet_address[i];
               act != exp {
                t.Errorf("Expected %#v got %#v\n", exp, act)
            }
        }
    }
}

func TestWrongInterfaceDetailsFetch(t *testing.T) {
    currentApi := &apiv0.ApiController{}
    router     := mux.NewRouter().StrictSlash(true)
    currentApi.Init(router)
    srv := httptest.NewServer(router)

    _, actualErr := fetchDetails(srv.URL, "__hope_there_is_no_such_interface")

    if actualErr == nil {
        t.Fatal("Expected no such interface error got nothing\n")
    }
}

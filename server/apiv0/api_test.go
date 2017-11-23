package apiv0

import (
    "bytes"
    "testing"
    "net"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "github.com/gorilla/mux"
)

func TestInterfaceListRequest(t *testing.T) {
    currentApi := &ApiController{}
    router     := mux.NewRouter().StrictSlash(true)
    currentApi.Init(router)
    srv := httptest.NewServer(router)

    url := srv.URL + "/" + VERSION + "/interfaces"
    resp, _ := http.DefaultClient.Get(url)
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        t.Fatalf("Expected status code 200 got %d\n", resp.StatusCode)
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)


    var actual interfaceListResponse
    jsonErr := json.Unmarshal(buf.Bytes(), &actual)
    if jsonErr != nil {
        t.Fatalf("Couldn't parse json from response %#v\n", jsonErr)
    }

    expected, _ := getInterfacesNames()

    for i := range expected.Interfaces {
        if act, exp := actual.Interfaces[i], expected.Interfaces[i];
           act != exp {
            t.Fatalf("Expected %#v got %#v\n", exp, act)
        }
    }
}

func TestInterfaceDetailsRequest(t *testing.T) {
    currentApi := &ApiController{}
    router     := mux.NewRouter().StrictSlash(true)
    currentApi.Init(router)
    srv := httptest.NewServer(router)

    ifis, errLib := net.Interfaces()
    if errLib != nil {
        t.Fatalf("Test cannot be complited due the %#v\n", errLib)
    }

    for _, ifi := range ifis {
        name := ifi.Name
        expected, _ := getInterfaceDetails(name)

        url := srv.URL + "/" + VERSION + "/interface/" + name
        resp, _ := http.DefaultClient.Get(url)
        defer resp.Body.Close()

        if resp.StatusCode != 200 {
            t.Fatalf("Expected status code 200 got %d\n", resp.StatusCode)
        }

        buf := new(bytes.Buffer)
        buf.ReadFrom(resp.Body)

        var actual interfaceResponse;
        jsonErr := json.Unmarshal(buf.Bytes(), &actual)
        if jsonErr != nil {
            t.Fatalf("Couldn't parse json from response %#v\n", jsonErr)
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

func TestWrongInterfaceDetailsRequest(t *testing.T) {
    currentApi := &ApiController{}
    router     := mux.NewRouter().StrictSlash(true)
    currentApi.Init(router)

    srv := httptest.NewServer(router)

    url := srv.URL + "/" + VERSION + "/interface/__hope_there_is_no_such_interface"
    resp, _ := http.DefaultClient.Get(url)
    defer resp.Body.Close()

    if resp.StatusCode != 404 {
        t.Fatalf("Expected status code 404 got %d\n", resp.StatusCode)
    }
}
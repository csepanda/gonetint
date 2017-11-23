package apiv0

import (
    "testing"
    "net"
)

// useless test just for testing
func TestGetInterfacesNames(t *testing.T) {
    expected, errLib := net.Interfaces()
    actualResponse, err := getInterfacesNames()

    checkServiceError(errLib, err, SERVER_SIDE_ERROR, t)

    actual := actualResponse.Interfaces
    for i := range expected {
        if exp, act := expected[i].Name, actual[i]; exp != act {
            t.Errorf("Expect %s got %s\n", exp, act)
        }
    }
}

func TestAllInterfacesDetails(t *testing.T) {
    expectedIfi, errLib := net.Interfaces()
    if errLib != nil {
        t.Fatalf("Test cannot be complited due the %#v\n", errLib)
    }

    for _, ifi := range expectedIfi {
        name := ifi.Name
        details, err := getInterfaceDetails(name)
        expectedAddrs, possibleError := ifi.Addrs()
        checkServiceError(possibleError, err, SERVER_SIDE_ERROR, t)
        for i := range expectedAddrs {
            if exp, act := expectedAddrs[i], details.Inet_address[i];
                exp.String() != act {
                t.Errorf("Expect %s got %s\n", exp, act)
            }
        }
    }
}

func TestWrongInterfaceDetails(t *testing.T) {
    _, err := getInterfaceDetails("")
    checkServiceError(clientError("Wrong interface name"), err,
                      CLIENT_REQUEST_ERROR, t)
}

func checkServiceError(expected, actual error,
    errType serviceErrorType, t *testing.T) {
    if actual != nil {
        if se, ok := actual.(*serviceError); ok {
            if se.errType != errType {
                t.Fatalf("Unexpected error type: %X\n", se.errType)
            }
        } else {
            t.Fatalf("Unexpected error: %#v\n", actual)
        }
    } else if expected != nil {
        t.Fatalf("Unhandled error %#v\n", expected)
    }
}

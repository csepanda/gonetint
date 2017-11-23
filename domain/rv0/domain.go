package rv0
/*  Response domain protocol for API major version 0.
    General domain types to unify server-client communication */

// details of specified interface
type InterfaceResponse struct {
    Name           string
    Hw_address     string
    Inet_address []string
    MTU            int
}

// list of network interfaces list
type InterfaceListResponse struct {
    Interfaces []string
}

type ErrorResponse struct {
    Error string
}

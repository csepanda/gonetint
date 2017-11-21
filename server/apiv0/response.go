package apiv0

// details of specified interface
type interfaceResponse struct {
    Name           string
    Hw_address     string
    Inet_address []string
    MTU            int
}

// list of network interfaces list
type interfaceListResponse struct {
    Interfaces []string
}

package apiv0

// details of specified interface
type interfaceResponse struct {
    name           string
    hw_address     string
    inet_address []string
    MTU            int
}

// list of network interfaces list
type interfaceListResponse struct {
    interfaces []string
}

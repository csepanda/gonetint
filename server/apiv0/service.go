package apiv0

import "net"

func getInterfacesNames() (interfaceListResponse, error) {
    ints, err := net.Interfaces()

    if err != nil {
        return interfaceListResponse{}, serverError(err.Error())
    }

    names := make([]string, len(ints))

    for i, v := range ints {
        names[i] = v.Name
    }

    return interfaceListResponse{names}, nil
}

func getInterfaceDetails(name string) (interfaceResponse, error) {
    ints, err := net.Interfaces()

    if err != nil {
        return interfaceResponse{}, serverError(err.Error())
    }

    var ifi *net.Interface
    for _, i := range ints {
        if i.Name == name {
            ifi = &i
            break
        }
    }

    if ifi == nil {
        message := "interface " + name + " was not found"
        return interfaceResponse{}, clientError(message)
    }

    addrs, err := getInterfaceAddressList(ifi)
    if err != nil {
        return interfaceResponse{}, serverError(err.Error())
    }

    return interfaceResponse{
        Name:         name,
        Hw_address:   ifi.HardwareAddr.String(),
        Inet_address: addrs,
        MTU:          ifi.MTU,
    }, nil
}

func getInterfaceAddressList(ifi *net.Interface) ([]string, error) {
    addrs, err := ifi.Addrs()
    if err != nil {
        return nil, serverError(err.Error())
    }

    addrsList := make([]string, len(addrs))
    for i, v := range addrs {
        addrsList[i] = v.String()
    }

    return addrsList, nil
}

package apiv0

import "net"

func getInterfacesNames() (interfaceListResponse, *serviceError) {
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


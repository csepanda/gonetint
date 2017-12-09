/* This Source Code Form is subject to the terms of the Mozilla
 * Public License, v. 2.0. If a copy of the MPL was not distributed
 * with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright © 2017 Andrey Bova                                       */
package apiv0

import (
    "net"
    "github.com/csepanda/gonetint/domain/rv0"
)

func getInterfacesNames() (rv0.InterfaceListResponse, error) {
    ints, err := net.Interfaces()

    if err != nil {
        return rv0.InterfaceListResponse{}, serverError(err.Error())
    }

    names := make([]string, len(ints))

    for i, v := range ints {
        names[i] = v.Name
    }

    return rv0.InterfaceListResponse{names}, nil
}

func getInterfaceDetails(name string) (rv0.InterfaceResponse, error) {
    ints, err := net.Interfaces()

    if err != nil {
        return rv0.InterfaceResponse{}, serverError(err.Error())
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
        return rv0.InterfaceResponse{}, clientError(message)
    }

    addrs, err := getInterfaceAddressList(ifi)
    if err != nil {
        return rv0.InterfaceResponse{}, serverError(err.Error())
    }

    return rv0.InterfaceResponse{
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

/* This Source Code Form is subject to the terms of the Mozilla
 * Public License, v. 2.0. If a copy of the MPL was not distributed
 * with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright © 2017 Andrey Bova                                       */
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

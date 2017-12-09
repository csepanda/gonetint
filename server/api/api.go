/* This Source Code Form is subject to the terms of the Mozilla
 * Public License, v. 2.0. If a copy of the MPL was not distributed
 * with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright © 2017 Andrey Bova                                       */
package api

import "github.com/gorilla/mux"

type API interface {
    Init(router *mux.Router) error
    GetVersion() string
}


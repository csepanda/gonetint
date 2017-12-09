/* This Source Code Form is subject to the terms of the Mozilla
 * Public License, v. 2.0. If a copy of the MPL was not distributed
 * with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright © 2017 Andrey Bova                                       */
package apiv0

const (
    SERVER_SIDE_ERROR    = 0xBEAF
    CLIENT_REQUEST_ERROR = 0xDEAD
)

type serviceErrorType int
type serviceError     struct {
    errType serviceErrorType
    message string
}

func (e *serviceError) Error() string {
    return e.message
}

func serverError(msg string) error {
    return &serviceError{SERVER_SIDE_ERROR, msg}
}

func clientError(msg string) error {
    return &serviceError{CLIENT_REQUEST_ERROR, msg}
}

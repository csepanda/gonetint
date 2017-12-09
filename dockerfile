# This Source Code Form is subject to the terms of the Mozilla
# Public License, v. 2.0. If a copy of the MPL was not distributed
# with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
# Copyright © 2017 Andrey Bova
FROM golang:1.8

EXPOSE 8080

WORKDIR /go/src/github.com/csepanda/gonetint
COPY . .

WORKDIR /go/src/github.com/csepanda/gonetint/server
RUN go-wrapper download 
RUN go-wrapper install

WORKDIR /go/src/github.com/csepanda/gonetint
ENV PATH ./:$PATH

RUN make

CMD ["./net_server"] 

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

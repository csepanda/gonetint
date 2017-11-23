Client-Server go-based NetworkInterface explorer
===

[![BUILD STATUS](https://travis-ci.org/csepanda/gonetint.svg?branch=master)](https://travis-ci.org/csepanda/gonetint)

> Test assigment to YADRO

# Description
System network interfaces exploration client-server tool.

# Installation
## Plain
```
git clone https://github.com/csepanda/gonetint.git
cd gonetint
make
```
## Docker image
```
docker pull csepanda/gonetint
```
# Usage

## Server

Start server on default port 9000: `$ ./net_server`

Start server on custom port: `$./net_server -port $YOUR_PORT`

### Server on docker
`$ docker run -p 8080:8080 gonetint net_server # -port $YOUR_PORT`

## Client
Fetch interfaces from machine where server runs: `$ ./net_client -server $MACHINE_HOST -port $SERVER_PORT list`

Show details of interface __eth0__: `$ ./net_client -server $MACHINE_HOST -port $SERVER_PORT show eth0`

To show complete help execute: `$ ./net_client -h`

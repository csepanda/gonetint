/* This Source Code Form is subject to the terms of the Mozilla
 * Public License, v. 2.0. If a copy of the MPL was not distributed
 * with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright © 2017 Andrey Bova                                       */
package main

const (
    HELP_STRING = `NAME:
    cli_net - display network interfaces and info. 

USAGE:
    cli_net [-server={host|ip}] [-port=port_num] command [command_args]

VERSION:
    0.1

OPTIONS
    -server=host|ip
        Specifies server to connect (localhost by default)
    -port=port_num
        Specifies port to connect (80 by default)
    -h
        Show this help
    -version
        Print this client version

COMMANDS:
    list
        List all system network interfaces
    show interface_name
        Shows detailed interface info about interface which name
        specified in command

EXIT_STATUS:
     0 - success
     1 - execution error
     2 - no arguments specified
     3 - no command specified
     4 - command usage error
    13 - fatal inner logic error; please report bug
`
    VERSION_STRING = "cli_net " + VERSION

    /* Error strings */
    NO_ARGS_ERROR = "error: no arguments specified\n"
    NO_CMD_ERROR  = "error: nothing to execute\n"
    NO_IFI_ERROR  = "error: no interface name specified\n"
    ACCEPT_NO_ARGS_ERROR = "error: command accept no arguments\n"
    FATAL_ERROR   = "fatal error: inner logic error\n"
)


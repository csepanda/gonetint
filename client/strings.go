package main

const (
    HELP_STRING = `NAME:
    cli_net - display network interfaces and info. 

USAGE:
    cli_net [-server={host|ip}] [-port=port_num] -cmd=command 

VERSION:
    0.1

OPTIONS
    -server=host|ip
        Specifies server to connect (localhost by default)
    -port=port_num
        Specifies port to connect (80 by default)
    -h
        Show this help
    -cmd=command
        Execute specified command. List of commands is below
    -version
        Print this client version

COMMANDS:
    list
        List all system network interfaces
    show
        Shows detailed interface info 
`
    VERSION_STRING = "cli_net " + VERSION
)


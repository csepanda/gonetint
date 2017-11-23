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

EXIT_STATUS:
    0 - success
    1 - execution error
    2 - no arguments specified
    3 - no command specified
`
    VERSION_STRING = "cli_net " + VERSION

    /* Error strings */
    NO_ARGS_ERROR = "Error: no arguments specified\n"
    NO_CMD_ERROR  = "Error: nothing to execute\n"

)


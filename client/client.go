package main

import (
    "os"
    "fmt"
    "flag"
    "strconv"
    "errors"
)

const (
    VERSION       = "v0.1"
    PRINT_VERSION = 0xDEAD
    EXECUTE_CMD   = 0xBEAF
    NO_ARG_FAIL   = -1
    NO_CMD_FAIL   = -2
    CMD_USE_FAIL  = -3
)

type actionType int
type parsedOptions struct {
    action actionType
    cmd    command

    err    error

    server string
    port   int
}

func main() {
    parsed := parseArgs()

    switch {
        case parsed.action == NO_ARG_FAIL:
            printErr(NO_ARGS_ERROR)
            os.Exit(2)
        case parsed.action == NO_CMD_FAIL:
            printErr(NO_CMD_ERROR)
            os.Exit(3)
        case parsed.action == CMD_USE_FAIL:
            printErr(parsed.err.Error())
            os.Exit(4)
        case parsed.action == PRINT_VERSION:
            fmt.Println(VERSION_STRING)
            return
    }

    switch parsed.cmd.GetName() {
        case "list":
            address := "http://" + parsed.server + ":" + strconv.Itoa(parsed.port)
            list, err := fetchList(address)
            if err != nil {
                printErr(err.Error())
                os.Exit(1)
            }

            for _, ifi := range list.Interfaces {
                fmt.Print(ifi, " ")
            }

            fmt.Println("")
        case "show":
            details, err := fetchDetails(parsed.server, parsed.port)
            if err != nil {
                printErr(err.Error())
                os.Exit(1)
            }
            fmt.Println(details)
    }

}

func parseArgs() parsedOptions {
    if len(os.Args) == 1 {
        return parsedOptions{action: NO_ARG_FAIL}
    }

    serverPtr  := flag.String ("server",  "localhost", "")
    portPtr    := flag.Int    ("port",    8080,        "")
    versionPtr := flag.Bool   ("version", false,       "")

    flag.Usage = usage
    flag.Parse()

    if *versionPtr {
        return parsedOptions{ action: PRINT_VERSION }
    }

    args      := os.Args[1:]
    argsCount := len(args)
    var cmd command
    var err error
    for i, arg := range args {
        switch arg {
            case "list":
                if i != argsCount - 1 {
                    return parsedOptions{
                        action: CMD_USE_FAIL,
                        err: errors.New(ACCEPT_NO_ARGS_ERROR),
                    }
                } else {
                    cmd = listCommand{}
                }
            case "show":
                cmd, err = parseShowCommand(args[i + 1:])
                break
        }
    }

    if err != nil {
        return parsedOptions{action: CMD_USE_FAIL, err: err}
    } else if cmd == nil {
        return parsedOptions{ action: NO_CMD_FAIL }
    }

    return parsedOptions{EXECUTE_CMD, cmd, nil, *serverPtr, *portPtr }
}

func parseShowCommand(args []string) (command, error) {
    if len(args) != 1 {
        return nil, errors.New(NO_IFI_ERROR)
    }

    return showCommand{args[0]}, nil
}

func usage() {
    fmt.Println(HELP_STRING)
}

func printErr(message string) {
    os.Stderr.WriteString(message)
}


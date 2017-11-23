package main

import (
    "os"
    "fmt"
    "flag"
    "strconv"
)

const (
    VERSION       = "v0.1"
    PRINT_VERSION = 0xDEAD
    EXECUTE_CMD   = 0xBEAF
    NO_ARG_FAIL   = -1
)

type actionType int
type parsedOptions struct {
    action actionType
    cmd    string

    server string
    port   int
}

func main() {
    parsed := parseArgs()

    switch {
        case parsed.action == NO_ARG_FAIL:
            printErr(NO_ARGS_ERROR)
            os.Exit(2)
        case parsed.action == EXECUTE_CMD && parsed.cmd == "":
            printErr(NO_CMD_ERROR)
            os.Exit(3)
        case parsed.action == PRINT_VERSION:
            fmt.Println(VERSION_STRING)
            return
    }

    switch parsed.cmd {
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

    serverPtr  := flag.String ("server",  "127.1", "")
    portPtr    := flag.Int    ("port",    8080,    "")
    cmdPtr     := flag.String ("cmd",     "",      "")
    versionPtr := flag.Bool   ("version", false,   "")

    flag.Usage = usage
    flag.Parse()

    if *versionPtr {
        return parsedOptions{ action: PRINT_VERSION }
    }

    return parsedOptions{EXECUTE_CMD, *cmdPtr, *serverPtr, *portPtr }
}

func usage() {
    fmt.Println(HELP_STRING)
}

func printErr(message string) {
    os.Stderr.WriteString(message)
}


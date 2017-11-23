package main

import (
    "os"
    "fmt"
    "flag"
)

const (
    VERSION       = "v0.1"
    PRINT_VERSION = 0xDEAD
    EXECUTE_CMD   = 0xBEAF
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

    if parsed.action == PRINT_VERSION {
        fmt.Println(VERSION_STRING)
        return
    }

    if parsed.cmd == "" {
        printErr("No command to execute specified\n")
    }
}

func parseArgs() parsedOptions {
    if len(os.Args) == 1 {
        printErr("Error: no arguments specified\n")
        os.Exit(2)
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


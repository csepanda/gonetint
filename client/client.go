package main

import (
    "os"
    "fmt"
    "net"
    "flag"
    "bytes"
    "errors"
    "strconv"
    "github.com/csepanda/gonetint/domain/rv0"
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
            if showCmd, ok := parsed.cmd.(showCommand); ok {
                address := "http://" + parsed.server + ":" + strconv.Itoa(parsed.port)
                details, err := fetchDetails(address, showCmd.interfaceName)
                if err != nil {
                    printErr(err.Error())
                    os.Exit(1)
                }

                str, err := detailsToString(details)
                if err != nil {
                    printErr(err.Error())
                    os.Exit(1)
                }

                fmt.Println(str)
            } else {
                printErr(FATAL_ERROR)
                os.Exit(13)
            }
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

func detailsToString(ifi rv0.InterfaceResponse) (string, error) {
    var buffer bytes.Buffer

    buffer.WriteString(ifi.Name)
    buffer.WriteString(":\n\tHw_address: ")
    buffer.WriteString(ifi.Hw_address)

    for _, addr := range ifi.Inet_address {
        ip, nw, err := net.ParseCIDR(addr)

        if err != nil {
            return "", err
        }

        if ip.To4() != nil {
            mask := fmt.Sprintf("%d.%d.%d.%d",
                nw.Mask[0], nw.Mask[1], nw.Mask[2], nw.Mask[3])
            buffer.WriteString("\n\tIPv6: ");
            buffer.WriteString(ip.String())
            buffer.WriteString(", mask: ");

            buffer.WriteString(mask)
        } else {
            buffer.WriteString("\n\tIPv4: ");
            buffer.WriteString(addr)
        }
    }

    buffer.WriteString("\n\tMTU: ")
    buffer.WriteString(strconv.Itoa(ifi.MTU))

    return buffer.String(), nil
}

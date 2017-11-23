package main

type command interface {
    GetName() string
}

type listCommand struct { }

func (cmd listCommand) GetName() string {
    return "list"
}

type showCommand struct {
    interfaceName string
}

func (cmd showCommand) GetName() string {
    return "show"
}

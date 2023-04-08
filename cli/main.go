package main

import (
	"github.com/fabiante/nodion-go-cli-example/cli/cmds"
)

func main() {
	if err := cmds.RootCmd.Execute(); err != nil {
		panic(err)
	}
}

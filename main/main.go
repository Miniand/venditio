package main

import (
	"github.com/Miniand/venditio"
	"github.com/Miniand/venditio/cmd"
)

func main() {
	v := venditio.New()
	cmd.Run(v.MustGet(cmd.DEP_COMMANDER).(cmd.Commander))
}

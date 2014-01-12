package main

import (
	"github.com/Miniand/venditio"
)

func main() {
	v := venditio.New()
	v.Get("/", func() string {
		return "moo!"
	})
	v.Run()
}

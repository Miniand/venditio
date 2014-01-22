package main

import (
	"github.com/Miniand/venditio"
	"github.com/Miniand/venditio/web"
)

func main() {
	v := venditio.New()
	web.Run(v)
}

package main

import (
	"flag"
)

type Stat struct {
	cpu int
	disk int
	memory int
}

func main() {
	var t string
	flag.StringVar(&t, "type", "client", "\"server\" or \"client\"")
	flag.Parse()

	switch t {
	case "client":
		break
	case "server":
		break
	}
}

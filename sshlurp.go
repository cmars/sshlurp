package main

import (
	. "github.com/cmars/sshlurp-go.crypto/ssh"
	"fmt"
	"os"
)

func usage(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}

func main() {
	config := &ClientConfig{
		User: "xbmc",
		Auth: []ClientAuth{},
	}
	if len(os.Args) < 2 {
		usage("Usage: <host:port>")
	}
	key, err := DialASlurp("tcp", os.Args[1], config)
	if err != nil {
		usage("Failed to dial: " + err.Error())
	}
	hk, err := ParseHostKey(key)
	if err != nil {
		usage("Failed to parse host key: " + err.Error())
	}
	fmt.Printf("%v\n", hk.Fingerprint())
	fmt.Printf("%v\n", hk.String())
}

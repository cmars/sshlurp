package main

import (
	"encoding/json"
	"fmt"
	. "github.com/cmars/sshlurp-go.crypto/ssh"
	"os"
)

func die(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}

type SlurpKey struct {
	Algo        string `json:",omitempty"`
	Fingerprint string `json:",omitempty"`
	HostKey     string `json:",omitempty"`
	Err         string `json:"err,omitempty"`
}

func main() {
	config := &ClientConfig{}
	if len(os.Args) < 2 {
		die("Usage: <host:port>")
	}
	var results []*SlurpKey
	for _, algo := range SlurpAlgos {
		SetSlurpAlgo(algo)
		key, err := DialASlurp("tcp", os.Args[1], config)
		if err != nil {
			results = append(results, &SlurpKey{Algo: algo, Err: err.Error()})
			continue
		}
		hk, err := ParseHostKey(key)
		if err != nil {
			results = append(results, &SlurpKey{Algo: algo, Err: err.Error()})
			continue
		}
		results = append(results, &SlurpKey{algo, hk.Fingerprint(), hk.KeyString(), ""})
	}
	out, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		die("Failed to format JSON: " + err.Error())
	}
	os.Stdout.Write(out)
}

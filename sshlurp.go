/* This program is free software. It comes without any warranty, to
 * the extent permitted by applicable law. You can redistribute it
 * and/or modify it under the terms of the Do What The Fuck You Want
 * To Public License, Version 2, as published by Sam Hocevar. See
 * http://www.wtfpl.net/ for more details.
 */

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


sshlurp
=======

Retrieve SSH public keys from remote hosts.

Install
-------

	$ go get github.com/cmars/sshlurp
	$ go install github.com/cmars/sshlurp

Usage
-----

	$ sshlurp host:port
	$ sshlurp localhost:22

No authentication necessary. sshlurp disconnects after it gets the key.

Why?
----

* Take inventory of your SSH host keys.
* Create alerts if public keys suddenly change.
* Audit your network for known weak keys, duplicates, other problems.

TODO
----

* Force public key algorithm.
* IP range collection in parallel, because goroutines.
* Extract more server details like version string, supported ciphers, etc.
* Output JSON format.
* Check against SSHFP DNS records.

License
-------
WTFPL


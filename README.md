
sshlurp
=======

Retrieve SSH public keys from remote hosts.

Build & Install
---------------

1. Install Go (see http://golang.org). Add $GOPATH/bin to your $PATH.

2. Download sshlurp sources into your $GOPATH.

	$ go get github.com/cmars/sshlurp

3. Build and install sshlurp binary into your $GOPATH.

	$ go install github.com/cmars/sshlurp

Usage
-----

	$ sshlurp
	Usage: <host:port>
	$ sshlurp somehost:22
	56:a2:99:9d:0b:67:33:88:0f:33:3f:64:68:81:4c:c4
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCm2sVAZuox2WgZYn+a01j/Q5CHsOhFR5CuVuI+AWaDdKAAW15YS+HlM6FVpZBJWL6XO2ev21bZV6pfu8O0FKnTfsA2cvS4IXq5lU3g65zDbObRaBJBjcdIMRcboLc9gWSBjSkF6uC++Kfmw4YanG5SW3e7FpwJdsIrTx8nKEh0omqUW966RxOcGmMs8mFaVn4/8aKM+ZTZGSDDcYNu84VUaS+dAsn45Qw2JXHEpCmgXF9QNQGw/6nJBi2QeWQ6o6OJs/5gl+NE0LY6fUcH0oxBJ2FxBoT3DjnDq4YLUeCvnH48JyDnCH6+ZB/R8apzRTaiEHGgw8kvI1pnqYtloxgX
	$

No authentication necessary. sshlurp disconnects after it gets the key.

Why?
----

* Take inventory of your SSH host keys.
* Create alerts if public keys suddenly change.
* Audit your network for known weak keys, duplicates, other problems.

TODO
----

* Force public key algorithm. Right now, it only seems to grab RSA.
* IP range collection in parallel, because goroutines.
* Extract more server details like version string, supported ciphers, etc.
* Check against SSHFP DNS records, blacklists, etc.
* Output JSON format.

License
-------
WTFPL



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
	[
		{
			"Algo": "ssh-rsa",
			"Fingerprint": "7a:d1:95:b1:e0:f0:95:b4:b6:1e:19:f6:38:c6:30:ad",
			"HostKey": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDVNIcPz4Kw0dJ5BlN5olQDNItgCYHf79nylPhjuibOHsSewyBdCDZD0FBvYYWMDwtv7gWKj6tHQIyRtNQmR3xKCVNDXur6ikWCA/YJgFIbgrdFewBqCpgYZCfIA4m27En7+CvkK8VJa285w26QFf1q3oOOGR+8Cpk57j0mBQ7eiW9grpNxKnwtdeR5CUTiXu6b5mLYBiue9NPlZM/1wz7tF5m6GNixnekLvDoqsrnVDClO0q5eUYSsweBncOrOVpbphAhlOpLPS7LsrU5Rvk+Kemd+PbOmmGEYxcqB2GcbGURsrBSw54Wwk20QGeQSOTcAikvS+tCYR7tKU/Kmerk9"
		},
		{
			"Algo": "ssh-dss",
			"Fingerprint": "aa:84:ab:1f:55:1a:68:9b:90:43:9f:06:14:03:cb:24",
			"HostKey": "AAAAB3NzaC1kc3MAAACBALFMp2WIoogkkBIl1tjXlz8quIxn/QSmPDPxa5HcwL/cSKg60JWisHm8M8S3xGyLuhn3nzmIqVc5raPVQXHNW8AG1zwRerFP+aRedS1t/yhUabOoWKd/tLEi7Ih1ZqQKNXu5AixXLaBZxSlc+yfWQGVSdYqVjrjDXgfe9lSjvqujAAAAFQDpAaY1zJJHtedLeyh7gcby5sqw1QAAAIAwe1b045M0B/iGPQJfhcPjr8SmiwvWGhE1FeTVtzu8flktWQTXrWj76+afx/vUUAoRpdhfzUK7LUOzwX2BaIRroMslvC0fOAKR/ioPAWQpHUO/+hBSwmYzDcC16S/HSvqi2OAaZ/PFZQtv/avi8h0Nq4nMQJ/IXRjaYNA0MKvWTwAAAIB3Lra3uWTdl8yfeYeUANXHeK9hU3MTJP4swZiSjiTHnoN79AKus0lU7nb75RtBRmH5PgEsd8P7D5/ZzLNOhOcESSfDWTo6F0p1izFIalLe3E/sCuGgCNg1673OKh9O/klRCY+h2GooOvuhW93wAU35AS2YZ+0jbhStzQ8p51BkBg=="
		},
		{
			"Algo": "ecdsa-sha2-nistp256",
			"Fingerprint": "96:bb:c8:1f:8d:c5:f8:b3:06:c2:d1:e3:da:69:a3:09",
			"HostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJv+wtO8RpXrej4ZAuCOJGoK+SDkj2lm59CSEybcQ3x9E1xqJeZkIEQTo6dvq22DRCfa7G/W1k0/sz7XO4Qzwcw="
		},
		{
			"Algo": "ecdsa-sha2-nistp384",
			"err": "ssh: no common algorithms"
		},
		{
			"Algo": "ecdsa-sha2-nistp521",
			"err": "ssh: no common algorithms"
		}
	]
	$

No authentication necessary. sshlurp disconnects after it gets the key.

Why?
----

* Take inventory of your SSH host keys.
* Create alerts if public keys suddenly change.
* Audit your network for known weak keys, duplicates, other problems.

TODO
----

* IP range collection in parallel, because goroutines.
* Extract more SSH server info like version string, supported ciphers, etc.
* Check against SSHFP DNS records, blacklists, etc.

License
-------
WTFPL


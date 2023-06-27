package main

import "flag"

var (
	listenAddr = flag.String("listen", ":4000", "Address to listen on")
)

package main

import (
	"flag"
)

type cmdArgs struct {
	host     string
	client   string
	interval int64
	timeout  int64
}

func setupFlags() *cmdArgs {
	args := new(cmdArgs)

	flag.StringVar(&args.host, "d", "google.com", "name of the domain")
	flag.StringVar(&args.client, "c", "1.1.1.1", "ip of the dns client")
	flag.Int64Var(&args.interval, "i", 1, "request interval in seconds")
	flag.Int64Var(&args.timeout, "t", 1000, "request timeout in milliseconds")

	flag.Parse()

	return args
}

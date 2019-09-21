package main

import (
	"fmt"
	"net"
	"os"

	"github.com/fiskeben/resolv"
)

func main() {
	var lookupURL string
	if len(os.Args) > 1 {
		lookupURL = os.Args[1]
	} else {
		lookupURL = "www.example.com"
	}

	ips, err := net.LookupIP(lookupURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}

	r, err := resolv.Config()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get DNS: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Server:         ", r.Nameservers[0])
	fmt.Println("Address:         " + r.Nameservers[0] + "#53")
	fmt.Print("\n")
	fmt.Println("Non-authoritative answer:")
	fmt.Println("Name:   ", lookupURL)
	fmt.Println("Address: ", ips[0])
	fmt.Print("\n")
}

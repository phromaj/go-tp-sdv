package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// Implémentation avec fmt.Sprintf()
func ipAddrToStringWithFmt(ip IPAddr) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Implémentation avec strconv
func ipAddrToStringWithStrconv(ip IPAddr) string {
	var ipString string
	for i, octet := range ip {
		ipString += strconv.Itoa(int(octet))
		if i < len(ip)-1 {
			ipString += "."
		}
	}
	return ipString
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ipAddrToStringWithFmt(ip))
		fmt.Printf("%v: %v\n", name, ipAddrToStringWithStrconv(ip))
	}
}

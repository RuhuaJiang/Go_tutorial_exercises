package main

import "fmt"
import "strings"
import "strconv"

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	var straddr []string
	for _, v := range ipaddr {
		straddr = append(straddr, strconv.Itoa(int(v)))
	}
	return strings.Join(straddr, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

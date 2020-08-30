package main

import (
	"fmt"
	"log"

	"github.com/Juniper/go-netconf/netconf"
)

func main() {
	fmt.Println("New netconf client")
	netconf.SetLog(netconf.NoopLog{})
	sesion, err := netconf.DialSSH("10.121.40.14", netconf.SSHConfigPassword("diag", "ciena123"))
	if err != nil {
		log.Fatal(err)
	}

	for _, cap := range sesion.ServerCapabilities {
		fmt.Println(cap)
	}
	fmt.Println(sesion.SessionID)
	sesion.Close()
}

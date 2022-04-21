package c3

import (
	"fmt"
	"log"
	"net"
)

func FindIP() {
	//if len(os.Args) != 2 {
	//	log.Fatal("No IP address argument provided.")
	//}
	arg := "120.230.105.35"
	ip := net.ParseIP(arg)
	if ip == nil {
		log.Fatal("Valid IP not detected. Value provided: " + arg)

	}
	hostnames, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range hostnames {
		fmt.Println(item)
	}
}

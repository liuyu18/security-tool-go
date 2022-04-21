package c5

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func FindAllDevices() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range devices {
		fmt.Println("name: ", device.Name)
		fmt.Println("des： ", device.Description)
		fmt.Println("addr: ", device.Addresses)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}

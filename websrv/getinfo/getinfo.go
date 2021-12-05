package getinfo

import (
	"fmt"
	"net"

	"os"
)

func Gethostname() string {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Get hostname Error: %v\n", err)
	}

	return name
}

func GetIpAddress() string {

	var ipaddress string
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("Get IPaddress Error: %v\n", err)
	}

	for _, values := range address {
		if ipnet, ok := values.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipaddress = ipnet.IP.String()
			}
		}
	}
	return ipaddress
}

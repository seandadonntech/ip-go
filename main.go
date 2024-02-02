package main

import (
	"fmt"
	"net"
	
)

func main() {
	var options string
	fmt.Print("ENTER OPTIONS: ")
	fmt.Scan(&options)

	if options == "1" {
		// Get local IP addresses
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				fmt.Println(ipNet.IP)
			}
		}
	} else {
		fmt.Println("Invalid command")
		main() // Exit the program with an error code
	}
}



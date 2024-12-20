package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func LoopNetwork() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		// handle err
		for _, addr := range addrs {
			fmt.Println(addr)
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println(ip)
			fmt.Println()
		}
	}
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	fmt.Println("Hello")
	fmt.Println(len(os.Args), os.Args)
	// if len(os.Args) <= 1 {
	// 	fmt.Println("Need test name")
	// 	return
	// }
	LoopNetwork()

	fmt.Println(GetOutboundIP())

	test_fibonacci()
}

package main

import (
	"fmt"
	"github.com/ugwis/dhcp4client"
	"log"
	"net"
)

func main() {
	m, err := net.ParseMAC("f4:0f:24:3c:6f:4c")
	if err != nil {
		log.Printf("MAC Error:%v\n", err)
	}

	c, err := dhcp4client.NewInetSock(dhcp4client.SetLocalAddr(net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 68}), dhcp4client.SetRemoteAddr(net.UDPAddr{IP: net.IPv4bcast, Port: 67}))
	if err != nil {
		log.Fatalf("Client Connection Generation:" + err.Error())
	}
	defer c.Close()

	exampleClient, err := dhcp4client.New(dhcp4client.HardwareAddr(m), dhcp4client.Connection(c))
	if err != nil {
		log.Fatalf("Error:%v\n", err)
	}
	defer exampleClient.Close()
	servers, err := exampleClient.Request()
	if err != nil {
		log.Fatalf("Client request failed: " + err.Error())
	}
	for _, server := range servers {
		fmt.Println(server)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type clientType map[string]bool

var clients = clientType{}

func (c clientType) keys(filter string) string {
	var output []string
	for key := range c {
		if key != filter {
			output = append(output, key)
		}
	}

	return strings.Join(output, ",")
}

// Server --
func Server() {
	localAddress := ":4242"
	if len(os.Args) > 2 {
		localAddress = os.Args[2]
	}

	addr, _ := net.ResolveUDPAddr("udp", localAddress)
	conn, _ := net.ListenUDP("udp", addr)

	for {
		buffer := make([]byte, 1024)
		bytesRead, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		incoming := string(buffer[0:bytesRead])
		fmt.Printf("[INCOMING] %s from %s", incoming, remoteAddr)
		if incoming != "register" {
			continue
		}

		clients[remoteAddr.String()] = true

		for client := range clients {
			resp := clients.keys(client)
			if len(resp) > 0 {
				r, _ := net.ResolveUDPAddr("udp", client)
				_, _ = conn.WriteTo([]byte(resp), r)
				fmt.Printf("[INFO] Responded to %s with %s\n", client, string(resp))
			}
		}
	}
}

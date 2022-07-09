package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Client --
func Client() {
	registerClient()
}

func registerClient() {
	signalAddress := os.Args[2]

	localAddress := ":50001" // default port
	if len(os.Args) > 3 {
		localAddress = os.Args[3]
	}

	remote, _ := net.ResolveUDPAddr("udp", signalAddress)
	local, _ := net.ResolveUDPAddr("udp", localAddress)
	conn, _ := net.ListenUDP("udp", local)
	go func() {
		bytesWritten, err := conn.WriteTo([]byte("register"), remote)
		if err != nil {
			panic(err)
		}

		fmt.Println(bytesWritten, " bytes written")
	}()

	listenToClient(conn, local.String())
}

func listenToClient(conn *net.UDPConn, local string) {
	for {
		fmt.Println("listening")
		buffer := make([]byte, 1024)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("[ERROR]", err)
			continue
		}

		fmt.Println("[INCOMING]", string(buffer[0:bytesRead]))
		if string(buffer[0:bytesRead]) == "Hello!" {
			continue
		}

		for _, a := range strings.Split(string(buffer[0:bytesRead]), ",") {
			if a != local {
				go handleUDPOnClients(conn, a)
			}
		}
	}
}

func handleUDPOnClients(conn *net.UDPConn, service string) {
	addr, _ := net.ResolveUDPAddr("udp", service)
	for {
		chat(conn, addr)
	}
}

func chat(conn *net.UDPConn, addr net.Addr) {
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("Error reading ", message)
		return
	}
	_, err = conn.WriteTo([]byte(message), addr)
	if err != nil {
		log.Fatalln("Error writing ", message)
		return
	}
}

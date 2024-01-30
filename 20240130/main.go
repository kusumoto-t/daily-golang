package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":53")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer udpConn.Close()

	fmt.Println("DNS server is listening on :53")

	for {
		buffer := make([]byte, 512)
		n, addr, err := udpConn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		go handleDNSQuery(udpConn, addr, buffer[:n])
	}
}

func handleDNSQuery(conn *net.UDPConn, addr *net.UDPAddr, query []byte) {
	var dnsHeader struct {
		ID      uint16
		Flags   uint16
		QDCount uint16
		ANCount uint16
		NSCount uint16
		ARCount uint16
	}

	// ヘッダーをパース
	dnsHeader.ID = binary.BigEndian.Uint16(query[0:2])
	dnsHeader.Flags = binary.BigEndian.Uint16(query[2:4])
	dnsHeader.QDCount = binary.BigEndian.Uint16(query[4:6])
	dnsHeader.ANCount = binary.BigEndian.Uint16(query[6:8])
	dnsHeader.NSCount = binary.BigEndian.Uint16(query[8:10])
	dnsHeader.ARCount = binary.BigEndian.Uint16(query[10:12])

	qnameStart := 12
	qnameEnd := qnameStart
	for query[qnameEnd] != 0 {
		qnameEnd += int(query[qnameEnd]) + 1
	}
}

package main

import (
	"encoding/binary"
	"fmt"
)

// DNSパケットパーサー
func main() {
	// 仮のDNSパケットデータ
	dnsPacketData := []byte{
		0x12, 0x34, // Transaction ID
		0x01, 0x00, // Flags
		0x00, 0x01, // Question Count
		0x00, 0x00, // Answer Record Count
		0x00, 0x00, // Authority Record Count
		0x00, 0x00, // Additional Record Count
		0x03, 0x77, 0x77, 0x77, // QNAME: www
		0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, // QNAME: google
		0x03, 0x63, 0x6f, 0x6d, // QNAME: com
		0x00,       // Null-terminator for QNAME
		0x00, 0x01, // QTYPE: A
		0x00, 0x01, // QCLASS: IN
	}

	var dnsHeader struct {
		ID      uint16
		Flags   uint16
		QDCount uint16
		ANCount uint16
		NSCount uint16
		ARCount uint16
	}

	// ヘッダーをパース
	dnsHeader.ID = binary.BigEndian.Uint16(dnsPacketData[0:2])
	dnsHeader.Flags = binary.BigEndian.Uint16(dnsPacketData[2:4])
	dnsHeader.QDCount = binary.BigEndian.Uint16(dnsPacketData[4:6])
	dnsHeader.ANCount = binary.BigEndian.Uint16(dnsPacketData[6:8])
	dnsHeader.NSCount = binary.BigEndian.Uint16(dnsPacketData[8:10])
	dnsHeader.ARCount = binary.BigEndian.Uint16(dnsPacketData[10:12])

	fmt.Printf("Transaction ID: %x\n", dnsHeader.ID)
	fmt.Printf("Flags: %x\n", dnsHeader.Flags)
	fmt.Printf("Question Count: %d\n", dnsHeader.QDCount)
	fmt.Printf("Answer Record Count: %d\n", dnsHeader.ANCount)
	fmt.Printf("Authority Record Count: %d\n", dnsHeader.NSCount)
	fmt.Printf("Additional Record Count: %d\n", dnsHeader.ARCount)
}

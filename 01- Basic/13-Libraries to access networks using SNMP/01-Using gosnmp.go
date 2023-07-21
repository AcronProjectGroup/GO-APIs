package main

import (
	"fmt"
	snmp "github.com/gosnmp/gosnmp"
	"log"
)

func main() {
	snmp.Default.Target = "10.0.4.1"
	if err := snmp.Default.Connect(); err != nil {
		log.Fatalf("Failed Connect: %v", err)
	}
	defer snmp.Default.Conn.Close()
	// SNMPv2-MIB::sysUpTime
	oid := []string{"1.3.6.1.2.1.1.3.0"}
	result, err := snmp.Default.Get(oid)
	if err != nil {
		log.Fatalf("Failed Get: %v", err)
	}
	for _, variable := range result.Variables {
		fmt.Printf("oid: %s ", variable.Name)
		fmt.Printf(": %d\n", snmp.ToBigInt(variable.Value))
	}
}

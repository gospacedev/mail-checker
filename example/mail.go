package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	sci := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	for sci.Scan(){
		checkDom(sci.Text())
	}

	if err := sci.Err(); err != nil {
		log.Fatal("Error: Can't read from input: \n", err)
	}
}

func checkDom(domain string)  {
	var hasMX, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf"){
			hasSPF = true
			sprRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord)
}

package main

import (
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"
)

/*
CheckDomainMX takes a domain name as a string, and returns email information
containing the domain name, whether the domain has MX records, whether the domain
has an SPF record, whether the domain has a DMARC record, the SPF record, and the
DMARC record to a specified config file.
*/
func CheckDomainMX(domain string, fileName string, fileType string, filePath string) {
	var HasMX, HasSPF, HasDMARC bool
	var SPRRecord, DMARCRecord string

	vp := viper.New()

	// setup config file
	vp.SetConfigName(fileName)
	vp.SetConfigType(fileType)
	vp.AddConfigPath(filePath)

	err := vp.ReadInConfig()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		HasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf") {
			HasSPF = true
			SPRRecord = record
			break
		}
	}

	DMARCRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range DMARCRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			HasDMARC = true
			DMARCRecord = record
			break
		}
	}

	vp.Set("domain", domain)
	vp.Set("HasMX", HasMX)
	vp.Set("HasSPF", HasSPF)
	vp.Set("HasDMARC", HasDMARC)
	vp.Set("SPRecord", SPRRecord)
	vp.Set("DMARCRecord", DMARCRecord)
	vp.WriteConfig()
}

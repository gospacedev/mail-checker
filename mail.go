package mail

import (
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"
)

// `DomMailInfo` is a struct with fields `Domain`, `HasMX`, `HasSPF`, `HasDMARC`, `SPRRecord`, and
// `DMARCRecord`.
// @property {string} Domain - The domain name
// @property {bool} HasMX - Does the domain have an MX record?
// @property {bool} HasSPF - Does the domain have an SPF record?
// @property {bool} HasDMARC - Does the domain have a DMARC record?
// @property {string} SPRRecord - The SPF record for the domain
// @property {string} DMARCRecord - The DMARC record for the domain.
type DomMailInfo struct {
	Domain      string `json:"Domain"`
	HasMX       bool   `json:"HasMX"`
	HasSPF      bool   `json:"HasSPF"`
	HasDMARC    bool   `json:"HasDMARC"`
	SPRRecord   string `json:"SPRRecord"`
	DMARCRecord string `json:"DMARCRecord"`
}

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
		log.Fatal(err)
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

	// parse info in config file
	vp.Set("domain", domain)
	vp.Set("HasMX", HasMX)
	vp.Set("HasSPF", HasSPF)
	vp.Set("HasDMARC", HasDMARC)
	vp.Set("SPRecord", SPRRecord)
	vp.Set("DMARCRecord", DMARCRecord)

	vp.WriteConfig()
}

# mail-checker
[![MIT Licence](https://img.shields.io/badge/license-MIT-blue)](https://opensource.org/licenses/mit-license.php)
[![Go Reference](https://pkg.go.dev/badge/github.com/gospacedev/mail-checker.svg)](https://pkg.go.dev/github.com/gospacedev/mail-checker)

Mail Checker extracts a domain's email DMARC and SPF records. 

## Usage
```
go get github.com/gospacedev/mail-checker
```

Example code:

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sci := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	for sci.Scan(){
		fmt.Println(CheckDomainMX(sci.Text()))
	}
	

	if err := sci.Err(); err != nil {
		log.Fatal("Error: Can't read from input: \n", err)
	}
}
```

Run the file and enter a domain name:

    Enter a domain name:
    google.com
    
Then it should show the email information:

    {
     "Domain": "google.com",
     "HasMX": true,
     "HasSPF": false,
     "HasDMARC": true,
     "SPRRecord": "",
     "DMARCRecord": "v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com"
     }

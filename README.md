# mail-checker

It checks the domain if it has a DMARC, 
a SPF Record, and a Mail Server

## Usage
```
go get github.com/gocrazygh/mail-checker
```

This is an example code:

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/gocrazygh/mail-checker"
)

func main() {
	sci := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	for sci.Scan(){
		mail.CheckDom(sci.Text())
	}

	if err := sci.Err(); err != nil {
		log.Fatal("Error: Can't read from input: \n", err)
	}
}
```

You can simply run it as:

    go run mail.go

This should show:

    domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord

Enter a domain:

    domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord
    mailchimp.com
    
Then it should show the values:

    domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord
    mailchimp.com
    mailchimp.com, true, false, , true, v=DMARC1; p=reject; rua=mailto:19ezfriw@ag.dmarcian.com; ruf=mailto:19ezfriw@fr.dmarcian.com
    
## Buy Me a Coffee
BTC: `1F5qqrV9bX8Z1eyvy6MBxyVCKnT8cc4Hpc`

## Acknowledgement

Akhil Sharma https://www.youtube.com/c/AkhilSharmaTech

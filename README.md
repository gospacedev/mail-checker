# mail-checker
[![MIT Licence](https://img.shields.io/badge/License-MIT-blue)](https://opensource.org/licenses/mit-license.php)

This checks a domain's email information, if it has DMARC, SPF Record, and or a Mail Server

Donate Bitcoin: `122K2s1DL1W75Y8ZAh4cEsfhRCAgeJDaVX`

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

# mail-checker
[![MIT Licence](https://img.shields.io/badge/license-MIT-blue)](https://opensource.org/licenses/mit-license.php)
[![Go Reference](https://pkg.go.dev/badge/github.com/gospacedev/mail-checker.svg)](https://pkg.go.dev/github.com/gospacedev/mail-checker)
[![Go Report Card](https://goreportcard.com/badge/github.com/gospacedev/mail-checker)](https://goreportcard.com/report/github.com/gospacedev/mail-checker)

Mail Checker extracts a domain's email DMARC and SPF records. 

## Usage
```
go get github.com/gospacedev/mail-checker
```

Example code:

```go
package main

import "github.com/gospacedev/mail-checker"

func main() {
	mail.CheckDomainMX("google.com", "config", "json", ".")
}

```
    
The mail information is outputed to the config file, Mail Checker 
can write to JSON, TOML, and YAML config files.

```json
{
  "dmarcrecord": "v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com",
  "domain": "google.com",
  "hasdmarc": true,
  "hasmx": true,
  "hasspf": false,
  "sprecord": ""
}
```

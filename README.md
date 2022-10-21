# mail-checker
[![MIT Licence](https://img.shields.io/badge/license-MIT-blue)](https://opensource.org/licenses/mit-license.php)
[![Go Reference](https://pkg.go.dev/badge/github.com/gospacedev/mail-checker.svg)](https://pkg.go.dev/github.com/gospacedev/mail-checker)
[![Go Report Card](https://goreportcard.com/badge/github.com/gospacedev/mail-checker)](https://goreportcard.com/report/github.com/gospacedev/mail-checker)

Mail Checker extracts a domain's DMARC and SPF Record.

## Usage
```
go get github.com/gospacedev/mail-checker
```

Mail Checker takes in the targeted domain and the config file info and returns 
the domain's mail information to the config file:

```go
package main

import "github.com/gospacedev/mail-checker"

func main() {
	mail.CheckDomainMX("google.com", "config", "json", ".")
}

```
    
The mail information is outputted to the config file, Mail Checker 
supports writing to JSON, TOML, and YAML config files:

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

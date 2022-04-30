# mail-checker

It checks the domain if it has a DMARC, 
a SPF Record, and a Mail Server

## Usage
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
    
## Acknowledgement

Akhil Sharma https://www.youtube.com/c/AkhilSharmaTech

### For Coffee
[![DOGE Donate](https://img.shields.io/badge/DOGE-For%20Coffee-green)](https://dogechain.info/address/DSr9btWDuDcdSg4yXkSMYjbRMLEUqp4ijt)`DSr9btWDuDcdSg4yXkSMYjbRMLEUqp4ijt`

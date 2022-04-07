# Mail Checker
> An mail checker

## How does it work
It scans the inputted domain and checks if it has
DMARC, SPF Record, and a Mail Server

## Usage
You can simply run it as:
```powershell
go run mail.go
```

This sould show
```powershell
domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord
```

Enter a domain
```powershell
domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord
mailchimp.com
```

Then it should show the values
```powershell
domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord
mailchimp.com
mailchimp.com, true, false, , true, v=DMARC1; p=reject; rua=mailto:19ezfriw@ag.dmarcian.com; ruf=mailto:19ezfriw@fr.dmarcian.com
```

## Credits

Credits to Akhil Sharma (https://www.youtube.com/c/AkhilSharmaTech)

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

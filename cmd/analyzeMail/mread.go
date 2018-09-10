package main

import (
	"github.com/mchirico/go_read/grab"
)

func main() {
	grab.FileParse("/var/log/mail.log", "/var/log/email.db")
}

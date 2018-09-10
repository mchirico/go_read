package main

import (
	"fmt"
	"github.com/mchirico/go_read/web"
	"os"
)

func main() {

	a := web.App{}

	fmt.Printf("host: %v\n",os.Getenv("host"))

	a.Initialize(
		"./junk2.db")


	a.Run("8080")
}


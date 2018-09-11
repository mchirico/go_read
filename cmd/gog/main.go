package main

import (
	"github.com/mchirico/go_read/pkg"
)

func main() {

	a := pkg.App{}
	a.Initilize()
	a.Run("8080")

}

package grab

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {

	m := FileParse("../mockfiles/mail.log")
	fmt.Println(m)

}

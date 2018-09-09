package grab

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {

	expected := map[string]int{"<hemanth@codestak.com>": 1,
		"<vishal@ibusinesssolution.com>": 1,
		"<ducky@cwstat.com>":             1,
		"<ducky@aipiggybot.io>":          1,
		"<sam@witssolutions.com>":        2,
		"<ducky@gmail.com>":              1}

	m := FileParse("../mockfiles/mail.log")
	fmt.Println(m)

	if !reflect.DeepEqual(expected, m) {
		t.Fatalf("Results not equal")
	}

}

package grab

import (
	"fmt"
	"os"
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

	dbFile := "./junk.db"
	os.Remove(dbFile)
	m := FileParse("../mockfiles/mail.log", dbFile)
	fmt.Println(m)

	if !reflect.DeepEqual(expected, m) {
		t.Fatalf("Results not equal")
	}

}

package grab

import (
	"fmt"
	"github.com/mchirico/go_read/sqlite"
	"log"
	"os"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func checkCount(dbFile string, stmt string) int {

	var count int

	sq := &sqlite.SQL{}

	sq.OpenDB(dbFile)
	defer sq.Close()

	rows := sq.Query(stmt)
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Count:", count)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return count
}

func TestRead(t *testing.T) {

	expected := map[string]int{"from=<gaurav.jha@acsicorp.com>": 1,
		"from=<sam@witssolutions.com>":             2,
		"from=<pushkar.prabhudesai@collabera.com>": 1,
		"from=<support@aipiggybot.io>":             1,
		"from=<sushant.singh@okayainc.com>":        1,
		"from=<ducky@gmail.com>":                   1,
		"from=<shishir.shekhar@intelliswift.com>":  1,
		"from=<hemanth@codestak.com>":              1,
		"from=<raj.a@vdartinc.com>":                1,
		"from=<pankaj.kumar@okayainc.com>":         2,
		"from=<info@aipiggybot.io>":                1,
		"from=<devesh.mishra@nlbservices.com>":     1,
		"from=<afzal.pakkal@panzersolutions.com>":  1,
		"from=<spameri@tiscali.it>":                5,
		"from=<praveen@sagebeans.com>":             1,
		"from=<lalit.yadav@okayainc.com>":          1,
		"from=<nkps@cwg.com>":                      1}

	dbFile := "./junk.db"
	os.Remove(dbFile)
	m := FileParse("../mockfiles/mail.log", dbFile)
	fmt.Println(m)

	if !reflect.DeepEqual(expected, m) {
		t.Fatalf("Results not equal: %v\n", m)
	}

	count := checkCount(dbFile, "select count(*) as count from mail;")
	if count != 23 {
		t.Fatalf("Count is off: %v\n", count)
	}

}

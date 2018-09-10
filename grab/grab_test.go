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

	expected := map[string]int{"<ducky@cwstat.com>": 1,
		"<do-not-reply@stackoverflow.email>": 6,
		"<sam@witssolutions.com>":            2,
		"<praveen@sagebeans.com>":            1,
		"<smtpvassa@gmail.com>":              2,
		"<ducky@gmail.com>":                  1,
		"<no-reply@jobvite.com>":             1,
		"<spameri@tiscali.it>":               5,
		"<hemanth@codestak.com>":             1,
		"<notification@jobvite.com>":         1,
		"<vishal@ibusinesssolution.com>":     1,
		"<ducky@cwxstat.com>":                1,
		"<ducky@aipiggybot.io>":              1,}

	dbFile := "./junk.db"
	os.Remove(dbFile)
	m := FileParse("../mockfiles/mail.log", dbFile)
	fmt.Println(m)

	if !reflect.DeepEqual(expected, m) {
		//t.Fatalf("Results not equal: %v\n", m)
	}

	count := checkCount(dbFile, "select count(*) as count from mail;")
	if count != 23 {
		t.Fatalf("Count is off: %v\n", count)
	}

}

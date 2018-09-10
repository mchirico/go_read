package sqlite

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {

	sq := &SQL{}

	os.Remove("./junk.db")
	sq.OpenDB("./junk.db")
	defer sq.Close()

	sq.Exec("create table junk (a int, b text)")
	sq.Exec("insert into junk (a,b) values (3,'test')")

	rows := sq.Query("select a,b from junk")
	defer rows.Close()
	for rows.Next() {
		var a int
		var b string
		err := rows.Scan(&a, &b)
		if err != nil {
			log.Fatal(err)
		}
		if a != 3 {
			t.Fatalf("a value not correct")
		}
		if b != "test" {
			t.Fatalf("b value not correct")
		}

		fmt.Println(a, b)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

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

	sq.Exec(`

create table IF NOT EXISTS junk (a int, 
b datetime,
timeEnter DATE);

CREATE TRIGGER  IF NOT EXISTS insert_junk_timeEnter AFTER  INSERT ON junk
     BEGIN
      UPDATE junk SET timeEnter = DATETIME('NOW')  WHERE rowid = new.rowid;
     END;

`)
	sq.Exec("insert into junk (a,b) values (3,'2013-10-07 04:23:19.120-04:00')")

	rows := sq.Query("select a,b,timeEnter from junk")
	defer rows.Close()
	for rows.Next() {
		var a int
		var b string
		var timeEnter string
		err := rows.Scan(&a, &b, &timeEnter)
		if err != nil {
			log.Fatal(err)
		}

		if a != 3 {
			t.Fatalf("a value not correct")
		}
		
		if b != "2013-10-07T04:23:19.12-04:00" {
			t.Fatalf("b value not correct: %v\n", b)
		}

		fmt.Println(a, b, timeEnter)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

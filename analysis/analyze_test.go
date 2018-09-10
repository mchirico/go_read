package analysis

import (
	"fmt"
	"github.com/mchirico/date/parse"
	"log"
	"os"
	"testing"

	"github.com/mchirico/go_read/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func TestDateparse(t *testing.T) {
	s := "Sep 10 19:13:10"
	tt, _ := parse.DateTimeParse(s).NewYork()

	email_time := tt.Format("2006-01-02 15:04:05-07:00")

	if email_time != "2018-09-10 19:13:10-04:00" {
		t.Fatalf("Time not working: %v\n", email_time)
	}
	fmt.Printf("%v\n", email_time)
	//2013-10-07 04:23:19.120-04:00

}

func Setup() int {

	count := 0

	sq := &sqlite.SQL{}

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


create table IF NOT EXISTS junk2 (a int, 
b datetime,
timeEnter DATE);

CREATE TRIGGER  IF NOT EXISTS insert_junk2_timeEnter AFTER  INSERT ON junk2
     BEGIN
      UPDATE junk2 SET timeEnter = DATETIME('NOW')  WHERE rowid = new.rowid;
     END;



`)
	sq.Exec("insert into junk (a,b) values (3,'2013-10-07 04:23:19.120-04:00')")
	sq.Exec("insert into junk (a,b) values (4,'2013-10-07 04:23:19.121-04:00')")

	sq.Exec("insert into junk2 (a,b) values (3,'2013-10-07 04:23:19.120-04:00')")
	sq.Exec("insert into junk2 (a,b) values (4,'2013-10-07 04:23:19.123-04:00')")

	sq.Exec(`

insert into junk (a,b)
select j2.a,j2.b from junk2 as j2 left outer join junk as j on
(j.a=j2.a and j.b=j2.b) where j.a is null

`)

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

		count += 1
		fmt.Println(a, b, timeEnter)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return count
}

func TestInsert(t *testing.T) {

	if Setup() != 3 {
		t.Fatalf("Count should equal 3: %v\n", Setup())
	}

}

func TestInsertData(t *testing.T) {

	file := "./junk.db"
	os.Remove(file)

	m := [][]string{}

	s := "Sep 10 19:13:10"
	tt, _ := parse.DateTimeParse(s).NewYork()
	email_time := tt.Format("2006-01-02 15:04:05-07:00")

	m = append(m, []string{"reject", "this is' onetext ",
		"<donkey@example.com>", email_time})

	m = append(m, []string{"discard", "this is one text ",
		"<donkey@example.com>", email_time})

	InsertData(file, m)

}

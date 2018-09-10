package analysis

import (
	"fmt"
	"github.com/mchirico/go_read/sqlite"
	"log"
	"strings"
)


func InsertData(file string, records [][]string) int {

	count := 0

	sq := &sqlite.SQL{}

	//os.Remove(file)
	sq.OpenDB(file)
	defer sq.Close()

	sq.Exec(`

create table IF NOT EXISTS mail (
tag text,
full text,
email text,
date datetime,
timeEnter DATE);

CREATE TRIGGER  IF NOT EXISTS insert_mail_timeEnter AFTER  INSERT ON mail
     BEGIN
      UPDATE mail SET timeEnter = DATETIME('NOW')  WHERE rowid = new.rowid;
     END;


create table IF NOT EXISTS mail2 (
tag text,
full text,
email text,
date datetime,
timeEnter DATE);


CREATE TRIGGER  IF NOT EXISTS insert_mail2_timeEnter AFTER  INSERT ON mail2
     BEGIN
      UPDATE mail2 SET timeEnter = DATETIME('NOW')  WHERE rowid = new.rowid;
     END;



`)

	for _, v := range records {

		full := strings.Replace(v[1], "'", "`", -1)
		stmt := fmt.Sprintf("insert into mail2 (tag,full,email,date) "+
			"values ('%v','%v','%v','%v')", v[0], full, v[2], v[3])
		sq.Exec(stmt)
	}

	sq.Exec(`

insert into mail (tag,full,email,date)
select j2.tag,j2.full,j2.email,j2.date from mail2 as j2 left outer join mail as j on
(j.tag=j2.tag and j.full=j2.full and j.email=j2.email and j.date=j2.date) where j.email is null

`)

	rows := sq.Query("select tag,full,email,date,timeEnter from mail")
	defer rows.Close()
	for rows.Next() {
		var tag string
		var full string
		var email string
		var date string
		var timeEnter string
		err := rows.Scan(&tag, &full, &email, &date, &timeEnter)
		if err != nil {
			log.Fatal(err)
		}

		count += 1
		fmt.Println(tag, full, email, date, timeEnter)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return count
}

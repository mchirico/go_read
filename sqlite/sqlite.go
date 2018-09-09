package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

type SQL struct {
	db  *sql.DB
	err error
	sync.Mutex
}

func (sq *SQL) OpenDB(file string) {
	sq.Mutex.Lock()
	defer sq.Mutex.Unlock()

	sq.db, sq.err = sql.Open("sqlite3", file)
	if sq.err != nil {
		log.Fatal(sq.err)
	}
}

func (sq *SQL) Exec(stmt string) {
	sq.Mutex.Lock()
	defer sq.Mutex.Unlock()

	_, sq.err = sq.db.Exec(stmt)
	if sq.err != nil {
		log.Printf("%q: %s\n", sq.err, stmt)
		return
	}
}

func (sq *SQL) Query(stmt string) *sql.Rows {
	sq.Mutex.Lock()
	defer sq.Mutex.Unlock()

	rows, err := sq.db.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}
	return rows

}

func (sq *SQL) Close() {
	sq.Mutex.Lock()
	defer sq.Mutex.Unlock()

	sq.db.Close()
}

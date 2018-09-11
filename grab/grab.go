package grab

import (
	"bufio"
	"fmt"
	"github.com/mchirico/date/parse"
	"github.com/mchirico/go_read/analysis"
	"log"
	"os"
	"regexp"
	"strings"
)

func matches(text string) bool {

	str := []string{"ladders",
		"linkedin", "spammer", "root"}

	for _, v := range str {

		if strings.Contains(text, v) {
			return true
		}
	}
	return false

}

func loop(text string, m map[string]int, tag string, records [][]string) [][]string {

	r, _ := regexp.Compile("from=<([a-z|-|_|.|0-9]+@+[a-z|-|0-9]+.[a-z|-]+)>")

	if matches(text) {
		return records
	}

	if r.MatchString(text) {
		s := text[0:15]
		tt, _ := parse.DateTimeParse(s).NewYork()

		email_time := tt.Format("2006-01-02 15:04:05-07:00")
		email := r.FindString(text)

		email = strings.Replace(email, "from=", "", -1)

		fmt.Printf("%v, \t %v, \t%v\n", email_time, tag, email)

		records = append(records, []string{tag, text, email, email_time})

		m[r.FindString(text)] += 1

	}

	return records
}

func FileParse(file string, dbFile string) map[string]int {

	fmt.Printf("file: %v\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		print("error opening file")
		return nil
	}
	defer f.Close()

	m := map[string]int{}

	records := [][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "discard:") {
			records = loop(scanner.Text(), m, "discard", records)
		}

		if strings.Contains(scanner.Text(), "reject: ") {
			records = loop(scanner.Text(), m, " reject", records)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	analysis.InsertData(dbFile, records)

	return m

}

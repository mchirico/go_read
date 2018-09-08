package grab

import (
	"bufio"
	"fmt"
	"github.com/mchirico/date/dateparse"
	"log"
	"os"
	"regexp"
	"strings"
)

func loop(text string, m map[string]int, tag string) {

	r, _ := regexp.Compile("<([a-z|-]+@+[a-z|-]+.[a-z|-]+)>")

	if strings.Contains(text, "reply") ||
		strings.Contains(text, "ladders") ||
		strings.Contains(text, "linkedin") ||
		strings.Contains(text, "spamer") ||
		strings.Contains(text, "root") ||
		strings.Contains(text, "mchirico") ||
		strings.Contains(text, "smtp") {
		return
	}

	if r.MatchString(text) {
		s := text[0:15]
		tt, _ := dateparse.DateTimeParse(s).NewYork()

		email_time := tt.Format("2006-01-02 3:04 PM")
		email := r.FindString(text)
		fmt.Printf("%v, \t %v, \t%v\n", email_time, tag, email)

		m[r.FindString(text)] += 1

	}
}

func FileParse(file string) map[string]int {

	fmt.Printf("file: %v\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		print("error opening file")
		return nil
	}
	defer f.Close()

	m := map[string]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "discard:") {
			loop(scanner.Text(), m, "discard")
		}

		if strings.Contains(scanner.Text(), "reject: ") {
			loop(scanner.Text(), m, " reject")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return m

}

package controller

import (
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://contests.arrl.org/contestmultipliers.php?a=wve")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	sections := doc.FindAll("td")

	for _, section := range sections {
		if len(section.Text()) <= 3 {
			if len(strings.TrimSpace(section.Text())) > 0 {
				// FIXME - replace Println with insert into psql
				fmt.Println(section.Text())
			}
		}
	}

}

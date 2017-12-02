package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	dlName := "all-postcode.csv"
	if _, err := os.Stat(dlName); os.IsNotExist(err) {
		// path/to/whatever does not exist
		download(dlName)
	}

	inp, err := os.Open(dlName)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("sheffield-postcode.csv")
	if err != nil {
		log.Fatal(err)
	}

	n := 0
	lines := bufio.NewScanner(inp)
	lines.Scan()
	fmt.Fprintln(out, lines.Text())
	for lines.Scan() {
		line := lines.Text()
		if strings.Contains(line, "E08000019") {
			fmt.Fprintln(out, lines.Text())
			n++
		}
	}
	fmt.Println(n, "records written")
}

func download(name string) {
	out, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err :=
		http.Get("http://geoportal.statistics.gov.uk/datasets/75edec484c5d49bcadd4893c0ebca0ff_0.csv")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(out, resp.Body)
	fmt.Println(n, "bytes copied")
}

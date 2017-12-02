package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	download()
}

func download() {
	out, err := os.Create("sheffield-postcode.csv")
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
	fmt.Print(n, " bytes copied\n")
}

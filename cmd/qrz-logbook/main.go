package main

import (
	"flag"
	"fmt"
	ql "github.com/xylo04/qrz-logbook"
	"os"
	"strings"
)

func main() {
	key := flag.String("key", "", "QRZ.com logbook API key")
	flag.Parse()
	if *key == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	status, err := ql.Status(key)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	println("The logbook is called", status.BookName)

	fetch, err2 := ql.Fetch(key)
	if err2 != nil {
		_, _ = fmt.Fprintln(os.Stderr, err2)
		os.Exit(1)
	}
	println("The logbook is has", fetch.Count, "records")
	lines := strings.Count(fetch.Adif, "\n") + 1
	println("The ADIF is has", lines, "lines")
}

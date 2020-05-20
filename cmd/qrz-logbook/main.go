package main

import (
	"flag"
	"fmt"
	ql "github.com/xylo04/qrz-logbook"
	"os"
)

func main() {
	key := flag.String("key", "", "QRZ.com logbook API key")
	flag.Parse()
	if *key == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	lookupResp, err := ql.Fetch(key)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	println("That logbook has", lookupResp.COUNT, "records")
}

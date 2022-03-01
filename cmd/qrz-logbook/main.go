package main

import (
	"context"
	"flag"
	"fmt"
	ql "github.com/k0swe/qrz-logbook"
	"os"
	"strconv"
	"strings"
)

func main() {
	ctx := context.Background()
	key := flag.String("key", "", "QRZ.com logbook API key")
	flag.Parse()
	if *key == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	status := getStatus(ctx, key)
	getRecords(ctx, key)

	myCall := strings.ReplaceAll(status.Callsign, "_", "/")
	adif := `<call:6>3L4SF <gridsquare:4>FM23 <mode:3>SSB <rst_sent:2>59 <rst_rcvd:2>59 
		<qso_date:8>20200521 <time_on:6>040315 <qso_date_off:8>20200521 <time_off:6>040400 
		<band:3>40m <freq:8>7.175950 <station_callsign:` + strconv.Itoa(len(myCall)) +
		`>` + myCall + ` <my_gridsquare:6>DM79LV <tx_pwr:2>40 <eor>`
	inserted := insertRecord(ctx, key, adif)
	deleteRecord(ctx, key, inserted.LogId)
}

func getStatus(ctx context.Context, key *string) *ql.StatusResponse {
	status, err := ql.Status(ctx, key)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	println("The logbook is called", status.BookName)
	return status
}

func getRecords(ctx context.Context, key *string) {
	fetch, err := ql.Fetch(ctx, key)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	println("The logbook is has", fetch.Count, "records")
	lines := strings.Count(fetch.Adif, "\n") + 1
	println("The ADIF is has", lines, "lines")
}

func insertRecord(ctx context.Context, key *string, adif string) *ql.InsertResponse {
	insert, err := ql.Insert(ctx, key, adif, true)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	println("Inserted a record, result", insert.Result, "with ID", insert.LogId)
	return insert
}

func deleteRecord(ctx context.Context, key *string, id string) {
	del, err := ql.Delete(ctx, key, []string{id})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	println("Deleted the record, result", del.Result)
}

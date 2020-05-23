package qrzlog

import (
	"context"
	"errors"
	"github.com/antihax/optional"
	"regexp"
	"strconv"
	"strings"
)

const agent = "xylo-go-1.0"

type FetchResponse struct {
	Adif   string
	Count  uint64
	Result string
}

func Fetch(key *string) (*FetchResponse, error) {
	client := newClient()
	opts := RootPostOpts{}
	apiResp, _, err := client.DefaultApi.RootPost(context.TODO(), *key, "FETCH", &opts)
	if err != nil {
		return nil, err
	}
	count, _ := strconv.ParseUint(apiResp.COUNT, 10, 64)
	adif := strings.Replace(apiResp.DATA, "ADIF=", "", -1)
	adif = strings.ReplaceAll(adif, "&lt;", "<")
	adif = strings.ReplaceAll(adif, "&gt;", ">")
	r := FetchResponse{
		Result: apiResp.RESULT,
		Count:  count,
		Adif:   adif,
	}
	if apiResp.RESULT == "FAIL" && apiResp.REASON != "" {
		// having a logbook with 0 QSOs will always result in a FAIL
		err = errors.New(apiResp.REASON)
	}
	return &r, err
}

type StatusResponse struct {
	Action    string
	BookId    string
	BookName  string
	Callsign  string
	Confirmed uint64
	Count     uint64
	DxccCount uint64
	EndDate   string
	Owner     string
	Result    string
	StartDate string
}

func Status(key *string) (*StatusResponse, error) {
	client := newClient()
	opts := RootPostOpts{}
	apiResp, _, err := client.DefaultApi.RootPost(context.TODO(), *key, "STATUS", &opts)
	if err != nil {
		return nil, err
	}
	count, _ := strconv.ParseUint(apiResp.COUNT, 10, 64)
	r := StatusResponse{
		Result: apiResp.RESULT,
		Count:  count,
	}

	var confirmedStr, dxccCountStr string
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)CONFIRMED=([\d]*)`), &confirmedStr)
	r.Confirmed, _ = strconv.ParseUint(confirmedStr, 10, 64)
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)DXCC_COUNT=(\d*)`), &dxccCountStr)
	r.DxccCount, _ = strconv.ParseUint(dxccCountStr, 10, 64)

	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)ACTION=(\w*)`), &r.Action)
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)BOOKID=(\d*)`), &r.BookId)
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)CALLSIGN=([\w-/]*)`), &r.Callsign)
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)END_DATE=([\d-]*)`), &r.EndDate)
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)OWNER=([\w-/]*)`), &r.Owner)
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)START_DATE=([\d-]*)`), &r.StartDate)

	// Put this at the end because it's the least strict regex
	findField(&apiResp.DATA, regexp.MustCompile(`(^|&)BOOK_NAME=(.*)`), &r.BookName)
	if apiResp.RESULT == "FAIL" && apiResp.REASON != "" {
		err = errors.New(apiResp.REASON)
	}
	return &r, err
}

type InsertResponse struct {
	Count  uint64
	Result string
	LogId  string
}

func Insert(key *string, adif string, replace bool) (*InsertResponse, error) {
	client := newClient()
	opts := RootPostOpts{
		ADIF: optional.NewString(adif),
	}
	if replace {
		opts.OPTION = optional.NewString("REPLACE")
	}
	apiResp, _, err := client.DefaultApi.RootPost(context.TODO(), *key, "INSERT", &opts)
	if err != nil {
		return nil, err
	}

	count, _ := strconv.ParseUint(apiResp.COUNT, 10, 64)
	r := InsertResponse{
		Result: apiResp.RESULT,
		LogId:  apiResp.LOGID,
		Count:  count,
	}
	if apiResp.RESULT == "FAIL" && apiResp.REASON != "" {
		err = errors.New(apiResp.REASON)
	}
	return &r, err
}

type DeleteResponse struct {
	Count  uint64
	Result string
	// The log IDs which were not deleted; Result will be PARTIAL
	LogIds  []string
}

func Delete(key *string, ids []string) (*DeleteResponse, error) {
	client := newClient()
	idsSlice := strings.Join(ids, ",")
	opts := RootPostOpts{
		LOGIDS: optional.NewString(idsSlice),
	}
	apiResp, _, err := client.DefaultApi.RootPost(context.TODO(), *key, "DELETE", &opts)
	if err != nil {
		return nil, err
	}

	count, _ := strconv.ParseUint(apiResp.COUNT, 10, 64)
	notDeleted := strings.Split(apiResp.LOGIDS, ",")
	r := DeleteResponse{
		Result: apiResp.RESULT,
		LogIds:  notDeleted,
		Count:  count,
	}
	if apiResp.RESULT == "FAIL" && apiResp.REASON != "" {
		err = errors.New(apiResp.REASON)
	}
	return &r, err
}

func newClient() *APIClient {
	config := NewConfiguration()
	config.UserAgent = agent
	return NewAPIClient(config)
}

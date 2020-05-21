package qrzlog

import (
	"context"
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
	return &r, err
}

func newClient() *APIClient {
	config := NewConfiguration()
	config.UserAgent = agent
	return NewAPIClient(config)
}

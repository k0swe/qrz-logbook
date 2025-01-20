package qrzlog

import (
	"net/url"
	"regexp"
	"strings"
)

func unmarshalForm(payload []byte, resp *Response) error {
	if resp == nil {
		resp = &Response{}
	}
	s := string(payload)
	// QRZ's Logbook API returns something resembling application/x-www-form-urlencoded, but
	// the fields are not escaped so attempting to use url.ParseQuery() directly fails.
	values, err := url.ParseQuery(s)
	if err != nil {
		return err
	}

	v1 := values.Get("RESULT")
	resp.RESULT = &v1
	v2 := values.Get("REASON")
	resp.REASON = &v2
	v3 := values.Get("LOGIDS")
	resp.LOGIDS = &v3
	v4 := values.Get("LOGID")
	resp.LOGID = &v4
	v5 := values.Get("COUNT")
	resp.COUNT = &v5

	trim := strings.Trim(s, "& \n")
	resp.DATA = &trim
	return nil
}

func findField(data *string, compile *regexp.Regexp, s *string) {

}

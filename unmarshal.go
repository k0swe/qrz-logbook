package qrzlog

import (
	"regexp"
	"strings"
)

func unmarshalForm(payload []byte, v interface{}) error {
	resp := v.(*Response)
	s := string(payload)
	// QRZ's Logbook API returns something resembling application/x-www-form-urlencoded, but
	// the fields are not escaped so attempting to use url.ParseQuery() directly fails.
	// TODO: This parser is a terrible hack and if you can think of something better, have at it.
	findField(&s, regexp.MustCompile(`(^|&)RESULT=(\w*)`), resp.RESULT)
	findField(&s, regexp.MustCompile(`(^|&)REASON=(\w*)`), resp.REASON)
	findField(&s, regexp.MustCompile(`(^|&)LOGIDS=([\d,]*)`), resp.LOGIDS)
	findField(&s, regexp.MustCompile(`(^|&)LOGID=(\d*)`), resp.LOGID)
	findField(&s, regexp.MustCompile(`(^|&)COUNT=(\d*)`), resp.COUNT)
	trim := strings.Trim(s, "& \n")
	resp.DATA = &trim
	return nil
}

func findField(payload *string, regex *regexp.Regexp, field *string) {
	submatch := regex.FindStringSubmatch(*payload)
	if submatch != nil {
		*field = submatch[2]
		*payload = strings.Replace(*payload, submatch[0], "", -1)
	}
}

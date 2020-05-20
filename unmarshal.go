package qrzlog

import (
	"regexp"
	"strings"
)

var (
	resultR = regexp.MustCompile(`RESULT=(\w*)`)
	reasonR = regexp.MustCompile(`REASON=(\w*)`)
	logIdsR = regexp.MustCompile(`LOGIDS=([\d,]*)`)
	logIdR  = regexp.MustCompile(`LOGID=(\d*)`)
	countR  = regexp.MustCompile(`COUNT=(\d*)`)
	dataR   = regexp.MustCompile(`DATA=(\w*)`)
	adifR   = regexp.MustCompile(`ADIF=`)
)

func unmarshalForm(payload []byte, v interface{}) error {
	resp := v.(*Response)
	s := string(payload)
	// QRZ's Logbook API returns something resembling application/x-www-form-urlencoded, but
	// the values are not escaped so attempting to use url.ParseQuery() directly fails.
	submatch := resultR.FindStringSubmatch(s)
	if submatch != nil {
		resp.RESULT = submatch[1]
		s = strings.Replace(s, submatch[0], "", -1)
	}
	submatch = reasonR.FindStringSubmatch(s)
	if submatch != nil {
		resp.REASON = submatch[1]
		s = strings.Replace(s, submatch[0], "", -1)
	}
	submatch = logIdsR.FindStringSubmatch(s)
	if submatch != nil {
		resp.LOGIDS = submatch[1]
		s = strings.Replace(s, submatch[0], "", -1)
	}
	submatch = logIdR.FindStringSubmatch(s)
	if submatch != nil {
		resp.LOGID = submatch[1]
		s = strings.Replace(s, submatch[0], "", -1)
	}
	submatch = countR.FindStringSubmatch(s)
	if submatch != nil {
		resp.COUNT = submatch[1]
		s = strings.Replace(s, submatch[0], "", -1)
	}
	submatch = dataR.FindStringSubmatch(s)
	if submatch != nil {
		resp.DATA = submatch[1]
		s = strings.Replace(s, submatch[0], "", -1)
	}
	submatch = adifR.FindStringSubmatch(s)
	if submatch != nil {
		adif := strings.Replace(s, "ADIF=", "", -1)
		adif = strings.ReplaceAll(adif, "&lt;", "<")
		adif = strings.ReplaceAll(adif, "&gt;", ">")
		resp.ADIF = adif
		s = ""
	}
	return nil
}

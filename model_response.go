/*
 * QRZ Logbook API
 *
 * This API provides methods for external programs to interact with the QRZ Logbook using an HTTP REST interface. The QRZ Logbook is a combination free and paid subscription service of QRZ. Some advanced features require a valid subscription while the majority of operations are free to all QRZ members. All users of the QRZ Logbook, regardless of their subscription status, may access, edit, update, and view their complete logs online at the QRZ website.
 *
 * API version: 1.0.0
 * Contact: flloyd@qrz.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qrzlog

// Response struct for Response
type Response struct {
	// OK when the operation succeeds, FAIL when the operation failed, AUTH when API Access lacked sufficient privileges for the operation, or other action-specific codes
	RESULT string `json:"RESULT,omitempty"`
	// Used with RESULT=FAIL to describe the specific reason for failure
	REASON string `json:"REASON,omitempty"`
	// A comma separated list of logid values that were affected by the action
	LOGIDS string `json:"LOGIDS,omitempty"`
	// The logid value of the record that was inserted or replaced. (Singular \"LOGID\" response only used by INSERT as it is a single record operation.)
	LOGID string `json:"LOGID,omitempty"`
	// The number of QSO records that were affected by the action
	COUNT string `json:"COUNT,omitempty"`
	// Used for action-specific data such as status reports
	DATA string `json:"DATA,omitempty"`
	// ADIF data containing the QSO’s that match the selection criteria, limited by the MAX parameter.
	ADIF string `json:"ADIF,omitempty"`
}

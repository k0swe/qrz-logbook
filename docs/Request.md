# Request

## Properties

| Name       | Type                  | Description                                        | Notes      |
| ---------- | --------------------- | -------------------------------------------------- | ---------- |
| **KEY**    | **string**            | A QRZ supplied logbook access key                  |
| **ACTION** | **string**            | Type of request, i.e. INSERT, DELETE, UPLOAD, etc. |
| **ADIF**   | Pointer to **string** | ADIF formatted input data                          | [optional] |
| **OPTION** | Pointer to **string** | Action-specific options                            | [optional] |
| **LOGIDS** | Pointer to **string** | A comma separated list of integer logid values     | [optional] |

## Methods

### NewRequest

`func NewRequest(kEY string, aCTION string, ) *Request`

NewRequest instantiates a new Request object This constructor will assign
default values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set of
required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object This constructor will
only assign default values to properties that have it defined, but it doesn't
guarantee that properties required by API are set

### GetKEY

`func (o *Request) GetKEY() string`

GetKEY returns the KEY field if non-nil, zero value otherwise.

### GetKEYOk

`func (o *Request) GetKEYOk() (*string, bool)`

GetKEYOk returns a tuple with the KEY field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetKEY

`func (o *Request) SetKEY(v string)`

SetKEY sets KEY field to given value.

### GetACTION

`func (o *Request) GetACTION() string`

GetACTION returns the ACTION field if non-nil, zero value otherwise.

### GetACTIONOk

`func (o *Request) GetACTIONOk() (*string, bool)`

GetACTIONOk returns a tuple with the ACTION field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetACTION

`func (o *Request) SetACTION(v string)`

SetACTION sets ACTION field to given value.

### GetADIF

`func (o *Request) GetADIF() string`

GetADIF returns the ADIF field if non-nil, zero value otherwise.

### GetADIFOk

`func (o *Request) GetADIFOk() (*string, bool)`

GetADIFOk returns a tuple with the ADIF field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetADIF

`func (o *Request) SetADIF(v string)`

SetADIF sets ADIF field to given value.

### HasADIF

`func (o *Request) HasADIF() bool`

HasADIF returns a boolean if a field has been set.

### GetOPTION

`func (o *Request) GetOPTION() string`

GetOPTION returns the OPTION field if non-nil, zero value otherwise.

### GetOPTIONOk

`func (o *Request) GetOPTIONOk() (*string, bool)`

GetOPTIONOk returns a tuple with the OPTION field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetOPTION

`func (o *Request) SetOPTION(v string)`

SetOPTION sets OPTION field to given value.

### HasOPTION

`func (o *Request) HasOPTION() bool`

HasOPTION returns a boolean if a field has been set.

### GetLOGIDS

`func (o *Request) GetLOGIDS() string`

GetLOGIDS returns the LOGIDS field if non-nil, zero value otherwise.

### GetLOGIDSOk

`func (o *Request) GetLOGIDSOk() (*string, bool)`

GetLOGIDSOk returns a tuple with the LOGIDS field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLOGIDS

`func (o *Request) SetLOGIDS(v string)`

SetLOGIDS sets LOGIDS field to given value.

### HasLOGIDS

`func (o *Request) HasLOGIDS() bool`

HasLOGIDS returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

# Response

## Properties

| Name       | Type                  | Description                                                                                                                                                         | Notes      |
| ---------- | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **RESULT** | Pointer to **string** | OK when the operation succeeds, FAIL when the operation failed, AUTH when API Access lacked sufficient privileges for the operation, or other action-specific codes | [optional] |
| **REASON** | Pointer to **string** | Used with RESULT&#x3D;FAIL to describe the specific reason for failure                                                                                              | [optional] |
| **LOGIDS** | Pointer to **string** | A comma separated list of logid values that were affected by the action                                                                                             | [optional] |
| **LOGID**  | Pointer to **string** | The logid value of the record that was inserted or replaced. (Singular \&quot;LOGID\&quot; response only used by INSERT as it is a single record operation.)        | [optional] |
| **COUNT**  | Pointer to **string** | The number of QSO records that were affected by the action                                                                                                          | [optional] |
| **DATA**   | Pointer to **string** | Used for action-specific data such as status reports                                                                                                                | [optional] |

## Methods

### NewResponse

`func NewResponse() *Response`

NewResponse instantiates a new Response object This constructor will assign
default values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set of
required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object This constructor will
only assign default values to properties that have it defined, but it doesn't
guarantee that properties required by API are set

### GetRESULT

`func (o *Response) GetRESULT() string`

GetRESULT returns the RESULT field if non-nil, zero value otherwise.

### GetRESULTOk

`func (o *Response) GetRESULTOk() (*string, bool)`

GetRESULTOk returns a tuple with the RESULT field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetRESULT

`func (o *Response) SetRESULT(v string)`

SetRESULT sets RESULT field to given value.

### HasRESULT

`func (o *Response) HasRESULT() bool`

HasRESULT returns a boolean if a field has been set.

### GetREASON

`func (o *Response) GetREASON() string`

GetREASON returns the REASON field if non-nil, zero value otherwise.

### GetREASONOk

`func (o *Response) GetREASONOk() (*string, bool)`

GetREASONOk returns a tuple with the REASON field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetREASON

`func (o *Response) SetREASON(v string)`

SetREASON sets REASON field to given value.

### HasREASON

`func (o *Response) HasREASON() bool`

HasREASON returns a boolean if a field has been set.

### GetLOGIDS

`func (o *Response) GetLOGIDS() string`

GetLOGIDS returns the LOGIDS field if non-nil, zero value otherwise.

### GetLOGIDSOk

`func (o *Response) GetLOGIDSOk() (*string, bool)`

GetLOGIDSOk returns a tuple with the LOGIDS field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLOGIDS

`func (o *Response) SetLOGIDS(v string)`

SetLOGIDS sets LOGIDS field to given value.

### HasLOGIDS

`func (o *Response) HasLOGIDS() bool`

HasLOGIDS returns a boolean if a field has been set.

### GetLOGID

`func (o *Response) GetLOGID() string`

GetLOGID returns the LOGID field if non-nil, zero value otherwise.

### GetLOGIDOk

`func (o *Response) GetLOGIDOk() (*string, bool)`

GetLOGIDOk returns a tuple with the LOGID field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLOGID

`func (o *Response) SetLOGID(v string)`

SetLOGID sets LOGID field to given value.

### HasLOGID

`func (o *Response) HasLOGID() bool`

HasLOGID returns a boolean if a field has been set.

### GetCOUNT

`func (o *Response) GetCOUNT() string`

GetCOUNT returns the COUNT field if non-nil, zero value otherwise.

### GetCOUNTOk

`func (o *Response) GetCOUNTOk() (*string, bool)`

GetCOUNTOk returns a tuple with the COUNT field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCOUNT

`func (o *Response) SetCOUNT(v string)`

SetCOUNT sets COUNT field to given value.

### HasCOUNT

`func (o *Response) HasCOUNT() bool`

HasCOUNT returns a boolean if a field has been set.

### GetDATA

`func (o *Response) GetDATA() string`

GetDATA returns the DATA field if non-nil, zero value otherwise.

### GetDATAOk

`func (o *Response) GetDATAOk() (*string, bool)`

GetDATAOk returns a tuple with the DATA field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetDATA

`func (o *Response) SetDATA(v string)`

SetDATA sets DATA field to given value.

### HasDATA

`func (o *Response) HasDATA() bool`

HasDATA returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

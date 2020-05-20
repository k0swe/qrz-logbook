# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RESULT** | **string** | OK when the operation succeeds, FAIL when the operation failed, AUTH when API Access lacked sufficient privileges for the operation, or other action-specific codes | [optional] 
**REASON** | **string** | Used with RESULT&#x3D;FAIL to describe the specific reason for failure | [optional] 
**LOGIDS** | **string** | A comma separated list of logid values that were affected by the action | [optional] 
**LOGID** | **string** | The logid value of the record that was inserted or replaced. (Singular \&quot;LOGID\&quot; response only used by INSERT as it is a single record operation.) | [optional] 
**COUNT** | **string** | The number of QSO records that were affected by the action | [optional] 
**DATA** | **string** | Used for action-specific data such as status reports | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



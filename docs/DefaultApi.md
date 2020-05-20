# \DefaultApi

All URIs are relative to *https://logbook.qrz.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoEverything**](DefaultApi.md#DoEverything) | **Post** / | The do-everything endpoint



## DoEverything

> Response DoEverything(ctx, kEY, aCTION, optional)

The do-everything endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kEY** | **string**| A QRZ supplied logbook access key | 
**aCTION** | **string**| Type of request, i.e. INSERT, DELETE, UPLOAD, etc. | 
 **optional** | ***DoEverythingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DoEverythingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aDIF** | **optional.String**| ADIF formatted input data | 
 **oPTION** | **optional.String**| Action-specific options | 
 **lOGIDS** | **optional.String**| A comma separated list of integer logid values | 

### Return type

[**Response**](Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


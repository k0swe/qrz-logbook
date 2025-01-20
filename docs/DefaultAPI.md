# \DefaultAPI

All URIs are relative to *https://logbook.qrz.com/api*

| Method                                 | HTTP request | Description                |
| -------------------------------------- | ------------ | -------------------------- |
| [**RootPost**](DefaultAPI.md#RootPost) | **Post** /   | The do-everything endpoint |

## RootPost

> Response
> RootPost(ctx).KEY(kEY).ACTION(aCTION).ADIF(aDIF).OPTION(oPTION).LOGIDS(lOGIDS).Execute()

The do-everything endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/k0swe/qrz-logbook"
)

func main() {
	kEY := "kEY_example" // string | A QRZ supplied logbook access key
	aCTION := "aCTION_example" // string | Type of request, i.e. INSERT, DELETE, UPLOAD, etc.
	aDIF := "aDIF_example" // string | ADIF formatted input data (optional)
	oPTION := "oPTION_example" // string | Action-specific options (optional)
	lOGIDS := "lOGIDS_example" // string | A comma separated list of integer logid values (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RootPost(context.Background()).KEY(kEY).ACTION(aCTION).ADIF(aDIF).OPTION(oPTION).LOGIDS(lOGIDS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RootPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RootPost`: Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RootPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRootPostRequest struct via
the builder pattern

| Name       | Type       | Description                                        | Notes |
| ---------- | ---------- | -------------------------------------------------- | ----- |
| **kEY**    | **string** | A QRZ supplied logbook access key                  |
| **aCTION** | **string** | Type of request, i.e. INSERT, DELETE, UPLOAD, etc. |
| **aDIF**   | **string** | ADIF formatted input data                          |
| **oPTION** | **string** | Action-specific options                            |
| **lOGIDS** | **string** | A comma separated list of integer logid values     |

### Return type

[**Response**](Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/x-www-form-urlencoded

[[Back to top]](#)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

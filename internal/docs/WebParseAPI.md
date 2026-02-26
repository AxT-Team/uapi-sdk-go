# \WebParseAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebTomarkdownAsyncStatus**](WebParseAPI.md#GetWebTomarkdownAsyncStatus) | **Get** /web/tomarkdown/async/{task_id} | 转换任务状态
[**GetWebparseExtractimages**](WebParseAPI.md#GetWebparseExtractimages) | **Get** /webparse/extractimages | 提取网页图片
[**GetWebparseMetadata**](WebParseAPI.md#GetWebparseMetadata) | **Get** /webparse/metadata | 提取网页元数据
[**PostWebTomarkdownAsync**](WebParseAPI.md#PostWebTomarkdownAsync) | **Post** /web/tomarkdown/async | 网页转 Markdown



## GetWebTomarkdownAsyncStatus

> GetWebTomarkdownAsyncStatus200Response GetWebTomarkdownAsyncStatus(ctx, taskId).Execute()

转换任务状态



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	taskId := "a1b2c3d4-e5f6-47a8-b9c0-d1e2f3a4b5c6" // string | 任务ID（由提交接口返回）

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebParseAPI.GetWebTomarkdownAsyncStatus(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebParseAPI.GetWebTomarkdownAsyncStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebTomarkdownAsyncStatus`: GetWebTomarkdownAsyncStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `WebParseAPI.GetWebTomarkdownAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | 任务ID（由提交接口返回） | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebTomarkdownAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWebTomarkdownAsyncStatus200Response**](GetWebTomarkdownAsyncStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebparseExtractimages

> GetWebparseExtractimages200Response GetWebparseExtractimages(ctx).Url(url).Execute()

提取网页图片



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	url := "https://cn.bing.com/" // string | 需要提取图片的网页URL

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebParseAPI.GetWebparseExtractimages(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebParseAPI.GetWebparseExtractimages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebparseExtractimages`: GetWebparseExtractimages200Response
	fmt.Fprintf(os.Stdout, "Response from `WebParseAPI.GetWebparseExtractimages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebparseExtractimagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | 需要提取图片的网页URL | 

### Return type

[**GetWebparseExtractimages200Response**](GetWebparseExtractimages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebparseMetadata

> GetWebparseMetadata200Response GetWebparseMetadata(ctx).Url(url).Execute()

提取网页元数据



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	url := "https://www.bilibili.com" // string | 需要提取元数据的网页URL

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebParseAPI.GetWebparseMetadata(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebParseAPI.GetWebparseMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebparseMetadata`: GetWebparseMetadata200Response
	fmt.Fprintf(os.Stdout, "Response from `WebParseAPI.GetWebparseMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebparseMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | 需要提取元数据的网页URL | 

### Return type

[**GetWebparseMetadata200Response**](GetWebparseMetadata200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebTomarkdownAsync

> PostWebTomarkdownAsync202Response PostWebTomarkdownAsync(ctx).Url(url).Execute()

网页转 Markdown



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	url := "https://example.com" // string | 需要转换的网页URL。URL必须经过编码。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebParseAPI.PostWebTomarkdownAsync(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebParseAPI.PostWebTomarkdownAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebTomarkdownAsync`: PostWebTomarkdownAsync202Response
	fmt.Fprintf(os.Stdout, "Response from `WebParseAPI.PostWebTomarkdownAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebTomarkdownAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | 需要转换的网页URL。URL必须经过编码。 | 

### Return type

[**PostWebTomarkdownAsync202Response**](PostWebTomarkdownAsync202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


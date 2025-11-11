# \ConvertAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConvertUnixtime**](ConvertAPI.md#GetConvertUnixtime) | **Get** /convert/unixtime | Unix时间戳与日期字符串双向转换
[**PostConvertJson**](ConvertAPI.md#PostConvertJson) | **Post** /convert/json | 美化并格式化JSON字符串



## GetConvertUnixtime

> GetConvertUnixtime200Response GetConvertUnixtime(ctx).Time(time).Execute()

Unix时间戳与日期字符串双向转换



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
	time := "1698380645" // string | 一个智能时间参数，可传入Unix时间戳（10位或13位）或标准日期字符串（如 '2023-10-27 10:30:00'），系统将自动识别并转换。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.GetConvertUnixtime(context.Background()).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.GetConvertUnixtime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConvertUnixtime`: GetConvertUnixtime200Response
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.GetConvertUnixtime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConvertUnixtimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **string** | 一个智能时间参数，可传入Unix时间戳（10位或13位）或标准日期字符串（如 &#39;2023-10-27 10:30:00&#39;），系统将自动识别并转换。 | 

### Return type

[**GetConvertUnixtime200Response**](GetConvertUnixtime200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConvertJson

> PostConvertJson200Response PostConvertJson(ctx).PostConvertJsonRequest(postConvertJsonRequest).Execute()

美化并格式化JSON字符串



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
	postConvertJsonRequest := *openapiclient.NewPostConvertJsonRequest("{"name":"John Doe","age":30,"isStudent":false,"courses":[{"title":"History","credits":3},{"title":"Math","credits":4}]}") // PostConvertJsonRequest | 这是一个JSON对象，里面必须包含一个名为 `content` 的字段。这个字段的值，就是你希望格式化的、原始的JSON字符串。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.PostConvertJson(context.Background()).PostConvertJsonRequest(postConvertJsonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.PostConvertJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConvertJson`: PostConvertJson200Response
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.PostConvertJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConvertJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postConvertJsonRequest** | [**PostConvertJsonRequest**](PostConvertJsonRequest.md) | 这是一个JSON对象，里面必须包含一个名为 &#x60;content&#x60; 的字段。这个字段的值，就是你希望格式化的、原始的JSON字符串。 | 

### Return type

[**PostConvertJson200Response**](PostConvertJson200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


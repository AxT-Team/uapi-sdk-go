# \StatusAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatusRatelimit**](StatusAPI.md#GetStatusRatelimit) | **Get** /status/ratelimit | 获取API限流器实时状态
[**GetStatusUsage**](StatusAPI.md#GetStatusUsage) | **Get** /status/usage | 获取API端点使用统计



## GetStatusRatelimit

> GetStatusRatelimit200Response GetStatusRatelimit(ctx).Authorization(authorization).Execute()

获取API限流器实时状态



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
	authorization := "Bearer sk-xxx" // string | Bearer类型的API密钥认证头。例如：`Bearer sk-xxx`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetStatusRatelimit(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetStatusRatelimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusRatelimit`: GetStatusRatelimit200Response
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetStatusRatelimit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRatelimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer类型的API密钥认证头。例如：&#x60;Bearer sk-xxx&#x60; | 

### Return type

[**GetStatusRatelimit200Response**](GetStatusRatelimit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusUsage

> GetStatusUsage200Response GetStatusUsage(ctx).Path(path).Execute()

获取API端点使用统计



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
	path := "/api/v1/image/motou" // string | （可选）如果你想查询某个特定的端点，请提供它的路径，例如 '/api/v1/image/motou'。如果留空，则返回所有端点的统计列表。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetStatusUsage(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetStatusUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusUsage`: GetStatusUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetStatusUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | （可选）如果你想查询某个特定的端点，请提供它的路径，例如 &#39;/api/v1/image/motou&#39;。如果留空，则返回所有端点的统计列表。 | 

### Return type

[**GetStatusUsage200Response**](GetStatusUsage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


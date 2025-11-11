# \ClipzyAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClipzyGet**](ClipzyAPI.md#GetClipzyGet) | **Get** /api/get | 步骤2 (方法一): 获取加密数据
[**GetClipzyRaw**](ClipzyAPI.md#GetClipzyRaw) | **Get** /api/raw/{id} | 步骤2 (方法二): 获取原始文本
[**PostClipzyStore**](ClipzyAPI.md#PostClipzyStore) | **Post** /api/store | 步骤1：上传加密数据



## GetClipzyGet

> GetClipzyGet200Response GetClipzyGet(ctx).Id(id).Execute()

步骤2 (方法一): 获取加密数据



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
	id := "PREF0Zv8Yj" // string | 片段的唯一 ID。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClipzyAPI.GetClipzyGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClipzyAPI.GetClipzyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClipzyGet`: GetClipzyGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ClipzyAPI.GetClipzyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClipzyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 片段的唯一 ID。 | 

### Return type

[**GetClipzyGet200Response**](GetClipzyGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClipzyRaw

> string GetClipzyRaw(ctx, id).Key(key).Execute()

步骤2 (方法二): 获取原始文本



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
	id := "PREF0Zv8Yj" // string | 片段的唯一 ID。
	key := "ES9tGP0v3e7oqzmAk3vMboxVOOBlvw9RS3DszeW675k=" // string | 用于解密的 Base64 编码的 AES 密钥。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClipzyAPI.GetClipzyRaw(context.Background(), id).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClipzyAPI.GetClipzyRaw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClipzyRaw`: string
	fmt.Fprintf(os.Stdout, "Response from `ClipzyAPI.GetClipzyRaw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 片段的唯一 ID。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClipzyRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | 用于解密的 Base64 编码的 AES 密钥。 | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClipzyStore

> PostClipzyStore200Response PostClipzyStore(ctx).PostClipzyStoreRequest(postClipzyStoreRequest).Execute()

步骤1：上传加密数据



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
	postClipzyStoreRequest := *openapiclient.NewPostClipzyStoreRequest("PIZQqgLgUgNg6gEQEYAUCuBrA7gSQBoDMAiigPakCOAYgQIYBKEA7AFbABqAEiAJYBmAYSJUAznHbsAcgBMBAFQDSAegCcALz4BZKHDwql9DADtOCqvQCCATQowAWhgVgYAUwBMAFgC8QA==") // PostClipzyStoreRequest | 包含加密数据和可选的TTL。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClipzyAPI.PostClipzyStore(context.Background()).PostClipzyStoreRequest(postClipzyStoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClipzyAPI.PostClipzyStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClipzyStore`: PostClipzyStore200Response
	fmt.Fprintf(os.Stdout, "Response from `ClipzyAPI.PostClipzyStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClipzyStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postClipzyStoreRequest** | [**PostClipzyStoreRequest**](PostClipzyStoreRequest.md) | 包含加密数据和可选的TTL。 | 

### Return type

[**PostClipzyStore200Response**](PostClipzyStore200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


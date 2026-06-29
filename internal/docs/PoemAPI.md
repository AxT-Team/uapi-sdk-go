# \PoemAPI

All URIs are relative to *https://uapis.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSaying**](PoemAPI.md#GetSaying) | **Get** /saying | 一言
[**GetSayingRandom**](PoemAPI.md#GetSayingRandom) | **Get** /saying/random | 一言（随机/每日/场景/此刻）



## GetSaying

> GetSaying200Response GetSaying(ctx).Execute()

一言



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoemAPI.GetSaying(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoemAPI.GetSaying``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSaying`: GetSaying200Response
	fmt.Fprintf(os.Stdout, "Response from `PoemAPI.GetSaying`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSayingRequest struct via the builder pattern


### Return type

[**GetSaying200Response**](GetSaying200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSayingRandom

> GetSayingRandom200Response GetSayingRandom(ctx).Mode(mode).Scene(scene).Source(source).Category(category).Tag(tag).Execute()

一言（随机/每日/场景/此刻）



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
	mode := "mode_example" // string | 运行模式。不传或 random 为随机一言；可选 daily、recommend、moment。 (optional) (default to "random")
	scene := "scene_example" // string | 推荐场景。当 mode=recommend 时必填，例如 night、morning、work 等。请从[支持的场景列表](#enum-list)中选择。 (optional)
	source := "source_example" // string | 语料来源过滤。支持重复传参，或使用逗号/分号分隔多个值。请从[支持的来源列表](#enum-list)中选择。 (optional)
	category := "category_example" // string | 分类过滤。支持重复传参，或使用逗号/分号分隔多个值。请从[支持的分类列表](#enum-list)中选择。 (optional)
	tag := "tag_example" // string | 标签过滤。支持重复传参，或使用逗号/分号分隔多个值。请从[支持的标签列表](#enum-list)中选择。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoemAPI.GetSayingRandom(context.Background()).Mode(mode).Scene(scene).Source(source).Category(category).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoemAPI.GetSayingRandom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSayingRandom`: GetSayingRandom200Response
	fmt.Fprintf(os.Stdout, "Response from `PoemAPI.GetSayingRandom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSayingRandomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | 运行模式。不传或 random 为随机一言；可选 daily、recommend、moment。 | [default to &quot;random&quot;]
 **scene** | **string** | 推荐场景。当 mode&#x3D;recommend 时必填，例如 night、morning、work 等。请从[支持的场景列表](#enum-list)中选择。 | 
 **source** | **string** | 语料来源过滤。支持重复传参，或使用逗号/分号分隔多个值。请从[支持的来源列表](#enum-list)中选择。 | 
 **category** | **string** | 分类过滤。支持重复传参，或使用逗号/分号分隔多个值。请从[支持的分类列表](#enum-list)中选择。 | 
 **tag** | **string** | 标签过滤。支持重复传参，或使用逗号/分号分隔多个值。请从[支持的标签列表](#enum-list)中选择。 | 

### Return type

[**GetSayingRandom200Response**](GetSayingRandom200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \DefaultApi

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchEngines**](DefaultApi.md#GetSearchEngines) | **Get** /search/engines | 搜索引擎配置
[**GetSensitiveWordAnalyzeQuery**](DefaultApi.md#GetSensitiveWordAnalyzeQuery) | **Get** /sensitive-word/analyze-query | 敏感词分析 (GET)
[**PostSearchAggregate**](DefaultApi.md#PostSearchAggregate) | **Post** /search/aggregate | 智能搜索
[**PostSensitiveWordAnalyze**](DefaultApi.md#PostSensitiveWordAnalyze) | **Post** /sensitive-word/analyze | 分析敏感词
[**PostSensitiveWordQuickCheck**](DefaultApi.md#PostSensitiveWordQuickCheck) | **Post** /text/profanitycheck | 敏感词检测（快速）



## GetSearchEngines

> GetSearchEngines200Response GetSearchEngines(ctx).Execute()

搜索引擎配置



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
	resp, r, err := apiClient.DefaultApi.GetSearchEngines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSearchEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchEngines`: GetSearchEngines200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSearchEngines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchEnginesRequest struct via the builder pattern


### Return type

[**GetSearchEngines200Response**](GetSearchEngines200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSensitiveWordAnalyzeQuery

> PostSensitiveWordAnalyze200Response GetSensitiveWordAnalyzeQuery(ctx).Keyword(keyword).Execute()

敏感词分析 (GET)



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
	keyword := "喵" // string | 要分析的关键词，最长50字符。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSensitiveWordAnalyzeQuery(context.Background()).Keyword(keyword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSensitiveWordAnalyzeQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSensitiveWordAnalyzeQuery`: PostSensitiveWordAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSensitiveWordAnalyzeQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSensitiveWordAnalyzeQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyword** | **string** | 要分析的关键词，最长50字符。 | 

### Return type

[**PostSensitiveWordAnalyze200Response**](PostSensitiveWordAnalyze200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSearchAggregate

> PostSearchAggregate200Response PostSearchAggregate(ctx).PostSearchAggregateRequest(postSearchAggregateRequest).Execute()

智能搜索



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
	postSearchAggregateRequest := *openapiclient.NewPostSearchAggregateRequest("Go最新的版本是多少") // PostSearchAggregateRequest | 包含搜索参数的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PostSearchAggregate(context.Background()).PostSearchAggregateRequest(postSearchAggregateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSearchAggregate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSearchAggregate`: PostSearchAggregate200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSearchAggregate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSearchAggregateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSearchAggregateRequest** | [**PostSearchAggregateRequest**](PostSearchAggregateRequest.md) | 包含搜索参数的JSON对象 | 

### Return type

[**PostSearchAggregate200Response**](PostSearchAggregate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSensitiveWordAnalyze

> PostSensitiveWordAnalyze200Response PostSensitiveWordAnalyze(ctx).PostSensitiveWordAnalyzeRequest(postSensitiveWordAnalyzeRequest).Execute()

分析敏感词



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
	postSensitiveWordAnalyzeRequest := *openapiclient.NewPostSensitiveWordAnalyzeRequest([]string{"Keywords_example"}) // PostSensitiveWordAnalyzeRequest | 包含待检测文本 'keywords' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PostSensitiveWordAnalyze(context.Background()).PostSensitiveWordAnalyzeRequest(postSensitiveWordAnalyzeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSensitiveWordAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSensitiveWordAnalyze`: PostSensitiveWordAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSensitiveWordAnalyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSensitiveWordAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSensitiveWordAnalyzeRequest** | [**PostSensitiveWordAnalyzeRequest**](PostSensitiveWordAnalyzeRequest.md) | 包含待检测文本 &#39;keywords&#39; 的JSON对象 | 

### Return type

[**PostSensitiveWordAnalyze200Response**](PostSensitiveWordAnalyze200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSensitiveWordQuickCheck

> PostSensitiveWordQuickCheck200Response PostSensitiveWordQuickCheck(ctx).PostSensitiveWordQuickCheckRequest(postSensitiveWordQuickCheckRequest).Execute()

敏感词检测（快速）



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
	postSensitiveWordQuickCheckRequest := *openapiclient.NewPostSensitiveWordQuickCheckRequest("这是一段友善的、不包含任何违禁词的文本。") // PostSensitiveWordQuickCheckRequest | 包含待检测文本 'text' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PostSensitiveWordQuickCheck(context.Background()).PostSensitiveWordQuickCheckRequest(postSensitiveWordQuickCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSensitiveWordQuickCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSensitiveWordQuickCheck`: PostSensitiveWordQuickCheck200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSensitiveWordQuickCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSensitiveWordQuickCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSensitiveWordQuickCheckRequest** | [**PostSensitiveWordQuickCheckRequest**](PostSensitiveWordQuickCheckRequest.md) | 包含待检测文本 &#39;text&#39; 的JSON对象 | 

### Return type

[**PostSensitiveWordQuickCheck200Response**](PostSensitiveWordQuickCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


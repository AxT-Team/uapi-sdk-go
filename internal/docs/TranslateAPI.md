# \TranslateAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAiTranslateLanguages**](TranslateAPI.md#GetAiTranslateLanguages) | **Get** /ai/translate/languages | AI翻译配置
[**PostAiTranslate**](TranslateAPI.md#PostAiTranslate) | **Post** /ai/translate | AI智能翻译
[**PostTranslateStream**](TranslateAPI.md#PostTranslateStream) | **Post** /translate/stream | 流式翻译（中英互译）
[**PostTranslateText**](TranslateAPI.md#PostTranslateText) | **Post** /translate/text | 翻译



## GetAiTranslateLanguages

> GetAiTranslateLanguages200Response GetAiTranslateLanguages(ctx).Execute()

AI翻译配置



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
	resp, r, err := apiClient.TranslateAPI.GetAiTranslateLanguages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.GetAiTranslateLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiTranslateLanguages`: GetAiTranslateLanguages200Response
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.GetAiTranslateLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiTranslateLanguagesRequest struct via the builder pattern


### Return type

[**GetAiTranslateLanguages200Response**](GetAiTranslateLanguages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiTranslate

> PostAiTranslate200Response PostAiTranslate(ctx).TargetLang(targetLang).PostAiTranslateRequest(postAiTranslateRequest).Execute()

AI智能翻译



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
	targetLang := "zh-CHS" // string | 目标语言代码。请从支持的语言列表中选择一个语言代码。
	postAiTranslateRequest := *openapiclient.NewPostAiTranslateRequest() // PostAiTranslateRequest | 包含翻译参数的JSON对象，支持单个文本或批量文本翻译

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslateAPI.PostAiTranslate(context.Background()).TargetLang(targetLang).PostAiTranslateRequest(postAiTranslateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.PostAiTranslate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiTranslate`: PostAiTranslate200Response
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.PostAiTranslate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiTranslateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetLang** | **string** | 目标语言代码。请从支持的语言列表中选择一个语言代码。 | 
 **postAiTranslateRequest** | [**PostAiTranslateRequest**](PostAiTranslateRequest.md) | 包含翻译参数的JSON对象，支持单个文本或批量文本翻译 | 

### Return type

[**PostAiTranslate200Response**](PostAiTranslate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTranslateStream

> string PostTranslateStream(ctx).PostTranslateStreamRequest(postTranslateStreamRequest).Execute()

流式翻译（中英互译）



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
	postTranslateStreamRequest := *openapiclient.NewPostTranslateStreamRequest("Hello, how are you?", "中文") // PostTranslateStreamRequest | 包含翻译参数的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslateAPI.PostTranslateStream(context.Background()).PostTranslateStreamRequest(postTranslateStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.PostTranslateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTranslateStream`: string
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.PostTranslateStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTranslateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTranslateStreamRequest** | [**PostTranslateStreamRequest**](PostTranslateStreamRequest.md) | 包含翻译参数的JSON对象 | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTranslateText

> PostTranslateText200Response PostTranslateText(ctx).ToLang(toLang).PostTranslateTextRequest(postTranslateTextRequest).Execute()

翻译



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
	toLang := "zh-CHS" // string | 目标语言代码。请从支持的语言列表中选择一个语言代码。
	postTranslateTextRequest := *openapiclient.NewPostTranslateTextRequest("hello world") // PostTranslateTextRequest | 包含待翻译文本的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslateAPI.PostTranslateText(context.Background()).ToLang(toLang).PostTranslateTextRequest(postTranslateTextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslateAPI.PostTranslateText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTranslateText`: PostTranslateText200Response
	fmt.Fprintf(os.Stdout, "Response from `TranslateAPI.PostTranslateText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTranslateTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toLang** | **string** | 目标语言代码。请从支持的语言列表中选择一个语言代码。 | 
 **postTranslateTextRequest** | [**PostTranslateTextRequest**](PostTranslateTextRequest.md) | 包含待翻译文本的JSON对象 | 

### Return type

[**PostTranslateText200Response**](PostTranslateText200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


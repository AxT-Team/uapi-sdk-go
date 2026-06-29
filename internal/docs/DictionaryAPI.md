# \DictionaryAPI

All URIs are relative to *https://uapis.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDictionaryAudio**](DictionaryAPI.md#GetDictionaryAudio) | **Get** /dictionary/audio | 单词发音
[**GetDictionaryLookup**](DictionaryAPI.md#GetDictionaryLookup) | **Get** /dictionary/lookup | 单词查询



## GetDictionaryAudio

> *os.File GetDictionaryAudio(ctx).Word(word).Accent(accent).Execute()

单词发音



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
	word := "present" // string | 要发音的英文单词，长度不超过 64 个字符。
	accent := "accent_example" // string | 口音偏好：uk（英式）或 us（美式），默认 uk。 (optional) (default to "uk")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DictionaryAPI.GetDictionaryAudio(context.Background()).Word(word).Accent(accent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.GetDictionaryAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDictionaryAudio`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.GetDictionaryAudio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDictionaryAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **word** | **string** | 要发音的英文单词，长度不超过 64 个字符。 | 
 **accent** | **string** | 口音偏好：uk（英式）或 us（美式），默认 uk。 | [default to &quot;uk&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/mpeg, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDictionaryLookup

> GetDictionaryLookup200Response GetDictionaryLookup(ctx).Word(word).Lang(lang).Execute()

单词查询



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
	word := "present" // string | 要查询的英文单词，长度不超过 64 个字符。
	lang := "lang_example" // string | 目标语种。目前仅支持 en（默认）。 (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DictionaryAPI.GetDictionaryLookup(context.Background()).Word(word).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.GetDictionaryLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDictionaryLookup`: GetDictionaryLookup200Response
	fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.GetDictionaryLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDictionaryLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **word** | **string** | 要查询的英文单词，长度不超过 64 个字符。 | 
 **lang** | **string** | 目标语种。目前仅支持 en（默认）。 | [default to &quot;en&quot;]

### Return type

[**GetDictionaryLookup200Response**](GetDictionaryLookup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


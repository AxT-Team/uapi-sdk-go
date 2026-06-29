# \DailyAPI

All URIs are relative to *https://uapis.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyNewsImage**](DailyAPI.md#GetDailyNewsImage) | **Get** /daily/news-image | 每日新闻图
[**GetDailyWord**](DailyAPI.md#GetDailyWord) | **Get** /daily/word | 每日单词



## GetDailyNewsImage

> *os.File GetDailyNewsImage(ctx).Execute()

每日新闻图



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
	resp, r, err := apiClient.DailyAPI.GetDailyNewsImage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyAPI.GetDailyNewsImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDailyNewsImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DailyAPI.GetDailyNewsImage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyNewsImageRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDailyWord

> GetDailyWord200Response GetDailyWord(ctx).Lang(lang).Category(category).Count(count).Date(date).Seed(seed).Example(example).Phonetic(phonetic).Define(define).Execute()

每日单词



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
	lang := "en" // string | 语种，目前支持 en，默认 en。 (optional) (default to "en")
	category := "category_example" // string | 词库范围：all/cet4/cet6/ielts/toefl/gre，默认 all。 (optional) (default to "all")
	count := int32(3) // int32 | 返回数量，1-20，默认 1。 (optional) (default to 1)
	date := "date_example" // string | 日期，格式 YYYY-MM-DD，作为每日单词的种子基准。 (optional)
	seed := int32(56) // int32 | 固定种子，结果可复现；不可与 date 同时使用。 (optional)
	example := true // bool | 是否返回例句，默认 true。 (optional) (default to true)
	phonetic := true // bool | 是否返回音标，默认 true。 (optional) (default to true)
	define := true // bool | 是否为每个单词附带详细释义（音标发音、中英释义、词形、词组、近义词、双语例句），默认 false。 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DailyAPI.GetDailyWord(context.Background()).Lang(lang).Category(category).Count(count).Date(date).Seed(seed).Example(example).Phonetic(phonetic).Define(define).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DailyAPI.GetDailyWord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDailyWord`: GetDailyWord200Response
	fmt.Fprintf(os.Stdout, "Response from `DailyAPI.GetDailyWord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyWordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lang** | **string** | 语种，目前支持 en，默认 en。 | [default to &quot;en&quot;]
 **category** | **string** | 词库范围：all/cet4/cet6/ielts/toefl/gre，默认 all。 | [default to &quot;all&quot;]
 **count** | **int32** | 返回数量，1-20，默认 1。 | [default to 1]
 **date** | **string** | 日期，格式 YYYY-MM-DD，作为每日单词的种子基准。 | 
 **seed** | **int32** | 固定种子，结果可复现；不可与 date 同时使用。 | 
 **example** | **bool** | 是否返回例句，默认 true。 | [default to true]
 **phonetic** | **bool** | 是否返回音标，默认 true。 | [default to true]
 **define** | **bool** | 是否为每个单词附带详细释义（音标发音、中英释义、词形、词组、近义词、双语例句），默认 false。 | [default to false]

### Return type

[**GetDailyWord200Response**](GetDailyWord200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


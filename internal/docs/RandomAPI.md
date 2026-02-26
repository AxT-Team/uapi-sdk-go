# \RandomAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnswerbookAsk**](RandomAPI.md#GetAnswerbookAsk) | **Get** /answerbook/ask | 答案之书
[**GetRandomImage**](RandomAPI.md#GetRandomImage) | **Get** /random/image | 随机图片
[**GetRandomString**](RandomAPI.md#GetRandomString) | **Get** /random/string | 随机字符串
[**PostAnswerbookAsk**](RandomAPI.md#PostAnswerbookAsk) | **Post** /answerbook/ask | 答案之书 (POST)



## GetAnswerbookAsk

> GetAnswerbookAsk200Response GetAnswerbookAsk(ctx).Question(question).Execute()

答案之书



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
	question := "我今天会有好运吗？" // string | 你想要提问的问题。问题不能为空。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RandomAPI.GetAnswerbookAsk(context.Background()).Question(question).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomAPI.GetAnswerbookAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnswerbookAsk`: GetAnswerbookAsk200Response
	fmt.Fprintf(os.Stdout, "Response from `RandomAPI.GetAnswerbookAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnswerbookAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **question** | **string** | 你想要提问的问题。问题不能为空。 | 

### Return type

[**GetAnswerbookAsk200Response**](GetAnswerbookAsk200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRandomImage

> *os.File GetRandomImage(ctx).Category(category).Type_(type_).Execute()

随机图片



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
	category := "acg" // string | （可选）指定图片主类别。  **支持的主类别：** - `acg`（二次元动漫，UapiPro服务器） - `landscape`（风景图，外部图床） - `anime`（混合动漫） - `pc_wallpaper`（电脑壁纸，外部图床） - `mobile_wallpaper`（手机壁纸，外部图床） - `general_anime`（动漫图，外部图床） - `ai_drawing`（AI绘画，外部图床） - `bq`（表情包/趣图，UapiPro服务器） - `furry`（福瑞，UapiPro服务器）  > [!TIP] > 如果不指定，将从所有图片中随机抽取（不包含 `ikun` 和 `ai_drawing`）。  (optional)
	type_ := "pc" // string | （可选，仅UapiPro服务器图片支持）指定图片子类别。  - **bq**: `xiongmao`, `waiguoren`, `maomao`, `ikun`, `eciyuan` - **acg**: `pc`, `mb` - **furry**: `z4k`, `szs8k`, `s4k`, `4k`  > [!TIP] > 外部图床类别和 `anime` 混合类别不支持 `type` 参数。  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RandomAPI.GetRandomImage(context.Background()).Category(category).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomAPI.GetRandomImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRandomImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RandomAPI.GetRandomImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRandomImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | （可选）指定图片主类别。  **支持的主类别：** - &#x60;acg&#x60;（二次元动漫，UapiPro服务器） - &#x60;landscape&#x60;（风景图，外部图床） - &#x60;anime&#x60;（混合动漫） - &#x60;pc_wallpaper&#x60;（电脑壁纸，外部图床） - &#x60;mobile_wallpaper&#x60;（手机壁纸，外部图床） - &#x60;general_anime&#x60;（动漫图，外部图床） - &#x60;ai_drawing&#x60;（AI绘画，外部图床） - &#x60;bq&#x60;（表情包/趣图，UapiPro服务器） - &#x60;furry&#x60;（福瑞，UapiPro服务器）  &gt; [!TIP] &gt; 如果不指定，将从所有图片中随机抽取（不包含 &#x60;ikun&#x60; 和 &#x60;ai_drawing&#x60;）。  | 
 **type_** | **string** | （可选，仅UapiPro服务器图片支持）指定图片子类别。  - **bq**: &#x60;xiongmao&#x60;, &#x60;waiguoren&#x60;, &#x60;maomao&#x60;, &#x60;ikun&#x60;, &#x60;eciyuan&#x60; - **acg**: &#x60;pc&#x60;, &#x60;mb&#x60; - **furry**: &#x60;z4k&#x60;, &#x60;szs8k&#x60;, &#x60;s4k&#x60;, &#x60;4k&#x60;  &gt; [!TIP] &gt; 外部图床类别和 &#x60;anime&#x60; 混合类别不支持 &#x60;type&#x60; 参数。  | 

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


## GetRandomString

> GetRandomString200Response GetRandomString(ctx).Length(length).Type_(type_).Execute()

随机字符串



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
	length := int32(32) // int32 | 你希望生成的字符串的长度。有效范围是 1 到 1024。 (optional) (default to 16)
	type_ := "alphanumeric" // string | 指定构成字符串的字符类型。 (optional) (default to "alphanumeric")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RandomAPI.GetRandomString(context.Background()).Length(length).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomAPI.GetRandomString``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRandomString`: GetRandomString200Response
	fmt.Fprintf(os.Stdout, "Response from `RandomAPI.GetRandomString`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRandomStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **int32** | 你希望生成的字符串的长度。有效范围是 1 到 1024。 | [default to 16]
 **type_** | **string** | 指定构成字符串的字符类型。 | [default to &quot;alphanumeric&quot;]

### Return type

[**GetRandomString200Response**](GetRandomString200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAnswerbookAsk

> PostAnswerbookAsk200Response PostAnswerbookAsk(ctx).PostAnswerbookAskRequest(postAnswerbookAskRequest).Execute()

答案之书 (POST)



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
	postAnswerbookAskRequest := *openapiclient.NewPostAnswerbookAskRequest("我应该接受这份工作吗？") // PostAnswerbookAskRequest | 包含问题的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RandomAPI.PostAnswerbookAsk(context.Background()).PostAnswerbookAskRequest(postAnswerbookAskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomAPI.PostAnswerbookAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAnswerbookAsk`: PostAnswerbookAsk200Response
	fmt.Fprintf(os.Stdout, "Response from `RandomAPI.PostAnswerbookAsk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAnswerbookAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postAnswerbookAskRequest** | [**PostAnswerbookAskRequest**](PostAnswerbookAskRequest.md) | 包含问题的JSON对象 | 

### Return type

[**PostAnswerbookAsk200Response**](PostAnswerbookAsk200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


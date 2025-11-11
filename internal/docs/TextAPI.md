# \TextAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTextMd5**](TextAPI.md#GetTextMd5) | **Get** /text/md5 | 计算文本的MD5哈希值(GET)
[**PostTextAesDecrypt**](TextAPI.md#PostTextAesDecrypt) | **Post** /text/aes/decrypt | 使用AES算法解密文本
[**PostTextAesEncrypt**](TextAPI.md#PostTextAesEncrypt) | **Post** /text/aes/encrypt | 使用AES算法加密文本
[**PostTextAnalyze**](TextAPI.md#PostTextAnalyze) | **Post** /text/analyze | 多维度分析文本内容
[**PostTextBase64Decode**](TextAPI.md#PostTextBase64Decode) | **Post** /text/base64/decode | 解码Base64编码的文本
[**PostTextBase64Encode**](TextAPI.md#PostTextBase64Encode) | **Post** /text/base64/encode | 将文本进行Base64编码
[**PostTextMd5**](TextAPI.md#PostTextMd5) | **Post** /text/md5 | 计算文本的MD5哈希值 (POST)
[**PostTextMd5Verify**](TextAPI.md#PostTextMd5Verify) | **Post** /text/md5/verify | 校验MD5哈希值



## GetTextMd5

> GetTextMd5200Response GetTextMd5(ctx).Text(text).Execute()

计算文本的MD5哈希值(GET)



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
	text := "hello world" // string | 需要计算哈希值的文本

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.GetTextMd5(context.Background()).Text(text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.GetTextMd5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTextMd5`: GetTextMd5200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.GetTextMd5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTextMd5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | 需要计算哈希值的文本 | 

### Return type

[**GetTextMd5200Response**](GetTextMd5200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextAesDecrypt

> PostTextAesDecrypt200Response PostTextAesDecrypt(ctx).PostTextAesDecryptRequest(postTextAesDecryptRequest).Execute()

使用AES算法解密文本



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
	postTextAesDecryptRequest := *openapiclient.NewPostTextAesDecryptRequest("a-secret-key-123", "uyzVKczxZi3HdoGfeuaAt4F2/20WSmwFzIhJWMmDIaxeu97nLqbsX3wdp+NnRw==", "1234567890abcdef") // PostTextAesDecryptRequest | 包含待解密文本 'text'、密钥 'key' 和随机数 'nonce' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextAesDecrypt(context.Background()).PostTextAesDecryptRequest(postTextAesDecryptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextAesDecrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextAesDecrypt`: PostTextAesDecrypt200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextAesDecrypt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextAesDecryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextAesDecryptRequest** | [**PostTextAesDecryptRequest**](PostTextAesDecryptRequest.md) | 包含待解密文本 &#39;text&#39;、密钥 &#39;key&#39; 和随机数 &#39;nonce&#39; 的JSON对象 | 

### Return type

[**PostTextAesDecrypt200Response**](PostTextAesDecrypt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextAesEncrypt

> PostTextAesEncrypt200Response PostTextAesEncrypt(ctx).PostTextAesEncryptRequest(postTextAesEncryptRequest).Execute()

使用AES算法加密文本



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
	postTextAesEncryptRequest := *openapiclient.NewPostTextAesEncryptRequest("a-secret-key-123", "top secret message") // PostTextAesEncryptRequest | 包含待加密文本 'text' 和密钥 'key' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextAesEncrypt(context.Background()).PostTextAesEncryptRequest(postTextAesEncryptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextAesEncrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextAesEncrypt`: PostTextAesEncrypt200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextAesEncrypt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextAesEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextAesEncryptRequest** | [**PostTextAesEncryptRequest**](PostTextAesEncryptRequest.md) | 包含待加密文本 &#39;text&#39; 和密钥 &#39;key&#39; 的JSON对象 | 

### Return type

[**PostTextAesEncrypt200Response**](PostTextAesEncrypt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextAnalyze

> PostTextAnalyze200Response PostTextAnalyze(ctx).PostTextAnalyzeRequest(postTextAnalyzeRequest).Execute()

多维度分析文本内容



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
	postTextAnalyzeRequest := *openapiclient.NewPostTextAnalyzeRequest("Hello world.
This is a sample sentence. It has multiple lines and words.") // PostTextAnalyzeRequest | 包含待分析文本 'text' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextAnalyze(context.Background()).PostTextAnalyzeRequest(postTextAnalyzeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextAnalyze`: PostTextAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextAnalyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextAnalyzeRequest** | [**PostTextAnalyzeRequest**](PostTextAnalyzeRequest.md) | 包含待分析文本 &#39;text&#39; 的JSON对象 | 

### Return type

[**PostTextAnalyze200Response**](PostTextAnalyze200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextBase64Decode

> PostTextBase64Decode200Response PostTextBase64Decode(ctx).PostTextBase64DecodeRequest(postTextBase64DecodeRequest).Execute()

解码Base64编码的文本



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
	postTextBase64DecodeRequest := *openapiclient.NewPostTextBase64DecodeRequest("aGVsbG8gd29ybGQ=") // PostTextBase64DecodeRequest | 包含待解码文本 'text' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextBase64Decode(context.Background()).PostTextBase64DecodeRequest(postTextBase64DecodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextBase64Decode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextBase64Decode`: PostTextBase64Decode200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextBase64Decode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextBase64DecodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextBase64DecodeRequest** | [**PostTextBase64DecodeRequest**](PostTextBase64DecodeRequest.md) | 包含待解码文本 &#39;text&#39; 的JSON对象 | 

### Return type

[**PostTextBase64Decode200Response**](PostTextBase64Decode200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextBase64Encode

> PostTextBase64Encode200Response PostTextBase64Encode(ctx).PostTextBase64EncodeRequest(postTextBase64EncodeRequest).Execute()

将文本进行Base64编码



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
	postTextBase64EncodeRequest := *openapiclient.NewPostTextBase64EncodeRequest("hello world") // PostTextBase64EncodeRequest | 包含待编码文本 'text' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextBase64Encode(context.Background()).PostTextBase64EncodeRequest(postTextBase64EncodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextBase64Encode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextBase64Encode`: PostTextBase64Encode200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextBase64Encode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextBase64EncodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextBase64EncodeRequest** | [**PostTextBase64EncodeRequest**](PostTextBase64EncodeRequest.md) | 包含待编码文本 &#39;text&#39; 的JSON对象 | 

### Return type

[**PostTextBase64Encode200Response**](PostTextBase64Encode200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextMd5

> GetTextMd5200Response PostTextMd5(ctx).PostTextMd5Request(postTextMd5Request).Execute()

计算文本的MD5哈希值 (POST)



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
	postTextMd5Request := *openapiclient.NewPostTextMd5Request("hello world") // PostTextMd5Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextMd5(context.Background()).PostTextMd5Request(postTextMd5Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextMd5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextMd5`: GetTextMd5200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextMd5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextMd5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextMd5Request** | [**PostTextMd5Request**](PostTextMd5Request.md) |  | 

### Return type

[**GetTextMd5200Response**](GetTextMd5200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextMd5Verify

> PostTextMd5Verify200Response PostTextMd5Verify(ctx).PostTextMd5VerifyRequest(postTextMd5VerifyRequest).Execute()

校验MD5哈希值



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
	postTextMd5VerifyRequest := *openapiclient.NewPostTextMd5VerifyRequest("5d41402abc4b2a76b9719d911017c592", "hello world") // PostTextMd5VerifyRequest | 包含待校验文本 'text' 和哈希值 'hash' 的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextMd5Verify(context.Background()).PostTextMd5VerifyRequest(postTextMd5VerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextMd5Verify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextMd5Verify`: PostTextMd5Verify200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextMd5Verify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextMd5VerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextMd5VerifyRequest** | [**PostTextMd5VerifyRequest**](PostTextMd5VerifyRequest.md) | 包含待校验文本 &#39;text&#39; 和哈希值 &#39;hash&#39; 的JSON对象 | 

### Return type

[**PostTextMd5Verify200Response**](PostTextMd5Verify200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


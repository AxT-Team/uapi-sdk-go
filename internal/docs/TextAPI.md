# \TextAPI

All URIs are relative to *https://uapis.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTextMd5**](TextAPI.md#GetTextMd5) | **Get** /text/md5 | MD5 哈希
[**PostTextAesDecrypt**](TextAPI.md#PostTextAesDecrypt) | **Post** /text/aes/decrypt | AES 解密
[**PostTextAesDecryptAdvanced**](TextAPI.md#PostTextAesDecryptAdvanced) | **Post** /text/aes/decrypt-advanced | AES高级解密
[**PostTextAesEncrypt**](TextAPI.md#PostTextAesEncrypt) | **Post** /text/aes/encrypt | AES 加密
[**PostTextAesEncryptAdvanced**](TextAPI.md#PostTextAesEncryptAdvanced) | **Post** /text/aes/encrypt-advanced | AES高级加密
[**PostTextAnalyze**](TextAPI.md#PostTextAnalyze) | **Post** /text/analyze | 文本分析
[**PostTextBase64Decode**](TextAPI.md#PostTextBase64Decode) | **Post** /text/base64/decode | Base64 解码
[**PostTextBase64Encode**](TextAPI.md#PostTextBase64Encode) | **Post** /text/base64/encode | Base64 编码
[**PostTextConvert**](TextAPI.md#PostTextConvert) | **Post** /text/convert | 格式转换
[**PostTextMarkdownToHtml**](TextAPI.md#PostTextMarkdownToHtml) | **Post** /text/markdown-to-html | Markdown 转 HTML
[**PostTextMarkdownToPdf**](TextAPI.md#PostTextMarkdownToPdf) | **Post** /text/markdown-to-pdf | Markdown 转 PDF
[**PostTextMd5**](TextAPI.md#PostTextMd5) | **Post** /text/md5 | MD5 哈希 (POST)
[**PostTextMd5Verify**](TextAPI.md#PostTextMd5Verify) | **Post** /text/md5/verify | MD5 校验



## GetTextMd5

> GetTextMd5200Response GetTextMd5(ctx).Text(text).Execute()

MD5 哈希



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

AES 解密



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
	postTextAesDecryptRequest := *openapiclient.NewPostTextAesDecryptRequest("a-secret-key-123", "uyzVKczxZi3HdoGfeuaAt4F2/20WSmwFzIhJWMmDIaxeu97nLqbsX3wdp+NnRw==") // PostTextAesDecryptRequest | 

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
 **postTextAesDecryptRequest** | [**PostTextAesDecryptRequest**](PostTextAesDecryptRequest.md) |  | 

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


## PostTextAesDecryptAdvanced

> PostTextAesDecryptAdvanced200Response PostTextAesDecryptAdvanced(ctx).PostTextAesDecryptAdvancedRequest(postTextAesDecryptAdvancedRequest).Execute()

AES高级解密



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
	postTextAesDecryptAdvancedRequest := *openapiclient.NewPostTextAesDecryptAdvancedRequest("my-super-secret-key", "GCM", "68vWkaxJPg1vx0LWJONpEfYdvW3Wz7V5uXiYg0WWfGJHIZWBmVVghHg=") // PostTextAesDecryptAdvancedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextAesDecryptAdvanced(context.Background()).PostTextAesDecryptAdvancedRequest(postTextAesDecryptAdvancedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextAesDecryptAdvanced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextAesDecryptAdvanced`: PostTextAesDecryptAdvanced200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextAesDecryptAdvanced`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextAesDecryptAdvancedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextAesDecryptAdvancedRequest** | [**PostTextAesDecryptAdvancedRequest**](PostTextAesDecryptAdvancedRequest.md) |  | 

### Return type

[**PostTextAesDecryptAdvanced200Response**](PostTextAesDecryptAdvanced200Response.md)

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

AES 加密



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
	postTextAesEncryptRequest := *openapiclient.NewPostTextAesEncryptRequest("a-secret-key-123", "这是一段需要加密的消息") // PostTextAesEncryptRequest | 

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
 **postTextAesEncryptRequest** | [**PostTextAesEncryptRequest**](PostTextAesEncryptRequest.md) |  | 

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


## PostTextAesEncryptAdvanced

> PostTextAesEncryptAdvanced200Response PostTextAesEncryptAdvanced(ctx).PostTextAesEncryptAdvancedRequest(postTextAesEncryptAdvancedRequest).Execute()

AES高级加密



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
	postTextAesEncryptAdvancedRequest := *openapiclient.NewPostTextAesEncryptAdvancedRequest("my-super-secret-key", "Hello, World! 你好世界！") // PostTextAesEncryptAdvancedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextAesEncryptAdvanced(context.Background()).PostTextAesEncryptAdvancedRequest(postTextAesEncryptAdvancedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextAesEncryptAdvanced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextAesEncryptAdvanced`: PostTextAesEncryptAdvanced200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextAesEncryptAdvanced`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextAesEncryptAdvancedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextAesEncryptAdvancedRequest** | [**PostTextAesEncryptAdvancedRequest**](PostTextAesEncryptAdvancedRequest.md) |  | 

### Return type

[**PostTextAesEncryptAdvanced200Response**](PostTextAesEncryptAdvanced200Response.md)

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

文本分析



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
This is a sample sentence. It has multiple lines and words.") // PostTextAnalyzeRequest | 

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
 **postTextAnalyzeRequest** | [**PostTextAnalyzeRequest**](PostTextAnalyzeRequest.md) |  | 

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

Base64 解码



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
	postTextBase64DecodeRequest := *openapiclient.NewPostTextBase64DecodeRequest("aGVsbG8gd29ybGQ=") // PostTextBase64DecodeRequest | 

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
 **postTextBase64DecodeRequest** | [**PostTextBase64DecodeRequest**](PostTextBase64DecodeRequest.md) |  | 

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

Base64 编码



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
	postTextBase64EncodeRequest := *openapiclient.NewPostTextBase64EncodeRequest("hello world") // PostTextBase64EncodeRequest | 

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
 **postTextBase64EncodeRequest** | [**PostTextBase64EncodeRequest**](PostTextBase64EncodeRequest.md) |  | 

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


## PostTextConvert

> PostTextConvert200Response PostTextConvert(ctx).PostTextConvertRequest(postTextConvertRequest).Execute()

格式转换



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
	postTextConvertRequest := *openapiclient.NewPostTextConvertRequest("plain", "hello world", "base64") // PostTextConvertRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextConvert(context.Background()).PostTextConvertRequest(postTextConvertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextConvert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextConvert`: PostTextConvert200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextConvert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextConvertRequest** | [**PostTextConvertRequest**](PostTextConvertRequest.md) |  | 

### Return type

[**PostTextConvert200Response**](PostTextConvert200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextMarkdownToHtml

> PostTextMarkdownToHtml200Response PostTextMarkdownToHtml(ctx).PostTextMarkdownToHtmlRequest(postTextMarkdownToHtmlRequest).Execute()

Markdown 转 HTML



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
	postTextMarkdownToHtmlRequest := *openapiclient.NewPostTextMarkdownToHtmlRequest("# 咖啡

**咖啡**（英语：*coffee*）是指将咖啡植物的种子（即[咖啡豆](https://baike.baidu.com/item/%E5%92%96%E5%95%A1%E8%B1%86/13579425)）经过烘焙磨粉后冲泡溶解制成的饮料，是世界上流行范围最为广泛的软性饮料之一。

> 野生咖啡原产于非洲和亚洲的热带地区，现今，未经烘焙的生咖啡豆作为世界上最大的出口农产品，有超过70个国家种植咖啡树。

![咖啡图片](https://images.unsplash.com/photo-1534234757579-8ad69d218ad4?q=80&w=1170&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D)

## 营养价值 (每100克)

| 营养成分 | 含量 | 每日参考摄入量 (DRI) |
| :--- | :--- | :--- |
| 热量 | 2 kJ (0.48 kcal) | - |
| 咖啡因 | 40 mg | - |
| 核黄素 (维生素B2) | 0.076 mg | 6% |
| 镁 | 3 mg | 1% |
| 钾 | 49 mg | 1% |

## 咖啡制作工序

咖啡从果实到杯中的饮品，需要经过严谨的步骤：

- [x] **采收**：有机械采收和人工采收两种方式。
- [x] **加工**：去除水果层，干燥到合适含水量（水洗法、日晒法、蜜处理等）。
- [ ] **烘焙**：生豆水分释放、体积膨胀。分为浅焙、中焙、深焙。

## 冲煮与萃取

所有的咖啡都是由磨好的咖啡粉和热水制出的。常见的分类包括：

1. **滴滤法**：如手冲咖啡。水与咖啡粉接触的机会只有一次。
2. **压力法**：如意式浓缩咖啡（Espresso），以8～9个大气压的力道萃取。
3. **浸泡法**：如冷泡咖啡，将研磨后的咖啡豆置于盛有冷水的玻璃瓶中静置。

## 相关技术协议

在计算机网络中，甚至有一个专门用于控制咖啡壶的协议（超文本咖啡壶控制协议，HTCPCP）：

```http
BREW /pot-1/coffee HTCPCP/1.0
Host: coffee-pot.local
Content-Type: message/coffeepot
Accept-Additions: cream, whole-milk
```") // PostTextMarkdownToHtmlRequest | 请求体使用 `application/json`。`text` 必填；`format` 和 `sanitize` 可选。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextMarkdownToHtml(context.Background()).PostTextMarkdownToHtmlRequest(postTextMarkdownToHtmlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextMarkdownToHtml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextMarkdownToHtml`: PostTextMarkdownToHtml200Response
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextMarkdownToHtml`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextMarkdownToHtmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextMarkdownToHtmlRequest** | [**PostTextMarkdownToHtmlRequest**](PostTextMarkdownToHtmlRequest.md) | 请求体使用 &#x60;application/json&#x60;。&#x60;text&#x60; 必填；&#x60;format&#x60; 和 &#x60;sanitize&#x60; 可选。 | 

### Return type

[**PostTextMarkdownToHtml200Response**](PostTextMarkdownToHtml200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextMarkdownToPdf

> *os.File PostTextMarkdownToPdf(ctx).PostTextMarkdownToPdfRequest(postTextMarkdownToPdfRequest).Execute()

Markdown 转 PDF



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
	postTextMarkdownToPdfRequest := *openapiclient.NewPostTextMarkdownToPdfRequest("# 咖啡

**咖啡**（英语：*coffee*）是指将咖啡植物的种子（即[咖啡豆](https://baike.baidu.com/item/%E5%92%96%E5%95%A1%E8%B1%86/13579425)）经过烘焙磨粉后冲泡溶解制成的饮料，是世界上流行范围最为广泛的软性饮料之一。

> 野生咖啡原产于非洲和亚洲的热带地区，现今，未经烘焙的生咖啡豆作为世界上最大的出口农产品，有超过70个国家种植咖啡树。

![咖啡图片](https://images.unsplash.com/photo-1534234757579-8ad69d218ad4?q=80&w=1170&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D)

## 营养价值 (每100克)

| 营养成分 | 含量 | 每日参考摄入量 (DRI) |
| :--- | :--- | :--- |
| 热量 | 2 kJ (0.48 kcal) | - |
| 咖啡因 | 40 mg | - |
| 核黄素 (维生素B2) | 0.076 mg | 6% |
| 镁 | 3 mg | 1% |
| 钾 | 49 mg | 1% |

## 咖啡制作工序

咖啡从果实到杯中的饮品，需要经过严谨的步骤：

- [x] **采收**：有机械采收和人工采收两种方式。
- [x] **加工**：去除水果层，干燥到合适含水量（水洗法、日晒法、蜜处理等）。
- [ ] **烘焙**：生豆水分释放、体积膨胀。分为浅焙、中焙、深焙。

## 冲煮与萃取

所有的咖啡都是由磨好的咖啡粉和热水制出的。常见的分类包括：

1. **滴滤法**：如手冲咖啡。水与咖啡粉接触的机会只有一次。
2. **压力法**：如意式浓缩咖啡（Espresso），以8～9个大气压的力道萃取。
3. **浸泡法**：如冷泡咖啡，将研磨后的咖啡豆置于盛有冷水的玻璃瓶中静置。

## 相关技术协议

在计算机网络中，甚至有一个专门用于控制咖啡壶的协议（超文本咖啡壶控制协议，HTCPCP）：

```http
BREW /pot-1/coffee HTCPCP/1.0
Host: coffee-pot.local
Content-Type: message/coffeepot
Accept-Additions: cream, whole-milk
```") // PostTextMarkdownToPdfRequest | 请求体使用 `application/json`。`text` 必填，`theme` 和 `paper_size` 可选。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextAPI.PostTextMarkdownToPdf(context.Background()).PostTextMarkdownToPdfRequest(postTextMarkdownToPdfRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextAPI.PostTextMarkdownToPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTextMarkdownToPdf`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TextAPI.PostTextMarkdownToPdf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextMarkdownToPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTextMarkdownToPdfRequest** | [**PostTextMarkdownToPdfRequest**](PostTextMarkdownToPdfRequest.md) | 请求体使用 &#x60;application/json&#x60;。&#x60;text&#x60; 必填，&#x60;theme&#x60; 和 &#x60;paper_size&#x60; 可选。 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextMd5

> GetTextMd5200Response PostTextMd5(ctx).PostTextMd5Request(postTextMd5Request).Execute()

MD5 哈希 (POST)



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

MD5 校验



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
	postTextMd5VerifyRequest := *openapiclient.NewPostTextMd5VerifyRequest("5d41402abc4b2a76b9719d911017c592", "hello world") // PostTextMd5VerifyRequest | 

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
 **postTextMd5VerifyRequest** | [**PostTextMd5VerifyRequest**](PostTextMd5VerifyRequest.md) |  | 

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


# \AIGCAPI

All URIs are relative to *https://uapis.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostWatermarkDecode**](AIGCAPI.md#PostWatermarkDecode) | **Post** /watermark/decode | 提取图片隐水印
[**PostWatermarkEmbed**](AIGCAPI.md#PostWatermarkEmbed) | **Post** /watermark/embed | 添加图片隐水印
[**PostWatermarkLabel**](AIGCAPI.md#PostWatermarkLabel) | **Post** /watermark/label | 添加 AI 生成内容标识
[**PostWatermarkProducerCode**](AIGCAPI.md#PostWatermarkProducerCode) | **Post** /watermark/producer-code | 生成 AIGC 服务提供者编码



## PostWatermarkDecode

> PostWatermarkDecode200Response PostWatermarkDecode(ctx).Ecc(ecc).File(file).ImageBase64(imageBase64).ModelType(modelType).Url(url).Execute()

提取图片隐水印



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
	ecc := "ecc_example" // string | 纠错强度，必须和嵌入时填的一致，否则无法正确提取。[查看各档位](#enum-list) (optional) (default to "BCH_4")
	file := os.NewFile(1234, "some_file") // *os.File | 要提取水印的图片文件，支持 PNG、JPEG、WebP。 (optional)
	imageBase64 := "imageBase64_example" // string | 图片的 Base64 编码，可携带或省略 data: 前缀。 (optional)
	modelType := "modelType_example" // string | 水印档位，必须和嵌入时用的一致，否则无法正确提取。[查看各档位](#enum-list) (optional) (default to "B")
	url := "url_example" // string | 图片链接，需确保公网可直接访问。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIGCAPI.PostWatermarkDecode(context.Background()).Ecc(ecc).File(file).ImageBase64(imageBase64).ModelType(modelType).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGCAPI.PostWatermarkDecode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWatermarkDecode`: PostWatermarkDecode200Response
	fmt.Fprintf(os.Stdout, "Response from `AIGCAPI.PostWatermarkDecode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWatermarkDecodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ecc** | **string** | 纠错强度，必须和嵌入时填的一致，否则无法正确提取。[查看各档位](#enum-list) | [default to &quot;BCH_4&quot;]
 **file** | ***os.File** | 要提取水印的图片文件，支持 PNG、JPEG、WebP。 | 
 **imageBase64** | **string** | 图片的 Base64 编码，可携带或省略 data: 前缀。 | 
 **modelType** | **string** | 水印档位，必须和嵌入时用的一致，否则无法正确提取。[查看各档位](#enum-list) | [default to &quot;B&quot;]
 **url** | **string** | 图片链接，需确保公网可直接访问。 | 

### Return type

[**PostWatermarkDecode200Response**](PostWatermarkDecode200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWatermarkEmbed

> PostWatermarkEmbed200Response PostWatermarkEmbed(ctx).Payload(payload).Ecc(ecc).File(file).ImageBase64(imageBase64).JpegQuality(jpegQuality).ModelType(modelType).OutFormat(outFormat).Strength(strength).Url(url).Execute()

添加图片隐水印



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
	payload := "payload_example" // string | 需要嵌入图片的隐形标识内容。
	ecc := "ecc_example" // string | 纠错强度，决定水印能抗多少损坏、最多能嵌入多少字符：纠错越强，图片被压缩、裁剪后越容易读回，但能嵌入的字符越少。不填默认 `BCH_4`。[查看各档位](#enum-list) (optional) (default to "BCH_4")
	file := os.NewFile(1234, "some_file") // *os.File | 要加水印的图片文件，支持 PNG、JPEG、WebP。 (optional)
	imageBase64 := "imageBase64_example" // string | 图片的 Base64 编码，可携带或省略 data: 前缀。 (optional)
	jpegQuality := int32(56) // int32 | 输出 JPEG 时的图像质量，范围 1 到 100。 (optional)
	modelType := "modelType_example" // string | 水印档位，在稳健性和画质之间取舍。不填默认 `B`。[查看各档位](#enum-list) (optional) (default to "B")
	outFormat := "outFormat_example" // string | 输出的图片格式。不填则默认保持与原图一致。 (optional)
	strength := float32(8.14) // float32 | 水印写入强度，默认 `1.0`。调高更不容易被压缩、转发破坏，但更可能被肉眼看出；调低更隐蔽，但抗损坏能力下降。 (optional) (default to 1)
	url := "url_example" // string | 图片链接，需确保公网可直接访问。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIGCAPI.PostWatermarkEmbed(context.Background()).Payload(payload).Ecc(ecc).File(file).ImageBase64(imageBase64).JpegQuality(jpegQuality).ModelType(modelType).OutFormat(outFormat).Strength(strength).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGCAPI.PostWatermarkEmbed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWatermarkEmbed`: PostWatermarkEmbed200Response
	fmt.Fprintf(os.Stdout, "Response from `AIGCAPI.PostWatermarkEmbed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWatermarkEmbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | **string** | 需要嵌入图片的隐形标识内容。 | 
 **ecc** | **string** | 纠错强度，决定水印能抗多少损坏、最多能嵌入多少字符：纠错越强，图片被压缩、裁剪后越容易读回，但能嵌入的字符越少。不填默认 &#x60;BCH_4&#x60;。[查看各档位](#enum-list) | [default to &quot;BCH_4&quot;]
 **file** | ***os.File** | 要加水印的图片文件，支持 PNG、JPEG、WebP。 | 
 **imageBase64** | **string** | 图片的 Base64 编码，可携带或省略 data: 前缀。 | 
 **jpegQuality** | **int32** | 输出 JPEG 时的图像质量，范围 1 到 100。 | 
 **modelType** | **string** | 水印档位，在稳健性和画质之间取舍。不填默认 &#x60;B&#x60;。[查看各档位](#enum-list) | [default to &quot;B&quot;]
 **outFormat** | **string** | 输出的图片格式。不填则默认保持与原图一致。 | 
 **strength** | **float32** | 水印写入强度，默认 &#x60;1.0&#x60;。调高更不容易被压缩、转发破坏，但更可能被肉眼看出；调低更隐蔽，但抗损坏能力下降。 | [default to 1]
 **url** | **string** | 图片链接，需确保公网可直接访问。 | 

### Return type

[**PostWatermarkEmbed200Response**](PostWatermarkEmbed200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWatermarkLabel

> PostWatermarkLabel200Response PostWatermarkLabel(ctx).ContentProducer(contentProducer).ContentPropagator(contentPropagator).EmbedWatermark(embedWatermark).ExplicitHeightRatio(explicitHeightRatio).ExplicitLabel(explicitLabel).ExplicitPosition(explicitPosition).ExplicitText(explicitText).File(file).ImageBase64(imageBase64).JpegQuality(jpegQuality).Label(label).OutFormat(outFormat).ProduceId(produceId).PropagateId(propagateId).SkipMetadata(skipMetadata).Url(url).WatermarkPayload(watermarkPayload).Execute()

添加 AI 生成内容标识



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
	contentProducer := "contentProducer_example" // string | 必需：生成此图片的服务提供者编码（27 位）。
	contentPropagator := "contentPropagator_example" // string | 负责内容传播的服务提供者编码（27 位，可选）。 (optional)
	embedWatermark := true // bool | 是否额外注入抗压缩的隐形水印。默认不开启。 (optional) (default to false)
	explicitHeightRatio := float32(8.14) // float32 | 角标文字高度占画面短边的比例。低于 0.05 会自动补偿至国标下限要求。 (optional)
	explicitLabel := true // bool | 是否叠加可见的角标文字标识。默认不开启。 (optional) (default to false)
	explicitPosition := "explicitPosition_example" // string | 角标所处的相对位置，默认为右下角。 (optional)
	explicitText := "explicitText_example" // string | 角标显示的具体文案，默认为“AI 生成”。 (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 待处理的图片文件，支持 PNG、JPEG、WebP。 (optional)
	imageBase64 := "imageBase64_example" // string | 图片的 Base64 编码，可携带或省略 data: 前缀。 (optional)
	jpegQuality := int32(56) // int32 | 输出 JPEG 时的图像质量，范围 1 到 100。 (optional)
	label := "label_example" // string | 生成场景分类：1 代表 AI 生成合成，2 代表人机协同，3 代表其他情况。默认取值为 1。 (optional)
	outFormat := "outFormat_example" // string | 输出的图片格式。不填则默认保持与原图一致。 (optional)
	produceId := "produceId_example" // string | 服务侧内部生成的内容编号（可选）。 (optional)
	propagateId := "propagateId_example" // string | 传播方侧的内容编号（可选）。 (optional)
	skipMetadata := true // bool | 是否跳过写入元数据标识。若设置为 true，则必须开启另外两项中的至少一项。 (optional) (default to false)
	url := "url_example" // string | 图片链接，需确保公网可直接访问。 (optional)
	watermarkPayload := "watermarkPayload_example" // string | 隐形水印中所记载的标识内容。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIGCAPI.PostWatermarkLabel(context.Background()).ContentProducer(contentProducer).ContentPropagator(contentPropagator).EmbedWatermark(embedWatermark).ExplicitHeightRatio(explicitHeightRatio).ExplicitLabel(explicitLabel).ExplicitPosition(explicitPosition).ExplicitText(explicitText).File(file).ImageBase64(imageBase64).JpegQuality(jpegQuality).Label(label).OutFormat(outFormat).ProduceId(produceId).PropagateId(propagateId).SkipMetadata(skipMetadata).Url(url).WatermarkPayload(watermarkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGCAPI.PostWatermarkLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWatermarkLabel`: PostWatermarkLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `AIGCAPI.PostWatermarkLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWatermarkLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentProducer** | **string** | 必需：生成此图片的服务提供者编码（27 位）。 | 
 **contentPropagator** | **string** | 负责内容传播的服务提供者编码（27 位，可选）。 | 
 **embedWatermark** | **bool** | 是否额外注入抗压缩的隐形水印。默认不开启。 | [default to false]
 **explicitHeightRatio** | **float32** | 角标文字高度占画面短边的比例。低于 0.05 会自动补偿至国标下限要求。 | 
 **explicitLabel** | **bool** | 是否叠加可见的角标文字标识。默认不开启。 | [default to false]
 **explicitPosition** | **string** | 角标所处的相对位置，默认为右下角。 | 
 **explicitText** | **string** | 角标显示的具体文案，默认为“AI 生成”。 | 
 **file** | ***os.File** | 待处理的图片文件，支持 PNG、JPEG、WebP。 | 
 **imageBase64** | **string** | 图片的 Base64 编码，可携带或省略 data: 前缀。 | 
 **jpegQuality** | **int32** | 输出 JPEG 时的图像质量，范围 1 到 100。 | 
 **label** | **string** | 生成场景分类：1 代表 AI 生成合成，2 代表人机协同，3 代表其他情况。默认取值为 1。 | 
 **outFormat** | **string** | 输出的图片格式。不填则默认保持与原图一致。 | 
 **produceId** | **string** | 服务侧内部生成的内容编号（可选）。 | 
 **propagateId** | **string** | 传播方侧的内容编号（可选）。 | 
 **skipMetadata** | **bool** | 是否跳过写入元数据标识。若设置为 true，则必须开启另外两项中的至少一项。 | [default to false]
 **url** | **string** | 图片链接，需确保公网可直接访问。 | 
 **watermarkPayload** | **string** | 隐形水印中所记载的标识内容。 | 

### Return type

[**PostWatermarkLabel200Response**](PostWatermarkLabel200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWatermarkProducerCode

> PostWatermarkProducerCode200Response PostWatermarkProducerCode(ctx).PostWatermarkProducerCodeRequest(postWatermarkProducerCodeRequest).Execute()

生成 AIGC 服务提供者编码



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
	postWatermarkProducerCodeRequest := *openapiclient.NewPostWatermarkProducerCodeRequest() // PostWatermarkProducerCodeRequest | 生成所需的身份信息，或用于校验的 27 位现成编码。支持 application/json。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIGCAPI.PostWatermarkProducerCode(context.Background()).PostWatermarkProducerCodeRequest(postWatermarkProducerCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGCAPI.PostWatermarkProducerCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWatermarkProducerCode`: PostWatermarkProducerCode200Response
	fmt.Fprintf(os.Stdout, "Response from `AIGCAPI.PostWatermarkProducerCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWatermarkProducerCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postWatermarkProducerCodeRequest** | [**PostWatermarkProducerCodeRequest**](PostWatermarkProducerCodeRequest.md) | 生成所需的身份信息，或用于校验的 27 位现成编码。支持 application/json。 | 

### Return type

[**PostWatermarkProducerCode200Response**](PostWatermarkProducerCode200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ImageAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvatarGravatar**](ImageAPI.md#GetAvatarGravatar) | **Get** /avatar/gravatar | 获取Gravatar头像
[**GetImageBingDaily**](ImageAPI.md#GetImageBingDaily) | **Get** /image/bing-daily | 必应壁纸
[**GetImageMotou**](ImageAPI.md#GetImageMotou) | **Get** /image/motou | 生成摸摸头GIF (QQ号)
[**GetImageQrcode**](ImageAPI.md#GetImageQrcode) | **Get** /image/qrcode | 生成二维码
[**GetImageTobase64**](ImageAPI.md#GetImageTobase64) | **Get** /image/tobase64 | 图片转 Base64
[**PostImageCompress**](ImageAPI.md#PostImageCompress) | **Post** /image/compress | 无损压缩图片
[**PostImageFrombase64**](ImageAPI.md#PostImageFrombase64) | **Post** /image/frombase64 | 通过Base64编码上传图片
[**PostImageMotou**](ImageAPI.md#PostImageMotou) | **Post** /image/motou | 生成摸摸头GIF
[**PostImageNsfw**](ImageAPI.md#PostImageNsfw) | **Post** /image/nsfw | 图片敏感检测
[**PostImageSpeechless**](ImageAPI.md#PostImageSpeechless) | **Post** /image/speechless | 生成你们怎么不说话了表情包
[**PostImageSvg**](ImageAPI.md#PostImageSvg) | **Post** /image/svg | SVG转图片



## GetAvatarGravatar

> *os.File GetAvatarGravatar(ctx).Email(email).Hash(hash).S(s).D(d).R(r).Execute()

获取Gravatar头像



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
	email := "shuakami@sdjz.wiki" // string | 用户的 Email 地址。如果未提供 `hash` 参数，则此参数为必需。 (optional)
	hash := "hash_example" // string | 用户 Email 地址的小写 MD5 哈希值。如果提供此参数，将忽略 `email` 参数。 (optional)
	s := int32(56) // int32 | 头像的尺寸，单位为像素。有效范围是 1 到 2048。 (optional) (default to 80)
	d := "d_example" // string | 当用户没有自己的 Gravatar 头像时，显示的默认头像类型。可选值包括 `mp`, `identicon`, `monsterid`, `wavatar`, `retro`, `robohash`, `blank`, `404`。 (optional) (default to "mp")
	r := "r_example" // string | 头像分级。可选值：`g`, `pg`, `r`, `x`。 (optional) (default to "g")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetAvatarGravatar(context.Background()).Email(email).Hash(hash).S(s).D(d).R(r).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetAvatarGravatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvatarGravatar`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetAvatarGravatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarGravatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | 用户的 Email 地址。如果未提供 &#x60;hash&#x60; 参数，则此参数为必需。 | 
 **hash** | **string** | 用户 Email 地址的小写 MD5 哈希值。如果提供此参数，将忽略 &#x60;email&#x60; 参数。 | 
 **s** | **int32** | 头像的尺寸，单位为像素。有效范围是 1 到 2048。 | [default to 80]
 **d** | **string** | 当用户没有自己的 Gravatar 头像时，显示的默认头像类型。可选值包括 &#x60;mp&#x60;, &#x60;identicon&#x60;, &#x60;monsterid&#x60;, &#x60;wavatar&#x60;, &#x60;retro&#x60;, &#x60;robohash&#x60;, &#x60;blank&#x60;, &#x60;404&#x60;。 | [default to &quot;mp&quot;]
 **r** | **string** | 头像分级。可选值：&#x60;g&#x60;, &#x60;pg&#x60;, &#x60;r&#x60;, &#x60;x&#x60;。 | [default to &quot;g&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageBingDaily

> *os.File GetImageBingDaily(ctx).Execute()

必应壁纸



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
	resp, r, err := apiClient.ImageAPI.GetImageBingDaily(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageBingDaily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageBingDaily`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageBingDaily`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageBingDailyRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageMotou

> *os.File GetImageMotou(ctx).Qq(qq).BgColor(bgColor).Execute()

生成摸摸头GIF (QQ号)



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
	qq := "10001" // string | 你想要摸头的对象的QQ号码。
	bgColor := "transparent" // string | GIF的背景颜色。留空则由后端服务决定默认值。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImageMotou(context.Background()).Qq(qq).BgColor(bgColor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageMotou``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageMotou`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageMotou`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageMotouRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qq** | **string** | 你想要摸头的对象的QQ号码。 | 
 **bgColor** | **string** | GIF的背景颜色。留空则由后端服务决定默认值。 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/gif, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageQrcode

> *os.File GetImageQrcode(ctx).Text(text).Size(size).Format(format).Transparent(transparent).Fgcolor(fgcolor).Bgcolor(bgcolor).Execute()

生成二维码



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
	text := "https://www.bilibili.com/video/BV1uT4y1P7CX/" // string | 你希望编码到二维码中的任何文本内容，比如一个URL、一段话或者一个JSON字符串。
	size := int32(512) // int32 | 二维码图片的边长（正方形），单位是像素。有效范围是 256 到 2048 之间。 (optional) (default to 256)
	format := "image" // string | 指定响应内容的格式。可选值为 `image`, `json`, `json_url`。 (optional) (default to "image")
	transparent := true // bool | 是否使用透明背景。启用后生成的 PNG 图片将具有 alpha 通道，背景透明。 (optional) (default to false)
	fgcolor := "fgcolor_example" // string | 二维码前景色（即二维码本身的颜色），使用十六进制格式。URL 中需要将 `#` 编码为 `%23`。 (optional) (default to "#000000")
	bgcolor := "bgcolor_example" // string | 二维码背景色，使用十六进制格式。当 `transparent=true` 时此参数会被忽略。URL 中需要将 `#` 编码为 `%23`。 (optional) (default to "#FFFFFF")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImageQrcode(context.Background()).Text(text).Size(size).Format(format).Transparent(transparent).Fgcolor(fgcolor).Bgcolor(bgcolor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageQrcode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageQrcode`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageQrcode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageQrcodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | 你希望编码到二维码中的任何文本内容，比如一个URL、一段话或者一个JSON字符串。 | 
 **size** | **int32** | 二维码图片的边长（正方形），单位是像素。有效范围是 256 到 2048 之间。 | [default to 256]
 **format** | **string** | 指定响应内容的格式。可选值为 &#x60;image&#x60;, &#x60;json&#x60;, &#x60;json_url&#x60;。 | [default to &quot;image&quot;]
 **transparent** | **bool** | 是否使用透明背景。启用后生成的 PNG 图片将具有 alpha 通道，背景透明。 | [default to false]
 **fgcolor** | **string** | 二维码前景色（即二维码本身的颜色），使用十六进制格式。URL 中需要将 &#x60;#&#x60; 编码为 &#x60;%23&#x60;。 | [default to &quot;#000000&quot;]
 **bgcolor** | **string** | 二维码背景色，使用十六进制格式。当 &#x60;transparent&#x3D;true&#x60; 时此参数会被忽略。URL 中需要将 &#x60;#&#x60; 编码为 &#x60;%23&#x60;。 | [default to &quot;#FFFFFF&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageTobase64

> GetImageTobase64200Response GetImageTobase64(ctx).Url(url).Execute()

图片转 Base64



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
	url := "https://ts3.tc.mm.bing.net/th?id=ORMS.44196851bb1757ec3f66572811fe8e07&pid=Wdp&w=612&h=304&qlt=90&c=1&rs=1&dpr=1.25&p=0" // string | 需要转换为Base64的、可公开访问的图片URL地址。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImageTobase64(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageTobase64``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageTobase64`: GetImageTobase64200Response
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageTobase64`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageTobase64Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | 需要转换为Base64的、可公开访问的图片URL地址。 | 

### Return type

[**GetImageTobase64200Response**](GetImageTobase64200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageCompress

> *os.File PostImageCompress(ctx).File(file).Level(level).Format(format).Execute()

无损压缩图片



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
	file := os.NewFile(1234, "some_file") // *os.File | 支持PNG, JPG, JPEG等常见图片格式。文件大小不超过15MB。
	level := int32(3) // int32 | 压缩强度 (1-5)，默认为 3。数字越小，压缩率越高。 (optional) (default to 3)
	format := "png" // string | 输出图片格式，可以是 'png' 或 'jpeg'。 (optional) (default to "png")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageCompress(context.Background()).File(file).Level(level).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageCompress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageCompress`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageCompress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageCompressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | 支持PNG, JPG, JPEG等常见图片格式。文件大小不超过15MB。 | 
 **level** | **int32** | 压缩强度 (1-5)，默认为 3。数字越小，压缩率越高。 | [default to 3]
 **format** | **string** | 输出图片格式，可以是 &#39;png&#39; 或 &#39;jpeg&#39;。 | [default to &quot;png&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageFrombase64

> PostImageFrombase64200Response PostImageFrombase64(ctx).PostImageFrombase64Request(postImageFrombase64Request).Execute()

通过Base64编码上传图片



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
	postImageFrombase64Request := *openapiclient.NewPostImageFrombase64Request("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII=") // PostImageFrombase64Request | 一个JSON对象，包含 `imageData` 字段，其值为你想要上传图片的完整Base64 Data URI。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageFrombase64(context.Background()).PostImageFrombase64Request(postImageFrombase64Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageFrombase64``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageFrombase64`: PostImageFrombase64200Response
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageFrombase64`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageFrombase64Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postImageFrombase64Request** | [**PostImageFrombase64Request**](PostImageFrombase64Request.md) | 一个JSON对象，包含 &#x60;imageData&#x60; 字段，其值为你想要上传图片的完整Base64 Data URI。 | 

### Return type

[**PostImageFrombase64200Response**](PostImageFrombase64200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageMotou

> *os.File PostImageMotou(ctx).ImageUrl(imageUrl).File(file).BgColor(bgColor).Execute()

生成摸摸头GIF



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
	imageUrl := "imageUrl_example" // string | 图片的URL地址。如果提供此项，将优先使用该URL的图片。 (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 上传的图片文件。支持JPG、PNG、GIF等常见格式。 (optional)
	bgColor := "bgColor_example" // string | GIF的背景颜色。可选值为 'white', 'black', 'transparent'。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageMotou(context.Background()).ImageUrl(imageUrl).File(file).BgColor(bgColor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageMotou``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageMotou`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageMotou`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageMotouRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageUrl** | **string** | 图片的URL地址。如果提供此项，将优先使用该URL的图片。 | 
 **file** | ***os.File** | 上传的图片文件。支持JPG、PNG、GIF等常见格式。 | 
 **bgColor** | **string** | GIF的背景颜色。可选值为 &#39;white&#39;, &#39;black&#39;, &#39;transparent&#39;。 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: image/gif, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageNsfw

> PostImageNsfw200Response PostImageNsfw(ctx).File(file).Url(url).Execute()

图片敏感检测



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
	file := os.NewFile(1234, "some_file") // *os.File | 要检测的图片文件。支持 JPG、JPEG、PNG、GIF、WebP 格式，最大 20MB。 (optional)
	url := "url_example" // string | 图片的 URL 地址。如果同时提供 file 和 url，将优先使用 file。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageNsfw(context.Background()).File(file).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageNsfw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageNsfw`: PostImageNsfw200Response
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageNsfw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageNsfwRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | 要检测的图片文件。支持 JPG、JPEG、PNG、GIF、WebP 格式，最大 20MB。 | 
 **url** | **string** | 图片的 URL 地址。如果同时提供 file 和 url，将优先使用 file。 | 

### Return type

[**PostImageNsfw200Response**](PostImageNsfw200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageSpeechless

> *os.File PostImageSpeechless(ctx).PostImageSpeechlessRequest(postImageSpeechlessRequest).Execute()

生成你们怎么不说话了表情包



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
	postImageSpeechlessRequest := *openapiclient.NewPostImageSpeechlessRequest() // PostImageSpeechlessRequest | 包含表情包文字内容的JSON对象。至少需要提供上方或下方文字之一。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageSpeechless(context.Background()).PostImageSpeechlessRequest(postImageSpeechlessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageSpeechless``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageSpeechless`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageSpeechless`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageSpeechlessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postImageSpeechlessRequest** | [**PostImageSpeechlessRequest**](PostImageSpeechlessRequest.md) | 包含表情包文字内容的JSON对象。至少需要提供上方或下方文字之一。 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageSvg

> *os.File PostImageSvg(ctx).Format(format).Width(width).Height(height).Quality(quality).File(file).Execute()

SVG转图片



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
	format := "format_example" // string | 输出图像的目标格式。支持的值：`png`, `jpeg`, `jpg`, `gif`, `tiff`, `bmp`。 (optional) (default to "png")
	width := int32(56) // int32 | 输出图像的宽度（像素）。如果省略，将根据 `height` 保持宽高比，或者使用 SVG 的原始宽度。 (optional)
	height := int32(56) // int32 | 输出图像的高度（像素）。如果省略，将根据 `width` 保持宽高比，或者使用 SVG 的原始高度。 (optional)
	quality := int32(56) // int32 | JPEG 图像的压缩质量（1-100）。仅当 `format` 为 `jpeg` 或 `jpg` 时有效。 (optional) (default to 90)
	file := os.NewFile(1234, "some_file") // *os.File | 支持SVG文件 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageSvg(context.Background()).Format(format).Width(width).Height(height).Quality(quality).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageSvg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageSvg`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageSvg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageSvgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | 输出图像的目标格式。支持的值：&#x60;png&#x60;, &#x60;jpeg&#x60;, &#x60;jpg&#x60;, &#x60;gif&#x60;, &#x60;tiff&#x60;, &#x60;bmp&#x60;。 | [default to &quot;png&quot;]
 **width** | **int32** | 输出图像的宽度（像素）。如果省略，将根据 &#x60;height&#x60; 保持宽高比，或者使用 SVG 的原始宽度。 | 
 **height** | **int32** | 输出图像的高度（像素）。如果省略，将根据 &#x60;width&#x60; 保持宽高比，或者使用 SVG 的原始高度。 | 
 **quality** | **int32** | JPEG 图像的压缩质量（1-100）。仅当 &#x60;format&#x60; 为 &#x60;jpeg&#x60; 或 &#x60;jpg&#x60; 时有效。 | [default to 90]
 **file** | ***os.File** | 支持SVG文件 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


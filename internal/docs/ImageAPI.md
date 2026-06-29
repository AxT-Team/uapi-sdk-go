# \ImageAPI

All URIs are relative to *https://uapis.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvatarGravatar**](ImageAPI.md#GetAvatarGravatar) | **Get** /avatar/gravatar | 获取Gravatar头像
[**GetImageBingDaily**](ImageAPI.md#GetImageBingDaily) | **Get** /image/bing-daily | 获取必应每日壁纸
[**GetImageBingDailyHistory**](ImageAPI.md#GetImageBingDailyHistory) | **Get** /image/bing-daily/history | 查询必应壁纸历史
[**GetImageMotou**](ImageAPI.md#GetImageMotou) | **Get** /image/motou | 生成摸摸头GIF (QQ号)
[**GetImageQrcode**](ImageAPI.md#GetImageQrcode) | **Get** /image/qrcode | 生成二维码
[**GetImageTobase64**](ImageAPI.md#GetImageTobase64) | **Get** /image/tobase64 | 图片转 Base64
[**PostImageCompress**](ImageAPI.md#PostImageCompress) | **Post** /image/compress | 无损压缩图片
[**PostImageDecode**](ImageAPI.md#PostImageDecode) | **Post** /image/decode | 解码并缩放图片
[**PostImageFrombase64**](ImageAPI.md#PostImageFrombase64) | **Post** /image/frombase64 | 通过Base64编码上传图片
[**PostImageMotou**](ImageAPI.md#PostImageMotou) | **Post** /image/motou | 生成摸摸头GIF
[**PostImageNsfw**](ImageAPI.md#PostImageNsfw) | **Post** /image/nsfw | 图片敏感检测
[**PostImageOcr**](ImageAPI.md#PostImageOcr) | **Post** /image/ocr | 通用 OCR 文字识别
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
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageBingDaily

> *os.File GetImageBingDaily(ctx).Date(date).Random(random).Resolution(resolution).Format(format).Execute()

获取必应每日壁纸



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
	date := "date_example" // string | 壁纸日期，格式是 `YYYY-MM-DD`。不传时返回当天壁纸。 (optional)
	random := true // bool | 是否每次请求随机返回一张历史壁纸。传 `true` 时生效；不能和 `date` 同时使用。不传或传 `false` 时保持默认当天/指定日期逻辑。 (optional) (default to false)
	resolution := "4k" // string | 返回图片的目标分辨率。可以传 `4k` 或 `1080`，不传时默认是 `4k`。 (optional) (default to "4k")
	format := "format_example" // string | 响应格式。可以传 `image`、`json` 或 `redirect`。不传时默认是 `image`。 (optional) (default to "image")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImageBingDaily(context.Background()).Date(date).Random(random).Resolution(resolution).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageBingDaily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageBingDaily`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageBingDaily`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageBingDailyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | 壁纸日期，格式是 &#x60;YYYY-MM-DD&#x60;。不传时返回当天壁纸。 | 
 **random** | **bool** | 是否每次请求随机返回一张历史壁纸。传 &#x60;true&#x60; 时生效；不能和 &#x60;date&#x60; 同时使用。不传或传 &#x60;false&#x60; 时保持默认当天/指定日期逻辑。 | [default to false]
 **resolution** | **string** | 返回图片的目标分辨率。可以传 &#x60;4k&#x60; 或 &#x60;1080&#x60;，不传时默认是 &#x60;4k&#x60;。 | [default to &quot;4k&quot;]
 **format** | **string** | 响应格式。可以传 &#x60;image&#x60;、&#x60;json&#x60; 或 &#x60;redirect&#x60;。不传时默认是 &#x60;image&#x60;。 | [default to &quot;image&quot;]

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


## GetImageBingDailyHistory

> GetImageBingDailyHistory200Response GetImageBingDailyHistory(ctx).Date(date).Resolution(resolution).Page(page).PageSize(pageSize).Execute()

查询必应壁纸历史



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
	date := "2026-04-05" // string | 壁纸日期，格式是 `YYYY-MM-DD`。传了以后会按日期精确查询，并且忽略 `page` 和 `page_size`。 (optional)
	resolution := "1080" // string | 返回图片的目标分辨率。可以传 `4k` 或 `1080`，不传时默认是 `4k`。 (optional) (default to "4k")
	page := int32(2) // int32 | 分页页码，必须是正整数。不传时默认是 `1`。只有在不传 `date` 时才生效。 (optional) (default to 1)
	pageSize := int32(10) // int32 | 每页条数，必须是正整数。不传时默认是 `30`，最大是 `100`。只有在不传 `date` 时才生效。 (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImageBingDailyHistory(context.Background()).Date(date).Resolution(resolution).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageBingDailyHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageBingDailyHistory`: GetImageBingDailyHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageBingDailyHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageBingDailyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | 壁纸日期，格式是 &#x60;YYYY-MM-DD&#x60;。传了以后会按日期精确查询，并且忽略 &#x60;page&#x60; 和 &#x60;page_size&#x60;。 | 
 **resolution** | **string** | 返回图片的目标分辨率。可以传 &#x60;4k&#x60; 或 &#x60;1080&#x60;，不传时默认是 &#x60;4k&#x60;。 | [default to &quot;4k&quot;]
 **page** | **int32** | 分页页码，必须是正整数。不传时默认是 &#x60;1&#x60;。只有在不传 &#x60;date&#x60; 时才生效。 | [default to 1]
 **pageSize** | **int32** | 每页条数，必须是正整数。不传时默认是 &#x60;30&#x60;，最大是 &#x60;100&#x60;。只有在不传 &#x60;date&#x60; 时才生效。 | [default to 30]

### Return type

[**GetImageBingDailyHistory200Response**](GetImageBingDailyHistory200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageDecode

> *os.File PostImageDecode(ctx).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).Format(format).ColorMode(colorMode).Fit(fit).Background(background).File(file).Url(url).Execute()

解码并缩放图片



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
	width := int32(800) // int32 | 目标宽度，单位是像素。可以单独传，也可以和 `height` 一起传。与 `max_width`、`max_height` 互斥。 (optional)
	height := int32(600) // int32 | 目标高度，单位是像素。可以单独传，也可以和 `width` 一起传。与 `max_width`、`max_height` 互斥。 (optional)
	maxWidth := int32(56) // int32 | 最大宽度，单位是像素。只有在不传 `width`、`height` 时才生效，会按原比例缩放。 (optional)
	maxHeight := int32(56) // int32 | 最大高度，单位是像素。只有在不传 `width`、`height` 时才生效，会按原比例缩放。 (optional)
	format := "format_example" // string | 输出格式。可以传 `bmp`、`rgb565` 或 `rgb888`，不传时默认是 `bmp`。 (optional) (default to "bmp")
	colorMode := "colorMode_example" // string | BMP 输出的颜色模式。只有在 `format=bmp` 时才生效，可以传 `RGB565` 或 `RGB888`，不传时默认是 `RGB888`。 (optional) (default to "RGB888")
	fit := "fit_example" // string | 缩放模式。可以传 `contain`、`cover` 或 `fill`，不传时默认是 `contain`。当传 `cover` 或 `fill` 时，`width` 和 `height` 都要传。 (optional) (default to "contain")
	background := "background_example" // string | 背景色。可以传 `black`、`white` 或 `#RRGGBB`，不传时默认是 `black`。 (optional) (default to "black")
	file := os.NewFile(1234, "some_file") // *os.File | 要处理的图片文件。这个接口适合直接上传 JPG、JPEG、PNG、WebP、BMP 等常见格式。 (optional)
	url := "url_example" // string | 要处理的图片链接。适合不方便直接上传文件时使用。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageDecode(context.Background()).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).Format(format).ColorMode(colorMode).Fit(fit).Background(background).File(file).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageDecode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageDecode`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageDecode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageDecodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | **int32** | 目标宽度，单位是像素。可以单独传，也可以和 &#x60;height&#x60; 一起传。与 &#x60;max_width&#x60;、&#x60;max_height&#x60; 互斥。 | 
 **height** | **int32** | 目标高度，单位是像素。可以单独传，也可以和 &#x60;width&#x60; 一起传。与 &#x60;max_width&#x60;、&#x60;max_height&#x60; 互斥。 | 
 **maxWidth** | **int32** | 最大宽度，单位是像素。只有在不传 &#x60;width&#x60;、&#x60;height&#x60; 时才生效，会按原比例缩放。 | 
 **maxHeight** | **int32** | 最大高度，单位是像素。只有在不传 &#x60;width&#x60;、&#x60;height&#x60; 时才生效，会按原比例缩放。 | 
 **format** | **string** | 输出格式。可以传 &#x60;bmp&#x60;、&#x60;rgb565&#x60; 或 &#x60;rgb888&#x60;，不传时默认是 &#x60;bmp&#x60;。 | [default to &quot;bmp&quot;]
 **colorMode** | **string** | BMP 输出的颜色模式。只有在 &#x60;format&#x3D;bmp&#x60; 时才生效，可以传 &#x60;RGB565&#x60; 或 &#x60;RGB888&#x60;，不传时默认是 &#x60;RGB888&#x60;。 | [default to &quot;RGB888&quot;]
 **fit** | **string** | 缩放模式。可以传 &#x60;contain&#x60;、&#x60;cover&#x60; 或 &#x60;fill&#x60;，不传时默认是 &#x60;contain&#x60;。当传 &#x60;cover&#x60; 或 &#x60;fill&#x60; 时，&#x60;width&#x60; 和 &#x60;height&#x60; 都要传。 | [default to &quot;contain&quot;]
 **background** | **string** | 背景色。可以传 &#x60;black&#x60;、&#x60;white&#x60; 或 &#x60;#RRGGBB&#x60;，不传时默认是 &#x60;black&#x60;。 | [default to &quot;black&quot;]
 **file** | ***os.File** | 要处理的图片文件。这个接口适合直接上传 JPG、JPEG、PNG、WebP、BMP 等常见格式。 | 
 **url** | **string** | 要处理的图片链接。适合不方便直接上传文件时使用。 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: image/bmp, application/octet-stream, application/json

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
	postImageFrombase64Request := *openapiclient.NewPostImageFrombase64Request("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII=") // PostImageFrombase64Request | 

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
 **postImageFrombase64Request** | [**PostImageFrombase64Request**](PostImageFrombase64Request.md) |  | 

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

> *os.File PostImageMotou(ctx).BgColor(bgColor).File(file).ImageUrl(imageUrl).Execute()

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
	bgColor := "bgColor_example" // string | GIF的背景颜色。可选值为 'white', 'black', 'transparent'。 (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 上传的图片文件。支持JPG、PNG、GIF等常见格式。 (optional)
	imageUrl := "imageUrl_example" // string | 图片的URL地址。如果提供此项，将优先使用该URL的图片。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageMotou(context.Background()).BgColor(bgColor).File(file).ImageUrl(imageUrl).Execute()
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
 **bgColor** | **string** | GIF的背景颜色。可选值为 &#39;white&#39;, &#39;black&#39;, &#39;transparent&#39;。 | 
 **file** | ***os.File** | 上传的图片文件。支持JPG、PNG、GIF等常见格式。 | 
 **imageUrl** | **string** | 图片的URL地址。如果提供此项，将优先使用该URL的图片。 | 

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

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageOcr

> PostImageOcr200Response PostImageOcr(ctx).EnableCls(enableCls).File(file).ImageBase64(imageBase64).ImageName(imageName).NeedLocation(needLocation).ReturnMarkdown(returnMarkdown).Url(url).Execute()

通用 OCR 文字识别



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
	enableCls := "enableCls_example" // string | 是否开启额外的文字方向校正。请传 `true` 或 `false`，不传时默认是 `false`。 (optional) (default to "false")
	file := os.NewFile(1234, "some_file") // *os.File | 待识别的图片文件。支持 JPG、JPEG、PNG、BMP、GIF、WebP 等常见格式，最大不超过 10MB。请勿与 url 或 image_base64 同时提交。 (optional)
	imageBase64 := "imageBase64_example" // string | 图片的 Base64 字符串。可以传完整 Data URI，也可以只传纯 Base64 内容。请勿与 file 或 url 同时提交。 (optional)
	imageName := "imageName_example" // string | 自定义图片文件名。传链接或纯 Base64 时建议一起传，便于保留或推断扩展名。 (optional)
	needLocation := "needLocation_example" // string | 是否返回文字坐标信息。请传 `true` 或 `false`，不传时默认是 `true`。 (optional) (default to "true")
	returnMarkdown := "returnMarkdown_example" // string | 是否额外返回整理后的 Markdown 文本。请传 `true` 或 `false`，不传时默认是 `false`。 (optional) (default to "false")
	url := "url_example" // string | 公网可直接访问的图片地址。请勿与 file 或 image_base64 同时提交。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.PostImageOcr(context.Background()).EnableCls(enableCls).File(file).ImageBase64(imageBase64).ImageName(imageName).NeedLocation(needLocation).ReturnMarkdown(returnMarkdown).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostImageOcr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageOcr`: PostImageOcr200Response
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.PostImageOcr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageOcrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableCls** | **string** | 是否开启额外的文字方向校正。请传 &#x60;true&#x60; 或 &#x60;false&#x60;，不传时默认是 &#x60;false&#x60;。 | [default to &quot;false&quot;]
 **file** | ***os.File** | 待识别的图片文件。支持 JPG、JPEG、PNG、BMP、GIF、WebP 等常见格式，最大不超过 10MB。请勿与 url 或 image_base64 同时提交。 | 
 **imageBase64** | **string** | 图片的 Base64 字符串。可以传完整 Data URI，也可以只传纯 Base64 内容。请勿与 file 或 url 同时提交。 | 
 **imageName** | **string** | 自定义图片文件名。传链接或纯 Base64 时建议一起传，便于保留或推断扩展名。 | 
 **needLocation** | **string** | 是否返回文字坐标信息。请传 &#x60;true&#x60; 或 &#x60;false&#x60;，不传时默认是 &#x60;true&#x60;。 | [default to &quot;true&quot;]
 **returnMarkdown** | **string** | 是否额外返回整理后的 Markdown 文本。请传 &#x60;true&#x60; 或 &#x60;false&#x60;，不传时默认是 &#x60;false&#x60;。 | [default to &quot;false&quot;]
 **url** | **string** | 公网可直接访问的图片地址。请勿与 file 或 image_base64 同时提交。 | 

### Return type

[**PostImageOcr200Response**](PostImageOcr200Response.md)

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
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \MiscAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoryProgrammer**](MiscAPI.md#GetHistoryProgrammer) | **Get** /history/programmer | 获取指定日期的程序员历史事件
[**GetHistoryProgrammerToday**](MiscAPI.md#GetHistoryProgrammerToday) | **Get** /history/programmer/today | 获取今天的程序员历史事件
[**GetMiscHotboard**](MiscAPI.md#GetMiscHotboard) | **Get** /misc/hotboard | 获取多平台实时热榜
[**GetMiscPhoneinfo**](MiscAPI.md#GetMiscPhoneinfo) | **Get** /misc/phoneinfo | 查询手机号码归属地信息
[**GetMiscRandomnumber**](MiscAPI.md#GetMiscRandomnumber) | **Get** /misc/randomnumber | 生成高度可定制的随机数
[**GetMiscTimestamp**](MiscAPI.md#GetMiscTimestamp) | **Get** /misc/timestamp | 转换时间戳 (旧版，推荐使用/convert/unixtime)
[**GetMiscTrackingCarriers**](MiscAPI.md#GetMiscTrackingCarriers) | **Get** /misc/tracking/carriers | 获取支持的快递公司列表
[**GetMiscTrackingDetect**](MiscAPI.md#GetMiscTrackingDetect) | **Get** /misc/tracking/detect | 识别快递公司
[**GetMiscTrackingQuery**](MiscAPI.md#GetMiscTrackingQuery) | **Get** /misc/tracking/query | 查询快递物流信息
[**GetMiscWeather**](MiscAPI.md#GetMiscWeather) | **Get** /misc/weather | 查询实时天气信息
[**GetMiscWorldtime**](MiscAPI.md#GetMiscWorldtime) | **Get** /misc/worldtime | 查询全球任意时区的时间



## GetHistoryProgrammer

> GetHistoryProgrammer200Response GetHistoryProgrammer(ctx).Month(month).Day(day).Execute()

获取指定日期的程序员历史事件



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
	month := int32(4) // int32 | 月份，1-12之间的整数。
	day := int32(4) // int32 | 日期，1-31之间的整数。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetHistoryProgrammer(context.Background()).Month(month).Day(day).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetHistoryProgrammer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoryProgrammer`: GetHistoryProgrammer200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetHistoryProgrammer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryProgrammerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **int32** | 月份，1-12之间的整数。 | 
 **day** | **int32** | 日期，1-31之间的整数。 | 

### Return type

[**GetHistoryProgrammer200Response**](GetHistoryProgrammer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryProgrammerToday

> GetHistoryProgrammerToday200Response GetHistoryProgrammerToday(ctx).Execute()

获取今天的程序员历史事件



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
	resp, r, err := apiClient.MiscAPI.GetHistoryProgrammerToday(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetHistoryProgrammerToday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoryProgrammerToday`: GetHistoryProgrammerToday200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetHistoryProgrammerToday`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryProgrammerTodayRequest struct via the builder pattern


### Return type

[**GetHistoryProgrammerToday200Response**](GetHistoryProgrammerToday200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscHotboard

> GetMiscHotboard200Response GetMiscHotboard(ctx).Type_(type_).Execute()

获取多平台实时热榜



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
	type_ := "weibo" // string | 你想要查询的热榜平台。支持多种主流平台类型，详见下方[可选值](#可选值)表格。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscHotboard(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscHotboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscHotboard`: GetMiscHotboard200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscHotboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscHotboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | 你想要查询的热榜平台。支持多种主流平台类型，详见下方[可选值](#可选值)表格。 | 

### Return type

[**GetMiscHotboard200Response**](GetMiscHotboard200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscPhoneinfo

> GetMiscPhoneinfo200Response GetMiscPhoneinfo(ctx).Phone(phone).Execute()

查询手机号码归属地信息



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
	phone := "13800138000" // string | 需要查询的11位中国大陆手机号码。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscPhoneinfo(context.Background()).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscPhoneinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscPhoneinfo`: GetMiscPhoneinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscPhoneinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscPhoneinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | 需要查询的11位中国大陆手机号码。 | 

### Return type

[**GetMiscPhoneinfo200Response**](GetMiscPhoneinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscRandomnumber

> GetMiscRandomnumber200Response GetMiscRandomnumber(ctx).Min(min).Max(max).Count(count).AllowRepeat(allowRepeat).AllowDecimal(allowDecimal).DecimalPlaces(decimalPlaces).Execute()

生成高度可定制的随机数



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
	min := int32(10) // int32 | 生成随机数的最小值（包含）。 (optional) (default to 1)
	max := int32(50) // int32 | 生成随机数的最大值（包含）。 (optional) (default to 100)
	count := int32(5) // int32 | 需要生成的随机数的数量。 (optional) (default to 1)
	allowRepeat := true // bool | 是否允许生成的多个数字中出现重复值。 (optional) (default to false)
	allowDecimal := true // bool | 是否生成小（浮点）数。如果为 false，则只生成整数。 (optional) (default to false)
	decimalPlaces := int32(2) // int32 | 如果 `allow_decimal=true`，这里可以指定小数的位数。 (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscRandomnumber(context.Background()).Min(min).Max(max).Count(count).AllowRepeat(allowRepeat).AllowDecimal(allowDecimal).DecimalPlaces(decimalPlaces).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscRandomnumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscRandomnumber`: GetMiscRandomnumber200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscRandomnumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscRandomnumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **min** | **int32** | 生成随机数的最小值（包含）。 | [default to 1]
 **max** | **int32** | 生成随机数的最大值（包含）。 | [default to 100]
 **count** | **int32** | 需要生成的随机数的数量。 | [default to 1]
 **allowRepeat** | **bool** | 是否允许生成的多个数字中出现重复值。 | [default to false]
 **allowDecimal** | **bool** | 是否生成小（浮点）数。如果为 false，则只生成整数。 | [default to false]
 **decimalPlaces** | **int32** | 如果 &#x60;allow_decimal&#x3D;true&#x60;，这里可以指定小数的位数。 | [default to 2]

### Return type

[**GetMiscRandomnumber200Response**](GetMiscRandomnumber200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscTimestamp

> GetMiscTimestamp200Response GetMiscTimestamp(ctx).Ts(ts).Execute()

转换时间戳 (旧版，推荐使用/convert/unixtime)



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
	ts := "1672531200" // string | 需要转换的Unix时间戳，支持10位（秒）或13位（毫秒）。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscTimestamp(context.Background()).Ts(ts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscTimestamp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscTimestamp`: GetMiscTimestamp200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscTimestamp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ts** | **string** | 需要转换的Unix时间戳，支持10位（秒）或13位（毫秒）。 | 

### Return type

[**GetMiscTimestamp200Response**](GetMiscTimestamp200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscTrackingCarriers

> GetMiscTrackingCarriers200Response GetMiscTrackingCarriers(ctx).Execute()

获取支持的快递公司列表



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
	resp, r, err := apiClient.MiscAPI.GetMiscTrackingCarriers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscTrackingCarriers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscTrackingCarriers`: GetMiscTrackingCarriers200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscTrackingCarriers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscTrackingCarriersRequest struct via the builder pattern


### Return type

[**GetMiscTrackingCarriers200Response**](GetMiscTrackingCarriers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscTrackingDetect

> GetMiscTrackingDetect200Response GetMiscTrackingDetect(ctx).TrackingNumber(trackingNumber).Execute()

识别快递公司



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
	trackingNumber := "trackingNumber_example" // string | 需要识别的快递单号。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscTrackingDetect(context.Background()).TrackingNumber(trackingNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscTrackingDetect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscTrackingDetect`: GetMiscTrackingDetect200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscTrackingDetect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscTrackingDetectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackingNumber** | **string** | 需要识别的快递单号。 | 

### Return type

[**GetMiscTrackingDetect200Response**](GetMiscTrackingDetect200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscTrackingQuery

> GetMiscTrackingQuery200Response GetMiscTrackingQuery(ctx).TrackingNumber(trackingNumber).CarrierCode(carrierCode).Execute()

查询快递物流信息



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
	trackingNumber := "trackingNumber_example" // string | 快递单号，通常是一串10-20位的数字或字母数字组合。
	carrierCode := "carrierCode_example" // string | 快递公司编码（可选）。不填写时系统会自动识别，填写后可加快查询速度。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscTrackingQuery(context.Background()).TrackingNumber(trackingNumber).CarrierCode(carrierCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscTrackingQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscTrackingQuery`: GetMiscTrackingQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscTrackingQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscTrackingQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackingNumber** | **string** | 快递单号，通常是一串10-20位的数字或字母数字组合。 | 
 **carrierCode** | **string** | 快递公司编码（可选）。不填写时系统会自动识别，填写后可加快查询速度。 | 

### Return type

[**GetMiscTrackingQuery200Response**](GetMiscTrackingQuery200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscWeather

> GetMiscWeather200Response GetMiscWeather(ctx).City(city).Adcode(adcode).Execute()

查询实时天气信息



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
	city := "北京" // string | 标准的城市名称，如 '北京', '上海市', '福田区'。请使用官方的省、市、区县行政区划名称。 (optional)
	adcode := "110000" // string | 高德地图的6位数字城市编码。例如，北京市的Adcode是 '110000'。使用Adcode查询更准确、更快速。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscWeather(context.Background()).City(city).Adcode(adcode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscWeather``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscWeather`: GetMiscWeather200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscWeather`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscWeatherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **city** | **string** | 标准的城市名称，如 &#39;北京&#39;, &#39;上海市&#39;, &#39;福田区&#39;。请使用官方的省、市、区县行政区划名称。 | 
 **adcode** | **string** | 高德地图的6位数字城市编码。例如，北京市的Adcode是 &#39;110000&#39;。使用Adcode查询更准确、更快速。 | 

### Return type

[**GetMiscWeather200Response**](GetMiscWeather200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscWorldtime

> GetMiscWorldtime200Response GetMiscWorldtime(ctx).City(city).Execute()

查询全球任意时区的时间



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
	city := "Asia/Shanghai" // string | 你需要查询的城市或地区，请使用标准的 IANA 时区数据库名称，例如 'Shanghai', 'Asia/Tokyo', 'America/New_York'。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscWorldtime(context.Background()).City(city).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscWorldtime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscWorldtime`: GetMiscWorldtime200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscWorldtime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscWorldtimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **city** | **string** | 你需要查询的城市或地区，请使用标准的 IANA 时区数据库名称，例如 &#39;Shanghai&#39;, &#39;Asia/Tokyo&#39;, &#39;America/New_York&#39;。 | 

### Return type

[**GetMiscWorldtime200Response**](GetMiscWorldtime200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


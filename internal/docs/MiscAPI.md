# \MiscAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoryProgrammer**](MiscAPI.md#GetHistoryProgrammer) | **Get** /history/programmer | 程序员历史事件
[**GetHistoryProgrammerToday**](MiscAPI.md#GetHistoryProgrammerToday) | **Get** /history/programmer/today | 程序员历史上的今天
[**GetMiscDistrict**](MiscAPI.md#GetMiscDistrict) | **Get** /misc/district | Adcode 国内外行政区域查询
[**GetMiscHolidayCalendar**](MiscAPI.md#GetMiscHolidayCalendar) | **Get** /misc/holiday-calendar | 查询节假日与万年历
[**GetMiscHotboard**](MiscAPI.md#GetMiscHotboard) | **Get** /misc/hotboard | 查询热榜
[**GetMiscLunartime**](MiscAPI.md#GetMiscLunartime) | **Get** /misc/lunartime | 查询农历时间
[**GetMiscPhoneinfo**](MiscAPI.md#GetMiscPhoneinfo) | **Get** /misc/phoneinfo | 查询手机归属地
[**GetMiscRandomnumber**](MiscAPI.md#GetMiscRandomnumber) | **Get** /misc/randomnumber | 随机数生成
[**GetMiscTimestamp**](MiscAPI.md#GetMiscTimestamp) | **Get** /misc/timestamp | 转换时间戳 (旧版，推荐使用/convert/unixtime)
[**GetMiscTrackingCarriers**](MiscAPI.md#GetMiscTrackingCarriers) | **Get** /misc/tracking/carriers | 获取支持的快递公司列表
[**GetMiscTrackingDetect**](MiscAPI.md#GetMiscTrackingDetect) | **Get** /misc/tracking/detect | 识别快递公司
[**GetMiscTrackingQuery**](MiscAPI.md#GetMiscTrackingQuery) | **Get** /misc/tracking/query | 查询快递物流信息
[**GetMiscWeather**](MiscAPI.md#GetMiscWeather) | **Get** /misc/weather | 查询天气
[**GetMiscWorldtime**](MiscAPI.md#GetMiscWorldtime) | **Get** /misc/worldtime | 查询世界时间
[**PostMiscDateDiff**](MiscAPI.md#PostMiscDateDiff) | **Post** /misc/date-diff | 计算两个日期之间的时间差值



## GetHistoryProgrammer

> GetHistoryProgrammer200Response GetHistoryProgrammer(ctx).Month(month).Day(day).Execute()

程序员历史事件



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

程序员历史上的今天



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


## GetMiscDistrict

> GetMiscDistrict200Response GetMiscDistrict(ctx).Keywords(keywords).Adcode(adcode).Lat(lat).Lng(lng).Level(level).Country(country).Limit(limit).Execute()

Adcode 国内外行政区域查询



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
	keywords := "上海" // string | 关键词搜索（城市名、区县名，支持中英文）。 (optional)
	adcode := "110000" // string | 中国行政区划代码精确查询（如 `110000`），同时返回下级行政区。 (optional)
	lat := float32(39.9) // float32 | 纬度，与 `lng` 配合使用，坐标反查附近地点。 (optional)
	lng := float32(116.4) // float32 | 经度，与 `lat` 配合使用。 (optional)
	level := "level_example" // string | 过滤行政级别。 (optional)
	country := "CN" // string | 过滤国家代码（ISO 3166-1 alpha-2），如 `CN`、`JP`、`US`、`GB`。 (optional)
	limit := int32(20) // int32 | 返回数量上限，默认 `20`，最大 `100`。 (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscDistrict(context.Background()).Keywords(keywords).Adcode(adcode).Lat(lat).Lng(lng).Level(level).Country(country).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscDistrict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscDistrict`: GetMiscDistrict200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscDistrict`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscDistrictRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keywords** | **string** | 关键词搜索（城市名、区县名，支持中英文）。 | 
 **adcode** | **string** | 中国行政区划代码精确查询（如 &#x60;110000&#x60;），同时返回下级行政区。 | 
 **lat** | **float32** | 纬度，与 &#x60;lng&#x60; 配合使用，坐标反查附近地点。 | 
 **lng** | **float32** | 经度，与 &#x60;lat&#x60; 配合使用。 | 
 **level** | **string** | 过滤行政级别。 | 
 **country** | **string** | 过滤国家代码（ISO 3166-1 alpha-2），如 &#x60;CN&#x60;、&#x60;JP&#x60;、&#x60;US&#x60;、&#x60;GB&#x60;。 | 
 **limit** | **int32** | 返回数量上限，默认 &#x60;20&#x60;，最大 &#x60;100&#x60;。 | [default to 20]

### Return type

[**GetMiscDistrict200Response**](GetMiscDistrict200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscHolidayCalendar

> GetMiscHolidayCalendar200Response GetMiscHolidayCalendar(ctx).Date(date).Month(month).Year(year).Timezone(timezone).HolidayType(holidayType).IncludeNearby(includeNearby).NearbyLimit(nearbyLimit).Execute()

查询节假日与万年历



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
	date := "2025-10-01" // string | 按天查询时填写这个参数，例如查某一天。格式：`YYYY-MM-DD`。和 `month`、`year` 三选一。 (optional)
	month := "month_example" // string | 按月查询时填写这个参数，例如查某个月。格式：`YYYY-MM`。和 `date`、`year` 三选一。 (optional)
	year := "year_example" // string | 按年查询时填写这个参数，例如查某一年。格式：`YYYY`。和 `date`、`month` 三选一。 (optional)
	timezone := "Asia/Shanghai" // string | 时区名称，默认 Asia/Shanghai。 (optional) (default to "Asia/Shanghai")
	holidayType := "legal" // string | 节日筛选类型，默认 all。 (optional) (default to "all")
	includeNearby := true // bool | 是否返回前后最近节日，仅 date 模式生效，默认 false。month/year 模式会忽略此参数。 (optional) (default to false)
	nearbyLimit := int32(7) // int32 | 返回最近节日数量限制，默认 7，最大 30。仅 date 模式 + include_nearby=true 生效。 (optional) (default to 7)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscHolidayCalendar(context.Background()).Date(date).Month(month).Year(year).Timezone(timezone).HolidayType(holidayType).IncludeNearby(includeNearby).NearbyLimit(nearbyLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscHolidayCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscHolidayCalendar`: GetMiscHolidayCalendar200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscHolidayCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscHolidayCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | 按天查询时填写这个参数，例如查某一天。格式：&#x60;YYYY-MM-DD&#x60;。和 &#x60;month&#x60;、&#x60;year&#x60; 三选一。 | 
 **month** | **string** | 按月查询时填写这个参数，例如查某个月。格式：&#x60;YYYY-MM&#x60;。和 &#x60;date&#x60;、&#x60;year&#x60; 三选一。 | 
 **year** | **string** | 按年查询时填写这个参数，例如查某一年。格式：&#x60;YYYY&#x60;。和 &#x60;date&#x60;、&#x60;month&#x60; 三选一。 | 
 **timezone** | **string** | 时区名称，默认 Asia/Shanghai。 | [default to &quot;Asia/Shanghai&quot;]
 **holidayType** | **string** | 节日筛选类型，默认 all。 | [default to &quot;all&quot;]
 **includeNearby** | **bool** | 是否返回前后最近节日，仅 date 模式生效，默认 false。month/year 模式会忽略此参数。 | [default to false]
 **nearbyLimit** | **int32** | 返回最近节日数量限制，默认 7，最大 30。仅 date 模式 + include_nearby&#x3D;true 生效。 | [default to 7]

### Return type

[**GetMiscHolidayCalendar200Response**](GetMiscHolidayCalendar200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiscHotboard

> GetMiscHotboard200Response GetMiscHotboard(ctx).Type_(type_).Time(time).Keyword(keyword).TimeStart(timeStart).TimeEnd(timeEnd).Limit(limit).Sources(sources).Execute()

查询热榜



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
	time := int64(1700000000000) // int64 | 时光机模式：毫秒时间戳，返回最接近该时间的热榜快照。不传则返回当前实时热榜。 (optional)
	keyword := "AI" // string | 搜索模式：搜索关键词，在历史热榜中搜索包含该关键词的条目。需配合 time_start 和 time_end 使用。 (optional)
	timeStart := int64(1699900000000) // int64 | 搜索模式必填：搜索起始时间戳（毫秒）。 (optional)
	timeEnd := int64(1700100000000) // int64 | 搜索模式必填：搜索结束时间戳（毫秒）。 (optional)
	limit := int32(50) // int32 | 搜索模式下最大返回条数，默认 50，最大 200。 (optional) (default to 50)
	sources := true // bool | 设为 true 时列出所有可用的历史数据源，忽略其他参数。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscHotboard(context.Background()).Type_(type_).Time(time).Keyword(keyword).TimeStart(timeStart).TimeEnd(timeEnd).Limit(limit).Sources(sources).Execute()
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
 **time** | **int64** | 时光机模式：毫秒时间戳，返回最接近该时间的热榜快照。不传则返回当前实时热榜。 | 
 **keyword** | **string** | 搜索模式：搜索关键词，在历史热榜中搜索包含该关键词的条目。需配合 time_start 和 time_end 使用。 | 
 **timeStart** | **int64** | 搜索模式必填：搜索起始时间戳（毫秒）。 | 
 **timeEnd** | **int64** | 搜索模式必填：搜索结束时间戳（毫秒）。 | 
 **limit** | **int32** | 搜索模式下最大返回条数，默认 50，最大 200。 | [default to 50]
 **sources** | **bool** | 设为 true 时列出所有可用的历史数据源，忽略其他参数。 | 

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


## GetMiscLunartime

> GetMiscLunartime200Response GetMiscLunartime(ctx).Ts(ts).Timezone(timezone).Execute()

查询农历时间



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
	ts := "1707537600" // string | Unix 时间戳，支持 10 位秒级或 13 位毫秒级。不传则默认当前时间。 (optional)
	timezone := "Asia/Shanghai" // string | 时区名称。支持 IANA 时区（如 Asia/Shanghai）和别名（Shanghai、Beijing）。默认 Asia/Shanghai。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscLunartime(context.Background()).Ts(ts).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.GetMiscLunartime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiscLunartime`: GetMiscLunartime200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.GetMiscLunartime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiscLunartimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ts** | **string** | Unix 时间戳，支持 10 位秒级或 13 位毫秒级。不传则默认当前时间。 | 
 **timezone** | **string** | 时区名称。支持 IANA 时区（如 Asia/Shanghai）和别名（Shanghai、Beijing）。默认 Asia/Shanghai。 | 

### Return type

[**GetMiscLunartime200Response**](GetMiscLunartime200Response.md)

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

查询手机归属地



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

随机数生成



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

> GetMiscTrackingQuery200Response GetMiscTrackingQuery(ctx).TrackingNumber(trackingNumber).CarrierCode(carrierCode).Phone(phone).Execute()

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
	phone := "phone_example" // string | 收件人手机尾号，4位数字（可选）。部分快递公司需要验证手机尾号才能查询详细物流信息。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscTrackingQuery(context.Background()).TrackingNumber(trackingNumber).CarrierCode(carrierCode).Phone(phone).Execute()
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
 **phone** | **string** | 收件人手机尾号，4位数字（可选）。部分快递公司需要验证手机尾号才能查询详细物流信息。 | 

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

> GetMiscWeather200Response GetMiscWeather(ctx).City(city).Adcode(adcode).Extended(extended).Forecast(forecast).Hourly(hourly).Minutely(minutely).Indices(indices).Lang(lang).Execute()

查询天气



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
	city := "北京" // string | 城市名称，支持中文（`北京`）和英文（`Tokyo`）。可选参数，不传时会尝试 IP 自动定位。 (optional)
	adcode := "adcode_example" // string | 城市行政区划代码（如 `110000`），优先级高于 city。可选参数，不传时会尝试 IP 自动定位。 (optional)
	extended := true // bool | 返回扩展气象字段（体感温度、能见度、气压、紫外线、降水量、云量、空气质量指数及污染物分项数据）。 (optional)
	forecast := true // bool | 返回多天预报数据（最多7天），含白天夜间天气、风向风力、日出日落等。 (optional)
	hourly := true // bool | 返回逐小时预报（24小时），含温度、天气、风向风速、湿度、降水概率等。 (optional)
	minutely := true // bool | 返回分钟级降水预报（仅国内城市），每5分钟一个数据点，共24个。 (optional)
	indices := true // bool | 返回18项生活指数（穿衣、紫外线、洗车、晾晒、空调、感冒、运动、舒适度、出行、钓鱼、过敏、防晒、心情、啤酒、雨伞、交通、空气净化器、花粉）。 (optional)
	lang := "lang_example" // string | 返回语言。`zh` 返回中文（默认），`en` 返回英文。城市名翻译覆盖 7000+ 城市。生活指数（`indices`）目前仅支持中文。 (optional) (default to "zh")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.GetMiscWeather(context.Background()).City(city).Adcode(adcode).Extended(extended).Forecast(forecast).Hourly(hourly).Minutely(minutely).Indices(indices).Lang(lang).Execute()
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
 **city** | **string** | 城市名称，支持中文（&#x60;北京&#x60;）和英文（&#x60;Tokyo&#x60;）。可选参数，不传时会尝试 IP 自动定位。 | 
 **adcode** | **string** | 城市行政区划代码（如 &#x60;110000&#x60;），优先级高于 city。可选参数，不传时会尝试 IP 自动定位。 | 
 **extended** | **bool** | 返回扩展气象字段（体感温度、能见度、气压、紫外线、降水量、云量、空气质量指数及污染物分项数据）。 | 
 **forecast** | **bool** | 返回多天预报数据（最多7天），含白天夜间天气、风向风力、日出日落等。 | 
 **hourly** | **bool** | 返回逐小时预报（24小时），含温度、天气、风向风速、湿度、降水概率等。 | 
 **minutely** | **bool** | 返回分钟级降水预报（仅国内城市），每5分钟一个数据点，共24个。 | 
 **indices** | **bool** | 返回18项生活指数（穿衣、紫外线、洗车、晾晒、空调、感冒、运动、舒适度、出行、钓鱼、过敏、防晒、心情、啤酒、雨伞、交通、空气净化器、花粉）。 | 
 **lang** | **string** | 返回语言。&#x60;zh&#x60; 返回中文（默认），&#x60;en&#x60; 返回英文。城市名翻译覆盖 7000+ 城市。生活指数（&#x60;indices&#x60;）目前仅支持中文。 | [default to &quot;zh&quot;]

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

查询世界时间



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


## PostMiscDateDiff

> PostMiscDateDiff200Response PostMiscDateDiff(ctx).PostMiscDateDiffRequest(postMiscDateDiffRequest).Execute()

计算两个日期之间的时间差值



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
	postMiscDateDiffRequest := *openapiclient.NewPostMiscDateDiffRequest("2025-01-01", "2025-12-31") // PostMiscDateDiffRequest | 包含日期信息的JSON对象

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscAPI.PostMiscDateDiff(context.Background()).PostMiscDateDiffRequest(postMiscDateDiffRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.PostMiscDateDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMiscDateDiff`: PostMiscDateDiff200Response
	fmt.Fprintf(os.Stdout, "Response from `MiscAPI.PostMiscDateDiff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMiscDateDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postMiscDateDiffRequest** | [**PostMiscDateDiffRequest**](PostMiscDateDiffRequest.md) | 包含日期信息的JSON对象 | 

### Return type

[**PostMiscDateDiff200Response**](PostMiscDateDiff200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


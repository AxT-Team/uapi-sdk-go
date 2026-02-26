# \GameAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGameEpicFree**](GameAPI.md#GetGameEpicFree) | **Get** /game/epic-free | Epic 免费游戏
[**GetGameMinecraftHistoryid**](GameAPI.md#GetGameMinecraftHistoryid) | **Get** /game/minecraft/historyid | 查询 MC 曾用名
[**GetGameMinecraftServerstatus**](GameAPI.md#GetGameMinecraftServerstatus) | **Get** /game/minecraft/serverstatus | 查询 MC 服务器
[**GetGameMinecraftUserinfo**](GameAPI.md#GetGameMinecraftUserinfo) | **Get** /game/minecraft/userinfo | 查询 MC 玩家
[**GetGameSteamSummary**](GameAPI.md#GetGameSteamSummary) | **Get** /game/steam/summary | 查询 Steam 用户



## GetGameEpicFree

> GetGameEpicFree200Response GetGameEpicFree(ctx).Execute()

Epic 免费游戏



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
	resp, r, err := apiClient.GameAPI.GetGameEpicFree(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetGameEpicFree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameEpicFree`: GetGameEpicFree200Response
	fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetGameEpicFree`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameEpicFreeRequest struct via the builder pattern


### Return type

[**GetGameEpicFree200Response**](GetGameEpicFree200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGameMinecraftHistoryid

> GetGameMinecraftHistoryid200Response GetGameMinecraftHistoryid(ctx).Name(name).Uuid(uuid).Execute()

查询 MC 曾用名



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
	name := "ExamplePlayer" // string | 玩家的 Minecraft 用户名。使用此参数查询时，会返回所有匹配用户的列表（包括当前用户名或曾用名匹配的玩家）。 (optional)
	uuid := "ee9b4ed1-aac1-491e-b761-1471be374b80" // string | 玩家的 Minecraft UUID，支持带连字符或不带连字符格式。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameAPI.GetGameMinecraftHistoryid(context.Background()).Name(name).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetGameMinecraftHistoryid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameMinecraftHistoryid`: GetGameMinecraftHistoryid200Response
	fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetGameMinecraftHistoryid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGameMinecraftHistoryidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 玩家的 Minecraft 用户名。使用此参数查询时，会返回所有匹配用户的列表（包括当前用户名或曾用名匹配的玩家）。 | 
 **uuid** | **string** | 玩家的 Minecraft UUID，支持带连字符或不带连字符格式。 | 

### Return type

[**GetGameMinecraftHistoryid200Response**](GetGameMinecraftHistoryid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGameMinecraftServerstatus

> GetGameMinecraftServerstatus200Response GetGameMinecraftServerstatus(ctx).Server(server).Execute()

查询 MC 服务器



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
	server := "hypixel.net" // string | Minecraft服务器的地址，可以是域名（如 `hypixel.net`）或 `IP:端口` 的形式（如 `mc.example.com:25565`）。如果省略端口，将默认使用 `25565`。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameAPI.GetGameMinecraftServerstatus(context.Background()).Server(server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetGameMinecraftServerstatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameMinecraftServerstatus`: GetGameMinecraftServerstatus200Response
	fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetGameMinecraftServerstatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGameMinecraftServerstatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server** | **string** | Minecraft服务器的地址，可以是域名（如 &#x60;hypixel.net&#x60;）或 &#x60;IP:端口&#x60; 的形式（如 &#x60;mc.example.com:25565&#x60;）。如果省略端口，将默认使用 &#x60;25565&#x60;。 | 

### Return type

[**GetGameMinecraftServerstatus200Response**](GetGameMinecraftServerstatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGameMinecraftUserinfo

> GetGameMinecraftUserinfo200Response GetGameMinecraftUserinfo(ctx).Username(username).Execute()

查询 MC 玩家



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
	username := "Notch" // string | 玩家的 Minecraft 游戏内名称（正版ID）。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameAPI.GetGameMinecraftUserinfo(context.Background()).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetGameMinecraftUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameMinecraftUserinfo`: GetGameMinecraftUserinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetGameMinecraftUserinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGameMinecraftUserinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | 玩家的 Minecraft 游戏内名称（正版ID）。 | 

### Return type

[**GetGameMinecraftUserinfo200Response**](GetGameMinecraftUserinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGameSteamSummary

> GetGameSteamSummary200Response GetGameSteamSummary(ctx).Steamid(steamid).Id(id).Id3(id3).Key(key).Execute()

查询 Steam 用户



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
	steamid := "76561197960287930" // string | 用户的 Steam 标识。可以是以下任意一种格式： - 纯数字的 **SteamID64** - 用户的 **自定义 URL 名称** (Vanity URL) - 完整的 **个人资料链接** (包含 SteamID64 或自定义名称) - 好友代码 (如 `22202`) (optional)
	id := "gabelogannewell" // string | 用户的 Steam 自定义URL名称（Vanity URL）。例如个人资料链接中 `/id/` 后面的部分。 (optional)
	id3 := "STEAM_0:0:22202" // string | 用户的 Steam ID3 格式标识符。传统的 Steam ID 格式，形如 STEAM_X:Y:Z。 (optional)
	key := "key_example" // string | 你的 Steam Web API Key。这是一个可选参数，如果提供，它将覆盖我们在后端配置的全局Key。这为你提供了更大的灵活性，但请务必注意Key的保密，不要在前端暴露。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameAPI.GetGameSteamSummary(context.Background()).Steamid(steamid).Id(id).Id3(id3).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetGameSteamSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameSteamSummary`: GetGameSteamSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetGameSteamSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGameSteamSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **steamid** | **string** | 用户的 Steam 标识。可以是以下任意一种格式： - 纯数字的 **SteamID64** - 用户的 **自定义 URL 名称** (Vanity URL) - 完整的 **个人资料链接** (包含 SteamID64 或自定义名称) - 好友代码 (如 &#x60;22202&#x60;) | 
 **id** | **string** | 用户的 Steam 自定义URL名称（Vanity URL）。例如个人资料链接中 &#x60;/id/&#x60; 后面的部分。 | 
 **id3** | **string** | 用户的 Steam ID3 格式标识符。传统的 Steam ID 格式，形如 STEAM_X:Y:Z。 | 
 **key** | **string** | 你的 Steam Web API Key。这是一个可选参数，如果提供，它将覆盖我们在后端配置的全局Key。这为你提供了更大的灵活性，但请务必注意Key的保密，不要在前端暴露。 | 

### Return type

[**GetGameSteamSummary200Response**](GetGameSteamSummary200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


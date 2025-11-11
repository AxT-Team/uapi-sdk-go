# \SocialAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGithubRepo**](SocialAPI.md#GetGithubRepo) | **Get** /github/repo | 获取GitHub仓库信息
[**GetSocialBilibiliArchives**](SocialAPI.md#GetSocialBilibiliArchives) | **Get** /social/bilibili/archives | 获取Bilibili用户投稿列表
[**GetSocialBilibiliLiveroom**](SocialAPI.md#GetSocialBilibiliLiveroom) | **Get** /social/bilibili/liveroom | 获取Bilibili直播间信息
[**GetSocialBilibiliReplies**](SocialAPI.md#GetSocialBilibiliReplies) | **Get** /social/bilibili/replies | 获取Bilibili视频评论
[**GetSocialBilibiliUserinfo**](SocialAPI.md#GetSocialBilibiliUserinfo) | **Get** /social/bilibili/userinfo | 查询Bilibili用户信息
[**GetSocialBilibiliVideoinfo**](SocialAPI.md#GetSocialBilibiliVideoinfo) | **Get** /social/bilibili/videoinfo | 获取Bilibili视频详细信息
[**GetSocialQqGroupinfo**](SocialAPI.md#GetSocialQqGroupinfo) | **Get** /social/qq/groupinfo | 获取QQ群名称、头像、简介
[**GetSocialQqUserinfo**](SocialAPI.md#GetSocialQqUserinfo) | **Get** /social/qq/userinfo | 独家获取QQ号头像、昵称



## GetGithubRepo

> GetGithubRepo200Response GetGithubRepo(ctx).Repo(repo).Execute()

获取GitHub仓库信息



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
	repo := "torvalds/linux" // string | 目标仓库的标识，格式为 `owner/repo`。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetGithubRepo(context.Background()).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetGithubRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGithubRepo`: GetGithubRepo200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetGithubRepo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | 目标仓库的标识，格式为 &#x60;owner/repo&#x60;。 | 

### Return type

[**GetGithubRepo200Response**](GetGithubRepo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialBilibiliArchives

> GetSocialBilibiliArchives200Response GetSocialBilibiliArchives(ctx).Mid(mid).Keywords(keywords).Orderby(orderby).Ps(ps).Pn(pn).Execute()

获取Bilibili用户投稿列表



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
	mid := "483307278" // string | B站用户的mid（用户ID）
	keywords := "keywords_example" // string | 搜索关键词，可为空 (optional)
	orderby := "pubdate" // string | 排序方式。`pubdate`=最新发布，`views`=最多播放 (optional)
	ps := "20" // string | 每页条数，默认20 (optional)
	pn := "1" // string | 页码，默认1 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialBilibiliArchives(context.Background()).Mid(mid).Keywords(keywords).Orderby(orderby).Ps(ps).Pn(pn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialBilibiliArchives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialBilibiliArchives`: GetSocialBilibiliArchives200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialBilibiliArchives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialBilibiliArchivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mid** | **string** | B站用户的mid（用户ID） | 
 **keywords** | **string** | 搜索关键词，可为空 | 
 **orderby** | **string** | 排序方式。&#x60;pubdate&#x60;&#x3D;最新发布，&#x60;views&#x60;&#x3D;最多播放 | 
 **ps** | **string** | 每页条数，默认20 | 
 **pn** | **string** | 页码，默认1 | 

### Return type

[**GetSocialBilibiliArchives200Response**](GetSocialBilibiliArchives200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialBilibiliLiveroom

> GetSocialBilibiliLiveroom200Response GetSocialBilibiliLiveroom(ctx).Mid(mid).RoomId(roomId).Execute()

获取Bilibili直播间信息



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
	mid := "672328094" // string | 主播的用户ID (`mid`)。与 `room_id` 任选其一。 (optional)
	roomId := "22637261" // string | 直播间ID，可以是长号（真实ID）或短号（靓号）。与 `mid` 任选其一。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialBilibiliLiveroom(context.Background()).Mid(mid).RoomId(roomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialBilibiliLiveroom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialBilibiliLiveroom`: GetSocialBilibiliLiveroom200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialBilibiliLiveroom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialBilibiliLiveroomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mid** | **string** | 主播的用户ID (&#x60;mid&#x60;)。与 &#x60;room_id&#x60; 任选其一。 | 
 **roomId** | **string** | 直播间ID，可以是长号（真实ID）或短号（靓号）。与 &#x60;mid&#x60; 任选其一。 | 

### Return type

[**GetSocialBilibiliLiveroom200Response**](GetSocialBilibiliLiveroom200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialBilibiliReplies

> GetSocialBilibiliReplies200Response GetSocialBilibiliReplies(ctx).Oid(oid).Sort(sort).Ps(ps).Pn(pn).Execute()

获取Bilibili视频评论



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
	oid := "1706416465" // string | 目标评论区的ID。对于视频，这通常就是它的 `aid`。
	sort := "1" // string | 排序方式。`0`=按时间, `1`=按点赞, `2`=按回复。默认为 `0`。 (optional)
	ps := "5" // string | 每页获取的评论数量，范围是1到20。默认为 `20`。 (optional)
	pn := "1" // string | 要获取的页码，从1开始。默认为 `1`。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialBilibiliReplies(context.Background()).Oid(oid).Sort(sort).Ps(ps).Pn(pn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialBilibiliReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialBilibiliReplies`: GetSocialBilibiliReplies200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialBilibiliReplies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialBilibiliRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oid** | **string** | 目标评论区的ID。对于视频，这通常就是它的 &#x60;aid&#x60;。 | 
 **sort** | **string** | 排序方式。&#x60;0&#x60;&#x3D;按时间, &#x60;1&#x60;&#x3D;按点赞, &#x60;2&#x60;&#x3D;按回复。默认为 &#x60;0&#x60;。 | 
 **ps** | **string** | 每页获取的评论数量，范围是1到20。默认为 &#x60;20&#x60;。 | 
 **pn** | **string** | 要获取的页码，从1开始。默认为 &#x60;1&#x60;。 | 

### Return type

[**GetSocialBilibiliReplies200Response**](GetSocialBilibiliReplies200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialBilibiliUserinfo

> GetSocialBilibiliUserinfo200Response GetSocialBilibiliUserinfo(ctx).Uid(uid).Execute()

查询Bilibili用户信息



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
	uid := "483307278" // string | Bilibili用户的UID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialBilibiliUserinfo(context.Background()).Uid(uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialBilibiliUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialBilibiliUserinfo`: GetSocialBilibiliUserinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialBilibiliUserinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialBilibiliUserinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** | Bilibili用户的UID | 

### Return type

[**GetSocialBilibiliUserinfo200Response**](GetSocialBilibiliUserinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialBilibiliVideoinfo

> GetSocialBilibiliVideoinfo200Response GetSocialBilibiliVideoinfo(ctx).Aid(aid).Bvid(bvid).Execute()

获取Bilibili视频详细信息



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
	aid := "75836761" // string | 视频的AV号 (aid)，纯数字格式。`aid`和`bvid`任选其一即可。 (optional)
	bvid := "BV17x411w79F" // string | 视频的BV号 (bvid)，例如 `BV117411r7R1`。`aid`和`bvid`任选其一即可。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialBilibiliVideoinfo(context.Background()).Aid(aid).Bvid(bvid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialBilibiliVideoinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialBilibiliVideoinfo`: GetSocialBilibiliVideoinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialBilibiliVideoinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialBilibiliVideoinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | 视频的AV号 (aid)，纯数字格式。&#x60;aid&#x60;和&#x60;bvid&#x60;任选其一即可。 | 
 **bvid** | **string** | 视频的BV号 (bvid)，例如 &#x60;BV117411r7R1&#x60;。&#x60;aid&#x60;和&#x60;bvid&#x60;任选其一即可。 | 

### Return type

[**GetSocialBilibiliVideoinfo200Response**](GetSocialBilibiliVideoinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialQqGroupinfo

> GetSocialQqGroupinfo200Response GetSocialQqGroupinfo(ctx).GroupId(groupId).Execute()

获取QQ群名称、头像、简介



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
	groupId := "526357265" // string | QQ群号，长度5-12位

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialQqGroupinfo(context.Background()).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialQqGroupinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialQqGroupinfo`: GetSocialQqGroupinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialQqGroupinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialQqGroupinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | QQ群号，长度5-12位 | 

### Return type

[**GetSocialQqGroupinfo200Response**](GetSocialQqGroupinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialQqUserinfo

> GetSocialQqUserinfo200Response GetSocialQqUserinfo(ctx).Qq(qq).Execute()

独家获取QQ号头像、昵称



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
	qq := "10001" // string | 需要查询的QQ号

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialQqUserinfo(context.Background()).Qq(qq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialQqUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialQqUserinfo`: GetSocialQqUserinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialQqUserinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialQqUserinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qq** | **string** | 需要查询的QQ号 | 

### Return type

[**GetSocialQqUserinfo200Response**](GetSocialQqUserinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


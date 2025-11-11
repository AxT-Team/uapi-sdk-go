# \NetworkAPI

All URIs are relative to *https://uapis.cn/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkDns**](NetworkAPI.md#GetNetworkDns) | **Get** /network/dns | 执行DNS解析查询
[**GetNetworkIcp**](NetworkAPI.md#GetNetworkIcp) | **Get** /network/icp | 查询域名ICP备案信息
[**GetNetworkIpinfo**](NetworkAPI.md#GetNetworkIpinfo) | **Get** /network/ipinfo | 查询指定IP或域名的归属信息
[**GetNetworkMyip**](NetworkAPI.md#GetNetworkMyip) | **Get** /network/myip | 获取你的公网IP及归属信息
[**GetNetworkPing**](NetworkAPI.md#GetNetworkPing) | **Get** /network/ping | 从服务器Ping指定主机
[**GetNetworkPingmyip**](NetworkAPI.md#GetNetworkPingmyip) | **Get** /network/pingmyip | 从服务器Ping你的客户端IP
[**GetNetworkPortscan**](NetworkAPI.md#GetNetworkPortscan) | **Get** /network/portscan | 扫描远程主机的指定端口
[**GetNetworkUrlstatus**](NetworkAPI.md#GetNetworkUrlstatus) | **Get** /network/urlstatus | 检查URL的可访问性状态
[**GetNetworkWhois**](NetworkAPI.md#GetNetworkWhois) | **Get** /network/whois | 查询域名的WHOIS注册信息
[**GetNetworkWxdomain**](NetworkAPI.md#GetNetworkWxdomain) | **Get** /network/wxdomain | 检查域名在微信中的访问状态



## GetNetworkDns

> GetNetworkDns200Response GetNetworkDns(ctx).Domain(domain).Type_(type_).Execute()

执行DNS解析查询



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
	domain := "cn.bing.com" // string | 你需要查询的域名，例如 'cn.bing.com'。
	type_ := "A" // string | 你想要查询的DNS记录类型。 (optional) (default to "A")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkDns(context.Background()).Domain(domain).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDns`: GetNetworkDns200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkDns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | 你需要查询的域名，例如 &#39;cn.bing.com&#39;。 | 
 **type_** | **string** | 你想要查询的DNS记录类型。 | [default to &quot;A&quot;]

### Return type

[**GetNetworkDns200Response**](GetNetworkDns200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkIcp

> GetNetworkIcp200Response GetNetworkIcp(ctx).Domain(domain).Execute()

查询域名ICP备案信息



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
	domain := "baidu.com" // string | 需要查询的域名或URL

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkIcp(context.Background()).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkIcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkIcp`: GetNetworkIcp200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkIcp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkIcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | 需要查询的域名或URL | 

### Return type

[**GetNetworkIcp200Response**](GetNetworkIcp200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkIpinfo

> GetNetworkIpinfo200Response GetNetworkIpinfo(ctx).Ip(ip).Source(source).Execute()

查询指定IP或域名的归属信息



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
	ip := "cn.bing.com" // string | 你需要查询的公网IP地址或域名（支持IPv4和IPv6）。
	source := "source_example" // string | 查询的数据源。如果留空，将使用默认的数据库。如果设置为 `commercial`，将调用商业级API，返回更详细的地理位置信息，但响应时间可能会稍长。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkIpinfo(context.Background()).Ip(ip).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkIpinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkIpinfo`: GetNetworkIpinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkIpinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkIpinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | 你需要查询的公网IP地址或域名（支持IPv4和IPv6）。 | 
 **source** | **string** | 查询的数据源。如果留空，将使用默认的数据库。如果设置为 &#x60;commercial&#x60;，将调用商业级API，返回更详细的地理位置信息，但响应时间可能会稍长。 | 

### Return type

[**GetNetworkIpinfo200Response**](GetNetworkIpinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMyip

> GetNetworkIpinfo200Response GetNetworkMyip(ctx).Source(source).Execute()

获取你的公网IP及归属信息



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
	source := "source_example" // string | 查询的数据源。如果留空，将使用默认的数据库。如果设置为 `commercial`，将调用商业级API，返回更详细的地理位置信息，但响应时间可能会稍长。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkMyip(context.Background()).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkMyip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkMyip`: GetNetworkIpinfo200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkMyip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMyipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | 查询的数据源。如果留空，将使用默认的数据库。如果设置为 &#x60;commercial&#x60;，将调用商业级API，返回更详细的地理位置信息，但响应时间可能会稍长。 | 

### Return type

[**GetNetworkIpinfo200Response**](GetNetworkIpinfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPing

> GetNetworkPing200Response GetNetworkPing(ctx).Host(host).Execute()

从服务器Ping指定主机



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
	host := "cn.bing.com" // string | 你需要 Ping 的目标主机，可以是域名或IP地址。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkPing(context.Background()).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPing`: GetNetworkPing200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkPing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | 你需要 Ping 的目标主机，可以是域名或IP地址。 | 

### Return type

[**GetNetworkPing200Response**](GetNetworkPing200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPingmyip

> GetNetworkPingmyip200Response GetNetworkPingmyip(ctx).Execute()

从服务器Ping你的客户端IP



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
	resp, r, err := apiClient.NetworkAPI.GetNetworkPingmyip(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkPingmyip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPingmyip`: GetNetworkPingmyip200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkPingmyip`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPingmyipRequest struct via the builder pattern


### Return type

[**GetNetworkPingmyip200Response**](GetNetworkPingmyip200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPortscan

> GetNetworkPortscan200Response GetNetworkPortscan(ctx).Host(host).Port(port).Protocol(protocol).Execute()

扫描远程主机的指定端口



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
	host := "cn.bing.com" // string | 需要扫描的目标主机，可以是域名或IP地址。
	port := int32(80) // int32 | 需要扫描的端口号，范围是 1 到 65535。
	protocol := "tcp" // string | 扫描使用的协议，可以是 'tcp' 或 'udp'。 (optional) (default to "tcp")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkPortscan(context.Background()).Host(host).Port(port).Protocol(protocol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkPortscan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPortscan`: GetNetworkPortscan200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkPortscan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPortscanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | 需要扫描的目标主机，可以是域名或IP地址。 | 
 **port** | **int32** | 需要扫描的端口号，范围是 1 到 65535。 | 
 **protocol** | **string** | 扫描使用的协议，可以是 &#39;tcp&#39; 或 &#39;udp&#39;。 | [default to &quot;tcp&quot;]

### Return type

[**GetNetworkPortscan200Response**](GetNetworkPortscan200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkUrlstatus

> GetNetworkUrlstatus200Response GetNetworkUrlstatus(ctx).Url(url).Execute()

检查URL的可访问性状态



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
	url := "https://cn.bing.com" // string | 你需要检查其可访问性状态的完整URL。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkUrlstatus(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkUrlstatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkUrlstatus`: GetNetworkUrlstatus200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkUrlstatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkUrlstatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | 你需要检查其可访问性状态的完整URL。 | 

### Return type

[**GetNetworkUrlstatus200Response**](GetNetworkUrlstatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWhois

> GetNetworkWhois200Response GetNetworkWhois(ctx).Domain(domain).Format(format).Execute()

查询域名的WHOIS注册信息



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
	domain := "google.com" // string | 你需要查询WHOIS信息的域名。
	format := "json" // string | 返回格式。留空或为 'text' 时返回原始WHOIS文本，设为 'json' 时返回结构化JSON。 (optional) (default to "text")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkWhois(context.Background()).Domain(domain).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkWhois``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWhois`: GetNetworkWhois200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkWhois`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWhoisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | 你需要查询WHOIS信息的域名。 | 
 **format** | **string** | 返回格式。留空或为 &#39;text&#39; 时返回原始WHOIS文本，设为 &#39;json&#39; 时返回结构化JSON。 | [default to &quot;text&quot;]

### Return type

[**GetNetworkWhois200Response**](GetNetworkWhois200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWxdomain

> GetNetworkWxdomain200Response GetNetworkWxdomain(ctx).Domain(domain).Execute()

检查域名在微信中的访问状态



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
	domain := "qq.com" // string | 需要查询的域名。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAPI.GetNetworkWxdomain(context.Background()).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkWxdomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWxdomain`: GetNetworkWxdomain200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkWxdomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWxdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | 需要查询的域名。 | 

### Return type

[**GetNetworkWxdomain200Response**](GetNetworkWxdomain200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# uapi-sdk-go

![Banner](https://raw.githubusercontent.com/AxT-Team/uapi-sdk-go/main/banner.png)

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![Docs](https://img.shields.io/badge/Docs-uapis.cn-2EAE5D?style=flat-square)](https://uapis.cn/)
[![Go Reference](https://pkg.go.dev/badge/github.com/AxT-Team/uapi-sdk-go.svg)](https://pkg.go.dev/github.com/AxT-Team/uapi-sdk-go)

> [!NOTE]
> 所有接口的 Go 示例都可以在 [UApi](https://uapis.cn/docs/introduction) 的接口文档页面，向下滚动至 **快速启动** 区块后直接复制。

## 快速开始

```bash
go get github.com/AxT-Team/uapi-sdk-go@latest
```

```go
package main

import (
	"fmt"
	"github.com/AxT-Team/uapi-sdk-go/uapi"
)

func main() {
	client := uapi.New("https://uapis.cn", "YOUR_API_KEY")
	info, err := client.Misc().GetMiscHotboard(map[string]any{"type": "weibo"})
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}
```

这个接口默认只要传 `type` 就可以拿当前热榜。`time`、`keyword`、`time_start`、`time_end`、`limit`、`sources` 都是按场景再传的可选参数。

## 特性

现在你不再需要反反复复的查阅文档了。

只需在 IDE 中键入 `client.`，所有核心模块——如 `Social()`、`Game()`、`Image()`——即刻同步展现。进一步输入即可直接定位到 `GetSocialQqUserinfo`（获取 QQ 用户信息） 这样的具体方法，其名称与文档的 `operationId` 严格保持一致，确保了开发过程的直观与高效。

所有方法签名只接受真实且必需的参数。当你在构建请求时，IDE 会即时提示 `qq`、`username` 等键名，这彻底杜绝了在 `map[string]any` 中因键名拼写错误而导致的运行时错误。

针对 401、404、429 等标准 HTTP 响应，SDK 已将其统一映射为具名的错误类型。这些错误均附带 `status`、`request_id`、`retry_after_seconds` 等关键上下文信息，确保你在日志中能第一时间准确、快速地诊断问题。

基础域名、请求超时和 `User-Agent` 已预设为合理的默认值。但你完全拥有控制权，可以通过 `uapi.Client.Builder()` 模式灵活覆盖 Token、BaseURL 等配置。

如果你需要查看字段细节或内部逻辑，仓库中的 `./internal` 目录同步保留了由 `openapi-generator` 生成的完整结构体，随时可供参考。

## 响应元信息

每次请求完成后，SDK 会自动把响应 Header 解析成结构化的 `ResponseMeta`，你不用自己拆原始字符串。

成功时可以通过 `client.LastResponseMeta()` 读取，失败时可以从具体错误对象的 `Meta` 字段读取，两条路径拿到的是同一套字段。

```go
package main

import (
	"fmt"

	"github.com/AxT-Team/uapi-sdk-go/uapi"
)

func main() {
	client := uapi.New("https://uapis.cn", "YOUR_API_KEY")

	// 成功路径
	_, err := client.Social().GetSocialQqUserinfo(map[string]any{"qq": "10001"})
	if err != nil {
		panic(err)
	}
	if meta := client.LastResponseMeta(); meta != nil {
		if meta.BalanceRemainingCents != nil {
			fmt.Printf("余额剩余: %d 分\n", *meta.BalanceRemainingCents)
		}
		if meta.QuotaRemainingCredits != nil {
			fmt.Printf("资源包剩余: %d 积分\n", *meta.QuotaRemainingCredits)
		}
		fmt.Printf("Request ID: %s\n", meta.RequestID)
	}

	// 失败路径
	_, err = client.Social().GetSocialQqUserinfo(map[string]any{"qq": "10001"})
	if e, ok := err.(*uapi.ServiceBusyError); ok && e.Meta != nil {
		if e.Meta.RetryAfterSeconds != nil {
			fmt.Printf("限流，%ds 后重试\n", *e.Meta.RetryAfterSeconds)
		}
		fmt.Printf("Request ID: %s\n", e.Meta.RequestID)
	}
}
```

常用字段一览：

| 字段 | 说明 |
|------|------|
| `BalanceRemainingCents` | 账户余额剩余（分） |
| `QuotaRemainingCredits` | 资源包剩余积分 |
| `VisitorQuotaRemainingCredits` | 访客配额剩余积分 |
| `RetryAfterSeconds` | 触发限流后的建议等待时长 |
| `RequestID` | 请求唯一 ID，排障时使用 |
| `DebitStatus` | 本次计费状态 |
| `RateLimitPolicies` / `RateLimits` | 完整结构化限流策略数据 |

## 错误模型概览

| HTTP 状态码 | SDK 错误类型            | 附加信息                                                                          |
|-------------|------------------------|------------------------------------------------------------------------------------|
| 401/403     | `AuthenticationError`  | `status`、`request_id`                                                             |
| 404         | `NotFound`             | `status`、`request_id`                                                             |
| 400         | `ValidationError`      | `status`、`request_id`、`details`（如缺少参数、格式错误）                          |
| 429         | `RateLimitError`       | `status`、`request_id`、`retry_after_seconds`（用于决定何时重试）                  |
| 5xx         | `ServerError`          | `status`、`request_id`                                                             |
| 其他 4xx    | `ApiError`             | `status`、`request_id`、`code`、`message`、`details`（用于快速定位具体业务错误）   |

## 其他 SDK

| 语言        | 仓库地址                                                     |
|-------------|--------------------------------------------------------------|
| Go（当前）          | https://github.com/AxT-Team/uapi-sdk-go                      |
| Python      | https://github.com/AxT-Team/uapi-sdk-python                  |
| TypeScript| https://github.com/AxT-Team/uapi-sdk-typescript           |
| Browser (TypeScript/JavaScript)| https://github.com/AxT-Team/uapi-browser-sdk        |
| Java        | https://github.com/AxT-Team/uapi-sdk-java                    |
| PHP         | https://github.com/AxT-Team/uapi-sdk-php                     |
| C#          | https://github.com/AxT-Team/uapi-sdk-csharp                  |
| C++         | https://github.com/AxT-Team/uapi-sdk-cpp                     |
| Rust        | https://github.com/AxT-Team/uapi-sdk-rust                    |

## 文档

访问 [UApi文档首页](https://uapis.cn/docs/introduction) 并选择任意接口，向下滚动到 **快速启动** 区块即可看到最新的 Go 示例代码。



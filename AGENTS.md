# AGENTS.md — uapi-sdk-go

This file tells AI coding agents how to use the **official Go SDK** for the
[uapis.cn](https://uapis.cn) public API platform.

## What this package is

Idiomatic Go client for UAPI. Generated from the live OpenAPI 3.1 spec at
<https://uapis.cn/openapi.json>, so the method list, parameter shapes, and
return types stay in lock-step with the real API.

```bash
go get github.com/AxT-Team/uapi-sdk-go@latest
```

## Quick start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/AxT-Team/uapi-sdk-go/uapi"
)

func main() {
    ctx := context.Background()

    // Free-tier endpoints don't need an API key
    client := uapi.NewClient("https://uapis.cn")
    weather, err := client.Misc.GetMiscWeather(ctx, &uapi.GetMiscWeatherRequest{
        City: "北京",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", weather)
}
```

The client is grouped by tag (`Misc`, `Network`, `Text`, `Image`, `Social`,
`Translate`, `Search`, …). Method names match the OpenAPI `operationId`.

## Authentication

Free-tier endpoints work with no key. Paid endpoints take a key:

```go
client := uapi.NewClient("https://uapis.cn", uapi.WithAPIKey("sk_…"))
```

## Errors

Every method returns `(response, error)`. On non-2xx, `error` is a
`*uapi.APIError` with `Code`, `Success`, `Error`, and `RequestID`. Surface
`err.Error` verbatim; do not retry on `400` or `401`.

## Rate limits

Headers `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`,
`Retry-After` are exposed via `uapi.RateLimitFrom(resp)`. Honor them — back
off on `429` and obey `Retry-After`.

## Discovery

For unknown endpoints, fetch <https://uapis.cn/openapi.json>. The
`operationId` maps directly to a Go method (`get-misc-weather` →
`client.Misc.GetMiscWeather`).

## Related repos

- MCP server: <https://github.com/AxT-Team/uapi-mcp> — same endpoints as MCP tools.
- Skills bundle: <https://github.com/AxT-Team/uapi-agent-skills>.
- Other languages: `uapi-sdk-typescript`, `uapi-sdk-python`, `uapi-sdk-rust`,
  `uapi-sdk-java`, `uapi-sdk-csharp`, `uapi-sdk-cpp`, `uapi-sdk-php`.

package uapi

import (
	"strings"
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
	baseURL string
	token   string
	cli     *fasthttp.Client
}

func New(baseURL, token string) *Client {
	return &Client{ baseURL: baseURL, token: token, cli: &fasthttp.Client{} }
}

func (c *Client) do(method, path string, q map[string]string, body any) (any, error) {
	url := c.baseURL + path
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req); defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.Set("Content-Type", "application/json")
	if c.token != "" { req.Header.Set("Authorization", "Bearer "+c.token) }
	if body != nil {
		b, _ := json.Marshal(body)
		req.SetBody(b)
	}

	if err := c.cli.DoTimeout(req, resp, 15*time.Second); err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	status := resp.StatusCode()
	if status >= 400 {
		return nil, mapError(status, resp.Body())
	}
	ct := string(resp.Header.ContentType())
	if ct == "application/json" {
		var v any
		if err := json.Unmarshal(resp.Body(), &v); err == nil { return v, nil }
	}
	return resp.Body(), nil
}

// Facade by tag
type ClipzyZaiXianJianTieBanApi struct { c *Client }
func (c *Client) ClipzyZaiXianJianTieBan() *ClipzyZaiXianJianTieBanApi { return &ClipzyZaiXianJianTieBanApi{c: c} }
// 步骤2 (方法一): 获取加密数据
func (api *ClipzyZaiXianJianTieBanApi) GetClipzyGet(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["id"]; ok { q["id"] = fmt.Sprint(v) }
	path := "/api/v1/api/get"
	return api.c.do("GET", path, q, nil)
}
// 步骤2 (方法二): 获取原始文本
func (api *ClipzyZaiXianJianTieBanApi) GetClipzyRaw(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["key"]; ok { q["key"] = fmt.Sprint(v) }
	path := "/api/v1/api/raw/{id}"
	if v, ok := args["id"]; ok { path = strings.ReplaceAll(path, "{"+"id"+"}", fmt.Sprint(v)) }
	return api.c.do("GET", path, q, nil)
}
// 步骤1：上传加密数据
func (api *ClipzyZaiXianJianTieBanApi) PostClipzyStore(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["compressedData"]; ok { body["compressedData"] = v }
	if v, ok := args["ttl"]; ok { body["ttl"] = v }
	path := "/api/v1/api/store"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type ConvertApi struct { c *Client }
func (c *Client) Convert() *ConvertApi { return &ConvertApi{c: c} }
// 时间戳转换
func (api *ConvertApi) GetConvertUnixtime(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["time"]; ok { q["time"] = fmt.Sprint(v) }
	path := "/api/v1/convert/unixtime"
	return api.c.do("GET", path, q, nil)
}
// JSON 格式化
func (api *ConvertApi) PostConvertJson(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["content"]; ok { body["content"] = v }
	path := "/api/v1/convert/json"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type DailyApi struct { c *Client }
func (c *Client) Daily() *DailyApi { return &DailyApi{c: c} }
// 每日新闻图
func (api *DailyApi) GetDailyNewsImage(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/daily/news-image"
	return api.c.do("GET", path, q, nil)
}
type GameApi struct { c *Client }
func (c *Client) Game() *GameApi { return &GameApi{c: c} }
// Epic 免费游戏
func (api *GameApi) GetGameEpicFree(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/game/epic-free"
	return api.c.do("GET", path, q, nil)
}
// 查询 MC 曾用名
func (api *GameApi) GetGameMinecraftHistoryid(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["name"]; ok { q["name"] = fmt.Sprint(v) }
	if v, ok := args["uuid"]; ok { q["uuid"] = fmt.Sprint(v) }
	path := "/api/v1/game/minecraft/historyid"
	return api.c.do("GET", path, q, nil)
}
// 查询 MC 服务器
func (api *GameApi) GetGameMinecraftServerstatus(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["server"]; ok { q["server"] = fmt.Sprint(v) }
	path := "/api/v1/game/minecraft/serverstatus"
	return api.c.do("GET", path, q, nil)
}
// 查询 MC 玩家
func (api *GameApi) GetGameMinecraftUserinfo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["username"]; ok { q["username"] = fmt.Sprint(v) }
	path := "/api/v1/game/minecraft/userinfo"
	return api.c.do("GET", path, q, nil)
}
// 查询 Steam 用户
func (api *GameApi) GetGameSteamSummary(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["steamid"]; ok { q["steamid"] = fmt.Sprint(v) }
	if v, ok := args["id"]; ok { q["id"] = fmt.Sprint(v) }
	if v, ok := args["id3"]; ok { q["id3"] = fmt.Sprint(v) }
	if v, ok := args["key"]; ok { q["key"] = fmt.Sprint(v) }
	path := "/api/v1/game/steam/summary"
	return api.c.do("GET", path, q, nil)
}
type ImageApi struct { c *Client }
func (c *Client) Image() *ImageApi { return &ImageApi{c: c} }
// 获取Gravatar头像
func (api *ImageApi) GetAvatarGravatar(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["email"]; ok { q["email"] = fmt.Sprint(v) }
	if v, ok := args["hash"]; ok { q["hash"] = fmt.Sprint(v) }
	if v, ok := args["s"]; ok { q["s"] = fmt.Sprint(v) }
	if v, ok := args["d"]; ok { q["d"] = fmt.Sprint(v) }
	if v, ok := args["r"]; ok { q["r"] = fmt.Sprint(v) }
	path := "/api/v1/avatar/gravatar"
	return api.c.do("GET", path, q, nil)
}
// 必应壁纸
func (api *ImageApi) GetImageBingDaily(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/image/bing-daily"
	return api.c.do("GET", path, q, nil)
}
// 摸头 GIF
func (api *ImageApi) GetImageMotou(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["qq"]; ok { q["qq"] = fmt.Sprint(v) }
	if v, ok := args["bg_color"]; ok { q["bg_color"] = fmt.Sprint(v) }
	path := "/api/v1/image/motou"
	return api.c.do("GET", path, q, nil)
}
// 生成二维码
func (api *ImageApi) GetImageQrcode(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["text"]; ok { q["text"] = fmt.Sprint(v) }
	if v, ok := args["size"]; ok { q["size"] = fmt.Sprint(v) }
	if v, ok := args["format"]; ok { q["format"] = fmt.Sprint(v) }
	path := "/api/v1/image/qrcode"
	return api.c.do("GET", path, q, nil)
}
// 图片转 Base64
func (api *ImageApi) GetImageTobase64(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["url"]; ok { q["url"] = fmt.Sprint(v) }
	path := "/api/v1/image/tobase64"
	return api.c.do("GET", path, q, nil)
}
// 无损压缩图片
func (api *ImageApi) PostImageCompress(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["level"]; ok { q["level"] = fmt.Sprint(v) }
	if v, ok := args["format"]; ok { q["format"] = fmt.Sprint(v) }
	body = make(map[string]any)
	if v, ok := args["file"]; ok { body["file"] = v }
	path := "/api/v1/image/compress"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 通过Base64编码上传图片
func (api *ImageApi) PostImageFrombase64(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["imageData"]; ok { body["imageData"] = v }
	path := "/api/v1/image/frombase64"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 摸头 GIF (上传)
func (api *ImageApi) PostImageMotou(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["bg_color"]; ok { body["bg_color"] = v }
	if v, ok := args["file"]; ok { body["file"] = v }
	if v, ok := args["image_url"]; ok { body["image_url"] = v }
	path := "/api/v1/image/motou"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 生成你们怎么不说话了表情包
func (api *ImageApi) PostImageSpeechless(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["bottom_text"]; ok { body["bottom_text"] = v }
	if v, ok := args["top_text"]; ok { body["top_text"] = v }
	path := "/api/v1/image/speechless"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// SVG转图片
func (api *ImageApi) PostImageSvg(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["format"]; ok { q["format"] = fmt.Sprint(v) }
	if v, ok := args["width"]; ok { q["width"] = fmt.Sprint(v) }
	if v, ok := args["height"]; ok { q["height"] = fmt.Sprint(v) }
	if v, ok := args["quality"]; ok { q["quality"] = fmt.Sprint(v) }
	body = make(map[string]any)
	if v, ok := args["file"]; ok { body["file"] = v }
	path := "/api/v1/image/svg"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type MiscApi struct { c *Client }
func (c *Client) Misc() *MiscApi { return &MiscApi{c: c} }
// 程序员历史事件
func (api *MiscApi) GetHistoryProgrammer(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["month"]; ok { q["month"] = fmt.Sprint(v) }
	if v, ok := args["day"]; ok { q["day"] = fmt.Sprint(v) }
	path := "/api/v1/history/programmer"
	return api.c.do("GET", path, q, nil)
}
// 程序员历史上的今天
func (api *MiscApi) GetHistoryProgrammerToday(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/history/programmer/today"
	return api.c.do("GET", path, q, nil)
}
// 查询热榜
func (api *MiscApi) GetMiscHotboard(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["type"]; ok { q["type"] = fmt.Sprint(v) }
	path := "/api/v1/misc/hotboard"
	return api.c.do("GET", path, q, nil)
}
// 查询手机归属地
func (api *MiscApi) GetMiscPhoneinfo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["phone"]; ok { q["phone"] = fmt.Sprint(v) }
	path := "/api/v1/misc/phoneinfo"
	return api.c.do("GET", path, q, nil)
}
// 随机数生成
func (api *MiscApi) GetMiscRandomnumber(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["min"]; ok { q["min"] = fmt.Sprint(v) }
	if v, ok := args["max"]; ok { q["max"] = fmt.Sprint(v) }
	if v, ok := args["count"]; ok { q["count"] = fmt.Sprint(v) }
	if v, ok := args["allow_repeat"]; ok { q["allow_repeat"] = fmt.Sprint(v) }
	if v, ok := args["allow_decimal"]; ok { q["allow_decimal"] = fmt.Sprint(v) }
	if v, ok := args["decimal_places"]; ok { q["decimal_places"] = fmt.Sprint(v) }
	path := "/api/v1/misc/randomnumber"
	return api.c.do("GET", path, q, nil)
}
// 转换时间戳 (旧版，推荐使用/convert/unixtime)
func (api *MiscApi) GetMiscTimestamp(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["ts"]; ok { q["ts"] = fmt.Sprint(v) }
	path := "/api/v1/misc/timestamp"
	return api.c.do("GET", path, q, nil)
}
// 获取支持的快递公司列表
func (api *MiscApi) GetMiscTrackingCarriers(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/misc/tracking/carriers"
	return api.c.do("GET", path, q, nil)
}
// 识别快递公司
func (api *MiscApi) GetMiscTrackingDetect(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["tracking_number"]; ok { q["tracking_number"] = fmt.Sprint(v) }
	path := "/api/v1/misc/tracking/detect"
	return api.c.do("GET", path, q, nil)
}
// 查询快递物流信息
func (api *MiscApi) GetMiscTrackingQuery(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["tracking_number"]; ok { q["tracking_number"] = fmt.Sprint(v) }
	if v, ok := args["carrier_code"]; ok { q["carrier_code"] = fmt.Sprint(v) }
	path := "/api/v1/misc/tracking/query"
	return api.c.do("GET", path, q, nil)
}
// 查询天气
func (api *MiscApi) GetMiscWeather(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["city"]; ok { q["city"] = fmt.Sprint(v) }
	if v, ok := args["adcode"]; ok { q["adcode"] = fmt.Sprint(v) }
	if v, ok := args["extended"]; ok { q["extended"] = fmt.Sprint(v) }
	if v, ok := args["indices"]; ok { q["indices"] = fmt.Sprint(v) }
	if v, ok := args["forecast"]; ok { q["forecast"] = fmt.Sprint(v) }
	path := "/api/v1/misc/weather"
	return api.c.do("GET", path, q, nil)
}
// 查询世界时间
func (api *MiscApi) GetMiscWorldtime(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["city"]; ok { q["city"] = fmt.Sprint(v) }
	path := "/api/v1/misc/worldtime"
	return api.c.do("GET", path, q, nil)
}
// 计算两个日期之间的时间差值
func (api *MiscApi) PostMiscDateDiff(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["end_date"]; ok { body["end_date"] = v }
	if v, ok := args["format"]; ok { body["format"] = v }
	if v, ok := args["start_date"]; ok { body["start_date"] = v }
	path := "/api/v1/misc/date-diff"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type NetworkApi struct { c *Client }
func (c *Client) Network() *NetworkApi { return &NetworkApi{c: c} }
// 执行DNS解析查询
func (api *NetworkApi) GetNetworkDns(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["domain"]; ok { q["domain"] = fmt.Sprint(v) }
	if v, ok := args["type"]; ok { q["type"] = fmt.Sprint(v) }
	path := "/api/v1/network/dns"
	return api.c.do("GET", path, q, nil)
}
// 查询域名ICP备案信息
func (api *NetworkApi) GetNetworkIcp(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["domain"]; ok { q["domain"] = fmt.Sprint(v) }
	path := "/api/v1/network/icp"
	return api.c.do("GET", path, q, nil)
}
// 查询 IP
func (api *NetworkApi) GetNetworkIpinfo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["ip"]; ok { q["ip"] = fmt.Sprint(v) }
	if v, ok := args["source"]; ok { q["source"] = fmt.Sprint(v) }
	path := "/api/v1/network/ipinfo"
	return api.c.do("GET", path, q, nil)
}
// 查询我的 IP
func (api *NetworkApi) GetNetworkMyip(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["source"]; ok { q["source"] = fmt.Sprint(v) }
	path := "/api/v1/network/myip"
	return api.c.do("GET", path, q, nil)
}
// Ping 主机
func (api *NetworkApi) GetNetworkPing(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["host"]; ok { q["host"] = fmt.Sprint(v) }
	path := "/api/v1/network/ping"
	return api.c.do("GET", path, q, nil)
}
// Ping 我的 IP
func (api *NetworkApi) GetNetworkPingmyip(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/network/pingmyip"
	return api.c.do("GET", path, q, nil)
}
// 端口扫描
func (api *NetworkApi) GetNetworkPortscan(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["host"]; ok { q["host"] = fmt.Sprint(v) }
	if v, ok := args["port"]; ok { q["port"] = fmt.Sprint(v) }
	if v, ok := args["protocol"]; ok { q["protocol"] = fmt.Sprint(v) }
	path := "/api/v1/network/portscan"
	return api.c.do("GET", path, q, nil)
}
// 检查URL的可访问性状态
func (api *NetworkApi) GetNetworkUrlstatus(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["url"]; ok { q["url"] = fmt.Sprint(v) }
	path := "/api/v1/network/urlstatus"
	return api.c.do("GET", path, q, nil)
}
// 查询域名的WHOIS注册信息
func (api *NetworkApi) GetNetworkWhois(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["domain"]; ok { q["domain"] = fmt.Sprint(v) }
	if v, ok := args["format"]; ok { q["format"] = fmt.Sprint(v) }
	path := "/api/v1/network/whois"
	return api.c.do("GET", path, q, nil)
}
// 检查域名在微信中的访问状态
func (api *NetworkApi) GetNetworkWxdomain(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["domain"]; ok { q["domain"] = fmt.Sprint(v) }
	path := "/api/v1/network/wxdomain"
	return api.c.do("GET", path, q, nil)
}
type PoemApi struct { c *Client }
func (c *Client) Poem() *PoemApi { return &PoemApi{c: c} }
// 一言
func (api *PoemApi) GetSaying(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/saying"
	return api.c.do("GET", path, q, nil)
}
type RandomApi struct { c *Client }
func (c *Client) Random() *RandomApi { return &RandomApi{c: c} }
// 答案之书
func (api *RandomApi) GetAnswerbookAsk(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["question"]; ok { q["question"] = fmt.Sprint(v) }
	path := "/api/v1/answerbook/ask"
	return api.c.do("GET", path, q, nil)
}
// 随机图片
func (api *RandomApi) GetRandomImage(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["category"]; ok { q["category"] = fmt.Sprint(v) }
	if v, ok := args["type"]; ok { q["type"] = fmt.Sprint(v) }
	path := "/api/v1/random/image"
	return api.c.do("GET", path, q, nil)
}
// 随机字符串
func (api *RandomApi) GetRandomString(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["length"]; ok { q["length"] = fmt.Sprint(v) }
	if v, ok := args["type"]; ok { q["type"] = fmt.Sprint(v) }
	path := "/api/v1/random/string"
	return api.c.do("GET", path, q, nil)
}
// 答案之书 (POST)
func (api *RandomApi) PostAnswerbookAsk(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["question"]; ok { body["question"] = v }
	path := "/api/v1/answerbook/ask"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type SocialApi struct { c *Client }
func (c *Client) Social() *SocialApi { return &SocialApi{c: c} }
// 查询 GitHub 仓库
func (api *SocialApi) GetGithubRepo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["repo"]; ok { q["repo"] = fmt.Sprint(v) }
	path := "/api/v1/github/repo"
	return api.c.do("GET", path, q, nil)
}
// 查询 B站投稿
func (api *SocialApi) GetSocialBilibiliArchives(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["mid"]; ok { q["mid"] = fmt.Sprint(v) }
	if v, ok := args["keywords"]; ok { q["keywords"] = fmt.Sprint(v) }
	if v, ok := args["orderby"]; ok { q["orderby"] = fmt.Sprint(v) }
	if v, ok := args["ps"]; ok { q["ps"] = fmt.Sprint(v) }
	if v, ok := args["pn"]; ok { q["pn"] = fmt.Sprint(v) }
	path := "/api/v1/social/bilibili/archives"
	return api.c.do("GET", path, q, nil)
}
// 查询 B站直播间
func (api *SocialApi) GetSocialBilibiliLiveroom(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["mid"]; ok { q["mid"] = fmt.Sprint(v) }
	if v, ok := args["room_id"]; ok { q["room_id"] = fmt.Sprint(v) }
	path := "/api/v1/social/bilibili/liveroom"
	return api.c.do("GET", path, q, nil)
}
// 查询 B站评论
func (api *SocialApi) GetSocialBilibiliReplies(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["oid"]; ok { q["oid"] = fmt.Sprint(v) }
	if v, ok := args["sort"]; ok { q["sort"] = fmt.Sprint(v) }
	if v, ok := args["ps"]; ok { q["ps"] = fmt.Sprint(v) }
	if v, ok := args["pn"]; ok { q["pn"] = fmt.Sprint(v) }
	path := "/api/v1/social/bilibili/replies"
	return api.c.do("GET", path, q, nil)
}
// 查询 B站用户
func (api *SocialApi) GetSocialBilibiliUserinfo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["uid"]; ok { q["uid"] = fmt.Sprint(v) }
	path := "/api/v1/social/bilibili/userinfo"
	return api.c.do("GET", path, q, nil)
}
// 查询 B站视频
func (api *SocialApi) GetSocialBilibiliVideoinfo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["aid"]; ok { q["aid"] = fmt.Sprint(v) }
	if v, ok := args["bvid"]; ok { q["bvid"] = fmt.Sprint(v) }
	path := "/api/v1/social/bilibili/videoinfo"
	return api.c.do("GET", path, q, nil)
}
// 查询 QQ 群信息
func (api *SocialApi) GetSocialQqGroupinfo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["group_id"]; ok { q["group_id"] = fmt.Sprint(v) }
	path := "/api/v1/social/qq/groupinfo"
	return api.c.do("GET", path, q, nil)
}
// 查询 QQ 信息
func (api *SocialApi) GetSocialQqUserinfo(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["qq"]; ok { q["qq"] = fmt.Sprint(v) }
	path := "/api/v1/social/qq/userinfo"
	return api.c.do("GET", path, q, nil)
}
type StatusApi struct { c *Client }
func (c *Client) Status() *StatusApi { return &StatusApi{c: c} }
// 限流状态
func (api *StatusApi) GetStatusRatelimit(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/status/ratelimit"
	return api.c.do("GET", path, q, nil)
}
// 获取API端点使用统计
func (api *StatusApi) GetStatusUsage(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["path"]; ok { q["path"] = fmt.Sprint(v) }
	path := "/api/v1/status/usage"
	return api.c.do("GET", path, q, nil)
}
type TextApi struct { c *Client }
func (c *Client) Text() *TextApi { return &TextApi{c: c} }
// MD5 哈希
func (api *TextApi) GetTextMd5(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["text"]; ok { q["text"] = fmt.Sprint(v) }
	path := "/api/v1/text/md5"
	return api.c.do("GET", path, q, nil)
}
// AES 解密
func (api *TextApi) PostTextAesDecrypt(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["key"]; ok { body["key"] = v }
	if v, ok := args["nonce"]; ok { body["nonce"] = v }
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/aes/decrypt"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// AES高级解密
func (api *TextApi) PostTextAesDecryptAdvanced(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["iv"]; ok { body["iv"] = v }
	if v, ok := args["key"]; ok { body["key"] = v }
	if v, ok := args["mode"]; ok { body["mode"] = v }
	if v, ok := args["padding"]; ok { body["padding"] = v }
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/aes/decrypt-advanced"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// AES 加密
func (api *TextApi) PostTextAesEncrypt(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["key"]; ok { body["key"] = v }
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/aes/encrypt"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// AES高级加密
func (api *TextApi) PostTextAesEncryptAdvanced(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["iv"]; ok { body["iv"] = v }
	if v, ok := args["key"]; ok { body["key"] = v }
	if v, ok := args["mode"]; ok { body["mode"] = v }
	if v, ok := args["output_format"]; ok { body["output_format"] = v }
	if v, ok := args["padding"]; ok { body["padding"] = v }
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/aes/encrypt-advanced"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 文本分析
func (api *TextApi) PostTextAnalyze(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/analyze"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// Base64 解码
func (api *TextApi) PostTextBase64Decode(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/base64/decode"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// Base64 编码
func (api *TextApi) PostTextBase64Encode(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/base64/encode"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 格式转换
func (api *TextApi) PostTextConvert(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["from"]; ok { body["from"] = v }
	if v, ok := args["options"]; ok { body["options"] = v }
	if v, ok := args["text"]; ok { body["text"] = v }
	if v, ok := args["to"]; ok { body["to"] = v }
	path := "/api/v1/text/convert"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// MD5 哈希 (POST)
func (api *TextApi) PostTextMd5(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/md5"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// MD5 校验
func (api *TextApi) PostTextMd5Verify(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["hash"]; ok { body["hash"] = v }
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/md5/verify"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type TranslateApi struct { c *Client }
func (c *Client) Translate() *TranslateApi { return &TranslateApi{c: c} }
// AI翻译配置
func (api *TranslateApi) GetAiTranslateLanguages(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/ai/translate/languages"
	return api.c.do("GET", path, q, nil)
}
// AI智能翻译
func (api *TranslateApi) PostAiTranslate(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["target_lang"]; ok { q["target_lang"] = fmt.Sprint(v) }
	body = make(map[string]any)
	if v, ok := args["context"]; ok { body["context"] = v }
	if v, ok := args["fast_mode"]; ok { body["fast_mode"] = v }
	if v, ok := args["max_concurrency"]; ok { body["max_concurrency"] = v }
	if v, ok := args["preserve_format"]; ok { body["preserve_format"] = v }
	if v, ok := args["source_lang"]; ok { body["source_lang"] = v }
	if v, ok := args["style"]; ok { body["style"] = v }
	if v, ok := args["text"]; ok { body["text"] = v }
	if v, ok := args["texts"]; ok { body["texts"] = v }
	path := "/api/v1/ai/translate"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 流式翻译（中英互译）
func (api *TranslateApi) PostTranslateStream(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["from_lang"]; ok { body["from_lang"] = v }
	if v, ok := args["query"]; ok { body["query"] = v }
	if v, ok := args["to_lang"]; ok { body["to_lang"] = v }
	if v, ok := args["tone"]; ok { body["tone"] = v }
	path := "/api/v1/translate/stream"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 翻译
func (api *TranslateApi) PostTranslateText(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["to_lang"]; ok { q["to_lang"] = fmt.Sprint(v) }
	body = make(map[string]any)
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/translate/text"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type WebparseApi struct { c *Client }
func (c *Client) Webparse() *WebparseApi { return &WebparseApi{c: c} }
// 转换任务状态
func (api *WebparseApi) GetWebTomarkdownAsyncStatus(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/web/tomarkdown/async/{task_id}"
	if v, ok := args["task_id"]; ok { path = strings.ReplaceAll(path, "{"+"task_id"+"}", fmt.Sprint(v)) }
	return api.c.do("GET", path, q, nil)
}
// 提取网页图片
func (api *WebparseApi) GetWebparseExtractimages(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["url"]; ok { q["url"] = fmt.Sprint(v) }
	path := "/api/v1/webparse/extractimages"
	return api.c.do("GET", path, q, nil)
}
// 网页元数据
func (api *WebparseApi) GetWebparseMetadata(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["url"]; ok { q["url"] = fmt.Sprint(v) }
	path := "/api/v1/webparse/metadata"
	return api.c.do("GET", path, q, nil)
}
// 网页转 Markdown
func (api *WebparseApi) PostWebTomarkdownAsync(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["url"]; ok { q["url"] = fmt.Sprint(v) }
	path := "/api/v1/web/tomarkdown/async"
	return api.c.do("POST", path, q, nil)
}
type MinGanCiShiBieApi struct { c *Client }
func (c *Client) MinGanCiShiBie() *MinGanCiShiBieApi { return &MinGanCiShiBieApi{c: c} }
// 敏感词分析 (GET)
func (api *MinGanCiShiBieApi) GetSensitiveWordAnalyzeQuery(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	if v, ok := args["keyword"]; ok { q["keyword"] = fmt.Sprint(v) }
	path := "/api/v1/sensitive-word/analyze-query"
	return api.c.do("GET", path, q, nil)
}
// 分析敏感词
func (api *MinGanCiShiBieApi) PostSensitiveWordAnalyze(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["keywords"]; ok { body["keywords"] = v }
	path := "/api/v1/sensitive-word/analyze"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
// 敏感词检测（快速）
func (api *MinGanCiShiBieApi) PostSensitiveWordQuickCheck(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["text"]; ok { body["text"] = v }
	path := "/api/v1/text/profanitycheck"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}
type ZhiNengSouSuoApi struct { c *Client }
func (c *Client) ZhiNengSouSuo() *ZhiNengSouSuoApi { return &ZhiNengSouSuoApi{c: c} }
// 搜索引擎配置
func (api *ZhiNengSouSuoApi) GetSearchEngines(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	path := "/api/v1/search/engines"
	return api.c.do("GET", path, q, nil)
}
// 智能搜索
func (api *ZhiNengSouSuoApi) PostSearchAggregate(args map[string]any) (any, error) {
	q := map[string]string{}
	var body map[string]any
	body = make(map[string]any)
	if v, ok := args["fetch_full"]; ok { body["fetch_full"] = v }
	if v, ok := args["filetype"]; ok { body["filetype"] = v }
	if v, ok := args["query"]; ok { body["query"] = v }
	if v, ok := args["site"]; ok { body["site"] = v }
	if v, ok := args["sort"]; ok { body["sort"] = v }
	if v, ok := args["time_range"]; ok { body["time_range"] = v }
	if v, ok := args["timeout_ms"]; ok { body["timeout_ms"] = v }
	path := "/api/v1/search/aggregate"
	if len(body) == 0 { body = nil }
	return api.c.do("POST", path, q, body)
}

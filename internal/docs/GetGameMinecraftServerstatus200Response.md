# GetGameMinecraftServerstatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FaviconUrl** | Pointer to **string** | 服务器图标的 Base64 Data URI。你可以直接在 &#x60;&lt;img&gt;&#x60; 标签的 &#x60;src&#x60; 属性中使用它。 | [optional] 
**Ip** | Pointer to **string** | 服务器解析后的IP地址。 | [optional] 
**MaxPlayers** | Pointer to **int32** | 服务器配置的最大玩家容量。 | [optional] 
**MotdClean** | Pointer to **string** | 纯文本格式的服务器MOTD（每日消息），去除了所有颜色和格式代码。 | [optional] 
**MotdHtml** | Pointer to **string** | HTML格式的服务器MOTD，保留了颜色和样式，方便你在网页上直接渲染。 | [optional] 
**Online** | Pointer to **bool** | 服务器当前是否在线。 | [optional] 
**Players** | Pointer to **int32** | 当前在线的玩家数量。 | [optional] 
**Port** | Pointer to **int32** | 服务器使用的端口。 | [optional] 
**Version** | Pointer to **string** | 服务器报告的版本信息。 | [optional] 

## Methods

### NewGetGameMinecraftServerstatus200Response

`func NewGetGameMinecraftServerstatus200Response() *GetGameMinecraftServerstatus200Response`

NewGetGameMinecraftServerstatus200Response instantiates a new GetGameMinecraftServerstatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameMinecraftServerstatus200ResponseWithDefaults

`func NewGetGameMinecraftServerstatus200ResponseWithDefaults() *GetGameMinecraftServerstatus200Response`

NewGetGameMinecraftServerstatus200ResponseWithDefaults instantiates a new GetGameMinecraftServerstatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaviconUrl

`func (o *GetGameMinecraftServerstatus200Response) GetFaviconUrl() string`

GetFaviconUrl returns the FaviconUrl field if non-nil, zero value otherwise.

### GetFaviconUrlOk

`func (o *GetGameMinecraftServerstatus200Response) GetFaviconUrlOk() (*string, bool)`

GetFaviconUrlOk returns a tuple with the FaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaviconUrl

`func (o *GetGameMinecraftServerstatus200Response) SetFaviconUrl(v string)`

SetFaviconUrl sets FaviconUrl field to given value.

### HasFaviconUrl

`func (o *GetGameMinecraftServerstatus200Response) HasFaviconUrl() bool`

HasFaviconUrl returns a boolean if a field has been set.

### GetIp

`func (o *GetGameMinecraftServerstatus200Response) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetGameMinecraftServerstatus200Response) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetGameMinecraftServerstatus200Response) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetGameMinecraftServerstatus200Response) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMaxPlayers

`func (o *GetGameMinecraftServerstatus200Response) GetMaxPlayers() int32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *GetGameMinecraftServerstatus200Response) GetMaxPlayersOk() (*int32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *GetGameMinecraftServerstatus200Response) SetMaxPlayers(v int32)`

SetMaxPlayers sets MaxPlayers field to given value.

### HasMaxPlayers

`func (o *GetGameMinecraftServerstatus200Response) HasMaxPlayers() bool`

HasMaxPlayers returns a boolean if a field has been set.

### GetMotdClean

`func (o *GetGameMinecraftServerstatus200Response) GetMotdClean() string`

GetMotdClean returns the MotdClean field if non-nil, zero value otherwise.

### GetMotdCleanOk

`func (o *GetGameMinecraftServerstatus200Response) GetMotdCleanOk() (*string, bool)`

GetMotdCleanOk returns a tuple with the MotdClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdClean

`func (o *GetGameMinecraftServerstatus200Response) SetMotdClean(v string)`

SetMotdClean sets MotdClean field to given value.

### HasMotdClean

`func (o *GetGameMinecraftServerstatus200Response) HasMotdClean() bool`

HasMotdClean returns a boolean if a field has been set.

### GetMotdHtml

`func (o *GetGameMinecraftServerstatus200Response) GetMotdHtml() string`

GetMotdHtml returns the MotdHtml field if non-nil, zero value otherwise.

### GetMotdHtmlOk

`func (o *GetGameMinecraftServerstatus200Response) GetMotdHtmlOk() (*string, bool)`

GetMotdHtmlOk returns a tuple with the MotdHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdHtml

`func (o *GetGameMinecraftServerstatus200Response) SetMotdHtml(v string)`

SetMotdHtml sets MotdHtml field to given value.

### HasMotdHtml

`func (o *GetGameMinecraftServerstatus200Response) HasMotdHtml() bool`

HasMotdHtml returns a boolean if a field has been set.

### GetOnline

`func (o *GetGameMinecraftServerstatus200Response) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *GetGameMinecraftServerstatus200Response) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *GetGameMinecraftServerstatus200Response) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *GetGameMinecraftServerstatus200Response) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetPlayers

`func (o *GetGameMinecraftServerstatus200Response) GetPlayers() int32`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GetGameMinecraftServerstatus200Response) GetPlayersOk() (*int32, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GetGameMinecraftServerstatus200Response) SetPlayers(v int32)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *GetGameMinecraftServerstatus200Response) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetPort

`func (o *GetGameMinecraftServerstatus200Response) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetGameMinecraftServerstatus200Response) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetGameMinecraftServerstatus200Response) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetGameMinecraftServerstatus200Response) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetVersion

`func (o *GetGameMinecraftServerstatus200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetGameMinecraftServerstatus200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetGameMinecraftServerstatus200Response) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetGameMinecraftServerstatus200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



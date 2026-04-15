# GetNetworkWhois200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Whois** | Pointer to **map[string]interface{}** | 结构化 WHOIS 信息，返回经过解析的 JSON 对象，通常包含域名信息、注册商信息、注册人信息以及注册日期、更新时间、到期时间等字段。 有些字段会因域名注册局和隐私保护设置而有所不同噢。 | [optional] 

## Methods

### NewGetNetworkWhois200Response

`func NewGetNetworkWhois200Response() *GetNetworkWhois200Response`

NewGetNetworkWhois200Response instantiates a new GetNetworkWhois200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWhois200ResponseWithDefaults

`func NewGetNetworkWhois200ResponseWithDefaults() *GetNetworkWhois200Response`

NewGetNetworkWhois200ResponseWithDefaults instantiates a new GetNetworkWhois200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhois

`func (o *GetNetworkWhois200Response) GetWhois() map[string]interface{}`

GetWhois returns the Whois field if non-nil, zero value otherwise.

### GetWhoisOk

`func (o *GetNetworkWhois200Response) GetWhoisOk() (*map[string]interface{}, bool)`

GetWhoisOk returns a tuple with the Whois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhois

`func (o *GetNetworkWhois200Response) SetWhois(v map[string]interface{})`

SetWhois sets Whois field to given value.

### HasWhois

`func (o *GetNetworkWhois200Response) HasWhois() bool`

HasWhois returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



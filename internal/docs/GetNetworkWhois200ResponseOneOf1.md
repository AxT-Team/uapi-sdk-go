# GetNetworkWhois200ResponseOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Whois** | Pointer to **map[string]interface{}** | ### 结构化WHOIS信息  返回经过解析的JSON对象，包含以下主要部分：  - **域名信息**: 包含域名ID、注册状态、DNS服务器等 - **注册商信息**: 注册服务商的详细信息 - **注册人信息**: 域名所有者的相关信息（可能因隐私保护而部分隐藏） - **重要日期**: 包括注册日期、更新日期和到期日期 | [optional] 

## Methods

### NewGetNetworkWhois200ResponseOneOf1

`func NewGetNetworkWhois200ResponseOneOf1() *GetNetworkWhois200ResponseOneOf1`

NewGetNetworkWhois200ResponseOneOf1 instantiates a new GetNetworkWhois200ResponseOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWhois200ResponseOneOf1WithDefaults

`func NewGetNetworkWhois200ResponseOneOf1WithDefaults() *GetNetworkWhois200ResponseOneOf1`

NewGetNetworkWhois200ResponseOneOf1WithDefaults instantiates a new GetNetworkWhois200ResponseOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhois

`func (o *GetNetworkWhois200ResponseOneOf1) GetWhois() map[string]interface{}`

GetWhois returns the Whois field if non-nil, zero value otherwise.

### GetWhoisOk

`func (o *GetNetworkWhois200ResponseOneOf1) GetWhoisOk() (*map[string]interface{}, bool)`

GetWhoisOk returns a tuple with the Whois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhois

`func (o *GetNetworkWhois200ResponseOneOf1) SetWhois(v map[string]interface{})`

SetWhois sets Whois field to given value.

### HasWhois

`func (o *GetNetworkWhois200ResponseOneOf1) HasWhois() bool`

HasWhois returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



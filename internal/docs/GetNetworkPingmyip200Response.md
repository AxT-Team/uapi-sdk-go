# GetNetworkPingmyip200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** | 当前客户端的公网 IP 地址。 | [optional] 
**Message** | Pointer to **string** | 操作结果说明。成功时通常会附带平均延迟信息。 | [optional] 
**PingSuccessful** | Pointer to **bool** | 是否成功完成对当前客户端 IP 的 Ping。 | [optional] 

## Methods

### NewGetNetworkPingmyip200Response

`func NewGetNetworkPingmyip200Response() *GetNetworkPingmyip200Response`

NewGetNetworkPingmyip200Response instantiates a new GetNetworkPingmyip200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkPingmyip200ResponseWithDefaults

`func NewGetNetworkPingmyip200ResponseWithDefaults() *GetNetworkPingmyip200Response`

NewGetNetworkPingmyip200ResponseWithDefaults instantiates a new GetNetworkPingmyip200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *GetNetworkPingmyip200Response) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *GetNetworkPingmyip200Response) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *GetNetworkPingmyip200Response) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *GetNetworkPingmyip200Response) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetMessage

`func (o *GetNetworkPingmyip200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetNetworkPingmyip200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetNetworkPingmyip200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetNetworkPingmyip200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPingSuccessful

`func (o *GetNetworkPingmyip200Response) GetPingSuccessful() bool`

GetPingSuccessful returns the PingSuccessful field if non-nil, zero value otherwise.

### GetPingSuccessfulOk

`func (o *GetNetworkPingmyip200Response) GetPingSuccessfulOk() (*bool, bool)`

GetPingSuccessfulOk returns a tuple with the PingSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSuccessful

`func (o *GetNetworkPingmyip200Response) SetPingSuccessful(v bool)`

SetPingSuccessful sets PingSuccessful field to given value.

### HasPingSuccessful

`func (o *GetNetworkPingmyip200Response) HasPingSuccessful() bool`

HasPingSuccessful returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetNetworkUrlstatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | 目标不可达或请求失败时固定为 0。 | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewGetNetworkUrlstatus200Response

`func NewGetNetworkUrlstatus200Response() *GetNetworkUrlstatus200Response`

NewGetNetworkUrlstatus200Response instantiates a new GetNetworkUrlstatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkUrlstatus200ResponseWithDefaults

`func NewGetNetworkUrlstatus200ResponseWithDefaults() *GetNetworkUrlstatus200Response`

NewGetNetworkUrlstatus200ResponseWithDefaults instantiates a new GetNetworkUrlstatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetNetworkUrlstatus200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkUrlstatus200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkUrlstatus200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkUrlstatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GetNetworkUrlstatus200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetNetworkUrlstatus200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetNetworkUrlstatus200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetNetworkUrlstatus200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



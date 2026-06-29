# GetGameSteamServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appid** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Servers** | Pointer to [**[]GetGameSteamServers200ResponseServersInner**](GetGameSteamServers200ResponseServersInner.md) |  | [optional] 

## Methods

### NewGetGameSteamServers200Response

`func NewGetGameSteamServers200Response() *GetGameSteamServers200Response`

NewGetGameSteamServers200Response instantiates a new GetGameSteamServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameSteamServers200ResponseWithDefaults

`func NewGetGameSteamServers200ResponseWithDefaults() *GetGameSteamServers200Response`

NewGetGameSteamServers200ResponseWithDefaults instantiates a new GetGameSteamServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppid

`func (o *GetGameSteamServers200Response) GetAppid() int32`

GetAppid returns the Appid field if non-nil, zero value otherwise.

### GetAppidOk

`func (o *GetGameSteamServers200Response) GetAppidOk() (*int32, bool)`

GetAppidOk returns a tuple with the Appid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppid

`func (o *GetGameSteamServers200Response) SetAppid(v int32)`

SetAppid sets Appid field to given value.

### HasAppid

`func (o *GetGameSteamServers200Response) HasAppid() bool`

HasAppid returns a boolean if a field has been set.

### GetCount

`func (o *GetGameSteamServers200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetGameSteamServers200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetGameSteamServers200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetGameSteamServers200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetQuery

`func (o *GetGameSteamServers200Response) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetGameSteamServers200Response) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetGameSteamServers200Response) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GetGameSteamServers200Response) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetServers

`func (o *GetGameSteamServers200Response) GetServers() []GetGameSteamServers200ResponseServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetGameSteamServers200Response) GetServersOk() (*[]GetGameSteamServers200ResponseServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetGameSteamServers200Response) SetServers(v []GetGameSteamServers200ResponseServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetGameSteamServers200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



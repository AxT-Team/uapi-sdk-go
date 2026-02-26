# GetGameMinecraftHistoryid200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | 【name 查询时返回】查询的用户名。 | [optional] 
**Count** | Pointer to **int32** | 【name 查询时返回】匹配到的用户数量，为 0 时表示未找到。 | [optional] 
**Results** | Pointer to [**[]GetGameMinecraftHistoryid200ResponseResultsInner**](GetGameMinecraftHistoryid200ResponseResultsInner.md) | 【name 查询时返回】匹配用户列表，包含当前用户名或曾用名匹配的所有玩家。 | [optional] 
**Id** | Pointer to **string** | 【uuid 查询时返回】玩家当前的用户名。 | [optional] 
**Uuid** | Pointer to **string** | 【uuid 查询时返回】被查询玩家的UUID（带连字符格式）。 | [optional] 
**NameNum** | Pointer to **int32** | 【uuid 查询时返回】历史名称的总数（包含当前名称）。 | [optional] 
**History** | Pointer to [**[]GetGameMinecraftHistoryid200ResponseHistoryInner**](GetGameMinecraftHistoryid200ResponseHistoryInner.md) | 【uuid 查询时返回】包含所有历史用户名的数组，按时间倒序排列。 | [optional] 

## Methods

### NewGetGameMinecraftHistoryid200Response

`func NewGetGameMinecraftHistoryid200Response() *GetGameMinecraftHistoryid200Response`

NewGetGameMinecraftHistoryid200Response instantiates a new GetGameMinecraftHistoryid200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameMinecraftHistoryid200ResponseWithDefaults

`func NewGetGameMinecraftHistoryid200ResponseWithDefaults() *GetGameMinecraftHistoryid200Response`

NewGetGameMinecraftHistoryid200ResponseWithDefaults instantiates a new GetGameMinecraftHistoryid200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *GetGameMinecraftHistoryid200Response) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetGameMinecraftHistoryid200Response) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetGameMinecraftHistoryid200Response) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GetGameMinecraftHistoryid200Response) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCount

`func (o *GetGameMinecraftHistoryid200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetGameMinecraftHistoryid200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetGameMinecraftHistoryid200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetGameMinecraftHistoryid200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *GetGameMinecraftHistoryid200Response) GetResults() []GetGameMinecraftHistoryid200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetGameMinecraftHistoryid200Response) GetResultsOk() (*[]GetGameMinecraftHistoryid200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetGameMinecraftHistoryid200Response) SetResults(v []GetGameMinecraftHistoryid200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetGameMinecraftHistoryid200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetId

`func (o *GetGameMinecraftHistoryid200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGameMinecraftHistoryid200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGameMinecraftHistoryid200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetGameMinecraftHistoryid200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetGameMinecraftHistoryid200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGameMinecraftHistoryid200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGameMinecraftHistoryid200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGameMinecraftHistoryid200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetNameNum

`func (o *GetGameMinecraftHistoryid200Response) GetNameNum() int32`

GetNameNum returns the NameNum field if non-nil, zero value otherwise.

### GetNameNumOk

`func (o *GetGameMinecraftHistoryid200Response) GetNameNumOk() (*int32, bool)`

GetNameNumOk returns a tuple with the NameNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameNum

`func (o *GetGameMinecraftHistoryid200Response) SetNameNum(v int32)`

SetNameNum sets NameNum field to given value.

### HasNameNum

`func (o *GetGameMinecraftHistoryid200Response) HasNameNum() bool`

HasNameNum returns a boolean if a field has been set.

### GetHistory

`func (o *GetGameMinecraftHistoryid200Response) GetHistory() []GetGameMinecraftHistoryid200ResponseHistoryInner`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GetGameMinecraftHistoryid200Response) GetHistoryOk() (*[]GetGameMinecraftHistoryid200ResponseHistoryInner, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GetGameMinecraftHistoryid200Response) SetHistory(v []GetGameMinecraftHistoryid200ResponseHistoryInner)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GetGameMinecraftHistoryid200Response) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



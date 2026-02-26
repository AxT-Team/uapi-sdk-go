# GetMiscHotboard200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetMiscHotboard200ResponseListInner**](GetMiscHotboard200ResponseListInner.md) | 热榜条目列表。 | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**SnapshotTime** | Pointer to **int32** | 时光机模式返回的快照实际时间戳（毫秒）。 | [optional] 
**Keyword** | Pointer to **string** | 搜索模式返回的搜索关键词。 | [optional] 
**Count** | Pointer to **int32** | 搜索模式返回的结果数量。 | [optional] 
**Results** | Pointer to [**[]GetMiscHotboard200ResponseResultsInner**](GetMiscHotboard200ResponseResultsInner.md) | 搜索模式返回的结果数组。 | [optional] 
**Sources** | Pointer to **[]string** | 数据源列表模式返回的可用历史数据源数组。 | [optional] 

## Methods

### NewGetMiscHotboard200Response

`func NewGetMiscHotboard200Response() *GetMiscHotboard200Response`

NewGetMiscHotboard200Response instantiates a new GetMiscHotboard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHotboard200ResponseWithDefaults

`func NewGetMiscHotboard200ResponseWithDefaults() *GetMiscHotboard200Response`

NewGetMiscHotboard200ResponseWithDefaults instantiates a new GetMiscHotboard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetMiscHotboard200Response) GetList() []GetMiscHotboard200ResponseListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetMiscHotboard200Response) GetListOk() (*[]GetMiscHotboard200ResponseListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetMiscHotboard200Response) SetList(v []GetMiscHotboard200ResponseListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetMiscHotboard200Response) HasList() bool`

HasList returns a boolean if a field has been set.

### GetType

`func (o *GetMiscHotboard200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMiscHotboard200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMiscHotboard200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMiscHotboard200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetMiscHotboard200Response) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetMiscHotboard200Response) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetMiscHotboard200Response) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetMiscHotboard200Response) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetSnapshotTime

`func (o *GetMiscHotboard200Response) GetSnapshotTime() int32`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *GetMiscHotboard200Response) GetSnapshotTimeOk() (*int32, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *GetMiscHotboard200Response) SetSnapshotTime(v int32)`

SetSnapshotTime sets SnapshotTime field to given value.

### HasSnapshotTime

`func (o *GetMiscHotboard200Response) HasSnapshotTime() bool`

HasSnapshotTime returns a boolean if a field has been set.

### GetKeyword

`func (o *GetMiscHotboard200Response) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *GetMiscHotboard200Response) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *GetMiscHotboard200Response) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *GetMiscHotboard200Response) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetCount

`func (o *GetMiscHotboard200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetMiscHotboard200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetMiscHotboard200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetMiscHotboard200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *GetMiscHotboard200Response) GetResults() []GetMiscHotboard200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetMiscHotboard200Response) GetResultsOk() (*[]GetMiscHotboard200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetMiscHotboard200Response) SetResults(v []GetMiscHotboard200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetMiscHotboard200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSources

`func (o *GetMiscHotboard200Response) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetMiscHotboard200Response) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetMiscHotboard200Response) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetMiscHotboard200Response) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



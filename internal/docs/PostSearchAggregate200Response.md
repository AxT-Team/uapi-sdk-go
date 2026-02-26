# PostSearchAggregate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | 实际执行的搜索查询 | [optional] 
**TotalResults** | Pointer to **int32** | 搜索结果总数 | [optional] 
**Results** | Pointer to [**[]PostSearchAggregate200ResponseResultsInner**](PostSearchAggregate200ResponseResultsInner.md) | 搜索结果列表 | [optional] 
**Sources** | Pointer to [**[]PostSearchAggregate200ResponseSourcesInner**](PostSearchAggregate200ResponseSourcesInner.md) | 各搜索源的结果统计 | [optional] 
**ProcessTimeMs** | Pointer to **int32** | 处理耗时（毫秒） | [optional] 
**Cached** | Pointer to **bool** | 结果是否来自缓存 | [optional] 

## Methods

### NewPostSearchAggregate200Response

`func NewPostSearchAggregate200Response() *PostSearchAggregate200Response`

NewPostSearchAggregate200Response instantiates a new PostSearchAggregate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSearchAggregate200ResponseWithDefaults

`func NewPostSearchAggregate200ResponseWithDefaults() *PostSearchAggregate200Response`

NewPostSearchAggregate200ResponseWithDefaults instantiates a new PostSearchAggregate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *PostSearchAggregate200Response) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PostSearchAggregate200Response) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PostSearchAggregate200Response) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PostSearchAggregate200Response) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTotalResults

`func (o *PostSearchAggregate200Response) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *PostSearchAggregate200Response) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *PostSearchAggregate200Response) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *PostSearchAggregate200Response) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResults

`func (o *PostSearchAggregate200Response) GetResults() []PostSearchAggregate200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PostSearchAggregate200Response) GetResultsOk() (*[]PostSearchAggregate200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PostSearchAggregate200Response) SetResults(v []PostSearchAggregate200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *PostSearchAggregate200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSources

`func (o *PostSearchAggregate200Response) GetSources() []PostSearchAggregate200ResponseSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PostSearchAggregate200Response) GetSourcesOk() (*[]PostSearchAggregate200ResponseSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PostSearchAggregate200Response) SetSources(v []PostSearchAggregate200ResponseSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PostSearchAggregate200Response) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetProcessTimeMs

`func (o *PostSearchAggregate200Response) GetProcessTimeMs() int32`

GetProcessTimeMs returns the ProcessTimeMs field if non-nil, zero value otherwise.

### GetProcessTimeMsOk

`func (o *PostSearchAggregate200Response) GetProcessTimeMsOk() (*int32, bool)`

GetProcessTimeMsOk returns a tuple with the ProcessTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTimeMs

`func (o *PostSearchAggregate200Response) SetProcessTimeMs(v int32)`

SetProcessTimeMs sets ProcessTimeMs field to given value.

### HasProcessTimeMs

`func (o *PostSearchAggregate200Response) HasProcessTimeMs() bool`

HasProcessTimeMs returns a boolean if a field has been set.

### GetCached

`func (o *PostSearchAggregate200Response) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *PostSearchAggregate200Response) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *PostSearchAggregate200Response) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *PostSearchAggregate200Response) HasCached() bool`

HasCached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



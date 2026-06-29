# PostSearchAggregate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PostSearchAggregate200ResponseMetadata**](PostSearchAggregate200ResponseMetadata.md) |  | [optional] 
**ProcessTimeMs** | Pointer to **int32** | 本次请求总耗时（毫秒） | [optional] 
**Query** | Pointer to **string** | 执行的搜索查询 | [optional] 
**Results** | Pointer to [**[]PostSearchAggregate200ResponseResultsInner**](PostSearchAggregate200ResponseResultsInner.md) | 搜索结果列表 | [optional] 
**Sources** | Pointer to [**[]PostSearchAggregate200ResponseSourcesInner**](PostSearchAggregate200ResponseSourcesInner.md) | 本次请求实际命中的搜索引擎信息 | [optional] 
**TotalResults** | Pointer to **int32** | 返回的搜索结果总数 | [optional] 

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

### GetMetadata

`func (o *PostSearchAggregate200Response) GetMetadata() PostSearchAggregate200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostSearchAggregate200Response) GetMetadataOk() (*PostSearchAggregate200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostSearchAggregate200Response) SetMetadata(v PostSearchAggregate200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostSearchAggregate200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



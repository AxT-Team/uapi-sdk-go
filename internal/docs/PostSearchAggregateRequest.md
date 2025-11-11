# PostSearchAggregateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | 搜索查询关键词，支持中英文 | 
**Site** | Pointer to **string** | 限制搜索特定网站，不需要 &#x60;site:&#x60; 前缀 | [optional] 
**Filetype** | Pointer to **string** | 限制文件类型，不需要 &#x60;filetype:&#x60; 前缀。支持 pdf、doc、docx、ppt、pptx、xls、xlsx、txt 等 | [optional] 
**FetchFull** | Pointer to **bool** | 是否获取页面完整正文（会影响响应时间） | [optional] [default to false]
**TimeoutMs** | Pointer to **int32** | 请求超时时间（毫秒），范围 1000-30000 | [optional] [default to 3000]

## Methods

### NewPostSearchAggregateRequest

`func NewPostSearchAggregateRequest(query string, ) *PostSearchAggregateRequest`

NewPostSearchAggregateRequest instantiates a new PostSearchAggregateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSearchAggregateRequestWithDefaults

`func NewPostSearchAggregateRequestWithDefaults() *PostSearchAggregateRequest`

NewPostSearchAggregateRequestWithDefaults instantiates a new PostSearchAggregateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *PostSearchAggregateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PostSearchAggregateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PostSearchAggregateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetSite

`func (o *PostSearchAggregateRequest) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PostSearchAggregateRequest) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PostSearchAggregateRequest) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PostSearchAggregateRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetFiletype

`func (o *PostSearchAggregateRequest) GetFiletype() string`

GetFiletype returns the Filetype field if non-nil, zero value otherwise.

### GetFiletypeOk

`func (o *PostSearchAggregateRequest) GetFiletypeOk() (*string, bool)`

GetFiletypeOk returns a tuple with the Filetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiletype

`func (o *PostSearchAggregateRequest) SetFiletype(v string)`

SetFiletype sets Filetype field to given value.

### HasFiletype

`func (o *PostSearchAggregateRequest) HasFiletype() bool`

HasFiletype returns a boolean if a field has been set.

### GetFetchFull

`func (o *PostSearchAggregateRequest) GetFetchFull() bool`

GetFetchFull returns the FetchFull field if non-nil, zero value otherwise.

### GetFetchFullOk

`func (o *PostSearchAggregateRequest) GetFetchFullOk() (*bool, bool)`

GetFetchFullOk returns a tuple with the FetchFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchFull

`func (o *PostSearchAggregateRequest) SetFetchFull(v bool)`

SetFetchFull sets FetchFull field to given value.

### HasFetchFull

`func (o *PostSearchAggregateRequest) HasFetchFull() bool`

HasFetchFull returns a boolean if a field has been set.

### GetTimeoutMs

`func (o *PostSearchAggregateRequest) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *PostSearchAggregateRequest) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *PostSearchAggregateRequest) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *PostSearchAggregateRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



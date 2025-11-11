# PostSearchAggregate200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | 结果标题 | [optional] 
**Url** | Pointer to **string** | 结果链接 | [optional] 
**Snippet** | Pointer to **string** | 结果摘要/描述 | [optional] 
**Domain** | Pointer to **string** | 来源域名 | [optional] 
**Source** | Pointer to **string** | 搜索引擎标识 | [optional] 
**Position** | Pointer to **int32** | 原始排名位置 | [optional] 
**Score** | Pointer to **float32** | 综合得分 (0-1，经过机器学习排序) | [optional] 
**PublishTime** | Pointer to **time.Time** | 发布时间 (ISO 8601 格式) | [optional] 
**Author** | Pointer to **NullableString** | 作者信息 | [optional] 

## Methods

### NewPostSearchAggregate200ResponseResultsInner

`func NewPostSearchAggregate200ResponseResultsInner() *PostSearchAggregate200ResponseResultsInner`

NewPostSearchAggregate200ResponseResultsInner instantiates a new PostSearchAggregate200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSearchAggregate200ResponseResultsInnerWithDefaults

`func NewPostSearchAggregate200ResponseResultsInnerWithDefaults() *PostSearchAggregate200ResponseResultsInner`

NewPostSearchAggregate200ResponseResultsInnerWithDefaults instantiates a new PostSearchAggregate200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PostSearchAggregate200ResponseResultsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostSearchAggregate200ResponseResultsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostSearchAggregate200ResponseResultsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *PostSearchAggregate200ResponseResultsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostSearchAggregate200ResponseResultsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostSearchAggregate200ResponseResultsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSnippet

`func (o *PostSearchAggregate200ResponseResultsInner) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *PostSearchAggregate200ResponseResultsInner) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *PostSearchAggregate200ResponseResultsInner) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetDomain

`func (o *PostSearchAggregate200ResponseResultsInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PostSearchAggregate200ResponseResultsInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PostSearchAggregate200ResponseResultsInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSource

`func (o *PostSearchAggregate200ResponseResultsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostSearchAggregate200ResponseResultsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PostSearchAggregate200ResponseResultsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPosition

`func (o *PostSearchAggregate200ResponseResultsInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PostSearchAggregate200ResponseResultsInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PostSearchAggregate200ResponseResultsInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetScore

`func (o *PostSearchAggregate200ResponseResultsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PostSearchAggregate200ResponseResultsInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *PostSearchAggregate200ResponseResultsInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPublishTime

`func (o *PostSearchAggregate200ResponseResultsInner) GetPublishTime() time.Time`

GetPublishTime returns the PublishTime field if non-nil, zero value otherwise.

### GetPublishTimeOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetPublishTimeOk() (*time.Time, bool)`

GetPublishTimeOk returns a tuple with the PublishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTime

`func (o *PostSearchAggregate200ResponseResultsInner) SetPublishTime(v time.Time)`

SetPublishTime sets PublishTime field to given value.

### HasPublishTime

`func (o *PostSearchAggregate200ResponseResultsInner) HasPublishTime() bool`

HasPublishTime returns a boolean if a field has been set.

### GetAuthor

`func (o *PostSearchAggregate200ResponseResultsInner) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PostSearchAggregate200ResponseResultsInner) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PostSearchAggregate200ResponseResultsInner) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PostSearchAggregate200ResponseResultsInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *PostSearchAggregate200ResponseResultsInner) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *PostSearchAggregate200ResponseResultsInner) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



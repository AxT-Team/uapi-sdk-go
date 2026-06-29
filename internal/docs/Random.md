# Random

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | 作者或说话者。 | [optional] 
**Authorinfo** | Pointer to [**RandomAuthorinfo**](RandomAuthorinfo.md) |  | [optional] 
**Category** | Pointer to **string** | 语录分类。 | [optional] 
**Content** | Pointer to **string** | 语录正文。 | [optional] 
**ContentLength** | Pointer to **int32** | 正文字数。 | [optional] 
**Corpus** | Pointer to **string** | 所属语料库标识。 | [optional] 
**CreatedAt** | Pointer to **string** | 语料入库时间戳（秒），部分语料返回。 | [optional] 
**MatchedTags** | Pointer to **[]string** | 命中的标签，部分模式返回。 | [optional] 
**Source** | Pointer to **string** | 语录出处或来源。 | [optional] 
**Uuid** | Pointer to **string** | 语录唯一标识。 | [optional] 

## Methods

### NewRandom

`func NewRandom() *Random`

NewRandom instantiates a new Random object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomWithDefaults

`func NewRandomWithDefaults() *Random`

NewRandomWithDefaults instantiates a new Random object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Random) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Random) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Random) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Random) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorinfo

`func (o *Random) GetAuthorinfo() RandomAuthorinfo`

GetAuthorinfo returns the Authorinfo field if non-nil, zero value otherwise.

### GetAuthorinfoOk

`func (o *Random) GetAuthorinfoOk() (*RandomAuthorinfo, bool)`

GetAuthorinfoOk returns a tuple with the Authorinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorinfo

`func (o *Random) SetAuthorinfo(v RandomAuthorinfo)`

SetAuthorinfo sets Authorinfo field to given value.

### HasAuthorinfo

`func (o *Random) HasAuthorinfo() bool`

HasAuthorinfo returns a boolean if a field has been set.

### GetCategory

`func (o *Random) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Random) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Random) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Random) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContent

`func (o *Random) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Random) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Random) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Random) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentLength

`func (o *Random) GetContentLength() int32`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *Random) GetContentLengthOk() (*int32, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *Random) SetContentLength(v int32)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *Random) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### GetCorpus

`func (o *Random) GetCorpus() string`

GetCorpus returns the Corpus field if non-nil, zero value otherwise.

### GetCorpusOk

`func (o *Random) GetCorpusOk() (*string, bool)`

GetCorpusOk returns a tuple with the Corpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorpus

`func (o *Random) SetCorpus(v string)`

SetCorpus sets Corpus field to given value.

### HasCorpus

`func (o *Random) HasCorpus() bool`

HasCorpus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Random) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Random) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Random) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Random) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMatchedTags

`func (o *Random) GetMatchedTags() []string`

GetMatchedTags returns the MatchedTags field if non-nil, zero value otherwise.

### GetMatchedTagsOk

`func (o *Random) GetMatchedTagsOk() (*[]string, bool)`

GetMatchedTagsOk returns a tuple with the MatchedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedTags

`func (o *Random) SetMatchedTags(v []string)`

SetMatchedTags sets MatchedTags field to given value.

### HasMatchedTags

`func (o *Random) HasMatchedTags() bool`

HasMatchedTags returns a boolean if a field has been set.

### GetSource

`func (o *Random) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Random) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Random) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Random) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUuid

`func (o *Random) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Random) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Random) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Random) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



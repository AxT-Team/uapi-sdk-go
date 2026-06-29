# GetSayingRandom200Response

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
**CurrentTime** | Pointer to **string** | 仅 moment 模式返回，服务器当前时间，ISO 8601 格式。 | [optional] 
**Date** | Pointer to **string** | 仅 daily 模式返回，对应日期，格式 YYYY-MM-DD。 | [optional] 
**Item** | Pointer to [**DailyRecommendMomentItem**](DailyRecommendMomentItem.md) |  | [optional] 
**Mode** | Pointer to **string** | 当前运行模式。 | [optional] 
**Scene** | Pointer to [**DailyRecommendMomentScene**](DailyRecommendMomentScene.md) |  | [optional] 
**Seed** | Pointer to **string** | 当次结果的确定性种子。 | [optional] 
**TimeSegment** | Pointer to **string** | 仅 moment 模式返回，命中的时段标识。 | [optional] 

## Methods

### NewGetSayingRandom200Response

`func NewGetSayingRandom200Response() *GetSayingRandom200Response`

NewGetSayingRandom200Response instantiates a new GetSayingRandom200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSayingRandom200ResponseWithDefaults

`func NewGetSayingRandom200ResponseWithDefaults() *GetSayingRandom200Response`

NewGetSayingRandom200ResponseWithDefaults instantiates a new GetSayingRandom200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *GetSayingRandom200Response) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetSayingRandom200Response) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetSayingRandom200Response) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetSayingRandom200Response) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorinfo

`func (o *GetSayingRandom200Response) GetAuthorinfo() RandomAuthorinfo`

GetAuthorinfo returns the Authorinfo field if non-nil, zero value otherwise.

### GetAuthorinfoOk

`func (o *GetSayingRandom200Response) GetAuthorinfoOk() (*RandomAuthorinfo, bool)`

GetAuthorinfoOk returns a tuple with the Authorinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorinfo

`func (o *GetSayingRandom200Response) SetAuthorinfo(v RandomAuthorinfo)`

SetAuthorinfo sets Authorinfo field to given value.

### HasAuthorinfo

`func (o *GetSayingRandom200Response) HasAuthorinfo() bool`

HasAuthorinfo returns a boolean if a field has been set.

### GetCategory

`func (o *GetSayingRandom200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetSayingRandom200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetSayingRandom200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetSayingRandom200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContent

`func (o *GetSayingRandom200Response) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetSayingRandom200Response) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetSayingRandom200Response) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetSayingRandom200Response) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentLength

`func (o *GetSayingRandom200Response) GetContentLength() int32`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *GetSayingRandom200Response) GetContentLengthOk() (*int32, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *GetSayingRandom200Response) SetContentLength(v int32)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *GetSayingRandom200Response) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### GetCorpus

`func (o *GetSayingRandom200Response) GetCorpus() string`

GetCorpus returns the Corpus field if non-nil, zero value otherwise.

### GetCorpusOk

`func (o *GetSayingRandom200Response) GetCorpusOk() (*string, bool)`

GetCorpusOk returns a tuple with the Corpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorpus

`func (o *GetSayingRandom200Response) SetCorpus(v string)`

SetCorpus sets Corpus field to given value.

### HasCorpus

`func (o *GetSayingRandom200Response) HasCorpus() bool`

HasCorpus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetSayingRandom200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetSayingRandom200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetSayingRandom200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetSayingRandom200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMatchedTags

`func (o *GetSayingRandom200Response) GetMatchedTags() []string`

GetMatchedTags returns the MatchedTags field if non-nil, zero value otherwise.

### GetMatchedTagsOk

`func (o *GetSayingRandom200Response) GetMatchedTagsOk() (*[]string, bool)`

GetMatchedTagsOk returns a tuple with the MatchedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedTags

`func (o *GetSayingRandom200Response) SetMatchedTags(v []string)`

SetMatchedTags sets MatchedTags field to given value.

### HasMatchedTags

`func (o *GetSayingRandom200Response) HasMatchedTags() bool`

HasMatchedTags returns a boolean if a field has been set.

### GetSource

`func (o *GetSayingRandom200Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetSayingRandom200Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetSayingRandom200Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetSayingRandom200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUuid

`func (o *GetSayingRandom200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetSayingRandom200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetSayingRandom200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetSayingRandom200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCurrentTime

`func (o *GetSayingRandom200Response) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *GetSayingRandom200Response) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *GetSayingRandom200Response) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *GetSayingRandom200Response) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetDate

`func (o *GetSayingRandom200Response) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetSayingRandom200Response) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetSayingRandom200Response) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetSayingRandom200Response) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetItem

`func (o *GetSayingRandom200Response) GetItem() DailyRecommendMomentItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GetSayingRandom200Response) GetItemOk() (*DailyRecommendMomentItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GetSayingRandom200Response) SetItem(v DailyRecommendMomentItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *GetSayingRandom200Response) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetMode

`func (o *GetSayingRandom200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetSayingRandom200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetSayingRandom200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetSayingRandom200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetScene

`func (o *GetSayingRandom200Response) GetScene() DailyRecommendMomentScene`

GetScene returns the Scene field if non-nil, zero value otherwise.

### GetSceneOk

`func (o *GetSayingRandom200Response) GetSceneOk() (*DailyRecommendMomentScene, bool)`

GetSceneOk returns a tuple with the Scene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScene

`func (o *GetSayingRandom200Response) SetScene(v DailyRecommendMomentScene)`

SetScene sets Scene field to given value.

### HasScene

`func (o *GetSayingRandom200Response) HasScene() bool`

HasScene returns a boolean if a field has been set.

### GetSeed

`func (o *GetSayingRandom200Response) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *GetSayingRandom200Response) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *GetSayingRandom200Response) SetSeed(v string)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *GetSayingRandom200Response) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetTimeSegment

`func (o *GetSayingRandom200Response) GetTimeSegment() string`

GetTimeSegment returns the TimeSegment field if non-nil, zero value otherwise.

### GetTimeSegmentOk

`func (o *GetSayingRandom200Response) GetTimeSegmentOk() (*string, bool)`

GetTimeSegmentOk returns a tuple with the TimeSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSegment

`func (o *GetSayingRandom200Response) SetTimeSegment(v string)`

SetTimeSegment sets TimeSegment field to given value.

### HasTimeSegment

`func (o *GetSayingRandom200Response) HasTimeSegment() bool`

HasTimeSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



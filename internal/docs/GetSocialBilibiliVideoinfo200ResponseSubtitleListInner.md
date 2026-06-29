# GetSocialBilibiliVideoinfo200ResponseSubtitleListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseSubtitleListInnerAuthor**](GetSocialBilibiliVideoinfo200ResponseSubtitleListInnerAuthor.md) |  | [optional] 
**AuthorMid** | Pointer to **float32** | 字幕作者 UID。 | [optional] 
**Id** | Pointer to **float32** | 字幕 ID。 | [optional] 
**IsLock** | Pointer to **bool** | 是否锁定。 | [optional] 
**Lan** | Pointer to **string** | 语言代码。 | [optional] 
**LanDoc** | Pointer to **string** | 语言名称。 | [optional] 
**SubtitleUrl** | Pointer to **string** | 字幕文件链接。 | [optional] 

## Methods

### NewGetSocialBilibiliVideoinfo200ResponseSubtitleListInner

`func NewGetSocialBilibiliVideoinfo200ResponseSubtitleListInner() *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner`

NewGetSocialBilibiliVideoinfo200ResponseSubtitleListInner instantiates a new GetSocialBilibiliVideoinfo200ResponseSubtitleListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialBilibiliVideoinfo200ResponseSubtitleListInnerWithDefaults

`func NewGetSocialBilibiliVideoinfo200ResponseSubtitleListInnerWithDefaults() *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner`

NewGetSocialBilibiliVideoinfo200ResponseSubtitleListInnerWithDefaults instantiates a new GetSocialBilibiliVideoinfo200ResponseSubtitleListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetAuthor() GetSocialBilibiliVideoinfo200ResponseSubtitleListInnerAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetAuthorOk() (*GetSocialBilibiliVideoinfo200ResponseSubtitleListInnerAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) SetAuthor(v GetSocialBilibiliVideoinfo200ResponseSubtitleListInnerAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorMid

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetAuthorMid() float32`

GetAuthorMid returns the AuthorMid field if non-nil, zero value otherwise.

### GetAuthorMidOk

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetAuthorMidOk() (*float32, bool)`

GetAuthorMidOk returns a tuple with the AuthorMid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMid

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) SetAuthorMid(v float32)`

SetAuthorMid sets AuthorMid field to given value.

### HasAuthorMid

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) HasAuthorMid() bool`

HasAuthorMid returns a boolean if a field has been set.

### GetId

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsLock

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetIsLock() bool`

GetIsLock returns the IsLock field if non-nil, zero value otherwise.

### GetIsLockOk

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetIsLockOk() (*bool, bool)`

GetIsLockOk returns a tuple with the IsLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLock

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) SetIsLock(v bool)`

SetIsLock sets IsLock field to given value.

### HasIsLock

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) HasIsLock() bool`

HasIsLock returns a boolean if a field has been set.

### GetLan

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLanDoc

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetLanDoc() string`

GetLanDoc returns the LanDoc field if non-nil, zero value otherwise.

### GetLanDocOk

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetLanDocOk() (*string, bool)`

GetLanDocOk returns a tuple with the LanDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanDoc

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) SetLanDoc(v string)`

SetLanDoc sets LanDoc field to given value.

### HasLanDoc

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) HasLanDoc() bool`

HasLanDoc returns a boolean if a field has been set.

### GetSubtitleUrl

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetSubtitleUrl() string`

GetSubtitleUrl returns the SubtitleUrl field if non-nil, zero value otherwise.

### GetSubtitleUrlOk

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) GetSubtitleUrlOk() (*string, bool)`

GetSubtitleUrlOk returns a tuple with the SubtitleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleUrl

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) SetSubtitleUrl(v string)`

SetSubtitleUrl sets SubtitleUrl field to given value.

### HasSubtitleUrl

`func (o *GetSocialBilibiliVideoinfo200ResponseSubtitleListInner) HasSubtitleUrl() bool`

HasSubtitleUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



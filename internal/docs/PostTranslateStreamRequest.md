# PostTranslateStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | 待翻译的文本内容 | 
**ToLang** | **string** | 目标语言，支持：中文、英文 | 
**FromLang** | Pointer to **string** | 源语言，支持：中文、英文、auto（自动检测）。默认为auto | [optional] [default to "auto"]
**Tone** | Pointer to **string** | 语气参数，可选 | [optional] 

## Methods

### NewPostTranslateStreamRequest

`func NewPostTranslateStreamRequest(query string, toLang string, ) *PostTranslateStreamRequest`

NewPostTranslateStreamRequest instantiates a new PostTranslateStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTranslateStreamRequestWithDefaults

`func NewPostTranslateStreamRequestWithDefaults() *PostTranslateStreamRequest`

NewPostTranslateStreamRequestWithDefaults instantiates a new PostTranslateStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *PostTranslateStreamRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PostTranslateStreamRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PostTranslateStreamRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetToLang

`func (o *PostTranslateStreamRequest) GetToLang() string`

GetToLang returns the ToLang field if non-nil, zero value otherwise.

### GetToLangOk

`func (o *PostTranslateStreamRequest) GetToLangOk() (*string, bool)`

GetToLangOk returns a tuple with the ToLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToLang

`func (o *PostTranslateStreamRequest) SetToLang(v string)`

SetToLang sets ToLang field to given value.


### GetFromLang

`func (o *PostTranslateStreamRequest) GetFromLang() string`

GetFromLang returns the FromLang field if non-nil, zero value otherwise.

### GetFromLangOk

`func (o *PostTranslateStreamRequest) GetFromLangOk() (*string, bool)`

GetFromLangOk returns a tuple with the FromLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLang

`func (o *PostTranslateStreamRequest) SetFromLang(v string)`

SetFromLang sets FromLang field to given value.

### HasFromLang

`func (o *PostTranslateStreamRequest) HasFromLang() bool`

HasFromLang returns a boolean if a field has been set.

### GetTone

`func (o *PostTranslateStreamRequest) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *PostTranslateStreamRequest) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *PostTranslateStreamRequest) SetTone(v string)`

SetTone sets Tone field to given value.

### HasTone

`func (o *PostTranslateStreamRequest) HasTone() bool`

HasTone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



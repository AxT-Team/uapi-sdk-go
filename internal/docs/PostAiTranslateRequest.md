# PostAiTranslateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | 待翻译的文本内容。最大长度10,000字符。 | 
**SourceLang** | Pointer to **string** | 源语言代码，可选。如果不指定，系统会自动检测源语言。 | [optional] 
**Style** | Pointer to **string** | 翻译风格，可选。支持casual(随意口语化)、professional(专业商务，默认)、academic(学术正式)、literary(文学艺术)。 | [optional] [default to "professional"]
**Context** | Pointer to **string** | 翻译上下文场景，可选。支持general(通用，默认)、business(商务)、technical(技术)、medical(医疗)、legal(法律)、marketing(市场营销)、entertainment(娱乐)、education(教育)、news(新闻)。 | [optional] [default to "general"]
**PreserveFormat** | Pointer to **bool** | 是否保留原文格式，包括换行、缩进等。 | [optional] [default to true]

## Methods

### NewPostAiTranslateRequest

`func NewPostAiTranslateRequest(text string, ) *PostAiTranslateRequest`

NewPostAiTranslateRequest instantiates a new PostAiTranslateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAiTranslateRequestWithDefaults

`func NewPostAiTranslateRequestWithDefaults() *PostAiTranslateRequest`

NewPostAiTranslateRequestWithDefaults instantiates a new PostAiTranslateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PostAiTranslateRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostAiTranslateRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostAiTranslateRequest) SetText(v string)`

SetText sets Text field to given value.


### GetSourceLang

`func (o *PostAiTranslateRequest) GetSourceLang() string`

GetSourceLang returns the SourceLang field if non-nil, zero value otherwise.

### GetSourceLangOk

`func (o *PostAiTranslateRequest) GetSourceLangOk() (*string, bool)`

GetSourceLangOk returns a tuple with the SourceLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLang

`func (o *PostAiTranslateRequest) SetSourceLang(v string)`

SetSourceLang sets SourceLang field to given value.

### HasSourceLang

`func (o *PostAiTranslateRequest) HasSourceLang() bool`

HasSourceLang returns a boolean if a field has been set.

### GetStyle

`func (o *PostAiTranslateRequest) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *PostAiTranslateRequest) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *PostAiTranslateRequest) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *PostAiTranslateRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetContext

`func (o *PostAiTranslateRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PostAiTranslateRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PostAiTranslateRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PostAiTranslateRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetPreserveFormat

`func (o *PostAiTranslateRequest) GetPreserveFormat() bool`

GetPreserveFormat returns the PreserveFormat field if non-nil, zero value otherwise.

### GetPreserveFormatOk

`func (o *PostAiTranslateRequest) GetPreserveFormatOk() (*bool, bool)`

GetPreserveFormatOk returns a tuple with the PreserveFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFormat

`func (o *PostAiTranslateRequest) SetPreserveFormat(v bool)`

SetPreserveFormat sets PreserveFormat field to given value.

### HasPreserveFormat

`func (o *PostAiTranslateRequest) HasPreserveFormat() bool`

HasPreserveFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



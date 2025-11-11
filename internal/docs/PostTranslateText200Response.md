# PostTranslateText200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLang** | Pointer to **string** | The source language detected. | [optional] 
**TranslatedText** | Pointer to **string** | The translated text. | [optional] 

## Methods

### NewPostTranslateText200Response

`func NewPostTranslateText200Response() *PostTranslateText200Response`

NewPostTranslateText200Response instantiates a new PostTranslateText200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTranslateText200ResponseWithDefaults

`func NewPostTranslateText200ResponseWithDefaults() *PostTranslateText200Response`

NewPostTranslateText200ResponseWithDefaults instantiates a new PostTranslateText200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLang

`func (o *PostTranslateText200Response) GetSourceLang() string`

GetSourceLang returns the SourceLang field if non-nil, zero value otherwise.

### GetSourceLangOk

`func (o *PostTranslateText200Response) GetSourceLangOk() (*string, bool)`

GetSourceLangOk returns a tuple with the SourceLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLang

`func (o *PostTranslateText200Response) SetSourceLang(v string)`

SetSourceLang sets SourceLang field to given value.

### HasSourceLang

`func (o *PostTranslateText200Response) HasSourceLang() bool`

HasSourceLang returns a boolean if a field has been set.

### GetTranslatedText

`func (o *PostTranslateText200Response) GetTranslatedText() string`

GetTranslatedText returns the TranslatedText field if non-nil, zero value otherwise.

### GetTranslatedTextOk

`func (o *PostTranslateText200Response) GetTranslatedTextOk() (*string, bool)`

GetTranslatedTextOk returns a tuple with the TranslatedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedText

`func (o *PostTranslateText200Response) SetTranslatedText(v string)`

SetTranslatedText sets TranslatedText field to given value.

### HasTranslatedText

`func (o *PostTranslateText200Response) HasTranslatedText() bool`

HasTranslatedText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



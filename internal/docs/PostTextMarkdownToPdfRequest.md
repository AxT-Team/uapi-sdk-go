# PostTextMarkdownToPdfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | 原始 Markdown 字符串，最大不超过 1MB。 | 
**Theme** | Pointer to **string** | PDF 的排版主题。可选 &#x60;github&#x60;、&#x60;minimal&#x60;、&#x60;light&#x60;、&#x60;dark&#x60;，默认是 &#x60;github&#x60;。 | [optional] [default to "github"]
**PaperSize** | Pointer to **string** | PDF 的纸张大小。可选 &#x60;A4&#x60; 或 &#x60;Letter&#x60;，默认是 &#x60;A4&#x60;。 | [optional] [default to "A4"]

## Methods

### NewPostTextMarkdownToPdfRequest

`func NewPostTextMarkdownToPdfRequest(text string, ) *PostTextMarkdownToPdfRequest`

NewPostTextMarkdownToPdfRequest instantiates a new PostTextMarkdownToPdfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextMarkdownToPdfRequestWithDefaults

`func NewPostTextMarkdownToPdfRequestWithDefaults() *PostTextMarkdownToPdfRequest`

NewPostTextMarkdownToPdfRequestWithDefaults instantiates a new PostTextMarkdownToPdfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PostTextMarkdownToPdfRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTextMarkdownToPdfRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTextMarkdownToPdfRequest) SetText(v string)`

SetText sets Text field to given value.


### GetTheme

`func (o *PostTextMarkdownToPdfRequest) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *PostTextMarkdownToPdfRequest) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *PostTextMarkdownToPdfRequest) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *PostTextMarkdownToPdfRequest) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetPaperSize

`func (o *PostTextMarkdownToPdfRequest) GetPaperSize() string`

GetPaperSize returns the PaperSize field if non-nil, zero value otherwise.

### GetPaperSizeOk

`func (o *PostTextMarkdownToPdfRequest) GetPaperSizeOk() (*string, bool)`

GetPaperSizeOk returns a tuple with the PaperSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperSize

`func (o *PostTextMarkdownToPdfRequest) SetPaperSize(v string)`

SetPaperSize sets PaperSize field to given value.

### HasPaperSize

`func (o *PostTextMarkdownToPdfRequest) HasPaperSize() bool`

HasPaperSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostTextMarkdownToHtmlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | 原始 Markdown 字符串，最大不超过 1MB。 | 
**Format** | Pointer to **string** | 响应格式。传 &#x60;json&#x60; 时返回 JSON 包裹的 HTML 片段；传 &#x60;html&#x60; 时直接返回 &#x60;text/html&#x60;，并且响应内容会自动带完整的网页结构，适合浏览器预览或直接保存为网页文件。默认是 &#x60;json&#x60;。 | [optional] [default to "json"]
**Sanitize** | Pointer to **bool** | 是否开启安全模式，过滤掉用户输入中的风险脚本。默认是 &#x60;true&#x60;。 | [optional] [default to true]

## Methods

### NewPostTextMarkdownToHtmlRequest

`func NewPostTextMarkdownToHtmlRequest(text string, ) *PostTextMarkdownToHtmlRequest`

NewPostTextMarkdownToHtmlRequest instantiates a new PostTextMarkdownToHtmlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextMarkdownToHtmlRequestWithDefaults

`func NewPostTextMarkdownToHtmlRequestWithDefaults() *PostTextMarkdownToHtmlRequest`

NewPostTextMarkdownToHtmlRequestWithDefaults instantiates a new PostTextMarkdownToHtmlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PostTextMarkdownToHtmlRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTextMarkdownToHtmlRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTextMarkdownToHtmlRequest) SetText(v string)`

SetText sets Text field to given value.


### GetFormat

`func (o *PostTextMarkdownToHtmlRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PostTextMarkdownToHtmlRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PostTextMarkdownToHtmlRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PostTextMarkdownToHtmlRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSanitize

`func (o *PostTextMarkdownToHtmlRequest) GetSanitize() bool`

GetSanitize returns the Sanitize field if non-nil, zero value otherwise.

### GetSanitizeOk

`func (o *PostTextMarkdownToHtmlRequest) GetSanitizeOk() (*bool, bool)`

GetSanitizeOk returns a tuple with the Sanitize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitize

`func (o *PostTextMarkdownToHtmlRequest) SetSanitize(v bool)`

SetSanitize sets Sanitize field to given value.

### HasSanitize

`func (o *PostTextMarkdownToHtmlRequest) HasSanitize() bool`

HasSanitize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



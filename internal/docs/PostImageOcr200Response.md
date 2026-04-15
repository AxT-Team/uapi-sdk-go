# PostImageOcr200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | 按阅读顺序拼接后的识别文本。 | [optional] 
**PlainText** | Pointer to **string** | 纯文本结果，适合做搜索、索引或直接展示。 | [optional] 
**Markdown** | Pointer to **string** | 根据图片中的标题、段落和表格整理出的 Markdown 文本。只有在 &#x60;return_markdown&#x3D;true&#x60; 时才会返回。 | [optional] 
**WordsResult** | Pointer to [**[]PostImageOcr200ResponseWordsResultInner**](PostImageOcr200ResponseWordsResultInner.md) | 逐段文字结果。适合做高亮、框选和逐项解析。 | [optional] 
**WordsResultNum** | Pointer to **int32** | 识别出的文字片段数量。 | [optional] 
**NeedLocation** | Pointer to **bool** | 本次响应是否包含坐标信息。 | [optional] 
**Timing** | Pointer to **map[string]interface{}** | 耗时拆分信息，适合做性能统计或排查。 | [optional] 
**Summary** | Pointer to **map[string]interface{}** | 识别结果的统计摘要。 | [optional] 
**Image** | Pointer to **map[string]interface{}** | 图片本身的基础信息。 | [optional] 
**Lines** | Pointer to **[]map[string]interface{}** | 按行组织的详细识别结果。 | [optional] 
**Blocks** | Pointer to **[]map[string]interface{}** | 按块组织的详细识别结果。 | [optional] 
**Pages** | Pointer to **[]map[string]interface{}** | 按页组织的详细识别结果。 | [optional] 
**Raw** | Pointer to **map[string]interface{}** | 补充识别结果对象，适合需要继续解析更多细节字段的场景。 | [optional] 

## Methods

### NewPostImageOcr200Response

`func NewPostImageOcr200Response() *PostImageOcr200Response`

NewPostImageOcr200Response instantiates a new PostImageOcr200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImageOcr200ResponseWithDefaults

`func NewPostImageOcr200ResponseWithDefaults() *PostImageOcr200Response`

NewPostImageOcr200ResponseWithDefaults instantiates a new PostImageOcr200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PostImageOcr200Response) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostImageOcr200Response) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostImageOcr200Response) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostImageOcr200Response) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPlainText

`func (o *PostImageOcr200Response) GetPlainText() string`

GetPlainText returns the PlainText field if non-nil, zero value otherwise.

### GetPlainTextOk

`func (o *PostImageOcr200Response) GetPlainTextOk() (*string, bool)`

GetPlainTextOk returns a tuple with the PlainText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainText

`func (o *PostImageOcr200Response) SetPlainText(v string)`

SetPlainText sets PlainText field to given value.

### HasPlainText

`func (o *PostImageOcr200Response) HasPlainText() bool`

HasPlainText returns a boolean if a field has been set.

### GetMarkdown

`func (o *PostImageOcr200Response) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *PostImageOcr200Response) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *PostImageOcr200Response) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.

### HasMarkdown

`func (o *PostImageOcr200Response) HasMarkdown() bool`

HasMarkdown returns a boolean if a field has been set.

### GetWordsResult

`func (o *PostImageOcr200Response) GetWordsResult() []PostImageOcr200ResponseWordsResultInner`

GetWordsResult returns the WordsResult field if non-nil, zero value otherwise.

### GetWordsResultOk

`func (o *PostImageOcr200Response) GetWordsResultOk() (*[]PostImageOcr200ResponseWordsResultInner, bool)`

GetWordsResultOk returns a tuple with the WordsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordsResult

`func (o *PostImageOcr200Response) SetWordsResult(v []PostImageOcr200ResponseWordsResultInner)`

SetWordsResult sets WordsResult field to given value.

### HasWordsResult

`func (o *PostImageOcr200Response) HasWordsResult() bool`

HasWordsResult returns a boolean if a field has been set.

### GetWordsResultNum

`func (o *PostImageOcr200Response) GetWordsResultNum() int32`

GetWordsResultNum returns the WordsResultNum field if non-nil, zero value otherwise.

### GetWordsResultNumOk

`func (o *PostImageOcr200Response) GetWordsResultNumOk() (*int32, bool)`

GetWordsResultNumOk returns a tuple with the WordsResultNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordsResultNum

`func (o *PostImageOcr200Response) SetWordsResultNum(v int32)`

SetWordsResultNum sets WordsResultNum field to given value.

### HasWordsResultNum

`func (o *PostImageOcr200Response) HasWordsResultNum() bool`

HasWordsResultNum returns a boolean if a field has been set.

### GetNeedLocation

`func (o *PostImageOcr200Response) GetNeedLocation() bool`

GetNeedLocation returns the NeedLocation field if non-nil, zero value otherwise.

### GetNeedLocationOk

`func (o *PostImageOcr200Response) GetNeedLocationOk() (*bool, bool)`

GetNeedLocationOk returns a tuple with the NeedLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedLocation

`func (o *PostImageOcr200Response) SetNeedLocation(v bool)`

SetNeedLocation sets NeedLocation field to given value.

### HasNeedLocation

`func (o *PostImageOcr200Response) HasNeedLocation() bool`

HasNeedLocation returns a boolean if a field has been set.

### GetTiming

`func (o *PostImageOcr200Response) GetTiming() map[string]interface{}`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *PostImageOcr200Response) GetTimingOk() (*map[string]interface{}, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *PostImageOcr200Response) SetTiming(v map[string]interface{})`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *PostImageOcr200Response) HasTiming() bool`

HasTiming returns a boolean if a field has been set.

### GetSummary

`func (o *PostImageOcr200Response) GetSummary() map[string]interface{}`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PostImageOcr200Response) GetSummaryOk() (*map[string]interface{}, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PostImageOcr200Response) SetSummary(v map[string]interface{})`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PostImageOcr200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetImage

`func (o *PostImageOcr200Response) GetImage() map[string]interface{}`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PostImageOcr200Response) GetImageOk() (*map[string]interface{}, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PostImageOcr200Response) SetImage(v map[string]interface{})`

SetImage sets Image field to given value.

### HasImage

`func (o *PostImageOcr200Response) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLines

`func (o *PostImageOcr200Response) GetLines() []map[string]interface{}`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *PostImageOcr200Response) GetLinesOk() (*[]map[string]interface{}, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *PostImageOcr200Response) SetLines(v []map[string]interface{})`

SetLines sets Lines field to given value.

### HasLines

`func (o *PostImageOcr200Response) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetBlocks

`func (o *PostImageOcr200Response) GetBlocks() []map[string]interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *PostImageOcr200Response) GetBlocksOk() (*[]map[string]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *PostImageOcr200Response) SetBlocks(v []map[string]interface{})`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *PostImageOcr200Response) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetPages

`func (o *PostImageOcr200Response) GetPages() []map[string]interface{}`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PostImageOcr200Response) GetPagesOk() (*[]map[string]interface{}, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PostImageOcr200Response) SetPages(v []map[string]interface{})`

SetPages sets Pages field to given value.

### HasPages

`func (o *PostImageOcr200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetRaw

`func (o *PostImageOcr200Response) GetRaw() map[string]interface{}`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *PostImageOcr200Response) GetRawOk() (*map[string]interface{}, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *PostImageOcr200Response) SetRaw(v map[string]interface{})`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *PostImageOcr200Response) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostTextConvertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | 待转换的文本内容 | 
**From** | **string** | 源格式类型 | 
**To** | **string** | 目标格式类型 | 
**Options** | Pointer to **map[string]interface{}** | 可选参数（预留，当前未使用） | [optional] 

## Methods

### NewPostTextConvertRequest

`func NewPostTextConvertRequest(text string, from string, to string, ) *PostTextConvertRequest`

NewPostTextConvertRequest instantiates a new PostTextConvertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextConvertRequestWithDefaults

`func NewPostTextConvertRequestWithDefaults() *PostTextConvertRequest`

NewPostTextConvertRequestWithDefaults instantiates a new PostTextConvertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PostTextConvertRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTextConvertRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTextConvertRequest) SetText(v string)`

SetText sets Text field to given value.


### GetFrom

`func (o *PostTextConvertRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PostTextConvertRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PostTextConvertRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *PostTextConvertRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PostTextConvertRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PostTextConvertRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetOptions

`func (o *PostTextConvertRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PostTextConvertRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PostTextConvertRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PostTextConvertRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



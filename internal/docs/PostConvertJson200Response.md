# PostConvertJson200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码，200代表操作成功。 | [optional] 
**Content** | Pointer to **string** | 格式化后的JSON字符串，带有标准缩进和换行。 | [optional] 

## Methods

### NewPostConvertJson200Response

`func NewPostConvertJson200Response() *PostConvertJson200Response`

NewPostConvertJson200Response instantiates a new PostConvertJson200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConvertJson200ResponseWithDefaults

`func NewPostConvertJson200ResponseWithDefaults() *PostConvertJson200Response`

NewPostConvertJson200ResponseWithDefaults instantiates a new PostConvertJson200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PostConvertJson200Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostConvertJson200Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostConvertJson200Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PostConvertJson200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContent

`func (o *PostConvertJson200Response) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PostConvertJson200Response) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PostConvertJson200Response) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PostConvertJson200Response) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostConvertJson400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | 机器可读的错误代码。 | [optional] 
**Details** | Pointer to **map[string]interface{}** | 包含错误详情的对象。 | [optional] 
**Message** | Pointer to **string** | 人类可读的错误描述信息。 | [optional] 

## Methods

### NewPostConvertJson400Response

`func NewPostConvertJson400Response() *PostConvertJson400Response`

NewPostConvertJson400Response instantiates a new PostConvertJson400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConvertJson400ResponseWithDefaults

`func NewPostConvertJson400ResponseWithDefaults() *PostConvertJson400Response`

NewPostConvertJson400ResponseWithDefaults instantiates a new PostConvertJson400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PostConvertJson400Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostConvertJson400Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostConvertJson400Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PostConvertJson400Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *PostConvertJson400Response) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PostConvertJson400Response) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PostConvertJson400Response) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PostConvertJson400Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *PostConvertJson400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostConvertJson400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostConvertJson400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostConvertJson400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



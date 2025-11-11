# GetImageTobase64200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base64** | Pointer to **string** | 转换后的完整Base64 Data URI，可以直接在CSS或HTML中使用。 | [optional] 
**Code** | Pointer to **int32** | 状态码，200代表成功。 | [optional] 
**Msg** | Pointer to **string** | 操作结果描述。 | [optional] 

## Methods

### NewGetImageTobase64200Response

`func NewGetImageTobase64200Response() *GetImageTobase64200Response`

NewGetImageTobase64200Response instantiates a new GetImageTobase64200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageTobase64200ResponseWithDefaults

`func NewGetImageTobase64200ResponseWithDefaults() *GetImageTobase64200Response`

NewGetImageTobase64200ResponseWithDefaults instantiates a new GetImageTobase64200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase64

`func (o *GetImageTobase64200Response) GetBase64() string`

GetBase64 returns the Base64 field if non-nil, zero value otherwise.

### GetBase64Ok

`func (o *GetImageTobase64200Response) GetBase64Ok() (*string, bool)`

GetBase64Ok returns a tuple with the Base64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64

`func (o *GetImageTobase64200Response) SetBase64(v string)`

SetBase64 sets Base64 field to given value.

### HasBase64

`func (o *GetImageTobase64200Response) HasBase64() bool`

HasBase64 returns a boolean if a field has been set.

### GetCode

`func (o *GetImageTobase64200Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetImageTobase64200Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetImageTobase64200Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetImageTobase64200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *GetImageTobase64200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetImageTobase64200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetImageTobase64200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetImageTobase64200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



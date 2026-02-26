# GetMiscLunartime200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 业务状态码，200 表示成功。 | [optional] 
**Message** | Pointer to **string** | 状态描述。 | [optional] 
**Data** | Pointer to [**GetMiscLunartime200ResponseData**](GetMiscLunartime200ResponseData.md) |  | [optional] 

## Methods

### NewGetMiscLunartime200Response

`func NewGetMiscLunartime200Response() *GetMiscLunartime200Response`

NewGetMiscLunartime200Response instantiates a new GetMiscLunartime200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscLunartime200ResponseWithDefaults

`func NewGetMiscLunartime200ResponseWithDefaults() *GetMiscLunartime200Response`

NewGetMiscLunartime200ResponseWithDefaults instantiates a new GetMiscLunartime200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetMiscLunartime200Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetMiscLunartime200Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetMiscLunartime200Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetMiscLunartime200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetMiscLunartime200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetMiscLunartime200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetMiscLunartime200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetMiscLunartime200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *GetMiscLunartime200Response) GetData() GetMiscLunartime200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMiscLunartime200Response) GetDataOk() (*GetMiscLunartime200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMiscLunartime200Response) SetData(v GetMiscLunartime200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetMiscLunartime200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



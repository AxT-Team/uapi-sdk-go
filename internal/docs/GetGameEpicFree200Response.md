# GetGameEpicFree200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码，200代表成功。 | [optional] 
**Data** | Pointer to [**[]GetGameEpicFree200ResponseDataInner**](GetGameEpicFree200ResponseDataInner.md) | 免费游戏列表数组。 | [optional] 

## Methods

### NewGetGameEpicFree200Response

`func NewGetGameEpicFree200Response() *GetGameEpicFree200Response`

NewGetGameEpicFree200Response instantiates a new GetGameEpicFree200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameEpicFree200ResponseWithDefaults

`func NewGetGameEpicFree200ResponseWithDefaults() *GetGameEpicFree200Response`

NewGetGameEpicFree200ResponseWithDefaults instantiates a new GetGameEpicFree200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetGameEpicFree200Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetGameEpicFree200Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetGameEpicFree200Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetGameEpicFree200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetGameEpicFree200Response) GetData() []GetGameEpicFree200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetGameEpicFree200Response) GetDataOk() (*[]GetGameEpicFree200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetGameEpicFree200Response) SetData(v []GetGameEpicFree200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetGameEpicFree200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



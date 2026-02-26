# PostTextConvert200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | 转换结果 | [optional] 
**From** | Pointer to **string** | 源格式 | [optional] 
**To** | Pointer to **string** | 目标格式 | [optional] 
**Length** | Pointer to **int32** | 结果长度 | [optional] 
**Info** | Pointer to **string** | 额外信息（如哈希不可逆提示） | [optional] 

## Methods

### NewPostTextConvert200Response

`func NewPostTextConvert200Response() *PostTextConvert200Response`

NewPostTextConvert200Response instantiates a new PostTextConvert200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextConvert200ResponseWithDefaults

`func NewPostTextConvert200ResponseWithDefaults() *PostTextConvert200Response`

NewPostTextConvert200ResponseWithDefaults instantiates a new PostTextConvert200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *PostTextConvert200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PostTextConvert200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PostTextConvert200Response) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *PostTextConvert200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFrom

`func (o *PostTextConvert200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PostTextConvert200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PostTextConvert200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PostTextConvert200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *PostTextConvert200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PostTextConvert200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PostTextConvert200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *PostTextConvert200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetLength

`func (o *PostTextConvert200Response) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PostTextConvert200Response) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PostTextConvert200Response) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PostTextConvert200Response) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetInfo

`func (o *PostTextConvert200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PostTextConvert200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PostTextConvert200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PostTextConvert200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



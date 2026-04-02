# PostAiTranslate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**PostAiTranslate200ResponseData**](PostAiTranslate200ResponseData.md) |  | [optional] 
**Performance** | Pointer to [**PostAiTranslate200ResponsePerformance**](PostAiTranslate200ResponsePerformance.md) |  | [optional] 
**IsBatch** | Pointer to **bool** | 是否为批量翻译请求。 | [optional] 

## Methods

### NewPostAiTranslate200Response

`func NewPostAiTranslate200Response() *PostAiTranslate200Response`

NewPostAiTranslate200Response instantiates a new PostAiTranslate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAiTranslate200ResponseWithDefaults

`func NewPostAiTranslate200ResponseWithDefaults() *PostAiTranslate200Response`

NewPostAiTranslate200ResponseWithDefaults instantiates a new PostAiTranslate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PostAiTranslate200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostAiTranslate200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostAiTranslate200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostAiTranslate200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *PostAiTranslate200Response) GetData() PostAiTranslate200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostAiTranslate200Response) GetDataOk() (*PostAiTranslate200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostAiTranslate200Response) SetData(v PostAiTranslate200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *PostAiTranslate200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPerformance

`func (o *PostAiTranslate200Response) GetPerformance() PostAiTranslate200ResponsePerformance`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *PostAiTranslate200Response) GetPerformanceOk() (*PostAiTranslate200ResponsePerformance, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *PostAiTranslate200Response) SetPerformance(v PostAiTranslate200ResponsePerformance)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *PostAiTranslate200Response) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetIsBatch

`func (o *PostAiTranslate200Response) GetIsBatch() bool`

GetIsBatch returns the IsBatch field if non-nil, zero value otherwise.

### GetIsBatchOk

`func (o *PostAiTranslate200Response) GetIsBatchOk() (*bool, bool)`

GetIsBatchOk returns a tuple with the IsBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBatch

`func (o *PostAiTranslate200Response) SetIsBatch(v bool)`

SetIsBatch sets IsBatch field to given value.

### HasIsBatch

`func (o *PostAiTranslate200Response) HasIsBatch() bool`

HasIsBatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



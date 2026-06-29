# PostWatermarkDecode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to **float32** | 检测结果的置信度，取值范围 0-1。 | [optional] 
**DecodeMs** | Pointer to **float32** | 解析耗时（毫秒）。 | [optional] 
**Payload** | Pointer to **string** | 还原出来的标识内容（仅在 present&#x3D;true 时返回）。 | [optional] 
**Present** | Pointer to **bool** | 是否在图片中检测到隐形水印。 | [optional] 

## Methods

### NewPostWatermarkDecode200Response

`func NewPostWatermarkDecode200Response() *PostWatermarkDecode200Response`

NewPostWatermarkDecode200Response instantiates a new PostWatermarkDecode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWatermarkDecode200ResponseWithDefaults

`func NewPostWatermarkDecode200ResponseWithDefaults() *PostWatermarkDecode200Response`

NewPostWatermarkDecode200ResponseWithDefaults instantiates a new PostWatermarkDecode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *PostWatermarkDecode200Response) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PostWatermarkDecode200Response) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PostWatermarkDecode200Response) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *PostWatermarkDecode200Response) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetDecodeMs

`func (o *PostWatermarkDecode200Response) GetDecodeMs() float32`

GetDecodeMs returns the DecodeMs field if non-nil, zero value otherwise.

### GetDecodeMsOk

`func (o *PostWatermarkDecode200Response) GetDecodeMsOk() (*float32, bool)`

GetDecodeMsOk returns a tuple with the DecodeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodeMs

`func (o *PostWatermarkDecode200Response) SetDecodeMs(v float32)`

SetDecodeMs sets DecodeMs field to given value.

### HasDecodeMs

`func (o *PostWatermarkDecode200Response) HasDecodeMs() bool`

HasDecodeMs returns a boolean if a field has been set.

### GetPayload

`func (o *PostWatermarkDecode200Response) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PostWatermarkDecode200Response) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PostWatermarkDecode200Response) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PostWatermarkDecode200Response) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPresent

`func (o *PostWatermarkDecode200Response) GetPresent() bool`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *PostWatermarkDecode200Response) GetPresentOk() (*bool, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *PostWatermarkDecode200Response) SetPresent(v bool)`

SetPresent sets Present field to given value.

### HasPresent

`func (o *PostWatermarkDecode200Response) HasPresent() bool`

HasPresent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostWatermarkLabel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | Pointer to **[]string** | 本次实际注入成功的标识层级，可能包含 &#39;watermark&#39;、&#39;explicit&#39;、&#39;metadata&#39;。 | [optional] 
**CapacityChars** | Pointer to **int32** | 当前配置下的隐形水印最大容量（若开启）。 | [optional] 
**ContentProducer** | Pointer to **string** | 成功写入的服务提供者编码。 | [optional] 
**Format** | Pointer to **string** | 实际输出的图片格式。 | [optional] 
**ImageBase64** | Pointer to **string** | 处理完成后的图片 Base64 编码。 | [optional] 
**ImageName** | Pointer to **string** | 原始图片文件名（若请求中包含则返回）。 | [optional] 
**WatermarkPayload** | Pointer to **string** | 成功嵌入的隐形水印内容（若开启）。 | [optional] 

## Methods

### NewPostWatermarkLabel200Response

`func NewPostWatermarkLabel200Response() *PostWatermarkLabel200Response`

NewPostWatermarkLabel200Response instantiates a new PostWatermarkLabel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWatermarkLabel200ResponseWithDefaults

`func NewPostWatermarkLabel200ResponseWithDefaults() *PostWatermarkLabel200Response`

NewPostWatermarkLabel200ResponseWithDefaults instantiates a new PostWatermarkLabel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *PostWatermarkLabel200Response) GetApplied() []string`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *PostWatermarkLabel200Response) GetAppliedOk() (*[]string, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *PostWatermarkLabel200Response) SetApplied(v []string)`

SetApplied sets Applied field to given value.

### HasApplied

`func (o *PostWatermarkLabel200Response) HasApplied() bool`

HasApplied returns a boolean if a field has been set.

### GetCapacityChars

`func (o *PostWatermarkLabel200Response) GetCapacityChars() int32`

GetCapacityChars returns the CapacityChars field if non-nil, zero value otherwise.

### GetCapacityCharsOk

`func (o *PostWatermarkLabel200Response) GetCapacityCharsOk() (*int32, bool)`

GetCapacityCharsOk returns a tuple with the CapacityChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityChars

`func (o *PostWatermarkLabel200Response) SetCapacityChars(v int32)`

SetCapacityChars sets CapacityChars field to given value.

### HasCapacityChars

`func (o *PostWatermarkLabel200Response) HasCapacityChars() bool`

HasCapacityChars returns a boolean if a field has been set.

### GetContentProducer

`func (o *PostWatermarkLabel200Response) GetContentProducer() string`

GetContentProducer returns the ContentProducer field if non-nil, zero value otherwise.

### GetContentProducerOk

`func (o *PostWatermarkLabel200Response) GetContentProducerOk() (*string, bool)`

GetContentProducerOk returns a tuple with the ContentProducer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProducer

`func (o *PostWatermarkLabel200Response) SetContentProducer(v string)`

SetContentProducer sets ContentProducer field to given value.

### HasContentProducer

`func (o *PostWatermarkLabel200Response) HasContentProducer() bool`

HasContentProducer returns a boolean if a field has been set.

### GetFormat

`func (o *PostWatermarkLabel200Response) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PostWatermarkLabel200Response) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PostWatermarkLabel200Response) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PostWatermarkLabel200Response) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetImageBase64

`func (o *PostWatermarkLabel200Response) GetImageBase64() string`

GetImageBase64 returns the ImageBase64 field if non-nil, zero value otherwise.

### GetImageBase64Ok

`func (o *PostWatermarkLabel200Response) GetImageBase64Ok() (*string, bool)`

GetImageBase64Ok returns a tuple with the ImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBase64

`func (o *PostWatermarkLabel200Response) SetImageBase64(v string)`

SetImageBase64 sets ImageBase64 field to given value.

### HasImageBase64

`func (o *PostWatermarkLabel200Response) HasImageBase64() bool`

HasImageBase64 returns a boolean if a field has been set.

### GetImageName

`func (o *PostWatermarkLabel200Response) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *PostWatermarkLabel200Response) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *PostWatermarkLabel200Response) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *PostWatermarkLabel200Response) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetWatermarkPayload

`func (o *PostWatermarkLabel200Response) GetWatermarkPayload() string`

GetWatermarkPayload returns the WatermarkPayload field if non-nil, zero value otherwise.

### GetWatermarkPayloadOk

`func (o *PostWatermarkLabel200Response) GetWatermarkPayloadOk() (*string, bool)`

GetWatermarkPayloadOk returns a tuple with the WatermarkPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkPayload

`func (o *PostWatermarkLabel200Response) SetWatermarkPayload(v string)`

SetWatermarkPayload sets WatermarkPayload field to given value.

### HasWatermarkPayload

`func (o *PostWatermarkLabel200Response) HasWatermarkPayload() bool`

HasWatermarkPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



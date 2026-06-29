# PostWatermarkEmbed200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityChars** | Pointer to **int32** | 在当前图片和配置下，能够嵌入的最大字符数。 | [optional] 
**EmbedMs** | Pointer to **float32** | 处理耗时（毫秒）。 | [optional] 
**Format** | Pointer to **string** | 实际输出的图片格式。 | [optional] 
**ImageBase64** | Pointer to **string** | 处理完成后的图片 Base64 编码。 | [optional] 
**ImageName** | Pointer to **string** | 原始图片文件名（若请求中包含则返回）。 | [optional] 
**Payload** | Pointer to **string** | 实际嵌入的标识内容。 | [optional] 

## Methods

### NewPostWatermarkEmbed200Response

`func NewPostWatermarkEmbed200Response() *PostWatermarkEmbed200Response`

NewPostWatermarkEmbed200Response instantiates a new PostWatermarkEmbed200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWatermarkEmbed200ResponseWithDefaults

`func NewPostWatermarkEmbed200ResponseWithDefaults() *PostWatermarkEmbed200Response`

NewPostWatermarkEmbed200ResponseWithDefaults instantiates a new PostWatermarkEmbed200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityChars

`func (o *PostWatermarkEmbed200Response) GetCapacityChars() int32`

GetCapacityChars returns the CapacityChars field if non-nil, zero value otherwise.

### GetCapacityCharsOk

`func (o *PostWatermarkEmbed200Response) GetCapacityCharsOk() (*int32, bool)`

GetCapacityCharsOk returns a tuple with the CapacityChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityChars

`func (o *PostWatermarkEmbed200Response) SetCapacityChars(v int32)`

SetCapacityChars sets CapacityChars field to given value.

### HasCapacityChars

`func (o *PostWatermarkEmbed200Response) HasCapacityChars() bool`

HasCapacityChars returns a boolean if a field has been set.

### GetEmbedMs

`func (o *PostWatermarkEmbed200Response) GetEmbedMs() float32`

GetEmbedMs returns the EmbedMs field if non-nil, zero value otherwise.

### GetEmbedMsOk

`func (o *PostWatermarkEmbed200Response) GetEmbedMsOk() (*float32, bool)`

GetEmbedMsOk returns a tuple with the EmbedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedMs

`func (o *PostWatermarkEmbed200Response) SetEmbedMs(v float32)`

SetEmbedMs sets EmbedMs field to given value.

### HasEmbedMs

`func (o *PostWatermarkEmbed200Response) HasEmbedMs() bool`

HasEmbedMs returns a boolean if a field has been set.

### GetFormat

`func (o *PostWatermarkEmbed200Response) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PostWatermarkEmbed200Response) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PostWatermarkEmbed200Response) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PostWatermarkEmbed200Response) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetImageBase64

`func (o *PostWatermarkEmbed200Response) GetImageBase64() string`

GetImageBase64 returns the ImageBase64 field if non-nil, zero value otherwise.

### GetImageBase64Ok

`func (o *PostWatermarkEmbed200Response) GetImageBase64Ok() (*string, bool)`

GetImageBase64Ok returns a tuple with the ImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBase64

`func (o *PostWatermarkEmbed200Response) SetImageBase64(v string)`

SetImageBase64 sets ImageBase64 field to given value.

### HasImageBase64

`func (o *PostWatermarkEmbed200Response) HasImageBase64() bool`

HasImageBase64 returns a boolean if a field has been set.

### GetImageName

`func (o *PostWatermarkEmbed200Response) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *PostWatermarkEmbed200Response) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *PostWatermarkEmbed200Response) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *PostWatermarkEmbed200Response) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetPayload

`func (o *PostWatermarkEmbed200Response) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PostWatermarkEmbed200Response) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PostWatermarkEmbed200Response) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PostWatermarkEmbed200Response) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostTextAesEncryptAdvancedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | 待加密的明文文本 | 
**Key** | **string** | 加密密钥（支持任意长度） | 
**Mode** | Pointer to **string** | 加密模式：GCM/CBC/ECB/CTR/OFB/CFB（可选，默认GCM） | [optional] 
**Padding** | Pointer to **string** | 填充方式：PKCS7/ZERO/NONE（可选，默认PKCS7） | [optional] 
**Iv** | Pointer to **string** | 自定义IV（可选，Base64编码，16字节）。GCM模式无需此参数 | [optional] 
**OutputFormat** | Pointer to **string** | 输出格式：base64（默认）或hex | [optional] 

## Methods

### NewPostTextAesEncryptAdvancedRequest

`func NewPostTextAesEncryptAdvancedRequest(text string, key string, ) *PostTextAesEncryptAdvancedRequest`

NewPostTextAesEncryptAdvancedRequest instantiates a new PostTextAesEncryptAdvancedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextAesEncryptAdvancedRequestWithDefaults

`func NewPostTextAesEncryptAdvancedRequestWithDefaults() *PostTextAesEncryptAdvancedRequest`

NewPostTextAesEncryptAdvancedRequestWithDefaults instantiates a new PostTextAesEncryptAdvancedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PostTextAesEncryptAdvancedRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTextAesEncryptAdvancedRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTextAesEncryptAdvancedRequest) SetText(v string)`

SetText sets Text field to given value.


### GetKey

`func (o *PostTextAesEncryptAdvancedRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostTextAesEncryptAdvancedRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostTextAesEncryptAdvancedRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetMode

`func (o *PostTextAesEncryptAdvancedRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostTextAesEncryptAdvancedRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostTextAesEncryptAdvancedRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostTextAesEncryptAdvancedRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPadding

`func (o *PostTextAesEncryptAdvancedRequest) GetPadding() string`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *PostTextAesEncryptAdvancedRequest) GetPaddingOk() (*string, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *PostTextAesEncryptAdvancedRequest) SetPadding(v string)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *PostTextAesEncryptAdvancedRequest) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetIv

`func (o *PostTextAesEncryptAdvancedRequest) GetIv() string`

GetIv returns the Iv field if non-nil, zero value otherwise.

### GetIvOk

`func (o *PostTextAesEncryptAdvancedRequest) GetIvOk() (*string, bool)`

GetIvOk returns a tuple with the Iv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIv

`func (o *PostTextAesEncryptAdvancedRequest) SetIv(v string)`

SetIv sets Iv field to given value.

### HasIv

`func (o *PostTextAesEncryptAdvancedRequest) HasIv() bool`

HasIv returns a boolean if a field has been set.

### GetOutputFormat

`func (o *PostTextAesEncryptAdvancedRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *PostTextAesEncryptAdvancedRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *PostTextAesEncryptAdvancedRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *PostTextAesEncryptAdvancedRequest) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



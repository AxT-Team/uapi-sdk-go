# PostTextAesDecryptAdvancedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | 待解密的密文（Base64编码）。此值来自加密接口返回的ciphertext字段 | 
**Key** | **string** | 解密密钥（必须与加密时相同） | 
**Mode** | **string** | 加密模式（必须与加密时相同）：GCM/CBC/ECB/CTR/OFB/CFB | 
**Padding** | Pointer to **string** | 填充方式（可选，必须与加密时相同）：PKCS7/ZERO/NONE。GCM模式默认为NONE | [optional] 
**Iv** | Pointer to **string** | 初始化向量（非GCM模式必须提供，Base64编码）。此值来自加密接口返回的iv字段 | [optional] 

## Methods

### NewPostTextAesDecryptAdvancedRequest

`func NewPostTextAesDecryptAdvancedRequest(text string, key string, mode string, ) *PostTextAesDecryptAdvancedRequest`

NewPostTextAesDecryptAdvancedRequest instantiates a new PostTextAesDecryptAdvancedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextAesDecryptAdvancedRequestWithDefaults

`func NewPostTextAesDecryptAdvancedRequestWithDefaults() *PostTextAesDecryptAdvancedRequest`

NewPostTextAesDecryptAdvancedRequestWithDefaults instantiates a new PostTextAesDecryptAdvancedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PostTextAesDecryptAdvancedRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTextAesDecryptAdvancedRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTextAesDecryptAdvancedRequest) SetText(v string)`

SetText sets Text field to given value.


### GetKey

`func (o *PostTextAesDecryptAdvancedRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostTextAesDecryptAdvancedRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostTextAesDecryptAdvancedRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetMode

`func (o *PostTextAesDecryptAdvancedRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostTextAesDecryptAdvancedRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostTextAesDecryptAdvancedRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPadding

`func (o *PostTextAesDecryptAdvancedRequest) GetPadding() string`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *PostTextAesDecryptAdvancedRequest) GetPaddingOk() (*string, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *PostTextAesDecryptAdvancedRequest) SetPadding(v string)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *PostTextAesDecryptAdvancedRequest) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetIv

`func (o *PostTextAesDecryptAdvancedRequest) GetIv() string`

GetIv returns the Iv field if non-nil, zero value otherwise.

### GetIvOk

`func (o *PostTextAesDecryptAdvancedRequest) GetIvOk() (*string, bool)`

GetIvOk returns a tuple with the Iv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIv

`func (o *PostTextAesDecryptAdvancedRequest) SetIv(v string)`

SetIv sets Iv field to given value.

### HasIv

`func (o *PostTextAesDecryptAdvancedRequest) HasIv() bool`

HasIv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



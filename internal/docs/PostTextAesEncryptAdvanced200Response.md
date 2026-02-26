# PostTextAesEncryptAdvanced200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** | 加密后的密文（Base64编码） | [optional] 
**Mode** | Pointer to **string** | 使用的加密模式 | [optional] 
**Padding** | Pointer to **string** | 使用的填充方式 | [optional] 
**Iv** | Pointer to **string** | 使用的IV（Base64编码）。GCM模式不返回此字段 | [optional] 

## Methods

### NewPostTextAesEncryptAdvanced200Response

`func NewPostTextAesEncryptAdvanced200Response() *PostTextAesEncryptAdvanced200Response`

NewPostTextAesEncryptAdvanced200Response instantiates a new PostTextAesEncryptAdvanced200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextAesEncryptAdvanced200ResponseWithDefaults

`func NewPostTextAesEncryptAdvanced200ResponseWithDefaults() *PostTextAesEncryptAdvanced200Response`

NewPostTextAesEncryptAdvanced200ResponseWithDefaults instantiates a new PostTextAesEncryptAdvanced200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *PostTextAesEncryptAdvanced200Response) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *PostTextAesEncryptAdvanced200Response) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *PostTextAesEncryptAdvanced200Response) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *PostTextAesEncryptAdvanced200Response) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetMode

`func (o *PostTextAesEncryptAdvanced200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostTextAesEncryptAdvanced200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostTextAesEncryptAdvanced200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostTextAesEncryptAdvanced200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPadding

`func (o *PostTextAesEncryptAdvanced200Response) GetPadding() string`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *PostTextAesEncryptAdvanced200Response) GetPaddingOk() (*string, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *PostTextAesEncryptAdvanced200Response) SetPadding(v string)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *PostTextAesEncryptAdvanced200Response) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetIv

`func (o *PostTextAesEncryptAdvanced200Response) GetIv() string`

GetIv returns the Iv field if non-nil, zero value otherwise.

### GetIvOk

`func (o *PostTextAesEncryptAdvanced200Response) GetIvOk() (*string, bool)`

GetIvOk returns a tuple with the Iv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIv

`func (o *PostTextAesEncryptAdvanced200Response) SetIv(v string)`

SetIv sets Iv field to given value.

### HasIv

`func (o *PostTextAesEncryptAdvanced200Response) HasIv() bool`

HasIv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



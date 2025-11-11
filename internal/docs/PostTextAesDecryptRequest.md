# PostTextAesDecryptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | 密钥，长度必须为16、24或32字节，对应AES-128/192/256。 | 
**Text** | **string** | Base64编码的密文。 | 
**Nonce** | **string** | 16�ֽڵ�IV/Nonce����Ϊ16���ַ� | 

## Methods

### NewPostTextAesDecryptRequest

`func NewPostTextAesDecryptRequest(key string, text string, nonce string, ) *PostTextAesDecryptRequest`

NewPostTextAesDecryptRequest instantiates a new PostTextAesDecryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextAesDecryptRequestWithDefaults

`func NewPostTextAesDecryptRequestWithDefaults() *PostTextAesDecryptRequest`

NewPostTextAesDecryptRequestWithDefaults instantiates a new PostTextAesDecryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostTextAesDecryptRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostTextAesDecryptRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostTextAesDecryptRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetText

`func (o *PostTextAesDecryptRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTextAesDecryptRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTextAesDecryptRequest) SetText(v string)`

SetText sets Text field to given value.


### GetNonce

`func (o *PostTextAesDecryptRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *PostTextAesDecryptRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *PostTextAesDecryptRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostTextAesEncryptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key must be 16, 24, or 32 bytes long to select AES-128, AES-192, or AES-256. | 
**Text** | **string** |  | 

## Methods

### NewPostTextAesEncryptRequest

`func NewPostTextAesEncryptRequest(key string, text string, ) *PostTextAesEncryptRequest`

NewPostTextAesEncryptRequest instantiates a new PostTextAesEncryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTextAesEncryptRequestWithDefaults

`func NewPostTextAesEncryptRequestWithDefaults() *PostTextAesEncryptRequest`

NewPostTextAesEncryptRequestWithDefaults instantiates a new PostTextAesEncryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostTextAesEncryptRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostTextAesEncryptRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostTextAesEncryptRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetText

`func (o *PostTextAesEncryptRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTextAesEncryptRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTextAesEncryptRequest) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



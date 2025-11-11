# PostClipzyStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressedData** | **string** | 必需：经过加密和 LZString 压缩后的 Base64 字符串。请参考文档首页的JS代码示例。 | 
**Ttl** | Pointer to **float32** | 可选：片段的留存时间（秒）。正数表示秒数（最大约30天），-1 表示永久存储。默认为 3600。 | [optional] 

## Methods

### NewPostClipzyStoreRequest

`func NewPostClipzyStoreRequest(compressedData string, ) *PostClipzyStoreRequest`

NewPostClipzyStoreRequest instantiates a new PostClipzyStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostClipzyStoreRequestWithDefaults

`func NewPostClipzyStoreRequestWithDefaults() *PostClipzyStoreRequest`

NewPostClipzyStoreRequestWithDefaults instantiates a new PostClipzyStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressedData

`func (o *PostClipzyStoreRequest) GetCompressedData() string`

GetCompressedData returns the CompressedData field if non-nil, zero value otherwise.

### GetCompressedDataOk

`func (o *PostClipzyStoreRequest) GetCompressedDataOk() (*string, bool)`

GetCompressedDataOk returns a tuple with the CompressedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedData

`func (o *PostClipzyStoreRequest) SetCompressedData(v string)`

SetCompressedData sets CompressedData field to given value.


### GetTtl

`func (o *PostClipzyStoreRequest) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PostClipzyStoreRequest) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PostClipzyStoreRequest) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PostClipzyStoreRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



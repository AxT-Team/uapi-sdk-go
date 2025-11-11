# PostImageFrombase64Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageData** | **string** | 图片的Base64 Data URI，必须包含MIME类型头。例如：&#x60;data:image/png;base64,...&#x60; | 

## Methods

### NewPostImageFrombase64Request

`func NewPostImageFrombase64Request(imageData string, ) *PostImageFrombase64Request`

NewPostImageFrombase64Request instantiates a new PostImageFrombase64Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImageFrombase64RequestWithDefaults

`func NewPostImageFrombase64RequestWithDefaults() *PostImageFrombase64Request`

NewPostImageFrombase64RequestWithDefaults instantiates a new PostImageFrombase64Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageData

`func (o *PostImageFrombase64Request) GetImageData() string`

GetImageData returns the ImageData field if non-nil, zero value otherwise.

### GetImageDataOk

`func (o *PostImageFrombase64Request) GetImageDataOk() (*string, bool)`

GetImageDataOk returns a tuple with the ImageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageData

`func (o *PostImageFrombase64Request) SetImageData(v string)`

SetImageData sets ImageData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



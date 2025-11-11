# PostImageSpeechlessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopText** | Pointer to **string** | 表情包上方的文字内容。你们怎么不说话了，是不是都在偷偷 _______ | [optional] 
**BottomText** | Pointer to **string** | 表情包下方的文字内容。求求你_______ | [optional] 

## Methods

### NewPostImageSpeechlessRequest

`func NewPostImageSpeechlessRequest() *PostImageSpeechlessRequest`

NewPostImageSpeechlessRequest instantiates a new PostImageSpeechlessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImageSpeechlessRequestWithDefaults

`func NewPostImageSpeechlessRequestWithDefaults() *PostImageSpeechlessRequest`

NewPostImageSpeechlessRequestWithDefaults instantiates a new PostImageSpeechlessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopText

`func (o *PostImageSpeechlessRequest) GetTopText() string`

GetTopText returns the TopText field if non-nil, zero value otherwise.

### GetTopTextOk

`func (o *PostImageSpeechlessRequest) GetTopTextOk() (*string, bool)`

GetTopTextOk returns a tuple with the TopText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopText

`func (o *PostImageSpeechlessRequest) SetTopText(v string)`

SetTopText sets TopText field to given value.

### HasTopText

`func (o *PostImageSpeechlessRequest) HasTopText() bool`

HasTopText returns a boolean if a field has been set.

### GetBottomText

`func (o *PostImageSpeechlessRequest) GetBottomText() string`

GetBottomText returns the BottomText field if non-nil, zero value otherwise.

### GetBottomTextOk

`func (o *PostImageSpeechlessRequest) GetBottomTextOk() (*string, bool)`

GetBottomTextOk returns a tuple with the BottomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomText

`func (o *PostImageSpeechlessRequest) SetBottomText(v string)`

SetBottomText sets BottomText field to given value.

### HasBottomText

`func (o *PostImageSpeechlessRequest) HasBottomText() bool`

HasBottomText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



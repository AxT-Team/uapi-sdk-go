# PostImageOcr200ResponseWordsResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**PostImageOcr200ResponseWordsResultInnerLocation**](PostImageOcr200ResponseWordsResultInnerLocation.md) |  | [optional] 
**Score** | Pointer to **float32** | 当前文字片段的置信度。部分结果会返回。 | [optional] 
**VertexesLocation** | Pointer to [**[]PostImageOcr200ResponseWordsResultInnerVertexesLocationInner**](PostImageOcr200ResponseWordsResultInnerVertexesLocationInner.md) | 当前文字片段的顶点坐标列表。只有在 &#x60;need_location&#x3D;true&#x60; 时才会返回。 | [optional] 
**Words** | Pointer to **string** | 当前文字片段的识别结果。 | [optional] 

## Methods

### NewPostImageOcr200ResponseWordsResultInner

`func NewPostImageOcr200ResponseWordsResultInner() *PostImageOcr200ResponseWordsResultInner`

NewPostImageOcr200ResponseWordsResultInner instantiates a new PostImageOcr200ResponseWordsResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImageOcr200ResponseWordsResultInnerWithDefaults

`func NewPostImageOcr200ResponseWordsResultInnerWithDefaults() *PostImageOcr200ResponseWordsResultInner`

NewPostImageOcr200ResponseWordsResultInnerWithDefaults instantiates a new PostImageOcr200ResponseWordsResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PostImageOcr200ResponseWordsResultInner) GetLocation() PostImageOcr200ResponseWordsResultInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostImageOcr200ResponseWordsResultInner) GetLocationOk() (*PostImageOcr200ResponseWordsResultInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostImageOcr200ResponseWordsResultInner) SetLocation(v PostImageOcr200ResponseWordsResultInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PostImageOcr200ResponseWordsResultInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetScore

`func (o *PostImageOcr200ResponseWordsResultInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PostImageOcr200ResponseWordsResultInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PostImageOcr200ResponseWordsResultInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *PostImageOcr200ResponseWordsResultInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetVertexesLocation

`func (o *PostImageOcr200ResponseWordsResultInner) GetVertexesLocation() []PostImageOcr200ResponseWordsResultInnerVertexesLocationInner`

GetVertexesLocation returns the VertexesLocation field if non-nil, zero value otherwise.

### GetVertexesLocationOk

`func (o *PostImageOcr200ResponseWordsResultInner) GetVertexesLocationOk() (*[]PostImageOcr200ResponseWordsResultInnerVertexesLocationInner, bool)`

GetVertexesLocationOk returns a tuple with the VertexesLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexesLocation

`func (o *PostImageOcr200ResponseWordsResultInner) SetVertexesLocation(v []PostImageOcr200ResponseWordsResultInnerVertexesLocationInner)`

SetVertexesLocation sets VertexesLocation field to given value.

### HasVertexesLocation

`func (o *PostImageOcr200ResponseWordsResultInner) HasVertexesLocation() bool`

HasVertexesLocation returns a boolean if a field has been set.

### GetWords

`func (o *PostImageOcr200ResponseWordsResultInner) GetWords() string`

GetWords returns the Words field if non-nil, zero value otherwise.

### GetWordsOk

`func (o *PostImageOcr200ResponseWordsResultInner) GetWordsOk() (*string, bool)`

GetWordsOk returns a tuple with the Words field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWords

`func (o *PostImageOcr200ResponseWordsResultInner) SetWords(v string)`

SetWords sets Words field to given value.

### HasWords

`func (o *PostImageOcr200ResponseWordsResultInner) HasWords() bool`

HasWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



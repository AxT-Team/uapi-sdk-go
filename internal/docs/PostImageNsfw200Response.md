# PostImageNsfw200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsfwScore** | Pointer to **float32** | NSFW 内容的置信度分数，范围 0-1，越高表示越可能是敏感内容。 | [optional] 
**NormalScore** | Pointer to **float32** | 正常内容的置信度分数，范围 0-1。 | [optional] 
**IsNsfw** | Pointer to **bool** | 是否判定为 NSFW 内容。 | [optional] 
**Label** | Pointer to **string** | 内容标签，&#39;nsfw&#39; 或 &#39;normal&#39;。 | [optional] 
**Suggestion** | Pointer to **string** | 处理建议：&#39;pass&#39;（通过）、&#39;review&#39;（人工复核）、&#39;block&#39;（拦截）。 | [optional] 
**RiskLevel** | Pointer to **string** | 风险等级：&#39;low&#39;、&#39;medium&#39;、&#39;high&#39;。 | [optional] 
**Confidence** | Pointer to **float32** | 模型对当前判断的置信度。 | [optional] 
**InferenceTimeMs** | Pointer to **float32** | 模型推理耗时，单位毫秒。 | [optional] 

## Methods

### NewPostImageNsfw200Response

`func NewPostImageNsfw200Response() *PostImageNsfw200Response`

NewPostImageNsfw200Response instantiates a new PostImageNsfw200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImageNsfw200ResponseWithDefaults

`func NewPostImageNsfw200ResponseWithDefaults() *PostImageNsfw200Response`

NewPostImageNsfw200ResponseWithDefaults instantiates a new PostImageNsfw200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNsfwScore

`func (o *PostImageNsfw200Response) GetNsfwScore() float32`

GetNsfwScore returns the NsfwScore field if non-nil, zero value otherwise.

### GetNsfwScoreOk

`func (o *PostImageNsfw200Response) GetNsfwScoreOk() (*float32, bool)`

GetNsfwScoreOk returns a tuple with the NsfwScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwScore

`func (o *PostImageNsfw200Response) SetNsfwScore(v float32)`

SetNsfwScore sets NsfwScore field to given value.

### HasNsfwScore

`func (o *PostImageNsfw200Response) HasNsfwScore() bool`

HasNsfwScore returns a boolean if a field has been set.

### GetNormalScore

`func (o *PostImageNsfw200Response) GetNormalScore() float32`

GetNormalScore returns the NormalScore field if non-nil, zero value otherwise.

### GetNormalScoreOk

`func (o *PostImageNsfw200Response) GetNormalScoreOk() (*float32, bool)`

GetNormalScoreOk returns a tuple with the NormalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalScore

`func (o *PostImageNsfw200Response) SetNormalScore(v float32)`

SetNormalScore sets NormalScore field to given value.

### HasNormalScore

`func (o *PostImageNsfw200Response) HasNormalScore() bool`

HasNormalScore returns a boolean if a field has been set.

### GetIsNsfw

`func (o *PostImageNsfw200Response) GetIsNsfw() bool`

GetIsNsfw returns the IsNsfw field if non-nil, zero value otherwise.

### GetIsNsfwOk

`func (o *PostImageNsfw200Response) GetIsNsfwOk() (*bool, bool)`

GetIsNsfwOk returns a tuple with the IsNsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNsfw

`func (o *PostImageNsfw200Response) SetIsNsfw(v bool)`

SetIsNsfw sets IsNsfw field to given value.

### HasIsNsfw

`func (o *PostImageNsfw200Response) HasIsNsfw() bool`

HasIsNsfw returns a boolean if a field has been set.

### GetLabel

`func (o *PostImageNsfw200Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostImageNsfw200Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostImageNsfw200Response) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostImageNsfw200Response) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSuggestion

`func (o *PostImageNsfw200Response) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *PostImageNsfw200Response) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *PostImageNsfw200Response) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.

### HasSuggestion

`func (o *PostImageNsfw200Response) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.

### GetRiskLevel

`func (o *PostImageNsfw200Response) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *PostImageNsfw200Response) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *PostImageNsfw200Response) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *PostImageNsfw200Response) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetConfidence

`func (o *PostImageNsfw200Response) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PostImageNsfw200Response) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PostImageNsfw200Response) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *PostImageNsfw200Response) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetInferenceTimeMs

`func (o *PostImageNsfw200Response) GetInferenceTimeMs() float32`

GetInferenceTimeMs returns the InferenceTimeMs field if non-nil, zero value otherwise.

### GetInferenceTimeMsOk

`func (o *PostImageNsfw200Response) GetInferenceTimeMsOk() (*float32, bool)`

GetInferenceTimeMsOk returns a tuple with the InferenceTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceTimeMs

`func (o *PostImageNsfw200Response) SetInferenceTimeMs(v float32)`

SetInferenceTimeMs sets InferenceTimeMs field to given value.

### HasInferenceTimeMs

`func (o *PostImageNsfw200Response) HasInferenceTimeMs() bool`

HasInferenceTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostSensitiveWordAnalyzeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keywords** | **[]string** | 要分析的关键词列表，单次最多100个，每个关键词最长50字符。 | 

## Methods

### NewPostSensitiveWordAnalyzeRequest

`func NewPostSensitiveWordAnalyzeRequest(keywords []string, ) *PostSensitiveWordAnalyzeRequest`

NewPostSensitiveWordAnalyzeRequest instantiates a new PostSensitiveWordAnalyzeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSensitiveWordAnalyzeRequestWithDefaults

`func NewPostSensitiveWordAnalyzeRequestWithDefaults() *PostSensitiveWordAnalyzeRequest`

NewPostSensitiveWordAnalyzeRequestWithDefaults instantiates a new PostSensitiveWordAnalyzeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywords

`func (o *PostSensitiveWordAnalyzeRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *PostSensitiveWordAnalyzeRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *PostSensitiveWordAnalyzeRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



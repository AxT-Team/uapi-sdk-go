# PostMiscDateDiffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | 开始日期，支持多种格式自动识别 | 
**EndDate** | **string** | 结束日期，支持多种格式自动识别 | 
**Format** | Pointer to **string** | 日期格式（可选），如DD-MM-YYYY。不指定则自动识别 | [optional] 

## Methods

### NewPostMiscDateDiffRequest

`func NewPostMiscDateDiffRequest(startDate string, endDate string, ) *PostMiscDateDiffRequest`

NewPostMiscDateDiffRequest instantiates a new PostMiscDateDiffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMiscDateDiffRequestWithDefaults

`func NewPostMiscDateDiffRequestWithDefaults() *PostMiscDateDiffRequest`

NewPostMiscDateDiffRequestWithDefaults instantiates a new PostMiscDateDiffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *PostMiscDateDiffRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PostMiscDateDiffRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PostMiscDateDiffRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *PostMiscDateDiffRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PostMiscDateDiffRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PostMiscDateDiffRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetFormat

`func (o *PostMiscDateDiffRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PostMiscDateDiffRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PostMiscDateDiffRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PostMiscDateDiffRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



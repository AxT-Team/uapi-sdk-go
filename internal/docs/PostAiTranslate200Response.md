# PostAiTranslate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**IsBatch** | Pointer to **bool** | 标识是否为批量翻译请求。 | [optional] 
**Data** | Pointer to [**PostAiTranslate200ResponseData**](PostAiTranslate200ResponseData.md) |  | [optional] 
**BatchData** | Pointer to [**[]PostAiTranslate200ResponseBatchDataInner**](PostAiTranslate200ResponseBatchDataInner.md) | 批量翻译结果列表，仅在批量翻译时返回。 | [optional] 
**BatchSummary** | Pointer to [**PostAiTranslate200ResponseBatchSummary**](PostAiTranslate200ResponseBatchSummary.md) |  | [optional] 
**Performance** | Pointer to [**PostAiTranslate200ResponsePerformance**](PostAiTranslate200ResponsePerformance.md) |  | [optional] 
**QualityMetrics** | Pointer to [**PostAiTranslate200ResponseQualityMetrics**](PostAiTranslate200ResponseQualityMetrics.md) |  | [optional] 

## Methods

### NewPostAiTranslate200Response

`func NewPostAiTranslate200Response() *PostAiTranslate200Response`

NewPostAiTranslate200Response instantiates a new PostAiTranslate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAiTranslate200ResponseWithDefaults

`func NewPostAiTranslate200ResponseWithDefaults() *PostAiTranslate200Response`

NewPostAiTranslate200ResponseWithDefaults instantiates a new PostAiTranslate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PostAiTranslate200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostAiTranslate200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostAiTranslate200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostAiTranslate200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetIsBatch

`func (o *PostAiTranslate200Response) GetIsBatch() bool`

GetIsBatch returns the IsBatch field if non-nil, zero value otherwise.

### GetIsBatchOk

`func (o *PostAiTranslate200Response) GetIsBatchOk() (*bool, bool)`

GetIsBatchOk returns a tuple with the IsBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBatch

`func (o *PostAiTranslate200Response) SetIsBatch(v bool)`

SetIsBatch sets IsBatch field to given value.

### HasIsBatch

`func (o *PostAiTranslate200Response) HasIsBatch() bool`

HasIsBatch returns a boolean if a field has been set.

### GetData

`func (o *PostAiTranslate200Response) GetData() PostAiTranslate200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostAiTranslate200Response) GetDataOk() (*PostAiTranslate200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostAiTranslate200Response) SetData(v PostAiTranslate200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *PostAiTranslate200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetBatchData

`func (o *PostAiTranslate200Response) GetBatchData() []PostAiTranslate200ResponseBatchDataInner`

GetBatchData returns the BatchData field if non-nil, zero value otherwise.

### GetBatchDataOk

`func (o *PostAiTranslate200Response) GetBatchDataOk() (*[]PostAiTranslate200ResponseBatchDataInner, bool)`

GetBatchDataOk returns a tuple with the BatchData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchData

`func (o *PostAiTranslate200Response) SetBatchData(v []PostAiTranslate200ResponseBatchDataInner)`

SetBatchData sets BatchData field to given value.

### HasBatchData

`func (o *PostAiTranslate200Response) HasBatchData() bool`

HasBatchData returns a boolean if a field has been set.

### GetBatchSummary

`func (o *PostAiTranslate200Response) GetBatchSummary() PostAiTranslate200ResponseBatchSummary`

GetBatchSummary returns the BatchSummary field if non-nil, zero value otherwise.

### GetBatchSummaryOk

`func (o *PostAiTranslate200Response) GetBatchSummaryOk() (*PostAiTranslate200ResponseBatchSummary, bool)`

GetBatchSummaryOk returns a tuple with the BatchSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSummary

`func (o *PostAiTranslate200Response) SetBatchSummary(v PostAiTranslate200ResponseBatchSummary)`

SetBatchSummary sets BatchSummary field to given value.

### HasBatchSummary

`func (o *PostAiTranslate200Response) HasBatchSummary() bool`

HasBatchSummary returns a boolean if a field has been set.

### GetPerformance

`func (o *PostAiTranslate200Response) GetPerformance() PostAiTranslate200ResponsePerformance`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *PostAiTranslate200Response) GetPerformanceOk() (*PostAiTranslate200ResponsePerformance, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *PostAiTranslate200Response) SetPerformance(v PostAiTranslate200ResponsePerformance)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *PostAiTranslate200Response) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetQualityMetrics

`func (o *PostAiTranslate200Response) GetQualityMetrics() PostAiTranslate200ResponseQualityMetrics`

GetQualityMetrics returns the QualityMetrics field if non-nil, zero value otherwise.

### GetQualityMetricsOk

`func (o *PostAiTranslate200Response) GetQualityMetricsOk() (*PostAiTranslate200ResponseQualityMetrics, bool)`

GetQualityMetricsOk returns a tuple with the QualityMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityMetrics

`func (o *PostAiTranslate200Response) SetQualityMetrics(v PostAiTranslate200ResponseQualityMetrics)`

SetQualityMetrics sets QualityMetrics field to given value.

### HasQualityMetrics

`func (o *PostAiTranslate200Response) HasQualityMetrics() bool`

HasQualityMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetMiscWeather200ResponseMinutelyPrecip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to **string** | 降水描述 | [optional] 
**UpdateTime** | Pointer to **string** | 更新时间 | [optional] 
**Data** | Pointer to [**[]GetMiscWeather200ResponseMinutelyPrecipDataInner**](GetMiscWeather200ResponseMinutelyPrecipDataInner.md) | 每5分钟一个数据点，共24个 | [optional] 

## Methods

### NewGetMiscWeather200ResponseMinutelyPrecip

`func NewGetMiscWeather200ResponseMinutelyPrecip() *GetMiscWeather200ResponseMinutelyPrecip`

NewGetMiscWeather200ResponseMinutelyPrecip instantiates a new GetMiscWeather200ResponseMinutelyPrecip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscWeather200ResponseMinutelyPrecipWithDefaults

`func NewGetMiscWeather200ResponseMinutelyPrecipWithDefaults() *GetMiscWeather200ResponseMinutelyPrecip`

NewGetMiscWeather200ResponseMinutelyPrecipWithDefaults instantiates a new GetMiscWeather200ResponseMinutelyPrecip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *GetMiscWeather200ResponseMinutelyPrecip) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetMiscWeather200ResponseMinutelyPrecip) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetMiscWeather200ResponseMinutelyPrecip) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetMiscWeather200ResponseMinutelyPrecip) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetMiscWeather200ResponseMinutelyPrecip) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetMiscWeather200ResponseMinutelyPrecip) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetMiscWeather200ResponseMinutelyPrecip) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetMiscWeather200ResponseMinutelyPrecip) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetData

`func (o *GetMiscWeather200ResponseMinutelyPrecip) GetData() []GetMiscWeather200ResponseMinutelyPrecipDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMiscWeather200ResponseMinutelyPrecip) GetDataOk() (*[]GetMiscWeather200ResponseMinutelyPrecipDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMiscWeather200ResponseMinutelyPrecip) SetData(v []GetMiscWeather200ResponseMinutelyPrecipDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetMiscWeather200ResponseMinutelyPrecip) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



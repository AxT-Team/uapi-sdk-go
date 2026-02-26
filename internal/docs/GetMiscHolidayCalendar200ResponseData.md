# GetMiscHolidayCalendar200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | 查询模式：day、month、year。 | [optional] 
**Query** | Pointer to [**GetMiscHolidayCalendar200ResponseDataQuery**](GetMiscHolidayCalendar200ResponseDataQuery.md) |  | [optional] 
**Summary** | Pointer to [**GetMiscHolidayCalendar200ResponseDataSummary**](GetMiscHolidayCalendar200ResponseDataSummary.md) |  | [optional] 
**Days** | Pointer to [**[]GetMiscHolidayCalendar200ResponseDataDaysInner**](GetMiscHolidayCalendar200ResponseDataDaysInner.md) | 日期明细列表。 | [optional] 
**Holidays** | Pointer to [**[]GetMiscHolidayCalendar200ResponseDataHolidaysInner**](GetMiscHolidayCalendar200ResponseDataHolidaysInner.md) | 节日事件列表。 | [optional] 
**Nearby** | Pointer to [**GetMiscHolidayCalendar200ResponseDataNearby**](GetMiscHolidayCalendar200ResponseDataNearby.md) |  | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseData

`func NewGetMiscHolidayCalendar200ResponseData() *GetMiscHolidayCalendar200ResponseData`

NewGetMiscHolidayCalendar200ResponseData instantiates a new GetMiscHolidayCalendar200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseDataWithDefaults

`func NewGetMiscHolidayCalendar200ResponseDataWithDefaults() *GetMiscHolidayCalendar200ResponseData`

NewGetMiscHolidayCalendar200ResponseDataWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetMiscHolidayCalendar200ResponseData) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetMiscHolidayCalendar200ResponseData) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetMiscHolidayCalendar200ResponseData) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetMiscHolidayCalendar200ResponseData) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetQuery

`func (o *GetMiscHolidayCalendar200ResponseData) GetQuery() GetMiscHolidayCalendar200ResponseDataQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetMiscHolidayCalendar200ResponseData) GetQueryOk() (*GetMiscHolidayCalendar200ResponseDataQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetMiscHolidayCalendar200ResponseData) SetQuery(v GetMiscHolidayCalendar200ResponseDataQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GetMiscHolidayCalendar200ResponseData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSummary

`func (o *GetMiscHolidayCalendar200ResponseData) GetSummary() GetMiscHolidayCalendar200ResponseDataSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetMiscHolidayCalendar200ResponseData) GetSummaryOk() (*GetMiscHolidayCalendar200ResponseDataSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetMiscHolidayCalendar200ResponseData) SetSummary(v GetMiscHolidayCalendar200ResponseDataSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetMiscHolidayCalendar200ResponseData) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDays

`func (o *GetMiscHolidayCalendar200ResponseData) GetDays() []GetMiscHolidayCalendar200ResponseDataDaysInner`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *GetMiscHolidayCalendar200ResponseData) GetDaysOk() (*[]GetMiscHolidayCalendar200ResponseDataDaysInner, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *GetMiscHolidayCalendar200ResponseData) SetDays(v []GetMiscHolidayCalendar200ResponseDataDaysInner)`

SetDays sets Days field to given value.

### HasDays

`func (o *GetMiscHolidayCalendar200ResponseData) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetHolidays

`func (o *GetMiscHolidayCalendar200ResponseData) GetHolidays() []GetMiscHolidayCalendar200ResponseDataHolidaysInner`

GetHolidays returns the Holidays field if non-nil, zero value otherwise.

### GetHolidaysOk

`func (o *GetMiscHolidayCalendar200ResponseData) GetHolidaysOk() (*[]GetMiscHolidayCalendar200ResponseDataHolidaysInner, bool)`

GetHolidaysOk returns a tuple with the Holidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidays

`func (o *GetMiscHolidayCalendar200ResponseData) SetHolidays(v []GetMiscHolidayCalendar200ResponseDataHolidaysInner)`

SetHolidays sets Holidays field to given value.

### HasHolidays

`func (o *GetMiscHolidayCalendar200ResponseData) HasHolidays() bool`

HasHolidays returns a boolean if a field has been set.

### GetNearby

`func (o *GetMiscHolidayCalendar200ResponseData) GetNearby() GetMiscHolidayCalendar200ResponseDataNearby`

GetNearby returns the Nearby field if non-nil, zero value otherwise.

### GetNearbyOk

`func (o *GetMiscHolidayCalendar200ResponseData) GetNearbyOk() (*GetMiscHolidayCalendar200ResponseDataNearby, bool)`

GetNearbyOk returns a tuple with the Nearby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearby

`func (o *GetMiscHolidayCalendar200ResponseData) SetNearby(v GetMiscHolidayCalendar200ResponseDataNearby)`

SetNearby sets Nearby field to given value.

### HasNearby

`func (o *GetMiscHolidayCalendar200ResponseData) HasNearby() bool`

HasNearby returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



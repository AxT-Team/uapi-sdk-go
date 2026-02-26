# GetMiscHolidayCalendar200ResponseDataSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalDays** | Pointer to **int32** | 查询范围内总天数。 | [optional] 
**WeekendDays** | Pointer to **int32** | 查询范围内周末天数。 | [optional] 
**Workdays** | Pointer to **int32** | 查询范围内工作日天数（含法定调休上班）。 | [optional] 
**RestDays** | Pointer to **int32** | 查询范围内休息日天数（含周末和法定休假）。 | [optional] 
**HolidayEvents** | Pointer to **int32** | 按 holiday_type 过滤后的节日事件总数。 | [optional] 
**LegalRestDays** | Pointer to **int32** | 法定休假日天数。 | [optional] 
**LegalWorkdays** | Pointer to **int32** | 法定调休上班天数。 | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseDataSummary

`func NewGetMiscHolidayCalendar200ResponseDataSummary() *GetMiscHolidayCalendar200ResponseDataSummary`

NewGetMiscHolidayCalendar200ResponseDataSummary instantiates a new GetMiscHolidayCalendar200ResponseDataSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseDataSummaryWithDefaults

`func NewGetMiscHolidayCalendar200ResponseDataSummaryWithDefaults() *GetMiscHolidayCalendar200ResponseDataSummary`

NewGetMiscHolidayCalendar200ResponseDataSummaryWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseDataSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetTotalDays() int32`

GetTotalDays returns the TotalDays field if non-nil, zero value otherwise.

### GetTotalDaysOk

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetTotalDaysOk() (*int32, bool)`

GetTotalDaysOk returns a tuple with the TotalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) SetTotalDays(v int32)`

SetTotalDays sets TotalDays field to given value.

### HasTotalDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) HasTotalDays() bool`

HasTotalDays returns a boolean if a field has been set.

### GetWeekendDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetWeekendDays() int32`

GetWeekendDays returns the WeekendDays field if non-nil, zero value otherwise.

### GetWeekendDaysOk

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetWeekendDaysOk() (*int32, bool)`

GetWeekendDaysOk returns a tuple with the WeekendDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekendDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) SetWeekendDays(v int32)`

SetWeekendDays sets WeekendDays field to given value.

### HasWeekendDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) HasWeekendDays() bool`

HasWeekendDays returns a boolean if a field has been set.

### GetWorkdays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetWorkdays() int32`

GetWorkdays returns the Workdays field if non-nil, zero value otherwise.

### GetWorkdaysOk

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetWorkdaysOk() (*int32, bool)`

GetWorkdaysOk returns a tuple with the Workdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkdays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) SetWorkdays(v int32)`

SetWorkdays sets Workdays field to given value.

### HasWorkdays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) HasWorkdays() bool`

HasWorkdays returns a boolean if a field has been set.

### GetRestDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetRestDays() int32`

GetRestDays returns the RestDays field if non-nil, zero value otherwise.

### GetRestDaysOk

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetRestDaysOk() (*int32, bool)`

GetRestDaysOk returns a tuple with the RestDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) SetRestDays(v int32)`

SetRestDays sets RestDays field to given value.

### HasRestDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) HasRestDays() bool`

HasRestDays returns a boolean if a field has been set.

### GetHolidayEvents

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetHolidayEvents() int32`

GetHolidayEvents returns the HolidayEvents field if non-nil, zero value otherwise.

### GetHolidayEventsOk

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetHolidayEventsOk() (*int32, bool)`

GetHolidayEventsOk returns a tuple with the HolidayEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayEvents

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) SetHolidayEvents(v int32)`

SetHolidayEvents sets HolidayEvents field to given value.

### HasHolidayEvents

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) HasHolidayEvents() bool`

HasHolidayEvents returns a boolean if a field has been set.

### GetLegalRestDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetLegalRestDays() int32`

GetLegalRestDays returns the LegalRestDays field if non-nil, zero value otherwise.

### GetLegalRestDaysOk

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetLegalRestDaysOk() (*int32, bool)`

GetLegalRestDaysOk returns a tuple with the LegalRestDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalRestDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) SetLegalRestDays(v int32)`

SetLegalRestDays sets LegalRestDays field to given value.

### HasLegalRestDays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) HasLegalRestDays() bool`

HasLegalRestDays returns a boolean if a field has been set.

### GetLegalWorkdays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetLegalWorkdays() int32`

GetLegalWorkdays returns the LegalWorkdays field if non-nil, zero value otherwise.

### GetLegalWorkdaysOk

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) GetLegalWorkdaysOk() (*int32, bool)`

GetLegalWorkdaysOk returns a tuple with the LegalWorkdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalWorkdays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) SetLegalWorkdays(v int32)`

SetLegalWorkdays sets LegalWorkdays field to given value.

### HasLegalWorkdays

`func (o *GetMiscHolidayCalendar200ResponseDataSummary) HasLegalWorkdays() bool`

HasLegalWorkdays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



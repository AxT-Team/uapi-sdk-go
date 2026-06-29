# GetMiscHolidayCalendar200ResponseQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | 日视图查询参数。date 模式下为 YYYY-MM-DD，其余模式下为空字符串。 | [optional] 
**ExcludePast** | Pointer to **bool** | 是否过滤今天之前已经过去的节日。 | [optional] 
**HolidayType** | Pointer to **string** | 节日筛选类型。 | [optional] 
**IncludeNearby** | Pointer to **bool** | 是否开启前后最近节日查询。 | [optional] 
**Month** | Pointer to **string** | 月视图查询参数。month 模式下为 YYYY-MM，其余模式下为空字符串。 | [optional] 
**NearbyLimit** | Pointer to **int32** | 前后最近节日返回数量上限。 | [optional] 
**Timezone** | Pointer to **string** | 实际生效的时区。 | [optional] 
**Year** | Pointer to **string** | 年视图查询参数。year 模式下为 YYYY，其余模式下为空字符串。 | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseQuery

`func NewGetMiscHolidayCalendar200ResponseQuery() *GetMiscHolidayCalendar200ResponseQuery`

NewGetMiscHolidayCalendar200ResponseQuery instantiates a new GetMiscHolidayCalendar200ResponseQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseQueryWithDefaults

`func NewGetMiscHolidayCalendar200ResponseQueryWithDefaults() *GetMiscHolidayCalendar200ResponseQuery`

NewGetMiscHolidayCalendar200ResponseQueryWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetExcludePast

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetExcludePast() bool`

GetExcludePast returns the ExcludePast field if non-nil, zero value otherwise.

### GetExcludePastOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetExcludePastOk() (*bool, bool)`

GetExcludePastOk returns a tuple with the ExcludePast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePast

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetExcludePast(v bool)`

SetExcludePast sets ExcludePast field to given value.

### HasExcludePast

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasExcludePast() bool`

HasExcludePast returns a boolean if a field has been set.

### GetHolidayType

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetHolidayType() string`

GetHolidayType returns the HolidayType field if non-nil, zero value otherwise.

### GetHolidayTypeOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetHolidayTypeOk() (*string, bool)`

GetHolidayTypeOk returns a tuple with the HolidayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayType

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetHolidayType(v string)`

SetHolidayType sets HolidayType field to given value.

### HasHolidayType

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasHolidayType() bool`

HasHolidayType returns a boolean if a field has been set.

### GetIncludeNearby

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetIncludeNearby() bool`

GetIncludeNearby returns the IncludeNearby field if non-nil, zero value otherwise.

### GetIncludeNearbyOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetIncludeNearbyOk() (*bool, bool)`

GetIncludeNearbyOk returns a tuple with the IncludeNearby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNearby

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetIncludeNearby(v bool)`

SetIncludeNearby sets IncludeNearby field to given value.

### HasIncludeNearby

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasIncludeNearby() bool`

HasIncludeNearby returns a boolean if a field has been set.

### GetMonth

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetNearbyLimit

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetNearbyLimit() int32`

GetNearbyLimit returns the NearbyLimit field if non-nil, zero value otherwise.

### GetNearbyLimitOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetNearbyLimitOk() (*int32, bool)`

GetNearbyLimitOk returns a tuple with the NearbyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearbyLimit

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetNearbyLimit(v int32)`

SetNearbyLimit sets NearbyLimit field to given value.

### HasNearbyLimit

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasNearbyLimit() bool`

HasNearbyLimit returns a boolean if a field has been set.

### GetTimezone

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetYear

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetMiscHolidayCalendar200ResponseQuery) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetMiscHolidayCalendar200ResponseQuery) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *GetMiscHolidayCalendar200ResponseQuery) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



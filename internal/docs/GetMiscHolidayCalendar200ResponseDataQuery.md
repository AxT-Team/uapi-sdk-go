# GetMiscHolidayCalendar200ResponseDataQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | 日视图查询参数（YYYY-MM-DD）。非日视图时可能为空。 | [optional] 
**Month** | Pointer to **string** | 月视图查询参数（YYYY-MM）。非月视图时可能为空。 | [optional] 
**Year** | Pointer to **string** | 年视图查询参数（YYYY）。非年视图时可能为空。 | [optional] 
**Timezone** | Pointer to **string** | 实际生效的时区。 | [optional] 
**HolidayType** | Pointer to **string** | 节日筛选类型。 | [optional] 
**IncludeNearby** | Pointer to **bool** | 是否开启前后最近节日查询。 | [optional] 
**NearbyLimit** | Pointer to **int32** | 前后最近节日返回数量上限。 | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseDataQuery

`func NewGetMiscHolidayCalendar200ResponseDataQuery() *GetMiscHolidayCalendar200ResponseDataQuery`

NewGetMiscHolidayCalendar200ResponseDataQuery instantiates a new GetMiscHolidayCalendar200ResponseDataQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseDataQueryWithDefaults

`func NewGetMiscHolidayCalendar200ResponseDataQueryWithDefaults() *GetMiscHolidayCalendar200ResponseDataQuery`

NewGetMiscHolidayCalendar200ResponseDataQueryWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseDataQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMonth

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetTimezone

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetHolidayType() string`

GetHolidayType returns the HolidayType field if non-nil, zero value otherwise.

### GetHolidayTypeOk

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetHolidayTypeOk() (*string, bool)`

GetHolidayTypeOk returns a tuple with the HolidayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) SetHolidayType(v string)`

SetHolidayType sets HolidayType field to given value.

### HasHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) HasHolidayType() bool`

HasHolidayType returns a boolean if a field has been set.

### GetIncludeNearby

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetIncludeNearby() bool`

GetIncludeNearby returns the IncludeNearby field if non-nil, zero value otherwise.

### GetIncludeNearbyOk

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetIncludeNearbyOk() (*bool, bool)`

GetIncludeNearbyOk returns a tuple with the IncludeNearby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNearby

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) SetIncludeNearby(v bool)`

SetIncludeNearby sets IncludeNearby field to given value.

### HasIncludeNearby

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) HasIncludeNearby() bool`

HasIncludeNearby returns a boolean if a field has been set.

### GetNearbyLimit

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetNearbyLimit() int32`

GetNearbyLimit returns the NearbyLimit field if non-nil, zero value otherwise.

### GetNearbyLimitOk

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) GetNearbyLimitOk() (*int32, bool)`

GetNearbyLimitOk returns a tuple with the NearbyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearbyLimit

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) SetNearbyLimit(v int32)`

SetNearbyLimit sets NearbyLimit field to given value.

### HasNearbyLimit

`func (o *GetMiscHolidayCalendar200ResponseDataQuery) HasNearbyLimit() bool`

HasNearbyLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetMiscHolidayCalendar200ResponseDataDaysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | 公历日期（YYYY-MM-DD）。 | [optional] 
**Year** | Pointer to **int32** | 公历年份。 | [optional] 
**Month** | Pointer to **int32** | 公历月份。 | [optional] 
**Day** | Pointer to **int32** | 公历日期（天）。 | [optional] 
**WeekdayCn** | Pointer to **string** | 中文星期，如星期三。 | [optional] 
**IsWeekend** | Pointer to **bool** | 是否为周末。 | [optional] 
**IsWorkday** | Pointer to **bool** | 是否为工作日（含法定调休上班日）。 | [optional] 
**IsRestDay** | Pointer to **bool** | 是否为休息日。 | [optional] 
**IsHoliday** | Pointer to **bool** | 当天是否存在节日/节气/法定事件。 | [optional] 
**LegalHolidayName** | Pointer to **string** | 法定节假日名称，无则为空。 | [optional] 
**LegalHolidayType** | Pointer to **string** | 法定假日类型：rest 或 workday_adjust。 | [optional] 
**SolarFestival** | Pointer to **string** | 公历节日名称，无则为空。 | [optional] 
**LunarFestival** | Pointer to **string** | 农历节日名称，无则为空。 | [optional] 
**SolarTerm** | Pointer to **string** | 节气名称，无则为空。 | [optional] 
**LunarYear** | Pointer to **int32** | 农历年份（数字）。 | [optional] 
**LunarMonth** | Pointer to **int32** | 农历月份（数字）。 | [optional] 
**LunarDay** | Pointer to **int32** | 农历日期（数字）。 | [optional] 
**LunarMonthName** | Pointer to **string** | 农历月份中文名称。 | [optional] 
**LunarDayName** | Pointer to **string** | 农历日期中文名称。 | [optional] 
**GanzhiYear** | Pointer to **string** | 干支年。 | [optional] 
**GanzhiMonth** | Pointer to **string** | 干支月。 | [optional] 
**GanzhiDay** | Pointer to **string** | 干支日。 | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseDataDaysInner

`func NewGetMiscHolidayCalendar200ResponseDataDaysInner() *GetMiscHolidayCalendar200ResponseDataDaysInner`

NewGetMiscHolidayCalendar200ResponseDataDaysInner instantiates a new GetMiscHolidayCalendar200ResponseDataDaysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseDataDaysInnerWithDefaults

`func NewGetMiscHolidayCalendar200ResponseDataDaysInnerWithDefaults() *GetMiscHolidayCalendar200ResponseDataDaysInner`

NewGetMiscHolidayCalendar200ResponseDataDaysInnerWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseDataDaysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetWeekdayCn

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetWeekdayCn() string`

GetWeekdayCn returns the WeekdayCn field if non-nil, zero value otherwise.

### GetWeekdayCnOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetWeekdayCnOk() (*string, bool)`

GetWeekdayCnOk returns a tuple with the WeekdayCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayCn

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetWeekdayCn(v string)`

SetWeekdayCn sets WeekdayCn field to given value.

### HasWeekdayCn

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasWeekdayCn() bool`

HasWeekdayCn returns a boolean if a field has been set.

### GetIsWeekend

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsWeekend() bool`

GetIsWeekend returns the IsWeekend field if non-nil, zero value otherwise.

### GetIsWeekendOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsWeekendOk() (*bool, bool)`

GetIsWeekendOk returns a tuple with the IsWeekend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeekend

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetIsWeekend(v bool)`

SetIsWeekend sets IsWeekend field to given value.

### HasIsWeekend

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasIsWeekend() bool`

HasIsWeekend returns a boolean if a field has been set.

### GetIsWorkday

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsWorkday() bool`

GetIsWorkday returns the IsWorkday field if non-nil, zero value otherwise.

### GetIsWorkdayOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsWorkdayOk() (*bool, bool)`

GetIsWorkdayOk returns a tuple with the IsWorkday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorkday

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetIsWorkday(v bool)`

SetIsWorkday sets IsWorkday field to given value.

### HasIsWorkday

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasIsWorkday() bool`

HasIsWorkday returns a boolean if a field has been set.

### GetIsRestDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsRestDay() bool`

GetIsRestDay returns the IsRestDay field if non-nil, zero value otherwise.

### GetIsRestDayOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsRestDayOk() (*bool, bool)`

GetIsRestDayOk returns a tuple with the IsRestDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetIsRestDay(v bool)`

SetIsRestDay sets IsRestDay field to given value.

### HasIsRestDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasIsRestDay() bool`

HasIsRestDay returns a boolean if a field has been set.

### GetIsHoliday

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsHoliday() bool`

GetIsHoliday returns the IsHoliday field if non-nil, zero value otherwise.

### GetIsHolidayOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetIsHolidayOk() (*bool, bool)`

GetIsHolidayOk returns a tuple with the IsHoliday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHoliday

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetIsHoliday(v bool)`

SetIsHoliday sets IsHoliday field to given value.

### HasIsHoliday

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasIsHoliday() bool`

HasIsHoliday returns a boolean if a field has been set.

### GetLegalHolidayName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLegalHolidayName() string`

GetLegalHolidayName returns the LegalHolidayName field if non-nil, zero value otherwise.

### GetLegalHolidayNameOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLegalHolidayNameOk() (*string, bool)`

GetLegalHolidayNameOk returns a tuple with the LegalHolidayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalHolidayName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLegalHolidayName(v string)`

SetLegalHolidayName sets LegalHolidayName field to given value.

### HasLegalHolidayName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLegalHolidayName() bool`

HasLegalHolidayName returns a boolean if a field has been set.

### GetLegalHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLegalHolidayType() string`

GetLegalHolidayType returns the LegalHolidayType field if non-nil, zero value otherwise.

### GetLegalHolidayTypeOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLegalHolidayTypeOk() (*string, bool)`

GetLegalHolidayTypeOk returns a tuple with the LegalHolidayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLegalHolidayType(v string)`

SetLegalHolidayType sets LegalHolidayType field to given value.

### HasLegalHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLegalHolidayType() bool`

HasLegalHolidayType returns a boolean if a field has been set.

### GetSolarFestival

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetSolarFestival() string`

GetSolarFestival returns the SolarFestival field if non-nil, zero value otherwise.

### GetSolarFestivalOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetSolarFestivalOk() (*string, bool)`

GetSolarFestivalOk returns a tuple with the SolarFestival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarFestival

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetSolarFestival(v string)`

SetSolarFestival sets SolarFestival field to given value.

### HasSolarFestival

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasSolarFestival() bool`

HasSolarFestival returns a boolean if a field has been set.

### GetLunarFestival

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarFestival() string`

GetLunarFestival returns the LunarFestival field if non-nil, zero value otherwise.

### GetLunarFestivalOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarFestivalOk() (*string, bool)`

GetLunarFestivalOk returns a tuple with the LunarFestival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarFestival

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLunarFestival(v string)`

SetLunarFestival sets LunarFestival field to given value.

### HasLunarFestival

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLunarFestival() bool`

HasLunarFestival returns a boolean if a field has been set.

### GetSolarTerm

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetSolarTerm() string`

GetSolarTerm returns the SolarTerm field if non-nil, zero value otherwise.

### GetSolarTermOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetSolarTermOk() (*string, bool)`

GetSolarTermOk returns a tuple with the SolarTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarTerm

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetSolarTerm(v string)`

SetSolarTerm sets SolarTerm field to given value.

### HasSolarTerm

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasSolarTerm() bool`

HasSolarTerm returns a boolean if a field has been set.

### GetLunarYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarYear() int32`

GetLunarYear returns the LunarYear field if non-nil, zero value otherwise.

### GetLunarYearOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarYearOk() (*int32, bool)`

GetLunarYearOk returns a tuple with the LunarYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLunarYear(v int32)`

SetLunarYear sets LunarYear field to given value.

### HasLunarYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLunarYear() bool`

HasLunarYear returns a boolean if a field has been set.

### GetLunarMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarMonth() int32`

GetLunarMonth returns the LunarMonth field if non-nil, zero value otherwise.

### GetLunarMonthOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarMonthOk() (*int32, bool)`

GetLunarMonthOk returns a tuple with the LunarMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLunarMonth(v int32)`

SetLunarMonth sets LunarMonth field to given value.

### HasLunarMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLunarMonth() bool`

HasLunarMonth returns a boolean if a field has been set.

### GetLunarDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarDay() int32`

GetLunarDay returns the LunarDay field if non-nil, zero value otherwise.

### GetLunarDayOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarDayOk() (*int32, bool)`

GetLunarDayOk returns a tuple with the LunarDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLunarDay(v int32)`

SetLunarDay sets LunarDay field to given value.

### HasLunarDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLunarDay() bool`

HasLunarDay returns a boolean if a field has been set.

### GetLunarMonthName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarMonthName() string`

GetLunarMonthName returns the LunarMonthName field if non-nil, zero value otherwise.

### GetLunarMonthNameOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarMonthNameOk() (*string, bool)`

GetLunarMonthNameOk returns a tuple with the LunarMonthName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarMonthName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLunarMonthName(v string)`

SetLunarMonthName sets LunarMonthName field to given value.

### HasLunarMonthName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLunarMonthName() bool`

HasLunarMonthName returns a boolean if a field has been set.

### GetLunarDayName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarDayName() string`

GetLunarDayName returns the LunarDayName field if non-nil, zero value otherwise.

### GetLunarDayNameOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetLunarDayNameOk() (*string, bool)`

GetLunarDayNameOk returns a tuple with the LunarDayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarDayName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetLunarDayName(v string)`

SetLunarDayName sets LunarDayName field to given value.

### HasLunarDayName

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasLunarDayName() bool`

HasLunarDayName returns a boolean if a field has been set.

### GetGanzhiYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetGanzhiYear() string`

GetGanzhiYear returns the GanzhiYear field if non-nil, zero value otherwise.

### GetGanzhiYearOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetGanzhiYearOk() (*string, bool)`

GetGanzhiYearOk returns a tuple with the GanzhiYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetGanzhiYear(v string)`

SetGanzhiYear sets GanzhiYear field to given value.

### HasGanzhiYear

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasGanzhiYear() bool`

HasGanzhiYear returns a boolean if a field has been set.

### GetGanzhiMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetGanzhiMonth() string`

GetGanzhiMonth returns the GanzhiMonth field if non-nil, zero value otherwise.

### GetGanzhiMonthOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetGanzhiMonthOk() (*string, bool)`

GetGanzhiMonthOk returns a tuple with the GanzhiMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetGanzhiMonth(v string)`

SetGanzhiMonth sets GanzhiMonth field to given value.

### HasGanzhiMonth

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasGanzhiMonth() bool`

HasGanzhiMonth returns a boolean if a field has been set.

### GetGanzhiDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetGanzhiDay() string`

GetGanzhiDay returns the GanzhiDay field if non-nil, zero value otherwise.

### GetGanzhiDayOk

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) GetGanzhiDayOk() (*string, bool)`

GetGanzhiDayOk returns a tuple with the GanzhiDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) SetGanzhiDay(v string)`

SetGanzhiDay sets GanzhiDay field to given value.

### HasGanzhiDay

`func (o *GetMiscHolidayCalendar200ResponseDataDaysInner) HasGanzhiDay() bool`

HasGanzhiDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetMiscHolidayCalendar200ResponseDaysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | 公历日期（YYYY-MM-DD）。 | [optional] 
**Day** | Pointer to **int32** | 公历日期（天）。 | [optional] 
**GanzhiDay** | Pointer to **string** | 干支日。 | [optional] 
**GanzhiMonth** | Pointer to **string** | 干支月。 | [optional] 
**GanzhiYear** | Pointer to **string** | 干支年。 | [optional] 
**IsHoliday** | Pointer to **bool** | 当天是否存在节日、节气或法定事件。 | [optional] 
**IsRestDay** | Pointer to **bool** | 是否为休息日。 | [optional] 
**IsWeekend** | Pointer to **bool** | 是否为周末。 | [optional] 
**IsWorkday** | Pointer to **bool** | 是否为工作日（含法定调休上班日）。 | [optional] 
**LegalHolidayName** | Pointer to **string** | 法定节假日名称，无则为空或不返回。 | [optional] 
**LegalHolidayType** | Pointer to **string** | 法定假日类型：rest 或 workday_adjust。 | [optional] 
**LunarDay** | Pointer to **int32** | 农历日期（数字）。 | [optional] 
**LunarDayName** | Pointer to **string** | 农历日期中文名称。 | [optional] 
**LunarFestival** | Pointer to **string** | 农历节日名称。有值时返回。 | [optional] 
**LunarMonth** | Pointer to **int32** | 农历月份（数字）。 | [optional] 
**LunarMonthName** | Pointer to **string** | 农历月份中文名称。 | [optional] 
**LunarYear** | Pointer to **int32** | 农历年份（数字）。 | [optional] 
**Month** | Pointer to **int32** | 公历月份。 | [optional] 
**SolarFestival** | Pointer to **string** | 公历节日名称。有值时返回。 | [optional] 
**SolarTerm** | Pointer to **string** | 节气名称。有值时返回。 | [optional] 
**WeekdayCn** | Pointer to **string** | 中文星期，如星期三。 | [optional] 
**Year** | Pointer to **int32** | 公历年份。 | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseDaysInner

`func NewGetMiscHolidayCalendar200ResponseDaysInner() *GetMiscHolidayCalendar200ResponseDaysInner`

NewGetMiscHolidayCalendar200ResponseDaysInner instantiates a new GetMiscHolidayCalendar200ResponseDaysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseDaysInnerWithDefaults

`func NewGetMiscHolidayCalendar200ResponseDaysInnerWithDefaults() *GetMiscHolidayCalendar200ResponseDaysInner`

NewGetMiscHolidayCalendar200ResponseDaysInnerWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseDaysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetGanzhiDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetGanzhiDay() string`

GetGanzhiDay returns the GanzhiDay field if non-nil, zero value otherwise.

### GetGanzhiDayOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetGanzhiDayOk() (*string, bool)`

GetGanzhiDayOk returns a tuple with the GanzhiDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetGanzhiDay(v string)`

SetGanzhiDay sets GanzhiDay field to given value.

### HasGanzhiDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasGanzhiDay() bool`

HasGanzhiDay returns a boolean if a field has been set.

### GetGanzhiMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetGanzhiMonth() string`

GetGanzhiMonth returns the GanzhiMonth field if non-nil, zero value otherwise.

### GetGanzhiMonthOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetGanzhiMonthOk() (*string, bool)`

GetGanzhiMonthOk returns a tuple with the GanzhiMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetGanzhiMonth(v string)`

SetGanzhiMonth sets GanzhiMonth field to given value.

### HasGanzhiMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasGanzhiMonth() bool`

HasGanzhiMonth returns a boolean if a field has been set.

### GetGanzhiYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetGanzhiYear() string`

GetGanzhiYear returns the GanzhiYear field if non-nil, zero value otherwise.

### GetGanzhiYearOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetGanzhiYearOk() (*string, bool)`

GetGanzhiYearOk returns a tuple with the GanzhiYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetGanzhiYear(v string)`

SetGanzhiYear sets GanzhiYear field to given value.

### HasGanzhiYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasGanzhiYear() bool`

HasGanzhiYear returns a boolean if a field has been set.

### GetIsHoliday

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsHoliday() bool`

GetIsHoliday returns the IsHoliday field if non-nil, zero value otherwise.

### GetIsHolidayOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsHolidayOk() (*bool, bool)`

GetIsHolidayOk returns a tuple with the IsHoliday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHoliday

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetIsHoliday(v bool)`

SetIsHoliday sets IsHoliday field to given value.

### HasIsHoliday

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasIsHoliday() bool`

HasIsHoliday returns a boolean if a field has been set.

### GetIsRestDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsRestDay() bool`

GetIsRestDay returns the IsRestDay field if non-nil, zero value otherwise.

### GetIsRestDayOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsRestDayOk() (*bool, bool)`

GetIsRestDayOk returns a tuple with the IsRestDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetIsRestDay(v bool)`

SetIsRestDay sets IsRestDay field to given value.

### HasIsRestDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasIsRestDay() bool`

HasIsRestDay returns a boolean if a field has been set.

### GetIsWeekend

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsWeekend() bool`

GetIsWeekend returns the IsWeekend field if non-nil, zero value otherwise.

### GetIsWeekendOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsWeekendOk() (*bool, bool)`

GetIsWeekendOk returns a tuple with the IsWeekend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeekend

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetIsWeekend(v bool)`

SetIsWeekend sets IsWeekend field to given value.

### HasIsWeekend

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasIsWeekend() bool`

HasIsWeekend returns a boolean if a field has been set.

### GetIsWorkday

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsWorkday() bool`

GetIsWorkday returns the IsWorkday field if non-nil, zero value otherwise.

### GetIsWorkdayOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetIsWorkdayOk() (*bool, bool)`

GetIsWorkdayOk returns a tuple with the IsWorkday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorkday

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetIsWorkday(v bool)`

SetIsWorkday sets IsWorkday field to given value.

### HasIsWorkday

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasIsWorkday() bool`

HasIsWorkday returns a boolean if a field has been set.

### GetLegalHolidayName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLegalHolidayName() string`

GetLegalHolidayName returns the LegalHolidayName field if non-nil, zero value otherwise.

### GetLegalHolidayNameOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLegalHolidayNameOk() (*string, bool)`

GetLegalHolidayNameOk returns a tuple with the LegalHolidayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalHolidayName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLegalHolidayName(v string)`

SetLegalHolidayName sets LegalHolidayName field to given value.

### HasLegalHolidayName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLegalHolidayName() bool`

HasLegalHolidayName returns a boolean if a field has been set.

### GetLegalHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLegalHolidayType() string`

GetLegalHolidayType returns the LegalHolidayType field if non-nil, zero value otherwise.

### GetLegalHolidayTypeOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLegalHolidayTypeOk() (*string, bool)`

GetLegalHolidayTypeOk returns a tuple with the LegalHolidayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLegalHolidayType(v string)`

SetLegalHolidayType sets LegalHolidayType field to given value.

### HasLegalHolidayType

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLegalHolidayType() bool`

HasLegalHolidayType returns a boolean if a field has been set.

### GetLunarDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarDay() int32`

GetLunarDay returns the LunarDay field if non-nil, zero value otherwise.

### GetLunarDayOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarDayOk() (*int32, bool)`

GetLunarDayOk returns a tuple with the LunarDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLunarDay(v int32)`

SetLunarDay sets LunarDay field to given value.

### HasLunarDay

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLunarDay() bool`

HasLunarDay returns a boolean if a field has been set.

### GetLunarDayName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarDayName() string`

GetLunarDayName returns the LunarDayName field if non-nil, zero value otherwise.

### GetLunarDayNameOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarDayNameOk() (*string, bool)`

GetLunarDayNameOk returns a tuple with the LunarDayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarDayName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLunarDayName(v string)`

SetLunarDayName sets LunarDayName field to given value.

### HasLunarDayName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLunarDayName() bool`

HasLunarDayName returns a boolean if a field has been set.

### GetLunarFestival

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarFestival() string`

GetLunarFestival returns the LunarFestival field if non-nil, zero value otherwise.

### GetLunarFestivalOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarFestivalOk() (*string, bool)`

GetLunarFestivalOk returns a tuple with the LunarFestival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarFestival

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLunarFestival(v string)`

SetLunarFestival sets LunarFestival field to given value.

### HasLunarFestival

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLunarFestival() bool`

HasLunarFestival returns a boolean if a field has been set.

### GetLunarMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarMonth() int32`

GetLunarMonth returns the LunarMonth field if non-nil, zero value otherwise.

### GetLunarMonthOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarMonthOk() (*int32, bool)`

GetLunarMonthOk returns a tuple with the LunarMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLunarMonth(v int32)`

SetLunarMonth sets LunarMonth field to given value.

### HasLunarMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLunarMonth() bool`

HasLunarMonth returns a boolean if a field has been set.

### GetLunarMonthName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarMonthName() string`

GetLunarMonthName returns the LunarMonthName field if non-nil, zero value otherwise.

### GetLunarMonthNameOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarMonthNameOk() (*string, bool)`

GetLunarMonthNameOk returns a tuple with the LunarMonthName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarMonthName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLunarMonthName(v string)`

SetLunarMonthName sets LunarMonthName field to given value.

### HasLunarMonthName

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLunarMonthName() bool`

HasLunarMonthName returns a boolean if a field has been set.

### GetLunarYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarYear() int32`

GetLunarYear returns the LunarYear field if non-nil, zero value otherwise.

### GetLunarYearOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetLunarYearOk() (*int32, bool)`

GetLunarYearOk returns a tuple with the LunarYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetLunarYear(v int32)`

SetLunarYear sets LunarYear field to given value.

### HasLunarYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasLunarYear() bool`

HasLunarYear returns a boolean if a field has been set.

### GetMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetSolarFestival

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetSolarFestival() string`

GetSolarFestival returns the SolarFestival field if non-nil, zero value otherwise.

### GetSolarFestivalOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetSolarFestivalOk() (*string, bool)`

GetSolarFestivalOk returns a tuple with the SolarFestival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarFestival

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetSolarFestival(v string)`

SetSolarFestival sets SolarFestival field to given value.

### HasSolarFestival

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasSolarFestival() bool`

HasSolarFestival returns a boolean if a field has been set.

### GetSolarTerm

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetSolarTerm() string`

GetSolarTerm returns the SolarTerm field if non-nil, zero value otherwise.

### GetSolarTermOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetSolarTermOk() (*string, bool)`

GetSolarTermOk returns a tuple with the SolarTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarTerm

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetSolarTerm(v string)`

SetSolarTerm sets SolarTerm field to given value.

### HasSolarTerm

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasSolarTerm() bool`

HasSolarTerm returns a boolean if a field has been set.

### GetWeekdayCn

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetWeekdayCn() string`

GetWeekdayCn returns the WeekdayCn field if non-nil, zero value otherwise.

### GetWeekdayCnOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetWeekdayCnOk() (*string, bool)`

GetWeekdayCnOk returns a tuple with the WeekdayCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayCn

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetWeekdayCn(v string)`

SetWeekdayCn sets WeekdayCn field to given value.

### HasWeekdayCn

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasWeekdayCn() bool`

HasWeekdayCn returns a boolean if a field has been set.

### GetYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *GetMiscHolidayCalendar200ResponseDaysInner) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



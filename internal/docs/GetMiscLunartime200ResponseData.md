# GetMiscLunartime200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryTimestamp** | Pointer to **string** | 原始 ts 入参。 | [optional] 
**QueryTimezone** | Pointer to **string** | 原始 timezone 入参。 | [optional] 
**Timezone** | Pointer to **string** | 解析后的时区。 | [optional] 
**Datetime** | Pointer to **string** | 本地化时间，格式 YYYY-MM-DD HH:mm:ss。 | [optional] 
**DatetimeRfc3339** | Pointer to **string** | RFC3339 时间格式。 | [optional] 
**TimestampUnix** | Pointer to **int32** | 秒级 Unix 时间戳。 | [optional] 
**Weekday** | Pointer to **string** | 星期英文。 | [optional] 
**WeekdayCn** | Pointer to **string** | 星期中文。 | [optional] 
**LunarYear** | Pointer to **int32** | 农历年份（数字）。 | [optional] 
**LunarMonth** | Pointer to **int32** | 农历月份（数字）。 | [optional] 
**LunarDay** | Pointer to **int32** | 农历日期（数字）。 | [optional] 
**IsLeapMonth** | Pointer to **bool** | 是否闰月。 | [optional] 
**LunarYearCn** | Pointer to **string** | 农历年份中文表示。 | [optional] 
**LunarMonthCn** | Pointer to **string** | 农历月份中文表示。 | [optional] 
**LunarDayCn** | Pointer to **string** | 农历日期中文表示。 | [optional] 
**GanzhiYear** | Pointer to **string** | 干支年。 | [optional] 
**GanzhiMonth** | Pointer to **string** | 干支月。 | [optional] 
**GanzhiDay** | Pointer to **string** | 干支日。 | [optional] 
**Zodiac** | Pointer to **string** | 生肖。 | [optional] 
**SolarTerm** | Pointer to **string** | 节气，无则为空字符串。 | [optional] 
**LunarFestivals** | Pointer to **[]string** | 农历节日数组。 | [optional] 
**SolarFestivals** | Pointer to **[]string** | 公历节日数组。 | [optional] 

## Methods

### NewGetMiscLunartime200ResponseData

`func NewGetMiscLunartime200ResponseData() *GetMiscLunartime200ResponseData`

NewGetMiscLunartime200ResponseData instantiates a new GetMiscLunartime200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscLunartime200ResponseDataWithDefaults

`func NewGetMiscLunartime200ResponseDataWithDefaults() *GetMiscLunartime200ResponseData`

NewGetMiscLunartime200ResponseDataWithDefaults instantiates a new GetMiscLunartime200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryTimestamp

`func (o *GetMiscLunartime200ResponseData) GetQueryTimestamp() string`

GetQueryTimestamp returns the QueryTimestamp field if non-nil, zero value otherwise.

### GetQueryTimestampOk

`func (o *GetMiscLunartime200ResponseData) GetQueryTimestampOk() (*string, bool)`

GetQueryTimestampOk returns a tuple with the QueryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTimestamp

`func (o *GetMiscLunartime200ResponseData) SetQueryTimestamp(v string)`

SetQueryTimestamp sets QueryTimestamp field to given value.

### HasQueryTimestamp

`func (o *GetMiscLunartime200ResponseData) HasQueryTimestamp() bool`

HasQueryTimestamp returns a boolean if a field has been set.

### GetQueryTimezone

`func (o *GetMiscLunartime200ResponseData) GetQueryTimezone() string`

GetQueryTimezone returns the QueryTimezone field if non-nil, zero value otherwise.

### GetQueryTimezoneOk

`func (o *GetMiscLunartime200ResponseData) GetQueryTimezoneOk() (*string, bool)`

GetQueryTimezoneOk returns a tuple with the QueryTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTimezone

`func (o *GetMiscLunartime200ResponseData) SetQueryTimezone(v string)`

SetQueryTimezone sets QueryTimezone field to given value.

### HasQueryTimezone

`func (o *GetMiscLunartime200ResponseData) HasQueryTimezone() bool`

HasQueryTimezone returns a boolean if a field has been set.

### GetTimezone

`func (o *GetMiscLunartime200ResponseData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetMiscLunartime200ResponseData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetMiscLunartime200ResponseData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetMiscLunartime200ResponseData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDatetime

`func (o *GetMiscLunartime200ResponseData) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetMiscLunartime200ResponseData) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetMiscLunartime200ResponseData) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetMiscLunartime200ResponseData) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetDatetimeRfc3339

`func (o *GetMiscLunartime200ResponseData) GetDatetimeRfc3339() string`

GetDatetimeRfc3339 returns the DatetimeRfc3339 field if non-nil, zero value otherwise.

### GetDatetimeRfc3339Ok

`func (o *GetMiscLunartime200ResponseData) GetDatetimeRfc3339Ok() (*string, bool)`

GetDatetimeRfc3339Ok returns a tuple with the DatetimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeRfc3339

`func (o *GetMiscLunartime200ResponseData) SetDatetimeRfc3339(v string)`

SetDatetimeRfc3339 sets DatetimeRfc3339 field to given value.

### HasDatetimeRfc3339

`func (o *GetMiscLunartime200ResponseData) HasDatetimeRfc3339() bool`

HasDatetimeRfc3339 returns a boolean if a field has been set.

### GetTimestampUnix

`func (o *GetMiscLunartime200ResponseData) GetTimestampUnix() int32`

GetTimestampUnix returns the TimestampUnix field if non-nil, zero value otherwise.

### GetTimestampUnixOk

`func (o *GetMiscLunartime200ResponseData) GetTimestampUnixOk() (*int32, bool)`

GetTimestampUnixOk returns a tuple with the TimestampUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUnix

`func (o *GetMiscLunartime200ResponseData) SetTimestampUnix(v int32)`

SetTimestampUnix sets TimestampUnix field to given value.

### HasTimestampUnix

`func (o *GetMiscLunartime200ResponseData) HasTimestampUnix() bool`

HasTimestampUnix returns a boolean if a field has been set.

### GetWeekday

`func (o *GetMiscLunartime200ResponseData) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *GetMiscLunartime200ResponseData) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *GetMiscLunartime200ResponseData) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *GetMiscLunartime200ResponseData) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.

### GetWeekdayCn

`func (o *GetMiscLunartime200ResponseData) GetWeekdayCn() string`

GetWeekdayCn returns the WeekdayCn field if non-nil, zero value otherwise.

### GetWeekdayCnOk

`func (o *GetMiscLunartime200ResponseData) GetWeekdayCnOk() (*string, bool)`

GetWeekdayCnOk returns a tuple with the WeekdayCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayCn

`func (o *GetMiscLunartime200ResponseData) SetWeekdayCn(v string)`

SetWeekdayCn sets WeekdayCn field to given value.

### HasWeekdayCn

`func (o *GetMiscLunartime200ResponseData) HasWeekdayCn() bool`

HasWeekdayCn returns a boolean if a field has been set.

### GetLunarYear

`func (o *GetMiscLunartime200ResponseData) GetLunarYear() int32`

GetLunarYear returns the LunarYear field if non-nil, zero value otherwise.

### GetLunarYearOk

`func (o *GetMiscLunartime200ResponseData) GetLunarYearOk() (*int32, bool)`

GetLunarYearOk returns a tuple with the LunarYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarYear

`func (o *GetMiscLunartime200ResponseData) SetLunarYear(v int32)`

SetLunarYear sets LunarYear field to given value.

### HasLunarYear

`func (o *GetMiscLunartime200ResponseData) HasLunarYear() bool`

HasLunarYear returns a boolean if a field has been set.

### GetLunarMonth

`func (o *GetMiscLunartime200ResponseData) GetLunarMonth() int32`

GetLunarMonth returns the LunarMonth field if non-nil, zero value otherwise.

### GetLunarMonthOk

`func (o *GetMiscLunartime200ResponseData) GetLunarMonthOk() (*int32, bool)`

GetLunarMonthOk returns a tuple with the LunarMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarMonth

`func (o *GetMiscLunartime200ResponseData) SetLunarMonth(v int32)`

SetLunarMonth sets LunarMonth field to given value.

### HasLunarMonth

`func (o *GetMiscLunartime200ResponseData) HasLunarMonth() bool`

HasLunarMonth returns a boolean if a field has been set.

### GetLunarDay

`func (o *GetMiscLunartime200ResponseData) GetLunarDay() int32`

GetLunarDay returns the LunarDay field if non-nil, zero value otherwise.

### GetLunarDayOk

`func (o *GetMiscLunartime200ResponseData) GetLunarDayOk() (*int32, bool)`

GetLunarDayOk returns a tuple with the LunarDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarDay

`func (o *GetMiscLunartime200ResponseData) SetLunarDay(v int32)`

SetLunarDay sets LunarDay field to given value.

### HasLunarDay

`func (o *GetMiscLunartime200ResponseData) HasLunarDay() bool`

HasLunarDay returns a boolean if a field has been set.

### GetIsLeapMonth

`func (o *GetMiscLunartime200ResponseData) GetIsLeapMonth() bool`

GetIsLeapMonth returns the IsLeapMonth field if non-nil, zero value otherwise.

### GetIsLeapMonthOk

`func (o *GetMiscLunartime200ResponseData) GetIsLeapMonthOk() (*bool, bool)`

GetIsLeapMonthOk returns a tuple with the IsLeapMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeapMonth

`func (o *GetMiscLunartime200ResponseData) SetIsLeapMonth(v bool)`

SetIsLeapMonth sets IsLeapMonth field to given value.

### HasIsLeapMonth

`func (o *GetMiscLunartime200ResponseData) HasIsLeapMonth() bool`

HasIsLeapMonth returns a boolean if a field has been set.

### GetLunarYearCn

`func (o *GetMiscLunartime200ResponseData) GetLunarYearCn() string`

GetLunarYearCn returns the LunarYearCn field if non-nil, zero value otherwise.

### GetLunarYearCnOk

`func (o *GetMiscLunartime200ResponseData) GetLunarYearCnOk() (*string, bool)`

GetLunarYearCnOk returns a tuple with the LunarYearCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarYearCn

`func (o *GetMiscLunartime200ResponseData) SetLunarYearCn(v string)`

SetLunarYearCn sets LunarYearCn field to given value.

### HasLunarYearCn

`func (o *GetMiscLunartime200ResponseData) HasLunarYearCn() bool`

HasLunarYearCn returns a boolean if a field has been set.

### GetLunarMonthCn

`func (o *GetMiscLunartime200ResponseData) GetLunarMonthCn() string`

GetLunarMonthCn returns the LunarMonthCn field if non-nil, zero value otherwise.

### GetLunarMonthCnOk

`func (o *GetMiscLunartime200ResponseData) GetLunarMonthCnOk() (*string, bool)`

GetLunarMonthCnOk returns a tuple with the LunarMonthCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarMonthCn

`func (o *GetMiscLunartime200ResponseData) SetLunarMonthCn(v string)`

SetLunarMonthCn sets LunarMonthCn field to given value.

### HasLunarMonthCn

`func (o *GetMiscLunartime200ResponseData) HasLunarMonthCn() bool`

HasLunarMonthCn returns a boolean if a field has been set.

### GetLunarDayCn

`func (o *GetMiscLunartime200ResponseData) GetLunarDayCn() string`

GetLunarDayCn returns the LunarDayCn field if non-nil, zero value otherwise.

### GetLunarDayCnOk

`func (o *GetMiscLunartime200ResponseData) GetLunarDayCnOk() (*string, bool)`

GetLunarDayCnOk returns a tuple with the LunarDayCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarDayCn

`func (o *GetMiscLunartime200ResponseData) SetLunarDayCn(v string)`

SetLunarDayCn sets LunarDayCn field to given value.

### HasLunarDayCn

`func (o *GetMiscLunartime200ResponseData) HasLunarDayCn() bool`

HasLunarDayCn returns a boolean if a field has been set.

### GetGanzhiYear

`func (o *GetMiscLunartime200ResponseData) GetGanzhiYear() string`

GetGanzhiYear returns the GanzhiYear field if non-nil, zero value otherwise.

### GetGanzhiYearOk

`func (o *GetMiscLunartime200ResponseData) GetGanzhiYearOk() (*string, bool)`

GetGanzhiYearOk returns a tuple with the GanzhiYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiYear

`func (o *GetMiscLunartime200ResponseData) SetGanzhiYear(v string)`

SetGanzhiYear sets GanzhiYear field to given value.

### HasGanzhiYear

`func (o *GetMiscLunartime200ResponseData) HasGanzhiYear() bool`

HasGanzhiYear returns a boolean if a field has been set.

### GetGanzhiMonth

`func (o *GetMiscLunartime200ResponseData) GetGanzhiMonth() string`

GetGanzhiMonth returns the GanzhiMonth field if non-nil, zero value otherwise.

### GetGanzhiMonthOk

`func (o *GetMiscLunartime200ResponseData) GetGanzhiMonthOk() (*string, bool)`

GetGanzhiMonthOk returns a tuple with the GanzhiMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiMonth

`func (o *GetMiscLunartime200ResponseData) SetGanzhiMonth(v string)`

SetGanzhiMonth sets GanzhiMonth field to given value.

### HasGanzhiMonth

`func (o *GetMiscLunartime200ResponseData) HasGanzhiMonth() bool`

HasGanzhiMonth returns a boolean if a field has been set.

### GetGanzhiDay

`func (o *GetMiscLunartime200ResponseData) GetGanzhiDay() string`

GetGanzhiDay returns the GanzhiDay field if non-nil, zero value otherwise.

### GetGanzhiDayOk

`func (o *GetMiscLunartime200ResponseData) GetGanzhiDayOk() (*string, bool)`

GetGanzhiDayOk returns a tuple with the GanzhiDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGanzhiDay

`func (o *GetMiscLunartime200ResponseData) SetGanzhiDay(v string)`

SetGanzhiDay sets GanzhiDay field to given value.

### HasGanzhiDay

`func (o *GetMiscLunartime200ResponseData) HasGanzhiDay() bool`

HasGanzhiDay returns a boolean if a field has been set.

### GetZodiac

`func (o *GetMiscLunartime200ResponseData) GetZodiac() string`

GetZodiac returns the Zodiac field if non-nil, zero value otherwise.

### GetZodiacOk

`func (o *GetMiscLunartime200ResponseData) GetZodiacOk() (*string, bool)`

GetZodiacOk returns a tuple with the Zodiac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZodiac

`func (o *GetMiscLunartime200ResponseData) SetZodiac(v string)`

SetZodiac sets Zodiac field to given value.

### HasZodiac

`func (o *GetMiscLunartime200ResponseData) HasZodiac() bool`

HasZodiac returns a boolean if a field has been set.

### GetSolarTerm

`func (o *GetMiscLunartime200ResponseData) GetSolarTerm() string`

GetSolarTerm returns the SolarTerm field if non-nil, zero value otherwise.

### GetSolarTermOk

`func (o *GetMiscLunartime200ResponseData) GetSolarTermOk() (*string, bool)`

GetSolarTermOk returns a tuple with the SolarTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarTerm

`func (o *GetMiscLunartime200ResponseData) SetSolarTerm(v string)`

SetSolarTerm sets SolarTerm field to given value.

### HasSolarTerm

`func (o *GetMiscLunartime200ResponseData) HasSolarTerm() bool`

HasSolarTerm returns a boolean if a field has been set.

### GetLunarFestivals

`func (o *GetMiscLunartime200ResponseData) GetLunarFestivals() []string`

GetLunarFestivals returns the LunarFestivals field if non-nil, zero value otherwise.

### GetLunarFestivalsOk

`func (o *GetMiscLunartime200ResponseData) GetLunarFestivalsOk() (*[]string, bool)`

GetLunarFestivalsOk returns a tuple with the LunarFestivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunarFestivals

`func (o *GetMiscLunartime200ResponseData) SetLunarFestivals(v []string)`

SetLunarFestivals sets LunarFestivals field to given value.

### HasLunarFestivals

`func (o *GetMiscLunartime200ResponseData) HasLunarFestivals() bool`

HasLunarFestivals returns a boolean if a field has been set.

### GetSolarFestivals

`func (o *GetMiscLunartime200ResponseData) GetSolarFestivals() []string`

GetSolarFestivals returns the SolarFestivals field if non-nil, zero value otherwise.

### GetSolarFestivalsOk

`func (o *GetMiscLunartime200ResponseData) GetSolarFestivalsOk() (*[]string, bool)`

GetSolarFestivalsOk returns a tuple with the SolarFestivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarFestivals

`func (o *GetMiscLunartime200ResponseData) SetSolarFestivals(v []string)`

SetSolarFestivals sets SolarFestivals field to given value.

### HasSolarFestivals

`func (o *GetMiscLunartime200ResponseData) HasSolarFestivals() bool`

HasSolarFestivals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



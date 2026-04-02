# GetMiscHolidayCalendar200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | 查询模式：day、month、year。 | [optional] 
**Query** | Pointer to [**GetMiscHolidayCalendar200ResponseQuery**](GetMiscHolidayCalendar200ResponseQuery.md) |  | [optional] 
**Summary** | Pointer to [**GetMiscHolidayCalendar200ResponseSummary**](GetMiscHolidayCalendar200ResponseSummary.md) |  | [optional] 
**Days** | Pointer to [**[]GetMiscHolidayCalendar200ResponseDaysInner**](GetMiscHolidayCalendar200ResponseDaysInner.md) | 日期明细列表。 | [optional] 
**Holidays** | Pointer to [**[]GetMiscHolidayCalendar200ResponseHolidaysInner**](GetMiscHolidayCalendar200ResponseHolidaysInner.md) | 节日事件列表。 | [optional] 
**Nearby** | Pointer to [**GetMiscHolidayCalendar200ResponseNearby**](GetMiscHolidayCalendar200ResponseNearby.md) |  | [optional] 

## Methods

### NewGetMiscHolidayCalendar200Response

`func NewGetMiscHolidayCalendar200Response() *GetMiscHolidayCalendar200Response`

NewGetMiscHolidayCalendar200Response instantiates a new GetMiscHolidayCalendar200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseWithDefaults

`func NewGetMiscHolidayCalendar200ResponseWithDefaults() *GetMiscHolidayCalendar200Response`

NewGetMiscHolidayCalendar200ResponseWithDefaults instantiates a new GetMiscHolidayCalendar200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetMiscHolidayCalendar200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetMiscHolidayCalendar200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetMiscHolidayCalendar200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetMiscHolidayCalendar200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetQuery

`func (o *GetMiscHolidayCalendar200Response) GetQuery() GetMiscHolidayCalendar200ResponseQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetMiscHolidayCalendar200Response) GetQueryOk() (*GetMiscHolidayCalendar200ResponseQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetMiscHolidayCalendar200Response) SetQuery(v GetMiscHolidayCalendar200ResponseQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GetMiscHolidayCalendar200Response) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSummary

`func (o *GetMiscHolidayCalendar200Response) GetSummary() GetMiscHolidayCalendar200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetMiscHolidayCalendar200Response) GetSummaryOk() (*GetMiscHolidayCalendar200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetMiscHolidayCalendar200Response) SetSummary(v GetMiscHolidayCalendar200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetMiscHolidayCalendar200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDays

`func (o *GetMiscHolidayCalendar200Response) GetDays() []GetMiscHolidayCalendar200ResponseDaysInner`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *GetMiscHolidayCalendar200Response) GetDaysOk() (*[]GetMiscHolidayCalendar200ResponseDaysInner, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *GetMiscHolidayCalendar200Response) SetDays(v []GetMiscHolidayCalendar200ResponseDaysInner)`

SetDays sets Days field to given value.

### HasDays

`func (o *GetMiscHolidayCalendar200Response) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetHolidays

`func (o *GetMiscHolidayCalendar200Response) GetHolidays() []GetMiscHolidayCalendar200ResponseHolidaysInner`

GetHolidays returns the Holidays field if non-nil, zero value otherwise.

### GetHolidaysOk

`func (o *GetMiscHolidayCalendar200Response) GetHolidaysOk() (*[]GetMiscHolidayCalendar200ResponseHolidaysInner, bool)`

GetHolidaysOk returns a tuple with the Holidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidays

`func (o *GetMiscHolidayCalendar200Response) SetHolidays(v []GetMiscHolidayCalendar200ResponseHolidaysInner)`

SetHolidays sets Holidays field to given value.

### HasHolidays

`func (o *GetMiscHolidayCalendar200Response) HasHolidays() bool`

HasHolidays returns a boolean if a field has been set.

### GetNearby

`func (o *GetMiscHolidayCalendar200Response) GetNearby() GetMiscHolidayCalendar200ResponseNearby`

GetNearby returns the Nearby field if non-nil, zero value otherwise.

### GetNearbyOk

`func (o *GetMiscHolidayCalendar200Response) GetNearbyOk() (*GetMiscHolidayCalendar200ResponseNearby, bool)`

GetNearbyOk returns a tuple with the Nearby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearby

`func (o *GetMiscHolidayCalendar200Response) SetNearby(v GetMiscHolidayCalendar200ResponseNearby)`

SetNearby sets Nearby field to given value.

### HasNearby

`func (o *GetMiscHolidayCalendar200Response) HasNearby() bool`

HasNearby returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



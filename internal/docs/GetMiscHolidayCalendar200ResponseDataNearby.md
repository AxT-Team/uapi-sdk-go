# GetMiscHolidayCalendar200ResponseDataNearby

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previous** | Pointer to [**[]GetMiscHolidayCalendar200ResponseDataNearbyPreviousInner**](GetMiscHolidayCalendar200ResponseDataNearbyPreviousInner.md) | 当前查询日期之前最近的节日列表（按时间倒序）。 | [optional] 
**Next** | Pointer to [**[]GetMiscHolidayCalendar200ResponseDataNearbyNextInner**](GetMiscHolidayCalendar200ResponseDataNearbyNextInner.md) | 当前查询日期之后最近的节日列表（按时间正序）。 | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseDataNearby

`func NewGetMiscHolidayCalendar200ResponseDataNearby() *GetMiscHolidayCalendar200ResponseDataNearby`

NewGetMiscHolidayCalendar200ResponseDataNearby instantiates a new GetMiscHolidayCalendar200ResponseDataNearby object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseDataNearbyWithDefaults

`func NewGetMiscHolidayCalendar200ResponseDataNearbyWithDefaults() *GetMiscHolidayCalendar200ResponseDataNearby`

NewGetMiscHolidayCalendar200ResponseDataNearbyWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseDataNearby object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevious

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) GetPrevious() []GetMiscHolidayCalendar200ResponseDataNearbyPreviousInner`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) GetPreviousOk() (*[]GetMiscHolidayCalendar200ResponseDataNearbyPreviousInner, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) SetPrevious(v []GetMiscHolidayCalendar200ResponseDataNearbyPreviousInner)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetNext

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) GetNext() []GetMiscHolidayCalendar200ResponseDataNearbyNextInner`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) GetNextOk() (*[]GetMiscHolidayCalendar200ResponseDataNearbyNextInner, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) SetNext(v []GetMiscHolidayCalendar200ResponseDataNearbyNextInner)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetMiscHolidayCalendar200ResponseDataNearby) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



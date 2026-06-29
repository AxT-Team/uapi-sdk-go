# GetMiscHolidayCalendar200ResponseNearby

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**[]GetMiscHolidayCalendar200ResponseNearbyNextInner**](GetMiscHolidayCalendar200ResponseNearbyNextInner.md) | 当前查询日期之后最近的节日列表（按时间正序）。 | [optional] 
**Previous** | Pointer to [**[]GetMiscHolidayCalendar200ResponseNearbyPreviousInner**](GetMiscHolidayCalendar200ResponseNearbyPreviousInner.md) | 当前查询日期之前最近的节日列表（按时间倒序）。 | [optional] 

## Methods

### NewGetMiscHolidayCalendar200ResponseNearby

`func NewGetMiscHolidayCalendar200ResponseNearby() *GetMiscHolidayCalendar200ResponseNearby`

NewGetMiscHolidayCalendar200ResponseNearby instantiates a new GetMiscHolidayCalendar200ResponseNearby object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHolidayCalendar200ResponseNearbyWithDefaults

`func NewGetMiscHolidayCalendar200ResponseNearbyWithDefaults() *GetMiscHolidayCalendar200ResponseNearby`

NewGetMiscHolidayCalendar200ResponseNearbyWithDefaults instantiates a new GetMiscHolidayCalendar200ResponseNearby object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *GetMiscHolidayCalendar200ResponseNearby) GetNext() []GetMiscHolidayCalendar200ResponseNearbyNextInner`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetMiscHolidayCalendar200ResponseNearby) GetNextOk() (*[]GetMiscHolidayCalendar200ResponseNearbyNextInner, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetMiscHolidayCalendar200ResponseNearby) SetNext(v []GetMiscHolidayCalendar200ResponseNearbyNextInner)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetMiscHolidayCalendar200ResponseNearby) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *GetMiscHolidayCalendar200ResponseNearby) GetPrevious() []GetMiscHolidayCalendar200ResponseNearbyPreviousInner`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetMiscHolidayCalendar200ResponseNearby) GetPreviousOk() (*[]GetMiscHolidayCalendar200ResponseNearbyPreviousInner, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetMiscHolidayCalendar200ResponseNearby) SetPrevious(v []GetMiscHolidayCalendar200ResponseNearbyPreviousInner)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *GetMiscHolidayCalendar200ResponseNearby) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



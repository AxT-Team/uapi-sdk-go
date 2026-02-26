# GetMiscWeather200ResponseMinutelyPrecipDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **string** | 预报时间 ISO8601 | [optional] 
**Precip** | Pointer to **float32** | 5分钟累计降水量 mm | [optional] 
**Type** | Pointer to **string** | 降水类型：rain / snow | [optional] 

## Methods

### NewGetMiscWeather200ResponseMinutelyPrecipDataInner

`func NewGetMiscWeather200ResponseMinutelyPrecipDataInner() *GetMiscWeather200ResponseMinutelyPrecipDataInner`

NewGetMiscWeather200ResponseMinutelyPrecipDataInner instantiates a new GetMiscWeather200ResponseMinutelyPrecipDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscWeather200ResponseMinutelyPrecipDataInnerWithDefaults

`func NewGetMiscWeather200ResponseMinutelyPrecipDataInnerWithDefaults() *GetMiscWeather200ResponseMinutelyPrecipDataInner`

NewGetMiscWeather200ResponseMinutelyPrecipDataInnerWithDefaults instantiates a new GetMiscWeather200ResponseMinutelyPrecipDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetPrecip

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) GetPrecip() float32`

GetPrecip returns the Precip field if non-nil, zero value otherwise.

### GetPrecipOk

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) GetPrecipOk() (*float32, bool)`

GetPrecipOk returns a tuple with the Precip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecip

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) SetPrecip(v float32)`

SetPrecip sets Precip field to given value.

### HasPrecip

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) HasPrecip() bool`

HasPrecip returns a boolean if a field has been set.

### GetType

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMiscWeather200ResponseMinutelyPrecipDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



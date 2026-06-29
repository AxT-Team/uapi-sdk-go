# DailyRecommendMoment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTime** | Pointer to **string** | 仅 moment 模式返回，服务器当前时间，ISO 8601 格式。 | [optional] 
**Date** | Pointer to **string** | 仅 daily 模式返回，对应日期，格式 YYYY-MM-DD。 | [optional] 
**Item** | Pointer to [**DailyRecommendMomentItem**](DailyRecommendMomentItem.md) |  | [optional] 
**Mode** | Pointer to **string** | 当前运行模式。 | [optional] 
**Scene** | Pointer to [**DailyRecommendMomentScene**](DailyRecommendMomentScene.md) |  | [optional] 
**Seed** | Pointer to **string** | 当次结果的确定性种子。 | [optional] 
**TimeSegment** | Pointer to **string** | 仅 moment 模式返回，命中的时段标识。 | [optional] 

## Methods

### NewDailyRecommendMoment

`func NewDailyRecommendMoment() *DailyRecommendMoment`

NewDailyRecommendMoment instantiates a new DailyRecommendMoment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyRecommendMomentWithDefaults

`func NewDailyRecommendMomentWithDefaults() *DailyRecommendMoment`

NewDailyRecommendMomentWithDefaults instantiates a new DailyRecommendMoment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentTime

`func (o *DailyRecommendMoment) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *DailyRecommendMoment) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *DailyRecommendMoment) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *DailyRecommendMoment) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetDate

`func (o *DailyRecommendMoment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DailyRecommendMoment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DailyRecommendMoment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *DailyRecommendMoment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetItem

`func (o *DailyRecommendMoment) GetItem() DailyRecommendMomentItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *DailyRecommendMoment) GetItemOk() (*DailyRecommendMomentItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *DailyRecommendMoment) SetItem(v DailyRecommendMomentItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *DailyRecommendMoment) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetMode

`func (o *DailyRecommendMoment) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DailyRecommendMoment) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DailyRecommendMoment) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DailyRecommendMoment) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetScene

`func (o *DailyRecommendMoment) GetScene() DailyRecommendMomentScene`

GetScene returns the Scene field if non-nil, zero value otherwise.

### GetSceneOk

`func (o *DailyRecommendMoment) GetSceneOk() (*DailyRecommendMomentScene, bool)`

GetSceneOk returns a tuple with the Scene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScene

`func (o *DailyRecommendMoment) SetScene(v DailyRecommendMomentScene)`

SetScene sets Scene field to given value.

### HasScene

`func (o *DailyRecommendMoment) HasScene() bool`

HasScene returns a boolean if a field has been set.

### GetSeed

`func (o *DailyRecommendMoment) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *DailyRecommendMoment) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *DailyRecommendMoment) SetSeed(v string)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *DailyRecommendMoment) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetTimeSegment

`func (o *DailyRecommendMoment) GetTimeSegment() string`

GetTimeSegment returns the TimeSegment field if non-nil, zero value otherwise.

### GetTimeSegmentOk

`func (o *DailyRecommendMoment) GetTimeSegmentOk() (*string, bool)`

GetTimeSegmentOk returns a tuple with the TimeSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSegment

`func (o *DailyRecommendMoment) SetTimeSegment(v string)`

SetTimeSegment sets TimeSegment field to given value.

### HasTimeSegment

`func (o *DailyRecommendMoment) HasTimeSegment() bool`

HasTimeSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



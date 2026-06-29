# GetMiscHotboard200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetMiscHotboard200ResponseOneOfListInner**](GetMiscHotboard200ResponseOneOfListInner.md) | 热榜条目列表。 | [optional] 
**SnapshotTime** | Pointer to **int32** | 时光机模式返回的快照实际时间戳（毫秒）。当前热榜模式下通常不返回。 | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** | 热榜更新时间。时光机模式下对应返回快照的更新时间。 | [optional] 

## Methods

### NewGetMiscHotboard200ResponseOneOf

`func NewGetMiscHotboard200ResponseOneOf() *GetMiscHotboard200ResponseOneOf`

NewGetMiscHotboard200ResponseOneOf instantiates a new GetMiscHotboard200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHotboard200ResponseOneOfWithDefaults

`func NewGetMiscHotboard200ResponseOneOfWithDefaults() *GetMiscHotboard200ResponseOneOf`

NewGetMiscHotboard200ResponseOneOfWithDefaults instantiates a new GetMiscHotboard200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetMiscHotboard200ResponseOneOf) GetList() []GetMiscHotboard200ResponseOneOfListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetMiscHotboard200ResponseOneOf) GetListOk() (*[]GetMiscHotboard200ResponseOneOfListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetMiscHotboard200ResponseOneOf) SetList(v []GetMiscHotboard200ResponseOneOfListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetMiscHotboard200ResponseOneOf) HasList() bool`

HasList returns a boolean if a field has been set.

### GetSnapshotTime

`func (o *GetMiscHotboard200ResponseOneOf) GetSnapshotTime() int32`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *GetMiscHotboard200ResponseOneOf) GetSnapshotTimeOk() (*int32, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *GetMiscHotboard200ResponseOneOf) SetSnapshotTime(v int32)`

SetSnapshotTime sets SnapshotTime field to given value.

### HasSnapshotTime

`func (o *GetMiscHotboard200ResponseOneOf) HasSnapshotTime() bool`

HasSnapshotTime returns a boolean if a field has been set.

### GetType

`func (o *GetMiscHotboard200ResponseOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMiscHotboard200ResponseOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMiscHotboard200ResponseOneOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMiscHotboard200ResponseOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetMiscHotboard200ResponseOneOf) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetMiscHotboard200ResponseOneOf) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetMiscHotboard200ResponseOneOf) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetMiscHotboard200ResponseOneOf) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



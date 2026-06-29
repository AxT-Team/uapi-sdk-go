# GetMiscTrackingQuery200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierCode** | Pointer to **string** | 快递公司编码 | [optional] 
**CarrierName** | Pointer to **string** | 快递公司名称 | [optional] 
**CompletedAt** | Pointer to **string** | 完成时间。仅已完成时返回签收或妥投对应的轨迹时间；未完成时为空字符串。 | [optional] 
**IsCompleted** | Pointer to **bool** | 快递是否已完成。仅当状态识别为已签收/已妥投/已完成时为 true。 | [optional] 
**Status** | Pointer to **string** | 快递状态中文名称，例如：待揽收、已揽收、运输中、派送中、已完成、异常、未知。 | [optional] 
**StatusCode** | Pointer to **string** | 快递状态编码。可能值：pending、picked_up、in_transit、out_for_delivery、delivered、exception、unknown。 | [optional] 
**TrackCount** | Pointer to **int32** | 物流轨迹数量 | [optional] 
**TrackingNumber** | Pointer to **string** | 快递单号 | [optional] 
**Tracks** | Pointer to [**[]GetMiscTrackingQuery200ResponseTracksInner**](GetMiscTrackingQuery200ResponseTracksInner.md) | 物流轨迹列表，按时间倒序排列 | [optional] 

## Methods

### NewGetMiscTrackingQuery200Response

`func NewGetMiscTrackingQuery200Response() *GetMiscTrackingQuery200Response`

NewGetMiscTrackingQuery200Response instantiates a new GetMiscTrackingQuery200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscTrackingQuery200ResponseWithDefaults

`func NewGetMiscTrackingQuery200ResponseWithDefaults() *GetMiscTrackingQuery200Response`

NewGetMiscTrackingQuery200ResponseWithDefaults instantiates a new GetMiscTrackingQuery200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierCode

`func (o *GetMiscTrackingQuery200Response) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *GetMiscTrackingQuery200Response) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *GetMiscTrackingQuery200Response) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *GetMiscTrackingQuery200Response) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *GetMiscTrackingQuery200Response) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *GetMiscTrackingQuery200Response) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *GetMiscTrackingQuery200Response) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *GetMiscTrackingQuery200Response) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetMiscTrackingQuery200Response) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetMiscTrackingQuery200Response) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetMiscTrackingQuery200Response) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetMiscTrackingQuery200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetIsCompleted

`func (o *GetMiscTrackingQuery200Response) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *GetMiscTrackingQuery200Response) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *GetMiscTrackingQuery200Response) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *GetMiscTrackingQuery200Response) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### GetStatus

`func (o *GetMiscTrackingQuery200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMiscTrackingQuery200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMiscTrackingQuery200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMiscTrackingQuery200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *GetMiscTrackingQuery200Response) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetMiscTrackingQuery200Response) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetMiscTrackingQuery200Response) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetMiscTrackingQuery200Response) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTrackCount

`func (o *GetMiscTrackingQuery200Response) GetTrackCount() int32`

GetTrackCount returns the TrackCount field if non-nil, zero value otherwise.

### GetTrackCountOk

`func (o *GetMiscTrackingQuery200Response) GetTrackCountOk() (*int32, bool)`

GetTrackCountOk returns a tuple with the TrackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackCount

`func (o *GetMiscTrackingQuery200Response) SetTrackCount(v int32)`

SetTrackCount sets TrackCount field to given value.

### HasTrackCount

`func (o *GetMiscTrackingQuery200Response) HasTrackCount() bool`

HasTrackCount returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *GetMiscTrackingQuery200Response) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GetMiscTrackingQuery200Response) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GetMiscTrackingQuery200Response) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GetMiscTrackingQuery200Response) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetTracks

`func (o *GetMiscTrackingQuery200Response) GetTracks() []GetMiscTrackingQuery200ResponseTracksInner`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *GetMiscTrackingQuery200Response) GetTracksOk() (*[]GetMiscTrackingQuery200ResponseTracksInner, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *GetMiscTrackingQuery200Response) SetTracks(v []GetMiscTrackingQuery200ResponseTracksInner)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *GetMiscTrackingQuery200Response) HasTracks() bool`

HasTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



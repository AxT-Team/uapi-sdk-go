# GetMiscTrackingQuery200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingNumber** | Pointer to **string** | 快递单号 | [optional] 
**CarrierCode** | Pointer to **string** | 快递公司编码 | [optional] 
**CarrierName** | Pointer to **string** | 快递公司名称 | [optional] 
**TrackCount** | Pointer to **int32** | 物流轨迹数量 | [optional] 
**Tracks** | Pointer to [**[]GetMiscTrackingQuery200ResponseDataTracksInner**](GetMiscTrackingQuery200ResponseDataTracksInner.md) | 物流轨迹列表，按时间倒序排列 | [optional] 

## Methods

### NewGetMiscTrackingQuery200ResponseData

`func NewGetMiscTrackingQuery200ResponseData() *GetMiscTrackingQuery200ResponseData`

NewGetMiscTrackingQuery200ResponseData instantiates a new GetMiscTrackingQuery200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscTrackingQuery200ResponseDataWithDefaults

`func NewGetMiscTrackingQuery200ResponseDataWithDefaults() *GetMiscTrackingQuery200ResponseData`

NewGetMiscTrackingQuery200ResponseDataWithDefaults instantiates a new GetMiscTrackingQuery200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingNumber

`func (o *GetMiscTrackingQuery200ResponseData) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GetMiscTrackingQuery200ResponseData) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GetMiscTrackingQuery200ResponseData) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GetMiscTrackingQuery200ResponseData) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetCarrierCode

`func (o *GetMiscTrackingQuery200ResponseData) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *GetMiscTrackingQuery200ResponseData) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *GetMiscTrackingQuery200ResponseData) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *GetMiscTrackingQuery200ResponseData) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *GetMiscTrackingQuery200ResponseData) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *GetMiscTrackingQuery200ResponseData) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *GetMiscTrackingQuery200ResponseData) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *GetMiscTrackingQuery200ResponseData) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetTrackCount

`func (o *GetMiscTrackingQuery200ResponseData) GetTrackCount() int32`

GetTrackCount returns the TrackCount field if non-nil, zero value otherwise.

### GetTrackCountOk

`func (o *GetMiscTrackingQuery200ResponseData) GetTrackCountOk() (*int32, bool)`

GetTrackCountOk returns a tuple with the TrackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackCount

`func (o *GetMiscTrackingQuery200ResponseData) SetTrackCount(v int32)`

SetTrackCount sets TrackCount field to given value.

### HasTrackCount

`func (o *GetMiscTrackingQuery200ResponseData) HasTrackCount() bool`

HasTrackCount returns a boolean if a field has been set.

### GetTracks

`func (o *GetMiscTrackingQuery200ResponseData) GetTracks() []GetMiscTrackingQuery200ResponseDataTracksInner`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *GetMiscTrackingQuery200ResponseData) GetTracksOk() (*[]GetMiscTrackingQuery200ResponseDataTracksInner, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *GetMiscTrackingQuery200ResponseData) SetTracks(v []GetMiscTrackingQuery200ResponseDataTracksInner)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *GetMiscTrackingQuery200ResponseData) HasTracks() bool`

HasTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



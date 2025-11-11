# GetMiscTrackingCarriers200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | 支持的快递公司总数 | [optional] 
**Carriers** | Pointer to [**[]GetMiscTrackingCarriers200ResponseDataCarriersInner**](GetMiscTrackingCarriers200ResponseDataCarriersInner.md) | 快递公司列表 | [optional] 

## Methods

### NewGetMiscTrackingCarriers200ResponseData

`func NewGetMiscTrackingCarriers200ResponseData() *GetMiscTrackingCarriers200ResponseData`

NewGetMiscTrackingCarriers200ResponseData instantiates a new GetMiscTrackingCarriers200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscTrackingCarriers200ResponseDataWithDefaults

`func NewGetMiscTrackingCarriers200ResponseDataWithDefaults() *GetMiscTrackingCarriers200ResponseData`

NewGetMiscTrackingCarriers200ResponseDataWithDefaults instantiates a new GetMiscTrackingCarriers200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetMiscTrackingCarriers200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMiscTrackingCarriers200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMiscTrackingCarriers200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMiscTrackingCarriers200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCarriers

`func (o *GetMiscTrackingCarriers200ResponseData) GetCarriers() []GetMiscTrackingCarriers200ResponseDataCarriersInner`

GetCarriers returns the Carriers field if non-nil, zero value otherwise.

### GetCarriersOk

`func (o *GetMiscTrackingCarriers200ResponseData) GetCarriersOk() (*[]GetMiscTrackingCarriers200ResponseDataCarriersInner, bool)`

GetCarriersOk returns a tuple with the Carriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriers

`func (o *GetMiscTrackingCarriers200ResponseData) SetCarriers(v []GetMiscTrackingCarriers200ResponseDataCarriersInner)`

SetCarriers sets Carriers field to given value.

### HasCarriers

`func (o *GetMiscTrackingCarriers200ResponseData) HasCarriers() bool`

HasCarriers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



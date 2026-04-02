# GetMiscTrackingCarriers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carriers** | Pointer to [**[]GetMiscTrackingCarriers200ResponseCarriersInner**](GetMiscTrackingCarriers200ResponseCarriersInner.md) | 快递公司列表 | [optional] 
**Total** | Pointer to **int32** | 支持的快递公司总数 | [optional] 

## Methods

### NewGetMiscTrackingCarriers200Response

`func NewGetMiscTrackingCarriers200Response() *GetMiscTrackingCarriers200Response`

NewGetMiscTrackingCarriers200Response instantiates a new GetMiscTrackingCarriers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscTrackingCarriers200ResponseWithDefaults

`func NewGetMiscTrackingCarriers200ResponseWithDefaults() *GetMiscTrackingCarriers200Response`

NewGetMiscTrackingCarriers200ResponseWithDefaults instantiates a new GetMiscTrackingCarriers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarriers

`func (o *GetMiscTrackingCarriers200Response) GetCarriers() []GetMiscTrackingCarriers200ResponseCarriersInner`

GetCarriers returns the Carriers field if non-nil, zero value otherwise.

### GetCarriersOk

`func (o *GetMiscTrackingCarriers200Response) GetCarriersOk() (*[]GetMiscTrackingCarriers200ResponseCarriersInner, bool)`

GetCarriersOk returns a tuple with the Carriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriers

`func (o *GetMiscTrackingCarriers200Response) SetCarriers(v []GetMiscTrackingCarriers200ResponseCarriersInner)`

SetCarriers sets Carriers field to given value.

### HasCarriers

`func (o *GetMiscTrackingCarriers200Response) HasCarriers() bool`

HasCarriers returns a boolean if a field has been set.

### GetTotal

`func (o *GetMiscTrackingCarriers200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMiscTrackingCarriers200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMiscTrackingCarriers200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMiscTrackingCarriers200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



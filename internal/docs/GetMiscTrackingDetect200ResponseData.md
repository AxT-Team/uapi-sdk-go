# GetMiscTrackingDetect200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingNumber** | Pointer to **string** | 查询的快递单号 | [optional] 
**CarrierCode** | Pointer to **string** | 最可能的快递公司编码 | [optional] 
**CarrierName** | Pointer to **string** | 最可能的快递公司名称 | [optional] 
**Alternatives** | Pointer to [**[]GetMiscTrackingDetect200ResponseDataAlternativesInner**](GetMiscTrackingDetect200ResponseDataAlternativesInner.md) | 其他可能的快递公司列表（如果存在） | [optional] 

## Methods

### NewGetMiscTrackingDetect200ResponseData

`func NewGetMiscTrackingDetect200ResponseData() *GetMiscTrackingDetect200ResponseData`

NewGetMiscTrackingDetect200ResponseData instantiates a new GetMiscTrackingDetect200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscTrackingDetect200ResponseDataWithDefaults

`func NewGetMiscTrackingDetect200ResponseDataWithDefaults() *GetMiscTrackingDetect200ResponseData`

NewGetMiscTrackingDetect200ResponseDataWithDefaults instantiates a new GetMiscTrackingDetect200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingNumber

`func (o *GetMiscTrackingDetect200ResponseData) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GetMiscTrackingDetect200ResponseData) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GetMiscTrackingDetect200ResponseData) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GetMiscTrackingDetect200ResponseData) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetCarrierCode

`func (o *GetMiscTrackingDetect200ResponseData) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *GetMiscTrackingDetect200ResponseData) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *GetMiscTrackingDetect200ResponseData) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *GetMiscTrackingDetect200ResponseData) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *GetMiscTrackingDetect200ResponseData) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *GetMiscTrackingDetect200ResponseData) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *GetMiscTrackingDetect200ResponseData) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *GetMiscTrackingDetect200ResponseData) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetAlternatives

`func (o *GetMiscTrackingDetect200ResponseData) GetAlternatives() []GetMiscTrackingDetect200ResponseDataAlternativesInner`

GetAlternatives returns the Alternatives field if non-nil, zero value otherwise.

### GetAlternativesOk

`func (o *GetMiscTrackingDetect200ResponseData) GetAlternativesOk() (*[]GetMiscTrackingDetect200ResponseDataAlternativesInner, bool)`

GetAlternativesOk returns a tuple with the Alternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatives

`func (o *GetMiscTrackingDetect200ResponseData) SetAlternatives(v []GetMiscTrackingDetect200ResponseDataAlternativesInner)`

SetAlternatives sets Alternatives field to given value.

### HasAlternatives

`func (o *GetMiscTrackingDetect200ResponseData) HasAlternatives() bool`

HasAlternatives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



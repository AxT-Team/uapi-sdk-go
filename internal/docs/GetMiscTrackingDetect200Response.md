# GetMiscTrackingDetect200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternatives** | Pointer to [**[]GetMiscTrackingDetect200ResponseAlternativesInner**](GetMiscTrackingDetect200ResponseAlternativesInner.md) | 其他可能的快递公司列表。如果没有备选项，会返回空数组。 | [optional] 
**CarrierCode** | Pointer to **string** | 识别出的快递公司编码 | [optional] 
**CarrierName** | Pointer to **string** | 识别出的快递公司名称 | [optional] 
**TrackingNumber** | Pointer to **string** | 查询的快递单号 | [optional] 

## Methods

### NewGetMiscTrackingDetect200Response

`func NewGetMiscTrackingDetect200Response() *GetMiscTrackingDetect200Response`

NewGetMiscTrackingDetect200Response instantiates a new GetMiscTrackingDetect200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscTrackingDetect200ResponseWithDefaults

`func NewGetMiscTrackingDetect200ResponseWithDefaults() *GetMiscTrackingDetect200Response`

NewGetMiscTrackingDetect200ResponseWithDefaults instantiates a new GetMiscTrackingDetect200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatives

`func (o *GetMiscTrackingDetect200Response) GetAlternatives() []GetMiscTrackingDetect200ResponseAlternativesInner`

GetAlternatives returns the Alternatives field if non-nil, zero value otherwise.

### GetAlternativesOk

`func (o *GetMiscTrackingDetect200Response) GetAlternativesOk() (*[]GetMiscTrackingDetect200ResponseAlternativesInner, bool)`

GetAlternativesOk returns a tuple with the Alternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatives

`func (o *GetMiscTrackingDetect200Response) SetAlternatives(v []GetMiscTrackingDetect200ResponseAlternativesInner)`

SetAlternatives sets Alternatives field to given value.

### HasAlternatives

`func (o *GetMiscTrackingDetect200Response) HasAlternatives() bool`

HasAlternatives returns a boolean if a field has been set.

### GetCarrierCode

`func (o *GetMiscTrackingDetect200Response) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *GetMiscTrackingDetect200Response) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *GetMiscTrackingDetect200Response) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *GetMiscTrackingDetect200Response) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *GetMiscTrackingDetect200Response) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *GetMiscTrackingDetect200Response) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *GetMiscTrackingDetect200Response) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *GetMiscTrackingDetect200Response) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *GetMiscTrackingDetect200Response) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GetMiscTrackingDetect200Response) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GetMiscTrackingDetect200Response) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GetMiscTrackingDetect200Response) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



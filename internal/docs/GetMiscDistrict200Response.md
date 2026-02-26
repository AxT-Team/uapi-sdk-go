# GetMiscDistrict200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | 结果总数。 | [optional] 
**Results** | Pointer to [**[]GetMiscDistrict200ResponseResultsInner**](GetMiscDistrict200ResponseResultsInner.md) | 结果列表。 | [optional] 

## Methods

### NewGetMiscDistrict200Response

`func NewGetMiscDistrict200Response() *GetMiscDistrict200Response`

NewGetMiscDistrict200Response instantiates a new GetMiscDistrict200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscDistrict200ResponseWithDefaults

`func NewGetMiscDistrict200ResponseWithDefaults() *GetMiscDistrict200Response`

NewGetMiscDistrict200ResponseWithDefaults instantiates a new GetMiscDistrict200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetMiscDistrict200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMiscDistrict200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMiscDistrict200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMiscDistrict200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResults

`func (o *GetMiscDistrict200Response) GetResults() []GetMiscDistrict200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetMiscDistrict200Response) GetResultsOk() (*[]GetMiscDistrict200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetMiscDistrict200Response) SetResults(v []GetMiscDistrict200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetMiscDistrict200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetStatusUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]GetStatusUsage200ResponseAnyOfEndpointsInner**](GetStatusUsage200ResponseAnyOfEndpointsInner.md) |  | [optional] 
**Unaggregated** | Pointer to [**GetStatusUsage200ResponseAnyOfUnaggregated**](GetStatusUsage200ResponseAnyOfUnaggregated.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetStatusUsage200Response

`func NewGetStatusUsage200Response() *GetStatusUsage200Response`

NewGetStatusUsage200Response instantiates a new GetStatusUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusUsage200ResponseWithDefaults

`func NewGetStatusUsage200ResponseWithDefaults() *GetStatusUsage200Response`

NewGetStatusUsage200ResponseWithDefaults instantiates a new GetStatusUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *GetStatusUsage200Response) GetEndpoints() []GetStatusUsage200ResponseAnyOfEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *GetStatusUsage200Response) GetEndpointsOk() (*[]GetStatusUsage200ResponseAnyOfEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *GetStatusUsage200Response) SetEndpoints(v []GetStatusUsage200ResponseAnyOfEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *GetStatusUsage200Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetUnaggregated

`func (o *GetStatusUsage200Response) GetUnaggregated() GetStatusUsage200ResponseAnyOfUnaggregated`

GetUnaggregated returns the Unaggregated field if non-nil, zero value otherwise.

### GetUnaggregatedOk

`func (o *GetStatusUsage200Response) GetUnaggregatedOk() (*GetStatusUsage200ResponseAnyOfUnaggregated, bool)`

GetUnaggregatedOk returns a tuple with the Unaggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnaggregated

`func (o *GetStatusUsage200Response) SetUnaggregated(v GetStatusUsage200ResponseAnyOfUnaggregated)`

SetUnaggregated sets Unaggregated field to given value.

### HasUnaggregated

`func (o *GetStatusUsage200Response) HasUnaggregated() bool`

HasUnaggregated returns a boolean if a field has been set.

### GetPath

`func (o *GetStatusUsage200Response) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetStatusUsage200Response) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetStatusUsage200Response) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GetStatusUsage200Response) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCount

`func (o *GetStatusUsage200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetStatusUsage200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetStatusUsage200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetStatusUsage200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



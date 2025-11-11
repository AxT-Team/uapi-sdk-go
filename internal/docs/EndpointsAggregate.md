# EndpointsAggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]EndpointsAggregateEndpointsInner**](EndpointsAggregateEndpointsInner.md) |  | [optional] 
**Unaggregated** | Pointer to [**EndpointsAggregateUnaggregated**](EndpointsAggregateUnaggregated.md) |  | [optional] 

## Methods

### NewEndpointsAggregate

`func NewEndpointsAggregate() *EndpointsAggregate`

NewEndpointsAggregate instantiates a new EndpointsAggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointsAggregateWithDefaults

`func NewEndpointsAggregateWithDefaults() *EndpointsAggregate`

NewEndpointsAggregateWithDefaults instantiates a new EndpointsAggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *EndpointsAggregate) GetEndpoints() []EndpointsAggregateEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *EndpointsAggregate) GetEndpointsOk() (*[]EndpointsAggregateEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *EndpointsAggregate) SetEndpoints(v []EndpointsAggregateEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *EndpointsAggregate) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetUnaggregated

`func (o *EndpointsAggregate) GetUnaggregated() EndpointsAggregateUnaggregated`

GetUnaggregated returns the Unaggregated field if non-nil, zero value otherwise.

### GetUnaggregatedOk

`func (o *EndpointsAggregate) GetUnaggregatedOk() (*EndpointsAggregateUnaggregated, bool)`

GetUnaggregatedOk returns a tuple with the Unaggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnaggregated

`func (o *EndpointsAggregate) SetUnaggregated(v EndpointsAggregateUnaggregated)`

SetUnaggregated sets Unaggregated field to given value.

### HasUnaggregated

`func (o *EndpointsAggregate) HasUnaggregated() bool`

HasUnaggregated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



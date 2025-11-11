# GetStatusRatelimit200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepts** | Pointer to **int32** | Total number of accepted requests | [optional] 
**InFlight** | Pointer to **int32** | Number of current in-flight requests | [optional] 
**LastUpdate** | Pointer to **string** | Last update time of the status | [optional] 
**Limit** | Pointer to **int32** | Current concurrency limit | [optional] 
**Load** | Pointer to **float32** | Calculated system load (in_flight / limit) | [optional] 
**MinRtt** | Pointer to **float32** | Minimum observed RTT in milliseconds | [optional] 
**Rejects** | Pointer to **int32** | Total number of rejected requests | [optional] 
**Rtt** | Pointer to **float32** | Smoothed RTT in milliseconds | [optional] 
**Throttled** | Pointer to **int32** | Total number of throttled requests | [optional] 

## Methods

### NewGetStatusRatelimit200Response

`func NewGetStatusRatelimit200Response() *GetStatusRatelimit200Response`

NewGetStatusRatelimit200Response instantiates a new GetStatusRatelimit200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusRatelimit200ResponseWithDefaults

`func NewGetStatusRatelimit200ResponseWithDefaults() *GetStatusRatelimit200Response`

NewGetStatusRatelimit200ResponseWithDefaults instantiates a new GetStatusRatelimit200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepts

`func (o *GetStatusRatelimit200Response) GetAccepts() int32`

GetAccepts returns the Accepts field if non-nil, zero value otherwise.

### GetAcceptsOk

`func (o *GetStatusRatelimit200Response) GetAcceptsOk() (*int32, bool)`

GetAcceptsOk returns a tuple with the Accepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepts

`func (o *GetStatusRatelimit200Response) SetAccepts(v int32)`

SetAccepts sets Accepts field to given value.

### HasAccepts

`func (o *GetStatusRatelimit200Response) HasAccepts() bool`

HasAccepts returns a boolean if a field has been set.

### GetInFlight

`func (o *GetStatusRatelimit200Response) GetInFlight() int32`

GetInFlight returns the InFlight field if non-nil, zero value otherwise.

### GetInFlightOk

`func (o *GetStatusRatelimit200Response) GetInFlightOk() (*int32, bool)`

GetInFlightOk returns a tuple with the InFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFlight

`func (o *GetStatusRatelimit200Response) SetInFlight(v int32)`

SetInFlight sets InFlight field to given value.

### HasInFlight

`func (o *GetStatusRatelimit200Response) HasInFlight() bool`

HasInFlight returns a boolean if a field has been set.

### GetLastUpdate

`func (o *GetStatusRatelimit200Response) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *GetStatusRatelimit200Response) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *GetStatusRatelimit200Response) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *GetStatusRatelimit200Response) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetLimit

`func (o *GetStatusRatelimit200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetStatusRatelimit200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetStatusRatelimit200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetStatusRatelimit200Response) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLoad

`func (o *GetStatusRatelimit200Response) GetLoad() float32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *GetStatusRatelimit200Response) GetLoadOk() (*float32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *GetStatusRatelimit200Response) SetLoad(v float32)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *GetStatusRatelimit200Response) HasLoad() bool`

HasLoad returns a boolean if a field has been set.

### GetMinRtt

`func (o *GetStatusRatelimit200Response) GetMinRtt() float32`

GetMinRtt returns the MinRtt field if non-nil, zero value otherwise.

### GetMinRttOk

`func (o *GetStatusRatelimit200Response) GetMinRttOk() (*float32, bool)`

GetMinRttOk returns a tuple with the MinRtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRtt

`func (o *GetStatusRatelimit200Response) SetMinRtt(v float32)`

SetMinRtt sets MinRtt field to given value.

### HasMinRtt

`func (o *GetStatusRatelimit200Response) HasMinRtt() bool`

HasMinRtt returns a boolean if a field has been set.

### GetRejects

`func (o *GetStatusRatelimit200Response) GetRejects() int32`

GetRejects returns the Rejects field if non-nil, zero value otherwise.

### GetRejectsOk

`func (o *GetStatusRatelimit200Response) GetRejectsOk() (*int32, bool)`

GetRejectsOk returns a tuple with the Rejects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejects

`func (o *GetStatusRatelimit200Response) SetRejects(v int32)`

SetRejects sets Rejects field to given value.

### HasRejects

`func (o *GetStatusRatelimit200Response) HasRejects() bool`

HasRejects returns a boolean if a field has been set.

### GetRtt

`func (o *GetStatusRatelimit200Response) GetRtt() float32`

GetRtt returns the Rtt field if non-nil, zero value otherwise.

### GetRttOk

`func (o *GetStatusRatelimit200Response) GetRttOk() (*float32, bool)`

GetRttOk returns a tuple with the Rtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtt

`func (o *GetStatusRatelimit200Response) SetRtt(v float32)`

SetRtt sets Rtt field to given value.

### HasRtt

`func (o *GetStatusRatelimit200Response) HasRtt() bool`

HasRtt returns a boolean if a field has been set.

### GetThrottled

`func (o *GetStatusRatelimit200Response) GetThrottled() int32`

GetThrottled returns the Throttled field if non-nil, zero value otherwise.

### GetThrottledOk

`func (o *GetStatusRatelimit200Response) GetThrottledOk() (*int32, bool)`

GetThrottledOk returns a tuple with the Throttled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottled

`func (o *GetStatusRatelimit200Response) SetThrottled(v int32)`

SetThrottled sets Throttled field to given value.

### HasThrottled

`func (o *GetStatusRatelimit200Response) HasThrottled() bool`

HasThrottled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



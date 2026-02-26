# GetNetworkPortscan200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PortStatus** | Pointer to **string** | \&quot;open\&quot;, \&quot;closed\&quot;, 或 \&quot;timeout\&quot; | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 

## Methods

### NewGetNetworkPortscan200Response

`func NewGetNetworkPortscan200Response() *GetNetworkPortscan200Response`

NewGetNetworkPortscan200Response instantiates a new GetNetworkPortscan200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkPortscan200ResponseWithDefaults

`func NewGetNetworkPortscan200ResponseWithDefaults() *GetNetworkPortscan200Response`

NewGetNetworkPortscan200ResponseWithDefaults instantiates a new GetNetworkPortscan200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GetNetworkPortscan200Response) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetNetworkPortscan200Response) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetNetworkPortscan200Response) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetNetworkPortscan200Response) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *GetNetworkPortscan200Response) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetNetworkPortscan200Response) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetNetworkPortscan200Response) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetNetworkPortscan200Response) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortStatus

`func (o *GetNetworkPortscan200Response) GetPortStatus() string`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *GetNetworkPortscan200Response) GetPortStatusOk() (*string, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *GetNetworkPortscan200Response) SetPortStatus(v string)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *GetNetworkPortscan200Response) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkPortscan200Response) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkPortscan200Response) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkPortscan200Response) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkPortscan200Response) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetNetworkDns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Records** | Pointer to [**[]GetNetworkDns200ResponseRecordsInner**](GetNetworkDns200ResponseRecordsInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetNetworkDns200Response

`func NewGetNetworkDns200Response() *GetNetworkDns200Response`

NewGetNetworkDns200Response instantiates a new GetNetworkDns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDns200ResponseWithDefaults

`func NewGetNetworkDns200ResponseWithDefaults() *GetNetworkDns200Response`

NewGetNetworkDns200ResponseWithDefaults instantiates a new GetNetworkDns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetNetworkDns200Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkDns200Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkDns200Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkDns200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDomain

`func (o *GetNetworkDns200Response) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetNetworkDns200Response) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetNetworkDns200Response) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetNetworkDns200Response) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetError

`func (o *GetNetworkDns200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetNetworkDns200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetNetworkDns200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetNetworkDns200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRecords

`func (o *GetNetworkDns200Response) GetRecords() []GetNetworkDns200ResponseRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GetNetworkDns200Response) GetRecordsOk() (*[]GetNetworkDns200ResponseRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GetNetworkDns200Response) SetRecords(v []GetNetworkDns200ResponseRecordsInner)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GetNetworkDns200Response) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkDns200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkDns200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkDns200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkDns200Response) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



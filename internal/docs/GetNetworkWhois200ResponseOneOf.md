# GetNetworkWhois200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Whois** | Pointer to **string** | **WHOIS原始文本**  返回未经处理的原始WHOIS查询结果文本。 | [optional] 

## Methods

### NewGetNetworkWhois200ResponseOneOf

`func NewGetNetworkWhois200ResponseOneOf() *GetNetworkWhois200ResponseOneOf`

NewGetNetworkWhois200ResponseOneOf instantiates a new GetNetworkWhois200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWhois200ResponseOneOfWithDefaults

`func NewGetNetworkWhois200ResponseOneOfWithDefaults() *GetNetworkWhois200ResponseOneOf`

NewGetNetworkWhois200ResponseOneOfWithDefaults instantiates a new GetNetworkWhois200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetNetworkWhois200ResponseOneOf) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkWhois200ResponseOneOf) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkWhois200ResponseOneOf) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkWhois200ResponseOneOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetWhois

`func (o *GetNetworkWhois200ResponseOneOf) GetWhois() string`

GetWhois returns the Whois field if non-nil, zero value otherwise.

### GetWhoisOk

`func (o *GetNetworkWhois200ResponseOneOf) GetWhoisOk() (*string, bool)`

GetWhoisOk returns a tuple with the Whois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhois

`func (o *GetNetworkWhois200ResponseOneOf) SetWhois(v string)`

SetWhois sets Whois field to given value.

### HasWhois

`func (o *GetNetworkWhois200ResponseOneOf) HasWhois() bool`

HasWhois returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



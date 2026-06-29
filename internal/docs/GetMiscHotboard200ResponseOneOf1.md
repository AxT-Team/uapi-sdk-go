# GetMiscHotboard200ResponseOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | 匹配到的结果数量。 | [optional] 
**Keyword** | Pointer to **string** | 搜索关键词。 | [optional] 
**Results** | Pointer to [**[]GetMiscHotboard200ResponseOneOf1ResultsInner**](GetMiscHotboard200ResponseOneOf1ResultsInner.md) | 搜索结果数组。 | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMiscHotboard200ResponseOneOf1

`func NewGetMiscHotboard200ResponseOneOf1() *GetMiscHotboard200ResponseOneOf1`

NewGetMiscHotboard200ResponseOneOf1 instantiates a new GetMiscHotboard200ResponseOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHotboard200ResponseOneOf1WithDefaults

`func NewGetMiscHotboard200ResponseOneOf1WithDefaults() *GetMiscHotboard200ResponseOneOf1`

NewGetMiscHotboard200ResponseOneOf1WithDefaults instantiates a new GetMiscHotboard200ResponseOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetMiscHotboard200ResponseOneOf1) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetMiscHotboard200ResponseOneOf1) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetMiscHotboard200ResponseOneOf1) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetMiscHotboard200ResponseOneOf1) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetKeyword

`func (o *GetMiscHotboard200ResponseOneOf1) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *GetMiscHotboard200ResponseOneOf1) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *GetMiscHotboard200ResponseOneOf1) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *GetMiscHotboard200ResponseOneOf1) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetResults

`func (o *GetMiscHotboard200ResponseOneOf1) GetResults() []GetMiscHotboard200ResponseOneOf1ResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetMiscHotboard200ResponseOneOf1) GetResultsOk() (*[]GetMiscHotboard200ResponseOneOf1ResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetMiscHotboard200ResponseOneOf1) SetResults(v []GetMiscHotboard200ResponseOneOf1ResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetMiscHotboard200ResponseOneOf1) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetType

`func (o *GetMiscHotboard200ResponseOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMiscHotboard200ResponseOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMiscHotboard200ResponseOneOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMiscHotboard200ResponseOneOf1) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



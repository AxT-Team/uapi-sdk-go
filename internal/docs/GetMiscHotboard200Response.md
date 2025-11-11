# GetMiscHotboard200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetMiscHotboard200ResponseListInner**](GetMiscHotboard200ResponseListInner.md) | 热榜条目列表。 | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMiscHotboard200Response

`func NewGetMiscHotboard200Response() *GetMiscHotboard200Response`

NewGetMiscHotboard200Response instantiates a new GetMiscHotboard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscHotboard200ResponseWithDefaults

`func NewGetMiscHotboard200ResponseWithDefaults() *GetMiscHotboard200Response`

NewGetMiscHotboard200ResponseWithDefaults instantiates a new GetMiscHotboard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetMiscHotboard200Response) GetList() []GetMiscHotboard200ResponseListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetMiscHotboard200Response) GetListOk() (*[]GetMiscHotboard200ResponseListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetMiscHotboard200Response) SetList(v []GetMiscHotboard200ResponseListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetMiscHotboard200Response) HasList() bool`

HasList returns a boolean if a field has been set.

### GetType

`func (o *GetMiscHotboard200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMiscHotboard200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMiscHotboard200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMiscHotboard200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetMiscHotboard200Response) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetMiscHotboard200Response) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetMiscHotboard200Response) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetMiscHotboard200Response) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



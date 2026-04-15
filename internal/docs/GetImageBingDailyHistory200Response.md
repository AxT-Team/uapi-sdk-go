# GetImageBingDailyHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolution** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]GetImageBingDailyHistory200ResponseItemsInner**](GetImageBingDailyHistory200ResponseItemsInner.md) |  | [optional] 
**Pagination** | Pointer to [**GetImageBingDailyHistory200ResponsePagination**](GetImageBingDailyHistory200ResponsePagination.md) |  | [optional] 

## Methods

### NewGetImageBingDailyHistory200Response

`func NewGetImageBingDailyHistory200Response() *GetImageBingDailyHistory200Response`

NewGetImageBingDailyHistory200Response instantiates a new GetImageBingDailyHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBingDailyHistory200ResponseWithDefaults

`func NewGetImageBingDailyHistory200ResponseWithDefaults() *GetImageBingDailyHistory200Response`

NewGetImageBingDailyHistory200ResponseWithDefaults instantiates a new GetImageBingDailyHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolution

`func (o *GetImageBingDailyHistory200Response) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *GetImageBingDailyHistory200Response) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *GetImageBingDailyHistory200Response) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *GetImageBingDailyHistory200Response) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetItems

`func (o *GetImageBingDailyHistory200Response) GetItems() []GetImageBingDailyHistory200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetImageBingDailyHistory200Response) GetItemsOk() (*[]GetImageBingDailyHistory200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetImageBingDailyHistory200Response) SetItems(v []GetImageBingDailyHistory200ResponseItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetImageBingDailyHistory200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPagination

`func (o *GetImageBingDailyHistory200Response) GetPagination() GetImageBingDailyHistory200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetImageBingDailyHistory200Response) GetPaginationOk() (*GetImageBingDailyHistory200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetImageBingDailyHistory200Response) SetPagination(v GetImageBingDailyHistory200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetImageBingDailyHistory200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



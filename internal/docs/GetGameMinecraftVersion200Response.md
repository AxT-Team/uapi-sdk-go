# GetGameMinecraftVersion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latest** | Pointer to [**GetGameMinecraftVersion200ResponseLatest**](GetGameMinecraftVersion200ResponseLatest.md) |  | [optional] 
**Versions** | Pointer to [**[]GetGameMinecraftVersion200ResponseVersionsInner**](GetGameMinecraftVersion200ResponseVersionsInner.md) |  | [optional] 

## Methods

### NewGetGameMinecraftVersion200Response

`func NewGetGameMinecraftVersion200Response() *GetGameMinecraftVersion200Response`

NewGetGameMinecraftVersion200Response instantiates a new GetGameMinecraftVersion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameMinecraftVersion200ResponseWithDefaults

`func NewGetGameMinecraftVersion200ResponseWithDefaults() *GetGameMinecraftVersion200Response`

NewGetGameMinecraftVersion200ResponseWithDefaults instantiates a new GetGameMinecraftVersion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatest

`func (o *GetGameMinecraftVersion200Response) GetLatest() GetGameMinecraftVersion200ResponseLatest`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *GetGameMinecraftVersion200Response) GetLatestOk() (*GetGameMinecraftVersion200ResponseLatest, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *GetGameMinecraftVersion200Response) SetLatest(v GetGameMinecraftVersion200ResponseLatest)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *GetGameMinecraftVersion200Response) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetVersions

`func (o *GetGameMinecraftVersion200Response) GetVersions() []GetGameMinecraftVersion200ResponseVersionsInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetGameMinecraftVersion200Response) GetVersionsOk() (*[]GetGameMinecraftVersion200ResponseVersionsInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetGameMinecraftVersion200Response) SetVersions(v []GetGameMinecraftVersion200ResponseVersionsInner)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GetGameMinecraftVersion200Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetDailyWord200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Words** | Pointer to [**[]GetDailyWord200ResponseWordsInner**](GetDailyWord200ResponseWordsInner.md) |  | [optional] 

## Methods

### NewGetDailyWord200Response

`func NewGetDailyWord200Response() *GetDailyWord200Response`

NewGetDailyWord200Response instantiates a new GetDailyWord200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyWord200ResponseWithDefaults

`func NewGetDailyWord200ResponseWithDefaults() *GetDailyWord200Response`

NewGetDailyWord200ResponseWithDefaults instantiates a new GetDailyWord200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *GetDailyWord200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetDailyWord200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetDailyWord200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetDailyWord200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCount

`func (o *GetDailyWord200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetDailyWord200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetDailyWord200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetDailyWord200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLang

`func (o *GetDailyWord200Response) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *GetDailyWord200Response) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *GetDailyWord200Response) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *GetDailyWord200Response) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetWords

`func (o *GetDailyWord200Response) GetWords() []GetDailyWord200ResponseWordsInner`

GetWords returns the Words field if non-nil, zero value otherwise.

### GetWordsOk

`func (o *GetDailyWord200Response) GetWordsOk() (*[]GetDailyWord200ResponseWordsInner, bool)`

GetWordsOk returns a tuple with the Words field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWords

`func (o *GetDailyWord200Response) SetWords(v []GetDailyWord200ResponseWordsInner)`

SetWords sets Words field to given value.

### HasWords

`func (o *GetDailyWord200Response) HasWords() bool`

HasWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



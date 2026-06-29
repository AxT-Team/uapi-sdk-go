# GetMiscMovieBoxOffice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetMiscMovieBoxOffice200ResponseListInner**](GetMiscMovieBoxOffice200ResponseListInner.md) | 影片排名列表 | [optional] 
**Market** | Pointer to [**GetMiscMovieBoxOffice200ResponseMarket**](GetMiscMovieBoxOffice200ResponseMarket.md) |  | [optional] 
**TotalItems** | Pointer to **int32** | 返回的影片数量 | [optional] 
**UpdateGapSeconds** | Pointer to **int32** | 上游数据刷新间隔（秒） | [optional] 
**UpdateTime** | Pointer to **string** | 数据更新时间的格式化字符串 | [optional] 
**UpdatedAt** | Pointer to **int32** | 数据更新时间戳（毫秒） | [optional] 

## Methods

### NewGetMiscMovieBoxOffice200Response

`func NewGetMiscMovieBoxOffice200Response() *GetMiscMovieBoxOffice200Response`

NewGetMiscMovieBoxOffice200Response instantiates a new GetMiscMovieBoxOffice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscMovieBoxOffice200ResponseWithDefaults

`func NewGetMiscMovieBoxOffice200ResponseWithDefaults() *GetMiscMovieBoxOffice200Response`

NewGetMiscMovieBoxOffice200ResponseWithDefaults instantiates a new GetMiscMovieBoxOffice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetMiscMovieBoxOffice200Response) GetList() []GetMiscMovieBoxOffice200ResponseListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetMiscMovieBoxOffice200Response) GetListOk() (*[]GetMiscMovieBoxOffice200ResponseListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetMiscMovieBoxOffice200Response) SetList(v []GetMiscMovieBoxOffice200ResponseListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetMiscMovieBoxOffice200Response) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMarket

`func (o *GetMiscMovieBoxOffice200Response) GetMarket() GetMiscMovieBoxOffice200ResponseMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GetMiscMovieBoxOffice200Response) GetMarketOk() (*GetMiscMovieBoxOffice200ResponseMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GetMiscMovieBoxOffice200Response) SetMarket(v GetMiscMovieBoxOffice200ResponseMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *GetMiscMovieBoxOffice200Response) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetTotalItems

`func (o *GetMiscMovieBoxOffice200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *GetMiscMovieBoxOffice200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *GetMiscMovieBoxOffice200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *GetMiscMovieBoxOffice200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetUpdateGapSeconds

`func (o *GetMiscMovieBoxOffice200Response) GetUpdateGapSeconds() int32`

GetUpdateGapSeconds returns the UpdateGapSeconds field if non-nil, zero value otherwise.

### GetUpdateGapSecondsOk

`func (o *GetMiscMovieBoxOffice200Response) GetUpdateGapSecondsOk() (*int32, bool)`

GetUpdateGapSecondsOk returns a tuple with the UpdateGapSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateGapSeconds

`func (o *GetMiscMovieBoxOffice200Response) SetUpdateGapSeconds(v int32)`

SetUpdateGapSeconds sets UpdateGapSeconds field to given value.

### HasUpdateGapSeconds

`func (o *GetMiscMovieBoxOffice200Response) HasUpdateGapSeconds() bool`

HasUpdateGapSeconds returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetMiscMovieBoxOffice200Response) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetMiscMovieBoxOffice200Response) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetMiscMovieBoxOffice200Response) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetMiscMovieBoxOffice200Response) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetMiscMovieBoxOffice200Response) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetMiscMovieBoxOffice200Response) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetMiscMovieBoxOffice200Response) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetMiscMovieBoxOffice200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



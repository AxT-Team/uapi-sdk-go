# GetGameEpicFree200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | 游戏的唯一标识ID。 | [optional] 
**Title** | Pointer to **string** | 游戏的完整标题名称。 | [optional] 
**Cover** | Pointer to **string** | 游戏封面图片的URL地址。 | [optional] 
**OriginalPrice** | Pointer to **float32** | 游戏的原价，单位为人民币元。 | [optional] 
**OriginalPriceDesc** | Pointer to **string** | 格式化后的原价描述字符串。 | [optional] 
**Description** | Pointer to **string** | 游戏的简介描述。 | [optional] 
**Seller** | Pointer to **string** | 游戏的发行商或销售商。 | [optional] 
**IsFreeNow** | Pointer to **bool** | 当前是否处于免费状态。 | [optional] 
**FreeStart** | Pointer to **string** | 免费开始时间的可读字符串格式。 | [optional] 
**FreeStartAt** | Pointer to **int32** | 免费开始时间的13位毫秒时间戳。 | [optional] 
**FreeEnd** | Pointer to **string** | 免费结束时间的可读字符串格式。 | [optional] 
**FreeEndAt** | Pointer to **int32** | 免费结束时间的13位毫秒时间戳。 | [optional] 
**Link** | Pointer to **string** | 游戏在Epic Games商店的详情页链接。 | [optional] 

## Methods

### NewGetGameEpicFree200ResponseDataInner

`func NewGetGameEpicFree200ResponseDataInner() *GetGameEpicFree200ResponseDataInner`

NewGetGameEpicFree200ResponseDataInner instantiates a new GetGameEpicFree200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameEpicFree200ResponseDataInnerWithDefaults

`func NewGetGameEpicFree200ResponseDataInnerWithDefaults() *GetGameEpicFree200ResponseDataInner`

NewGetGameEpicFree200ResponseDataInnerWithDefaults instantiates a new GetGameEpicFree200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGameEpicFree200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGameEpicFree200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGameEpicFree200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetGameEpicFree200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetGameEpicFree200ResponseDataInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetGameEpicFree200ResponseDataInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetGameEpicFree200ResponseDataInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetGameEpicFree200ResponseDataInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCover

`func (o *GetGameEpicFree200ResponseDataInner) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *GetGameEpicFree200ResponseDataInner) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *GetGameEpicFree200ResponseDataInner) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *GetGameEpicFree200ResponseDataInner) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *GetGameEpicFree200ResponseDataInner) GetOriginalPrice() float32`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *GetGameEpicFree200ResponseDataInner) GetOriginalPriceOk() (*float32, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *GetGameEpicFree200ResponseDataInner) SetOriginalPrice(v float32)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *GetGameEpicFree200ResponseDataInner) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetOriginalPriceDesc

`func (o *GetGameEpicFree200ResponseDataInner) GetOriginalPriceDesc() string`

GetOriginalPriceDesc returns the OriginalPriceDesc field if non-nil, zero value otherwise.

### GetOriginalPriceDescOk

`func (o *GetGameEpicFree200ResponseDataInner) GetOriginalPriceDescOk() (*string, bool)`

GetOriginalPriceDescOk returns a tuple with the OriginalPriceDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPriceDesc

`func (o *GetGameEpicFree200ResponseDataInner) SetOriginalPriceDesc(v string)`

SetOriginalPriceDesc sets OriginalPriceDesc field to given value.

### HasOriginalPriceDesc

`func (o *GetGameEpicFree200ResponseDataInner) HasOriginalPriceDesc() bool`

HasOriginalPriceDesc returns a boolean if a field has been set.

### GetDescription

`func (o *GetGameEpicFree200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGameEpicFree200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGameEpicFree200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGameEpicFree200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeller

`func (o *GetGameEpicFree200ResponseDataInner) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *GetGameEpicFree200ResponseDataInner) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *GetGameEpicFree200ResponseDataInner) SetSeller(v string)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *GetGameEpicFree200ResponseDataInner) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetIsFreeNow

`func (o *GetGameEpicFree200ResponseDataInner) GetIsFreeNow() bool`

GetIsFreeNow returns the IsFreeNow field if non-nil, zero value otherwise.

### GetIsFreeNowOk

`func (o *GetGameEpicFree200ResponseDataInner) GetIsFreeNowOk() (*bool, bool)`

GetIsFreeNowOk returns a tuple with the IsFreeNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeNow

`func (o *GetGameEpicFree200ResponseDataInner) SetIsFreeNow(v bool)`

SetIsFreeNow sets IsFreeNow field to given value.

### HasIsFreeNow

`func (o *GetGameEpicFree200ResponseDataInner) HasIsFreeNow() bool`

HasIsFreeNow returns a boolean if a field has been set.

### GetFreeStart

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeStart() string`

GetFreeStart returns the FreeStart field if non-nil, zero value otherwise.

### GetFreeStartOk

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeStartOk() (*string, bool)`

GetFreeStartOk returns a tuple with the FreeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStart

`func (o *GetGameEpicFree200ResponseDataInner) SetFreeStart(v string)`

SetFreeStart sets FreeStart field to given value.

### HasFreeStart

`func (o *GetGameEpicFree200ResponseDataInner) HasFreeStart() bool`

HasFreeStart returns a boolean if a field has been set.

### GetFreeStartAt

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeStartAt() int32`

GetFreeStartAt returns the FreeStartAt field if non-nil, zero value otherwise.

### GetFreeStartAtOk

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeStartAtOk() (*int32, bool)`

GetFreeStartAtOk returns a tuple with the FreeStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStartAt

`func (o *GetGameEpicFree200ResponseDataInner) SetFreeStartAt(v int32)`

SetFreeStartAt sets FreeStartAt field to given value.

### HasFreeStartAt

`func (o *GetGameEpicFree200ResponseDataInner) HasFreeStartAt() bool`

HasFreeStartAt returns a boolean if a field has been set.

### GetFreeEnd

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeEnd() string`

GetFreeEnd returns the FreeEnd field if non-nil, zero value otherwise.

### GetFreeEndOk

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeEndOk() (*string, bool)`

GetFreeEndOk returns a tuple with the FreeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeEnd

`func (o *GetGameEpicFree200ResponseDataInner) SetFreeEnd(v string)`

SetFreeEnd sets FreeEnd field to given value.

### HasFreeEnd

`func (o *GetGameEpicFree200ResponseDataInner) HasFreeEnd() bool`

HasFreeEnd returns a boolean if a field has been set.

### GetFreeEndAt

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeEndAt() int32`

GetFreeEndAt returns the FreeEndAt field if non-nil, zero value otherwise.

### GetFreeEndAtOk

`func (o *GetGameEpicFree200ResponseDataInner) GetFreeEndAtOk() (*int32, bool)`

GetFreeEndAtOk returns a tuple with the FreeEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeEndAt

`func (o *GetGameEpicFree200ResponseDataInner) SetFreeEndAt(v int32)`

SetFreeEndAt sets FreeEndAt field to given value.

### HasFreeEndAt

`func (o *GetGameEpicFree200ResponseDataInner) HasFreeEndAt() bool`

HasFreeEndAt returns a boolean if a field has been set.

### GetLink

`func (o *GetGameEpicFree200ResponseDataInner) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GetGameEpicFree200ResponseDataInner) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GetGameEpicFree200ResponseDataInner) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *GetGameEpicFree200ResponseDataInner) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



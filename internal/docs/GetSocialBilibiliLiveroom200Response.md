# GetSocialBilibiliLiveroom200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **float32** | 主播的用户ID (mid)。 | [optional] 
**RoomId** | Pointer to **float32** | 直播间的真实房间号（长号）。 | [optional] 
**ShortId** | Pointer to **float32** | 直播间的短号（靓号）。如果没有设置，则为0。 | [optional] 
**Attention** | Pointer to **float32** | 主播的粉丝数（关注数量）。 | [optional] 
**Online** | Pointer to **float32** | 直播间当前的人气值。注意这不是真实在线人数。 | [optional] 
**LiveStatus** | Pointer to **float32** | 直播状态。0:未开播, 1:直播中, 2:轮播中。 | [optional] 
**AreaId** | Pointer to **float32** | 分区ID。 | [optional] 
**ParentAreaName** | Pointer to **string** | 父分区名称。 | [optional] 
**AreaName** | Pointer to **string** | 子分区名称。 | [optional] 
**Background** | Pointer to **string** | 直播间背景图的URL。 | [optional] 
**Title** | Pointer to **string** | 当前直播间的标题。 | [optional] 
**UserCover** | Pointer to **string** | 用户设置的直播间封面URL。 | [optional] 
**Description** | Pointer to **string** | 直播间公告或描述，支持换行符。 | [optional] 
**LiveTime** | Pointer to **string** | 本次直播开始的时间，格式为 &#x60;YYYY-MM-DD HH:mm:ss&#x60;。如果未开播，则为空字符串。 | [optional] 
**Tags** | Pointer to **string** | 直播间设置的标签，以逗号分隔。 | [optional] 
**HotWords** | Pointer to **[]string** | 直播间热词列表，通常用于弹幕互动。 | [optional] 
**NewPendants** | Pointer to **map[string]interface{}** | 主播佩戴的头像框、大航海等级等信息，结构可能比较复杂。 | [optional] 

## Methods

### NewGetSocialBilibiliLiveroom200Response

`func NewGetSocialBilibiliLiveroom200Response() *GetSocialBilibiliLiveroom200Response`

NewGetSocialBilibiliLiveroom200Response instantiates a new GetSocialBilibiliLiveroom200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialBilibiliLiveroom200ResponseWithDefaults

`func NewGetSocialBilibiliLiveroom200ResponseWithDefaults() *GetSocialBilibiliLiveroom200Response`

NewGetSocialBilibiliLiveroom200ResponseWithDefaults instantiates a new GetSocialBilibiliLiveroom200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *GetSocialBilibiliLiveroom200Response) GetUid() float32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetSocialBilibiliLiveroom200Response) GetUidOk() (*float32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetSocialBilibiliLiveroom200Response) SetUid(v float32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetSocialBilibiliLiveroom200Response) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetRoomId

`func (o *GetSocialBilibiliLiveroom200Response) GetRoomId() float32`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *GetSocialBilibiliLiveroom200Response) GetRoomIdOk() (*float32, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *GetSocialBilibiliLiveroom200Response) SetRoomId(v float32)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *GetSocialBilibiliLiveroom200Response) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetShortId

`func (o *GetSocialBilibiliLiveroom200Response) GetShortId() float32`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *GetSocialBilibiliLiveroom200Response) GetShortIdOk() (*float32, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *GetSocialBilibiliLiveroom200Response) SetShortId(v float32)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *GetSocialBilibiliLiveroom200Response) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetAttention

`func (o *GetSocialBilibiliLiveroom200Response) GetAttention() float32`

GetAttention returns the Attention field if non-nil, zero value otherwise.

### GetAttentionOk

`func (o *GetSocialBilibiliLiveroom200Response) GetAttentionOk() (*float32, bool)`

GetAttentionOk returns a tuple with the Attention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttention

`func (o *GetSocialBilibiliLiveroom200Response) SetAttention(v float32)`

SetAttention sets Attention field to given value.

### HasAttention

`func (o *GetSocialBilibiliLiveroom200Response) HasAttention() bool`

HasAttention returns a boolean if a field has been set.

### GetOnline

`func (o *GetSocialBilibiliLiveroom200Response) GetOnline() float32`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *GetSocialBilibiliLiveroom200Response) GetOnlineOk() (*float32, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *GetSocialBilibiliLiveroom200Response) SetOnline(v float32)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *GetSocialBilibiliLiveroom200Response) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetLiveStatus

`func (o *GetSocialBilibiliLiveroom200Response) GetLiveStatus() float32`

GetLiveStatus returns the LiveStatus field if non-nil, zero value otherwise.

### GetLiveStatusOk

`func (o *GetSocialBilibiliLiveroom200Response) GetLiveStatusOk() (*float32, bool)`

GetLiveStatusOk returns a tuple with the LiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStatus

`func (o *GetSocialBilibiliLiveroom200Response) SetLiveStatus(v float32)`

SetLiveStatus sets LiveStatus field to given value.

### HasLiveStatus

`func (o *GetSocialBilibiliLiveroom200Response) HasLiveStatus() bool`

HasLiveStatus returns a boolean if a field has been set.

### GetAreaId

`func (o *GetSocialBilibiliLiveroom200Response) GetAreaId() float32`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *GetSocialBilibiliLiveroom200Response) GetAreaIdOk() (*float32, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *GetSocialBilibiliLiveroom200Response) SetAreaId(v float32)`

SetAreaId sets AreaId field to given value.

### HasAreaId

`func (o *GetSocialBilibiliLiveroom200Response) HasAreaId() bool`

HasAreaId returns a boolean if a field has been set.

### GetParentAreaName

`func (o *GetSocialBilibiliLiveroom200Response) GetParentAreaName() string`

GetParentAreaName returns the ParentAreaName field if non-nil, zero value otherwise.

### GetParentAreaNameOk

`func (o *GetSocialBilibiliLiveroom200Response) GetParentAreaNameOk() (*string, bool)`

GetParentAreaNameOk returns a tuple with the ParentAreaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAreaName

`func (o *GetSocialBilibiliLiveroom200Response) SetParentAreaName(v string)`

SetParentAreaName sets ParentAreaName field to given value.

### HasParentAreaName

`func (o *GetSocialBilibiliLiveroom200Response) HasParentAreaName() bool`

HasParentAreaName returns a boolean if a field has been set.

### GetAreaName

`func (o *GetSocialBilibiliLiveroom200Response) GetAreaName() string`

GetAreaName returns the AreaName field if non-nil, zero value otherwise.

### GetAreaNameOk

`func (o *GetSocialBilibiliLiveroom200Response) GetAreaNameOk() (*string, bool)`

GetAreaNameOk returns a tuple with the AreaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaName

`func (o *GetSocialBilibiliLiveroom200Response) SetAreaName(v string)`

SetAreaName sets AreaName field to given value.

### HasAreaName

`func (o *GetSocialBilibiliLiveroom200Response) HasAreaName() bool`

HasAreaName returns a boolean if a field has been set.

### GetBackground

`func (o *GetSocialBilibiliLiveroom200Response) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *GetSocialBilibiliLiveroom200Response) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *GetSocialBilibiliLiveroom200Response) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *GetSocialBilibiliLiveroom200Response) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetTitle

`func (o *GetSocialBilibiliLiveroom200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetSocialBilibiliLiveroom200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetSocialBilibiliLiveroom200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetSocialBilibiliLiveroom200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUserCover

`func (o *GetSocialBilibiliLiveroom200Response) GetUserCover() string`

GetUserCover returns the UserCover field if non-nil, zero value otherwise.

### GetUserCoverOk

`func (o *GetSocialBilibiliLiveroom200Response) GetUserCoverOk() (*string, bool)`

GetUserCoverOk returns a tuple with the UserCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCover

`func (o *GetSocialBilibiliLiveroom200Response) SetUserCover(v string)`

SetUserCover sets UserCover field to given value.

### HasUserCover

`func (o *GetSocialBilibiliLiveroom200Response) HasUserCover() bool`

HasUserCover returns a boolean if a field has been set.

### GetDescription

`func (o *GetSocialBilibiliLiveroom200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSocialBilibiliLiveroom200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSocialBilibiliLiveroom200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSocialBilibiliLiveroom200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLiveTime

`func (o *GetSocialBilibiliLiveroom200Response) GetLiveTime() string`

GetLiveTime returns the LiveTime field if non-nil, zero value otherwise.

### GetLiveTimeOk

`func (o *GetSocialBilibiliLiveroom200Response) GetLiveTimeOk() (*string, bool)`

GetLiveTimeOk returns a tuple with the LiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveTime

`func (o *GetSocialBilibiliLiveroom200Response) SetLiveTime(v string)`

SetLiveTime sets LiveTime field to given value.

### HasLiveTime

`func (o *GetSocialBilibiliLiveroom200Response) HasLiveTime() bool`

HasLiveTime returns a boolean if a field has been set.

### GetTags

`func (o *GetSocialBilibiliLiveroom200Response) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetSocialBilibiliLiveroom200Response) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetSocialBilibiliLiveroom200Response) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetSocialBilibiliLiveroom200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHotWords

`func (o *GetSocialBilibiliLiveroom200Response) GetHotWords() []string`

GetHotWords returns the HotWords field if non-nil, zero value otherwise.

### GetHotWordsOk

`func (o *GetSocialBilibiliLiveroom200Response) GetHotWordsOk() (*[]string, bool)`

GetHotWordsOk returns a tuple with the HotWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotWords

`func (o *GetSocialBilibiliLiveroom200Response) SetHotWords(v []string)`

SetHotWords sets HotWords field to given value.

### HasHotWords

`func (o *GetSocialBilibiliLiveroom200Response) HasHotWords() bool`

HasHotWords returns a boolean if a field has been set.

### GetNewPendants

`func (o *GetSocialBilibiliLiveroom200Response) GetNewPendants() map[string]interface{}`

GetNewPendants returns the NewPendants field if non-nil, zero value otherwise.

### GetNewPendantsOk

`func (o *GetSocialBilibiliLiveroom200Response) GetNewPendantsOk() (*map[string]interface{}, bool)`

GetNewPendantsOk returns a tuple with the NewPendants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPendants

`func (o *GetSocialBilibiliLiveroom200Response) SetNewPendants(v map[string]interface{})`

SetNewPendants sets NewPendants field to given value.

### HasNewPendants

`func (o *GetSocialBilibiliLiveroom200Response) HasNewPendants() bool`

HasNewPendants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



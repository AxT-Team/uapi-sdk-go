# GetSocialQqUserinfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **int32** | 年龄 | [optional] 
**AvatarUrl** | Pointer to **string** | 头像URL | [optional] 
**BigClubLevel** | Pointer to **int32** | QQ大会员等级 | [optional] 
**Email** | Pointer to **string** | QQ邮箱 | [optional] 
**GreenDiamondLevel** | Pointer to **int32** | 绿钻等级（可选） | [optional] 
**IsBigClub** | Pointer to **bool** | 是否为QQ大会员用户 | [optional] 
**IsSvip** | Pointer to **bool** | 是否为SVIP用户 | [optional] 
**IsVip** | Pointer to **bool** | 是否为VIP用户 | [optional] 
**IsYearsVip** | Pointer to **bool** | 是否为年费VIP用户 | [optional] 
**LastUpdated** | Pointer to **string** | 最后更新时间（ISO 8601格式） | [optional] 
**Location** | Pointer to **string** | 地理位置（省市） | [optional] 
**LongNick** | Pointer to **string** | 个性签名 | [optional] 
**LoverVipLevel** | Pointer to **int32** | 情侣/恋人类会员等级（可选） | [optional] 
**Nickname** | Pointer to **string** | 用户昵称 | [optional] 
**PrivilegeIcons** | Pointer to [**GetSocialQqUserinfo200ResponsePrivilegeIcons**](GetSocialQqUserinfo200ResponsePrivilegeIcons.md) |  | [optional] 
**Qid** | Pointer to **string** | QQ个性域名 | [optional] 
**Qq** | Pointer to **string** | QQ号 | [optional] 
**QqLevel** | Pointer to **NullableInt32** | QQ等级。用户隐藏时返回 null | [optional] 
**RegTime** | Pointer to **string** | 注册时间（ISO 8601格式） | [optional] 
**Sex** | Pointer to **string** | 性别 | [optional] 
**VideoVipLevel** | Pointer to **int32** | 腾讯影视会员等级（可选） | [optional] 
**VipLevel** | Pointer to **int32** | VIP等级 | [optional] 
**VipStatus** | Pointer to **int32** | 会员开通状态 | [optional] 
**VipType** | Pointer to **int32** | 会员类型 | [optional] 
**YellowDiamondLevel** | Pointer to **int32** | 黄钻等级（可选） | [optional] 

## Methods

### NewGetSocialQqUserinfo200Response

`func NewGetSocialQqUserinfo200Response() *GetSocialQqUserinfo200Response`

NewGetSocialQqUserinfo200Response instantiates a new GetSocialQqUserinfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialQqUserinfo200ResponseWithDefaults

`func NewGetSocialQqUserinfo200ResponseWithDefaults() *GetSocialQqUserinfo200Response`

NewGetSocialQqUserinfo200ResponseWithDefaults instantiates a new GetSocialQqUserinfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *GetSocialQqUserinfo200Response) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *GetSocialQqUserinfo200Response) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *GetSocialQqUserinfo200Response) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *GetSocialQqUserinfo200Response) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *GetSocialQqUserinfo200Response) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetSocialQqUserinfo200Response) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetSocialQqUserinfo200Response) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetSocialQqUserinfo200Response) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBigClubLevel

`func (o *GetSocialQqUserinfo200Response) GetBigClubLevel() int32`

GetBigClubLevel returns the BigClubLevel field if non-nil, zero value otherwise.

### GetBigClubLevelOk

`func (o *GetSocialQqUserinfo200Response) GetBigClubLevelOk() (*int32, bool)`

GetBigClubLevelOk returns a tuple with the BigClubLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigClubLevel

`func (o *GetSocialQqUserinfo200Response) SetBigClubLevel(v int32)`

SetBigClubLevel sets BigClubLevel field to given value.

### HasBigClubLevel

`func (o *GetSocialQqUserinfo200Response) HasBigClubLevel() bool`

HasBigClubLevel returns a boolean if a field has been set.

### GetEmail

`func (o *GetSocialQqUserinfo200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetSocialQqUserinfo200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetSocialQqUserinfo200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetSocialQqUserinfo200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGreenDiamondLevel

`func (o *GetSocialQqUserinfo200Response) GetGreenDiamondLevel() int32`

GetGreenDiamondLevel returns the GreenDiamondLevel field if non-nil, zero value otherwise.

### GetGreenDiamondLevelOk

`func (o *GetSocialQqUserinfo200Response) GetGreenDiamondLevelOk() (*int32, bool)`

GetGreenDiamondLevelOk returns a tuple with the GreenDiamondLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreenDiamondLevel

`func (o *GetSocialQqUserinfo200Response) SetGreenDiamondLevel(v int32)`

SetGreenDiamondLevel sets GreenDiamondLevel field to given value.

### HasGreenDiamondLevel

`func (o *GetSocialQqUserinfo200Response) HasGreenDiamondLevel() bool`

HasGreenDiamondLevel returns a boolean if a field has been set.

### GetIsBigClub

`func (o *GetSocialQqUserinfo200Response) GetIsBigClub() bool`

GetIsBigClub returns the IsBigClub field if non-nil, zero value otherwise.

### GetIsBigClubOk

`func (o *GetSocialQqUserinfo200Response) GetIsBigClubOk() (*bool, bool)`

GetIsBigClubOk returns a tuple with the IsBigClub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBigClub

`func (o *GetSocialQqUserinfo200Response) SetIsBigClub(v bool)`

SetIsBigClub sets IsBigClub field to given value.

### HasIsBigClub

`func (o *GetSocialQqUserinfo200Response) HasIsBigClub() bool`

HasIsBigClub returns a boolean if a field has been set.

### GetIsSvip

`func (o *GetSocialQqUserinfo200Response) GetIsSvip() bool`

GetIsSvip returns the IsSvip field if non-nil, zero value otherwise.

### GetIsSvipOk

`func (o *GetSocialQqUserinfo200Response) GetIsSvipOk() (*bool, bool)`

GetIsSvipOk returns a tuple with the IsSvip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSvip

`func (o *GetSocialQqUserinfo200Response) SetIsSvip(v bool)`

SetIsSvip sets IsSvip field to given value.

### HasIsSvip

`func (o *GetSocialQqUserinfo200Response) HasIsSvip() bool`

HasIsSvip returns a boolean if a field has been set.

### GetIsVip

`func (o *GetSocialQqUserinfo200Response) GetIsVip() bool`

GetIsVip returns the IsVip field if non-nil, zero value otherwise.

### GetIsVipOk

`func (o *GetSocialQqUserinfo200Response) GetIsVipOk() (*bool, bool)`

GetIsVipOk returns a tuple with the IsVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVip

`func (o *GetSocialQqUserinfo200Response) SetIsVip(v bool)`

SetIsVip sets IsVip field to given value.

### HasIsVip

`func (o *GetSocialQqUserinfo200Response) HasIsVip() bool`

HasIsVip returns a boolean if a field has been set.

### GetIsYearsVip

`func (o *GetSocialQqUserinfo200Response) GetIsYearsVip() bool`

GetIsYearsVip returns the IsYearsVip field if non-nil, zero value otherwise.

### GetIsYearsVipOk

`func (o *GetSocialQqUserinfo200Response) GetIsYearsVipOk() (*bool, bool)`

GetIsYearsVipOk returns a tuple with the IsYearsVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYearsVip

`func (o *GetSocialQqUserinfo200Response) SetIsYearsVip(v bool)`

SetIsYearsVip sets IsYearsVip field to given value.

### HasIsYearsVip

`func (o *GetSocialQqUserinfo200Response) HasIsYearsVip() bool`

HasIsYearsVip returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetSocialQqUserinfo200Response) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetSocialQqUserinfo200Response) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetSocialQqUserinfo200Response) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetSocialQqUserinfo200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLocation

`func (o *GetSocialQqUserinfo200Response) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetSocialQqUserinfo200Response) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetSocialQqUserinfo200Response) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetSocialQqUserinfo200Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLongNick

`func (o *GetSocialQqUserinfo200Response) GetLongNick() string`

GetLongNick returns the LongNick field if non-nil, zero value otherwise.

### GetLongNickOk

`func (o *GetSocialQqUserinfo200Response) GetLongNickOk() (*string, bool)`

GetLongNickOk returns a tuple with the LongNick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongNick

`func (o *GetSocialQqUserinfo200Response) SetLongNick(v string)`

SetLongNick sets LongNick field to given value.

### HasLongNick

`func (o *GetSocialQqUserinfo200Response) HasLongNick() bool`

HasLongNick returns a boolean if a field has been set.

### GetLoverVipLevel

`func (o *GetSocialQqUserinfo200Response) GetLoverVipLevel() int32`

GetLoverVipLevel returns the LoverVipLevel field if non-nil, zero value otherwise.

### GetLoverVipLevelOk

`func (o *GetSocialQqUserinfo200Response) GetLoverVipLevelOk() (*int32, bool)`

GetLoverVipLevelOk returns a tuple with the LoverVipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoverVipLevel

`func (o *GetSocialQqUserinfo200Response) SetLoverVipLevel(v int32)`

SetLoverVipLevel sets LoverVipLevel field to given value.

### HasLoverVipLevel

`func (o *GetSocialQqUserinfo200Response) HasLoverVipLevel() bool`

HasLoverVipLevel returns a boolean if a field has been set.

### GetNickname

`func (o *GetSocialQqUserinfo200Response) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *GetSocialQqUserinfo200Response) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *GetSocialQqUserinfo200Response) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *GetSocialQqUserinfo200Response) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPrivilegeIcons

`func (o *GetSocialQqUserinfo200Response) GetPrivilegeIcons() GetSocialQqUserinfo200ResponsePrivilegeIcons`

GetPrivilegeIcons returns the PrivilegeIcons field if non-nil, zero value otherwise.

### GetPrivilegeIconsOk

`func (o *GetSocialQqUserinfo200Response) GetPrivilegeIconsOk() (*GetSocialQqUserinfo200ResponsePrivilegeIcons, bool)`

GetPrivilegeIconsOk returns a tuple with the PrivilegeIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeIcons

`func (o *GetSocialQqUserinfo200Response) SetPrivilegeIcons(v GetSocialQqUserinfo200ResponsePrivilegeIcons)`

SetPrivilegeIcons sets PrivilegeIcons field to given value.

### HasPrivilegeIcons

`func (o *GetSocialQqUserinfo200Response) HasPrivilegeIcons() bool`

HasPrivilegeIcons returns a boolean if a field has been set.

### GetQid

`func (o *GetSocialQqUserinfo200Response) GetQid() string`

GetQid returns the Qid field if non-nil, zero value otherwise.

### GetQidOk

`func (o *GetSocialQqUserinfo200Response) GetQidOk() (*string, bool)`

GetQidOk returns a tuple with the Qid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQid

`func (o *GetSocialQqUserinfo200Response) SetQid(v string)`

SetQid sets Qid field to given value.

### HasQid

`func (o *GetSocialQqUserinfo200Response) HasQid() bool`

HasQid returns a boolean if a field has been set.

### GetQq

`func (o *GetSocialQqUserinfo200Response) GetQq() string`

GetQq returns the Qq field if non-nil, zero value otherwise.

### GetQqOk

`func (o *GetSocialQqUserinfo200Response) GetQqOk() (*string, bool)`

GetQqOk returns a tuple with the Qq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQq

`func (o *GetSocialQqUserinfo200Response) SetQq(v string)`

SetQq sets Qq field to given value.

### HasQq

`func (o *GetSocialQqUserinfo200Response) HasQq() bool`

HasQq returns a boolean if a field has been set.

### GetQqLevel

`func (o *GetSocialQqUserinfo200Response) GetQqLevel() int32`

GetQqLevel returns the QqLevel field if non-nil, zero value otherwise.

### GetQqLevelOk

`func (o *GetSocialQqUserinfo200Response) GetQqLevelOk() (*int32, bool)`

GetQqLevelOk returns a tuple with the QqLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQqLevel

`func (o *GetSocialQqUserinfo200Response) SetQqLevel(v int32)`

SetQqLevel sets QqLevel field to given value.

### HasQqLevel

`func (o *GetSocialQqUserinfo200Response) HasQqLevel() bool`

HasQqLevel returns a boolean if a field has been set.

### SetQqLevelNil

`func (o *GetSocialQqUserinfo200Response) SetQqLevelNil(b bool)`

 SetQqLevelNil sets the value for QqLevel to be an explicit nil

### UnsetQqLevel
`func (o *GetSocialQqUserinfo200Response) UnsetQqLevel()`

UnsetQqLevel ensures that no value is present for QqLevel, not even an explicit nil
### GetRegTime

`func (o *GetSocialQqUserinfo200Response) GetRegTime() string`

GetRegTime returns the RegTime field if non-nil, zero value otherwise.

### GetRegTimeOk

`func (o *GetSocialQqUserinfo200Response) GetRegTimeOk() (*string, bool)`

GetRegTimeOk returns a tuple with the RegTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegTime

`func (o *GetSocialQqUserinfo200Response) SetRegTime(v string)`

SetRegTime sets RegTime field to given value.

### HasRegTime

`func (o *GetSocialQqUserinfo200Response) HasRegTime() bool`

HasRegTime returns a boolean if a field has been set.

### GetSex

`func (o *GetSocialQqUserinfo200Response) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *GetSocialQqUserinfo200Response) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *GetSocialQqUserinfo200Response) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *GetSocialQqUserinfo200Response) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetVideoVipLevel

`func (o *GetSocialQqUserinfo200Response) GetVideoVipLevel() int32`

GetVideoVipLevel returns the VideoVipLevel field if non-nil, zero value otherwise.

### GetVideoVipLevelOk

`func (o *GetSocialQqUserinfo200Response) GetVideoVipLevelOk() (*int32, bool)`

GetVideoVipLevelOk returns a tuple with the VideoVipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoVipLevel

`func (o *GetSocialQqUserinfo200Response) SetVideoVipLevel(v int32)`

SetVideoVipLevel sets VideoVipLevel field to given value.

### HasVideoVipLevel

`func (o *GetSocialQqUserinfo200Response) HasVideoVipLevel() bool`

HasVideoVipLevel returns a boolean if a field has been set.

### GetVipLevel

`func (o *GetSocialQqUserinfo200Response) GetVipLevel() int32`

GetVipLevel returns the VipLevel field if non-nil, zero value otherwise.

### GetVipLevelOk

`func (o *GetSocialQqUserinfo200Response) GetVipLevelOk() (*int32, bool)`

GetVipLevelOk returns a tuple with the VipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipLevel

`func (o *GetSocialQqUserinfo200Response) SetVipLevel(v int32)`

SetVipLevel sets VipLevel field to given value.

### HasVipLevel

`func (o *GetSocialQqUserinfo200Response) HasVipLevel() bool`

HasVipLevel returns a boolean if a field has been set.

### GetVipStatus

`func (o *GetSocialQqUserinfo200Response) GetVipStatus() int32`

GetVipStatus returns the VipStatus field if non-nil, zero value otherwise.

### GetVipStatusOk

`func (o *GetSocialQqUserinfo200Response) GetVipStatusOk() (*int32, bool)`

GetVipStatusOk returns a tuple with the VipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipStatus

`func (o *GetSocialQqUserinfo200Response) SetVipStatus(v int32)`

SetVipStatus sets VipStatus field to given value.

### HasVipStatus

`func (o *GetSocialQqUserinfo200Response) HasVipStatus() bool`

HasVipStatus returns a boolean if a field has been set.

### GetVipType

`func (o *GetSocialQqUserinfo200Response) GetVipType() int32`

GetVipType returns the VipType field if non-nil, zero value otherwise.

### GetVipTypeOk

`func (o *GetSocialQqUserinfo200Response) GetVipTypeOk() (*int32, bool)`

GetVipTypeOk returns a tuple with the VipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipType

`func (o *GetSocialQqUserinfo200Response) SetVipType(v int32)`

SetVipType sets VipType field to given value.

### HasVipType

`func (o *GetSocialQqUserinfo200Response) HasVipType() bool`

HasVipType returns a boolean if a field has been set.

### GetYellowDiamondLevel

`func (o *GetSocialQqUserinfo200Response) GetYellowDiamondLevel() int32`

GetYellowDiamondLevel returns the YellowDiamondLevel field if non-nil, zero value otherwise.

### GetYellowDiamondLevelOk

`func (o *GetSocialQqUserinfo200Response) GetYellowDiamondLevelOk() (*int32, bool)`

GetYellowDiamondLevelOk returns a tuple with the YellowDiamondLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYellowDiamondLevel

`func (o *GetSocialQqUserinfo200Response) SetYellowDiamondLevel(v int32)`

SetYellowDiamondLevel sets YellowDiamondLevel field to given value.

### HasYellowDiamondLevel

`func (o *GetSocialQqUserinfo200Response) HasYellowDiamondLevel() bool`

HasYellowDiamondLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



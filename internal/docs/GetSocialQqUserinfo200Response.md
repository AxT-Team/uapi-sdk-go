# GetSocialQqUserinfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qq** | Pointer to **string** | QQ号 | [optional] 
**Nickname** | Pointer to **string** | 用户昵称 | [optional] 
**LongNick** | Pointer to **string** | 个性签名 | [optional] 
**AvatarUrl** | Pointer to **string** | 头像URL | [optional] 
**Age** | Pointer to **int32** | 年龄 | [optional] 
**Sex** | Pointer to **string** | 性别 | [optional] 
**Qid** | Pointer to **string** | QQ个性域名 | [optional] 
**QqLevel** | Pointer to **int32** | QQ等级 | [optional] 
**Location** | Pointer to **string** | 地理位置（省市） | [optional] 
**Email** | Pointer to **string** | QQ邮箱 | [optional] 
**IsVip** | Pointer to **bool** | 是否为VIP用户 | [optional] 
**VipLevel** | Pointer to **int32** | VIP等级 | [optional] 
**RegTime** | Pointer to **string** | 注册时间（ISO 8601格式） | [optional] 
**LastUpdated** | Pointer to **string** | 最后更新时间（ISO 8601格式） | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



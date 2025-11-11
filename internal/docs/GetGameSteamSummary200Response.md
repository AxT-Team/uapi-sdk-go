# GetGameSteamSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | 32x32 像素的小尺寸头像URL。 | [optional] 
**Avatarfull** | Pointer to **string** | 184x184 像素的大尺寸头像URL。 | [optional] 
**Avatarmedium** | Pointer to **string** | 64x64 像素的中等尺寸头像URL。 | [optional] 
**Code** | Pointer to **int32** | 状态码，200代表成功。 | [optional] 
**Communityvisibilitystate** | Pointer to **int32** | 社区资料的可见性状态: 1&#x3D;私密, 3&#x3D;公开。 | [optional] 
**Loccountrycode** | Pointer to **string** | 用户个人资料中设置的国家代码 (ISO 3166-1)，前提是用户已设置并公开。 | [optional] 
**Personaname** | Pointer to **string** | 玩家的当前昵称。 | [optional] 
**Personastate** | Pointer to **int32** | 用户当前的在线状态: 0-离线, 1-在线, 2-忙碌, 3-离开, 4-打盹, 5-想交易, 6-想玩。 | [optional] 
**Primaryclanid** | Pointer to **string** | 玩家设置的主要部落的64位ID。 | [optional] 
**Profilestate** | Pointer to **int32** | 如果用户设置了个人资料，则为1。 | [optional] 
**Profileurl** | Pointer to **string** | 用户的Steam社区个人资料页完整URL。 | [optional] 
**Realname** | Pointer to **string** | 用户的真实姓名，前提是用户已设置并公开。 | [optional] 
**Steamid** | Pointer to **string** | 被查询用户的64位SteamID。 | [optional] 
**Timecreated** | Pointer to **int32** | 账户创建时的Unix时间戳（秒）。 | [optional] 
**TimecreatedStr** | Pointer to **string** | 我们为你格式化好的账户创建时间，更直观。 | [optional] 

## Methods

### NewGetGameSteamSummary200Response

`func NewGetGameSteamSummary200Response() *GetGameSteamSummary200Response`

NewGetGameSteamSummary200Response instantiates a new GetGameSteamSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameSteamSummary200ResponseWithDefaults

`func NewGetGameSteamSummary200ResponseWithDefaults() *GetGameSteamSummary200Response`

NewGetGameSteamSummary200ResponseWithDefaults instantiates a new GetGameSteamSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *GetGameSteamSummary200Response) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetGameSteamSummary200Response) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetGameSteamSummary200Response) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *GetGameSteamSummary200Response) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetAvatarfull

`func (o *GetGameSteamSummary200Response) GetAvatarfull() string`

GetAvatarfull returns the Avatarfull field if non-nil, zero value otherwise.

### GetAvatarfullOk

`func (o *GetGameSteamSummary200Response) GetAvatarfullOk() (*string, bool)`

GetAvatarfullOk returns a tuple with the Avatarfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarfull

`func (o *GetGameSteamSummary200Response) SetAvatarfull(v string)`

SetAvatarfull sets Avatarfull field to given value.

### HasAvatarfull

`func (o *GetGameSteamSummary200Response) HasAvatarfull() bool`

HasAvatarfull returns a boolean if a field has been set.

### GetAvatarmedium

`func (o *GetGameSteamSummary200Response) GetAvatarmedium() string`

GetAvatarmedium returns the Avatarmedium field if non-nil, zero value otherwise.

### GetAvatarmediumOk

`func (o *GetGameSteamSummary200Response) GetAvatarmediumOk() (*string, bool)`

GetAvatarmediumOk returns a tuple with the Avatarmedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarmedium

`func (o *GetGameSteamSummary200Response) SetAvatarmedium(v string)`

SetAvatarmedium sets Avatarmedium field to given value.

### HasAvatarmedium

`func (o *GetGameSteamSummary200Response) HasAvatarmedium() bool`

HasAvatarmedium returns a boolean if a field has been set.

### GetCode

`func (o *GetGameSteamSummary200Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetGameSteamSummary200Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetGameSteamSummary200Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetGameSteamSummary200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCommunityvisibilitystate

`func (o *GetGameSteamSummary200Response) GetCommunityvisibilitystate() int32`

GetCommunityvisibilitystate returns the Communityvisibilitystate field if non-nil, zero value otherwise.

### GetCommunityvisibilitystateOk

`func (o *GetGameSteamSummary200Response) GetCommunityvisibilitystateOk() (*int32, bool)`

GetCommunityvisibilitystateOk returns a tuple with the Communityvisibilitystate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityvisibilitystate

`func (o *GetGameSteamSummary200Response) SetCommunityvisibilitystate(v int32)`

SetCommunityvisibilitystate sets Communityvisibilitystate field to given value.

### HasCommunityvisibilitystate

`func (o *GetGameSteamSummary200Response) HasCommunityvisibilitystate() bool`

HasCommunityvisibilitystate returns a boolean if a field has been set.

### GetLoccountrycode

`func (o *GetGameSteamSummary200Response) GetLoccountrycode() string`

GetLoccountrycode returns the Loccountrycode field if non-nil, zero value otherwise.

### GetLoccountrycodeOk

`func (o *GetGameSteamSummary200Response) GetLoccountrycodeOk() (*string, bool)`

GetLoccountrycodeOk returns a tuple with the Loccountrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoccountrycode

`func (o *GetGameSteamSummary200Response) SetLoccountrycode(v string)`

SetLoccountrycode sets Loccountrycode field to given value.

### HasLoccountrycode

`func (o *GetGameSteamSummary200Response) HasLoccountrycode() bool`

HasLoccountrycode returns a boolean if a field has been set.

### GetPersonaname

`func (o *GetGameSteamSummary200Response) GetPersonaname() string`

GetPersonaname returns the Personaname field if non-nil, zero value otherwise.

### GetPersonanameOk

`func (o *GetGameSteamSummary200Response) GetPersonanameOk() (*string, bool)`

GetPersonanameOk returns a tuple with the Personaname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaname

`func (o *GetGameSteamSummary200Response) SetPersonaname(v string)`

SetPersonaname sets Personaname field to given value.

### HasPersonaname

`func (o *GetGameSteamSummary200Response) HasPersonaname() bool`

HasPersonaname returns a boolean if a field has been set.

### GetPersonastate

`func (o *GetGameSteamSummary200Response) GetPersonastate() int32`

GetPersonastate returns the Personastate field if non-nil, zero value otherwise.

### GetPersonastateOk

`func (o *GetGameSteamSummary200Response) GetPersonastateOk() (*int32, bool)`

GetPersonastateOk returns a tuple with the Personastate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonastate

`func (o *GetGameSteamSummary200Response) SetPersonastate(v int32)`

SetPersonastate sets Personastate field to given value.

### HasPersonastate

`func (o *GetGameSteamSummary200Response) HasPersonastate() bool`

HasPersonastate returns a boolean if a field has been set.

### GetPrimaryclanid

`func (o *GetGameSteamSummary200Response) GetPrimaryclanid() string`

GetPrimaryclanid returns the Primaryclanid field if non-nil, zero value otherwise.

### GetPrimaryclanidOk

`func (o *GetGameSteamSummary200Response) GetPrimaryclanidOk() (*string, bool)`

GetPrimaryclanidOk returns a tuple with the Primaryclanid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryclanid

`func (o *GetGameSteamSummary200Response) SetPrimaryclanid(v string)`

SetPrimaryclanid sets Primaryclanid field to given value.

### HasPrimaryclanid

`func (o *GetGameSteamSummary200Response) HasPrimaryclanid() bool`

HasPrimaryclanid returns a boolean if a field has been set.

### GetProfilestate

`func (o *GetGameSteamSummary200Response) GetProfilestate() int32`

GetProfilestate returns the Profilestate field if non-nil, zero value otherwise.

### GetProfilestateOk

`func (o *GetGameSteamSummary200Response) GetProfilestateOk() (*int32, bool)`

GetProfilestateOk returns a tuple with the Profilestate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilestate

`func (o *GetGameSteamSummary200Response) SetProfilestate(v int32)`

SetProfilestate sets Profilestate field to given value.

### HasProfilestate

`func (o *GetGameSteamSummary200Response) HasProfilestate() bool`

HasProfilestate returns a boolean if a field has been set.

### GetProfileurl

`func (o *GetGameSteamSummary200Response) GetProfileurl() string`

GetProfileurl returns the Profileurl field if non-nil, zero value otherwise.

### GetProfileurlOk

`func (o *GetGameSteamSummary200Response) GetProfileurlOk() (*string, bool)`

GetProfileurlOk returns a tuple with the Profileurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileurl

`func (o *GetGameSteamSummary200Response) SetProfileurl(v string)`

SetProfileurl sets Profileurl field to given value.

### HasProfileurl

`func (o *GetGameSteamSummary200Response) HasProfileurl() bool`

HasProfileurl returns a boolean if a field has been set.

### GetRealname

`func (o *GetGameSteamSummary200Response) GetRealname() string`

GetRealname returns the Realname field if non-nil, zero value otherwise.

### GetRealnameOk

`func (o *GetGameSteamSummary200Response) GetRealnameOk() (*string, bool)`

GetRealnameOk returns a tuple with the Realname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealname

`func (o *GetGameSteamSummary200Response) SetRealname(v string)`

SetRealname sets Realname field to given value.

### HasRealname

`func (o *GetGameSteamSummary200Response) HasRealname() bool`

HasRealname returns a boolean if a field has been set.

### GetSteamid

`func (o *GetGameSteamSummary200Response) GetSteamid() string`

GetSteamid returns the Steamid field if non-nil, zero value otherwise.

### GetSteamidOk

`func (o *GetGameSteamSummary200Response) GetSteamidOk() (*string, bool)`

GetSteamidOk returns a tuple with the Steamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamid

`func (o *GetGameSteamSummary200Response) SetSteamid(v string)`

SetSteamid sets Steamid field to given value.

### HasSteamid

`func (o *GetGameSteamSummary200Response) HasSteamid() bool`

HasSteamid returns a boolean if a field has been set.

### GetTimecreated

`func (o *GetGameSteamSummary200Response) GetTimecreated() int32`

GetTimecreated returns the Timecreated field if non-nil, zero value otherwise.

### GetTimecreatedOk

`func (o *GetGameSteamSummary200Response) GetTimecreatedOk() (*int32, bool)`

GetTimecreatedOk returns a tuple with the Timecreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecreated

`func (o *GetGameSteamSummary200Response) SetTimecreated(v int32)`

SetTimecreated sets Timecreated field to given value.

### HasTimecreated

`func (o *GetGameSteamSummary200Response) HasTimecreated() bool`

HasTimecreated returns a boolean if a field has been set.

### GetTimecreatedStr

`func (o *GetGameSteamSummary200Response) GetTimecreatedStr() string`

GetTimecreatedStr returns the TimecreatedStr field if non-nil, zero value otherwise.

### GetTimecreatedStrOk

`func (o *GetGameSteamSummary200Response) GetTimecreatedStrOk() (*string, bool)`

GetTimecreatedStrOk returns a tuple with the TimecreatedStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecreatedStr

`func (o *GetGameSteamSummary200Response) SetTimecreatedStr(v string)`

SetTimecreatedStr sets TimecreatedStr field to given value.

### HasTimecreatedStr

`func (o *GetGameSteamSummary200Response) HasTimecreatedStr() bool`

HasTimecreatedStr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



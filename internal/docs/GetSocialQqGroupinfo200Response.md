# GetSocialQqGroupinfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | 群号 | [optional] 
**GroupName** | Pointer to **string** | 群名称 | [optional] 
**AvatarUrl** | Pointer to **string** | 群头像URL（标准尺寸100x100） | [optional] 
**Description** | Pointer to **string** | 群描述/简介 | [optional] 
**Tag** | Pointer to **string** | 群标签 | [optional] 
**JoinUrl** | Pointer to **string** | 加群链接（QR码URL） | [optional] 
**LastUpdated** | Pointer to **string** | 最后更新时间（ISO 8601格式） | [optional] 
**MemberCount** | Pointer to **int32** | 当前成员数 | [optional] 
**MaxMemberCount** | Pointer to **int32** | 最大成员数 | [optional] 
**ActiveMemberNum** | Pointer to **int32** | 活跃成员数（可选，部分群有此数据） | [optional] 
**OwnerUin** | Pointer to **string** | 群主QQ号（可选） | [optional] 
**OwnerUid** | Pointer to **string** | 群主UID（可选） | [optional] 
**CreateTime** | Pointer to **int32** | 建群时间戳（Unix时间戳，可选） | [optional] 
**CreateTimeStr** | Pointer to **string** | 建群时间格式化字符串（可选） | [optional] 
**GroupGrade** | Pointer to **int32** | 群等级（可选） | [optional] 
**GroupMemo** | Pointer to **string** | 群公告/简介（可选） | [optional] 
**CertType** | Pointer to **int32** | 认证类型（0&#x3D;未认证，可选） | [optional] 
**CertText** | Pointer to **string** | 认证说明文本（可选） | [optional] 

## Methods

### NewGetSocialQqGroupinfo200Response

`func NewGetSocialQqGroupinfo200Response() *GetSocialQqGroupinfo200Response`

NewGetSocialQqGroupinfo200Response instantiates a new GetSocialQqGroupinfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialQqGroupinfo200ResponseWithDefaults

`func NewGetSocialQqGroupinfo200ResponseWithDefaults() *GetSocialQqGroupinfo200Response`

NewGetSocialQqGroupinfo200ResponseWithDefaults instantiates a new GetSocialQqGroupinfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GetSocialQqGroupinfo200Response) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GetSocialQqGroupinfo200Response) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GetSocialQqGroupinfo200Response) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GetSocialQqGroupinfo200Response) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *GetSocialQqGroupinfo200Response) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetSocialQqGroupinfo200Response) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetSocialQqGroupinfo200Response) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetSocialQqGroupinfo200Response) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *GetSocialQqGroupinfo200Response) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetSocialQqGroupinfo200Response) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetSocialQqGroupinfo200Response) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetSocialQqGroupinfo200Response) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetDescription

`func (o *GetSocialQqGroupinfo200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSocialQqGroupinfo200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSocialQqGroupinfo200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSocialQqGroupinfo200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTag

`func (o *GetSocialQqGroupinfo200Response) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetSocialQqGroupinfo200Response) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetSocialQqGroupinfo200Response) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetSocialQqGroupinfo200Response) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetJoinUrl

`func (o *GetSocialQqGroupinfo200Response) GetJoinUrl() string`

GetJoinUrl returns the JoinUrl field if non-nil, zero value otherwise.

### GetJoinUrlOk

`func (o *GetSocialQqGroupinfo200Response) GetJoinUrlOk() (*string, bool)`

GetJoinUrlOk returns a tuple with the JoinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinUrl

`func (o *GetSocialQqGroupinfo200Response) SetJoinUrl(v string)`

SetJoinUrl sets JoinUrl field to given value.

### HasJoinUrl

`func (o *GetSocialQqGroupinfo200Response) HasJoinUrl() bool`

HasJoinUrl returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetSocialQqGroupinfo200Response) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetSocialQqGroupinfo200Response) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetSocialQqGroupinfo200Response) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetSocialQqGroupinfo200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMemberCount

`func (o *GetSocialQqGroupinfo200Response) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *GetSocialQqGroupinfo200Response) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *GetSocialQqGroupinfo200Response) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *GetSocialQqGroupinfo200Response) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetMaxMemberCount

`func (o *GetSocialQqGroupinfo200Response) GetMaxMemberCount() int32`

GetMaxMemberCount returns the MaxMemberCount field if non-nil, zero value otherwise.

### GetMaxMemberCountOk

`func (o *GetSocialQqGroupinfo200Response) GetMaxMemberCountOk() (*int32, bool)`

GetMaxMemberCountOk returns a tuple with the MaxMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemberCount

`func (o *GetSocialQqGroupinfo200Response) SetMaxMemberCount(v int32)`

SetMaxMemberCount sets MaxMemberCount field to given value.

### HasMaxMemberCount

`func (o *GetSocialQqGroupinfo200Response) HasMaxMemberCount() bool`

HasMaxMemberCount returns a boolean if a field has been set.

### GetActiveMemberNum

`func (o *GetSocialQqGroupinfo200Response) GetActiveMemberNum() int32`

GetActiveMemberNum returns the ActiveMemberNum field if non-nil, zero value otherwise.

### GetActiveMemberNumOk

`func (o *GetSocialQqGroupinfo200Response) GetActiveMemberNumOk() (*int32, bool)`

GetActiveMemberNumOk returns a tuple with the ActiveMemberNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMemberNum

`func (o *GetSocialQqGroupinfo200Response) SetActiveMemberNum(v int32)`

SetActiveMemberNum sets ActiveMemberNum field to given value.

### HasActiveMemberNum

`func (o *GetSocialQqGroupinfo200Response) HasActiveMemberNum() bool`

HasActiveMemberNum returns a boolean if a field has been set.

### GetOwnerUin

`func (o *GetSocialQqGroupinfo200Response) GetOwnerUin() string`

GetOwnerUin returns the OwnerUin field if non-nil, zero value otherwise.

### GetOwnerUinOk

`func (o *GetSocialQqGroupinfo200Response) GetOwnerUinOk() (*string, bool)`

GetOwnerUinOk returns a tuple with the OwnerUin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUin

`func (o *GetSocialQqGroupinfo200Response) SetOwnerUin(v string)`

SetOwnerUin sets OwnerUin field to given value.

### HasOwnerUin

`func (o *GetSocialQqGroupinfo200Response) HasOwnerUin() bool`

HasOwnerUin returns a boolean if a field has been set.

### GetOwnerUid

`func (o *GetSocialQqGroupinfo200Response) GetOwnerUid() string`

GetOwnerUid returns the OwnerUid field if non-nil, zero value otherwise.

### GetOwnerUidOk

`func (o *GetSocialQqGroupinfo200Response) GetOwnerUidOk() (*string, bool)`

GetOwnerUidOk returns a tuple with the OwnerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUid

`func (o *GetSocialQqGroupinfo200Response) SetOwnerUid(v string)`

SetOwnerUid sets OwnerUid field to given value.

### HasOwnerUid

`func (o *GetSocialQqGroupinfo200Response) HasOwnerUid() bool`

HasOwnerUid returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetSocialQqGroupinfo200Response) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetSocialQqGroupinfo200Response) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetSocialQqGroupinfo200Response) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetSocialQqGroupinfo200Response) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreateTimeStr

`func (o *GetSocialQqGroupinfo200Response) GetCreateTimeStr() string`

GetCreateTimeStr returns the CreateTimeStr field if non-nil, zero value otherwise.

### GetCreateTimeStrOk

`func (o *GetSocialQqGroupinfo200Response) GetCreateTimeStrOk() (*string, bool)`

GetCreateTimeStrOk returns a tuple with the CreateTimeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStr

`func (o *GetSocialQqGroupinfo200Response) SetCreateTimeStr(v string)`

SetCreateTimeStr sets CreateTimeStr field to given value.

### HasCreateTimeStr

`func (o *GetSocialQqGroupinfo200Response) HasCreateTimeStr() bool`

HasCreateTimeStr returns a boolean if a field has been set.

### GetGroupGrade

`func (o *GetSocialQqGroupinfo200Response) GetGroupGrade() int32`

GetGroupGrade returns the GroupGrade field if non-nil, zero value otherwise.

### GetGroupGradeOk

`func (o *GetSocialQqGroupinfo200Response) GetGroupGradeOk() (*int32, bool)`

GetGroupGradeOk returns a tuple with the GroupGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGrade

`func (o *GetSocialQqGroupinfo200Response) SetGroupGrade(v int32)`

SetGroupGrade sets GroupGrade field to given value.

### HasGroupGrade

`func (o *GetSocialQqGroupinfo200Response) HasGroupGrade() bool`

HasGroupGrade returns a boolean if a field has been set.

### GetGroupMemo

`func (o *GetSocialQqGroupinfo200Response) GetGroupMemo() string`

GetGroupMemo returns the GroupMemo field if non-nil, zero value otherwise.

### GetGroupMemoOk

`func (o *GetSocialQqGroupinfo200Response) GetGroupMemoOk() (*string, bool)`

GetGroupMemoOk returns a tuple with the GroupMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemo

`func (o *GetSocialQqGroupinfo200Response) SetGroupMemo(v string)`

SetGroupMemo sets GroupMemo field to given value.

### HasGroupMemo

`func (o *GetSocialQqGroupinfo200Response) HasGroupMemo() bool`

HasGroupMemo returns a boolean if a field has been set.

### GetCertType

`func (o *GetSocialQqGroupinfo200Response) GetCertType() int32`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *GetSocialQqGroupinfo200Response) GetCertTypeOk() (*int32, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *GetSocialQqGroupinfo200Response) SetCertType(v int32)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *GetSocialQqGroupinfo200Response) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetCertText

`func (o *GetSocialQqGroupinfo200Response) GetCertText() string`

GetCertText returns the CertText field if non-nil, zero value otherwise.

### GetCertTextOk

`func (o *GetSocialQqGroupinfo200Response) GetCertTextOk() (*string, bool)`

GetCertTextOk returns a tuple with the CertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertText

`func (o *GetSocialQqGroupinfo200Response) SetCertText(v string)`

SetCertText sets CertText field to given value.

### HasCertText

`func (o *GetSocialQqGroupinfo200Response) HasCertText() bool`

HasCertText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



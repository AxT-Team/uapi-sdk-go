# GetSocialBilibiliReplies200ResponseRepliesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rpid** | Pointer to **float32** | 评论的唯一ID (Reply ID)。 | [optional] 
**Oid** | Pointer to **float32** | 评论区对象ID，即视频的aid。 | [optional] 
**Mid** | Pointer to **float32** | 发表评论的用户的mid。 | [optional] 
**Root** | Pointer to **float32** | 根评论的rpid。如果为0，表示这条评论是根评论。 | [optional] 
**Parent** | Pointer to **float32** | 回复的父级评论的rpid。如果为0，表示是根评论。 | [optional] 
**Count** | Pointer to **float32** | 这条评论下的回复（楼中楼）数量。 | [optional] 
**Ctime** | Pointer to **float32** | 评论发送时间的Unix时间戳（秒）。 | [optional] 
**Like** | Pointer to **float32** | 该评论获得的点赞数。 | [optional] 
**Member** | Pointer to [**GetSocialBilibiliReplies200ResponseRepliesInnerMember**](GetSocialBilibiliReplies200ResponseRepliesInnerMember.md) |  | [optional] 
**Content** | Pointer to [**GetSocialBilibiliReplies200ResponseRepliesInnerContent**](GetSocialBilibiliReplies200ResponseRepliesInnerContent.md) |  | [optional] 
**Replies** | Pointer to **[]map[string]interface{}** | 楼中楼回复列表。结构与顶层评论对象一致，但通常此数组为空，需要单独请求。 | [optional] 

## Methods

### NewGetSocialBilibiliReplies200ResponseRepliesInner

`func NewGetSocialBilibiliReplies200ResponseRepliesInner() *GetSocialBilibiliReplies200ResponseRepliesInner`

NewGetSocialBilibiliReplies200ResponseRepliesInner instantiates a new GetSocialBilibiliReplies200ResponseRepliesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialBilibiliReplies200ResponseRepliesInnerWithDefaults

`func NewGetSocialBilibiliReplies200ResponseRepliesInnerWithDefaults() *GetSocialBilibiliReplies200ResponseRepliesInner`

NewGetSocialBilibiliReplies200ResponseRepliesInnerWithDefaults instantiates a new GetSocialBilibiliReplies200ResponseRepliesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRpid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetRpid() float32`

GetRpid returns the Rpid field if non-nil, zero value otherwise.

### GetRpidOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetRpidOk() (*float32, bool)`

GetRpidOk returns a tuple with the Rpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetRpid(v float32)`

SetRpid sets Rpid field to given value.

### HasRpid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasRpid() bool`

HasRpid returns a boolean if a field has been set.

### GetOid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetOid() float32`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetOidOk() (*float32, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetOid(v float32)`

SetOid sets Oid field to given value.

### HasOid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetMid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetMid() float32`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetMidOk() (*float32, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetMid(v float32)`

SetMid sets Mid field to given value.

### HasMid

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetRoot

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetRoot() float32`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetRootOk() (*float32, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetRoot(v float32)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetParent

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetParent() float32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetParentOk() (*float32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetParent(v float32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetCount

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCtime

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetCtime() float32`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetCtimeOk() (*float32, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetCtime(v float32)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetLike

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetLike() float32`

GetLike returns the Like field if non-nil, zero value otherwise.

### GetLikeOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetLikeOk() (*float32, bool)`

GetLikeOk returns a tuple with the Like field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLike

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetLike(v float32)`

SetLike sets Like field to given value.

### HasLike

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasLike() bool`

HasLike returns a boolean if a field has been set.

### GetMember

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetMember() GetSocialBilibiliReplies200ResponseRepliesInnerMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetMemberOk() (*GetSocialBilibiliReplies200ResponseRepliesInnerMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetMember(v GetSocialBilibiliReplies200ResponseRepliesInnerMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetContent

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetContent() GetSocialBilibiliReplies200ResponseRepliesInnerContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetContentOk() (*GetSocialBilibiliReplies200ResponseRepliesInnerContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetContent(v GetSocialBilibiliReplies200ResponseRepliesInnerContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetReplies

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetReplies() []map[string]interface{}`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) GetRepliesOk() (*[]map[string]interface{}, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) SetReplies(v []map[string]interface{})`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *GetSocialBilibiliReplies200ResponseRepliesInner) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



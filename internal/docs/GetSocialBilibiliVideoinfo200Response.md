# GetSocialBilibiliVideoinfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bvid** | Pointer to **string** | 稿件的BV号。 | [optional] 
**Aid** | Pointer to **float32** | 稿件的AV号。 | [optional] 
**Videos** | Pointer to **float32** | 稿件分P总数。如果是单P视频，则为1。 | [optional] 
**Tid** | Pointer to **float32** | 视频所属的子分区 ID。 | [optional] 
**Tname** | Pointer to **string** | 视频所属的子分区名称。 | [optional] 
**Copyright** | Pointer to **float32** | 视频类型。1代表原创，2代表转载。 | [optional] 
**Pic** | Pointer to **string** | 稿件封面图片的URL。这是一个可以直接在网页上展示的链接。 | [optional] 
**Title** | Pointer to **string** | 稿件的标题。 | [optional] 
**Pubdate** | Pointer to **float32** | 稿件发布时间的Unix时间戳（秒）。 | [optional] 
**Ctime** | Pointer to **float32** | 用户投稿时间的Unix时间戳（秒）。 | [optional] 
**Desc** | Pointer to **string** | 视频简介。可能会包含HTML换行符。 | [optional] 
**DescV2** | Pointer to [**[]GetSocialBilibiliVideoinfo200ResponseDescV2Inner**](GetSocialBilibiliVideoinfo200ResponseDescV2Inner.md) | 结构化简介片段。 | [optional] 
**State** | Pointer to **float32** | 视频状态码。 | [optional] 
**Duration** | Pointer to **float32** | 稿件总时长（所有分P累加），单位为秒。 | [optional] 
**Rights** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseRights**](GetSocialBilibiliVideoinfo200ResponseRights.md) |  | [optional] 
**Owner** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseOwner**](GetSocialBilibiliVideoinfo200ResponseOwner.md) |  | [optional] 
**Stat** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseStat**](GetSocialBilibiliVideoinfo200ResponseStat.md) |  | [optional] 
**Dynamic** | Pointer to **string** | 投稿时附带的动态文字。 | [optional] 
**Cid** | Pointer to **float32** | 主分P的 CID（弹幕 ID）。 | [optional] 
**Dimension** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseDimension**](GetSocialBilibiliVideoinfo200ResponseDimension.md) |  | [optional] 
**NoCache** | Pointer to **bool** | 不缓存标记。 | [optional] 
**Pages** | Pointer to [**[]GetSocialBilibiliVideoinfo200ResponsePagesInner**](GetSocialBilibiliVideoinfo200ResponsePagesInner.md) | 视频分P列表。即使是单P视频，该数组也包含一个元素。 | [optional] 
**Subtitle** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseSubtitle**](GetSocialBilibiliVideoinfo200ResponseSubtitle.md) |  | [optional] 
**Staff** | Pointer to [**[]GetSocialBilibiliVideoinfo200ResponseStaffInner**](GetSocialBilibiliVideoinfo200ResponseStaffInner.md) | 联合投稿成员列表。 | [optional] 
**UgcSeason** | Pointer to [**NullableGetSocialBilibiliVideoinfo200ResponseUgcSeason**](GetSocialBilibiliVideoinfo200ResponseUgcSeason.md) |  | [optional] 
**IsChargeableSeason** | Pointer to **bool** | 是否为付费合集。 | [optional] 
**IsStory** | Pointer to **bool** | 是否为剧情类视频。 | [optional] 
**HonorReply** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseHonorReply**](GetSocialBilibiliVideoinfo200ResponseHonorReply.md) |  | [optional] 

## Methods

### NewGetSocialBilibiliVideoinfo200Response

`func NewGetSocialBilibiliVideoinfo200Response() *GetSocialBilibiliVideoinfo200Response`

NewGetSocialBilibiliVideoinfo200Response instantiates a new GetSocialBilibiliVideoinfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialBilibiliVideoinfo200ResponseWithDefaults

`func NewGetSocialBilibiliVideoinfo200ResponseWithDefaults() *GetSocialBilibiliVideoinfo200Response`

NewGetSocialBilibiliVideoinfo200ResponseWithDefaults instantiates a new GetSocialBilibiliVideoinfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBvid

`func (o *GetSocialBilibiliVideoinfo200Response) GetBvid() string`

GetBvid returns the Bvid field if non-nil, zero value otherwise.

### GetBvidOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetBvidOk() (*string, bool)`

GetBvidOk returns a tuple with the Bvid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBvid

`func (o *GetSocialBilibiliVideoinfo200Response) SetBvid(v string)`

SetBvid sets Bvid field to given value.

### HasBvid

`func (o *GetSocialBilibiliVideoinfo200Response) HasBvid() bool`

HasBvid returns a boolean if a field has been set.

### GetAid

`func (o *GetSocialBilibiliVideoinfo200Response) GetAid() float32`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetAidOk() (*float32, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *GetSocialBilibiliVideoinfo200Response) SetAid(v float32)`

SetAid sets Aid field to given value.

### HasAid

`func (o *GetSocialBilibiliVideoinfo200Response) HasAid() bool`

HasAid returns a boolean if a field has been set.

### GetVideos

`func (o *GetSocialBilibiliVideoinfo200Response) GetVideos() float32`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetVideosOk() (*float32, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *GetSocialBilibiliVideoinfo200Response) SetVideos(v float32)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *GetSocialBilibiliVideoinfo200Response) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetTid

`func (o *GetSocialBilibiliVideoinfo200Response) GetTid() float32`

GetTid returns the Tid field if non-nil, zero value otherwise.

### GetTidOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetTidOk() (*float32, bool)`

GetTidOk returns a tuple with the Tid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTid

`func (o *GetSocialBilibiliVideoinfo200Response) SetTid(v float32)`

SetTid sets Tid field to given value.

### HasTid

`func (o *GetSocialBilibiliVideoinfo200Response) HasTid() bool`

HasTid returns a boolean if a field has been set.

### GetTname

`func (o *GetSocialBilibiliVideoinfo200Response) GetTname() string`

GetTname returns the Tname field if non-nil, zero value otherwise.

### GetTnameOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetTnameOk() (*string, bool)`

GetTnameOk returns a tuple with the Tname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTname

`func (o *GetSocialBilibiliVideoinfo200Response) SetTname(v string)`

SetTname sets Tname field to given value.

### HasTname

`func (o *GetSocialBilibiliVideoinfo200Response) HasTname() bool`

HasTname returns a boolean if a field has been set.

### GetCopyright

`func (o *GetSocialBilibiliVideoinfo200Response) GetCopyright() float32`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetCopyrightOk() (*float32, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *GetSocialBilibiliVideoinfo200Response) SetCopyright(v float32)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *GetSocialBilibiliVideoinfo200Response) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### GetPic

`func (o *GetSocialBilibiliVideoinfo200Response) GetPic() string`

GetPic returns the Pic field if non-nil, zero value otherwise.

### GetPicOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetPicOk() (*string, bool)`

GetPicOk returns a tuple with the Pic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPic

`func (o *GetSocialBilibiliVideoinfo200Response) SetPic(v string)`

SetPic sets Pic field to given value.

### HasPic

`func (o *GetSocialBilibiliVideoinfo200Response) HasPic() bool`

HasPic returns a boolean if a field has been set.

### GetTitle

`func (o *GetSocialBilibiliVideoinfo200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetSocialBilibiliVideoinfo200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetSocialBilibiliVideoinfo200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPubdate

`func (o *GetSocialBilibiliVideoinfo200Response) GetPubdate() float32`

GetPubdate returns the Pubdate field if non-nil, zero value otherwise.

### GetPubdateOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetPubdateOk() (*float32, bool)`

GetPubdateOk returns a tuple with the Pubdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubdate

`func (o *GetSocialBilibiliVideoinfo200Response) SetPubdate(v float32)`

SetPubdate sets Pubdate field to given value.

### HasPubdate

`func (o *GetSocialBilibiliVideoinfo200Response) HasPubdate() bool`

HasPubdate returns a boolean if a field has been set.

### GetCtime

`func (o *GetSocialBilibiliVideoinfo200Response) GetCtime() float32`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetCtimeOk() (*float32, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *GetSocialBilibiliVideoinfo200Response) SetCtime(v float32)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *GetSocialBilibiliVideoinfo200Response) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetDesc

`func (o *GetSocialBilibiliVideoinfo200Response) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *GetSocialBilibiliVideoinfo200Response) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *GetSocialBilibiliVideoinfo200Response) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetDescV2

`func (o *GetSocialBilibiliVideoinfo200Response) GetDescV2() []GetSocialBilibiliVideoinfo200ResponseDescV2Inner`

GetDescV2 returns the DescV2 field if non-nil, zero value otherwise.

### GetDescV2Ok

`func (o *GetSocialBilibiliVideoinfo200Response) GetDescV2Ok() (*[]GetSocialBilibiliVideoinfo200ResponseDescV2Inner, bool)`

GetDescV2Ok returns a tuple with the DescV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescV2

`func (o *GetSocialBilibiliVideoinfo200Response) SetDescV2(v []GetSocialBilibiliVideoinfo200ResponseDescV2Inner)`

SetDescV2 sets DescV2 field to given value.

### HasDescV2

`func (o *GetSocialBilibiliVideoinfo200Response) HasDescV2() bool`

HasDescV2 returns a boolean if a field has been set.

### GetState

`func (o *GetSocialBilibiliVideoinfo200Response) GetState() float32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetStateOk() (*float32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSocialBilibiliVideoinfo200Response) SetState(v float32)`

SetState sets State field to given value.

### HasState

`func (o *GetSocialBilibiliVideoinfo200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDuration

`func (o *GetSocialBilibiliVideoinfo200Response) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetSocialBilibiliVideoinfo200Response) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetSocialBilibiliVideoinfo200Response) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetRights

`func (o *GetSocialBilibiliVideoinfo200Response) GetRights() GetSocialBilibiliVideoinfo200ResponseRights`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetRightsOk() (*GetSocialBilibiliVideoinfo200ResponseRights, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *GetSocialBilibiliVideoinfo200Response) SetRights(v GetSocialBilibiliVideoinfo200ResponseRights)`

SetRights sets Rights field to given value.

### HasRights

`func (o *GetSocialBilibiliVideoinfo200Response) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetOwner

`func (o *GetSocialBilibiliVideoinfo200Response) GetOwner() GetSocialBilibiliVideoinfo200ResponseOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetOwnerOk() (*GetSocialBilibiliVideoinfo200ResponseOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetSocialBilibiliVideoinfo200Response) SetOwner(v GetSocialBilibiliVideoinfo200ResponseOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetSocialBilibiliVideoinfo200Response) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetStat

`func (o *GetSocialBilibiliVideoinfo200Response) GetStat() GetSocialBilibiliVideoinfo200ResponseStat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetStatOk() (*GetSocialBilibiliVideoinfo200ResponseStat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *GetSocialBilibiliVideoinfo200Response) SetStat(v GetSocialBilibiliVideoinfo200ResponseStat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *GetSocialBilibiliVideoinfo200Response) HasStat() bool`

HasStat returns a boolean if a field has been set.

### GetDynamic

`func (o *GetSocialBilibiliVideoinfo200Response) GetDynamic() string`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetDynamicOk() (*string, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *GetSocialBilibiliVideoinfo200Response) SetDynamic(v string)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *GetSocialBilibiliVideoinfo200Response) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetCid

`func (o *GetSocialBilibiliVideoinfo200Response) GetCid() float32`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetCidOk() (*float32, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *GetSocialBilibiliVideoinfo200Response) SetCid(v float32)`

SetCid sets Cid field to given value.

### HasCid

`func (o *GetSocialBilibiliVideoinfo200Response) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetDimension

`func (o *GetSocialBilibiliVideoinfo200Response) GetDimension() GetSocialBilibiliVideoinfo200ResponseDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetDimensionOk() (*GetSocialBilibiliVideoinfo200ResponseDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *GetSocialBilibiliVideoinfo200Response) SetDimension(v GetSocialBilibiliVideoinfo200ResponseDimension)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *GetSocialBilibiliVideoinfo200Response) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetNoCache

`func (o *GetSocialBilibiliVideoinfo200Response) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *GetSocialBilibiliVideoinfo200Response) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *GetSocialBilibiliVideoinfo200Response) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetPages

`func (o *GetSocialBilibiliVideoinfo200Response) GetPages() []GetSocialBilibiliVideoinfo200ResponsePagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetPagesOk() (*[]GetSocialBilibiliVideoinfo200ResponsePagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetSocialBilibiliVideoinfo200Response) SetPages(v []GetSocialBilibiliVideoinfo200ResponsePagesInner)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetSocialBilibiliVideoinfo200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetSubtitle

`func (o *GetSocialBilibiliVideoinfo200Response) GetSubtitle() GetSocialBilibiliVideoinfo200ResponseSubtitle`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetSubtitleOk() (*GetSocialBilibiliVideoinfo200ResponseSubtitle, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *GetSocialBilibiliVideoinfo200Response) SetSubtitle(v GetSocialBilibiliVideoinfo200ResponseSubtitle)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *GetSocialBilibiliVideoinfo200Response) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetStaff

`func (o *GetSocialBilibiliVideoinfo200Response) GetStaff() []GetSocialBilibiliVideoinfo200ResponseStaffInner`

GetStaff returns the Staff field if non-nil, zero value otherwise.

### GetStaffOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetStaffOk() (*[]GetSocialBilibiliVideoinfo200ResponseStaffInner, bool)`

GetStaffOk returns a tuple with the Staff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaff

`func (o *GetSocialBilibiliVideoinfo200Response) SetStaff(v []GetSocialBilibiliVideoinfo200ResponseStaffInner)`

SetStaff sets Staff field to given value.

### HasStaff

`func (o *GetSocialBilibiliVideoinfo200Response) HasStaff() bool`

HasStaff returns a boolean if a field has been set.

### GetUgcSeason

`func (o *GetSocialBilibiliVideoinfo200Response) GetUgcSeason() GetSocialBilibiliVideoinfo200ResponseUgcSeason`

GetUgcSeason returns the UgcSeason field if non-nil, zero value otherwise.

### GetUgcSeasonOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetUgcSeasonOk() (*GetSocialBilibiliVideoinfo200ResponseUgcSeason, bool)`

GetUgcSeasonOk returns a tuple with the UgcSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUgcSeason

`func (o *GetSocialBilibiliVideoinfo200Response) SetUgcSeason(v GetSocialBilibiliVideoinfo200ResponseUgcSeason)`

SetUgcSeason sets UgcSeason field to given value.

### HasUgcSeason

`func (o *GetSocialBilibiliVideoinfo200Response) HasUgcSeason() bool`

HasUgcSeason returns a boolean if a field has been set.

### SetUgcSeasonNil

`func (o *GetSocialBilibiliVideoinfo200Response) SetUgcSeasonNil(b bool)`

 SetUgcSeasonNil sets the value for UgcSeason to be an explicit nil

### UnsetUgcSeason
`func (o *GetSocialBilibiliVideoinfo200Response) UnsetUgcSeason()`

UnsetUgcSeason ensures that no value is present for UgcSeason, not even an explicit nil
### GetIsChargeableSeason

`func (o *GetSocialBilibiliVideoinfo200Response) GetIsChargeableSeason() bool`

GetIsChargeableSeason returns the IsChargeableSeason field if non-nil, zero value otherwise.

### GetIsChargeableSeasonOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetIsChargeableSeasonOk() (*bool, bool)`

GetIsChargeableSeasonOk returns a tuple with the IsChargeableSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChargeableSeason

`func (o *GetSocialBilibiliVideoinfo200Response) SetIsChargeableSeason(v bool)`

SetIsChargeableSeason sets IsChargeableSeason field to given value.

### HasIsChargeableSeason

`func (o *GetSocialBilibiliVideoinfo200Response) HasIsChargeableSeason() bool`

HasIsChargeableSeason returns a boolean if a field has been set.

### GetIsStory

`func (o *GetSocialBilibiliVideoinfo200Response) GetIsStory() bool`

GetIsStory returns the IsStory field if non-nil, zero value otherwise.

### GetIsStoryOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetIsStoryOk() (*bool, bool)`

GetIsStoryOk returns a tuple with the IsStory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStory

`func (o *GetSocialBilibiliVideoinfo200Response) SetIsStory(v bool)`

SetIsStory sets IsStory field to given value.

### HasIsStory

`func (o *GetSocialBilibiliVideoinfo200Response) HasIsStory() bool`

HasIsStory returns a boolean if a field has been set.

### GetHonorReply

`func (o *GetSocialBilibiliVideoinfo200Response) GetHonorReply() GetSocialBilibiliVideoinfo200ResponseHonorReply`

GetHonorReply returns the HonorReply field if non-nil, zero value otherwise.

### GetHonorReplyOk

`func (o *GetSocialBilibiliVideoinfo200Response) GetHonorReplyOk() (*GetSocialBilibiliVideoinfo200ResponseHonorReply, bool)`

GetHonorReplyOk returns a tuple with the HonorReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorReply

`func (o *GetSocialBilibiliVideoinfo200Response) SetHonorReply(v GetSocialBilibiliVideoinfo200ResponseHonorReply)`

SetHonorReply sets HonorReply field to given value.

### HasHonorReply

`func (o *GetSocialBilibiliVideoinfo200Response) HasHonorReply() bool`

HasHonorReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



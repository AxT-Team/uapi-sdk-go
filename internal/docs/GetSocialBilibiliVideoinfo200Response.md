# GetSocialBilibiliVideoinfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bvid** | Pointer to **string** | 稿件的BV号。 | [optional] 
**Aid** | Pointer to **float32** | 稿件的AV号。 | [optional] 
**Videos** | Pointer to **float32** | 稿件分P总数。如果是单P视频，则为1。 | [optional] 
**Tname** | Pointer to **string** | 视频所属的子分区名称。 | [optional] 
**Copyright** | Pointer to **float32** | 视频类型。1代表原创，2代表转载。 | [optional] 
**Pic** | Pointer to **string** | 稿件封面图片的URL。这是一个可以直接在网页上展示的链接。 | [optional] 
**Title** | Pointer to **string** | 稿件的标题。 | [optional] 
**Pubdate** | Pointer to **float32** | 稿件发布时间的Unix时间戳（秒）。 | [optional] 
**Ctime** | Pointer to **float32** | 用户投稿时间的Unix时间戳（秒）。 | [optional] 
**Desc** | Pointer to **string** | 视频简介。可能会包含HTML换行符。 | [optional] 
**Duration** | Pointer to **float32** | 稿件总时长（所有分P累加），单位为秒。 | [optional] 
**Owner** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseOwner**](GetSocialBilibiliVideoinfo200ResponseOwner.md) |  | [optional] 
**Stat** | Pointer to [**GetSocialBilibiliVideoinfo200ResponseStat**](GetSocialBilibiliVideoinfo200ResponseStat.md) |  | [optional] 
**Pages** | Pointer to [**[]GetSocialBilibiliVideoinfo200ResponsePagesInner**](GetSocialBilibiliVideoinfo200ResponsePagesInner.md) | 视频分P列表。即使是单P视频，该数组也包含一个元素。 | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



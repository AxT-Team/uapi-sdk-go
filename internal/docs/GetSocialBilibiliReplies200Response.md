# GetSocialBilibiliReplies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**GetSocialBilibiliReplies200ResponsePage**](GetSocialBilibiliReplies200ResponsePage.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | 评论区配置。不同视频或不同权限下可能为 null。 | [optional] 
**Hots** | Pointer to **[]map[string]interface{}** | 热门评论列表。结构与 &#x60;replies&#x60; 中的对象一致。如果当前页是第一页，且有热门评论，则此数组非空。 | [optional] 
**Replies** | Pointer to [**[]GetSocialBilibiliReplies200ResponseRepliesInner**](GetSocialBilibiliReplies200ResponseRepliesInner.md) | 当前页的评论列表。 | [optional] 
**Upper** | Pointer to **map[string]interface{}** | UP 主相关信息。无数据时为 null。 | [optional] 
**Top** | Pointer to **map[string]interface{}** | 置顶评论信息。没有置顶评论时为 null。 | [optional] 
**Notice** | Pointer to **map[string]interface{}** | 评论区公告信息。没有公告时为 null。 | [optional] 
**Vote** | Pointer to **float32** | 评论区投票相关状态值。没有投票时通常为 0。 | [optional] 
**Folder** | Pointer to **map[string]interface{}** | 评论折叠相关信息。没有数据时为 null。 | [optional] 
**Control** | Pointer to **map[string]interface{}** | 评论区控制信息。没有数据时为 null。 | [optional] 
**Cursor** | Pointer to **map[string]interface{}** | 游标翻页信息。部分场景下为 null。 | [optional] 

## Methods

### NewGetSocialBilibiliReplies200Response

`func NewGetSocialBilibiliReplies200Response() *GetSocialBilibiliReplies200Response`

NewGetSocialBilibiliReplies200Response instantiates a new GetSocialBilibiliReplies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialBilibiliReplies200ResponseWithDefaults

`func NewGetSocialBilibiliReplies200ResponseWithDefaults() *GetSocialBilibiliReplies200Response`

NewGetSocialBilibiliReplies200ResponseWithDefaults instantiates a new GetSocialBilibiliReplies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetSocialBilibiliReplies200Response) GetPage() GetSocialBilibiliReplies200ResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSocialBilibiliReplies200Response) GetPageOk() (*GetSocialBilibiliReplies200ResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSocialBilibiliReplies200Response) SetPage(v GetSocialBilibiliReplies200ResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSocialBilibiliReplies200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetConfig

`func (o *GetSocialBilibiliReplies200Response) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetSocialBilibiliReplies200Response) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetSocialBilibiliReplies200Response) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetSocialBilibiliReplies200Response) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetSocialBilibiliReplies200Response) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetSocialBilibiliReplies200Response) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetHots

`func (o *GetSocialBilibiliReplies200Response) GetHots() []map[string]interface{}`

GetHots returns the Hots field if non-nil, zero value otherwise.

### GetHotsOk

`func (o *GetSocialBilibiliReplies200Response) GetHotsOk() (*[]map[string]interface{}, bool)`

GetHotsOk returns a tuple with the Hots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHots

`func (o *GetSocialBilibiliReplies200Response) SetHots(v []map[string]interface{})`

SetHots sets Hots field to given value.

### HasHots

`func (o *GetSocialBilibiliReplies200Response) HasHots() bool`

HasHots returns a boolean if a field has been set.

### SetHotsNil

`func (o *GetSocialBilibiliReplies200Response) SetHotsNil(b bool)`

 SetHotsNil sets the value for Hots to be an explicit nil

### UnsetHots
`func (o *GetSocialBilibiliReplies200Response) UnsetHots()`

UnsetHots ensures that no value is present for Hots, not even an explicit nil
### GetReplies

`func (o *GetSocialBilibiliReplies200Response) GetReplies() []GetSocialBilibiliReplies200ResponseRepliesInner`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *GetSocialBilibiliReplies200Response) GetRepliesOk() (*[]GetSocialBilibiliReplies200ResponseRepliesInner, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *GetSocialBilibiliReplies200Response) SetReplies(v []GetSocialBilibiliReplies200ResponseRepliesInner)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *GetSocialBilibiliReplies200Response) HasReplies() bool`

HasReplies returns a boolean if a field has been set.

### GetUpper

`func (o *GetSocialBilibiliReplies200Response) GetUpper() map[string]interface{}`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *GetSocialBilibiliReplies200Response) GetUpperOk() (*map[string]interface{}, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *GetSocialBilibiliReplies200Response) SetUpper(v map[string]interface{})`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *GetSocialBilibiliReplies200Response) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### SetUpperNil

`func (o *GetSocialBilibiliReplies200Response) SetUpperNil(b bool)`

 SetUpperNil sets the value for Upper to be an explicit nil

### UnsetUpper
`func (o *GetSocialBilibiliReplies200Response) UnsetUpper()`

UnsetUpper ensures that no value is present for Upper, not even an explicit nil
### GetTop

`func (o *GetSocialBilibiliReplies200Response) GetTop() map[string]interface{}`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *GetSocialBilibiliReplies200Response) GetTopOk() (*map[string]interface{}, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *GetSocialBilibiliReplies200Response) SetTop(v map[string]interface{})`

SetTop sets Top field to given value.

### HasTop

`func (o *GetSocialBilibiliReplies200Response) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTopNil

`func (o *GetSocialBilibiliReplies200Response) SetTopNil(b bool)`

 SetTopNil sets the value for Top to be an explicit nil

### UnsetTop
`func (o *GetSocialBilibiliReplies200Response) UnsetTop()`

UnsetTop ensures that no value is present for Top, not even an explicit nil
### GetNotice

`func (o *GetSocialBilibiliReplies200Response) GetNotice() map[string]interface{}`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *GetSocialBilibiliReplies200Response) GetNoticeOk() (*map[string]interface{}, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *GetSocialBilibiliReplies200Response) SetNotice(v map[string]interface{})`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *GetSocialBilibiliReplies200Response) HasNotice() bool`

HasNotice returns a boolean if a field has been set.

### SetNoticeNil

`func (o *GetSocialBilibiliReplies200Response) SetNoticeNil(b bool)`

 SetNoticeNil sets the value for Notice to be an explicit nil

### UnsetNotice
`func (o *GetSocialBilibiliReplies200Response) UnsetNotice()`

UnsetNotice ensures that no value is present for Notice, not even an explicit nil
### GetVote

`func (o *GetSocialBilibiliReplies200Response) GetVote() float32`

GetVote returns the Vote field if non-nil, zero value otherwise.

### GetVoteOk

`func (o *GetSocialBilibiliReplies200Response) GetVoteOk() (*float32, bool)`

GetVoteOk returns a tuple with the Vote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVote

`func (o *GetSocialBilibiliReplies200Response) SetVote(v float32)`

SetVote sets Vote field to given value.

### HasVote

`func (o *GetSocialBilibiliReplies200Response) HasVote() bool`

HasVote returns a boolean if a field has been set.

### GetFolder

`func (o *GetSocialBilibiliReplies200Response) GetFolder() map[string]interface{}`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *GetSocialBilibiliReplies200Response) GetFolderOk() (*map[string]interface{}, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *GetSocialBilibiliReplies200Response) SetFolder(v map[string]interface{})`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *GetSocialBilibiliReplies200Response) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *GetSocialBilibiliReplies200Response) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *GetSocialBilibiliReplies200Response) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetControl

`func (o *GetSocialBilibiliReplies200Response) GetControl() map[string]interface{}`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *GetSocialBilibiliReplies200Response) GetControlOk() (*map[string]interface{}, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *GetSocialBilibiliReplies200Response) SetControl(v map[string]interface{})`

SetControl sets Control field to given value.

### HasControl

`func (o *GetSocialBilibiliReplies200Response) HasControl() bool`

HasControl returns a boolean if a field has been set.

### SetControlNil

`func (o *GetSocialBilibiliReplies200Response) SetControlNil(b bool)`

 SetControlNil sets the value for Control to be an explicit nil

### UnsetControl
`func (o *GetSocialBilibiliReplies200Response) UnsetControl()`

UnsetControl ensures that no value is present for Control, not even an explicit nil
### GetCursor

`func (o *GetSocialBilibiliReplies200Response) GetCursor() map[string]interface{}`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetSocialBilibiliReplies200Response) GetCursorOk() (*map[string]interface{}, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetSocialBilibiliReplies200Response) SetCursor(v map[string]interface{})`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetSocialBilibiliReplies200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *GetSocialBilibiliReplies200Response) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *GetSocialBilibiliReplies200Response) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



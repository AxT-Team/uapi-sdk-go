# GetSocialBilibiliReplies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**GetSocialBilibiliReplies200ResponsePage**](GetSocialBilibiliReplies200ResponsePage.md) |  | [optional] 
**Hots** | Pointer to **[]map[string]interface{}** | 热门评论列表。结构与 &#x60;replies&#x60; 中的对象一致。如果当前页是第一页，且有热门评论，则此数组非空。 | [optional] 
**Replies** | Pointer to [**[]GetSocialBilibiliReplies200ResponseRepliesInner**](GetSocialBilibiliReplies200ResponseRepliesInner.md) | 当前页的评论列表。 | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



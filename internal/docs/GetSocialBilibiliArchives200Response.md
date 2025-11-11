# GetSocialBilibiliArchives200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | 总稿件数 | [optional] 
**Page** | Pointer to **int32** | 当前页码 | [optional] 
**Size** | Pointer to **int32** | 每页数量 | [optional] 
**Videos** | Pointer to [**[]GetSocialBilibiliArchives200ResponseVideosInner**](GetSocialBilibiliArchives200ResponseVideosInner.md) | 视频列表 | [optional] 

## Methods

### NewGetSocialBilibiliArchives200Response

`func NewGetSocialBilibiliArchives200Response() *GetSocialBilibiliArchives200Response`

NewGetSocialBilibiliArchives200Response instantiates a new GetSocialBilibiliArchives200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSocialBilibiliArchives200ResponseWithDefaults

`func NewGetSocialBilibiliArchives200ResponseWithDefaults() *GetSocialBilibiliArchives200Response`

NewGetSocialBilibiliArchives200ResponseWithDefaults instantiates a new GetSocialBilibiliArchives200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetSocialBilibiliArchives200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetSocialBilibiliArchives200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetSocialBilibiliArchives200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetSocialBilibiliArchives200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *GetSocialBilibiliArchives200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSocialBilibiliArchives200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSocialBilibiliArchives200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSocialBilibiliArchives200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *GetSocialBilibiliArchives200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetSocialBilibiliArchives200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetSocialBilibiliArchives200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetSocialBilibiliArchives200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVideos

`func (o *GetSocialBilibiliArchives200Response) GetVideos() []GetSocialBilibiliArchives200ResponseVideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *GetSocialBilibiliArchives200Response) GetVideosOk() (*[]GetSocialBilibiliArchives200ResponseVideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *GetSocialBilibiliArchives200Response) SetVideos(v []GetSocialBilibiliArchives200ResponseVideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *GetSocialBilibiliArchives200Response) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



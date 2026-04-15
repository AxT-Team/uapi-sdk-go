# GetGithubUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** | GitHub 登录名。 | [optional] 
**Name** | Pointer to **string** | 用户公开显示的名称。 | [optional] 
**Bio** | Pointer to **string** | 用户个人简介。 | [optional] 
**Company** | Pointer to **string** | 用户填写的公司或组织信息。 | [optional] 
**Location** | Pointer to **string** | 用户公开展示的地理位置。 | [optional] 
**Blog** | Pointer to **string** | 用户填写的网站或博客地址。 | [optional] 
**TwitterUsername** | Pointer to **string** | 用户绑定的 X（Twitter）用户名。 | [optional] 
**Email** | Pointer to **string** | 用户公开可见的邮箱地址。 | [optional] 
**HtmlUrl** | Pointer to **string** | GitHub 个人主页链接。 | [optional] 
**AvatarUrl** | Pointer to **string** | 用户头像图片链接。 | [optional] 
**Type** | Pointer to **string** | 账号类型，常见值为 User 或 Organization。 | [optional] 
**PublicRepos** | Pointer to **int32** | 公开仓库数量。 | [optional] 
**PublicGists** | Pointer to **int32** | 公开 Gist 数量。 | [optional] 
**Followers** | Pointer to **int32** | 关注该用户的人数。 | [optional] 
**Following** | Pointer to **int32** | 该用户正在关注的人数。 | [optional] 
**CreatedAt** | Pointer to **time.Time** | GitHub 账号创建时间（ISO 8601）。 | [optional] 
**UpdatedAt** | Pointer to **time.Time** | 用户资料最近更新时间（ISO 8601）。 | [optional] 
**Organizations** | Pointer to [**[]GetGithubUser200ResponseOrganizationsInner**](GetGithubUser200ResponseOrganizationsInner.md) | 用户公开加入的组织列表 | [optional] 
**Activity** | Pointer to [**GetGithubUser200ResponseActivity**](GetGithubUser200ResponseActivity.md) |  | [optional] 

## Methods

### NewGetGithubUser200Response

`func NewGetGithubUser200Response() *GetGithubUser200Response`

NewGetGithubUser200Response instantiates a new GetGithubUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGithubUser200ResponseWithDefaults

`func NewGetGithubUser200ResponseWithDefaults() *GetGithubUser200Response`

NewGetGithubUser200ResponseWithDefaults instantiates a new GetGithubUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *GetGithubUser200Response) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GetGithubUser200Response) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GetGithubUser200Response) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *GetGithubUser200Response) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *GetGithubUser200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGithubUser200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGithubUser200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGithubUser200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBio

`func (o *GetGithubUser200Response) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *GetGithubUser200Response) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *GetGithubUser200Response) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *GetGithubUser200Response) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetCompany

`func (o *GetGithubUser200Response) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GetGithubUser200Response) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GetGithubUser200Response) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GetGithubUser200Response) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *GetGithubUser200Response) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetGithubUser200Response) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetGithubUser200Response) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetGithubUser200Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBlog

`func (o *GetGithubUser200Response) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *GetGithubUser200Response) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *GetGithubUser200Response) SetBlog(v string)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *GetGithubUser200Response) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *GetGithubUser200Response) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *GetGithubUser200Response) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *GetGithubUser200Response) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *GetGithubUser200Response) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### GetEmail

`func (o *GetGithubUser200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetGithubUser200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetGithubUser200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetGithubUser200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GetGithubUser200Response) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GetGithubUser200Response) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GetGithubUser200Response) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GetGithubUser200Response) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *GetGithubUser200Response) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetGithubUser200Response) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetGithubUser200Response) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetGithubUser200Response) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetType

`func (o *GetGithubUser200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGithubUser200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGithubUser200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetGithubUser200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPublicRepos

`func (o *GetGithubUser200Response) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *GetGithubUser200Response) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *GetGithubUser200Response) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.

### HasPublicRepos

`func (o *GetGithubUser200Response) HasPublicRepos() bool`

HasPublicRepos returns a boolean if a field has been set.

### GetPublicGists

`func (o *GetGithubUser200Response) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *GetGithubUser200Response) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *GetGithubUser200Response) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.

### HasPublicGists

`func (o *GetGithubUser200Response) HasPublicGists() bool`

HasPublicGists returns a boolean if a field has been set.

### GetFollowers

`func (o *GetGithubUser200Response) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *GetGithubUser200Response) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *GetGithubUser200Response) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *GetGithubUser200Response) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFollowing

`func (o *GetGithubUser200Response) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *GetGithubUser200Response) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *GetGithubUser200Response) SetFollowing(v int32)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *GetGithubUser200Response) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetGithubUser200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetGithubUser200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetGithubUser200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetGithubUser200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetGithubUser200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetGithubUser200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetGithubUser200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetGithubUser200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizations

`func (o *GetGithubUser200Response) GetOrganizations() []GetGithubUser200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GetGithubUser200Response) GetOrganizationsOk() (*[]GetGithubUser200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GetGithubUser200Response) SetOrganizations(v []GetGithubUser200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *GetGithubUser200Response) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetActivity

`func (o *GetGithubUser200Response) GetActivity() GetGithubUser200ResponseActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *GetGithubUser200Response) GetActivityOk() (*GetGithubUser200ResponseActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *GetGithubUser200Response) SetActivity(v GetGithubUser200ResponseActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *GetGithubUser200Response) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



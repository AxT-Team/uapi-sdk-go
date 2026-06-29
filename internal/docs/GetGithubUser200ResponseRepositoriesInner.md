# GetGithubUser200ResponseRepositoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | 该仓库是否已归档。 | [optional] 
**Description** | Pointer to **string** | 仓库简介。 | [optional] 
**Fork** | Pointer to **bool** | 该仓库是否为 fork。 | [optional] 
**Forks** | Pointer to **int32** | Fork 数量。 | [optional] 
**FullName** | Pointer to **string** | 包含所有者的完整仓库名。 | [optional] 
**Homepage** | Pointer to **string** | 仓库填写的官网地址。 | [optional] 
**HtmlUrl** | Pointer to **string** | 仓库主页链接。 | [optional] 
**Language** | Pointer to **string** | 仓库主语言。 | [optional] 
**Name** | Pointer to **string** | 仓库名。 | [optional] 
**PushedAt** | Pointer to **time.Time** | 最近一次 push 时间。 | [optional] 
**Stargazers** | Pointer to **int32** | Star 数量。 | [optional] 
**UpdatedAt** | Pointer to **time.Time** | 最近一次资料更新时间。 | [optional] 
**Visibility** | Pointer to **string** | 仓库可见性。 | [optional] 

## Methods

### NewGetGithubUser200ResponseRepositoriesInner

`func NewGetGithubUser200ResponseRepositoriesInner() *GetGithubUser200ResponseRepositoriesInner`

NewGetGithubUser200ResponseRepositoriesInner instantiates a new GetGithubUser200ResponseRepositoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGithubUser200ResponseRepositoriesInnerWithDefaults

`func NewGetGithubUser200ResponseRepositoriesInnerWithDefaults() *GetGithubUser200ResponseRepositoriesInner`

NewGetGithubUser200ResponseRepositoriesInnerWithDefaults instantiates a new GetGithubUser200ResponseRepositoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *GetGithubUser200ResponseRepositoriesInner) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *GetGithubUser200ResponseRepositoriesInner) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *GetGithubUser200ResponseRepositoriesInner) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDescription

`func (o *GetGithubUser200ResponseRepositoriesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGithubUser200ResponseRepositoriesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGithubUser200ResponseRepositoriesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFork

`func (o *GetGithubUser200ResponseRepositoriesInner) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *GetGithubUser200ResponseRepositoriesInner) SetFork(v bool)`

SetFork sets Fork field to given value.

### HasFork

`func (o *GetGithubUser200ResponseRepositoriesInner) HasFork() bool`

HasFork returns a boolean if a field has been set.

### GetForks

`func (o *GetGithubUser200ResponseRepositoriesInner) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *GetGithubUser200ResponseRepositoriesInner) SetForks(v int32)`

SetForks sets Forks field to given value.

### HasForks

`func (o *GetGithubUser200ResponseRepositoriesInner) HasForks() bool`

HasForks returns a boolean if a field has been set.

### GetFullName

`func (o *GetGithubUser200ResponseRepositoriesInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetGithubUser200ResponseRepositoriesInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetGithubUser200ResponseRepositoriesInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHomepage

`func (o *GetGithubUser200ResponseRepositoriesInner) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *GetGithubUser200ResponseRepositoriesInner) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *GetGithubUser200ResponseRepositoriesInner) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GetGithubUser200ResponseRepositoriesInner) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GetGithubUser200ResponseRepositoriesInner) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GetGithubUser200ResponseRepositoriesInner) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetLanguage

`func (o *GetGithubUser200ResponseRepositoriesInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetGithubUser200ResponseRepositoriesInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetGithubUser200ResponseRepositoriesInner) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *GetGithubUser200ResponseRepositoriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGithubUser200ResponseRepositoriesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGithubUser200ResponseRepositoriesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPushedAt

`func (o *GetGithubUser200ResponseRepositoriesInner) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *GetGithubUser200ResponseRepositoriesInner) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.

### HasPushedAt

`func (o *GetGithubUser200ResponseRepositoriesInner) HasPushedAt() bool`

HasPushedAt returns a boolean if a field has been set.

### GetStargazers

`func (o *GetGithubUser200ResponseRepositoriesInner) GetStargazers() int32`

GetStargazers returns the Stargazers field if non-nil, zero value otherwise.

### GetStargazersOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetStargazersOk() (*int32, bool)`

GetStargazersOk returns a tuple with the Stargazers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazers

`func (o *GetGithubUser200ResponseRepositoriesInner) SetStargazers(v int32)`

SetStargazers sets Stargazers field to given value.

### HasStargazers

`func (o *GetGithubUser200ResponseRepositoriesInner) HasStargazers() bool`

HasStargazers returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetGithubUser200ResponseRepositoriesInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetGithubUser200ResponseRepositoriesInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetGithubUser200ResponseRepositoriesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibility

`func (o *GetGithubUser200ResponseRepositoriesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetGithubUser200ResponseRepositoriesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetGithubUser200ResponseRepositoriesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetGithubUser200ResponseRepositoriesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



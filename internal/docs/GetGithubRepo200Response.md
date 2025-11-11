# GetGithubRepo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Homepage** | Pointer to **string** |  | [optional] 
**DefaultBranch** | Pointer to **string** |  | [optional] 
**PrimaryBranch** | Pointer to **string** |  | [optional] 
**DefaultBranchSha** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Fork** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Stargazers** | Pointer to **int32** |  | [optional] 
**Forks** | Pointer to **int32** |  | [optional] 
**OpenIssues** | Pointer to **int32** |  | [optional] 
**Watchers** | Pointer to **int32** |  | [optional] 
**PushedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Languages** | Pointer to **map[string]int32** |  | [optional] 
**Collaborators** | Pointer to [**[]GetGithubRepo200ResponseCollaboratorsInner**](GetGithubRepo200ResponseCollaboratorsInner.md) |  | [optional] 
**Maintainers** | Pointer to [**[]GetGithubRepo200ResponseCollaboratorsInner**](GetGithubRepo200ResponseCollaboratorsInner.md) |  | [optional] 

## Methods

### NewGetGithubRepo200Response

`func NewGetGithubRepo200Response() *GetGithubRepo200Response`

NewGetGithubRepo200Response instantiates a new GetGithubRepo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGithubRepo200ResponseWithDefaults

`func NewGetGithubRepo200ResponseWithDefaults() *GetGithubRepo200Response`

NewGetGithubRepo200ResponseWithDefaults instantiates a new GetGithubRepo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *GetGithubRepo200Response) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetGithubRepo200Response) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetGithubRepo200Response) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetGithubRepo200Response) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *GetGithubRepo200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGithubRepo200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGithubRepo200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGithubRepo200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomepage

`func (o *GetGithubRepo200Response) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *GetGithubRepo200Response) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *GetGithubRepo200Response) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *GetGithubRepo200Response) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *GetGithubRepo200Response) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *GetGithubRepo200Response) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *GetGithubRepo200Response) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *GetGithubRepo200Response) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetPrimaryBranch

`func (o *GetGithubRepo200Response) GetPrimaryBranch() string`

GetPrimaryBranch returns the PrimaryBranch field if non-nil, zero value otherwise.

### GetPrimaryBranchOk

`func (o *GetGithubRepo200Response) GetPrimaryBranchOk() (*string, bool)`

GetPrimaryBranchOk returns a tuple with the PrimaryBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBranch

`func (o *GetGithubRepo200Response) SetPrimaryBranch(v string)`

SetPrimaryBranch sets PrimaryBranch field to given value.

### HasPrimaryBranch

`func (o *GetGithubRepo200Response) HasPrimaryBranch() bool`

HasPrimaryBranch returns a boolean if a field has been set.

### GetDefaultBranchSha

`func (o *GetGithubRepo200Response) GetDefaultBranchSha() string`

GetDefaultBranchSha returns the DefaultBranchSha field if non-nil, zero value otherwise.

### GetDefaultBranchShaOk

`func (o *GetGithubRepo200Response) GetDefaultBranchShaOk() (*string, bool)`

GetDefaultBranchShaOk returns a tuple with the DefaultBranchSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchSha

`func (o *GetGithubRepo200Response) SetDefaultBranchSha(v string)`

SetDefaultBranchSha sets DefaultBranchSha field to given value.

### HasDefaultBranchSha

`func (o *GetGithubRepo200Response) HasDefaultBranchSha() bool`

HasDefaultBranchSha returns a boolean if a field has been set.

### GetVisibility

`func (o *GetGithubRepo200Response) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetGithubRepo200Response) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetGithubRepo200Response) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetGithubRepo200Response) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetArchived

`func (o *GetGithubRepo200Response) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *GetGithubRepo200Response) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *GetGithubRepo200Response) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *GetGithubRepo200Response) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDisabled

`func (o *GetGithubRepo200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetGithubRepo200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetGithubRepo200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetGithubRepo200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFork

`func (o *GetGithubRepo200Response) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *GetGithubRepo200Response) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *GetGithubRepo200Response) SetFork(v bool)`

SetFork sets Fork field to given value.

### HasFork

`func (o *GetGithubRepo200Response) HasFork() bool`

HasFork returns a boolean if a field has been set.

### GetLanguage

`func (o *GetGithubRepo200Response) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetGithubRepo200Response) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetGithubRepo200Response) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetGithubRepo200Response) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTopics

`func (o *GetGithubRepo200Response) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *GetGithubRepo200Response) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *GetGithubRepo200Response) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *GetGithubRepo200Response) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetLicense

`func (o *GetGithubRepo200Response) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *GetGithubRepo200Response) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *GetGithubRepo200Response) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *GetGithubRepo200Response) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetStargazers

`func (o *GetGithubRepo200Response) GetStargazers() int32`

GetStargazers returns the Stargazers field if non-nil, zero value otherwise.

### GetStargazersOk

`func (o *GetGithubRepo200Response) GetStargazersOk() (*int32, bool)`

GetStargazersOk returns a tuple with the Stargazers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazers

`func (o *GetGithubRepo200Response) SetStargazers(v int32)`

SetStargazers sets Stargazers field to given value.

### HasStargazers

`func (o *GetGithubRepo200Response) HasStargazers() bool`

HasStargazers returns a boolean if a field has been set.

### GetForks

`func (o *GetGithubRepo200Response) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *GetGithubRepo200Response) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *GetGithubRepo200Response) SetForks(v int32)`

SetForks sets Forks field to given value.

### HasForks

`func (o *GetGithubRepo200Response) HasForks() bool`

HasForks returns a boolean if a field has been set.

### GetOpenIssues

`func (o *GetGithubRepo200Response) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *GetGithubRepo200Response) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *GetGithubRepo200Response) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.

### HasOpenIssues

`func (o *GetGithubRepo200Response) HasOpenIssues() bool`

HasOpenIssues returns a boolean if a field has been set.

### GetWatchers

`func (o *GetGithubRepo200Response) GetWatchers() int32`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *GetGithubRepo200Response) GetWatchersOk() (*int32, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *GetGithubRepo200Response) SetWatchers(v int32)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *GetGithubRepo200Response) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetPushedAt

`func (o *GetGithubRepo200Response) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *GetGithubRepo200Response) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *GetGithubRepo200Response) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.

### HasPushedAt

`func (o *GetGithubRepo200Response) HasPushedAt() bool`

HasPushedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetGithubRepo200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetGithubRepo200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetGithubRepo200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetGithubRepo200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetGithubRepo200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetGithubRepo200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetGithubRepo200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetGithubRepo200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLanguages

`func (o *GetGithubRepo200Response) GetLanguages() map[string]int32`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *GetGithubRepo200Response) GetLanguagesOk() (*map[string]int32, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *GetGithubRepo200Response) SetLanguages(v map[string]int32)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *GetGithubRepo200Response) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetCollaborators

`func (o *GetGithubRepo200Response) GetCollaborators() []GetGithubRepo200ResponseCollaboratorsInner`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *GetGithubRepo200Response) GetCollaboratorsOk() (*[]GetGithubRepo200ResponseCollaboratorsInner, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *GetGithubRepo200Response) SetCollaborators(v []GetGithubRepo200ResponseCollaboratorsInner)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *GetGithubRepo200Response) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.

### SetCollaboratorsNil

`func (o *GetGithubRepo200Response) SetCollaboratorsNil(b bool)`

 SetCollaboratorsNil sets the value for Collaborators to be an explicit nil

### UnsetCollaborators
`func (o *GetGithubRepo200Response) UnsetCollaborators()`

UnsetCollaborators ensures that no value is present for Collaborators, not even an explicit nil
### GetMaintainers

`func (o *GetGithubRepo200Response) GetMaintainers() []GetGithubRepo200ResponseCollaboratorsInner`

GetMaintainers returns the Maintainers field if non-nil, zero value otherwise.

### GetMaintainersOk

`func (o *GetGithubRepo200Response) GetMaintainersOk() (*[]GetGithubRepo200ResponseCollaboratorsInner, bool)`

GetMaintainersOk returns a tuple with the Maintainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainers

`func (o *GetGithubRepo200Response) SetMaintainers(v []GetGithubRepo200ResponseCollaboratorsInner)`

SetMaintainers sets Maintainers field to given value.

### HasMaintainers

`func (o *GetGithubRepo200Response) HasMaintainers() bool`

HasMaintainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



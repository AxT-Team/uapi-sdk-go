# GetGithubUser200ResponseActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContributionCalendar** | Pointer to [**GetGithubUser200ResponseActivityContributionCalendar**](GetGithubUser200ResponseActivityContributionCalendar.md) |  | [optional] 
**From** | Pointer to **string** | 统计开始日期。 | [optional] 
**Organization** | Pointer to **string** | 按组织过滤时对应的组织登录名。 | [optional] 
**Scope** | Pointer to **string** | 活动统计范围，常见值为 all 或 organization。 | [optional] 
**Timeline** | Pointer to [**[]GetGithubUser200ResponseActivityTimelineInner**](GetGithubUser200ResponseActivityTimelineInner.md) | 按月份聚合后的贡献时间线。 | [optional] 
**To** | Pointer to **string** | 统计结束日期。 | [optional] 
**TotalCommitContributions** | Pointer to **int32** | Commit 贡献总数。 | [optional] 
**TotalContributions** | Pointer to **int32** | 统计范围内的总贡献数。 | [optional] 
**TotalIssueContributions** | Pointer to **int32** | Issue 贡献总数。 | [optional] 
**TotalPullRequestContributions** | Pointer to **int32** | Pull Request 贡献总数。 | [optional] 
**TotalPullRequestReviewContributions** | Pointer to **int32** | Pull Request Review 贡献总数。 | [optional] 

## Methods

### NewGetGithubUser200ResponseActivity

`func NewGetGithubUser200ResponseActivity() *GetGithubUser200ResponseActivity`

NewGetGithubUser200ResponseActivity instantiates a new GetGithubUser200ResponseActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGithubUser200ResponseActivityWithDefaults

`func NewGetGithubUser200ResponseActivityWithDefaults() *GetGithubUser200ResponseActivity`

NewGetGithubUser200ResponseActivityWithDefaults instantiates a new GetGithubUser200ResponseActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributionCalendar

`func (o *GetGithubUser200ResponseActivity) GetContributionCalendar() GetGithubUser200ResponseActivityContributionCalendar`

GetContributionCalendar returns the ContributionCalendar field if non-nil, zero value otherwise.

### GetContributionCalendarOk

`func (o *GetGithubUser200ResponseActivity) GetContributionCalendarOk() (*GetGithubUser200ResponseActivityContributionCalendar, bool)`

GetContributionCalendarOk returns a tuple with the ContributionCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributionCalendar

`func (o *GetGithubUser200ResponseActivity) SetContributionCalendar(v GetGithubUser200ResponseActivityContributionCalendar)`

SetContributionCalendar sets ContributionCalendar field to given value.

### HasContributionCalendar

`func (o *GetGithubUser200ResponseActivity) HasContributionCalendar() bool`

HasContributionCalendar returns a boolean if a field has been set.

### GetFrom

`func (o *GetGithubUser200ResponseActivity) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetGithubUser200ResponseActivity) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetGithubUser200ResponseActivity) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetGithubUser200ResponseActivity) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetOrganization

`func (o *GetGithubUser200ResponseActivity) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetGithubUser200ResponseActivity) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetGithubUser200ResponseActivity) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetGithubUser200ResponseActivity) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScope

`func (o *GetGithubUser200ResponseActivity) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetGithubUser200ResponseActivity) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetGithubUser200ResponseActivity) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetGithubUser200ResponseActivity) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTimeline

`func (o *GetGithubUser200ResponseActivity) GetTimeline() []GetGithubUser200ResponseActivityTimelineInner`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *GetGithubUser200ResponseActivity) GetTimelineOk() (*[]GetGithubUser200ResponseActivityTimelineInner, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *GetGithubUser200ResponseActivity) SetTimeline(v []GetGithubUser200ResponseActivityTimelineInner)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *GetGithubUser200ResponseActivity) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetTo

`func (o *GetGithubUser200ResponseActivity) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetGithubUser200ResponseActivity) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetGithubUser200ResponseActivity) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetGithubUser200ResponseActivity) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotalCommitContributions

`func (o *GetGithubUser200ResponseActivity) GetTotalCommitContributions() int32`

GetTotalCommitContributions returns the TotalCommitContributions field if non-nil, zero value otherwise.

### GetTotalCommitContributionsOk

`func (o *GetGithubUser200ResponseActivity) GetTotalCommitContributionsOk() (*int32, bool)`

GetTotalCommitContributionsOk returns a tuple with the TotalCommitContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommitContributions

`func (o *GetGithubUser200ResponseActivity) SetTotalCommitContributions(v int32)`

SetTotalCommitContributions sets TotalCommitContributions field to given value.

### HasTotalCommitContributions

`func (o *GetGithubUser200ResponseActivity) HasTotalCommitContributions() bool`

HasTotalCommitContributions returns a boolean if a field has been set.

### GetTotalContributions

`func (o *GetGithubUser200ResponseActivity) GetTotalContributions() int32`

GetTotalContributions returns the TotalContributions field if non-nil, zero value otherwise.

### GetTotalContributionsOk

`func (o *GetGithubUser200ResponseActivity) GetTotalContributionsOk() (*int32, bool)`

GetTotalContributionsOk returns a tuple with the TotalContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContributions

`func (o *GetGithubUser200ResponseActivity) SetTotalContributions(v int32)`

SetTotalContributions sets TotalContributions field to given value.

### HasTotalContributions

`func (o *GetGithubUser200ResponseActivity) HasTotalContributions() bool`

HasTotalContributions returns a boolean if a field has been set.

### GetTotalIssueContributions

`func (o *GetGithubUser200ResponseActivity) GetTotalIssueContributions() int32`

GetTotalIssueContributions returns the TotalIssueContributions field if non-nil, zero value otherwise.

### GetTotalIssueContributionsOk

`func (o *GetGithubUser200ResponseActivity) GetTotalIssueContributionsOk() (*int32, bool)`

GetTotalIssueContributionsOk returns a tuple with the TotalIssueContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIssueContributions

`func (o *GetGithubUser200ResponseActivity) SetTotalIssueContributions(v int32)`

SetTotalIssueContributions sets TotalIssueContributions field to given value.

### HasTotalIssueContributions

`func (o *GetGithubUser200ResponseActivity) HasTotalIssueContributions() bool`

HasTotalIssueContributions returns a boolean if a field has been set.

### GetTotalPullRequestContributions

`func (o *GetGithubUser200ResponseActivity) GetTotalPullRequestContributions() int32`

GetTotalPullRequestContributions returns the TotalPullRequestContributions field if non-nil, zero value otherwise.

### GetTotalPullRequestContributionsOk

`func (o *GetGithubUser200ResponseActivity) GetTotalPullRequestContributionsOk() (*int32, bool)`

GetTotalPullRequestContributionsOk returns a tuple with the TotalPullRequestContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPullRequestContributions

`func (o *GetGithubUser200ResponseActivity) SetTotalPullRequestContributions(v int32)`

SetTotalPullRequestContributions sets TotalPullRequestContributions field to given value.

### HasTotalPullRequestContributions

`func (o *GetGithubUser200ResponseActivity) HasTotalPullRequestContributions() bool`

HasTotalPullRequestContributions returns a boolean if a field has been set.

### GetTotalPullRequestReviewContributions

`func (o *GetGithubUser200ResponseActivity) GetTotalPullRequestReviewContributions() int32`

GetTotalPullRequestReviewContributions returns the TotalPullRequestReviewContributions field if non-nil, zero value otherwise.

### GetTotalPullRequestReviewContributionsOk

`func (o *GetGithubUser200ResponseActivity) GetTotalPullRequestReviewContributionsOk() (*int32, bool)`

GetTotalPullRequestReviewContributionsOk returns a tuple with the TotalPullRequestReviewContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPullRequestReviewContributions

`func (o *GetGithubUser200ResponseActivity) SetTotalPullRequestReviewContributions(v int32)`

SetTotalPullRequestReviewContributions sets TotalPullRequestReviewContributions field to given value.

### HasTotalPullRequestReviewContributions

`func (o *GetGithubUser200ResponseActivity) HasTotalPullRequestReviewContributions() bool`

HasTotalPullRequestReviewContributions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



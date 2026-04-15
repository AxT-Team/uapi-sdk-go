# GetGithubUser200ResponseActivityContributionCalendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Colors** | Pointer to **[]string** | 贡献等级对应的颜色列表。 | [optional] 
**IsHalloween** | Pointer to **bool** | 是否启用了 GitHub 万圣节配色主题。 | [optional] 
**TotalContributions** | Pointer to **int32** | 贡献日历中的总贡献数。 | [optional] 
**Weeks** | Pointer to [**[]GetGithubUser200ResponseActivityContributionCalendarWeeksInner**](GetGithubUser200ResponseActivityContributionCalendarWeeksInner.md) | 按周排列的贡献日历列数据。 | [optional] 

## Methods

### NewGetGithubUser200ResponseActivityContributionCalendar

`func NewGetGithubUser200ResponseActivityContributionCalendar() *GetGithubUser200ResponseActivityContributionCalendar`

NewGetGithubUser200ResponseActivityContributionCalendar instantiates a new GetGithubUser200ResponseActivityContributionCalendar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGithubUser200ResponseActivityContributionCalendarWithDefaults

`func NewGetGithubUser200ResponseActivityContributionCalendarWithDefaults() *GetGithubUser200ResponseActivityContributionCalendar`

NewGetGithubUser200ResponseActivityContributionCalendarWithDefaults instantiates a new GetGithubUser200ResponseActivityContributionCalendar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColors

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetColors() []string`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetColorsOk() (*[]string, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *GetGithubUser200ResponseActivityContributionCalendar) SetColors(v []string)`

SetColors sets Colors field to given value.

### HasColors

`func (o *GetGithubUser200ResponseActivityContributionCalendar) HasColors() bool`

HasColors returns a boolean if a field has been set.

### GetIsHalloween

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetIsHalloween() bool`

GetIsHalloween returns the IsHalloween field if non-nil, zero value otherwise.

### GetIsHalloweenOk

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetIsHalloweenOk() (*bool, bool)`

GetIsHalloweenOk returns a tuple with the IsHalloween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHalloween

`func (o *GetGithubUser200ResponseActivityContributionCalendar) SetIsHalloween(v bool)`

SetIsHalloween sets IsHalloween field to given value.

### HasIsHalloween

`func (o *GetGithubUser200ResponseActivityContributionCalendar) HasIsHalloween() bool`

HasIsHalloween returns a boolean if a field has been set.

### GetTotalContributions

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetTotalContributions() int32`

GetTotalContributions returns the TotalContributions field if non-nil, zero value otherwise.

### GetTotalContributionsOk

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetTotalContributionsOk() (*int32, bool)`

GetTotalContributionsOk returns a tuple with the TotalContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContributions

`func (o *GetGithubUser200ResponseActivityContributionCalendar) SetTotalContributions(v int32)`

SetTotalContributions sets TotalContributions field to given value.

### HasTotalContributions

`func (o *GetGithubUser200ResponseActivityContributionCalendar) HasTotalContributions() bool`

HasTotalContributions returns a boolean if a field has been set.

### GetWeeks

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetWeeks() []GetGithubUser200ResponseActivityContributionCalendarWeeksInner`

GetWeeks returns the Weeks field if non-nil, zero value otherwise.

### GetWeeksOk

`func (o *GetGithubUser200ResponseActivityContributionCalendar) GetWeeksOk() (*[]GetGithubUser200ResponseActivityContributionCalendarWeeksInner, bool)`

GetWeeksOk returns a tuple with the Weeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeks

`func (o *GetGithubUser200ResponseActivityContributionCalendar) SetWeeks(v []GetGithubUser200ResponseActivityContributionCalendarWeeksInner)`

SetWeeks sets Weeks field to given value.

### HasWeeks

`func (o *GetGithubUser200ResponseActivityContributionCalendar) HasWeeks() bool`

HasWeeks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



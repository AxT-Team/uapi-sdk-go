# FormatJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Market** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Copyright** | Pointer to **string** |  | [optional] 
**CopyrightLink** | Pointer to **string** |  | [optional] 
**QuizId** | Pointer to **string** |  | [optional] 
**Trivia** | Pointer to [**FormatJsonTrivia**](FormatJsonTrivia.md) |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**ImageUrl4k** | Pointer to **string** |  | [optional] 
**ImageUrl1080** | Pointer to **string** |  | [optional] 
**FetchedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewFormatJson

`func NewFormatJson() *FormatJson`

NewFormatJson instantiates a new FormatJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormatJsonWithDefaults

`func NewFormatJsonWithDefaults() *FormatJson`

NewFormatJsonWithDefaults instantiates a new FormatJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *FormatJson) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FormatJson) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FormatJson) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *FormatJson) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMarket

`func (o *FormatJson) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *FormatJson) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *FormatJson) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *FormatJson) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetTitle

`func (o *FormatJson) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormatJson) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormatJson) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormatJson) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSubtitle

`func (o *FormatJson) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *FormatJson) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *FormatJson) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *FormatJson) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetHeadline

`func (o *FormatJson) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *FormatJson) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *FormatJson) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *FormatJson) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetDescription

`func (o *FormatJson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormatJson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormatJson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FormatJson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCopyright

`func (o *FormatJson) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *FormatJson) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *FormatJson) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *FormatJson) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### GetCopyrightLink

`func (o *FormatJson) GetCopyrightLink() string`

GetCopyrightLink returns the CopyrightLink field if non-nil, zero value otherwise.

### GetCopyrightLinkOk

`func (o *FormatJson) GetCopyrightLinkOk() (*string, bool)`

GetCopyrightLinkOk returns a tuple with the CopyrightLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightLink

`func (o *FormatJson) SetCopyrightLink(v string)`

SetCopyrightLink sets CopyrightLink field to given value.

### HasCopyrightLink

`func (o *FormatJson) HasCopyrightLink() bool`

HasCopyrightLink returns a boolean if a field has been set.

### GetQuizId

`func (o *FormatJson) GetQuizId() string`

GetQuizId returns the QuizId field if non-nil, zero value otherwise.

### GetQuizIdOk

`func (o *FormatJson) GetQuizIdOk() (*string, bool)`

GetQuizIdOk returns a tuple with the QuizId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuizId

`func (o *FormatJson) SetQuizId(v string)`

SetQuizId sets QuizId field to given value.

### HasQuizId

`func (o *FormatJson) HasQuizId() bool`

HasQuizId returns a boolean if a field has been set.

### GetTrivia

`func (o *FormatJson) GetTrivia() FormatJsonTrivia`

GetTrivia returns the Trivia field if non-nil, zero value otherwise.

### GetTriviaOk

`func (o *FormatJson) GetTriviaOk() (*FormatJsonTrivia, bool)`

GetTriviaOk returns a tuple with the Trivia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrivia

`func (o *FormatJson) SetTrivia(v FormatJsonTrivia)`

SetTrivia sets Trivia field to given value.

### HasTrivia

`func (o *FormatJson) HasTrivia() bool`

HasTrivia returns a boolean if a field has been set.

### GetResolution

`func (o *FormatJson) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *FormatJson) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *FormatJson) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *FormatJson) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetImageUrl

`func (o *FormatJson) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *FormatJson) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *FormatJson) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *FormatJson) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetImageUrl4k

`func (o *FormatJson) GetImageUrl4k() string`

GetImageUrl4k returns the ImageUrl4k field if non-nil, zero value otherwise.

### GetImageUrl4kOk

`func (o *FormatJson) GetImageUrl4kOk() (*string, bool)`

GetImageUrl4kOk returns a tuple with the ImageUrl4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl4k

`func (o *FormatJson) SetImageUrl4k(v string)`

SetImageUrl4k sets ImageUrl4k field to given value.

### HasImageUrl4k

`func (o *FormatJson) HasImageUrl4k() bool`

HasImageUrl4k returns a boolean if a field has been set.

### GetImageUrl1080

`func (o *FormatJson) GetImageUrl1080() string`

GetImageUrl1080 returns the ImageUrl1080 field if non-nil, zero value otherwise.

### GetImageUrl1080Ok

`func (o *FormatJson) GetImageUrl1080Ok() (*string, bool)`

GetImageUrl1080Ok returns a tuple with the ImageUrl1080 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl1080

`func (o *FormatJson) SetImageUrl1080(v string)`

SetImageUrl1080 sets ImageUrl1080 field to given value.

### HasImageUrl1080

`func (o *FormatJson) HasImageUrl1080() bool`

HasImageUrl1080 returns a boolean if a field has been set.

### GetFetchedAt

`func (o *FormatJson) GetFetchedAt() string`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *FormatJson) GetFetchedAtOk() (*string, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *FormatJson) SetFetchedAt(v string)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *FormatJson) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FormatJson) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormatJson) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormatJson) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormatJson) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetMiscMovieBoxOffice200ResponseListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgSeatView** | Pointer to **string** | 上座率 | [optional] 
**AvgShowView** | Pointer to **string** | 场均人次 | [optional] 
**BoxOffice** | Pointer to **string** | 实时综合票房，带单位 | [optional] 
**BoxOfficeRate** | Pointer to **string** | 实时综合票房占比 | [optional] 
**DetailUrl** | Pointer to **string** | 电影详情页 URL | [optional] 
**MovieId** | Pointer to **int32** | 影片 ID | [optional] 
**MovieName** | Pointer to **string** | 影片名称 | [optional] 
**Rank** | Pointer to **int32** | 排名，从 1 开始 | [optional] 
**ReleaseDays** | Pointer to **int32** | 已上映天数。仅当 release_info 可解析为“上映N天”时返回 | [optional] 
**ReleaseInfo** | Pointer to **string** | 上游上映信息原文 | [optional] 
**ReleaseStatus** | Pointer to **string** | 结构化上映状态，可取 released、preview、re_release 或 other | [optional] 
**ShowCount** | Pointer to **int32** | 排片场次 | [optional] 
**ShowCountRate** | Pointer to **string** | 排片占比 | [optional] 
**SplitBoxOffice** | Pointer to **string** | 实时分账票房，带单位 | [optional] 
**SplitBoxOfficeRate** | Pointer to **string** | 实时分账票房占比 | [optional] 
**SumBoxOffice** | Pointer to **string** | 累计综合票房 | [optional] 
**SumSplitBoxOffice** | Pointer to **string** | 累计分账票房 | [optional] 

## Methods

### NewGetMiscMovieBoxOffice200ResponseListInner

`func NewGetMiscMovieBoxOffice200ResponseListInner() *GetMiscMovieBoxOffice200ResponseListInner`

NewGetMiscMovieBoxOffice200ResponseListInner instantiates a new GetMiscMovieBoxOffice200ResponseListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscMovieBoxOffice200ResponseListInnerWithDefaults

`func NewGetMiscMovieBoxOffice200ResponseListInnerWithDefaults() *GetMiscMovieBoxOffice200ResponseListInner`

NewGetMiscMovieBoxOffice200ResponseListInnerWithDefaults instantiates a new GetMiscMovieBoxOffice200ResponseListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgSeatView

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetAvgSeatView() string`

GetAvgSeatView returns the AvgSeatView field if non-nil, zero value otherwise.

### GetAvgSeatViewOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetAvgSeatViewOk() (*string, bool)`

GetAvgSeatViewOk returns a tuple with the AvgSeatView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSeatView

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetAvgSeatView(v string)`

SetAvgSeatView sets AvgSeatView field to given value.

### HasAvgSeatView

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasAvgSeatView() bool`

HasAvgSeatView returns a boolean if a field has been set.

### GetAvgShowView

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetAvgShowView() string`

GetAvgShowView returns the AvgShowView field if non-nil, zero value otherwise.

### GetAvgShowViewOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetAvgShowViewOk() (*string, bool)`

GetAvgShowViewOk returns a tuple with the AvgShowView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgShowView

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetAvgShowView(v string)`

SetAvgShowView sets AvgShowView field to given value.

### HasAvgShowView

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasAvgShowView() bool`

HasAvgShowView returns a boolean if a field has been set.

### GetBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetBoxOffice() string`

GetBoxOffice returns the BoxOffice field if non-nil, zero value otherwise.

### GetBoxOfficeOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetBoxOfficeOk() (*string, bool)`

GetBoxOfficeOk returns a tuple with the BoxOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetBoxOffice(v string)`

SetBoxOffice sets BoxOffice field to given value.

### HasBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasBoxOffice() bool`

HasBoxOffice returns a boolean if a field has been set.

### GetBoxOfficeRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetBoxOfficeRate() string`

GetBoxOfficeRate returns the BoxOfficeRate field if non-nil, zero value otherwise.

### GetBoxOfficeRateOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetBoxOfficeRateOk() (*string, bool)`

GetBoxOfficeRateOk returns a tuple with the BoxOfficeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxOfficeRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetBoxOfficeRate(v string)`

SetBoxOfficeRate sets BoxOfficeRate field to given value.

### HasBoxOfficeRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasBoxOfficeRate() bool`

HasBoxOfficeRate returns a boolean if a field has been set.

### GetDetailUrl

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetDetailUrl() string`

GetDetailUrl returns the DetailUrl field if non-nil, zero value otherwise.

### GetDetailUrlOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetDetailUrlOk() (*string, bool)`

GetDetailUrlOk returns a tuple with the DetailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailUrl

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetDetailUrl(v string)`

SetDetailUrl sets DetailUrl field to given value.

### HasDetailUrl

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasDetailUrl() bool`

HasDetailUrl returns a boolean if a field has been set.

### GetMovieId

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetMovieId() int32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetMovieIdOk() (*int32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetMovieId(v int32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.

### GetMovieName

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetMovieName() string`

GetMovieName returns the MovieName field if non-nil, zero value otherwise.

### GetMovieNameOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetMovieNameOk() (*string, bool)`

GetMovieNameOk returns a tuple with the MovieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieName

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetMovieName(v string)`

SetMovieName sets MovieName field to given value.

### HasMovieName

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasMovieName() bool`

HasMovieName returns a boolean if a field has been set.

### GetRank

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetReleaseDays

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetReleaseDays() int32`

GetReleaseDays returns the ReleaseDays field if non-nil, zero value otherwise.

### GetReleaseDaysOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetReleaseDaysOk() (*int32, bool)`

GetReleaseDaysOk returns a tuple with the ReleaseDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDays

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetReleaseDays(v int32)`

SetReleaseDays sets ReleaseDays field to given value.

### HasReleaseDays

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasReleaseDays() bool`

HasReleaseDays returns a boolean if a field has been set.

### GetReleaseInfo

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetReleaseInfo() string`

GetReleaseInfo returns the ReleaseInfo field if non-nil, zero value otherwise.

### GetReleaseInfoOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetReleaseInfoOk() (*string, bool)`

GetReleaseInfoOk returns a tuple with the ReleaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseInfo

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetReleaseInfo(v string)`

SetReleaseInfo sets ReleaseInfo field to given value.

### HasReleaseInfo

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasReleaseInfo() bool`

HasReleaseInfo returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetReleaseStatus() string`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetReleaseStatusOk() (*string, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetReleaseStatus(v string)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetShowCount

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetShowCount() int32`

GetShowCount returns the ShowCount field if non-nil, zero value otherwise.

### GetShowCountOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetShowCountOk() (*int32, bool)`

GetShowCountOk returns a tuple with the ShowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCount

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetShowCount(v int32)`

SetShowCount sets ShowCount field to given value.

### HasShowCount

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasShowCount() bool`

HasShowCount returns a boolean if a field has been set.

### GetShowCountRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetShowCountRate() string`

GetShowCountRate returns the ShowCountRate field if non-nil, zero value otherwise.

### GetShowCountRateOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetShowCountRateOk() (*string, bool)`

GetShowCountRateOk returns a tuple with the ShowCountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCountRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetShowCountRate(v string)`

SetShowCountRate sets ShowCountRate field to given value.

### HasShowCountRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasShowCountRate() bool`

HasShowCountRate returns a boolean if a field has been set.

### GetSplitBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSplitBoxOffice() string`

GetSplitBoxOffice returns the SplitBoxOffice field if non-nil, zero value otherwise.

### GetSplitBoxOfficeOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSplitBoxOfficeOk() (*string, bool)`

GetSplitBoxOfficeOk returns a tuple with the SplitBoxOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetSplitBoxOffice(v string)`

SetSplitBoxOffice sets SplitBoxOffice field to given value.

### HasSplitBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasSplitBoxOffice() bool`

HasSplitBoxOffice returns a boolean if a field has been set.

### GetSplitBoxOfficeRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSplitBoxOfficeRate() string`

GetSplitBoxOfficeRate returns the SplitBoxOfficeRate field if non-nil, zero value otherwise.

### GetSplitBoxOfficeRateOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSplitBoxOfficeRateOk() (*string, bool)`

GetSplitBoxOfficeRateOk returns a tuple with the SplitBoxOfficeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitBoxOfficeRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetSplitBoxOfficeRate(v string)`

SetSplitBoxOfficeRate sets SplitBoxOfficeRate field to given value.

### HasSplitBoxOfficeRate

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasSplitBoxOfficeRate() bool`

HasSplitBoxOfficeRate returns a boolean if a field has been set.

### GetSumBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSumBoxOffice() string`

GetSumBoxOffice returns the SumBoxOffice field if non-nil, zero value otherwise.

### GetSumBoxOfficeOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSumBoxOfficeOk() (*string, bool)`

GetSumBoxOfficeOk returns a tuple with the SumBoxOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetSumBoxOffice(v string)`

SetSumBoxOffice sets SumBoxOffice field to given value.

### HasSumBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasSumBoxOffice() bool`

HasSumBoxOffice returns a boolean if a field has been set.

### GetSumSplitBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSumSplitBoxOffice() string`

GetSumSplitBoxOffice returns the SumSplitBoxOffice field if non-nil, zero value otherwise.

### GetSumSplitBoxOfficeOk

`func (o *GetMiscMovieBoxOffice200ResponseListInner) GetSumSplitBoxOfficeOk() (*string, bool)`

GetSumSplitBoxOfficeOk returns a tuple with the SumSplitBoxOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumSplitBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) SetSumSplitBoxOffice(v string)`

SetSumSplitBoxOffice sets SumSplitBoxOffice field to given value.

### HasSumSplitBoxOffice

`func (o *GetMiscMovieBoxOffice200ResponseListInner) HasSumSplitBoxOffice() bool`

HasSumSplitBoxOffice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



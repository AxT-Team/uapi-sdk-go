# GetWebTomarkdownAsyncStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Elapsed** | Pointer to **float32** |  | [optional] 
**CompletedAt** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**Result** | Pointer to [**GetWebTomarkdownAsyncStatus200ResponseAnyOf2Result**](GetWebTomarkdownAsyncStatus200ResponseAnyOf2Result.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewGetWebTomarkdownAsyncStatus200Response

`func NewGetWebTomarkdownAsyncStatus200Response() *GetWebTomarkdownAsyncStatus200Response`

NewGetWebTomarkdownAsyncStatus200Response instantiates a new GetWebTomarkdownAsyncStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebTomarkdownAsyncStatus200ResponseWithDefaults

`func NewGetWebTomarkdownAsyncStatus200ResponseWithDefaults() *GetWebTomarkdownAsyncStatus200Response`

NewGetWebTomarkdownAsyncStatus200ResponseWithDefaults instantiates a new GetWebTomarkdownAsyncStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *GetWebTomarkdownAsyncStatus200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetWebTomarkdownAsyncStatus200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetWebTomarkdownAsyncStatus200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetStatus

`func (o *GetWebTomarkdownAsyncStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetWebTomarkdownAsyncStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetWebTomarkdownAsyncStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GetWebTomarkdownAsyncStatus200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetWebTomarkdownAsyncStatus200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetWebTomarkdownAsyncStatus200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetProgress

`func (o *GetWebTomarkdownAsyncStatus200Response) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *GetWebTomarkdownAsyncStatus200Response) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *GetWebTomarkdownAsyncStatus200Response) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMessage

`func (o *GetWebTomarkdownAsyncStatus200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetWebTomarkdownAsyncStatus200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetWebTomarkdownAsyncStatus200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetElapsed

`func (o *GetWebTomarkdownAsyncStatus200Response) GetElapsed() float32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetElapsedOk() (*float32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *GetWebTomarkdownAsyncStatus200Response) SetElapsed(v float32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *GetWebTomarkdownAsyncStatus200Response) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetWebTomarkdownAsyncStatus200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetDuration

`func (o *GetWebTomarkdownAsyncStatus200Response) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetWebTomarkdownAsyncStatus200Response) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetWebTomarkdownAsyncStatus200Response) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetResult

`func (o *GetWebTomarkdownAsyncStatus200Response) GetResult() GetWebTomarkdownAsyncStatus200ResponseAnyOf2Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetResultOk() (*GetWebTomarkdownAsyncStatus200ResponseAnyOf2Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetWebTomarkdownAsyncStatus200Response) SetResult(v GetWebTomarkdownAsyncStatus200ResponseAnyOf2Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetWebTomarkdownAsyncStatus200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *GetWebTomarkdownAsyncStatus200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetWebTomarkdownAsyncStatus200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetWebTomarkdownAsyncStatus200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetWebTomarkdownAsyncStatus200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



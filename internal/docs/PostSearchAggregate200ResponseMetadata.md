# PostSearchAggregate200ResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestParams** | Pointer to [**PostSearchAggregate200ResponseMetadataRequestParams**](PostSearchAggregate200ResponseMetadataRequestParams.md) |  | [optional] 
**DedupeRemoved** | Pointer to **int32** | 去重后移除的结果数 | [optional] 
**RerankApplied** | Pointer to **bool** | 是否执行了排序重排 | [optional] 
**ContentFetched** | Pointer to **int32** | 额外抓取正文的结果数 | [optional] 

## Methods

### NewPostSearchAggregate200ResponseMetadata

`func NewPostSearchAggregate200ResponseMetadata() *PostSearchAggregate200ResponseMetadata`

NewPostSearchAggregate200ResponseMetadata instantiates a new PostSearchAggregate200ResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSearchAggregate200ResponseMetadataWithDefaults

`func NewPostSearchAggregate200ResponseMetadataWithDefaults() *PostSearchAggregate200ResponseMetadata`

NewPostSearchAggregate200ResponseMetadataWithDefaults instantiates a new PostSearchAggregate200ResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestParams

`func (o *PostSearchAggregate200ResponseMetadata) GetRequestParams() PostSearchAggregate200ResponseMetadataRequestParams`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *PostSearchAggregate200ResponseMetadata) GetRequestParamsOk() (*PostSearchAggregate200ResponseMetadataRequestParams, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *PostSearchAggregate200ResponseMetadata) SetRequestParams(v PostSearchAggregate200ResponseMetadataRequestParams)`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *PostSearchAggregate200ResponseMetadata) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### GetDedupeRemoved

`func (o *PostSearchAggregate200ResponseMetadata) GetDedupeRemoved() int32`

GetDedupeRemoved returns the DedupeRemoved field if non-nil, zero value otherwise.

### GetDedupeRemovedOk

`func (o *PostSearchAggregate200ResponseMetadata) GetDedupeRemovedOk() (*int32, bool)`

GetDedupeRemovedOk returns a tuple with the DedupeRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupeRemoved

`func (o *PostSearchAggregate200ResponseMetadata) SetDedupeRemoved(v int32)`

SetDedupeRemoved sets DedupeRemoved field to given value.

### HasDedupeRemoved

`func (o *PostSearchAggregate200ResponseMetadata) HasDedupeRemoved() bool`

HasDedupeRemoved returns a boolean if a field has been set.

### GetRerankApplied

`func (o *PostSearchAggregate200ResponseMetadata) GetRerankApplied() bool`

GetRerankApplied returns the RerankApplied field if non-nil, zero value otherwise.

### GetRerankAppliedOk

`func (o *PostSearchAggregate200ResponseMetadata) GetRerankAppliedOk() (*bool, bool)`

GetRerankAppliedOk returns a tuple with the RerankApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerankApplied

`func (o *PostSearchAggregate200ResponseMetadata) SetRerankApplied(v bool)`

SetRerankApplied sets RerankApplied field to given value.

### HasRerankApplied

`func (o *PostSearchAggregate200ResponseMetadata) HasRerankApplied() bool`

HasRerankApplied returns a boolean if a field has been set.

### GetContentFetched

`func (o *PostSearchAggregate200ResponseMetadata) GetContentFetched() int32`

GetContentFetched returns the ContentFetched field if non-nil, zero value otherwise.

### GetContentFetchedOk

`func (o *PostSearchAggregate200ResponseMetadata) GetContentFetchedOk() (*int32, bool)`

GetContentFetchedOk returns a tuple with the ContentFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFetched

`func (o *PostSearchAggregate200ResponseMetadata) SetContentFetched(v int32)`

SetContentFetched sets ContentFetched field to given value.

### HasContentFetched

`func (o *PostSearchAggregate200ResponseMetadata) HasContentFetched() bool`

HasContentFetched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



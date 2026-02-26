# GetConvertUnixtime200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | 标准格式（YYYY-MM-DD HH:mm:ss）的日期时间字符串。 | [optional] 
**Timestamp** | Pointer to **int32** | 转换后的10位秒级Unix时间戳。 | [optional] 

## Methods

### NewGetConvertUnixtime200Response

`func NewGetConvertUnixtime200Response() *GetConvertUnixtime200Response`

NewGetConvertUnixtime200Response instantiates a new GetConvertUnixtime200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConvertUnixtime200ResponseWithDefaults

`func NewGetConvertUnixtime200ResponseWithDefaults() *GetConvertUnixtime200Response`

NewGetConvertUnixtime200ResponseWithDefaults instantiates a new GetConvertUnixtime200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetConvertUnixtime200Response) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetConvertUnixtime200Response) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetConvertUnixtime200Response) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetConvertUnixtime200Response) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetConvertUnixtime200Response) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetConvertUnixtime200Response) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetConvertUnixtime200Response) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetConvertUnixtime200Response) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetSearchEngines200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to [**GetSearchEngines200ResponseEngine**](GetSearchEngines200ResponseEngine.md) |  | [optional] 
**Limits** | Pointer to [**GetSearchEngines200ResponseLimits**](GetSearchEngines200ResponseLimits.md) |  | [optional] 
**SupportedParameters** | Pointer to **[]string** | 支持的所有参数说明列表 | [optional] 

## Methods

### NewGetSearchEngines200Response

`func NewGetSearchEngines200Response() *GetSearchEngines200Response`

NewGetSearchEngines200Response instantiates a new GetSearchEngines200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSearchEngines200ResponseWithDefaults

`func NewGetSearchEngines200ResponseWithDefaults() *GetSearchEngines200Response`

NewGetSearchEngines200ResponseWithDefaults instantiates a new GetSearchEngines200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *GetSearchEngines200Response) GetEngine() GetSearchEngines200ResponseEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *GetSearchEngines200Response) GetEngineOk() (*GetSearchEngines200ResponseEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *GetSearchEngines200Response) SetEngine(v GetSearchEngines200ResponseEngine)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *GetSearchEngines200Response) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetLimits

`func (o *GetSearchEngines200Response) GetLimits() GetSearchEngines200ResponseLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *GetSearchEngines200Response) GetLimitsOk() (*GetSearchEngines200ResponseLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *GetSearchEngines200Response) SetLimits(v GetSearchEngines200ResponseLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *GetSearchEngines200Response) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetSupportedParameters

`func (o *GetSearchEngines200Response) GetSupportedParameters() []string`

GetSupportedParameters returns the SupportedParameters field if non-nil, zero value otherwise.

### GetSupportedParametersOk

`func (o *GetSearchEngines200Response) GetSupportedParametersOk() (*[]string, bool)`

GetSupportedParametersOk returns a tuple with the SupportedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedParameters

`func (o *GetSearchEngines200Response) SetSupportedParameters(v []string)`

SetSupportedParameters sets SupportedParameters field to given value.

### HasSupportedParameters

`func (o *GetSearchEngines200Response) HasSupportedParameters() bool`

HasSupportedParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



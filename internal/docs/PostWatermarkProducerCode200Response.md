# PostWatermarkProducerCode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to **string** | 解析出的证件绑定方式。 | [optional] 
**Code** | Pointer to **string** | 标准的 27 位服务提供者编码。 | [optional] 
**Identifier** | Pointer to **string** | 剔除补位后的主体证件原始明文。 | [optional] 
**ModelCode** | Pointer to **string** | 解析出的模型/应用码（启用扩展时存在）。 | [optional] 
**ServiceExtension** | Pointer to **bool** | 编码中是否启用了服务扩展段。 | [optional] 
**ServiceType** | Pointer to **string** | 解析出的服务角色类型（启用扩展时存在）。 | [optional] 
**SubjectCode** | Pointer to **string** | 包含补位逻辑在内的 18 位主体特征段。 | [optional] 
**SubjectType** | Pointer to **string** | 解析出的主体类型。 | [optional] 
**Valid** | Pointer to **bool** | 该编码是否合规合法。 | [optional] 

## Methods

### NewPostWatermarkProducerCode200Response

`func NewPostWatermarkProducerCode200Response() *PostWatermarkProducerCode200Response`

NewPostWatermarkProducerCode200Response instantiates a new PostWatermarkProducerCode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWatermarkProducerCode200ResponseWithDefaults

`func NewPostWatermarkProducerCode200ResponseWithDefaults() *PostWatermarkProducerCode200Response`

NewPostWatermarkProducerCode200ResponseWithDefaults instantiates a new PostWatermarkProducerCode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *PostWatermarkProducerCode200Response) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *PostWatermarkProducerCode200Response) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *PostWatermarkProducerCode200Response) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *PostWatermarkProducerCode200Response) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetCode

`func (o *PostWatermarkProducerCode200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostWatermarkProducerCode200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostWatermarkProducerCode200Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PostWatermarkProducerCode200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIdentifier

`func (o *PostWatermarkProducerCode200Response) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PostWatermarkProducerCode200Response) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PostWatermarkProducerCode200Response) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PostWatermarkProducerCode200Response) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetModelCode

`func (o *PostWatermarkProducerCode200Response) GetModelCode() string`

GetModelCode returns the ModelCode field if non-nil, zero value otherwise.

### GetModelCodeOk

`func (o *PostWatermarkProducerCode200Response) GetModelCodeOk() (*string, bool)`

GetModelCodeOk returns a tuple with the ModelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCode

`func (o *PostWatermarkProducerCode200Response) SetModelCode(v string)`

SetModelCode sets ModelCode field to given value.

### HasModelCode

`func (o *PostWatermarkProducerCode200Response) HasModelCode() bool`

HasModelCode returns a boolean if a field has been set.

### GetServiceExtension

`func (o *PostWatermarkProducerCode200Response) GetServiceExtension() bool`

GetServiceExtension returns the ServiceExtension field if non-nil, zero value otherwise.

### GetServiceExtensionOk

`func (o *PostWatermarkProducerCode200Response) GetServiceExtensionOk() (*bool, bool)`

GetServiceExtensionOk returns a tuple with the ServiceExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExtension

`func (o *PostWatermarkProducerCode200Response) SetServiceExtension(v bool)`

SetServiceExtension sets ServiceExtension field to given value.

### HasServiceExtension

`func (o *PostWatermarkProducerCode200Response) HasServiceExtension() bool`

HasServiceExtension returns a boolean if a field has been set.

### GetServiceType

`func (o *PostWatermarkProducerCode200Response) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PostWatermarkProducerCode200Response) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PostWatermarkProducerCode200Response) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PostWatermarkProducerCode200Response) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSubjectCode

`func (o *PostWatermarkProducerCode200Response) GetSubjectCode() string`

GetSubjectCode returns the SubjectCode field if non-nil, zero value otherwise.

### GetSubjectCodeOk

`func (o *PostWatermarkProducerCode200Response) GetSubjectCodeOk() (*string, bool)`

GetSubjectCodeOk returns a tuple with the SubjectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCode

`func (o *PostWatermarkProducerCode200Response) SetSubjectCode(v string)`

SetSubjectCode sets SubjectCode field to given value.

### HasSubjectCode

`func (o *PostWatermarkProducerCode200Response) HasSubjectCode() bool`

HasSubjectCode returns a boolean if a field has been set.

### GetSubjectType

`func (o *PostWatermarkProducerCode200Response) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *PostWatermarkProducerCode200Response) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *PostWatermarkProducerCode200Response) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *PostWatermarkProducerCode200Response) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetValid

`func (o *PostWatermarkProducerCode200Response) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *PostWatermarkProducerCode200Response) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *PostWatermarkProducerCode200Response) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *PostWatermarkProducerCode200Response) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



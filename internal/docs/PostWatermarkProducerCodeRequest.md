# PostWatermarkProducerCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to **string** | 主体绑定的证件类型。组织需使用统一社会信用代码；个人可选身份证、手机号、护照或网号。 | [optional] 
**Code** | Pointer to **string** | 待校验的 27 位现成编码。填写后接口将直接执行合法性校验。 | [optional] 
**Identifier** | Pointer to **string** | 证件号实际内容。长度需匹配选择的类型（如统一社会信用代码 18 位、手机号 11 位）。 | [optional] 
**ModelCode** | Pointer to **string** | 4 位自定义模型或应用码（可选）。未提供时扩展段将默认填充 00000。 | [optional] 
**ServiceType** | Pointer to **string** | 服务角色类型（仅在提供模型应用码时一同生效）。 | [optional] 
**SubjectType** | Pointer to **string** | 主体类型是组织还是个人。 | [optional] 

## Methods

### NewPostWatermarkProducerCodeRequest

`func NewPostWatermarkProducerCodeRequest() *PostWatermarkProducerCodeRequest`

NewPostWatermarkProducerCodeRequest instantiates a new PostWatermarkProducerCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWatermarkProducerCodeRequestWithDefaults

`func NewPostWatermarkProducerCodeRequestWithDefaults() *PostWatermarkProducerCodeRequest`

NewPostWatermarkProducerCodeRequestWithDefaults instantiates a new PostWatermarkProducerCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *PostWatermarkProducerCodeRequest) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *PostWatermarkProducerCodeRequest) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *PostWatermarkProducerCodeRequest) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *PostWatermarkProducerCodeRequest) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetCode

`func (o *PostWatermarkProducerCodeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostWatermarkProducerCodeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostWatermarkProducerCodeRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PostWatermarkProducerCodeRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIdentifier

`func (o *PostWatermarkProducerCodeRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PostWatermarkProducerCodeRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PostWatermarkProducerCodeRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PostWatermarkProducerCodeRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetModelCode

`func (o *PostWatermarkProducerCodeRequest) GetModelCode() string`

GetModelCode returns the ModelCode field if non-nil, zero value otherwise.

### GetModelCodeOk

`func (o *PostWatermarkProducerCodeRequest) GetModelCodeOk() (*string, bool)`

GetModelCodeOk returns a tuple with the ModelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCode

`func (o *PostWatermarkProducerCodeRequest) SetModelCode(v string)`

SetModelCode sets ModelCode field to given value.

### HasModelCode

`func (o *PostWatermarkProducerCodeRequest) HasModelCode() bool`

HasModelCode returns a boolean if a field has been set.

### GetServiceType

`func (o *PostWatermarkProducerCodeRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PostWatermarkProducerCodeRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PostWatermarkProducerCodeRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PostWatermarkProducerCodeRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSubjectType

`func (o *PostWatermarkProducerCodeRequest) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *PostWatermarkProducerCodeRequest) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *PostWatermarkProducerCodeRequest) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *PostWatermarkProducerCodeRequest) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



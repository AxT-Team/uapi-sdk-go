# GetImageQrcode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | 图片的URL。当&#x60;format&#x3D;json_url&#x60;时是临时公网URL；当&#x60;format&#x3D;json&#x60;时是Base64 Data URI。 | [optional] 

## Methods

### NewGetImageQrcode200Response

`func NewGetImageQrcode200Response() *GetImageQrcode200Response`

NewGetImageQrcode200Response instantiates a new GetImageQrcode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageQrcode200ResponseWithDefaults

`func NewGetImageQrcode200ResponseWithDefaults() *GetImageQrcode200Response`

NewGetImageQrcode200ResponseWithDefaults instantiates a new GetImageQrcode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GetImageQrcode200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetImageQrcode200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetImageQrcode200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetImageQrcode200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



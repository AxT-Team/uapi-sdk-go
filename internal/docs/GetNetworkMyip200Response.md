# GetNetworkMyip200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | 你的公网IP地址 | [optional] 
**Region** | Pointer to **string** | 地理位置，格式：国家 省份 城市 | [optional] 
**Isp** | Pointer to **string** | 运营商名称 | [optional] 
**Llc** | Pointer to **string** | 归属机构 | [optional] 
**Asn** | Pointer to **string** | 自治系统编号 | [optional] 
**Latitude** | Pointer to **float32** | 纬度 | [optional] 
**Longitude** | Pointer to **float32** | 经度 | [optional] 
**Beginip** | Pointer to **string** | IP段起始地址（标准查询） | [optional] 
**Endip** | Pointer to **string** | IP段结束地址（标准查询） | [optional] 
**District** | Pointer to **string** | 行政区（商业查询） | [optional] 

## Methods

### NewGetNetworkMyip200Response

`func NewGetNetworkMyip200Response() *GetNetworkMyip200Response`

NewGetNetworkMyip200Response instantiates a new GetNetworkMyip200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkMyip200ResponseWithDefaults

`func NewGetNetworkMyip200ResponseWithDefaults() *GetNetworkMyip200Response`

NewGetNetworkMyip200ResponseWithDefaults instantiates a new GetNetworkMyip200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GetNetworkMyip200Response) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetNetworkMyip200Response) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetNetworkMyip200Response) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetNetworkMyip200Response) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetRegion

`func (o *GetNetworkMyip200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetNetworkMyip200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetNetworkMyip200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetNetworkMyip200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIsp

`func (o *GetNetworkMyip200Response) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *GetNetworkMyip200Response) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *GetNetworkMyip200Response) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *GetNetworkMyip200Response) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetLlc

`func (o *GetNetworkMyip200Response) GetLlc() string`

GetLlc returns the Llc field if non-nil, zero value otherwise.

### GetLlcOk

`func (o *GetNetworkMyip200Response) GetLlcOk() (*string, bool)`

GetLlcOk returns a tuple with the Llc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlc

`func (o *GetNetworkMyip200Response) SetLlc(v string)`

SetLlc sets Llc field to given value.

### HasLlc

`func (o *GetNetworkMyip200Response) HasLlc() bool`

HasLlc returns a boolean if a field has been set.

### GetAsn

`func (o *GetNetworkMyip200Response) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *GetNetworkMyip200Response) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *GetNetworkMyip200Response) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *GetNetworkMyip200Response) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetLatitude

`func (o *GetNetworkMyip200Response) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GetNetworkMyip200Response) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GetNetworkMyip200Response) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GetNetworkMyip200Response) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *GetNetworkMyip200Response) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GetNetworkMyip200Response) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GetNetworkMyip200Response) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GetNetworkMyip200Response) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetBeginip

`func (o *GetNetworkMyip200Response) GetBeginip() string`

GetBeginip returns the Beginip field if non-nil, zero value otherwise.

### GetBeginipOk

`func (o *GetNetworkMyip200Response) GetBeginipOk() (*string, bool)`

GetBeginipOk returns a tuple with the Beginip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginip

`func (o *GetNetworkMyip200Response) SetBeginip(v string)`

SetBeginip sets Beginip field to given value.

### HasBeginip

`func (o *GetNetworkMyip200Response) HasBeginip() bool`

HasBeginip returns a boolean if a field has been set.

### GetEndip

`func (o *GetNetworkMyip200Response) GetEndip() string`

GetEndip returns the Endip field if non-nil, zero value otherwise.

### GetEndipOk

`func (o *GetNetworkMyip200Response) GetEndipOk() (*string, bool)`

GetEndipOk returns a tuple with the Endip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndip

`func (o *GetNetworkMyip200Response) SetEndip(v string)`

SetEndip sets Endip field to given value.

### HasEndip

`func (o *GetNetworkMyip200Response) HasEndip() bool`

HasEndip returns a boolean if a field has been set.

### GetDistrict

`func (o *GetNetworkMyip200Response) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *GetNetworkMyip200Response) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *GetNetworkMyip200Response) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *GetNetworkMyip200Response) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



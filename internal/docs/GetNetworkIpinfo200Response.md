# GetNetworkIpinfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **string** | 自治系统编号 (由GeoLite2或商业版提供) | [optional] 
**Beginip** | Pointer to **string** | IP范围起始 (仅在默认查询中提供) | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Endip** | Pointer to **string** | IP范围结束 (仅在默认查询中提供) | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Isp** | Pointer to **string** | 运营商 | [optional] 
**Latitude** | Pointer to **float32** |  | [optional] 
**Llc** | Pointer to **string** | 归属 | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**Region** | Pointer to **string** | 格式：国家 省份 城市 | [optional] 
**District** | Pointer to **string** | 行政区 (仅在商业查询中提供) | [optional] 
**AreaCode** | Pointer to **string** | 行政区划代码 (仅在商业查询中提供) | [optional] 
**CityCode** | Pointer to **string** | 城市区号 (仅在商业查询中提供) | [optional] 
**ZipCode** | Pointer to **string** | 邮政编码 (仅在商业查询中提供) | [optional] 
**TimeZone** | Pointer to **string** | 时区 (仅在商业查询中提供) | [optional] 
**Scenes** | Pointer to **string** | 应用场景 (仅在商业查询中提供) | [optional] 
**Elevation** | Pointer to **string** | 海拔（米）(仅在商业查询中提供) | [optional] 
**WeatherStation** | Pointer to **string** | 气象站代码 (仅在商业查询中提供) | [optional] 

## Methods

### NewGetNetworkIpinfo200Response

`func NewGetNetworkIpinfo200Response() *GetNetworkIpinfo200Response`

NewGetNetworkIpinfo200Response instantiates a new GetNetworkIpinfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkIpinfo200ResponseWithDefaults

`func NewGetNetworkIpinfo200ResponseWithDefaults() *GetNetworkIpinfo200Response`

NewGetNetworkIpinfo200ResponseWithDefaults instantiates a new GetNetworkIpinfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *GetNetworkIpinfo200Response) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *GetNetworkIpinfo200Response) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *GetNetworkIpinfo200Response) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *GetNetworkIpinfo200Response) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetBeginip

`func (o *GetNetworkIpinfo200Response) GetBeginip() string`

GetBeginip returns the Beginip field if non-nil, zero value otherwise.

### GetBeginipOk

`func (o *GetNetworkIpinfo200Response) GetBeginipOk() (*string, bool)`

GetBeginipOk returns a tuple with the Beginip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginip

`func (o *GetNetworkIpinfo200Response) SetBeginip(v string)`

SetBeginip sets Beginip field to given value.

### HasBeginip

`func (o *GetNetworkIpinfo200Response) HasBeginip() bool`

HasBeginip returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkIpinfo200Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkIpinfo200Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkIpinfo200Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkIpinfo200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEndip

`func (o *GetNetworkIpinfo200Response) GetEndip() string`

GetEndip returns the Endip field if non-nil, zero value otherwise.

### GetEndipOk

`func (o *GetNetworkIpinfo200Response) GetEndipOk() (*string, bool)`

GetEndipOk returns a tuple with the Endip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndip

`func (o *GetNetworkIpinfo200Response) SetEndip(v string)`

SetEndip sets Endip field to given value.

### HasEndip

`func (o *GetNetworkIpinfo200Response) HasEndip() bool`

HasEndip returns a boolean if a field has been set.

### GetIp

`func (o *GetNetworkIpinfo200Response) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetNetworkIpinfo200Response) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetNetworkIpinfo200Response) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetNetworkIpinfo200Response) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsp

`func (o *GetNetworkIpinfo200Response) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *GetNetworkIpinfo200Response) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *GetNetworkIpinfo200Response) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *GetNetworkIpinfo200Response) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetLatitude

`func (o *GetNetworkIpinfo200Response) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GetNetworkIpinfo200Response) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GetNetworkIpinfo200Response) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GetNetworkIpinfo200Response) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLlc

`func (o *GetNetworkIpinfo200Response) GetLlc() string`

GetLlc returns the Llc field if non-nil, zero value otherwise.

### GetLlcOk

`func (o *GetNetworkIpinfo200Response) GetLlcOk() (*string, bool)`

GetLlcOk returns a tuple with the Llc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlc

`func (o *GetNetworkIpinfo200Response) SetLlc(v string)`

SetLlc sets Llc field to given value.

### HasLlc

`func (o *GetNetworkIpinfo200Response) HasLlc() bool`

HasLlc returns a boolean if a field has been set.

### GetLongitude

`func (o *GetNetworkIpinfo200Response) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GetNetworkIpinfo200Response) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GetNetworkIpinfo200Response) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GetNetworkIpinfo200Response) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetRegion

`func (o *GetNetworkIpinfo200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetNetworkIpinfo200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetNetworkIpinfo200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetNetworkIpinfo200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDistrict

`func (o *GetNetworkIpinfo200Response) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *GetNetworkIpinfo200Response) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *GetNetworkIpinfo200Response) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *GetNetworkIpinfo200Response) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetAreaCode

`func (o *GetNetworkIpinfo200Response) GetAreaCode() string`

GetAreaCode returns the AreaCode field if non-nil, zero value otherwise.

### GetAreaCodeOk

`func (o *GetNetworkIpinfo200Response) GetAreaCodeOk() (*string, bool)`

GetAreaCodeOk returns a tuple with the AreaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCode

`func (o *GetNetworkIpinfo200Response) SetAreaCode(v string)`

SetAreaCode sets AreaCode field to given value.

### HasAreaCode

`func (o *GetNetworkIpinfo200Response) HasAreaCode() bool`

HasAreaCode returns a boolean if a field has been set.

### GetCityCode

`func (o *GetNetworkIpinfo200Response) GetCityCode() string`

GetCityCode returns the CityCode field if non-nil, zero value otherwise.

### GetCityCodeOk

`func (o *GetNetworkIpinfo200Response) GetCityCodeOk() (*string, bool)`

GetCityCodeOk returns a tuple with the CityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityCode

`func (o *GetNetworkIpinfo200Response) SetCityCode(v string)`

SetCityCode sets CityCode field to given value.

### HasCityCode

`func (o *GetNetworkIpinfo200Response) HasCityCode() bool`

HasCityCode returns a boolean if a field has been set.

### GetZipCode

`func (o *GetNetworkIpinfo200Response) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GetNetworkIpinfo200Response) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GetNetworkIpinfo200Response) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *GetNetworkIpinfo200Response) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetNetworkIpinfo200Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetNetworkIpinfo200Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetNetworkIpinfo200Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetNetworkIpinfo200Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetScenes

`func (o *GetNetworkIpinfo200Response) GetScenes() string`

GetScenes returns the Scenes field if non-nil, zero value otherwise.

### GetScenesOk

`func (o *GetNetworkIpinfo200Response) GetScenesOk() (*string, bool)`

GetScenesOk returns a tuple with the Scenes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenes

`func (o *GetNetworkIpinfo200Response) SetScenes(v string)`

SetScenes sets Scenes field to given value.

### HasScenes

`func (o *GetNetworkIpinfo200Response) HasScenes() bool`

HasScenes returns a boolean if a field has been set.

### GetElevation

`func (o *GetNetworkIpinfo200Response) GetElevation() string`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *GetNetworkIpinfo200Response) GetElevationOk() (*string, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *GetNetworkIpinfo200Response) SetElevation(v string)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *GetNetworkIpinfo200Response) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetWeatherStation

`func (o *GetNetworkIpinfo200Response) GetWeatherStation() string`

GetWeatherStation returns the WeatherStation field if non-nil, zero value otherwise.

### GetWeatherStationOk

`func (o *GetNetworkIpinfo200Response) GetWeatherStationOk() (*string, bool)`

GetWeatherStationOk returns a tuple with the WeatherStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeatherStation

`func (o *GetNetworkIpinfo200Response) SetWeatherStation(v string)`

SetWeatherStation sets WeatherStation field to given value.

### HasWeatherStation

`func (o *GetNetworkIpinfo200Response) HasWeatherStation() bool`

HasWeatherStation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



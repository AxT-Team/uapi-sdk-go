# GetMiscDistrict200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | 地区名称。 | [optional] 
**Level** | Pointer to **string** | 行政级别：province / city / district / street。 | [optional] 
**Country** | Pointer to **string** | 国家名称。 | [optional] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 国家代码。 | [optional] 
**Province** | Pointer to **string** | 省/州（中国数据）或一级行政区（国际数据）。 | [optional] 
**City** | Pointer to **string** | 市（仅中国数据）。 | [optional] 
**District** | Pointer to **string** | 区/县（仅中国数据）。 | [optional] 
**Street** | Pointer to **string** | 街道/乡镇（仅中国数据）。 | [optional] 
**Adcode** | Pointer to **string** | 行政区划代码（仅中国数据）。 | [optional] 
**Citycode** | Pointer to **string** | 城市区号（仅中国数据）。 | [optional] 
**Center** | Pointer to [**GetMiscDistrict200ResponseResultsInnerCenter**](GetMiscDistrict200ResponseResultsInnerCenter.md) |  | [optional] 
**Population** | Pointer to **int32** | 人口（仅国际城市数据）。 | [optional] 
**Timezone** | Pointer to **string** | 时区（仅国际城市数据），如 Asia/Tokyo。 | [optional] 

## Methods

### NewGetMiscDistrict200ResponseResultsInner

`func NewGetMiscDistrict200ResponseResultsInner() *GetMiscDistrict200ResponseResultsInner`

NewGetMiscDistrict200ResponseResultsInner instantiates a new GetMiscDistrict200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscDistrict200ResponseResultsInnerWithDefaults

`func NewGetMiscDistrict200ResponseResultsInnerWithDefaults() *GetMiscDistrict200ResponseResultsInner`

NewGetMiscDistrict200ResponseResultsInnerWithDefaults instantiates a new GetMiscDistrict200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetMiscDistrict200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMiscDistrict200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMiscDistrict200ResponseResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLevel

`func (o *GetMiscDistrict200ResponseResultsInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetMiscDistrict200ResponseResultsInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetMiscDistrict200ResponseResultsInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetCountry

`func (o *GetMiscDistrict200ResponseResultsInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetMiscDistrict200ResponseResultsInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetMiscDistrict200ResponseResultsInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *GetMiscDistrict200ResponseResultsInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GetMiscDistrict200ResponseResultsInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GetMiscDistrict200ResponseResultsInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetProvince

`func (o *GetMiscDistrict200ResponseResultsInner) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *GetMiscDistrict200ResponseResultsInner) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *GetMiscDistrict200ResponseResultsInner) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetCity

`func (o *GetMiscDistrict200ResponseResultsInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetMiscDistrict200ResponseResultsInner) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetMiscDistrict200ResponseResultsInner) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistrict

`func (o *GetMiscDistrict200ResponseResultsInner) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *GetMiscDistrict200ResponseResultsInner) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *GetMiscDistrict200ResponseResultsInner) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetStreet

`func (o *GetMiscDistrict200ResponseResultsInner) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GetMiscDistrict200ResponseResultsInner) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GetMiscDistrict200ResponseResultsInner) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetAdcode

`func (o *GetMiscDistrict200ResponseResultsInner) GetAdcode() string`

GetAdcode returns the Adcode field if non-nil, zero value otherwise.

### GetAdcodeOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetAdcodeOk() (*string, bool)`

GetAdcodeOk returns a tuple with the Adcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdcode

`func (o *GetMiscDistrict200ResponseResultsInner) SetAdcode(v string)`

SetAdcode sets Adcode field to given value.

### HasAdcode

`func (o *GetMiscDistrict200ResponseResultsInner) HasAdcode() bool`

HasAdcode returns a boolean if a field has been set.

### GetCitycode

`func (o *GetMiscDistrict200ResponseResultsInner) GetCitycode() string`

GetCitycode returns the Citycode field if non-nil, zero value otherwise.

### GetCitycodeOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetCitycodeOk() (*string, bool)`

GetCitycodeOk returns a tuple with the Citycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitycode

`func (o *GetMiscDistrict200ResponseResultsInner) SetCitycode(v string)`

SetCitycode sets Citycode field to given value.

### HasCitycode

`func (o *GetMiscDistrict200ResponseResultsInner) HasCitycode() bool`

HasCitycode returns a boolean if a field has been set.

### GetCenter

`func (o *GetMiscDistrict200ResponseResultsInner) GetCenter() GetMiscDistrict200ResponseResultsInnerCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetCenterOk() (*GetMiscDistrict200ResponseResultsInnerCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *GetMiscDistrict200ResponseResultsInner) SetCenter(v GetMiscDistrict200ResponseResultsInnerCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *GetMiscDistrict200ResponseResultsInner) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetPopulation

`func (o *GetMiscDistrict200ResponseResultsInner) GetPopulation() int32`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetPopulationOk() (*int32, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *GetMiscDistrict200ResponseResultsInner) SetPopulation(v int32)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *GetMiscDistrict200ResponseResultsInner) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.

### GetTimezone

`func (o *GetMiscDistrict200ResponseResultsInner) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetMiscDistrict200ResponseResultsInner) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetMiscDistrict200ResponseResultsInner) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetMiscDistrict200ResponseResultsInner) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



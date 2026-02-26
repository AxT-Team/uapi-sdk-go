# GetMiscWeather200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Province** | Pointer to **string** | 省份 | [optional] 
**City** | Pointer to **string** | 城市名 | [optional] 
**Adcode** | Pointer to **string** | 行政区划代码（部分数据源可能为空） | [optional] 
**Weather** | Pointer to **string** | 天气状况描述。默认返回中文，传 &#x60;lang&#x3D;en&#x60; 时返回英文。非固定枚举。 | [optional] 
**Temperature** | Pointer to **float32** | 当前温度 °C | [optional] 
**WindDirection** | Pointer to **string** | 风向 | [optional] 
**WindPower** | Pointer to **string** | 风力等级 | [optional] 
**Humidity** | Pointer to **float32** | 相对湿度 % | [optional] 
**ReportTime** | Pointer to **string** | 数据更新时间 | [optional] 
**FeelsLike** | Pointer to **float32** | 体感温度 °C（extended&#x3D;true 时返回） | [optional] 
**Visibility** | Pointer to **float32** | 能见度 km（extended&#x3D;true 时返回） | [optional] 
**Pressure** | Pointer to **float32** | 气压 hPa（extended&#x3D;true 时返回） | [optional] 
**Uv** | Pointer to **float32** | 紫外线指数（extended&#x3D;true 时返回） | [optional] 
**Precipitation** | Pointer to **float32** | 当前降水量 mm（extended&#x3D;true 时返回） | [optional] 
**Cloud** | Pointer to **float32** | 云量 %（extended&#x3D;true 时返回） | [optional] 
**Aqi** | Pointer to **float32** | 空气质量指数 0-500（extended&#x3D;true 时返回） | [optional] 
**AqiLevel** | Pointer to **float32** | AQI 等级 1-6（extended&#x3D;true 时返回） | [optional] 
**AqiCategory** | Pointer to **string** | AQI 等级描述（优/良/轻度污染/中度污染/重度污染/严重污染）（extended&#x3D;true 时返回） | [optional] 
**AqiPrimary** | Pointer to **string** | 主要污染物（如 PM2.5、PM10、O3 等）（extended&#x3D;true 时返回） | [optional] 
**AirPollutants** | Pointer to [**GetMiscWeather200ResponseAirPollutants**](GetMiscWeather200ResponseAirPollutants.md) |  | [optional] 
**TempMax** | Pointer to **float32** | 当天最高温 °C（forecast&#x3D;true 时返回） | [optional] 
**TempMin** | Pointer to **float32** | 当天最低温 °C（forecast&#x3D;true 时返回） | [optional] 
**Forecast** | Pointer to [**[]GetMiscWeather200ResponseForecastInner**](GetMiscWeather200ResponseForecastInner.md) | 多天天气预报，最多7天（forecast&#x3D;true 时返回） | [optional] 
**HourlyForecast** | Pointer to [**[]GetMiscWeather200ResponseHourlyForecastInner**](GetMiscWeather200ResponseHourlyForecastInner.md) | 逐小时预报，最多24小时（hourly&#x3D;true 时返回） | [optional] 
**MinutelyPrecip** | Pointer to [**GetMiscWeather200ResponseMinutelyPrecip**](GetMiscWeather200ResponseMinutelyPrecip.md) |  | [optional] 
**LifeIndices** | Pointer to [**GetMiscWeather200ResponseLifeIndices**](GetMiscWeather200ResponseLifeIndices.md) |  | [optional] 

## Methods

### NewGetMiscWeather200Response

`func NewGetMiscWeather200Response() *GetMiscWeather200Response`

NewGetMiscWeather200Response instantiates a new GetMiscWeather200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscWeather200ResponseWithDefaults

`func NewGetMiscWeather200ResponseWithDefaults() *GetMiscWeather200Response`

NewGetMiscWeather200ResponseWithDefaults instantiates a new GetMiscWeather200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvince

`func (o *GetMiscWeather200Response) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *GetMiscWeather200Response) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *GetMiscWeather200Response) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *GetMiscWeather200Response) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetCity

`func (o *GetMiscWeather200Response) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetMiscWeather200Response) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetMiscWeather200Response) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetMiscWeather200Response) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetAdcode

`func (o *GetMiscWeather200Response) GetAdcode() string`

GetAdcode returns the Adcode field if non-nil, zero value otherwise.

### GetAdcodeOk

`func (o *GetMiscWeather200Response) GetAdcodeOk() (*string, bool)`

GetAdcodeOk returns a tuple with the Adcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdcode

`func (o *GetMiscWeather200Response) SetAdcode(v string)`

SetAdcode sets Adcode field to given value.

### HasAdcode

`func (o *GetMiscWeather200Response) HasAdcode() bool`

HasAdcode returns a boolean if a field has been set.

### GetWeather

`func (o *GetMiscWeather200Response) GetWeather() string`

GetWeather returns the Weather field if non-nil, zero value otherwise.

### GetWeatherOk

`func (o *GetMiscWeather200Response) GetWeatherOk() (*string, bool)`

GetWeatherOk returns a tuple with the Weather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeather

`func (o *GetMiscWeather200Response) SetWeather(v string)`

SetWeather sets Weather field to given value.

### HasWeather

`func (o *GetMiscWeather200Response) HasWeather() bool`

HasWeather returns a boolean if a field has been set.

### GetTemperature

`func (o *GetMiscWeather200Response) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetMiscWeather200Response) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetMiscWeather200Response) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetMiscWeather200Response) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetWindDirection

`func (o *GetMiscWeather200Response) GetWindDirection() string`

GetWindDirection returns the WindDirection field if non-nil, zero value otherwise.

### GetWindDirectionOk

`func (o *GetMiscWeather200Response) GetWindDirectionOk() (*string, bool)`

GetWindDirectionOk returns a tuple with the WindDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindDirection

`func (o *GetMiscWeather200Response) SetWindDirection(v string)`

SetWindDirection sets WindDirection field to given value.

### HasWindDirection

`func (o *GetMiscWeather200Response) HasWindDirection() bool`

HasWindDirection returns a boolean if a field has been set.

### GetWindPower

`func (o *GetMiscWeather200Response) GetWindPower() string`

GetWindPower returns the WindPower field if non-nil, zero value otherwise.

### GetWindPowerOk

`func (o *GetMiscWeather200Response) GetWindPowerOk() (*string, bool)`

GetWindPowerOk returns a tuple with the WindPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindPower

`func (o *GetMiscWeather200Response) SetWindPower(v string)`

SetWindPower sets WindPower field to given value.

### HasWindPower

`func (o *GetMiscWeather200Response) HasWindPower() bool`

HasWindPower returns a boolean if a field has been set.

### GetHumidity

`func (o *GetMiscWeather200Response) GetHumidity() float32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *GetMiscWeather200Response) GetHumidityOk() (*float32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *GetMiscWeather200Response) SetHumidity(v float32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *GetMiscWeather200Response) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetReportTime

`func (o *GetMiscWeather200Response) GetReportTime() string`

GetReportTime returns the ReportTime field if non-nil, zero value otherwise.

### GetReportTimeOk

`func (o *GetMiscWeather200Response) GetReportTimeOk() (*string, bool)`

GetReportTimeOk returns a tuple with the ReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTime

`func (o *GetMiscWeather200Response) SetReportTime(v string)`

SetReportTime sets ReportTime field to given value.

### HasReportTime

`func (o *GetMiscWeather200Response) HasReportTime() bool`

HasReportTime returns a boolean if a field has been set.

### GetFeelsLike

`func (o *GetMiscWeather200Response) GetFeelsLike() float32`

GetFeelsLike returns the FeelsLike field if non-nil, zero value otherwise.

### GetFeelsLikeOk

`func (o *GetMiscWeather200Response) GetFeelsLikeOk() (*float32, bool)`

GetFeelsLikeOk returns a tuple with the FeelsLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeelsLike

`func (o *GetMiscWeather200Response) SetFeelsLike(v float32)`

SetFeelsLike sets FeelsLike field to given value.

### HasFeelsLike

`func (o *GetMiscWeather200Response) HasFeelsLike() bool`

HasFeelsLike returns a boolean if a field has been set.

### GetVisibility

`func (o *GetMiscWeather200Response) GetVisibility() float32`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetMiscWeather200Response) GetVisibilityOk() (*float32, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetMiscWeather200Response) SetVisibility(v float32)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetMiscWeather200Response) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetPressure

`func (o *GetMiscWeather200Response) GetPressure() float32`

GetPressure returns the Pressure field if non-nil, zero value otherwise.

### GetPressureOk

`func (o *GetMiscWeather200Response) GetPressureOk() (*float32, bool)`

GetPressureOk returns a tuple with the Pressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressure

`func (o *GetMiscWeather200Response) SetPressure(v float32)`

SetPressure sets Pressure field to given value.

### HasPressure

`func (o *GetMiscWeather200Response) HasPressure() bool`

HasPressure returns a boolean if a field has been set.

### GetUv

`func (o *GetMiscWeather200Response) GetUv() float32`

GetUv returns the Uv field if non-nil, zero value otherwise.

### GetUvOk

`func (o *GetMiscWeather200Response) GetUvOk() (*float32, bool)`

GetUvOk returns a tuple with the Uv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUv

`func (o *GetMiscWeather200Response) SetUv(v float32)`

SetUv sets Uv field to given value.

### HasUv

`func (o *GetMiscWeather200Response) HasUv() bool`

HasUv returns a boolean if a field has been set.

### GetPrecipitation

`func (o *GetMiscWeather200Response) GetPrecipitation() float32`

GetPrecipitation returns the Precipitation field if non-nil, zero value otherwise.

### GetPrecipitationOk

`func (o *GetMiscWeather200Response) GetPrecipitationOk() (*float32, bool)`

GetPrecipitationOk returns a tuple with the Precipitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecipitation

`func (o *GetMiscWeather200Response) SetPrecipitation(v float32)`

SetPrecipitation sets Precipitation field to given value.

### HasPrecipitation

`func (o *GetMiscWeather200Response) HasPrecipitation() bool`

HasPrecipitation returns a boolean if a field has been set.

### GetCloud

`func (o *GetMiscWeather200Response) GetCloud() float32`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetMiscWeather200Response) GetCloudOk() (*float32, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetMiscWeather200Response) SetCloud(v float32)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetMiscWeather200Response) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetAqi

`func (o *GetMiscWeather200Response) GetAqi() float32`

GetAqi returns the Aqi field if non-nil, zero value otherwise.

### GetAqiOk

`func (o *GetMiscWeather200Response) GetAqiOk() (*float32, bool)`

GetAqiOk returns a tuple with the Aqi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAqi

`func (o *GetMiscWeather200Response) SetAqi(v float32)`

SetAqi sets Aqi field to given value.

### HasAqi

`func (o *GetMiscWeather200Response) HasAqi() bool`

HasAqi returns a boolean if a field has been set.

### GetAqiLevel

`func (o *GetMiscWeather200Response) GetAqiLevel() float32`

GetAqiLevel returns the AqiLevel field if non-nil, zero value otherwise.

### GetAqiLevelOk

`func (o *GetMiscWeather200Response) GetAqiLevelOk() (*float32, bool)`

GetAqiLevelOk returns a tuple with the AqiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAqiLevel

`func (o *GetMiscWeather200Response) SetAqiLevel(v float32)`

SetAqiLevel sets AqiLevel field to given value.

### HasAqiLevel

`func (o *GetMiscWeather200Response) HasAqiLevel() bool`

HasAqiLevel returns a boolean if a field has been set.

### GetAqiCategory

`func (o *GetMiscWeather200Response) GetAqiCategory() string`

GetAqiCategory returns the AqiCategory field if non-nil, zero value otherwise.

### GetAqiCategoryOk

`func (o *GetMiscWeather200Response) GetAqiCategoryOk() (*string, bool)`

GetAqiCategoryOk returns a tuple with the AqiCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAqiCategory

`func (o *GetMiscWeather200Response) SetAqiCategory(v string)`

SetAqiCategory sets AqiCategory field to given value.

### HasAqiCategory

`func (o *GetMiscWeather200Response) HasAqiCategory() bool`

HasAqiCategory returns a boolean if a field has been set.

### GetAqiPrimary

`func (o *GetMiscWeather200Response) GetAqiPrimary() string`

GetAqiPrimary returns the AqiPrimary field if non-nil, zero value otherwise.

### GetAqiPrimaryOk

`func (o *GetMiscWeather200Response) GetAqiPrimaryOk() (*string, bool)`

GetAqiPrimaryOk returns a tuple with the AqiPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAqiPrimary

`func (o *GetMiscWeather200Response) SetAqiPrimary(v string)`

SetAqiPrimary sets AqiPrimary field to given value.

### HasAqiPrimary

`func (o *GetMiscWeather200Response) HasAqiPrimary() bool`

HasAqiPrimary returns a boolean if a field has been set.

### GetAirPollutants

`func (o *GetMiscWeather200Response) GetAirPollutants() GetMiscWeather200ResponseAirPollutants`

GetAirPollutants returns the AirPollutants field if non-nil, zero value otherwise.

### GetAirPollutantsOk

`func (o *GetMiscWeather200Response) GetAirPollutantsOk() (*GetMiscWeather200ResponseAirPollutants, bool)`

GetAirPollutantsOk returns a tuple with the AirPollutants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirPollutants

`func (o *GetMiscWeather200Response) SetAirPollutants(v GetMiscWeather200ResponseAirPollutants)`

SetAirPollutants sets AirPollutants field to given value.

### HasAirPollutants

`func (o *GetMiscWeather200Response) HasAirPollutants() bool`

HasAirPollutants returns a boolean if a field has been set.

### GetTempMax

`func (o *GetMiscWeather200Response) GetTempMax() float32`

GetTempMax returns the TempMax field if non-nil, zero value otherwise.

### GetTempMaxOk

`func (o *GetMiscWeather200Response) GetTempMaxOk() (*float32, bool)`

GetTempMaxOk returns a tuple with the TempMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMax

`func (o *GetMiscWeather200Response) SetTempMax(v float32)`

SetTempMax sets TempMax field to given value.

### HasTempMax

`func (o *GetMiscWeather200Response) HasTempMax() bool`

HasTempMax returns a boolean if a field has been set.

### GetTempMin

`func (o *GetMiscWeather200Response) GetTempMin() float32`

GetTempMin returns the TempMin field if non-nil, zero value otherwise.

### GetTempMinOk

`func (o *GetMiscWeather200Response) GetTempMinOk() (*float32, bool)`

GetTempMinOk returns a tuple with the TempMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMin

`func (o *GetMiscWeather200Response) SetTempMin(v float32)`

SetTempMin sets TempMin field to given value.

### HasTempMin

`func (o *GetMiscWeather200Response) HasTempMin() bool`

HasTempMin returns a boolean if a field has been set.

### GetForecast

`func (o *GetMiscWeather200Response) GetForecast() []GetMiscWeather200ResponseForecastInner`

GetForecast returns the Forecast field if non-nil, zero value otherwise.

### GetForecastOk

`func (o *GetMiscWeather200Response) GetForecastOk() (*[]GetMiscWeather200ResponseForecastInner, bool)`

GetForecastOk returns a tuple with the Forecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecast

`func (o *GetMiscWeather200Response) SetForecast(v []GetMiscWeather200ResponseForecastInner)`

SetForecast sets Forecast field to given value.

### HasForecast

`func (o *GetMiscWeather200Response) HasForecast() bool`

HasForecast returns a boolean if a field has been set.

### GetHourlyForecast

`func (o *GetMiscWeather200Response) GetHourlyForecast() []GetMiscWeather200ResponseHourlyForecastInner`

GetHourlyForecast returns the HourlyForecast field if non-nil, zero value otherwise.

### GetHourlyForecastOk

`func (o *GetMiscWeather200Response) GetHourlyForecastOk() (*[]GetMiscWeather200ResponseHourlyForecastInner, bool)`

GetHourlyForecastOk returns a tuple with the HourlyForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyForecast

`func (o *GetMiscWeather200Response) SetHourlyForecast(v []GetMiscWeather200ResponseHourlyForecastInner)`

SetHourlyForecast sets HourlyForecast field to given value.

### HasHourlyForecast

`func (o *GetMiscWeather200Response) HasHourlyForecast() bool`

HasHourlyForecast returns a boolean if a field has been set.

### GetMinutelyPrecip

`func (o *GetMiscWeather200Response) GetMinutelyPrecip() GetMiscWeather200ResponseMinutelyPrecip`

GetMinutelyPrecip returns the MinutelyPrecip field if non-nil, zero value otherwise.

### GetMinutelyPrecipOk

`func (o *GetMiscWeather200Response) GetMinutelyPrecipOk() (*GetMiscWeather200ResponseMinutelyPrecip, bool)`

GetMinutelyPrecipOk returns a tuple with the MinutelyPrecip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutelyPrecip

`func (o *GetMiscWeather200Response) SetMinutelyPrecip(v GetMiscWeather200ResponseMinutelyPrecip)`

SetMinutelyPrecip sets MinutelyPrecip field to given value.

### HasMinutelyPrecip

`func (o *GetMiscWeather200Response) HasMinutelyPrecip() bool`

HasMinutelyPrecip returns a boolean if a field has been set.

### GetLifeIndices

`func (o *GetMiscWeather200Response) GetLifeIndices() GetMiscWeather200ResponseLifeIndices`

GetLifeIndices returns the LifeIndices field if non-nil, zero value otherwise.

### GetLifeIndicesOk

`func (o *GetMiscWeather200Response) GetLifeIndicesOk() (*GetMiscWeather200ResponseLifeIndices, bool)`

GetLifeIndicesOk returns a tuple with the LifeIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeIndices

`func (o *GetMiscWeather200Response) SetLifeIndices(v GetMiscWeather200ResponseLifeIndices)`

SetLifeIndices sets LifeIndices field to given value.

### HasLifeIndices

`func (o *GetMiscWeather200Response) HasLifeIndices() bool`

HasLifeIndices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



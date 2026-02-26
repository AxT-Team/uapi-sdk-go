# GetMiscWeather200ResponseForecastInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | 日期 YYYY-MM-DD | [optional] 
**Week** | Pointer to **string** | 星期几（&#x60;lang&#x3D;en&#x60; 时返回英文星期） | [optional] 
**TempMax** | Pointer to **float32** | 最高温度 °C | [optional] 
**TempMin** | Pointer to **float32** | 最低温度 °C | [optional] 
**WeatherDay** | Pointer to **string** | 白天天气（&#x60;lang&#x3D;en&#x60; 时返回英文） | [optional] 
**WeatherNight** | Pointer to **string** | 夜间天气（&#x60;lang&#x3D;en&#x60; 时返回英文） | [optional] 
**WindDirDay** | Pointer to **string** | 白天风向（可选，&#x60;lang&#x3D;en&#x60; 时返回英文） | [optional] 
**WindDirNight** | Pointer to **string** | 夜间风向（可选，&#x60;lang&#x3D;en&#x60; 时返回英文） | [optional] 
**WindScaleDay** | Pointer to **string** | 白天风力（可选，&#x60;lang&#x3D;en&#x60; 时返回英文） | [optional] 
**WindScaleNight** | Pointer to **string** | 夜间风力（可选，&#x60;lang&#x3D;en&#x60; 时返回英文） | [optional] 
**WindSpeedDay** | Pointer to **float32** | 白天风速 km/h（可选） | [optional] 
**Humidity** | Pointer to **float32** | 湿度 %（可选） | [optional] 
**Precip** | Pointer to **float32** | 降水量 mm（可选） | [optional] 
**Visibility** | Pointer to **float32** | 能见度 km（可选） | [optional] 
**UvIndex** | Pointer to **float32** | 紫外线指数（可选） | [optional] 
**Sunrise** | Pointer to **string** | 日出时间 HH:MM（可选） | [optional] 
**Sunset** | Pointer to **string** | 日落时间 HH:MM（可选） | [optional] 

## Methods

### NewGetMiscWeather200ResponseForecastInner

`func NewGetMiscWeather200ResponseForecastInner() *GetMiscWeather200ResponseForecastInner`

NewGetMiscWeather200ResponseForecastInner instantiates a new GetMiscWeather200ResponseForecastInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscWeather200ResponseForecastInnerWithDefaults

`func NewGetMiscWeather200ResponseForecastInnerWithDefaults() *GetMiscWeather200ResponseForecastInner`

NewGetMiscWeather200ResponseForecastInnerWithDefaults instantiates a new GetMiscWeather200ResponseForecastInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetMiscWeather200ResponseForecastInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetMiscWeather200ResponseForecastInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetMiscWeather200ResponseForecastInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetMiscWeather200ResponseForecastInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetWeek

`func (o *GetMiscWeather200ResponseForecastInner) GetWeek() string`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWeekOk() (*string, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *GetMiscWeather200ResponseForecastInner) SetWeek(v string)`

SetWeek sets Week field to given value.

### HasWeek

`func (o *GetMiscWeather200ResponseForecastInner) HasWeek() bool`

HasWeek returns a boolean if a field has been set.

### GetTempMax

`func (o *GetMiscWeather200ResponseForecastInner) GetTempMax() float32`

GetTempMax returns the TempMax field if non-nil, zero value otherwise.

### GetTempMaxOk

`func (o *GetMiscWeather200ResponseForecastInner) GetTempMaxOk() (*float32, bool)`

GetTempMaxOk returns a tuple with the TempMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMax

`func (o *GetMiscWeather200ResponseForecastInner) SetTempMax(v float32)`

SetTempMax sets TempMax field to given value.

### HasTempMax

`func (o *GetMiscWeather200ResponseForecastInner) HasTempMax() bool`

HasTempMax returns a boolean if a field has been set.

### GetTempMin

`func (o *GetMiscWeather200ResponseForecastInner) GetTempMin() float32`

GetTempMin returns the TempMin field if non-nil, zero value otherwise.

### GetTempMinOk

`func (o *GetMiscWeather200ResponseForecastInner) GetTempMinOk() (*float32, bool)`

GetTempMinOk returns a tuple with the TempMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMin

`func (o *GetMiscWeather200ResponseForecastInner) SetTempMin(v float32)`

SetTempMin sets TempMin field to given value.

### HasTempMin

`func (o *GetMiscWeather200ResponseForecastInner) HasTempMin() bool`

HasTempMin returns a boolean if a field has been set.

### GetWeatherDay

`func (o *GetMiscWeather200ResponseForecastInner) GetWeatherDay() string`

GetWeatherDay returns the WeatherDay field if non-nil, zero value otherwise.

### GetWeatherDayOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWeatherDayOk() (*string, bool)`

GetWeatherDayOk returns a tuple with the WeatherDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeatherDay

`func (o *GetMiscWeather200ResponseForecastInner) SetWeatherDay(v string)`

SetWeatherDay sets WeatherDay field to given value.

### HasWeatherDay

`func (o *GetMiscWeather200ResponseForecastInner) HasWeatherDay() bool`

HasWeatherDay returns a boolean if a field has been set.

### GetWeatherNight

`func (o *GetMiscWeather200ResponseForecastInner) GetWeatherNight() string`

GetWeatherNight returns the WeatherNight field if non-nil, zero value otherwise.

### GetWeatherNightOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWeatherNightOk() (*string, bool)`

GetWeatherNightOk returns a tuple with the WeatherNight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeatherNight

`func (o *GetMiscWeather200ResponseForecastInner) SetWeatherNight(v string)`

SetWeatherNight sets WeatherNight field to given value.

### HasWeatherNight

`func (o *GetMiscWeather200ResponseForecastInner) HasWeatherNight() bool`

HasWeatherNight returns a boolean if a field has been set.

### GetWindDirDay

`func (o *GetMiscWeather200ResponseForecastInner) GetWindDirDay() string`

GetWindDirDay returns the WindDirDay field if non-nil, zero value otherwise.

### GetWindDirDayOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWindDirDayOk() (*string, bool)`

GetWindDirDayOk returns a tuple with the WindDirDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindDirDay

`func (o *GetMiscWeather200ResponseForecastInner) SetWindDirDay(v string)`

SetWindDirDay sets WindDirDay field to given value.

### HasWindDirDay

`func (o *GetMiscWeather200ResponseForecastInner) HasWindDirDay() bool`

HasWindDirDay returns a boolean if a field has been set.

### GetWindDirNight

`func (o *GetMiscWeather200ResponseForecastInner) GetWindDirNight() string`

GetWindDirNight returns the WindDirNight field if non-nil, zero value otherwise.

### GetWindDirNightOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWindDirNightOk() (*string, bool)`

GetWindDirNightOk returns a tuple with the WindDirNight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindDirNight

`func (o *GetMiscWeather200ResponseForecastInner) SetWindDirNight(v string)`

SetWindDirNight sets WindDirNight field to given value.

### HasWindDirNight

`func (o *GetMiscWeather200ResponseForecastInner) HasWindDirNight() bool`

HasWindDirNight returns a boolean if a field has been set.

### GetWindScaleDay

`func (o *GetMiscWeather200ResponseForecastInner) GetWindScaleDay() string`

GetWindScaleDay returns the WindScaleDay field if non-nil, zero value otherwise.

### GetWindScaleDayOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWindScaleDayOk() (*string, bool)`

GetWindScaleDayOk returns a tuple with the WindScaleDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindScaleDay

`func (o *GetMiscWeather200ResponseForecastInner) SetWindScaleDay(v string)`

SetWindScaleDay sets WindScaleDay field to given value.

### HasWindScaleDay

`func (o *GetMiscWeather200ResponseForecastInner) HasWindScaleDay() bool`

HasWindScaleDay returns a boolean if a field has been set.

### GetWindScaleNight

`func (o *GetMiscWeather200ResponseForecastInner) GetWindScaleNight() string`

GetWindScaleNight returns the WindScaleNight field if non-nil, zero value otherwise.

### GetWindScaleNightOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWindScaleNightOk() (*string, bool)`

GetWindScaleNightOk returns a tuple with the WindScaleNight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindScaleNight

`func (o *GetMiscWeather200ResponseForecastInner) SetWindScaleNight(v string)`

SetWindScaleNight sets WindScaleNight field to given value.

### HasWindScaleNight

`func (o *GetMiscWeather200ResponseForecastInner) HasWindScaleNight() bool`

HasWindScaleNight returns a boolean if a field has been set.

### GetWindSpeedDay

`func (o *GetMiscWeather200ResponseForecastInner) GetWindSpeedDay() float32`

GetWindSpeedDay returns the WindSpeedDay field if non-nil, zero value otherwise.

### GetWindSpeedDayOk

`func (o *GetMiscWeather200ResponseForecastInner) GetWindSpeedDayOk() (*float32, bool)`

GetWindSpeedDayOk returns a tuple with the WindSpeedDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindSpeedDay

`func (o *GetMiscWeather200ResponseForecastInner) SetWindSpeedDay(v float32)`

SetWindSpeedDay sets WindSpeedDay field to given value.

### HasWindSpeedDay

`func (o *GetMiscWeather200ResponseForecastInner) HasWindSpeedDay() bool`

HasWindSpeedDay returns a boolean if a field has been set.

### GetHumidity

`func (o *GetMiscWeather200ResponseForecastInner) GetHumidity() float32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *GetMiscWeather200ResponseForecastInner) GetHumidityOk() (*float32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *GetMiscWeather200ResponseForecastInner) SetHumidity(v float32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *GetMiscWeather200ResponseForecastInner) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetPrecip

`func (o *GetMiscWeather200ResponseForecastInner) GetPrecip() float32`

GetPrecip returns the Precip field if non-nil, zero value otherwise.

### GetPrecipOk

`func (o *GetMiscWeather200ResponseForecastInner) GetPrecipOk() (*float32, bool)`

GetPrecipOk returns a tuple with the Precip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecip

`func (o *GetMiscWeather200ResponseForecastInner) SetPrecip(v float32)`

SetPrecip sets Precip field to given value.

### HasPrecip

`func (o *GetMiscWeather200ResponseForecastInner) HasPrecip() bool`

HasPrecip returns a boolean if a field has been set.

### GetVisibility

`func (o *GetMiscWeather200ResponseForecastInner) GetVisibility() float32`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetMiscWeather200ResponseForecastInner) GetVisibilityOk() (*float32, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetMiscWeather200ResponseForecastInner) SetVisibility(v float32)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetMiscWeather200ResponseForecastInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetUvIndex

`func (o *GetMiscWeather200ResponseForecastInner) GetUvIndex() float32`

GetUvIndex returns the UvIndex field if non-nil, zero value otherwise.

### GetUvIndexOk

`func (o *GetMiscWeather200ResponseForecastInner) GetUvIndexOk() (*float32, bool)`

GetUvIndexOk returns a tuple with the UvIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUvIndex

`func (o *GetMiscWeather200ResponseForecastInner) SetUvIndex(v float32)`

SetUvIndex sets UvIndex field to given value.

### HasUvIndex

`func (o *GetMiscWeather200ResponseForecastInner) HasUvIndex() bool`

HasUvIndex returns a boolean if a field has been set.

### GetSunrise

`func (o *GetMiscWeather200ResponseForecastInner) GetSunrise() string`

GetSunrise returns the Sunrise field if non-nil, zero value otherwise.

### GetSunriseOk

`func (o *GetMiscWeather200ResponseForecastInner) GetSunriseOk() (*string, bool)`

GetSunriseOk returns a tuple with the Sunrise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunrise

`func (o *GetMiscWeather200ResponseForecastInner) SetSunrise(v string)`

SetSunrise sets Sunrise field to given value.

### HasSunrise

`func (o *GetMiscWeather200ResponseForecastInner) HasSunrise() bool`

HasSunrise returns a boolean if a field has been set.

### GetSunset

`func (o *GetMiscWeather200ResponseForecastInner) GetSunset() string`

GetSunset returns the Sunset field if non-nil, zero value otherwise.

### GetSunsetOk

`func (o *GetMiscWeather200ResponseForecastInner) GetSunsetOk() (*string, bool)`

GetSunsetOk returns a tuple with the Sunset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunset

`func (o *GetMiscWeather200ResponseForecastInner) SetSunset(v string)`

SetSunset sets Sunset field to given value.

### HasSunset

`func (o *GetMiscWeather200ResponseForecastInner) HasSunset() bool`

HasSunset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



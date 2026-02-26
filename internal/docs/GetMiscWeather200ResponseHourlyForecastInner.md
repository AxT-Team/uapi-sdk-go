# GetMiscWeather200ResponseHourlyForecastInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **string** | 预报时间（ISO8601 或 YYYY-MM-DD HH:MM） | [optional] 
**Temperature** | Pointer to **float32** | 温度 °C | [optional] 
**Weather** | Pointer to **string** | 天气状况 | [optional] 
**WindDirection** | Pointer to **string** | 风向（可选） | [optional] 
**WindSpeed** | Pointer to **float32** | 风速 km/h（可选） | [optional] 
**WindScale** | Pointer to **string** | 风力等级（可选） | [optional] 
**Humidity** | Pointer to **float32** | 湿度 %（可选） | [optional] 
**Precip** | Pointer to **float32** | 降水量 mm（可选） | [optional] 
**Pressure** | Pointer to **float32** | 气压 hPa（可选） | [optional] 
**Cloud** | Pointer to **float32** | 云量 %（可选） | [optional] 
**FeelsLike** | Pointer to **float32** | 体感温度 °C（可选） | [optional] 
**DewPoint** | Pointer to **float32** | 露点温度 °C（可选） | [optional] 
**Visibility** | Pointer to **float32** | 能见度 km（可选） | [optional] 
**Pop** | Pointer to **float32** | 降水概率 %（可选） | [optional] 
**UvIndex** | Pointer to **float32** | 紫外线指数（可选） | [optional] 

## Methods

### NewGetMiscWeather200ResponseHourlyForecastInner

`func NewGetMiscWeather200ResponseHourlyForecastInner() *GetMiscWeather200ResponseHourlyForecastInner`

NewGetMiscWeather200ResponseHourlyForecastInner instantiates a new GetMiscWeather200ResponseHourlyForecastInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiscWeather200ResponseHourlyForecastInnerWithDefaults

`func NewGetMiscWeather200ResponseHourlyForecastInnerWithDefaults() *GetMiscWeather200ResponseHourlyForecastInner`

NewGetMiscWeather200ResponseHourlyForecastInnerWithDefaults instantiates a new GetMiscWeather200ResponseHourlyForecastInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTemperature

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetWeather

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWeather() string`

GetWeather returns the Weather field if non-nil, zero value otherwise.

### GetWeatherOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWeatherOk() (*string, bool)`

GetWeatherOk returns a tuple with the Weather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeather

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetWeather(v string)`

SetWeather sets Weather field to given value.

### HasWeather

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasWeather() bool`

HasWeather returns a boolean if a field has been set.

### GetWindDirection

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWindDirection() string`

GetWindDirection returns the WindDirection field if non-nil, zero value otherwise.

### GetWindDirectionOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWindDirectionOk() (*string, bool)`

GetWindDirectionOk returns a tuple with the WindDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindDirection

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetWindDirection(v string)`

SetWindDirection sets WindDirection field to given value.

### HasWindDirection

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasWindDirection() bool`

HasWindDirection returns a boolean if a field has been set.

### GetWindSpeed

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWindSpeed() float32`

GetWindSpeed returns the WindSpeed field if non-nil, zero value otherwise.

### GetWindSpeedOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWindSpeedOk() (*float32, bool)`

GetWindSpeedOk returns a tuple with the WindSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindSpeed

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetWindSpeed(v float32)`

SetWindSpeed sets WindSpeed field to given value.

### HasWindSpeed

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasWindSpeed() bool`

HasWindSpeed returns a boolean if a field has been set.

### GetWindScale

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWindScale() string`

GetWindScale returns the WindScale field if non-nil, zero value otherwise.

### GetWindScaleOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetWindScaleOk() (*string, bool)`

GetWindScaleOk returns a tuple with the WindScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindScale

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetWindScale(v string)`

SetWindScale sets WindScale field to given value.

### HasWindScale

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasWindScale() bool`

HasWindScale returns a boolean if a field has been set.

### GetHumidity

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetHumidity() float32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetHumidityOk() (*float32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetHumidity(v float32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetPrecip

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetPrecip() float32`

GetPrecip returns the Precip field if non-nil, zero value otherwise.

### GetPrecipOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetPrecipOk() (*float32, bool)`

GetPrecipOk returns a tuple with the Precip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecip

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetPrecip(v float32)`

SetPrecip sets Precip field to given value.

### HasPrecip

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasPrecip() bool`

HasPrecip returns a boolean if a field has been set.

### GetPressure

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetPressure() float32`

GetPressure returns the Pressure field if non-nil, zero value otherwise.

### GetPressureOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetPressureOk() (*float32, bool)`

GetPressureOk returns a tuple with the Pressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressure

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetPressure(v float32)`

SetPressure sets Pressure field to given value.

### HasPressure

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasPressure() bool`

HasPressure returns a boolean if a field has been set.

### GetCloud

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetCloud() float32`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetCloudOk() (*float32, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetCloud(v float32)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetFeelsLike

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetFeelsLike() float32`

GetFeelsLike returns the FeelsLike field if non-nil, zero value otherwise.

### GetFeelsLikeOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetFeelsLikeOk() (*float32, bool)`

GetFeelsLikeOk returns a tuple with the FeelsLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeelsLike

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetFeelsLike(v float32)`

SetFeelsLike sets FeelsLike field to given value.

### HasFeelsLike

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasFeelsLike() bool`

HasFeelsLike returns a boolean if a field has been set.

### GetDewPoint

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetDewPoint() float32`

GetDewPoint returns the DewPoint field if non-nil, zero value otherwise.

### GetDewPointOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetDewPointOk() (*float32, bool)`

GetDewPointOk returns a tuple with the DewPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDewPoint

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetDewPoint(v float32)`

SetDewPoint sets DewPoint field to given value.

### HasDewPoint

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasDewPoint() bool`

HasDewPoint returns a boolean if a field has been set.

### GetVisibility

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetVisibility() float32`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetVisibilityOk() (*float32, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetVisibility(v float32)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetPop

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetPop() float32`

GetPop returns the Pop field if non-nil, zero value otherwise.

### GetPopOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetPopOk() (*float32, bool)`

GetPopOk returns a tuple with the Pop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPop

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetPop(v float32)`

SetPop sets Pop field to given value.

### HasPop

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasPop() bool`

HasPop returns a boolean if a field has been set.

### GetUvIndex

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetUvIndex() float32`

GetUvIndex returns the UvIndex field if non-nil, zero value otherwise.

### GetUvIndexOk

`func (o *GetMiscWeather200ResponseHourlyForecastInner) GetUvIndexOk() (*float32, bool)`

GetUvIndexOk returns a tuple with the UvIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUvIndex

`func (o *GetMiscWeather200ResponseHourlyForecastInner) SetUvIndex(v float32)`

SetUvIndex sets UvIndex field to given value.

### HasUvIndex

`func (o *GetMiscWeather200ResponseHourlyForecastInner) HasUvIndex() bool`

HasUvIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



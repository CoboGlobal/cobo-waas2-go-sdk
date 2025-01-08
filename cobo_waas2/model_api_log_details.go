/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiLogDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLogDetails{}

// ApiLogDetails The information about an API log.
type ApiLogDetails struct {
	// A unique identifier for the API log, used for tracking.
	LogId string `json:"log_id"`
	// The HTTP method used for the API request.
	ApiMethod string `json:"api_method"`
	// The endpoint of the API request.
	ApiEndpoint string `json:"api_endpoint"`
	// The HTTP status code returned by the API request.
	StatusCode int32 `json:"status_code"`
	// The client's IP address that made the API request.
	IpAddress string `json:"ip_address"`
	// The time when the API request was created, in Unix timestamp format, measured in milliseconds.
	RequestTimestamp int64 `json:"request_timestamp"`
	// The API key used to call the API. For more details, refer to [API key](https://www.cobo.com/developers/v2/guides/overview/cobo-auth#api-key).
	ApiKey string `json:"api_key"`
	// The response body of the API request.
	ResponseBody string `json:"response_body"`
	// The query parameters of the API request.
	QueryParams string `json:"query_params"`
	// The request body of the API request.
	RequestBody string `json:"request_body"`
}

type _ApiLogDetails ApiLogDetails

// NewApiLogDetails instantiates a new ApiLogDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLogDetails(logId string, apiMethod string, apiEndpoint string, statusCode int32, ipAddress string, requestTimestamp int64, apiKey string, responseBody string, queryParams string, requestBody string) *ApiLogDetails {
	this := ApiLogDetails{}
	this.LogId = logId
	this.ApiMethod = apiMethod
	this.ApiEndpoint = apiEndpoint
	this.StatusCode = statusCode
	this.IpAddress = ipAddress
	this.RequestTimestamp = requestTimestamp
	this.ApiKey = apiKey
	this.ResponseBody = responseBody
	this.QueryParams = queryParams
	this.RequestBody = requestBody
	return &this
}

// NewApiLogDetailsWithDefaults instantiates a new ApiLogDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLogDetailsWithDefaults() *ApiLogDetails {
	this := ApiLogDetails{}
	return &this
}

// GetLogId returns the LogId field value
func (o *ApiLogDetails) GetLogId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogId
}

// GetLogIdOk returns a tuple with the LogId field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogId, true
}

// SetLogId sets field value
func (o *ApiLogDetails) SetLogId(v string) {
	o.LogId = v
}

// GetApiMethod returns the ApiMethod field value
func (o *ApiLogDetails) GetApiMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiMethod
}

// GetApiMethodOk returns a tuple with the ApiMethod field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetApiMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiMethod, true
}

// SetApiMethod sets field value
func (o *ApiLogDetails) SetApiMethod(v string) {
	o.ApiMethod = v
}

// GetApiEndpoint returns the ApiEndpoint field value
func (o *ApiLogDetails) GetApiEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiEndpoint
}

// GetApiEndpointOk returns a tuple with the ApiEndpoint field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetApiEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiEndpoint, true
}

// SetApiEndpoint sets field value
func (o *ApiLogDetails) SetApiEndpoint(v string) {
	o.ApiEndpoint = v
}

// GetStatusCode returns the StatusCode field value
func (o *ApiLogDetails) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *ApiLogDetails) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetIpAddress returns the IpAddress field value
func (o *ApiLogDetails) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *ApiLogDetails) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetRequestTimestamp returns the RequestTimestamp field value
func (o *ApiLogDetails) GetRequestTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetRequestTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTimestamp, true
}

// SetRequestTimestamp sets field value
func (o *ApiLogDetails) SetRequestTimestamp(v int64) {
	o.RequestTimestamp = v
}

// GetApiKey returns the ApiKey field value
func (o *ApiLogDetails) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *ApiLogDetails) SetApiKey(v string) {
	o.ApiKey = v
}

// GetResponseBody returns the ResponseBody field value
func (o *ApiLogDetails) GetResponseBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseBody
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetResponseBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseBody, true
}

// SetResponseBody sets field value
func (o *ApiLogDetails) SetResponseBody(v string) {
	o.ResponseBody = v
}

// GetQueryParams returns the QueryParams field value
func (o *ApiLogDetails) GetQueryParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryParams
}

// GetQueryParamsOk returns a tuple with the QueryParams field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetQueryParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryParams, true
}

// SetQueryParams sets field value
func (o *ApiLogDetails) SetQueryParams(v string) {
	o.QueryParams = v
}

// GetRequestBody returns the RequestBody field value
func (o *ApiLogDetails) GetRequestBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestBody
}

// GetRequestBodyOk returns a tuple with the RequestBody field value
// and a boolean to check if the value has been set.
func (o *ApiLogDetails) GetRequestBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestBody, true
}

// SetRequestBody sets field value
func (o *ApiLogDetails) SetRequestBody(v string) {
	o.RequestBody = v
}

func (o ApiLogDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLogDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["log_id"] = o.LogId
	toSerialize["api_method"] = o.ApiMethod
	toSerialize["api_endpoint"] = o.ApiEndpoint
	toSerialize["status_code"] = o.StatusCode
	toSerialize["ip_address"] = o.IpAddress
	toSerialize["request_timestamp"] = o.RequestTimestamp
	toSerialize["api_key"] = o.ApiKey
	toSerialize["response_body"] = o.ResponseBody
	toSerialize["query_params"] = o.QueryParams
	toSerialize["request_body"] = o.RequestBody
	return toSerialize, nil
}

func (o *ApiLogDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"log_id",
		"api_method",
		"api_endpoint",
		"status_code",
		"ip_address",
		"request_timestamp",
		"api_key",
		"response_body",
		"query_params",
		"request_body",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiLogDetails := _ApiLogDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiLogDetails)

	if err != nil {
		return err
	}

	*o = ApiLogDetails(varApiLogDetails)

	return err
}

type NullableApiLogDetails struct {
	value *ApiLogDetails
	isSet bool
}

func (v NullableApiLogDetails) Get() *ApiLogDetails {
	return v.value
}

func (v *NullableApiLogDetails) Set(val *ApiLogDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLogDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLogDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLogDetails(val *ApiLogDetails) *NullableApiLogDetails {
	return &NullableApiLogDetails{value: val, isSet: true}
}

func (v NullableApiLogDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLogDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



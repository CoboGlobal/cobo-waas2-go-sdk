/*
Cobo Wallet as a Service 2.0

Cobo WaaS 2.0 enables you to programmatically access Cobo's full suite of crypto wallet technologies with powerful and flexible access controls.  # Wallet technologies - Custodial Wallet - MPC Wallet - Smart Contract Wallet (Based on Safe{Wallet}) - Exchange Wallet  # Risk Control technologies - Workflow - Access Control List (ACL)  # Risk Control targets - Wallet Management   - User/team and their permission management   - Risk control configurations, e.g. whitelist, blacklist, rate-limiting etc. - Blockchain Interaction   - Crypto transfer   - Smart Contract Invocation  # Important HTTPS only. RESTful, resource oriented  # Get Started Set up your APIs or get authorization  # Authentication and Authorization CoboAuth  # Request and Response application/json  # Error Handling  ### Common error codes | Error Code | Description | | -- | -- |  ### API-specific error codes For error codes that are dedicated to a specific API, see the Error codes section in each API specification, for example, /v3/wallets.  # Rate and Usage Limiting  # Idempotent Request  # Pagination # Support [Developer Hub](https://cobo.com/developers) 

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponse{}

// ErrorResponse The data for error response.
type ErrorResponse struct {
	// Indicates if the API operation was successful. Always false for errors.
	Success bool `json:"success"`
	// A machine-readable error code.`
	ErrorCode int32 `json:"error_code"`
	// A human-readable error description for users.
	ErrorDescription string `json:"error_description"`
	// A unique ID for the error log, mainly used for debugging.
	ErrorId string `json:"error_id"`
}

type _ErrorResponse ErrorResponse

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse(success bool, errorCode int32, errorDescription string, errorId string) *ErrorResponse {
	this := ErrorResponse{}
	this.Success = success
	this.ErrorCode = errorCode
	this.ErrorDescription = errorDescription
	this.ErrorId = errorId
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	var success bool = false
	this.Success = success
	return &this
}

// GetSuccess returns the Success field value
func (o *ErrorResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ErrorResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetErrorCode returns the ErrorCode field value
func (o *ErrorResponse) GetErrorCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ErrorResponse) SetErrorCode(v int32) {
	o.ErrorCode = v
}

// GetErrorDescription returns the ErrorDescription field value
func (o *ErrorResponse) GetErrorDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorDescription, true
}

// SetErrorDescription sets field value
func (o *ErrorResponse) SetErrorDescription(v string) {
	o.ErrorDescription = v
}

// GetErrorId returns the ErrorId field value
func (o *ErrorResponse) GetErrorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorId, true
}

// SetErrorId sets field value
func (o *ErrorResponse) SetErrorId(v string) {
	o.ErrorId = v
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["error_description"] = o.ErrorDescription
	toSerialize["error_id"] = o.ErrorId
	return toSerialize, nil
}

func (o *ErrorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"error_code",
		"error_description",
		"error_id",
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

	varErrorResponse := _ErrorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponse)

	if err != nil {
		return err
	}

	*o = ErrorResponse(varErrorResponse)

	return err
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



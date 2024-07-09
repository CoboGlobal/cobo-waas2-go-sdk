/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the GetToken200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetToken200Response{}

// GetToken200Response struct for GetToken200Response
type GetToken200Response struct {
	// The token that can be used to access protected resources.
	AccessToken *string `json:"access_token,omitempty"`
	// The type of the token issued.
	TokenType *string `json:"token_type,omitempty"`
	// The scope of the access token.
	Scope *string `json:"scope,omitempty"`
	// The time in seconds before the access token expires.
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// The token that can be used to obtain a new access token.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// NewGetToken200Response instantiates a new GetToken200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetToken200Response() *GetToken200Response {
	this := GetToken200Response{}
	return &this
}

// NewGetToken200ResponseWithDefaults instantiates a new GetToken200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetToken200ResponseWithDefaults() *GetToken200Response {
	this := GetToken200Response{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *GetToken200Response) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken200Response) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *GetToken200Response) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *GetToken200Response) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *GetToken200Response) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken200Response) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *GetToken200Response) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *GetToken200Response) SetTokenType(v string) {
	o.TokenType = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GetToken200Response) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken200Response) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GetToken200Response) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GetToken200Response) SetScope(v string) {
	o.Scope = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *GetToken200Response) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken200Response) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *GetToken200Response) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *GetToken200Response) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *GetToken200Response) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken200Response) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *GetToken200Response) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *GetToken200Response) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o GetToken200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetToken200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return toSerialize, nil
}

type NullableGetToken200Response struct {
	value *GetToken200Response
	isSet bool
}

func (v NullableGetToken200Response) Get() *GetToken200Response {
	return v.value
}

func (v *NullableGetToken200Response) Set(val *GetToken200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetToken200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetToken200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetToken200Response(val *GetToken200Response) *NullableGetToken200Response {
	return &NullableGetToken200Response{value: val, isSet: true}
}

func (v NullableGetToken200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetToken200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



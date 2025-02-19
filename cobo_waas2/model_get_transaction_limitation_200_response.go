/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the GetTransactionLimitation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransactionLimitation200Response{}

// GetTransactionLimitation200Response struct for GetTransactionLimitation200Response
type GetTransactionLimitation200Response struct {
	// A list of virtual asset service providers (VASP) you can select as the transaction source or destination.
	VaspList []Vasp `json:"vasp_list,omitempty"`
	// Indicates whether the transaction amount exceeds a predefined threshold. If exceeded, additional information is required when filling Travel Rule details. - `true`: Threshold exceeded. - `false`: Threshold not exceeded. 
	IsThresholdReached *bool `json:"is_threshold_reached,omitempty"`
	// A human-readable, time-sensitive message to be signed by the wallet owner. The message contains key information including the wallet address, a unique nonce, and a timestamp. Signing this message confirms ownership of the wallet and allows the operation to proceed. 
	SelfCustodyWalletChallenge *string `json:"self_custody_wallet_challenge,omitempty"`
	// A list of self-custody wallet providers you can select as the transaction source or destination.
	ConnectWalletList []string `json:"connect_wallet_list,omitempty"`
}

// NewGetTransactionLimitation200Response instantiates a new GetTransactionLimitation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionLimitation200Response() *GetTransactionLimitation200Response {
	this := GetTransactionLimitation200Response{}
	return &this
}

// NewGetTransactionLimitation200ResponseWithDefaults instantiates a new GetTransactionLimitation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionLimitation200ResponseWithDefaults() *GetTransactionLimitation200Response {
	this := GetTransactionLimitation200Response{}
	return &this
}

// GetVaspList returns the VaspList field value if set, zero value otherwise.
func (o *GetTransactionLimitation200Response) GetVaspList() []Vasp {
	if o == nil || IsNil(o.VaspList) {
		var ret []Vasp
		return ret
	}
	return o.VaspList
}

// GetVaspListOk returns a tuple with the VaspList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionLimitation200Response) GetVaspListOk() ([]Vasp, bool) {
	if o == nil || IsNil(o.VaspList) {
		return nil, false
	}
	return o.VaspList, true
}

// HasVaspList returns a boolean if a field has been set.
func (o *GetTransactionLimitation200Response) HasVaspList() bool {
	if o != nil && !IsNil(o.VaspList) {
		return true
	}

	return false
}

// SetVaspList gets a reference to the given []Vasp and assigns it to the VaspList field.
func (o *GetTransactionLimitation200Response) SetVaspList(v []Vasp) {
	o.VaspList = v
}

// GetIsThresholdReached returns the IsThresholdReached field value if set, zero value otherwise.
func (o *GetTransactionLimitation200Response) GetIsThresholdReached() bool {
	if o == nil || IsNil(o.IsThresholdReached) {
		var ret bool
		return ret
	}
	return *o.IsThresholdReached
}

// GetIsThresholdReachedOk returns a tuple with the IsThresholdReached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionLimitation200Response) GetIsThresholdReachedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsThresholdReached) {
		return nil, false
	}
	return o.IsThresholdReached, true
}

// HasIsThresholdReached returns a boolean if a field has been set.
func (o *GetTransactionLimitation200Response) HasIsThresholdReached() bool {
	if o != nil && !IsNil(o.IsThresholdReached) {
		return true
	}

	return false
}

// SetIsThresholdReached gets a reference to the given bool and assigns it to the IsThresholdReached field.
func (o *GetTransactionLimitation200Response) SetIsThresholdReached(v bool) {
	o.IsThresholdReached = &v
}

// GetSelfCustodyWalletChallenge returns the SelfCustodyWalletChallenge field value if set, zero value otherwise.
func (o *GetTransactionLimitation200Response) GetSelfCustodyWalletChallenge() string {
	if o == nil || IsNil(o.SelfCustodyWalletChallenge) {
		var ret string
		return ret
	}
	return *o.SelfCustodyWalletChallenge
}

// GetSelfCustodyWalletChallengeOk returns a tuple with the SelfCustodyWalletChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionLimitation200Response) GetSelfCustodyWalletChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfCustodyWalletChallenge) {
		return nil, false
	}
	return o.SelfCustodyWalletChallenge, true
}

// HasSelfCustodyWalletChallenge returns a boolean if a field has been set.
func (o *GetTransactionLimitation200Response) HasSelfCustodyWalletChallenge() bool {
	if o != nil && !IsNil(o.SelfCustodyWalletChallenge) {
		return true
	}

	return false
}

// SetSelfCustodyWalletChallenge gets a reference to the given string and assigns it to the SelfCustodyWalletChallenge field.
func (o *GetTransactionLimitation200Response) SetSelfCustodyWalletChallenge(v string) {
	o.SelfCustodyWalletChallenge = &v
}

// GetConnectWalletList returns the ConnectWalletList field value if set, zero value otherwise.
func (o *GetTransactionLimitation200Response) GetConnectWalletList() []string {
	if o == nil || IsNil(o.ConnectWalletList) {
		var ret []string
		return ret
	}
	return o.ConnectWalletList
}

// GetConnectWalletListOk returns a tuple with the ConnectWalletList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionLimitation200Response) GetConnectWalletListOk() ([]string, bool) {
	if o == nil || IsNil(o.ConnectWalletList) {
		return nil, false
	}
	return o.ConnectWalletList, true
}

// HasConnectWalletList returns a boolean if a field has been set.
func (o *GetTransactionLimitation200Response) HasConnectWalletList() bool {
	if o != nil && !IsNil(o.ConnectWalletList) {
		return true
	}

	return false
}

// SetConnectWalletList gets a reference to the given []string and assigns it to the ConnectWalletList field.
func (o *GetTransactionLimitation200Response) SetConnectWalletList(v []string) {
	o.ConnectWalletList = v
}

func (o GetTransactionLimitation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransactionLimitation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VaspList) {
		toSerialize["vasp_list"] = o.VaspList
	}
	if !IsNil(o.IsThresholdReached) {
		toSerialize["is_threshold_reached"] = o.IsThresholdReached
	}
	if !IsNil(o.SelfCustodyWalletChallenge) {
		toSerialize["self_custody_wallet_challenge"] = o.SelfCustodyWalletChallenge
	}
	if !IsNil(o.ConnectWalletList) {
		toSerialize["connect_wallet_list"] = o.ConnectWalletList
	}
	return toSerialize, nil
}

type NullableGetTransactionLimitation200Response struct {
	value *GetTransactionLimitation200Response
	isSet bool
}

func (v NullableGetTransactionLimitation200Response) Get() *GetTransactionLimitation200Response {
	return v.value
}

func (v *NullableGetTransactionLimitation200Response) Set(val *GetTransactionLimitation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionLimitation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionLimitation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionLimitation200Response(val *GetTransactionLimitation200Response) *NullableGetTransactionLimitation200Response {
	return &NullableGetTransactionLimitation200Response{value: val, isSet: true}
}

func (v NullableGetTransactionLimitation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionLimitation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



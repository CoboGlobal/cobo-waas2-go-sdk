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

// checks if the AssetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetInfo{}

// AssetInfo The asset information.
type AssetInfo struct {
	// The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account.
	AssetId string `json:"asset_id"`
	// The asset symbol. You can use the value for display purposes.
	DisplayCode *string `json:"display_code,omitempty"`
	// The description of the asset.
	Description *string `json:"description,omitempty"`
	// The URL of the asset icon.
	IconUrl *string `json:"icon_url,omitempty"`
}

type _AssetInfo AssetInfo

// NewAssetInfo instantiates a new AssetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetInfo(assetId string) *AssetInfo {
	this := AssetInfo{}
	this.AssetId = assetId
	return &this
}

// NewAssetInfoWithDefaults instantiates a new AssetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetInfoWithDefaults() *AssetInfo {
	this := AssetInfo{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *AssetInfo) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *AssetInfo) SetAssetId(v string) {
	o.AssetId = v
}

// GetDisplayCode returns the DisplayCode field value if set, zero value otherwise.
func (o *AssetInfo) GetDisplayCode() string {
	if o == nil || IsNil(o.DisplayCode) {
		var ret string
		return ret
	}
	return *o.DisplayCode
}

// GetDisplayCodeOk returns a tuple with the DisplayCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetDisplayCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayCode) {
		return nil, false
	}
	return o.DisplayCode, true
}

// HasDisplayCode returns a boolean if a field has been set.
func (o *AssetInfo) HasDisplayCode() bool {
	if o != nil && !IsNil(o.DisplayCode) {
		return true
	}

	return false
}

// SetDisplayCode gets a reference to the given string and assigns it to the DisplayCode field.
func (o *AssetInfo) SetDisplayCode(v string) {
	o.DisplayCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssetInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssetInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssetInfo) SetDescription(v string) {
	o.Description = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *AssetInfo) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfo) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *AssetInfo) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *AssetInfo) SetIconUrl(v string) {
	o.IconUrl = &v
}

func (o AssetInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_id"] = o.AssetId
	if !IsNil(o.DisplayCode) {
		toSerialize["display_code"] = o.DisplayCode
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	return toSerialize, nil
}

func (o *AssetInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_id",
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

	varAssetInfo := _AssetInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssetInfo)

	if err != nil {
		return err
	}

	*o = AssetInfo(varAssetInfo)

	return err
}

type NullableAssetInfo struct {
	value *AssetInfo
	isSet bool
}

func (v NullableAssetInfo) Get() *AssetInfo {
	return v.value
}

func (v *NullableAssetInfo) Set(val *AssetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetInfo(val *AssetInfo) *NullableAssetInfo {
	return &NullableAssetInfo{value: val, isSet: true}
}

func (v NullableAssetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



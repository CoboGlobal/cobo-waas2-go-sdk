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

// checks if the Transfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transfer{}

// Transfer The base data for transfer transaction.
type Transfer struct {
	// Unique id of the request.
	RequestId string `json:"request_id"`
	RequestType string `json:"request_type"`
	Source TransferSource `json:"source"`
	// ID of the token. Unique in all chains scope.
	TokenId string `json:"token_id"`
	// Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5) 
	Amount string `json:"amount"`
	Destination TransferDestination `json:"destination"`
	// The category names for transfer.
	CategoryNames []string `json:"category_names,omitempty"`
	// The description for transfer.
	Description *string `json:"description,omitempty"`
	Fee *TransactionFee `json:"fee,omitempty"`
}

type _Transfer Transfer

// NewTransfer instantiates a new Transfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfer(requestId string, requestType string, source TransferSource, tokenId string, amount string, destination TransferDestination) *Transfer {
	this := Transfer{}
	this.RequestId = requestId
	this.RequestType = requestType
	this.Source = source
	this.TokenId = tokenId
	this.Amount = amount
	this.Destination = destination
	return &this
}

// NewTransferWithDefaults instantiates a new Transfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferWithDefaults() *Transfer {
	this := Transfer{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *Transfer) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *Transfer) SetRequestId(v string) {
	o.RequestId = v
}

// GetRequestType returns the RequestType field value
func (o *Transfer) GetRequestType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetRequestTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *Transfer) SetRequestType(v string) {
	o.RequestType = v
}

// GetSource returns the Source field value
func (o *Transfer) GetSource() TransferSource {
	if o == nil {
		var ret TransferSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetSourceOk() (*TransferSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Transfer) SetSource(v TransferSource) {
	o.Source = v
}

// GetTokenId returns the TokenId field value
func (o *Transfer) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *Transfer) SetTokenId(v string) {
	o.TokenId = v
}

// GetAmount returns the Amount field value
func (o *Transfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Transfer) SetAmount(v string) {
	o.Amount = v
}

// GetDestination returns the Destination field value
func (o *Transfer) GetDestination() TransferDestination {
	if o == nil {
		var ret TransferDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetDestinationOk() (*TransferDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *Transfer) SetDestination(v TransferDestination) {
	o.Destination = v
}

// GetCategoryNames returns the CategoryNames field value if set, zero value otherwise.
func (o *Transfer) GetCategoryNames() []string {
	if o == nil || IsNil(o.CategoryNames) {
		var ret []string
		return ret
	}
	return o.CategoryNames
}

// GetCategoryNamesOk returns a tuple with the CategoryNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetCategoryNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoryNames) {
		return nil, false
	}
	return o.CategoryNames, true
}

// HasCategoryNames returns a boolean if a field has been set.
func (o *Transfer) HasCategoryNames() bool {
	if o != nil && !IsNil(o.CategoryNames) {
		return true
	}

	return false
}

// SetCategoryNames gets a reference to the given []string and assigns it to the CategoryNames field.
func (o *Transfer) SetCategoryNames(v []string) {
	o.CategoryNames = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Transfer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Transfer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Transfer) SetDescription(v string) {
	o.Description = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *Transfer) GetFee() TransactionFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetFeeOk() (*TransactionFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *Transfer) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionFee and assigns it to the Fee field.
func (o *Transfer) SetFee(v TransactionFee) {
	o.Fee = &v
}

func (o Transfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["request_type"] = o.RequestType
	toSerialize["source"] = o.Source
	toSerialize["token_id"] = o.TokenId
	toSerialize["amount"] = o.Amount
	toSerialize["destination"] = o.Destination
	if !IsNil(o.CategoryNames) {
		toSerialize["category_names"] = o.CategoryNames
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	return toSerialize, nil
}

func (o *Transfer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"request_type",
		"source",
		"token_id",
		"amount",
		"destination",
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

	varTransfer := _Transfer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransfer)

	if err != nil {
		return err
	}

	*o = Transfer(varTransfer)

	return err
}

type NullableTransfer struct {
	value *Transfer
	isSet bool
}

func (v NullableTransfer) Get() *Transfer {
	return v.value
}

func (v *NullableTransfer) Set(val *Transfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfer(val *Transfer) *NullableTransfer {
	return &NullableTransfer{value: val, isSet: true}
}

func (v NullableTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



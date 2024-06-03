# ContractCallDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressStr** | Pointer to **string** | To address | [optional] 
**Value** | Pointer to **string** | Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5)  | [optional] 
**Calldata** | Pointer to **string** | calldata for this transaction. Commonly used as part of contract interaction.  | [optional] 

## Methods

### NewContractCallDestination

`func NewContractCallDestination() *ContractCallDestination`

NewContractCallDestination instantiates a new ContractCallDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallDestinationWithDefaults

`func NewContractCallDestinationWithDefaults() *ContractCallDestination`

NewContractCallDestinationWithDefaults instantiates a new ContractCallDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressStr

`func (o *ContractCallDestination) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *ContractCallDestination) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *ContractCallDestination) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.

### HasAddressStr

`func (o *ContractCallDestination) HasAddressStr() bool`

HasAddressStr returns a boolean if a field has been set.

### GetValue

`func (o *ContractCallDestination) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContractCallDestination) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContractCallDestination) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContractCallDestination) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *ContractCallDestination) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *ContractCallDestination) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *ContractCallDestination) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.

### HasCalldata

`func (o *ContractCallDestination) HasCalldata() bool`

HasCalldata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



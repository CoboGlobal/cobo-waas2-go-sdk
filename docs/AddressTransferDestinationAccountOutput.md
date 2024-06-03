# AddressTransferDestinationAccountOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressStr** | Pointer to **string** | Destination address | [optional] 
**Memo** | Pointer to **string** | Destination address memo | [optional] 
**Amount** | Pointer to **string** | Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5)  | [optional] 

## Methods

### NewAddressTransferDestinationAccountOutput

`func NewAddressTransferDestinationAccountOutput() *AddressTransferDestinationAccountOutput`

NewAddressTransferDestinationAccountOutput instantiates a new AddressTransferDestinationAccountOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferDestinationAccountOutputWithDefaults

`func NewAddressTransferDestinationAccountOutputWithDefaults() *AddressTransferDestinationAccountOutput`

NewAddressTransferDestinationAccountOutputWithDefaults instantiates a new AddressTransferDestinationAccountOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressStr

`func (o *AddressTransferDestinationAccountOutput) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *AddressTransferDestinationAccountOutput) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *AddressTransferDestinationAccountOutput) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.

### HasAddressStr

`func (o *AddressTransferDestinationAccountOutput) HasAddressStr() bool`

HasAddressStr returns a boolean if a field has been set.

### GetMemo

`func (o *AddressTransferDestinationAccountOutput) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AddressTransferDestinationAccountOutput) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AddressTransferDestinationAccountOutput) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *AddressTransferDestinationAccountOutput) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAmount

`func (o *AddressTransferDestinationAccountOutput) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransferDestinationAccountOutput) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransferDestinationAccountOutput) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AddressTransferDestinationAccountOutput) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



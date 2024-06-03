# TransferDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransferDestinationType**](TransferDestinationType.md) |  | 
**AccountOutput** | Pointer to [**AddressTransferDestinationAccountOutput**](AddressTransferDestinationAccountOutput.md) |  | [optional] 
**UtxoOutputs** | Pointer to [**AddressTransferDestinationUtxoOutputs**](AddressTransferDestinationUtxoOutputs.md) |  | [optional] 
**WalletId** | **string** | Unique id of the wallet to transfer to. | 
**SubWalletId** | **string** | Exchange trading account or any sub wallet info for transfer. | 
**Amount** | Pointer to **string** | Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5)  | [optional] 

## Methods

### NewTransferDestination

`func NewTransferDestination(destinationType TransferDestinationType, walletId string, subWalletId string, ) *TransferDestination`

NewTransferDestination instantiates a new TransferDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferDestinationWithDefaults

`func NewTransferDestinationWithDefaults() *TransferDestination`

NewTransferDestinationWithDefaults instantiates a new TransferDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransferDestination) GetDestinationType() TransferDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransferDestination) GetDestinationTypeOk() (*TransferDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransferDestination) SetDestinationType(v TransferDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetAccountOutput

`func (o *TransferDestination) GetAccountOutput() AddressTransferDestinationAccountOutput`

GetAccountOutput returns the AccountOutput field if non-nil, zero value otherwise.

### GetAccountOutputOk

`func (o *TransferDestination) GetAccountOutputOk() (*AddressTransferDestinationAccountOutput, bool)`

GetAccountOutputOk returns a tuple with the AccountOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOutput

`func (o *TransferDestination) SetAccountOutput(v AddressTransferDestinationAccountOutput)`

SetAccountOutput sets AccountOutput field to given value.

### HasAccountOutput

`func (o *TransferDestination) HasAccountOutput() bool`

HasAccountOutput returns a boolean if a field has been set.

### GetUtxoOutputs

`func (o *TransferDestination) GetUtxoOutputs() AddressTransferDestinationUtxoOutputs`

GetUtxoOutputs returns the UtxoOutputs field if non-nil, zero value otherwise.

### GetUtxoOutputsOk

`func (o *TransferDestination) GetUtxoOutputsOk() (*AddressTransferDestinationUtxoOutputs, bool)`

GetUtxoOutputsOk returns a tuple with the UtxoOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoOutputs

`func (o *TransferDestination) SetUtxoOutputs(v AddressTransferDestinationUtxoOutputs)`

SetUtxoOutputs sets UtxoOutputs field to given value.

### HasUtxoOutputs

`func (o *TransferDestination) HasUtxoOutputs() bool`

HasUtxoOutputs returns a boolean if a field has been set.

### GetWalletId

`func (o *TransferDestination) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransferDestination) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransferDestination) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetSubWalletId

`func (o *TransferDestination) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *TransferDestination) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *TransferDestination) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.


### GetAmount

`func (o *TransferDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransferDestination) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



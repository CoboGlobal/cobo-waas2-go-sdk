# TransferDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransferDestinationType**](TransferDestinationType.md) |  | 
**AddressStr** | **string** | Destination address | 
**Memo** | Pointer to **string** | Destination address memo | [optional] 
**WalletId** | **string** | Unique id of the wallet to transfer to. | 
**SubWalletId** | **string** | Exchange trading account or any sub wallet info for transfer. | 

## Methods

### NewTransferDestination

`func NewTransferDestination(destinationType TransferDestinationType, addressStr string, walletId string, subWalletId string, ) *TransferDestination`

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


### GetAddressStr

`func (o *TransferDestination) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *TransferDestination) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *TransferDestination) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


### GetMemo

`func (o *TransferDestination) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransferDestination) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransferDestination) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransferDestination) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



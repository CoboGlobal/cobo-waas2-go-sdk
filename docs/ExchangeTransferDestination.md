# ExchangeTransferDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransferDestinationType**](TransferDestinationType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**SubWalletId** | **string** | The exchange trading account or the sub-wallet ID. | 
**Amount** | Pointer to **string** | The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 

## Methods

### NewExchangeTransferDestination

`func NewExchangeTransferDestination(destinationType TransferDestinationType, walletId string, subWalletId string, ) *ExchangeTransferDestination`

NewExchangeTransferDestination instantiates a new ExchangeTransferDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeTransferDestinationWithDefaults

`func NewExchangeTransferDestinationWithDefaults() *ExchangeTransferDestination`

NewExchangeTransferDestinationWithDefaults instantiates a new ExchangeTransferDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *ExchangeTransferDestination) GetDestinationType() TransferDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ExchangeTransferDestination) GetDestinationTypeOk() (*TransferDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ExchangeTransferDestination) SetDestinationType(v TransferDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetWalletId

`func (o *ExchangeTransferDestination) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExchangeTransferDestination) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExchangeTransferDestination) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetSubWalletId

`func (o *ExchangeTransferDestination) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *ExchangeTransferDestination) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *ExchangeTransferDestination) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.


### GetAmount

`func (o *ExchangeTransferDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExchangeTransferDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExchangeTransferDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExchangeTransferDestination) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



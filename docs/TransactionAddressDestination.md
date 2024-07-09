# TransactionAddressDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
**AssetId** | Pointer to **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | [optional] 
**AccountOutput** | Pointer to [**TransactionAddressDestinationAccountOutput**](TransactionAddressDestinationAccountOutput.md) |  | [optional] 
**UtxoOutputs** | Pointer to [**TransactionAddressDestinationUtxoOutputs**](TransactionAddressDestinationUtxoOutputs.md) |  | [optional] 

## Methods

### NewTransactionAddressDestination

`func NewTransactionAddressDestination(destinationType TransactionDestinationType, tokenId string, ) *TransactionAddressDestination`

NewTransactionAddressDestination instantiates a new TransactionAddressDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAddressDestinationWithDefaults

`func NewTransactionAddressDestinationWithDefaults() *TransactionAddressDestination`

NewTransactionAddressDestinationWithDefaults instantiates a new TransactionAddressDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionAddressDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionAddressDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionAddressDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetTokenId

`func (o *TransactionAddressDestination) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionAddressDestination) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionAddressDestination) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAssetId

`func (o *TransactionAddressDestination) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionAddressDestination) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionAddressDestination) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *TransactionAddressDestination) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAccountOutput

`func (o *TransactionAddressDestination) GetAccountOutput() TransactionAddressDestinationAccountOutput`

GetAccountOutput returns the AccountOutput field if non-nil, zero value otherwise.

### GetAccountOutputOk

`func (o *TransactionAddressDestination) GetAccountOutputOk() (*TransactionAddressDestinationAccountOutput, bool)`

GetAccountOutputOk returns a tuple with the AccountOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOutput

`func (o *TransactionAddressDestination) SetAccountOutput(v TransactionAddressDestinationAccountOutput)`

SetAccountOutput sets AccountOutput field to given value.

### HasAccountOutput

`func (o *TransactionAddressDestination) HasAccountOutput() bool`

HasAccountOutput returns a boolean if a field has been set.

### GetUtxoOutputs

`func (o *TransactionAddressDestination) GetUtxoOutputs() TransactionAddressDestinationUtxoOutputs`

GetUtxoOutputs returns the UtxoOutputs field if non-nil, zero value otherwise.

### GetUtxoOutputsOk

`func (o *TransactionAddressDestination) GetUtxoOutputsOk() (*TransactionAddressDestinationUtxoOutputs, bool)`

GetUtxoOutputsOk returns a tuple with the UtxoOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoOutputs

`func (o *TransactionAddressDestination) SetUtxoOutputs(v TransactionAddressDestinationUtxoOutputs)`

SetUtxoOutputs sets UtxoOutputs field to given value.

### HasUtxoOutputs

`func (o *TransactionAddressDestination) HasUtxoOutputs() bool`

HasUtxoOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



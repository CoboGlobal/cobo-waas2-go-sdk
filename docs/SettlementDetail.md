# SettlementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The fiat currency for the settlement. | [optional] 
**TokenId** | Pointer to **string** | The ID of the cryptocurrency settled. | [optional] 
**ChainId** | Pointer to **string** | The ID of the blockchain network on which the settlement occurred. | [optional] 
**Amount** | Pointer to **string** | The settlement amount.  - If &#x60;token_id&#x60; is specified, this represents the settlement amount in the specified cryptocurrency.  - If &#x60;token_id&#x60; is not specified, this represents the settlement amount in the specified fiat currency.  | [optional] 
**Status** | Pointer to [**SettleStatus**](SettleStatus.md) |  | [optional] 
**BankAccount** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this settlement request. Each transaction represents a separate blockchain operation related to the settlement process. | [optional] 

## Methods

### NewSettlementDetail

`func NewSettlementDetail() *SettlementDetail`

NewSettlementDetail instantiates a new SettlementDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementDetailWithDefaults

`func NewSettlementDetailWithDefaults() *SettlementDetail`

NewSettlementDetailWithDefaults instantiates a new SettlementDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *SettlementDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SettlementDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SettlementDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SettlementDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTokenId

`func (o *SettlementDetail) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SettlementDetail) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SettlementDetail) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *SettlementDetail) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetChainId

`func (o *SettlementDetail) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SettlementDetail) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SettlementDetail) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *SettlementDetail) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetAmount

`func (o *SettlementDetail) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SettlementDetail) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SettlementDetail) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SettlementDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStatus

`func (o *SettlementDetail) GetStatus() SettleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SettlementDetail) GetStatusOk() (*SettleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SettlementDetail) SetStatus(v SettleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SettlementDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBankAccount

`func (o *SettlementDetail) GetBankAccount() BankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *SettlementDetail) GetBankAccountOk() (*BankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *SettlementDetail) SetBankAccount(v BankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *SettlementDetail) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetTransactions

`func (o *SettlementDetail) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *SettlementDetail) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *SettlementDetail) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *SettlementDetail) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



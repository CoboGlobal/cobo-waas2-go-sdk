# SettlementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The fiat currency for the settlement. | [optional] 
**TokenId** | Pointer to **string** | The ID of the cryptocurrency settled. | [optional] 
**ChainId** | Pointer to **string** | The ID of the blockchain network on which the settlement occurred. | [optional] 
**MerchantId** | Pointer to **string** | The ID of the merchant associated with this settlement. | [optional] 
**Amount** | Pointer to **string** | The settlement amount. - If &#x60;payout_channel&#x60; is set to &#x60;Crypto&#x60;, this represents the settlement amount in the specified cryptocurrency. - If &#x60;payout_channel&#x60; is set to &#x60;OffRamp&#x60;, this represents the settlement amount in the specified fiat currency.  | [optional] 
**SettledAmount** | Pointer to **string** | The settled amount of this settlement detail.  - If &#x60;payout_channel&#x60; is set to &#x60;Crypto&#x60;, this represents the actual settled amount in the specified cryptocurrency.  - If &#x60;payout_channel&#x60; is set to &#x60;OffRamp&#x60;, this represents the actual settled amount in the specified fiat currency.  | [optional] 
**Status** | Pointer to [**SettleStatus**](SettleStatus.md) |  | [optional] 
**BankAccount** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this settlement request. Each transaction represents a separate blockchain operation related to the settlement process. | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The creation time of the settlement, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The last update time of the settlement, represented as a UNIX timestamp in seconds. | [optional] 
**CryptoAddressId** | Pointer to **string** | The ID of the crypto address used for crypto withdrawal. | [optional] 
**PayoutChannel** | Pointer to [**PayoutChannel**](PayoutChannel.md) |  | [optional] 
**AcquiringType** | Pointer to [**AcquiringType**](AcquiringType.md) |  | [optional] 

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

### GetMerchantId

`func (o *SettlementDetail) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *SettlementDetail) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *SettlementDetail) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *SettlementDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

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

### GetSettledAmount

`func (o *SettlementDetail) GetSettledAmount() string`

GetSettledAmount returns the SettledAmount field if non-nil, zero value otherwise.

### GetSettledAmountOk

`func (o *SettlementDetail) GetSettledAmountOk() (*string, bool)`

GetSettledAmountOk returns a tuple with the SettledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAmount

`func (o *SettlementDetail) SetSettledAmount(v string)`

SetSettledAmount sets SettledAmount field to given value.

### HasSettledAmount

`func (o *SettlementDetail) HasSettledAmount() bool`

HasSettledAmount returns a boolean if a field has been set.

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

### GetCreatedTimestamp

`func (o *SettlementDetail) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SettlementDetail) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SettlementDetail) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *SettlementDetail) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SettlementDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SettlementDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SettlementDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *SettlementDetail) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetCryptoAddressId

`func (o *SettlementDetail) GetCryptoAddressId() string`

GetCryptoAddressId returns the CryptoAddressId field if non-nil, zero value otherwise.

### GetCryptoAddressIdOk

`func (o *SettlementDetail) GetCryptoAddressIdOk() (*string, bool)`

GetCryptoAddressIdOk returns a tuple with the CryptoAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddressId

`func (o *SettlementDetail) SetCryptoAddressId(v string)`

SetCryptoAddressId sets CryptoAddressId field to given value.

### HasCryptoAddressId

`func (o *SettlementDetail) HasCryptoAddressId() bool`

HasCryptoAddressId returns a boolean if a field has been set.

### GetPayoutChannel

`func (o *SettlementDetail) GetPayoutChannel() PayoutChannel`

GetPayoutChannel returns the PayoutChannel field if non-nil, zero value otherwise.

### GetPayoutChannelOk

`func (o *SettlementDetail) GetPayoutChannelOk() (*PayoutChannel, bool)`

GetPayoutChannelOk returns a tuple with the PayoutChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutChannel

`func (o *SettlementDetail) SetPayoutChannel(v PayoutChannel)`

SetPayoutChannel sets PayoutChannel field to given value.

### HasPayoutChannel

`func (o *SettlementDetail) HasPayoutChannel() bool`

HasPayoutChannel returns a boolean if a field has been set.

### GetAcquiringType

`func (o *SettlementDetail) GetAcquiringType() AcquiringType`

GetAcquiringType returns the AcquiringType field if non-nil, zero value otherwise.

### GetAcquiringTypeOk

`func (o *SettlementDetail) GetAcquiringTypeOk() (*AcquiringType, bool)`

GetAcquiringTypeOk returns a tuple with the AcquiringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiringType

`func (o *SettlementDetail) SetAcquiringType(v AcquiringType)`

SetAcquiringType sets AcquiringType field to given value.

### HasAcquiringType

`func (o *SettlementDetail) HasAcquiringType() bool`

HasAcquiringType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



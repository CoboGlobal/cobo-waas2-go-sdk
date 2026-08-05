# PaymentPayoutRecipientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The recipient&#39;s cryptocurrency address. Use an address that you have registered as a destination and that has been approved for payouts. Call [List destination entries](https://www.cobo.com/payments/en/api-references/payment/list-destination-entries) with &#x60;entry_type&#x60; set to &#x60;Address&#x60; to retrieve your registered wallet addresses and each address&#39;s &#x60;chain_id&#x60;. A destination entry returns the &#x60;address&#x60; and its &#x60;chain_id&#x60;, not a &#x60;token_id&#x60; -- confirm that the network indicated by &#x60;chain_id&#x60; matches the network encoded in the &#x60;token_id&#x60; you use, since an address on the wrong network for that token cannot complete the transfer. | [optional] 
**TokenId** | Pointer to **string** | The cryptocurrency ID of the token to send (for example, &#x60;ETH_USDT&#x60;, &#x60;TRON_USDT&#x60;). When this token ID and the &#x60;token_id&#x60; in &#x60;payout_params&#x60; represent the same token on different chains, Cobo automatically executes a cross-chain transfer. You can retrieve the full list of supported token IDs by calling [List supported tokens](https://www.cobo.com/payments/en/api-references/payment/list-supported-tokens). | [optional] 
**Currency** | Pointer to **string** | The fiat currency of the bank account selected in &#x60;bank_account_id&#x60; -- the two must match. This endpoint currently accepts only &#x60;USD&#x60;. | [optional] 
**BankAccountId** | Pointer to **string** | The ID of the bank account to which the payout will be sent. This field is required only when the payout channel is &#x60;OffRamp&#x60;. You can retrieve the bank account ID by calling [List destination entries](https://www.cobo.com/payments/en/api-references/payment/list-destination-entries). | [optional] 
**TransferViaVa** | Pointer to **bool** | For OffRamp payout, whether the payout is transferred to a registered bank account via a virtual account (VA) or directly. - &#x60;true&#x60;: The payout is transferred to a registered bank account via a VA (virtual account). - &#x60;false&#x60;: The payout is transferred directly to a registered bank account.  | [optional] 

## Methods

### NewPaymentPayoutRecipientInfo

`func NewPaymentPayoutRecipientInfo() *PaymentPayoutRecipientInfo`

NewPaymentPayoutRecipientInfo instantiates a new PaymentPayoutRecipientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutRecipientInfoWithDefaults

`func NewPaymentPayoutRecipientInfoWithDefaults() *PaymentPayoutRecipientInfo`

NewPaymentPayoutRecipientInfoWithDefaults instantiates a new PaymentPayoutRecipientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PaymentPayoutRecipientInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentPayoutRecipientInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentPayoutRecipientInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaymentPayoutRecipientInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *PaymentPayoutRecipientInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentPayoutRecipientInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentPayoutRecipientInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *PaymentPayoutRecipientInfo) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentPayoutRecipientInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentPayoutRecipientInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentPayoutRecipientInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentPayoutRecipientInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBankAccountId

`func (o *PaymentPayoutRecipientInfo) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *PaymentPayoutRecipientInfo) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *PaymentPayoutRecipientInfo) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *PaymentPayoutRecipientInfo) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetTransferViaVa

`func (o *PaymentPayoutRecipientInfo) GetTransferViaVa() bool`

GetTransferViaVa returns the TransferViaVa field if non-nil, zero value otherwise.

### GetTransferViaVaOk

`func (o *PaymentPayoutRecipientInfo) GetTransferViaVaOk() (*bool, bool)`

GetTransferViaVaOk returns a tuple with the TransferViaVa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferViaVa

`func (o *PaymentPayoutRecipientInfo) SetTransferViaVa(v bool)`

SetTransferViaVa sets TransferViaVa field to given value.

### HasTransferViaVa

`func (o *PaymentPayoutRecipientInfo) HasTransferViaVa() bool`

HasTransferViaVa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



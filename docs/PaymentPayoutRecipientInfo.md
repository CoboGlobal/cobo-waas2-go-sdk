# PaymentPayoutRecipientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The recipient&#39;s wallet address where the payout will be sent. | [optional] 
**TokenId** | Pointer to **string** | The token ID for the cryptocurrency to be sent to the recipient.  If &#x60;recipient_info.token_id&#x60; is on a different chain than &#x60;payout_param.token_id&#x60;, the token will be automatically bridged to the chain specified in &#x60;recipient_info.token_id&#x60;.  | [optional] 
**Currency** | Pointer to **string** | The fiat currency of the bank account to which the payout will be sent. | [optional] 
**BankAccountId** | Pointer to **string** | The ID of the bank account to which the payout will be sent. You can retrieve the bank account ID by calling [List destination entries](https://www.cobo.com/payments/en/api-references/payment/list-destination-entries). | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentBankAccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountId** | **string** | The destination bank account ID. | 
**Currency** | **string** | The fiat currency of the bank account. | 
**TotalBalance** | **string** | The total balance of the bank account. | 
**AvailableBalance** | **string** | The available balance of the bank account. | 
**LockedBalance** | **string** | The locked balance of the bank account. | 

## Methods

### NewPaymentBankAccountBalance

`func NewPaymentBankAccountBalance(bankAccountId string, currency string, totalBalance string, availableBalance string, lockedBalance string, ) *PaymentBankAccountBalance`

NewPaymentBankAccountBalance instantiates a new PaymentBankAccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBankAccountBalanceWithDefaults

`func NewPaymentBankAccountBalanceWithDefaults() *PaymentBankAccountBalance`

NewPaymentBankAccountBalanceWithDefaults instantiates a new PaymentBankAccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountId

`func (o *PaymentBankAccountBalance) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *PaymentBankAccountBalance) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *PaymentBankAccountBalance) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.


### GetCurrency

`func (o *PaymentBankAccountBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentBankAccountBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentBankAccountBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotalBalance

`func (o *PaymentBankAccountBalance) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *PaymentBankAccountBalance) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *PaymentBankAccountBalance) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.


### GetAvailableBalance

`func (o *PaymentBankAccountBalance) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *PaymentBankAccountBalance) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *PaymentBankAccountBalance) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetLockedBalance

`func (o *PaymentBankAccountBalance) GetLockedBalance() string`

GetLockedBalance returns the LockedBalance field if non-nil, zero value otherwise.

### GetLockedBalanceOk

`func (o *PaymentBankAccountBalance) GetLockedBalanceOk() (*string, bool)`

GetLockedBalanceOk returns a tuple with the LockedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBalance

`func (o *PaymentBankAccountBalance) SetLockedBalance(v string)`

SetLockedBalance sets LockedBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



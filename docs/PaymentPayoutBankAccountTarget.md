# PaymentPayoutBankAccountTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | **string** | The type of the payout target. | 
**BankAccountId** | **string** | The recipient bank account ID. | 
**AccountAlias** | Pointer to **string** | The alias of the recipient bank account. | [optional] 
**BankName** | Pointer to **string** | The recipient bank name. | [optional] 
**Currency** | Pointer to **string** | The currency of the recipient bank account. | [optional] 

## Methods

### NewPaymentPayoutBankAccountTarget

`func NewPaymentPayoutBankAccountTarget(targetType string, bankAccountId string, ) *PaymentPayoutBankAccountTarget`

NewPaymentPayoutBankAccountTarget instantiates a new PaymentPayoutBankAccountTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutBankAccountTargetWithDefaults

`func NewPaymentPayoutBankAccountTargetWithDefaults() *PaymentPayoutBankAccountTarget`

NewPaymentPayoutBankAccountTargetWithDefaults instantiates a new PaymentPayoutBankAccountTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetType

`func (o *PaymentPayoutBankAccountTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *PaymentPayoutBankAccountTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *PaymentPayoutBankAccountTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetBankAccountId

`func (o *PaymentPayoutBankAccountTarget) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *PaymentPayoutBankAccountTarget) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *PaymentPayoutBankAccountTarget) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.


### GetAccountAlias

`func (o *PaymentPayoutBankAccountTarget) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *PaymentPayoutBankAccountTarget) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *PaymentPayoutBankAccountTarget) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *PaymentPayoutBankAccountTarget) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetBankName

`func (o *PaymentPayoutBankAccountTarget) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentPayoutBankAccountTarget) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentPayoutBankAccountTarget) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PaymentPayoutBankAccountTarget) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentPayoutBankAccountTarget) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentPayoutBankAccountTarget) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentPayoutBankAccountTarget) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentPayoutBankAccountTarget) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentPayoutRequestedTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | **string** | The type of the payout target. | 
**Address** | **string** | The recipient address. | 
**ChainId** | **string** | The chain ID of the recipient address. | 
**BankAccountId** | **string** | The recipient bank account ID. | 
**AccountAlias** | Pointer to **string** | The alias of the recipient bank account. | [optional] 
**BankName** | Pointer to **string** | The recipient bank name. | [optional] 
**Currency** | Pointer to **string** | The currency of the recipient bank account. | [optional] 

## Methods

### NewPaymentPayoutRequestedTarget

`func NewPaymentPayoutRequestedTarget(targetType string, address string, chainId string, bankAccountId string, ) *PaymentPayoutRequestedTarget`

NewPaymentPayoutRequestedTarget instantiates a new PaymentPayoutRequestedTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutRequestedTargetWithDefaults

`func NewPaymentPayoutRequestedTargetWithDefaults() *PaymentPayoutRequestedTarget`

NewPaymentPayoutRequestedTargetWithDefaults instantiates a new PaymentPayoutRequestedTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetType

`func (o *PaymentPayoutRequestedTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *PaymentPayoutRequestedTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *PaymentPayoutRequestedTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetAddress

`func (o *PaymentPayoutRequestedTarget) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentPayoutRequestedTarget) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentPayoutRequestedTarget) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *PaymentPayoutRequestedTarget) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PaymentPayoutRequestedTarget) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PaymentPayoutRequestedTarget) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetBankAccountId

`func (o *PaymentPayoutRequestedTarget) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *PaymentPayoutRequestedTarget) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *PaymentPayoutRequestedTarget) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.


### GetAccountAlias

`func (o *PaymentPayoutRequestedTarget) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *PaymentPayoutRequestedTarget) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *PaymentPayoutRequestedTarget) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *PaymentPayoutRequestedTarget) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetBankName

`func (o *PaymentPayoutRequestedTarget) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentPayoutRequestedTarget) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentPayoutRequestedTarget) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PaymentPayoutRequestedTarget) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentPayoutRequestedTarget) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentPayoutRequestedTarget) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentPayoutRequestedTarget) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentPayoutRequestedTarget) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



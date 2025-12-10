# UpdateDestinationBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | **string** | The alias of the bank account. | 
**AccountNumber** | **string** | The bank account number. | 
**SwiftCode** | **string** | The SWIFT or BIC code of the bank. | 
**Currency** | **string** | The currency of the bank account. | 
**BeneficiaryName** | **string** | The name of the account holder. | 
**BeneficiaryAddress** | **string** | The address of the account holder. | 
**BankName** | **string** | The name of the bank. | 
**BankAddress** | **string** | The address of the bank. | 
**IbanCode** | Pointer to **string** | The IBAN code of the bank account. | [optional] 
**FurtherCredit** | Pointer to **string** | The further credit of the bank account. | [optional] 
**IntermediaryBankInfo** | Pointer to [**IntermediaryBankInfo**](IntermediaryBankInfo.md) |  | [optional] 

## Methods

### NewUpdateDestinationBankAccount

`func NewUpdateDestinationBankAccount(accountAlias string, accountNumber string, swiftCode string, currency string, beneficiaryName string, beneficiaryAddress string, bankName string, bankAddress string, ) *UpdateDestinationBankAccount`

NewUpdateDestinationBankAccount instantiates a new UpdateDestinationBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDestinationBankAccountWithDefaults

`func NewUpdateDestinationBankAccountWithDefaults() *UpdateDestinationBankAccount`

NewUpdateDestinationBankAccountWithDefaults instantiates a new UpdateDestinationBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *UpdateDestinationBankAccount) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *UpdateDestinationBankAccount) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *UpdateDestinationBankAccount) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.


### GetAccountNumber

`func (o *UpdateDestinationBankAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UpdateDestinationBankAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UpdateDestinationBankAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetSwiftCode

`func (o *UpdateDestinationBankAccount) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *UpdateDestinationBankAccount) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *UpdateDestinationBankAccount) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.


### GetCurrency

`func (o *UpdateDestinationBankAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateDestinationBankAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateDestinationBankAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBeneficiaryName

`func (o *UpdateDestinationBankAccount) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *UpdateDestinationBankAccount) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *UpdateDestinationBankAccount) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.


### GetBeneficiaryAddress

`func (o *UpdateDestinationBankAccount) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *UpdateDestinationBankAccount) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *UpdateDestinationBankAccount) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.


### GetBankName

`func (o *UpdateDestinationBankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *UpdateDestinationBankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *UpdateDestinationBankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankAddress

`func (o *UpdateDestinationBankAccount) GetBankAddress() string`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *UpdateDestinationBankAccount) GetBankAddressOk() (*string, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *UpdateDestinationBankAccount) SetBankAddress(v string)`

SetBankAddress sets BankAddress field to given value.


### GetIbanCode

`func (o *UpdateDestinationBankAccount) GetIbanCode() string`

GetIbanCode returns the IbanCode field if non-nil, zero value otherwise.

### GetIbanCodeOk

`func (o *UpdateDestinationBankAccount) GetIbanCodeOk() (*string, bool)`

GetIbanCodeOk returns a tuple with the IbanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanCode

`func (o *UpdateDestinationBankAccount) SetIbanCode(v string)`

SetIbanCode sets IbanCode field to given value.

### HasIbanCode

`func (o *UpdateDestinationBankAccount) HasIbanCode() bool`

HasIbanCode returns a boolean if a field has been set.

### GetFurtherCredit

`func (o *UpdateDestinationBankAccount) GetFurtherCredit() string`

GetFurtherCredit returns the FurtherCredit field if non-nil, zero value otherwise.

### GetFurtherCreditOk

`func (o *UpdateDestinationBankAccount) GetFurtherCreditOk() (*string, bool)`

GetFurtherCreditOk returns a tuple with the FurtherCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurtherCredit

`func (o *UpdateDestinationBankAccount) SetFurtherCredit(v string)`

SetFurtherCredit sets FurtherCredit field to given value.

### HasFurtherCredit

`func (o *UpdateDestinationBankAccount) HasFurtherCredit() bool`

HasFurtherCredit returns a boolean if a field has been set.

### GetIntermediaryBankInfo

`func (o *UpdateDestinationBankAccount) GetIntermediaryBankInfo() IntermediaryBankInfo`

GetIntermediaryBankInfo returns the IntermediaryBankInfo field if non-nil, zero value otherwise.

### GetIntermediaryBankInfoOk

`func (o *UpdateDestinationBankAccount) GetIntermediaryBankInfoOk() (*IntermediaryBankInfo, bool)`

GetIntermediaryBankInfoOk returns a tuple with the IntermediaryBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBankInfo

`func (o *UpdateDestinationBankAccount) SetIntermediaryBankInfo(v IntermediaryBankInfo)`

SetIntermediaryBankInfo sets IntermediaryBankInfo field to given value.

### HasIntermediaryBankInfo

`func (o *UpdateDestinationBankAccount) HasIntermediaryBankInfo() bool`

HasIntermediaryBankInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



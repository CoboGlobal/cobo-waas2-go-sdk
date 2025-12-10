# CreateDestinationBankAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | The destination ID. | 
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

### NewCreateDestinationBankAccountRequest

`func NewCreateDestinationBankAccountRequest(destinationId string, accountAlias string, accountNumber string, swiftCode string, currency string, beneficiaryName string, beneficiaryAddress string, bankName string, bankAddress string, ) *CreateDestinationBankAccountRequest`

NewCreateDestinationBankAccountRequest instantiates a new CreateDestinationBankAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationBankAccountRequestWithDefaults

`func NewCreateDestinationBankAccountRequestWithDefaults() *CreateDestinationBankAccountRequest`

NewCreateDestinationBankAccountRequestWithDefaults instantiates a new CreateDestinationBankAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *CreateDestinationBankAccountRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CreateDestinationBankAccountRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CreateDestinationBankAccountRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetAccountAlias

`func (o *CreateDestinationBankAccountRequest) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *CreateDestinationBankAccountRequest) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *CreateDestinationBankAccountRequest) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.


### GetAccountNumber

`func (o *CreateDestinationBankAccountRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateDestinationBankAccountRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateDestinationBankAccountRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetSwiftCode

`func (o *CreateDestinationBankAccountRequest) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *CreateDestinationBankAccountRequest) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *CreateDestinationBankAccountRequest) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.


### GetCurrency

`func (o *CreateDestinationBankAccountRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateDestinationBankAccountRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateDestinationBankAccountRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBeneficiaryName

`func (o *CreateDestinationBankAccountRequest) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *CreateDestinationBankAccountRequest) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *CreateDestinationBankAccountRequest) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.


### GetBeneficiaryAddress

`func (o *CreateDestinationBankAccountRequest) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *CreateDestinationBankAccountRequest) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *CreateDestinationBankAccountRequest) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.


### GetBankName

`func (o *CreateDestinationBankAccountRequest) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *CreateDestinationBankAccountRequest) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *CreateDestinationBankAccountRequest) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankAddress

`func (o *CreateDestinationBankAccountRequest) GetBankAddress() string`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *CreateDestinationBankAccountRequest) GetBankAddressOk() (*string, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *CreateDestinationBankAccountRequest) SetBankAddress(v string)`

SetBankAddress sets BankAddress field to given value.


### GetIbanCode

`func (o *CreateDestinationBankAccountRequest) GetIbanCode() string`

GetIbanCode returns the IbanCode field if non-nil, zero value otherwise.

### GetIbanCodeOk

`func (o *CreateDestinationBankAccountRequest) GetIbanCodeOk() (*string, bool)`

GetIbanCodeOk returns a tuple with the IbanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanCode

`func (o *CreateDestinationBankAccountRequest) SetIbanCode(v string)`

SetIbanCode sets IbanCode field to given value.

### HasIbanCode

`func (o *CreateDestinationBankAccountRequest) HasIbanCode() bool`

HasIbanCode returns a boolean if a field has been set.

### GetFurtherCredit

`func (o *CreateDestinationBankAccountRequest) GetFurtherCredit() string`

GetFurtherCredit returns the FurtherCredit field if non-nil, zero value otherwise.

### GetFurtherCreditOk

`func (o *CreateDestinationBankAccountRequest) GetFurtherCreditOk() (*string, bool)`

GetFurtherCreditOk returns a tuple with the FurtherCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurtherCredit

`func (o *CreateDestinationBankAccountRequest) SetFurtherCredit(v string)`

SetFurtherCredit sets FurtherCredit field to given value.

### HasFurtherCredit

`func (o *CreateDestinationBankAccountRequest) HasFurtherCredit() bool`

HasFurtherCredit returns a boolean if a field has been set.

### GetIntermediaryBankInfo

`func (o *CreateDestinationBankAccountRequest) GetIntermediaryBankInfo() IntermediaryBankInfo`

GetIntermediaryBankInfo returns the IntermediaryBankInfo field if non-nil, zero value otherwise.

### GetIntermediaryBankInfoOk

`func (o *CreateDestinationBankAccountRequest) GetIntermediaryBankInfoOk() (*IntermediaryBankInfo, bool)`

GetIntermediaryBankInfoOk returns a tuple with the IntermediaryBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBankInfo

`func (o *CreateDestinationBankAccountRequest) SetIntermediaryBankInfo(v IntermediaryBankInfo)`

SetIntermediaryBankInfo sets IntermediaryBankInfo field to given value.

### HasIntermediaryBankInfo

`func (o *CreateDestinationBankAccountRequest) HasIntermediaryBankInfo() bool`

HasIntermediaryBankInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



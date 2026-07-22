# CreateDestinationBankAccount

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
**Country** | Pointer to **string** | The country, in ISO 3166-1 alpha-3 format. | [optional] 
**City** | Pointer to **string** | Beneficiary&#39;s city. | [optional] 
**PaymentMethod** | Pointer to [**BankAccountPaymentMethod**](BankAccountPaymentMethod.md) |  | [optional] 
**HolderType** | Pointer to [**BankAccountHolderType**](BankAccountHolderType.md) |  | [optional] 
**BeneficiaryProvince** | Pointer to **string** | The province or state of the beneficiary. Required when &#x60;payment_method&#x60; is &#x60;Swift&#x60;. Cannot be a pure number or contain Chinese characters.  | [optional] 
**BeneficiaryPostCode** | Pointer to **string** | The postal code of the beneficiary. Required when &#x60;payment_method&#x60; is &#x60;Swift&#x60;.  | [optional] 
**BankAccountName** | Pointer to **string** | The bank account name. Cannot contain Chinese characters.  | [optional] 
**BankBranchCode** | Pointer to **string** | The branch code. Required when &#x60;payment_method&#x60; is &#x60;Local&#x60; (HK only).  | [optional] 
**BankCountry** | Pointer to **string** | The country, in ISO 3166-1 alpha-3 format. | [optional] 
**BankProvince** | Pointer to **string** | The province or state of the bank. Cannot be a pure number or contain Chinese characters.  | [optional] 
**ContractFileId** | Pointer to **string** | The file ID of the contract document (e.g., cooperation agreement) that proves the business relationship between you and the beneficiary, which you can retrieve by calling [Upload file](https://www.cobo.com/developers/v2/api-references/payment/upload-file).  | [optional] 

## Methods

### NewCreateDestinationBankAccount

`func NewCreateDestinationBankAccount(accountAlias string, accountNumber string, swiftCode string, currency string, beneficiaryName string, beneficiaryAddress string, bankName string, bankAddress string, ) *CreateDestinationBankAccount`

NewCreateDestinationBankAccount instantiates a new CreateDestinationBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationBankAccountWithDefaults

`func NewCreateDestinationBankAccountWithDefaults() *CreateDestinationBankAccount`

NewCreateDestinationBankAccountWithDefaults instantiates a new CreateDestinationBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *CreateDestinationBankAccount) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *CreateDestinationBankAccount) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *CreateDestinationBankAccount) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.


### GetAccountNumber

`func (o *CreateDestinationBankAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateDestinationBankAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateDestinationBankAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetSwiftCode

`func (o *CreateDestinationBankAccount) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *CreateDestinationBankAccount) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *CreateDestinationBankAccount) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.


### GetCurrency

`func (o *CreateDestinationBankAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateDestinationBankAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateDestinationBankAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBeneficiaryName

`func (o *CreateDestinationBankAccount) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *CreateDestinationBankAccount) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *CreateDestinationBankAccount) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.


### GetBeneficiaryAddress

`func (o *CreateDestinationBankAccount) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *CreateDestinationBankAccount) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *CreateDestinationBankAccount) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.


### GetBankName

`func (o *CreateDestinationBankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *CreateDestinationBankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *CreateDestinationBankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankAddress

`func (o *CreateDestinationBankAccount) GetBankAddress() string`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *CreateDestinationBankAccount) GetBankAddressOk() (*string, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *CreateDestinationBankAccount) SetBankAddress(v string)`

SetBankAddress sets BankAddress field to given value.


### GetIbanCode

`func (o *CreateDestinationBankAccount) GetIbanCode() string`

GetIbanCode returns the IbanCode field if non-nil, zero value otherwise.

### GetIbanCodeOk

`func (o *CreateDestinationBankAccount) GetIbanCodeOk() (*string, bool)`

GetIbanCodeOk returns a tuple with the IbanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanCode

`func (o *CreateDestinationBankAccount) SetIbanCode(v string)`

SetIbanCode sets IbanCode field to given value.

### HasIbanCode

`func (o *CreateDestinationBankAccount) HasIbanCode() bool`

HasIbanCode returns a boolean if a field has been set.

### GetFurtherCredit

`func (o *CreateDestinationBankAccount) GetFurtherCredit() string`

GetFurtherCredit returns the FurtherCredit field if non-nil, zero value otherwise.

### GetFurtherCreditOk

`func (o *CreateDestinationBankAccount) GetFurtherCreditOk() (*string, bool)`

GetFurtherCreditOk returns a tuple with the FurtherCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurtherCredit

`func (o *CreateDestinationBankAccount) SetFurtherCredit(v string)`

SetFurtherCredit sets FurtherCredit field to given value.

### HasFurtherCredit

`func (o *CreateDestinationBankAccount) HasFurtherCredit() bool`

HasFurtherCredit returns a boolean if a field has been set.

### GetIntermediaryBankInfo

`func (o *CreateDestinationBankAccount) GetIntermediaryBankInfo() IntermediaryBankInfo`

GetIntermediaryBankInfo returns the IntermediaryBankInfo field if non-nil, zero value otherwise.

### GetIntermediaryBankInfoOk

`func (o *CreateDestinationBankAccount) GetIntermediaryBankInfoOk() (*IntermediaryBankInfo, bool)`

GetIntermediaryBankInfoOk returns a tuple with the IntermediaryBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBankInfo

`func (o *CreateDestinationBankAccount) SetIntermediaryBankInfo(v IntermediaryBankInfo)`

SetIntermediaryBankInfo sets IntermediaryBankInfo field to given value.

### HasIntermediaryBankInfo

`func (o *CreateDestinationBankAccount) HasIntermediaryBankInfo() bool`

HasIntermediaryBankInfo returns a boolean if a field has been set.

### GetCountry

`func (o *CreateDestinationBankAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateDestinationBankAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateDestinationBankAccount) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateDestinationBankAccount) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *CreateDestinationBankAccount) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateDestinationBankAccount) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreateDestinationBankAccount) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CreateDestinationBankAccount) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CreateDestinationBankAccount) GetPaymentMethod() BankAccountPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreateDestinationBankAccount) GetPaymentMethodOk() (*BankAccountPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreateDestinationBankAccount) SetPaymentMethod(v BankAccountPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CreateDestinationBankAccount) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetHolderType

`func (o *CreateDestinationBankAccount) GetHolderType() BankAccountHolderType`

GetHolderType returns the HolderType field if non-nil, zero value otherwise.

### GetHolderTypeOk

`func (o *CreateDestinationBankAccount) GetHolderTypeOk() (*BankAccountHolderType, bool)`

GetHolderTypeOk returns a tuple with the HolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderType

`func (o *CreateDestinationBankAccount) SetHolderType(v BankAccountHolderType)`

SetHolderType sets HolderType field to given value.

### HasHolderType

`func (o *CreateDestinationBankAccount) HasHolderType() bool`

HasHolderType returns a boolean if a field has been set.

### GetBeneficiaryProvince

`func (o *CreateDestinationBankAccount) GetBeneficiaryProvince() string`

GetBeneficiaryProvince returns the BeneficiaryProvince field if non-nil, zero value otherwise.

### GetBeneficiaryProvinceOk

`func (o *CreateDestinationBankAccount) GetBeneficiaryProvinceOk() (*string, bool)`

GetBeneficiaryProvinceOk returns a tuple with the BeneficiaryProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryProvince

`func (o *CreateDestinationBankAccount) SetBeneficiaryProvince(v string)`

SetBeneficiaryProvince sets BeneficiaryProvince field to given value.

### HasBeneficiaryProvince

`func (o *CreateDestinationBankAccount) HasBeneficiaryProvince() bool`

HasBeneficiaryProvince returns a boolean if a field has been set.

### GetBeneficiaryPostCode

`func (o *CreateDestinationBankAccount) GetBeneficiaryPostCode() string`

GetBeneficiaryPostCode returns the BeneficiaryPostCode field if non-nil, zero value otherwise.

### GetBeneficiaryPostCodeOk

`func (o *CreateDestinationBankAccount) GetBeneficiaryPostCodeOk() (*string, bool)`

GetBeneficiaryPostCodeOk returns a tuple with the BeneficiaryPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPostCode

`func (o *CreateDestinationBankAccount) SetBeneficiaryPostCode(v string)`

SetBeneficiaryPostCode sets BeneficiaryPostCode field to given value.

### HasBeneficiaryPostCode

`func (o *CreateDestinationBankAccount) HasBeneficiaryPostCode() bool`

HasBeneficiaryPostCode returns a boolean if a field has been set.

### GetBankAccountName

`func (o *CreateDestinationBankAccount) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *CreateDestinationBankAccount) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *CreateDestinationBankAccount) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountName

`func (o *CreateDestinationBankAccount) HasBankAccountName() bool`

HasBankAccountName returns a boolean if a field has been set.

### GetBankBranchCode

`func (o *CreateDestinationBankAccount) GetBankBranchCode() string`

GetBankBranchCode returns the BankBranchCode field if non-nil, zero value otherwise.

### GetBankBranchCodeOk

`func (o *CreateDestinationBankAccount) GetBankBranchCodeOk() (*string, bool)`

GetBankBranchCodeOk returns a tuple with the BankBranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranchCode

`func (o *CreateDestinationBankAccount) SetBankBranchCode(v string)`

SetBankBranchCode sets BankBranchCode field to given value.

### HasBankBranchCode

`func (o *CreateDestinationBankAccount) HasBankBranchCode() bool`

HasBankBranchCode returns a boolean if a field has been set.

### GetBankCountry

`func (o *CreateDestinationBankAccount) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *CreateDestinationBankAccount) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *CreateDestinationBankAccount) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.

### HasBankCountry

`func (o *CreateDestinationBankAccount) HasBankCountry() bool`

HasBankCountry returns a boolean if a field has been set.

### GetBankProvince

`func (o *CreateDestinationBankAccount) GetBankProvince() string`

GetBankProvince returns the BankProvince field if non-nil, zero value otherwise.

### GetBankProvinceOk

`func (o *CreateDestinationBankAccount) GetBankProvinceOk() (*string, bool)`

GetBankProvinceOk returns a tuple with the BankProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProvince

`func (o *CreateDestinationBankAccount) SetBankProvince(v string)`

SetBankProvince sets BankProvince field to given value.

### HasBankProvince

`func (o *CreateDestinationBankAccount) HasBankProvince() bool`

HasBankProvince returns a boolean if a field has been set.

### GetContractFileId

`func (o *CreateDestinationBankAccount) GetContractFileId() string`

GetContractFileId returns the ContractFileId field if non-nil, zero value otherwise.

### GetContractFileIdOk

`func (o *CreateDestinationBankAccount) GetContractFileIdOk() (*string, bool)`

GetContractFileIdOk returns a tuple with the ContractFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractFileId

`func (o *CreateDestinationBankAccount) SetContractFileId(v string)`

SetContractFileId sets ContractFileId field to given value.

### HasContractFileId

`func (o *CreateDestinationBankAccount) HasContractFileId() bool`

HasContractFileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



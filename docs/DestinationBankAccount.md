# DestinationBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountId** | **string** | The destination bank account ID. | 
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
**BankAccountStatus** | [**BankAccountStatus**](BankAccountStatus.md) |  | 
**CreatedTimestamp** | Pointer to **int32** | The created time of the bank account, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the bank account, represented as a UNIX timestamp in seconds. | [optional] 
**Country** | Pointer to **string** | Beneficiary&#39;s country, in ISO 3166-1 alpha-3 format. | [optional] 
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

### NewDestinationBankAccount

`func NewDestinationBankAccount(bankAccountId string, accountAlias string, accountNumber string, swiftCode string, currency string, beneficiaryName string, beneficiaryAddress string, bankName string, bankAddress string, bankAccountStatus BankAccountStatus, ) *DestinationBankAccount`

NewDestinationBankAccount instantiates a new DestinationBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationBankAccountWithDefaults

`func NewDestinationBankAccountWithDefaults() *DestinationBankAccount`

NewDestinationBankAccountWithDefaults instantiates a new DestinationBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountId

`func (o *DestinationBankAccount) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *DestinationBankAccount) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *DestinationBankAccount) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.


### GetAccountAlias

`func (o *DestinationBankAccount) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *DestinationBankAccount) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *DestinationBankAccount) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.


### GetAccountNumber

`func (o *DestinationBankAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DestinationBankAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DestinationBankAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetSwiftCode

`func (o *DestinationBankAccount) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *DestinationBankAccount) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *DestinationBankAccount) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.


### GetCurrency

`func (o *DestinationBankAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DestinationBankAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DestinationBankAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBeneficiaryName

`func (o *DestinationBankAccount) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *DestinationBankAccount) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *DestinationBankAccount) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.


### GetBeneficiaryAddress

`func (o *DestinationBankAccount) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *DestinationBankAccount) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *DestinationBankAccount) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.


### GetBankName

`func (o *DestinationBankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *DestinationBankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *DestinationBankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankAddress

`func (o *DestinationBankAccount) GetBankAddress() string`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *DestinationBankAccount) GetBankAddressOk() (*string, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *DestinationBankAccount) SetBankAddress(v string)`

SetBankAddress sets BankAddress field to given value.


### GetIbanCode

`func (o *DestinationBankAccount) GetIbanCode() string`

GetIbanCode returns the IbanCode field if non-nil, zero value otherwise.

### GetIbanCodeOk

`func (o *DestinationBankAccount) GetIbanCodeOk() (*string, bool)`

GetIbanCodeOk returns a tuple with the IbanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanCode

`func (o *DestinationBankAccount) SetIbanCode(v string)`

SetIbanCode sets IbanCode field to given value.

### HasIbanCode

`func (o *DestinationBankAccount) HasIbanCode() bool`

HasIbanCode returns a boolean if a field has been set.

### GetFurtherCredit

`func (o *DestinationBankAccount) GetFurtherCredit() string`

GetFurtherCredit returns the FurtherCredit field if non-nil, zero value otherwise.

### GetFurtherCreditOk

`func (o *DestinationBankAccount) GetFurtherCreditOk() (*string, bool)`

GetFurtherCreditOk returns a tuple with the FurtherCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurtherCredit

`func (o *DestinationBankAccount) SetFurtherCredit(v string)`

SetFurtherCredit sets FurtherCredit field to given value.

### HasFurtherCredit

`func (o *DestinationBankAccount) HasFurtherCredit() bool`

HasFurtherCredit returns a boolean if a field has been set.

### GetIntermediaryBankInfo

`func (o *DestinationBankAccount) GetIntermediaryBankInfo() IntermediaryBankInfo`

GetIntermediaryBankInfo returns the IntermediaryBankInfo field if non-nil, zero value otherwise.

### GetIntermediaryBankInfoOk

`func (o *DestinationBankAccount) GetIntermediaryBankInfoOk() (*IntermediaryBankInfo, bool)`

GetIntermediaryBankInfoOk returns a tuple with the IntermediaryBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBankInfo

`func (o *DestinationBankAccount) SetIntermediaryBankInfo(v IntermediaryBankInfo)`

SetIntermediaryBankInfo sets IntermediaryBankInfo field to given value.

### HasIntermediaryBankInfo

`func (o *DestinationBankAccount) HasIntermediaryBankInfo() bool`

HasIntermediaryBankInfo returns a boolean if a field has been set.

### GetBankAccountStatus

`func (o *DestinationBankAccount) GetBankAccountStatus() BankAccountStatus`

GetBankAccountStatus returns the BankAccountStatus field if non-nil, zero value otherwise.

### GetBankAccountStatusOk

`func (o *DestinationBankAccount) GetBankAccountStatusOk() (*BankAccountStatus, bool)`

GetBankAccountStatusOk returns a tuple with the BankAccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountStatus

`func (o *DestinationBankAccount) SetBankAccountStatus(v BankAccountStatus)`

SetBankAccountStatus sets BankAccountStatus field to given value.


### GetCreatedTimestamp

`func (o *DestinationBankAccount) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DestinationBankAccount) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DestinationBankAccount) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *DestinationBankAccount) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DestinationBankAccount) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DestinationBankAccount) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DestinationBankAccount) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *DestinationBankAccount) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetCountry

`func (o *DestinationBankAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DestinationBankAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DestinationBankAccount) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DestinationBankAccount) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *DestinationBankAccount) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DestinationBankAccount) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DestinationBankAccount) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DestinationBankAccount) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *DestinationBankAccount) GetPaymentMethod() BankAccountPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *DestinationBankAccount) GetPaymentMethodOk() (*BankAccountPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *DestinationBankAccount) SetPaymentMethod(v BankAccountPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *DestinationBankAccount) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetHolderType

`func (o *DestinationBankAccount) GetHolderType() BankAccountHolderType`

GetHolderType returns the HolderType field if non-nil, zero value otherwise.

### GetHolderTypeOk

`func (o *DestinationBankAccount) GetHolderTypeOk() (*BankAccountHolderType, bool)`

GetHolderTypeOk returns a tuple with the HolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderType

`func (o *DestinationBankAccount) SetHolderType(v BankAccountHolderType)`

SetHolderType sets HolderType field to given value.

### HasHolderType

`func (o *DestinationBankAccount) HasHolderType() bool`

HasHolderType returns a boolean if a field has been set.

### GetBeneficiaryProvince

`func (o *DestinationBankAccount) GetBeneficiaryProvince() string`

GetBeneficiaryProvince returns the BeneficiaryProvince field if non-nil, zero value otherwise.

### GetBeneficiaryProvinceOk

`func (o *DestinationBankAccount) GetBeneficiaryProvinceOk() (*string, bool)`

GetBeneficiaryProvinceOk returns a tuple with the BeneficiaryProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryProvince

`func (o *DestinationBankAccount) SetBeneficiaryProvince(v string)`

SetBeneficiaryProvince sets BeneficiaryProvince field to given value.

### HasBeneficiaryProvince

`func (o *DestinationBankAccount) HasBeneficiaryProvince() bool`

HasBeneficiaryProvince returns a boolean if a field has been set.

### GetBeneficiaryPostCode

`func (o *DestinationBankAccount) GetBeneficiaryPostCode() string`

GetBeneficiaryPostCode returns the BeneficiaryPostCode field if non-nil, zero value otherwise.

### GetBeneficiaryPostCodeOk

`func (o *DestinationBankAccount) GetBeneficiaryPostCodeOk() (*string, bool)`

GetBeneficiaryPostCodeOk returns a tuple with the BeneficiaryPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPostCode

`func (o *DestinationBankAccount) SetBeneficiaryPostCode(v string)`

SetBeneficiaryPostCode sets BeneficiaryPostCode field to given value.

### HasBeneficiaryPostCode

`func (o *DestinationBankAccount) HasBeneficiaryPostCode() bool`

HasBeneficiaryPostCode returns a boolean if a field has been set.

### GetBankAccountName

`func (o *DestinationBankAccount) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *DestinationBankAccount) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *DestinationBankAccount) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountName

`func (o *DestinationBankAccount) HasBankAccountName() bool`

HasBankAccountName returns a boolean if a field has been set.

### GetBankBranchCode

`func (o *DestinationBankAccount) GetBankBranchCode() string`

GetBankBranchCode returns the BankBranchCode field if non-nil, zero value otherwise.

### GetBankBranchCodeOk

`func (o *DestinationBankAccount) GetBankBranchCodeOk() (*string, bool)`

GetBankBranchCodeOk returns a tuple with the BankBranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranchCode

`func (o *DestinationBankAccount) SetBankBranchCode(v string)`

SetBankBranchCode sets BankBranchCode field to given value.

### HasBankBranchCode

`func (o *DestinationBankAccount) HasBankBranchCode() bool`

HasBankBranchCode returns a boolean if a field has been set.

### GetBankCountry

`func (o *DestinationBankAccount) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *DestinationBankAccount) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *DestinationBankAccount) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.

### HasBankCountry

`func (o *DestinationBankAccount) HasBankCountry() bool`

HasBankCountry returns a boolean if a field has been set.

### GetBankProvince

`func (o *DestinationBankAccount) GetBankProvince() string`

GetBankProvince returns the BankProvince field if non-nil, zero value otherwise.

### GetBankProvinceOk

`func (o *DestinationBankAccount) GetBankProvinceOk() (*string, bool)`

GetBankProvinceOk returns a tuple with the BankProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProvince

`func (o *DestinationBankAccount) SetBankProvince(v string)`

SetBankProvince sets BankProvince field to given value.

### HasBankProvince

`func (o *DestinationBankAccount) HasBankProvince() bool`

HasBankProvince returns a boolean if a field has been set.

### GetContractFileId

`func (o *DestinationBankAccount) GetContractFileId() string`

GetContractFileId returns the ContractFileId field if non-nil, zero value otherwise.

### GetContractFileIdOk

`func (o *DestinationBankAccount) GetContractFileIdOk() (*string, bool)`

GetContractFileIdOk returns a tuple with the ContractFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractFileId

`func (o *DestinationBankAccount) SetContractFileId(v string)`

SetContractFileId sets ContractFileId field to given value.

### HasContractFileId

`func (o *DestinationBankAccount) HasContractFileId() bool`

HasContractFileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



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
**BankCity** | Pointer to **string** | The city of the bank. | [optional] 
**RoutingValue** | Pointer to **string** | The routing value of the bank account. | [optional] 
**ContractFileId** | Pointer to **string** | The file ID of the contract document (e.g., cooperation agreement) that proves the business relationship between you and the beneficiary, which you can retrieve by calling [Upload file](https://www.cobo.com/developers/v2/api-references/payment/upload-file).  | [optional] 

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

### GetCountry

`func (o *UpdateDestinationBankAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateDestinationBankAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateDestinationBankAccount) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateDestinationBankAccount) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *UpdateDestinationBankAccount) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateDestinationBankAccount) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateDestinationBankAccount) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateDestinationBankAccount) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UpdateDestinationBankAccount) GetPaymentMethod() BankAccountPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UpdateDestinationBankAccount) GetPaymentMethodOk() (*BankAccountPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UpdateDestinationBankAccount) SetPaymentMethod(v BankAccountPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UpdateDestinationBankAccount) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetHolderType

`func (o *UpdateDestinationBankAccount) GetHolderType() BankAccountHolderType`

GetHolderType returns the HolderType field if non-nil, zero value otherwise.

### GetHolderTypeOk

`func (o *UpdateDestinationBankAccount) GetHolderTypeOk() (*BankAccountHolderType, bool)`

GetHolderTypeOk returns a tuple with the HolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderType

`func (o *UpdateDestinationBankAccount) SetHolderType(v BankAccountHolderType)`

SetHolderType sets HolderType field to given value.

### HasHolderType

`func (o *UpdateDestinationBankAccount) HasHolderType() bool`

HasHolderType returns a boolean if a field has been set.

### GetBeneficiaryProvince

`func (o *UpdateDestinationBankAccount) GetBeneficiaryProvince() string`

GetBeneficiaryProvince returns the BeneficiaryProvince field if non-nil, zero value otherwise.

### GetBeneficiaryProvinceOk

`func (o *UpdateDestinationBankAccount) GetBeneficiaryProvinceOk() (*string, bool)`

GetBeneficiaryProvinceOk returns a tuple with the BeneficiaryProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryProvince

`func (o *UpdateDestinationBankAccount) SetBeneficiaryProvince(v string)`

SetBeneficiaryProvince sets BeneficiaryProvince field to given value.

### HasBeneficiaryProvince

`func (o *UpdateDestinationBankAccount) HasBeneficiaryProvince() bool`

HasBeneficiaryProvince returns a boolean if a field has been set.

### GetBeneficiaryPostCode

`func (o *UpdateDestinationBankAccount) GetBeneficiaryPostCode() string`

GetBeneficiaryPostCode returns the BeneficiaryPostCode field if non-nil, zero value otherwise.

### GetBeneficiaryPostCodeOk

`func (o *UpdateDestinationBankAccount) GetBeneficiaryPostCodeOk() (*string, bool)`

GetBeneficiaryPostCodeOk returns a tuple with the BeneficiaryPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPostCode

`func (o *UpdateDestinationBankAccount) SetBeneficiaryPostCode(v string)`

SetBeneficiaryPostCode sets BeneficiaryPostCode field to given value.

### HasBeneficiaryPostCode

`func (o *UpdateDestinationBankAccount) HasBeneficiaryPostCode() bool`

HasBeneficiaryPostCode returns a boolean if a field has been set.

### GetBankAccountName

`func (o *UpdateDestinationBankAccount) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *UpdateDestinationBankAccount) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *UpdateDestinationBankAccount) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountName

`func (o *UpdateDestinationBankAccount) HasBankAccountName() bool`

HasBankAccountName returns a boolean if a field has been set.

### GetBankBranchCode

`func (o *UpdateDestinationBankAccount) GetBankBranchCode() string`

GetBankBranchCode returns the BankBranchCode field if non-nil, zero value otherwise.

### GetBankBranchCodeOk

`func (o *UpdateDestinationBankAccount) GetBankBranchCodeOk() (*string, bool)`

GetBankBranchCodeOk returns a tuple with the BankBranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranchCode

`func (o *UpdateDestinationBankAccount) SetBankBranchCode(v string)`

SetBankBranchCode sets BankBranchCode field to given value.

### HasBankBranchCode

`func (o *UpdateDestinationBankAccount) HasBankBranchCode() bool`

HasBankBranchCode returns a boolean if a field has been set.

### GetBankCountry

`func (o *UpdateDestinationBankAccount) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *UpdateDestinationBankAccount) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *UpdateDestinationBankAccount) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.

### HasBankCountry

`func (o *UpdateDestinationBankAccount) HasBankCountry() bool`

HasBankCountry returns a boolean if a field has been set.

### GetBankProvince

`func (o *UpdateDestinationBankAccount) GetBankProvince() string`

GetBankProvince returns the BankProvince field if non-nil, zero value otherwise.

### GetBankProvinceOk

`func (o *UpdateDestinationBankAccount) GetBankProvinceOk() (*string, bool)`

GetBankProvinceOk returns a tuple with the BankProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProvince

`func (o *UpdateDestinationBankAccount) SetBankProvince(v string)`

SetBankProvince sets BankProvince field to given value.

### HasBankProvince

`func (o *UpdateDestinationBankAccount) HasBankProvince() bool`

HasBankProvince returns a boolean if a field has been set.

### GetBankCity

`func (o *UpdateDestinationBankAccount) GetBankCity() string`

GetBankCity returns the BankCity field if non-nil, zero value otherwise.

### GetBankCityOk

`func (o *UpdateDestinationBankAccount) GetBankCityOk() (*string, bool)`

GetBankCityOk returns a tuple with the BankCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCity

`func (o *UpdateDestinationBankAccount) SetBankCity(v string)`

SetBankCity sets BankCity field to given value.

### HasBankCity

`func (o *UpdateDestinationBankAccount) HasBankCity() bool`

HasBankCity returns a boolean if a field has been set.

### GetRoutingValue

`func (o *UpdateDestinationBankAccount) GetRoutingValue() string`

GetRoutingValue returns the RoutingValue field if non-nil, zero value otherwise.

### GetRoutingValueOk

`func (o *UpdateDestinationBankAccount) GetRoutingValueOk() (*string, bool)`

GetRoutingValueOk returns a tuple with the RoutingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingValue

`func (o *UpdateDestinationBankAccount) SetRoutingValue(v string)`

SetRoutingValue sets RoutingValue field to given value.

### HasRoutingValue

`func (o *UpdateDestinationBankAccount) HasRoutingValue() bool`

HasRoutingValue returns a boolean if a field has been set.

### GetContractFileId

`func (o *UpdateDestinationBankAccount) GetContractFileId() string`

GetContractFileId returns the ContractFileId field if non-nil, zero value otherwise.

### GetContractFileIdOk

`func (o *UpdateDestinationBankAccount) GetContractFileIdOk() (*string, bool)`

GetContractFileIdOk returns a tuple with the ContractFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractFileId

`func (o *UpdateDestinationBankAccount) SetContractFileId(v string)`

SetContractFileId sets ContractFileId field to given value.

### HasContractFileId

`func (o *UpdateDestinationBankAccount) HasContractFileId() bool`

HasContractFileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



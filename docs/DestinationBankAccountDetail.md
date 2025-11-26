# DestinationBankAccountDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | The destination ID. | 
**DestinationName** | **string** | The name of the destination. | 
**DestinationType** | [**DestinationType**](DestinationType.md) |  | 
**DestinationEmail** | Pointer to **string** | The email of the destination. | [optional] 
**DestinationCountry** | Pointer to **string** | The country of the destination, in ISO 3166-1 alpha-3 format. | [optional] 
**DestinationContactAddress** | Pointer to **string** | The contact address of the destination. | [optional] 
**DestinationMerchantId** | Pointer to **string** | The ID of the merchant linked to the destination. | [optional] 
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

## Methods

### NewDestinationBankAccountDetail

`func NewDestinationBankAccountDetail(destinationId string, destinationName string, destinationType DestinationType, bankAccountId string, accountAlias string, accountNumber string, swiftCode string, currency string, beneficiaryName string, beneficiaryAddress string, bankName string, bankAddress string, bankAccountStatus BankAccountStatus, ) *DestinationBankAccountDetail`

NewDestinationBankAccountDetail instantiates a new DestinationBankAccountDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationBankAccountDetailWithDefaults

`func NewDestinationBankAccountDetailWithDefaults() *DestinationBankAccountDetail`

NewDestinationBankAccountDetailWithDefaults instantiates a new DestinationBankAccountDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *DestinationBankAccountDetail) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationBankAccountDetail) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationBankAccountDetail) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetDestinationName

`func (o *DestinationBankAccountDetail) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *DestinationBankAccountDetail) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *DestinationBankAccountDetail) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetDestinationType

`func (o *DestinationBankAccountDetail) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *DestinationBankAccountDetail) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *DestinationBankAccountDetail) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetDestinationEmail

`func (o *DestinationBankAccountDetail) GetDestinationEmail() string`

GetDestinationEmail returns the DestinationEmail field if non-nil, zero value otherwise.

### GetDestinationEmailOk

`func (o *DestinationBankAccountDetail) GetDestinationEmailOk() (*string, bool)`

GetDestinationEmailOk returns a tuple with the DestinationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationEmail

`func (o *DestinationBankAccountDetail) SetDestinationEmail(v string)`

SetDestinationEmail sets DestinationEmail field to given value.

### HasDestinationEmail

`func (o *DestinationBankAccountDetail) HasDestinationEmail() bool`

HasDestinationEmail returns a boolean if a field has been set.

### GetDestinationCountry

`func (o *DestinationBankAccountDetail) GetDestinationCountry() string`

GetDestinationCountry returns the DestinationCountry field if non-nil, zero value otherwise.

### GetDestinationCountryOk

`func (o *DestinationBankAccountDetail) GetDestinationCountryOk() (*string, bool)`

GetDestinationCountryOk returns a tuple with the DestinationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCountry

`func (o *DestinationBankAccountDetail) SetDestinationCountry(v string)`

SetDestinationCountry sets DestinationCountry field to given value.

### HasDestinationCountry

`func (o *DestinationBankAccountDetail) HasDestinationCountry() bool`

HasDestinationCountry returns a boolean if a field has been set.

### GetDestinationContactAddress

`func (o *DestinationBankAccountDetail) GetDestinationContactAddress() string`

GetDestinationContactAddress returns the DestinationContactAddress field if non-nil, zero value otherwise.

### GetDestinationContactAddressOk

`func (o *DestinationBankAccountDetail) GetDestinationContactAddressOk() (*string, bool)`

GetDestinationContactAddressOk returns a tuple with the DestinationContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContactAddress

`func (o *DestinationBankAccountDetail) SetDestinationContactAddress(v string)`

SetDestinationContactAddress sets DestinationContactAddress field to given value.

### HasDestinationContactAddress

`func (o *DestinationBankAccountDetail) HasDestinationContactAddress() bool`

HasDestinationContactAddress returns a boolean if a field has been set.

### GetDestinationMerchantId

`func (o *DestinationBankAccountDetail) GetDestinationMerchantId() string`

GetDestinationMerchantId returns the DestinationMerchantId field if non-nil, zero value otherwise.

### GetDestinationMerchantIdOk

`func (o *DestinationBankAccountDetail) GetDestinationMerchantIdOk() (*string, bool)`

GetDestinationMerchantIdOk returns a tuple with the DestinationMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMerchantId

`func (o *DestinationBankAccountDetail) SetDestinationMerchantId(v string)`

SetDestinationMerchantId sets DestinationMerchantId field to given value.

### HasDestinationMerchantId

`func (o *DestinationBankAccountDetail) HasDestinationMerchantId() bool`

HasDestinationMerchantId returns a boolean if a field has been set.

### GetBankAccountId

`func (o *DestinationBankAccountDetail) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *DestinationBankAccountDetail) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *DestinationBankAccountDetail) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.


### GetAccountAlias

`func (o *DestinationBankAccountDetail) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *DestinationBankAccountDetail) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *DestinationBankAccountDetail) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.


### GetAccountNumber

`func (o *DestinationBankAccountDetail) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DestinationBankAccountDetail) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DestinationBankAccountDetail) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetSwiftCode

`func (o *DestinationBankAccountDetail) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *DestinationBankAccountDetail) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *DestinationBankAccountDetail) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.


### GetCurrency

`func (o *DestinationBankAccountDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DestinationBankAccountDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DestinationBankAccountDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBeneficiaryName

`func (o *DestinationBankAccountDetail) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *DestinationBankAccountDetail) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *DestinationBankAccountDetail) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.


### GetBeneficiaryAddress

`func (o *DestinationBankAccountDetail) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *DestinationBankAccountDetail) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *DestinationBankAccountDetail) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.


### GetBankName

`func (o *DestinationBankAccountDetail) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *DestinationBankAccountDetail) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *DestinationBankAccountDetail) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankAddress

`func (o *DestinationBankAccountDetail) GetBankAddress() string`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *DestinationBankAccountDetail) GetBankAddressOk() (*string, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *DestinationBankAccountDetail) SetBankAddress(v string)`

SetBankAddress sets BankAddress field to given value.


### GetIbanCode

`func (o *DestinationBankAccountDetail) GetIbanCode() string`

GetIbanCode returns the IbanCode field if non-nil, zero value otherwise.

### GetIbanCodeOk

`func (o *DestinationBankAccountDetail) GetIbanCodeOk() (*string, bool)`

GetIbanCodeOk returns a tuple with the IbanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanCode

`func (o *DestinationBankAccountDetail) SetIbanCode(v string)`

SetIbanCode sets IbanCode field to given value.

### HasIbanCode

`func (o *DestinationBankAccountDetail) HasIbanCode() bool`

HasIbanCode returns a boolean if a field has been set.

### GetFurtherCredit

`func (o *DestinationBankAccountDetail) GetFurtherCredit() string`

GetFurtherCredit returns the FurtherCredit field if non-nil, zero value otherwise.

### GetFurtherCreditOk

`func (o *DestinationBankAccountDetail) GetFurtherCreditOk() (*string, bool)`

GetFurtherCreditOk returns a tuple with the FurtherCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurtherCredit

`func (o *DestinationBankAccountDetail) SetFurtherCredit(v string)`

SetFurtherCredit sets FurtherCredit field to given value.

### HasFurtherCredit

`func (o *DestinationBankAccountDetail) HasFurtherCredit() bool`

HasFurtherCredit returns a boolean if a field has been set.

### GetIntermediaryBankInfo

`func (o *DestinationBankAccountDetail) GetIntermediaryBankInfo() IntermediaryBankInfo`

GetIntermediaryBankInfo returns the IntermediaryBankInfo field if non-nil, zero value otherwise.

### GetIntermediaryBankInfoOk

`func (o *DestinationBankAccountDetail) GetIntermediaryBankInfoOk() (*IntermediaryBankInfo, bool)`

GetIntermediaryBankInfoOk returns a tuple with the IntermediaryBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryBankInfo

`func (o *DestinationBankAccountDetail) SetIntermediaryBankInfo(v IntermediaryBankInfo)`

SetIntermediaryBankInfo sets IntermediaryBankInfo field to given value.

### HasIntermediaryBankInfo

`func (o *DestinationBankAccountDetail) HasIntermediaryBankInfo() bool`

HasIntermediaryBankInfo returns a boolean if a field has been set.

### GetBankAccountStatus

`func (o *DestinationBankAccountDetail) GetBankAccountStatus() BankAccountStatus`

GetBankAccountStatus returns the BankAccountStatus field if non-nil, zero value otherwise.

### GetBankAccountStatusOk

`func (o *DestinationBankAccountDetail) GetBankAccountStatusOk() (*BankAccountStatus, bool)`

GetBankAccountStatusOk returns a tuple with the BankAccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountStatus

`func (o *DestinationBankAccountDetail) SetBankAccountStatus(v BankAccountStatus)`

SetBankAccountStatus sets BankAccountStatus field to given value.


### GetCreatedTimestamp

`func (o *DestinationBankAccountDetail) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DestinationBankAccountDetail) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DestinationBankAccountDetail) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *DestinationBankAccountDetail) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DestinationBankAccountDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DestinationBankAccountDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DestinationBankAccountDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *DestinationBankAccountDetail) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



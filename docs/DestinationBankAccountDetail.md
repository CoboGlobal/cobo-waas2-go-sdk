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
**Tag** | Pointer to [**NullableDestinationBankAccountTag**](DestinationBankAccountTag.md) |  | [optional] 
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
**BankCity** | Pointer to **string** | The city of the bank. | [optional] 
**RoutingValue** | Pointer to **string** | The routing value of the bank account. | [optional] 
**ContractFileId** | Pointer to **string** | The file ID of the contract document (e.g., cooperation agreement) that proves the business relationship between you and the beneficiary, which you can retrieve by calling [Upload file](https://www.cobo.com/developers/v2/api-references/payment/upload-file).  | [optional] 
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


### GetTag

`func (o *DestinationBankAccountDetail) GetTag() DestinationBankAccountTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DestinationBankAccountDetail) GetTagOk() (*DestinationBankAccountTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DestinationBankAccountDetail) SetTag(v DestinationBankAccountTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DestinationBankAccountDetail) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *DestinationBankAccountDetail) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *DestinationBankAccountDetail) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
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


### GetCountry

`func (o *DestinationBankAccountDetail) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DestinationBankAccountDetail) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DestinationBankAccountDetail) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DestinationBankAccountDetail) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *DestinationBankAccountDetail) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DestinationBankAccountDetail) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DestinationBankAccountDetail) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DestinationBankAccountDetail) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *DestinationBankAccountDetail) GetPaymentMethod() BankAccountPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *DestinationBankAccountDetail) GetPaymentMethodOk() (*BankAccountPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *DestinationBankAccountDetail) SetPaymentMethod(v BankAccountPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *DestinationBankAccountDetail) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetHolderType

`func (o *DestinationBankAccountDetail) GetHolderType() BankAccountHolderType`

GetHolderType returns the HolderType field if non-nil, zero value otherwise.

### GetHolderTypeOk

`func (o *DestinationBankAccountDetail) GetHolderTypeOk() (*BankAccountHolderType, bool)`

GetHolderTypeOk returns a tuple with the HolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderType

`func (o *DestinationBankAccountDetail) SetHolderType(v BankAccountHolderType)`

SetHolderType sets HolderType field to given value.

### HasHolderType

`func (o *DestinationBankAccountDetail) HasHolderType() bool`

HasHolderType returns a boolean if a field has been set.

### GetBeneficiaryProvince

`func (o *DestinationBankAccountDetail) GetBeneficiaryProvince() string`

GetBeneficiaryProvince returns the BeneficiaryProvince field if non-nil, zero value otherwise.

### GetBeneficiaryProvinceOk

`func (o *DestinationBankAccountDetail) GetBeneficiaryProvinceOk() (*string, bool)`

GetBeneficiaryProvinceOk returns a tuple with the BeneficiaryProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryProvince

`func (o *DestinationBankAccountDetail) SetBeneficiaryProvince(v string)`

SetBeneficiaryProvince sets BeneficiaryProvince field to given value.

### HasBeneficiaryProvince

`func (o *DestinationBankAccountDetail) HasBeneficiaryProvince() bool`

HasBeneficiaryProvince returns a boolean if a field has been set.

### GetBeneficiaryPostCode

`func (o *DestinationBankAccountDetail) GetBeneficiaryPostCode() string`

GetBeneficiaryPostCode returns the BeneficiaryPostCode field if non-nil, zero value otherwise.

### GetBeneficiaryPostCodeOk

`func (o *DestinationBankAccountDetail) GetBeneficiaryPostCodeOk() (*string, bool)`

GetBeneficiaryPostCodeOk returns a tuple with the BeneficiaryPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPostCode

`func (o *DestinationBankAccountDetail) SetBeneficiaryPostCode(v string)`

SetBeneficiaryPostCode sets BeneficiaryPostCode field to given value.

### HasBeneficiaryPostCode

`func (o *DestinationBankAccountDetail) HasBeneficiaryPostCode() bool`

HasBeneficiaryPostCode returns a boolean if a field has been set.

### GetBankAccountName

`func (o *DestinationBankAccountDetail) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *DestinationBankAccountDetail) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *DestinationBankAccountDetail) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountName

`func (o *DestinationBankAccountDetail) HasBankAccountName() bool`

HasBankAccountName returns a boolean if a field has been set.

### GetBankBranchCode

`func (o *DestinationBankAccountDetail) GetBankBranchCode() string`

GetBankBranchCode returns the BankBranchCode field if non-nil, zero value otherwise.

### GetBankBranchCodeOk

`func (o *DestinationBankAccountDetail) GetBankBranchCodeOk() (*string, bool)`

GetBankBranchCodeOk returns a tuple with the BankBranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranchCode

`func (o *DestinationBankAccountDetail) SetBankBranchCode(v string)`

SetBankBranchCode sets BankBranchCode field to given value.

### HasBankBranchCode

`func (o *DestinationBankAccountDetail) HasBankBranchCode() bool`

HasBankBranchCode returns a boolean if a field has been set.

### GetBankCountry

`func (o *DestinationBankAccountDetail) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *DestinationBankAccountDetail) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *DestinationBankAccountDetail) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.

### HasBankCountry

`func (o *DestinationBankAccountDetail) HasBankCountry() bool`

HasBankCountry returns a boolean if a field has been set.

### GetBankProvince

`func (o *DestinationBankAccountDetail) GetBankProvince() string`

GetBankProvince returns the BankProvince field if non-nil, zero value otherwise.

### GetBankProvinceOk

`func (o *DestinationBankAccountDetail) GetBankProvinceOk() (*string, bool)`

GetBankProvinceOk returns a tuple with the BankProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProvince

`func (o *DestinationBankAccountDetail) SetBankProvince(v string)`

SetBankProvince sets BankProvince field to given value.

### HasBankProvince

`func (o *DestinationBankAccountDetail) HasBankProvince() bool`

HasBankProvince returns a boolean if a field has been set.

### GetBankCity

`func (o *DestinationBankAccountDetail) GetBankCity() string`

GetBankCity returns the BankCity field if non-nil, zero value otherwise.

### GetBankCityOk

`func (o *DestinationBankAccountDetail) GetBankCityOk() (*string, bool)`

GetBankCityOk returns a tuple with the BankCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCity

`func (o *DestinationBankAccountDetail) SetBankCity(v string)`

SetBankCity sets BankCity field to given value.

### HasBankCity

`func (o *DestinationBankAccountDetail) HasBankCity() bool`

HasBankCity returns a boolean if a field has been set.

### GetRoutingValue

`func (o *DestinationBankAccountDetail) GetRoutingValue() string`

GetRoutingValue returns the RoutingValue field if non-nil, zero value otherwise.

### GetRoutingValueOk

`func (o *DestinationBankAccountDetail) GetRoutingValueOk() (*string, bool)`

GetRoutingValueOk returns a tuple with the RoutingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingValue

`func (o *DestinationBankAccountDetail) SetRoutingValue(v string)`

SetRoutingValue sets RoutingValue field to given value.

### HasRoutingValue

`func (o *DestinationBankAccountDetail) HasRoutingValue() bool`

HasRoutingValue returns a boolean if a field has been set.

### GetContractFileId

`func (o *DestinationBankAccountDetail) GetContractFileId() string`

GetContractFileId returns the ContractFileId field if non-nil, zero value otherwise.

### GetContractFileIdOk

`func (o *DestinationBankAccountDetail) GetContractFileIdOk() (*string, bool)`

GetContractFileIdOk returns a tuple with the ContractFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractFileId

`func (o *DestinationBankAccountDetail) SetContractFileId(v string)`

SetContractFileId sets ContractFileId field to given value.

### HasContractFileId

`func (o *DestinationBankAccountDetail) HasContractFileId() bool`

HasContractFileId returns a boolean if a field has been set.

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



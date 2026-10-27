# FiatEntryKycDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiatEntryKycId** | **string** | The fiat entry KYC ID. | 
**Status** | [**FiatEntryKycStatus**](FiatEntryKycStatus.md) |  | 
**MerchantType** | [**FiatEntryKycMerchantType**](FiatEntryKycMerchantType.md) |  | 
**Country** | **string** | The country/region of the fiat entry KYC, in ISO 3166-1 alpha-3 format. | 
**Industry** | **[]string** | The industry categories of the fiat entry KYC. | 
**Email** | Pointer to **string** | The email address. | [optional] 
**Phone** | Pointer to **string** | The phone number. | [optional] 
**CompanyInfo** | Pointer to [**FiatEntryKycCompanyInfo**](FiatEntryKycCompanyInfo.md) |  | [optional] 
**IndividualInfo** | Pointer to [**FiatEntryKycPersonInfo**](FiatEntryKycPersonInfo.md) |  | [optional] 
**CreatedTimestamp** | **int64** | The creation timestamp in Unix seconds. | 
**UpdatedTimestamp** | Pointer to **int64** | The last update timestamp in Unix seconds. | [optional] 

## Methods

### NewFiatEntryKycDetail

`func NewFiatEntryKycDetail(fiatEntryKycId string, status FiatEntryKycStatus, merchantType FiatEntryKycMerchantType, country string, industry []string, createdTimestamp int64, ) *FiatEntryKycDetail`

NewFiatEntryKycDetail instantiates a new FiatEntryKycDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatEntryKycDetailWithDefaults

`func NewFiatEntryKycDetailWithDefaults() *FiatEntryKycDetail`

NewFiatEntryKycDetailWithDefaults instantiates a new FiatEntryKycDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiatEntryKycId

`func (o *FiatEntryKycDetail) GetFiatEntryKycId() string`

GetFiatEntryKycId returns the FiatEntryKycId field if non-nil, zero value otherwise.

### GetFiatEntryKycIdOk

`func (o *FiatEntryKycDetail) GetFiatEntryKycIdOk() (*string, bool)`

GetFiatEntryKycIdOk returns a tuple with the FiatEntryKycId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatEntryKycId

`func (o *FiatEntryKycDetail) SetFiatEntryKycId(v string)`

SetFiatEntryKycId sets FiatEntryKycId field to given value.


### GetStatus

`func (o *FiatEntryKycDetail) GetStatus() FiatEntryKycStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiatEntryKycDetail) GetStatusOk() (*FiatEntryKycStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiatEntryKycDetail) SetStatus(v FiatEntryKycStatus)`

SetStatus sets Status field to given value.


### GetMerchantType

`func (o *FiatEntryKycDetail) GetMerchantType() FiatEntryKycMerchantType`

GetMerchantType returns the MerchantType field if non-nil, zero value otherwise.

### GetMerchantTypeOk

`func (o *FiatEntryKycDetail) GetMerchantTypeOk() (*FiatEntryKycMerchantType, bool)`

GetMerchantTypeOk returns a tuple with the MerchantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantType

`func (o *FiatEntryKycDetail) SetMerchantType(v FiatEntryKycMerchantType)`

SetMerchantType sets MerchantType field to given value.


### GetCountry

`func (o *FiatEntryKycDetail) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FiatEntryKycDetail) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FiatEntryKycDetail) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetIndustry

`func (o *FiatEntryKycDetail) GetIndustry() []string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *FiatEntryKycDetail) GetIndustryOk() (*[]string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *FiatEntryKycDetail) SetIndustry(v []string)`

SetIndustry sets Industry field to given value.


### GetEmail

`func (o *FiatEntryKycDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FiatEntryKycDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FiatEntryKycDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FiatEntryKycDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *FiatEntryKycDetail) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *FiatEntryKycDetail) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *FiatEntryKycDetail) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *FiatEntryKycDetail) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCompanyInfo

`func (o *FiatEntryKycDetail) GetCompanyInfo() FiatEntryKycCompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *FiatEntryKycDetail) GetCompanyInfoOk() (*FiatEntryKycCompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *FiatEntryKycDetail) SetCompanyInfo(v FiatEntryKycCompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.

### HasCompanyInfo

`func (o *FiatEntryKycDetail) HasCompanyInfo() bool`

HasCompanyInfo returns a boolean if a field has been set.

### GetIndividualInfo

`func (o *FiatEntryKycDetail) GetIndividualInfo() FiatEntryKycPersonInfo`

GetIndividualInfo returns the IndividualInfo field if non-nil, zero value otherwise.

### GetIndividualInfoOk

`func (o *FiatEntryKycDetail) GetIndividualInfoOk() (*FiatEntryKycPersonInfo, bool)`

GetIndividualInfoOk returns a tuple with the IndividualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualInfo

`func (o *FiatEntryKycDetail) SetIndividualInfo(v FiatEntryKycPersonInfo)`

SetIndividualInfo sets IndividualInfo field to given value.

### HasIndividualInfo

`func (o *FiatEntryKycDetail) HasIndividualInfo() bool`

HasIndividualInfo returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *FiatEntryKycDetail) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FiatEntryKycDetail) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FiatEntryKycDetail) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *FiatEntryKycDetail) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *FiatEntryKycDetail) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *FiatEntryKycDetail) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *FiatEntryKycDetail) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



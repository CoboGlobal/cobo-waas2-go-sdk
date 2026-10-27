# UpdateFiatEntryKyc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantType** | Pointer to [**FiatEntryKycMerchantType**](FiatEntryKycMerchantType.md) |  | [optional] 
**Country** | Pointer to **string** | The country/region of the fiat entry KYC, in ISO 3166-1 alpha-3 format. | [optional] 
**Industry** | Pointer to **[]string** | The industry categories of the fiat entry KYC. | [optional] 
**CompanyInfo** | Pointer to [**FiatEntryKycCompanyInfo**](FiatEntryKycCompanyInfo.md) |  | [optional] 
**IndividualInfo** | Pointer to [**FiatEntryKycPersonInfo**](FiatEntryKycPersonInfo.md) |  | [optional] 

## Methods

### NewUpdateFiatEntryKyc

`func NewUpdateFiatEntryKyc() *UpdateFiatEntryKyc`

NewUpdateFiatEntryKyc instantiates a new UpdateFiatEntryKyc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFiatEntryKycWithDefaults

`func NewUpdateFiatEntryKycWithDefaults() *UpdateFiatEntryKyc`

NewUpdateFiatEntryKycWithDefaults instantiates a new UpdateFiatEntryKyc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantType

`func (o *UpdateFiatEntryKyc) GetMerchantType() FiatEntryKycMerchantType`

GetMerchantType returns the MerchantType field if non-nil, zero value otherwise.

### GetMerchantTypeOk

`func (o *UpdateFiatEntryKyc) GetMerchantTypeOk() (*FiatEntryKycMerchantType, bool)`

GetMerchantTypeOk returns a tuple with the MerchantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantType

`func (o *UpdateFiatEntryKyc) SetMerchantType(v FiatEntryKycMerchantType)`

SetMerchantType sets MerchantType field to given value.

### HasMerchantType

`func (o *UpdateFiatEntryKyc) HasMerchantType() bool`

HasMerchantType returns a boolean if a field has been set.

### GetCountry

`func (o *UpdateFiatEntryKyc) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateFiatEntryKyc) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateFiatEntryKyc) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateFiatEntryKyc) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIndustry

`func (o *UpdateFiatEntryKyc) GetIndustry() []string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *UpdateFiatEntryKyc) GetIndustryOk() (*[]string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *UpdateFiatEntryKyc) SetIndustry(v []string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *UpdateFiatEntryKyc) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetCompanyInfo

`func (o *UpdateFiatEntryKyc) GetCompanyInfo() FiatEntryKycCompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *UpdateFiatEntryKyc) GetCompanyInfoOk() (*FiatEntryKycCompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *UpdateFiatEntryKyc) SetCompanyInfo(v FiatEntryKycCompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.

### HasCompanyInfo

`func (o *UpdateFiatEntryKyc) HasCompanyInfo() bool`

HasCompanyInfo returns a boolean if a field has been set.

### GetIndividualInfo

`func (o *UpdateFiatEntryKyc) GetIndividualInfo() FiatEntryKycPersonInfo`

GetIndividualInfo returns the IndividualInfo field if non-nil, zero value otherwise.

### GetIndividualInfoOk

`func (o *UpdateFiatEntryKyc) GetIndividualInfoOk() (*FiatEntryKycPersonInfo, bool)`

GetIndividualInfoOk returns a tuple with the IndividualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualInfo

`func (o *UpdateFiatEntryKyc) SetIndividualInfo(v FiatEntryKycPersonInfo)`

SetIndividualInfo sets IndividualInfo field to given value.

### HasIndividualInfo

`func (o *UpdateFiatEntryKyc) HasIndividualInfo() bool`

HasIndividualInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



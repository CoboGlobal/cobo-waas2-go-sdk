# CreateFiatEntryKyc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantType** | [**FiatEntryKycMerchantType**](FiatEntryKycMerchantType.md) |  | 
**Country** | **string** | The country/region of the fiat entry KYC, in ISO 3166-1 alpha-3 format. | 
**Industry** | **[]string** | The industry categories of the fiat entry KYC. | 
**CompanyInfo** | Pointer to [**FiatEntryKycCompanyInfo**](FiatEntryKycCompanyInfo.md) |  | [optional] 
**IndividualInfo** | Pointer to [**FiatEntryKycPersonInfo**](FiatEntryKycPersonInfo.md) |  | [optional] 

## Methods

### NewCreateFiatEntryKyc

`func NewCreateFiatEntryKyc(merchantType FiatEntryKycMerchantType, country string, industry []string, ) *CreateFiatEntryKyc`

NewCreateFiatEntryKyc instantiates a new CreateFiatEntryKyc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFiatEntryKycWithDefaults

`func NewCreateFiatEntryKycWithDefaults() *CreateFiatEntryKyc`

NewCreateFiatEntryKycWithDefaults instantiates a new CreateFiatEntryKyc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantType

`func (o *CreateFiatEntryKyc) GetMerchantType() FiatEntryKycMerchantType`

GetMerchantType returns the MerchantType field if non-nil, zero value otherwise.

### GetMerchantTypeOk

`func (o *CreateFiatEntryKyc) GetMerchantTypeOk() (*FiatEntryKycMerchantType, bool)`

GetMerchantTypeOk returns a tuple with the MerchantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantType

`func (o *CreateFiatEntryKyc) SetMerchantType(v FiatEntryKycMerchantType)`

SetMerchantType sets MerchantType field to given value.


### GetCountry

`func (o *CreateFiatEntryKyc) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateFiatEntryKyc) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateFiatEntryKyc) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetIndustry

`func (o *CreateFiatEntryKyc) GetIndustry() []string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *CreateFiatEntryKyc) GetIndustryOk() (*[]string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *CreateFiatEntryKyc) SetIndustry(v []string)`

SetIndustry sets Industry field to given value.


### GetCompanyInfo

`func (o *CreateFiatEntryKyc) GetCompanyInfo() FiatEntryKycCompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *CreateFiatEntryKyc) GetCompanyInfoOk() (*FiatEntryKycCompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *CreateFiatEntryKyc) SetCompanyInfo(v FiatEntryKycCompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.

### HasCompanyInfo

`func (o *CreateFiatEntryKyc) HasCompanyInfo() bool`

HasCompanyInfo returns a boolean if a field has been set.

### GetIndividualInfo

`func (o *CreateFiatEntryKyc) GetIndividualInfo() FiatEntryKycPersonInfo`

GetIndividualInfo returns the IndividualInfo field if non-nil, zero value otherwise.

### GetIndividualInfoOk

`func (o *CreateFiatEntryKyc) GetIndividualInfoOk() (*FiatEntryKycPersonInfo, bool)`

GetIndividualInfoOk returns a tuple with the IndividualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualInfo

`func (o *CreateFiatEntryKyc) SetIndividualInfo(v FiatEntryKycPersonInfo)`

SetIndividualInfo sets IndividualInfo field to given value.

### HasIndividualInfo

`func (o *CreateFiatEntryKyc) HasIndividualInfo() bool`

HasIndividualInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SubmitMerchantKyc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The merchant email address. | 
**Phone** | **string** | The merchant phone number. | 
**MerchantType** | [**MerchantKycMerchantType**](MerchantKycMerchantType.md) |  | 
**Country** | **string** | The country/region of the merchant, in ISO 3166-1 alpha-3 format. | 
**Industry** | **[]string** | The industry categories of the merchant. | 
**CompanyInfo** | [**MerchantKycCompanyInfo**](MerchantKycCompanyInfo.md) |  | 

## Methods

### NewSubmitMerchantKyc

`func NewSubmitMerchantKyc(email string, phone string, merchantType MerchantKycMerchantType, country string, industry []string, companyInfo MerchantKycCompanyInfo, ) *SubmitMerchantKyc`

NewSubmitMerchantKyc instantiates a new SubmitMerchantKyc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitMerchantKycWithDefaults

`func NewSubmitMerchantKycWithDefaults() *SubmitMerchantKyc`

NewSubmitMerchantKycWithDefaults instantiates a new SubmitMerchantKyc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubmitMerchantKyc) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubmitMerchantKyc) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubmitMerchantKyc) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *SubmitMerchantKyc) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SubmitMerchantKyc) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SubmitMerchantKyc) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetMerchantType

`func (o *SubmitMerchantKyc) GetMerchantType() MerchantKycMerchantType`

GetMerchantType returns the MerchantType field if non-nil, zero value otherwise.

### GetMerchantTypeOk

`func (o *SubmitMerchantKyc) GetMerchantTypeOk() (*MerchantKycMerchantType, bool)`

GetMerchantTypeOk returns a tuple with the MerchantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantType

`func (o *SubmitMerchantKyc) SetMerchantType(v MerchantKycMerchantType)`

SetMerchantType sets MerchantType field to given value.


### GetCountry

`func (o *SubmitMerchantKyc) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SubmitMerchantKyc) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SubmitMerchantKyc) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetIndustry

`func (o *SubmitMerchantKyc) GetIndustry() []string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *SubmitMerchantKyc) GetIndustryOk() (*[]string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *SubmitMerchantKyc) SetIndustry(v []string)`

SetIndustry sets Industry field to given value.


### GetCompanyInfo

`func (o *SubmitMerchantKyc) GetCompanyInfo() MerchantKycCompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *SubmitMerchantKyc) GetCompanyInfoOk() (*MerchantKycCompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *SubmitMerchantKyc) SetCompanyInfo(v MerchantKycCompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



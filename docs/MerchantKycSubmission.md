# MerchantKycSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KycSubmissionId** | **string** | The KYC submission ID. | 
**MerchantId** | **string** | The merchant ID. | 
**Status** | [**MerchantKycStatus**](MerchantKycStatus.md) |  | 
**MerchantType** | [**MerchantKycMerchantType**](MerchantKycMerchantType.md) |  | 
**Country** | **string** | The country/region of the merchant, in ISO 3166-1 alpha-3 format. | 
**Industry** | **[]string** | The industry categories of the merchant. | 
**CompanyInfo** | Pointer to [**MerchantKycCompanyInfo**](MerchantKycCompanyInfo.md) |  | [optional] 
**IndividualInfo** | Pointer to [**MerchantKycPersonInfo**](MerchantKycPersonInfo.md) |  | [optional] 
**CreatedTimestamp** | **int64** | The creation timestamp in Unix seconds. | 
**UpdatedTimestamp** | Pointer to **int64** | The last update timestamp in Unix seconds. | [optional] 

## Methods

### NewMerchantKycSubmission

`func NewMerchantKycSubmission(kycSubmissionId string, merchantId string, status MerchantKycStatus, merchantType MerchantKycMerchantType, country string, industry []string, createdTimestamp int64, ) *MerchantKycSubmission`

NewMerchantKycSubmission instantiates a new MerchantKycSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantKycSubmissionWithDefaults

`func NewMerchantKycSubmissionWithDefaults() *MerchantKycSubmission`

NewMerchantKycSubmissionWithDefaults instantiates a new MerchantKycSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKycSubmissionId

`func (o *MerchantKycSubmission) GetKycSubmissionId() string`

GetKycSubmissionId returns the KycSubmissionId field if non-nil, zero value otherwise.

### GetKycSubmissionIdOk

`func (o *MerchantKycSubmission) GetKycSubmissionIdOk() (*string, bool)`

GetKycSubmissionIdOk returns a tuple with the KycSubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycSubmissionId

`func (o *MerchantKycSubmission) SetKycSubmissionId(v string)`

SetKycSubmissionId sets KycSubmissionId field to given value.


### GetMerchantId

`func (o *MerchantKycSubmission) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *MerchantKycSubmission) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *MerchantKycSubmission) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetStatus

`func (o *MerchantKycSubmission) GetStatus() MerchantKycStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantKycSubmission) GetStatusOk() (*MerchantKycStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantKycSubmission) SetStatus(v MerchantKycStatus)`

SetStatus sets Status field to given value.


### GetMerchantType

`func (o *MerchantKycSubmission) GetMerchantType() MerchantKycMerchantType`

GetMerchantType returns the MerchantType field if non-nil, zero value otherwise.

### GetMerchantTypeOk

`func (o *MerchantKycSubmission) GetMerchantTypeOk() (*MerchantKycMerchantType, bool)`

GetMerchantTypeOk returns a tuple with the MerchantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantType

`func (o *MerchantKycSubmission) SetMerchantType(v MerchantKycMerchantType)`

SetMerchantType sets MerchantType field to given value.


### GetCountry

`func (o *MerchantKycSubmission) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MerchantKycSubmission) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MerchantKycSubmission) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetIndustry

`func (o *MerchantKycSubmission) GetIndustry() []string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *MerchantKycSubmission) GetIndustryOk() (*[]string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *MerchantKycSubmission) SetIndustry(v []string)`

SetIndustry sets Industry field to given value.


### GetCompanyInfo

`func (o *MerchantKycSubmission) GetCompanyInfo() MerchantKycCompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *MerchantKycSubmission) GetCompanyInfoOk() (*MerchantKycCompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *MerchantKycSubmission) SetCompanyInfo(v MerchantKycCompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.

### HasCompanyInfo

`func (o *MerchantKycSubmission) HasCompanyInfo() bool`

HasCompanyInfo returns a boolean if a field has been set.

### GetIndividualInfo

`func (o *MerchantKycSubmission) GetIndividualInfo() MerchantKycPersonInfo`

GetIndividualInfo returns the IndividualInfo field if non-nil, zero value otherwise.

### GetIndividualInfoOk

`func (o *MerchantKycSubmission) GetIndividualInfoOk() (*MerchantKycPersonInfo, bool)`

GetIndividualInfoOk returns a tuple with the IndividualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualInfo

`func (o *MerchantKycSubmission) SetIndividualInfo(v MerchantKycPersonInfo)`

SetIndividualInfo sets IndividualInfo field to given value.

### HasIndividualInfo

`func (o *MerchantKycSubmission) HasIndividualInfo() bool`

HasIndividualInfo returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *MerchantKycSubmission) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *MerchantKycSubmission) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *MerchantKycSubmission) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *MerchantKycSubmission) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *MerchantKycSubmission) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *MerchantKycSubmission) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *MerchantKycSubmission) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



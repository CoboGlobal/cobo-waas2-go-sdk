# MerchantKycCompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyType** | [**MerchantKycCompanyType**](MerchantKycCompanyType.md) |  | 
**Listed** | **bool** | Whether the company is listed. | 
**Attachments** | [**[]MerchantKycCompanyAttachment**](MerchantKycCompanyAttachment.md) | The attachments of the company. | 
**OperationAddress** | [**MerchantKycAddress**](MerchantKycAddress.md) |  | 
**IdentifyNo** | **string** | The company identification number. | 
**CompanyName** | **string** | The company name in local language. | 
**CompanyNameEn** | **string** | The company name in English. | 
**EstablishDate** | **string** | The establishment date of the company. | 
**CommencementDate** | **string** | The commencement date of the company. | 
**ValidPeriod** | **string** | The valid period of the company registration. | 
**RegisterAddress** | [**MerchantKycAddress**](MerchantKycAddress.md) |  | 
**LegalInfo** | [**MerchantKycPersonInfo**](MerchantKycPersonInfo.md) |  | 
**UboInfos** | [**[]MerchantKycPersonInfo**](MerchantKycPersonInfo.md) | The ultimate beneficial owner information. | 
**OnlineStoreUrl** | Pointer to **string** | The online store URL. Required when merchant type is B2B. | [optional] 

## Methods

### NewMerchantKycCompanyInfo

`func NewMerchantKycCompanyInfo(companyType MerchantKycCompanyType, listed bool, attachments []MerchantKycCompanyAttachment, operationAddress MerchantKycAddress, identifyNo string, companyName string, companyNameEn string, establishDate string, commencementDate string, validPeriod string, registerAddress MerchantKycAddress, legalInfo MerchantKycPersonInfo, uboInfos []MerchantKycPersonInfo, ) *MerchantKycCompanyInfo`

NewMerchantKycCompanyInfo instantiates a new MerchantKycCompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantKycCompanyInfoWithDefaults

`func NewMerchantKycCompanyInfoWithDefaults() *MerchantKycCompanyInfo`

NewMerchantKycCompanyInfoWithDefaults instantiates a new MerchantKycCompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyType

`func (o *MerchantKycCompanyInfo) GetCompanyType() MerchantKycCompanyType`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *MerchantKycCompanyInfo) GetCompanyTypeOk() (*MerchantKycCompanyType, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *MerchantKycCompanyInfo) SetCompanyType(v MerchantKycCompanyType)`

SetCompanyType sets CompanyType field to given value.


### GetListed

`func (o *MerchantKycCompanyInfo) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *MerchantKycCompanyInfo) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *MerchantKycCompanyInfo) SetListed(v bool)`

SetListed sets Listed field to given value.


### GetAttachments

`func (o *MerchantKycCompanyInfo) GetAttachments() []MerchantKycCompanyAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MerchantKycCompanyInfo) GetAttachmentsOk() (*[]MerchantKycCompanyAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MerchantKycCompanyInfo) SetAttachments(v []MerchantKycCompanyAttachment)`

SetAttachments sets Attachments field to given value.


### GetOperationAddress

`func (o *MerchantKycCompanyInfo) GetOperationAddress() MerchantKycAddress`

GetOperationAddress returns the OperationAddress field if non-nil, zero value otherwise.

### GetOperationAddressOk

`func (o *MerchantKycCompanyInfo) GetOperationAddressOk() (*MerchantKycAddress, bool)`

GetOperationAddressOk returns a tuple with the OperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAddress

`func (o *MerchantKycCompanyInfo) SetOperationAddress(v MerchantKycAddress)`

SetOperationAddress sets OperationAddress field to given value.


### GetIdentifyNo

`func (o *MerchantKycCompanyInfo) GetIdentifyNo() string`

GetIdentifyNo returns the IdentifyNo field if non-nil, zero value otherwise.

### GetIdentifyNoOk

`func (o *MerchantKycCompanyInfo) GetIdentifyNoOk() (*string, bool)`

GetIdentifyNoOk returns a tuple with the IdentifyNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifyNo

`func (o *MerchantKycCompanyInfo) SetIdentifyNo(v string)`

SetIdentifyNo sets IdentifyNo field to given value.


### GetCompanyName

`func (o *MerchantKycCompanyInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MerchantKycCompanyInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *MerchantKycCompanyInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyNameEn

`func (o *MerchantKycCompanyInfo) GetCompanyNameEn() string`

GetCompanyNameEn returns the CompanyNameEn field if non-nil, zero value otherwise.

### GetCompanyNameEnOk

`func (o *MerchantKycCompanyInfo) GetCompanyNameEnOk() (*string, bool)`

GetCompanyNameEnOk returns a tuple with the CompanyNameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNameEn

`func (o *MerchantKycCompanyInfo) SetCompanyNameEn(v string)`

SetCompanyNameEn sets CompanyNameEn field to given value.


### GetEstablishDate

`func (o *MerchantKycCompanyInfo) GetEstablishDate() string`

GetEstablishDate returns the EstablishDate field if non-nil, zero value otherwise.

### GetEstablishDateOk

`func (o *MerchantKycCompanyInfo) GetEstablishDateOk() (*string, bool)`

GetEstablishDateOk returns a tuple with the EstablishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishDate

`func (o *MerchantKycCompanyInfo) SetEstablishDate(v string)`

SetEstablishDate sets EstablishDate field to given value.


### GetCommencementDate

`func (o *MerchantKycCompanyInfo) GetCommencementDate() string`

GetCommencementDate returns the CommencementDate field if non-nil, zero value otherwise.

### GetCommencementDateOk

`func (o *MerchantKycCompanyInfo) GetCommencementDateOk() (*string, bool)`

GetCommencementDateOk returns a tuple with the CommencementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommencementDate

`func (o *MerchantKycCompanyInfo) SetCommencementDate(v string)`

SetCommencementDate sets CommencementDate field to given value.


### GetValidPeriod

`func (o *MerchantKycCompanyInfo) GetValidPeriod() string`

GetValidPeriod returns the ValidPeriod field if non-nil, zero value otherwise.

### GetValidPeriodOk

`func (o *MerchantKycCompanyInfo) GetValidPeriodOk() (*string, bool)`

GetValidPeriodOk returns a tuple with the ValidPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPeriod

`func (o *MerchantKycCompanyInfo) SetValidPeriod(v string)`

SetValidPeriod sets ValidPeriod field to given value.


### GetRegisterAddress

`func (o *MerchantKycCompanyInfo) GetRegisterAddress() MerchantKycAddress`

GetRegisterAddress returns the RegisterAddress field if non-nil, zero value otherwise.

### GetRegisterAddressOk

`func (o *MerchantKycCompanyInfo) GetRegisterAddressOk() (*MerchantKycAddress, bool)`

GetRegisterAddressOk returns a tuple with the RegisterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterAddress

`func (o *MerchantKycCompanyInfo) SetRegisterAddress(v MerchantKycAddress)`

SetRegisterAddress sets RegisterAddress field to given value.


### GetLegalInfo

`func (o *MerchantKycCompanyInfo) GetLegalInfo() MerchantKycPersonInfo`

GetLegalInfo returns the LegalInfo field if non-nil, zero value otherwise.

### GetLegalInfoOk

`func (o *MerchantKycCompanyInfo) GetLegalInfoOk() (*MerchantKycPersonInfo, bool)`

GetLegalInfoOk returns a tuple with the LegalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalInfo

`func (o *MerchantKycCompanyInfo) SetLegalInfo(v MerchantKycPersonInfo)`

SetLegalInfo sets LegalInfo field to given value.


### GetUboInfos

`func (o *MerchantKycCompanyInfo) GetUboInfos() []MerchantKycPersonInfo`

GetUboInfos returns the UboInfos field if non-nil, zero value otherwise.

### GetUboInfosOk

`func (o *MerchantKycCompanyInfo) GetUboInfosOk() (*[]MerchantKycPersonInfo, bool)`

GetUboInfosOk returns a tuple with the UboInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUboInfos

`func (o *MerchantKycCompanyInfo) SetUboInfos(v []MerchantKycPersonInfo)`

SetUboInfos sets UboInfos field to given value.


### GetOnlineStoreUrl

`func (o *MerchantKycCompanyInfo) GetOnlineStoreUrl() string`

GetOnlineStoreUrl returns the OnlineStoreUrl field if non-nil, zero value otherwise.

### GetOnlineStoreUrlOk

`func (o *MerchantKycCompanyInfo) GetOnlineStoreUrlOk() (*string, bool)`

GetOnlineStoreUrlOk returns a tuple with the OnlineStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStoreUrl

`func (o *MerchantKycCompanyInfo) SetOnlineStoreUrl(v string)`

SetOnlineStoreUrl sets OnlineStoreUrl field to given value.

### HasOnlineStoreUrl

`func (o *MerchantKycCompanyInfo) HasOnlineStoreUrl() bool`

HasOnlineStoreUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



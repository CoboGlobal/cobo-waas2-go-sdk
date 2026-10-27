# FiatEntryKycCompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyType** | [**FiatEntryKycCompanyType**](FiatEntryKycCompanyType.md) |  | 
**Listed** | **bool** | Whether the company is listed. | 
**Attachments** | [**[]FiatEntryKycCompanyAttachment**](FiatEntryKycCompanyAttachment.md) | The attachments of the company. | 
**OperationAddress** | [**FiatEntryKycAddress**](FiatEntryKycAddress.md) |  | 
**IdentifyNo** | **string** | The company identification number. | 
**CompanyName** | **string** | The company name in local language. | 
**CompanyNameEn** | **string** | The company name in English. | 
**EstablishDate** | **string** | The establishment date of the company. | 
**CommencementDate** | **string** | The commencement date of the company. | 
**ValidPeriod** | **string** | The valid period of the company registration. | 
**RegisterAddress** | [**FiatEntryKycAddress**](FiatEntryKycAddress.md) |  | 
**LegalInfo** | [**FiatEntryKycPersonInfo**](FiatEntryKycPersonInfo.md) |  | 
**UboInfos** | [**[]FiatEntryKycPersonInfo**](FiatEntryKycPersonInfo.md) | The ultimate beneficial owner information. | 
**OnlineStoreUrl** | Pointer to **string** | The online store URL. Required when merchant type is B2B. | [optional] 

## Methods

### NewFiatEntryKycCompanyInfo

`func NewFiatEntryKycCompanyInfo(companyType FiatEntryKycCompanyType, listed bool, attachments []FiatEntryKycCompanyAttachment, operationAddress FiatEntryKycAddress, identifyNo string, companyName string, companyNameEn string, establishDate string, commencementDate string, validPeriod string, registerAddress FiatEntryKycAddress, legalInfo FiatEntryKycPersonInfo, uboInfos []FiatEntryKycPersonInfo, ) *FiatEntryKycCompanyInfo`

NewFiatEntryKycCompanyInfo instantiates a new FiatEntryKycCompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatEntryKycCompanyInfoWithDefaults

`func NewFiatEntryKycCompanyInfoWithDefaults() *FiatEntryKycCompanyInfo`

NewFiatEntryKycCompanyInfoWithDefaults instantiates a new FiatEntryKycCompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyType

`func (o *FiatEntryKycCompanyInfo) GetCompanyType() FiatEntryKycCompanyType`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *FiatEntryKycCompanyInfo) GetCompanyTypeOk() (*FiatEntryKycCompanyType, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *FiatEntryKycCompanyInfo) SetCompanyType(v FiatEntryKycCompanyType)`

SetCompanyType sets CompanyType field to given value.


### GetListed

`func (o *FiatEntryKycCompanyInfo) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *FiatEntryKycCompanyInfo) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *FiatEntryKycCompanyInfo) SetListed(v bool)`

SetListed sets Listed field to given value.


### GetAttachments

`func (o *FiatEntryKycCompanyInfo) GetAttachments() []FiatEntryKycCompanyAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *FiatEntryKycCompanyInfo) GetAttachmentsOk() (*[]FiatEntryKycCompanyAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *FiatEntryKycCompanyInfo) SetAttachments(v []FiatEntryKycCompanyAttachment)`

SetAttachments sets Attachments field to given value.


### GetOperationAddress

`func (o *FiatEntryKycCompanyInfo) GetOperationAddress() FiatEntryKycAddress`

GetOperationAddress returns the OperationAddress field if non-nil, zero value otherwise.

### GetOperationAddressOk

`func (o *FiatEntryKycCompanyInfo) GetOperationAddressOk() (*FiatEntryKycAddress, bool)`

GetOperationAddressOk returns a tuple with the OperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAddress

`func (o *FiatEntryKycCompanyInfo) SetOperationAddress(v FiatEntryKycAddress)`

SetOperationAddress sets OperationAddress field to given value.


### GetIdentifyNo

`func (o *FiatEntryKycCompanyInfo) GetIdentifyNo() string`

GetIdentifyNo returns the IdentifyNo field if non-nil, zero value otherwise.

### GetIdentifyNoOk

`func (o *FiatEntryKycCompanyInfo) GetIdentifyNoOk() (*string, bool)`

GetIdentifyNoOk returns a tuple with the IdentifyNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifyNo

`func (o *FiatEntryKycCompanyInfo) SetIdentifyNo(v string)`

SetIdentifyNo sets IdentifyNo field to given value.


### GetCompanyName

`func (o *FiatEntryKycCompanyInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *FiatEntryKycCompanyInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *FiatEntryKycCompanyInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyNameEn

`func (o *FiatEntryKycCompanyInfo) GetCompanyNameEn() string`

GetCompanyNameEn returns the CompanyNameEn field if non-nil, zero value otherwise.

### GetCompanyNameEnOk

`func (o *FiatEntryKycCompanyInfo) GetCompanyNameEnOk() (*string, bool)`

GetCompanyNameEnOk returns a tuple with the CompanyNameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNameEn

`func (o *FiatEntryKycCompanyInfo) SetCompanyNameEn(v string)`

SetCompanyNameEn sets CompanyNameEn field to given value.


### GetEstablishDate

`func (o *FiatEntryKycCompanyInfo) GetEstablishDate() string`

GetEstablishDate returns the EstablishDate field if non-nil, zero value otherwise.

### GetEstablishDateOk

`func (o *FiatEntryKycCompanyInfo) GetEstablishDateOk() (*string, bool)`

GetEstablishDateOk returns a tuple with the EstablishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishDate

`func (o *FiatEntryKycCompanyInfo) SetEstablishDate(v string)`

SetEstablishDate sets EstablishDate field to given value.


### GetCommencementDate

`func (o *FiatEntryKycCompanyInfo) GetCommencementDate() string`

GetCommencementDate returns the CommencementDate field if non-nil, zero value otherwise.

### GetCommencementDateOk

`func (o *FiatEntryKycCompanyInfo) GetCommencementDateOk() (*string, bool)`

GetCommencementDateOk returns a tuple with the CommencementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommencementDate

`func (o *FiatEntryKycCompanyInfo) SetCommencementDate(v string)`

SetCommencementDate sets CommencementDate field to given value.


### GetValidPeriod

`func (o *FiatEntryKycCompanyInfo) GetValidPeriod() string`

GetValidPeriod returns the ValidPeriod field if non-nil, zero value otherwise.

### GetValidPeriodOk

`func (o *FiatEntryKycCompanyInfo) GetValidPeriodOk() (*string, bool)`

GetValidPeriodOk returns a tuple with the ValidPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPeriod

`func (o *FiatEntryKycCompanyInfo) SetValidPeriod(v string)`

SetValidPeriod sets ValidPeriod field to given value.


### GetRegisterAddress

`func (o *FiatEntryKycCompanyInfo) GetRegisterAddress() FiatEntryKycAddress`

GetRegisterAddress returns the RegisterAddress field if non-nil, zero value otherwise.

### GetRegisterAddressOk

`func (o *FiatEntryKycCompanyInfo) GetRegisterAddressOk() (*FiatEntryKycAddress, bool)`

GetRegisterAddressOk returns a tuple with the RegisterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterAddress

`func (o *FiatEntryKycCompanyInfo) SetRegisterAddress(v FiatEntryKycAddress)`

SetRegisterAddress sets RegisterAddress field to given value.


### GetLegalInfo

`func (o *FiatEntryKycCompanyInfo) GetLegalInfo() FiatEntryKycPersonInfo`

GetLegalInfo returns the LegalInfo field if non-nil, zero value otherwise.

### GetLegalInfoOk

`func (o *FiatEntryKycCompanyInfo) GetLegalInfoOk() (*FiatEntryKycPersonInfo, bool)`

GetLegalInfoOk returns a tuple with the LegalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalInfo

`func (o *FiatEntryKycCompanyInfo) SetLegalInfo(v FiatEntryKycPersonInfo)`

SetLegalInfo sets LegalInfo field to given value.


### GetUboInfos

`func (o *FiatEntryKycCompanyInfo) GetUboInfos() []FiatEntryKycPersonInfo`

GetUboInfos returns the UboInfos field if non-nil, zero value otherwise.

### GetUboInfosOk

`func (o *FiatEntryKycCompanyInfo) GetUboInfosOk() (*[]FiatEntryKycPersonInfo, bool)`

GetUboInfosOk returns a tuple with the UboInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUboInfos

`func (o *FiatEntryKycCompanyInfo) SetUboInfos(v []FiatEntryKycPersonInfo)`

SetUboInfos sets UboInfos field to given value.


### GetOnlineStoreUrl

`func (o *FiatEntryKycCompanyInfo) GetOnlineStoreUrl() string`

GetOnlineStoreUrl returns the OnlineStoreUrl field if non-nil, zero value otherwise.

### GetOnlineStoreUrlOk

`func (o *FiatEntryKycCompanyInfo) GetOnlineStoreUrlOk() (*string, bool)`

GetOnlineStoreUrlOk returns a tuple with the OnlineStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStoreUrl

`func (o *FiatEntryKycCompanyInfo) SetOnlineStoreUrl(v string)`

SetOnlineStoreUrl sets OnlineStoreUrl field to given value.

### HasOnlineStoreUrl

`func (o *FiatEntryKycCompanyInfo) HasOnlineStoreUrl() bool`

HasOnlineStoreUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MerchantKycPersonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name or title of an identification document. | 
**NameEn** | **string** | The English-language equivalent of the identification document&#39;s name. | 
**IdNumber** | **string** | The unique identification number associated with the identification document. | 
**DateOfBirth** | **string** | The date of birth of the individual, usually in the format YYYYMMDD.  | 
**IssueDate** | **string** | The issue date refers to the date when a document, such as an identification card or passport, was officially issued or granted, usually in the format YYYYMMDD.  | 
**ExpirationDate** | **string** | The expiration date refers to the date when a document, such as an identification card or passport, is no longer valid or legally usable, usually in the format YYYYMMDD.  | 
**Attachments** | [**[]MerchantKycPersonAttachment**](MerchantKycPersonAttachment.md) | Additional files or documents associated with a message or record to provide extra information or context.  | 
**ResidentialAddress** | [**MerchantKycAddress**](MerchantKycAddress.md) |  | 

## Methods

### NewMerchantKycPersonInfo

`func NewMerchantKycPersonInfo(name string, nameEn string, idNumber string, dateOfBirth string, issueDate string, expirationDate string, attachments []MerchantKycPersonAttachment, residentialAddress MerchantKycAddress, ) *MerchantKycPersonInfo`

NewMerchantKycPersonInfo instantiates a new MerchantKycPersonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantKycPersonInfoWithDefaults

`func NewMerchantKycPersonInfoWithDefaults() *MerchantKycPersonInfo`

NewMerchantKycPersonInfoWithDefaults instantiates a new MerchantKycPersonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantKycPersonInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantKycPersonInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantKycPersonInfo) SetName(v string)`

SetName sets Name field to given value.


### GetNameEn

`func (o *MerchantKycPersonInfo) GetNameEn() string`

GetNameEn returns the NameEn field if non-nil, zero value otherwise.

### GetNameEnOk

`func (o *MerchantKycPersonInfo) GetNameEnOk() (*string, bool)`

GetNameEnOk returns a tuple with the NameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameEn

`func (o *MerchantKycPersonInfo) SetNameEn(v string)`

SetNameEn sets NameEn field to given value.


### GetIdNumber

`func (o *MerchantKycPersonInfo) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *MerchantKycPersonInfo) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *MerchantKycPersonInfo) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.


### GetDateOfBirth

`func (o *MerchantKycPersonInfo) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *MerchantKycPersonInfo) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *MerchantKycPersonInfo) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetIssueDate

`func (o *MerchantKycPersonInfo) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *MerchantKycPersonInfo) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *MerchantKycPersonInfo) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetExpirationDate

`func (o *MerchantKycPersonInfo) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *MerchantKycPersonInfo) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *MerchantKycPersonInfo) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetAttachments

`func (o *MerchantKycPersonInfo) GetAttachments() []MerchantKycPersonAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MerchantKycPersonInfo) GetAttachmentsOk() (*[]MerchantKycPersonAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MerchantKycPersonInfo) SetAttachments(v []MerchantKycPersonAttachment)`

SetAttachments sets Attachments field to given value.


### GetResidentialAddress

`func (o *MerchantKycPersonInfo) GetResidentialAddress() MerchantKycAddress`

GetResidentialAddress returns the ResidentialAddress field if non-nil, zero value otherwise.

### GetResidentialAddressOk

`func (o *MerchantKycPersonInfo) GetResidentialAddressOk() (*MerchantKycAddress, bool)`

GetResidentialAddressOk returns a tuple with the ResidentialAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialAddress

`func (o *MerchantKycPersonInfo) SetResidentialAddress(v MerchantKycAddress)`

SetResidentialAddress sets ResidentialAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



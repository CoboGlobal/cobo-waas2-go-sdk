# MerchantKycCompanyAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The AWS file link of the uploaded file, which you can retrieve by calling [Upload file](https://www.cobo.com/developers/v2/api-references/payment/upload-file).  | 
**FileType** | [**MerchantKycCompanyAttachmentFileType**](MerchantKycCompanyAttachmentFileType.md) |  | 

## Methods

### NewMerchantKycCompanyAttachment

`func NewMerchantKycCompanyAttachment(fileId string, fileType MerchantKycCompanyAttachmentFileType, ) *MerchantKycCompanyAttachment`

NewMerchantKycCompanyAttachment instantiates a new MerchantKycCompanyAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantKycCompanyAttachmentWithDefaults

`func NewMerchantKycCompanyAttachmentWithDefaults() *MerchantKycCompanyAttachment`

NewMerchantKycCompanyAttachmentWithDefaults instantiates a new MerchantKycCompanyAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *MerchantKycCompanyAttachment) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MerchantKycCompanyAttachment) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MerchantKycCompanyAttachment) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileType

`func (o *MerchantKycCompanyAttachment) GetFileType() MerchantKycCompanyAttachmentFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *MerchantKycCompanyAttachment) GetFileTypeOk() (*MerchantKycCompanyAttachmentFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *MerchantKycCompanyAttachment) SetFileType(v MerchantKycCompanyAttachmentFileType)`

SetFileType sets FileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



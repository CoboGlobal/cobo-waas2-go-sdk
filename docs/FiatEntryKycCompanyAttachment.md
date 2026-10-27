# FiatEntryKycCompanyAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The file ID of the uploaded file, which you can retrieve by calling [Upload file v2](https://www.cobo.com/developers/v2/api-references/payment/upload-file-v2).  | 
**FileType** | [**FiatEntryKycCompanyAttachmentFileType**](FiatEntryKycCompanyAttachmentFileType.md) |  | 

## Methods

### NewFiatEntryKycCompanyAttachment

`func NewFiatEntryKycCompanyAttachment(fileId string, fileType FiatEntryKycCompanyAttachmentFileType, ) *FiatEntryKycCompanyAttachment`

NewFiatEntryKycCompanyAttachment instantiates a new FiatEntryKycCompanyAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatEntryKycCompanyAttachmentWithDefaults

`func NewFiatEntryKycCompanyAttachmentWithDefaults() *FiatEntryKycCompanyAttachment`

NewFiatEntryKycCompanyAttachmentWithDefaults instantiates a new FiatEntryKycCompanyAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *FiatEntryKycCompanyAttachment) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FiatEntryKycCompanyAttachment) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FiatEntryKycCompanyAttachment) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileType

`func (o *FiatEntryKycCompanyAttachment) GetFileType() FiatEntryKycCompanyAttachmentFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FiatEntryKycCompanyAttachment) GetFileTypeOk() (*FiatEntryKycCompanyAttachmentFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FiatEntryKycCompanyAttachment) SetFileType(v FiatEntryKycCompanyAttachmentFileType)`

SetFileType sets FileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



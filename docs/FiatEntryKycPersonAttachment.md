# FiatEntryKycPersonAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The AWS file link of the uploaded file, which you can retrieve by calling [Upload file v2](https://www.cobo.com/developers/v2/api-references/payment/upload-file-v2).  | 
**FileType** | [**FiatEntryKycPersonAttachmentFileType**](FiatEntryKycPersonAttachmentFileType.md) |  | 

## Methods

### NewFiatEntryKycPersonAttachment

`func NewFiatEntryKycPersonAttachment(fileId string, fileType FiatEntryKycPersonAttachmentFileType, ) *FiatEntryKycPersonAttachment`

NewFiatEntryKycPersonAttachment instantiates a new FiatEntryKycPersonAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatEntryKycPersonAttachmentWithDefaults

`func NewFiatEntryKycPersonAttachmentWithDefaults() *FiatEntryKycPersonAttachment`

NewFiatEntryKycPersonAttachmentWithDefaults instantiates a new FiatEntryKycPersonAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *FiatEntryKycPersonAttachment) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FiatEntryKycPersonAttachment) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FiatEntryKycPersonAttachment) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileType

`func (o *FiatEntryKycPersonAttachment) GetFileType() FiatEntryKycPersonAttachmentFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FiatEntryKycPersonAttachment) GetFileTypeOk() (*FiatEntryKycPersonAttachmentFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FiatEntryKycPersonAttachment) SetFileType(v FiatEntryKycPersonAttachmentFileType)`

SetFileType sets FileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MerchantKycPersonAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The AWS file link of the uploaded file, which you can retrieve by calling [Upload file](https://www.cobo.com/developers/v2/api-references/payment/upload-file).  | 
**FileType** | [**MerchantKycPersonAttachmentFileType**](MerchantKycPersonAttachmentFileType.md) |  | 

## Methods

### NewMerchantKycPersonAttachment

`func NewMerchantKycPersonAttachment(fileId string, fileType MerchantKycPersonAttachmentFileType, ) *MerchantKycPersonAttachment`

NewMerchantKycPersonAttachment instantiates a new MerchantKycPersonAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantKycPersonAttachmentWithDefaults

`func NewMerchantKycPersonAttachmentWithDefaults() *MerchantKycPersonAttachment`

NewMerchantKycPersonAttachmentWithDefaults instantiates a new MerchantKycPersonAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *MerchantKycPersonAttachment) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MerchantKycPersonAttachment) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MerchantKycPersonAttachment) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileType

`func (o *MerchantKycPersonAttachment) GetFileType() MerchantKycPersonAttachmentFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *MerchantKycPersonAttachment) GetFileTypeOk() (*MerchantKycPersonAttachmentFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *MerchantKycPersonAttachment) SetFileType(v MerchantKycPersonAttachmentFileType)`

SetFileType sets FileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



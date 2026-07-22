# PaymentUploadedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | The AWS file link of the uploaded file. | 
**ExpiredTimestamp** | **int64** | The time when the uploaded file link expires, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewPaymentUploadedFile

`func NewPaymentUploadedFile(fileId string, expiredTimestamp int64, ) *PaymentUploadedFile`

NewPaymentUploadedFile instantiates a new PaymentUploadedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentUploadedFileWithDefaults

`func NewPaymentUploadedFileWithDefaults() *PaymentUploadedFile`

NewPaymentUploadedFileWithDefaults instantiates a new PaymentUploadedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *PaymentUploadedFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PaymentUploadedFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PaymentUploadedFile) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetExpiredTimestamp

`func (o *PaymentUploadedFile) GetExpiredTimestamp() int64`

GetExpiredTimestamp returns the ExpiredTimestamp field if non-nil, zero value otherwise.

### GetExpiredTimestampOk

`func (o *PaymentUploadedFile) GetExpiredTimestampOk() (*int64, bool)`

GetExpiredTimestampOk returns a tuple with the ExpiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimestamp

`func (o *PaymentUploadedFile) SetExpiredTimestamp(v int64)`

SetExpiredTimestamp sets ExpiredTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



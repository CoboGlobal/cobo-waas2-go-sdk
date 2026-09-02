# PaymentUploadFileV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** | The original file name, including the file extension. | 
**FileContent** | **string** | The file content, encoded in Base64. | 

## Methods

### NewPaymentUploadFileV2

`func NewPaymentUploadFileV2(fileName string, fileContent string, ) *PaymentUploadFileV2`

NewPaymentUploadFileV2 instantiates a new PaymentUploadFileV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentUploadFileV2WithDefaults

`func NewPaymentUploadFileV2WithDefaults() *PaymentUploadFileV2`

NewPaymentUploadFileV2WithDefaults instantiates a new PaymentUploadFileV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *PaymentUploadFileV2) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *PaymentUploadFileV2) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *PaymentUploadFileV2) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileContent

`func (o *PaymentUploadFileV2) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *PaymentUploadFileV2) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *PaymentUploadFileV2) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TransactionUserApprovalDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name the user who audited this message. | [optional] 
**Email** | Pointer to **string** | The email the user who audited this message. | [optional] 
**Pubkey** | Pointer to **string** | The Cobo Guard public key of the user who audited this message. | [optional] 
**Result** | Pointer to [**TransactionApprovalResult**](TransactionApprovalResult.md) |  | [optional] 
**Signature** | Pointer to **string** | The signature of the audited message. | [optional] 
**Language** | Pointer to **string** | The language of the audited message. | [optional] 
**MessageVersion** | Pointer to **string** | The version of the audited message. | [optional] 
**Message** | Pointer to **string** | The audited message. | [optional] 
**ExtraMessage** | Pointer to **string** | The extra audited message. | [optional] 

## Methods

### NewTransactionUserApprovalDetail

`func NewTransactionUserApprovalDetail() *TransactionUserApprovalDetail`

NewTransactionUserApprovalDetail instantiates a new TransactionUserApprovalDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionUserApprovalDetailWithDefaults

`func NewTransactionUserApprovalDetailWithDefaults() *TransactionUserApprovalDetail`

NewTransactionUserApprovalDetailWithDefaults instantiates a new TransactionUserApprovalDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransactionUserApprovalDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionUserApprovalDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionUserApprovalDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionUserApprovalDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *TransactionUserApprovalDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TransactionUserApprovalDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TransactionUserApprovalDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TransactionUserApprovalDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPubkey

`func (o *TransactionUserApprovalDetail) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *TransactionUserApprovalDetail) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *TransactionUserApprovalDetail) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *TransactionUserApprovalDetail) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetResult

`func (o *TransactionUserApprovalDetail) GetResult() TransactionApprovalResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TransactionUserApprovalDetail) GetResultOk() (*TransactionApprovalResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TransactionUserApprovalDetail) SetResult(v TransactionApprovalResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *TransactionUserApprovalDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSignature

`func (o *TransactionUserApprovalDetail) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TransactionUserApprovalDetail) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TransactionUserApprovalDetail) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *TransactionUserApprovalDetail) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetLanguage

`func (o *TransactionUserApprovalDetail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TransactionUserApprovalDetail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TransactionUserApprovalDetail) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TransactionUserApprovalDetail) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMessageVersion

`func (o *TransactionUserApprovalDetail) GetMessageVersion() string`

GetMessageVersion returns the MessageVersion field if non-nil, zero value otherwise.

### GetMessageVersionOk

`func (o *TransactionUserApprovalDetail) GetMessageVersionOk() (*string, bool)`

GetMessageVersionOk returns a tuple with the MessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVersion

`func (o *TransactionUserApprovalDetail) SetMessageVersion(v string)`

SetMessageVersion sets MessageVersion field to given value.

### HasMessageVersion

`func (o *TransactionUserApprovalDetail) HasMessageVersion() bool`

HasMessageVersion returns a boolean if a field has been set.

### GetMessage

`func (o *TransactionUserApprovalDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionUserApprovalDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionUserApprovalDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionUserApprovalDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExtraMessage

`func (o *TransactionUserApprovalDetail) GetExtraMessage() string`

GetExtraMessage returns the ExtraMessage field if non-nil, zero value otherwise.

### GetExtraMessageOk

`func (o *TransactionUserApprovalDetail) GetExtraMessageOk() (*string, bool)`

GetExtraMessageOk returns a tuple with the ExtraMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraMessage

`func (o *TransactionUserApprovalDetail) SetExtraMessage(v string)`

SetExtraMessage sets ExtraMessage field to given value.

### HasExtraMessage

`func (o *TransactionUserApprovalDetail) HasExtraMessage() bool`

HasExtraMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



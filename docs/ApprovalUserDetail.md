# ApprovalUserDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the user who approved the transaction. | [optional] 
**Email** | Pointer to **string** | The email address of the user who approved the transaction. | [optional] 
**Pubkey** | Pointer to **string** | The public key of the user who approved the transaction. | [optional] 
**Signature** | Pointer to **string** | The signature of the transaction approval. | [optional] 
**StatementUuid** | Pointer to **string** | The UUID of the statement associated with the transaction approval. | [optional] 
**Result** | Pointer to [**ApprovalResult**](ApprovalResult.md) |  | [optional] 
**ApprovalResultCode** | Pointer to **int32** | The integer value representing the result of the approval. | [optional] 
**CreatedTime** | Pointer to **int32** | The timestamp when the approval was created. | [optional] 
**TemplateVersion** | Pointer to **string** | The version of the template used for the transaction approval. | [optional] 
**HeaderTitle** | Pointer to **string** | The title of the header for the transaction approval. | [optional] 
**IsForSign** | Pointer to **bool** | Indicates whether the approval is for signing. | [optional] 
**ShowInfo** | Pointer to **string** | Additional information to show for the transaction approval. | [optional] 
**Language** | Pointer to **string** | The language used for the transaction approval. | [optional] 
**MessageVersion** | Pointer to **string** | The version of the message format used for the transaction approval. | [optional] 
**Message** | Pointer to **string** | The message associated with the transaction approval. | [optional] 
**ExtraMessage** | Pointer to **string** | Any additional message or information related to the transaction approval. | [optional] 

## Methods

### NewApprovalUserDetail

`func NewApprovalUserDetail() *ApprovalUserDetail`

NewApprovalUserDetail instantiates a new ApprovalUserDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalUserDetailWithDefaults

`func NewApprovalUserDetailWithDefaults() *ApprovalUserDetail`

NewApprovalUserDetailWithDefaults instantiates a new ApprovalUserDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApprovalUserDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalUserDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalUserDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalUserDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ApprovalUserDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApprovalUserDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApprovalUserDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApprovalUserDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPubkey

`func (o *ApprovalUserDetail) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ApprovalUserDetail) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ApprovalUserDetail) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *ApprovalUserDetail) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetSignature

`func (o *ApprovalUserDetail) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ApprovalUserDetail) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ApprovalUserDetail) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ApprovalUserDetail) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetStatementUuid

`func (o *ApprovalUserDetail) GetStatementUuid() string`

GetStatementUuid returns the StatementUuid field if non-nil, zero value otherwise.

### GetStatementUuidOk

`func (o *ApprovalUserDetail) GetStatementUuidOk() (*string, bool)`

GetStatementUuidOk returns a tuple with the StatementUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementUuid

`func (o *ApprovalUserDetail) SetStatementUuid(v string)`

SetStatementUuid sets StatementUuid field to given value.

### HasStatementUuid

`func (o *ApprovalUserDetail) HasStatementUuid() bool`

HasStatementUuid returns a boolean if a field has been set.

### GetResult

`func (o *ApprovalUserDetail) GetResult() ApprovalResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApprovalUserDetail) GetResultOk() (*ApprovalResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApprovalUserDetail) SetResult(v ApprovalResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ApprovalUserDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetApprovalResultCode

`func (o *ApprovalUserDetail) GetApprovalResultCode() int32`

GetApprovalResultCode returns the ApprovalResultCode field if non-nil, zero value otherwise.

### GetApprovalResultCodeOk

`func (o *ApprovalUserDetail) GetApprovalResultCodeOk() (*int32, bool)`

GetApprovalResultCodeOk returns a tuple with the ApprovalResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalResultCode

`func (o *ApprovalUserDetail) SetApprovalResultCode(v int32)`

SetApprovalResultCode sets ApprovalResultCode field to given value.

### HasApprovalResultCode

`func (o *ApprovalUserDetail) HasApprovalResultCode() bool`

HasApprovalResultCode returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ApprovalUserDetail) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ApprovalUserDetail) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ApprovalUserDetail) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ApprovalUserDetail) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *ApprovalUserDetail) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *ApprovalUserDetail) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *ApprovalUserDetail) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *ApprovalUserDetail) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetHeaderTitle

`func (o *ApprovalUserDetail) GetHeaderTitle() string`

GetHeaderTitle returns the HeaderTitle field if non-nil, zero value otherwise.

### GetHeaderTitleOk

`func (o *ApprovalUserDetail) GetHeaderTitleOk() (*string, bool)`

GetHeaderTitleOk returns a tuple with the HeaderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTitle

`func (o *ApprovalUserDetail) SetHeaderTitle(v string)`

SetHeaderTitle sets HeaderTitle field to given value.

### HasHeaderTitle

`func (o *ApprovalUserDetail) HasHeaderTitle() bool`

HasHeaderTitle returns a boolean if a field has been set.

### GetIsForSign

`func (o *ApprovalUserDetail) GetIsForSign() bool`

GetIsForSign returns the IsForSign field if non-nil, zero value otherwise.

### GetIsForSignOk

`func (o *ApprovalUserDetail) GetIsForSignOk() (*bool, bool)`

GetIsForSignOk returns a tuple with the IsForSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForSign

`func (o *ApprovalUserDetail) SetIsForSign(v bool)`

SetIsForSign sets IsForSign field to given value.

### HasIsForSign

`func (o *ApprovalUserDetail) HasIsForSign() bool`

HasIsForSign returns a boolean if a field has been set.

### GetShowInfo

`func (o *ApprovalUserDetail) GetShowInfo() string`

GetShowInfo returns the ShowInfo field if non-nil, zero value otherwise.

### GetShowInfoOk

`func (o *ApprovalUserDetail) GetShowInfoOk() (*string, bool)`

GetShowInfoOk returns a tuple with the ShowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInfo

`func (o *ApprovalUserDetail) SetShowInfo(v string)`

SetShowInfo sets ShowInfo field to given value.

### HasShowInfo

`func (o *ApprovalUserDetail) HasShowInfo() bool`

HasShowInfo returns a boolean if a field has been set.

### GetLanguage

`func (o *ApprovalUserDetail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ApprovalUserDetail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ApprovalUserDetail) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ApprovalUserDetail) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMessageVersion

`func (o *ApprovalUserDetail) GetMessageVersion() string`

GetMessageVersion returns the MessageVersion field if non-nil, zero value otherwise.

### GetMessageVersionOk

`func (o *ApprovalUserDetail) GetMessageVersionOk() (*string, bool)`

GetMessageVersionOk returns a tuple with the MessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVersion

`func (o *ApprovalUserDetail) SetMessageVersion(v string)`

SetMessageVersion sets MessageVersion field to given value.

### HasMessageVersion

`func (o *ApprovalUserDetail) HasMessageVersion() bool`

HasMessageVersion returns a boolean if a field has been set.

### GetMessage

`func (o *ApprovalUserDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApprovalUserDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApprovalUserDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApprovalUserDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExtraMessage

`func (o *ApprovalUserDetail) GetExtraMessage() string`

GetExtraMessage returns the ExtraMessage field if non-nil, zero value otherwise.

### GetExtraMessageOk

`func (o *ApprovalUserDetail) GetExtraMessageOk() (*string, bool)`

GetExtraMessageOk returns a tuple with the ExtraMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraMessage

`func (o *ApprovalUserDetail) SetExtraMessage(v string)`

SetExtraMessage sets ExtraMessage field to given value.

### HasExtraMessage

`func (o *ApprovalUserDetail) HasExtraMessage() bool`

HasExtraMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



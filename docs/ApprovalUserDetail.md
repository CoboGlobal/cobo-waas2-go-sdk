# ApprovalUserDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserEmail** | Pointer to **string** | The email address of the user who approved the transaction. | [optional] 
**Pubkey** | Pointer to **string** | The public key of the user who approved the transaction. | [optional] 
**Signature** | Pointer to **string** | The signature of the transaction approval. | [optional] 
**StatementUuid** | Pointer to **string** | The UUID of the statement associated with the transaction approval. | [optional] 
**Result** | Pointer to [**ApprovalResult**](ApprovalResult.md) |  | [optional] 
**CreatedTime** | Pointer to **int32** | The timestamp when the approval was created. | [optional] 
**TemplateVersion** | Pointer to **string** | The version of the template used for the transaction approval. | [optional] 
**HeaderTitle** | Pointer to **string** | The title of the header for the transaction approval. | [optional] 
**IsForSign** | Pointer to **bool** | Indicates whether the approval is for signing. | [optional] 
**ShowInfo** | Pointer to [**ApprovalShowInfo**](ApprovalShowInfo.md) |  | [optional] 

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

### GetUserEmail

`func (o *ApprovalUserDetail) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ApprovalUserDetail) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ApprovalUserDetail) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *ApprovalUserDetail) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

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

`func (o *ApprovalUserDetail) GetShowInfo() ApprovalShowInfo`

GetShowInfo returns the ShowInfo field if non-nil, zero value otherwise.

### GetShowInfoOk

`func (o *ApprovalUserDetail) GetShowInfoOk() (*ApprovalShowInfo, bool)`

GetShowInfoOk returns a tuple with the ShowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInfo

`func (o *ApprovalUserDetail) SetShowInfo(v ApprovalShowInfo)`

SetShowInfo sets ShowInfo field to given value.

### HasShowInfo

`func (o *ApprovalUserDetail) HasShowInfo() bool`

HasShowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TransactionApprovalDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spender** | Pointer to [**TransactionRoleApprovalDetail**](TransactionRoleApprovalDetail.md) |  | [optional] 
**Approver** | Pointer to [**TransactionRoleApprovalDetail**](TransactionRoleApprovalDetail.md) |  | [optional] 
**AddressOwner** | Pointer to [**TransactionRoleApprovalDetail**](TransactionRoleApprovalDetail.md) |  | [optional] 

## Methods

### NewTransactionApprovalDetail

`func NewTransactionApprovalDetail() *TransactionApprovalDetail`

NewTransactionApprovalDetail instantiates a new TransactionApprovalDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionApprovalDetailWithDefaults

`func NewTransactionApprovalDetailWithDefaults() *TransactionApprovalDetail`

NewTransactionApprovalDetailWithDefaults instantiates a new TransactionApprovalDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpender

`func (o *TransactionApprovalDetail) GetSpender() TransactionRoleApprovalDetail`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *TransactionApprovalDetail) GetSpenderOk() (*TransactionRoleApprovalDetail, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *TransactionApprovalDetail) SetSpender(v TransactionRoleApprovalDetail)`

SetSpender sets Spender field to given value.

### HasSpender

`func (o *TransactionApprovalDetail) HasSpender() bool`

HasSpender returns a boolean if a field has been set.

### GetApprover

`func (o *TransactionApprovalDetail) GetApprover() TransactionRoleApprovalDetail`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *TransactionApprovalDetail) GetApproverOk() (*TransactionRoleApprovalDetail, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *TransactionApprovalDetail) SetApprover(v TransactionRoleApprovalDetail)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *TransactionApprovalDetail) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### GetAddressOwner

`func (o *TransactionApprovalDetail) GetAddressOwner() TransactionRoleApprovalDetail`

GetAddressOwner returns the AddressOwner field if non-nil, zero value otherwise.

### GetAddressOwnerOk

`func (o *TransactionApprovalDetail) GetAddressOwnerOk() (*TransactionRoleApprovalDetail, bool)`

GetAddressOwnerOk returns a tuple with the AddressOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOwner

`func (o *TransactionApprovalDetail) SetAddressOwner(v TransactionRoleApprovalDetail)`

SetAddressOwner sets AddressOwner field to given value.

### HasAddressOwner

`func (o *TransactionApprovalDetail) HasAddressOwner() bool`

HasAddressOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



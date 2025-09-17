# ApprovalDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** | Transaction ID. | [optional] 
**CoboId** | Pointer to **string** | Cobo ID used to track a transaction. | [optional] 
**RequestId** | Pointer to **string** | Request ID used to track a transaction request. | [optional] 
**AddressOwner** | Pointer to [**RoleDetail**](RoleDetail.md) |  | [optional] 
**Spender** | Pointer to [**RoleDetail**](RoleDetail.md) |  | [optional] 
**Approver** | Pointer to [**RoleDetail**](RoleDetail.md) |  | [optional] 

## Methods

### NewApprovalDetail

`func NewApprovalDetail() *ApprovalDetail`

NewApprovalDetail instantiates a new ApprovalDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalDetailWithDefaults

`func NewApprovalDetailWithDefaults() *ApprovalDetail`

NewApprovalDetailWithDefaults instantiates a new ApprovalDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *ApprovalDetail) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ApprovalDetail) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ApprovalDetail) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ApprovalDetail) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetCoboId

`func (o *ApprovalDetail) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *ApprovalDetail) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *ApprovalDetail) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *ApprovalDetail) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### GetRequestId

`func (o *ApprovalDetail) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ApprovalDetail) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ApprovalDetail) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ApprovalDetail) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetAddressOwner

`func (o *ApprovalDetail) GetAddressOwner() RoleDetail`

GetAddressOwner returns the AddressOwner field if non-nil, zero value otherwise.

### GetAddressOwnerOk

`func (o *ApprovalDetail) GetAddressOwnerOk() (*RoleDetail, bool)`

GetAddressOwnerOk returns a tuple with the AddressOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOwner

`func (o *ApprovalDetail) SetAddressOwner(v RoleDetail)`

SetAddressOwner sets AddressOwner field to given value.

### HasAddressOwner

`func (o *ApprovalDetail) HasAddressOwner() bool`

HasAddressOwner returns a boolean if a field has been set.

### GetSpender

`func (o *ApprovalDetail) GetSpender() RoleDetail`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *ApprovalDetail) GetSpenderOk() (*RoleDetail, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *ApprovalDetail) SetSpender(v RoleDetail)`

SetSpender sets Spender field to given value.

### HasSpender

`func (o *ApprovalDetail) HasSpender() bool`

HasSpender returns a boolean if a field has been set.

### GetApprover

`func (o *ApprovalDetail) GetApprover() RoleDetail`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *ApprovalDetail) GetApproverOk() (*RoleDetail, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *ApprovalDetail) SetApprover(v RoleDetail)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *ApprovalDetail) HasApprover() bool`

HasApprover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



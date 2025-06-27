# ApprovalDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** | The transaction ID. | [optional] 
**CoboId** | Pointer to **string** | The Cobo ID, which can be used to track a transaction. | [optional] 
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | [optional] 
**Result** | Pointer to [**ApprovalTransactionResult**](ApprovalTransactionResult.md) |  | [optional] 
**Threshold** | Pointer to **int32** | The threshold for the transaction approval. | [optional] 
**UserDetails** | Pointer to [**[]ApprovalUserDetail**](ApprovalUserDetail.md) |  | [optional] 

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

### GetResult

`func (o *ApprovalDetail) GetResult() ApprovalTransactionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApprovalDetail) GetResultOk() (*ApprovalTransactionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApprovalDetail) SetResult(v ApprovalTransactionResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ApprovalDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetThreshold

`func (o *ApprovalDetail) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApprovalDetail) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApprovalDetail) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ApprovalDetail) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUserDetails

`func (o *ApprovalDetail) GetUserDetails() []ApprovalUserDetail`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *ApprovalDetail) GetUserDetailsOk() (*[]ApprovalUserDetail, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *ApprovalDetail) SetUserDetails(v []ApprovalUserDetail)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *ApprovalDetail) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



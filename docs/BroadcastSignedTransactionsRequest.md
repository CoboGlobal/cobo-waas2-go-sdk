# BroadcastSignedTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionIds** | Pointer to **[]string** | The transaction IDs of the signed transactions to be broadcast. You can retrieve the transactions corresponding to a staking activity by calling [Get staking activity details](https://www.cobo.com/developers/v2/api-references/stakings/get-staking-activity-details). | [optional] 

## Methods

### NewBroadcastSignedTransactionsRequest

`func NewBroadcastSignedTransactionsRequest() *BroadcastSignedTransactionsRequest`

NewBroadcastSignedTransactionsRequest instantiates a new BroadcastSignedTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastSignedTransactionsRequestWithDefaults

`func NewBroadcastSignedTransactionsRequestWithDefaults() *BroadcastSignedTransactionsRequest`

NewBroadcastSignedTransactionsRequestWithDefaults instantiates a new BroadcastSignedTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionIds

`func (o *BroadcastSignedTransactionsRequest) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *BroadcastSignedTransactionsRequest) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *BroadcastSignedTransactionsRequest) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *BroadcastSignedTransactionsRequest) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



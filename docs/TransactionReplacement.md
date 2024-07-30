# TransactionReplacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplacedByType** | Pointer to [**ReplaceType**](ReplaceType.md) |  | [optional] 
**ReplacedByTransactionId** | Pointer to **string** | The ID of the transaction that this transaction was replaced by. | [optional] 
**ReplacedByTransactionHash** | Pointer to **string** | The hash of the transaction that this transaction was replaced by. | [optional] 
**ReplacedType** | Pointer to [**ReplaceType**](ReplaceType.md) |  | [optional] 
**ReplacedTransactionId** | Pointer to **string** | The ID of the transaction that this transaction replaced. | [optional] 
**ReplacedTransactionHash** | Pointer to **string** | The hash of the transaction that this transaction replaced. | [optional] 

## Methods

### NewTransactionReplacement

`func NewTransactionReplacement() *TransactionReplacement`

NewTransactionReplacement instantiates a new TransactionReplacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReplacementWithDefaults

`func NewTransactionReplacementWithDefaults() *TransactionReplacement`

NewTransactionReplacementWithDefaults instantiates a new TransactionReplacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplacedByType

`func (o *TransactionReplacement) GetReplacedByType() ReplaceType`

GetReplacedByType returns the ReplacedByType field if non-nil, zero value otherwise.

### GetReplacedByTypeOk

`func (o *TransactionReplacement) GetReplacedByTypeOk() (*ReplaceType, bool)`

GetReplacedByTypeOk returns a tuple with the ReplacedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedByType

`func (o *TransactionReplacement) SetReplacedByType(v ReplaceType)`

SetReplacedByType sets ReplacedByType field to given value.

### HasReplacedByType

`func (o *TransactionReplacement) HasReplacedByType() bool`

HasReplacedByType returns a boolean if a field has been set.

### GetReplacedByTransactionId

`func (o *TransactionReplacement) GetReplacedByTransactionId() string`

GetReplacedByTransactionId returns the ReplacedByTransactionId field if non-nil, zero value otherwise.

### GetReplacedByTransactionIdOk

`func (o *TransactionReplacement) GetReplacedByTransactionIdOk() (*string, bool)`

GetReplacedByTransactionIdOk returns a tuple with the ReplacedByTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedByTransactionId

`func (o *TransactionReplacement) SetReplacedByTransactionId(v string)`

SetReplacedByTransactionId sets ReplacedByTransactionId field to given value.

### HasReplacedByTransactionId

`func (o *TransactionReplacement) HasReplacedByTransactionId() bool`

HasReplacedByTransactionId returns a boolean if a field has been set.

### GetReplacedByTransactionHash

`func (o *TransactionReplacement) GetReplacedByTransactionHash() string`

GetReplacedByTransactionHash returns the ReplacedByTransactionHash field if non-nil, zero value otherwise.

### GetReplacedByTransactionHashOk

`func (o *TransactionReplacement) GetReplacedByTransactionHashOk() (*string, bool)`

GetReplacedByTransactionHashOk returns a tuple with the ReplacedByTransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedByTransactionHash

`func (o *TransactionReplacement) SetReplacedByTransactionHash(v string)`

SetReplacedByTransactionHash sets ReplacedByTransactionHash field to given value.

### HasReplacedByTransactionHash

`func (o *TransactionReplacement) HasReplacedByTransactionHash() bool`

HasReplacedByTransactionHash returns a boolean if a field has been set.

### GetReplacedType

`func (o *TransactionReplacement) GetReplacedType() ReplaceType`

GetReplacedType returns the ReplacedType field if non-nil, zero value otherwise.

### GetReplacedTypeOk

`func (o *TransactionReplacement) GetReplacedTypeOk() (*ReplaceType, bool)`

GetReplacedTypeOk returns a tuple with the ReplacedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedType

`func (o *TransactionReplacement) SetReplacedType(v ReplaceType)`

SetReplacedType sets ReplacedType field to given value.

### HasReplacedType

`func (o *TransactionReplacement) HasReplacedType() bool`

HasReplacedType returns a boolean if a field has been set.

### GetReplacedTransactionId

`func (o *TransactionReplacement) GetReplacedTransactionId() string`

GetReplacedTransactionId returns the ReplacedTransactionId field if non-nil, zero value otherwise.

### GetReplacedTransactionIdOk

`func (o *TransactionReplacement) GetReplacedTransactionIdOk() (*string, bool)`

GetReplacedTransactionIdOk returns a tuple with the ReplacedTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedTransactionId

`func (o *TransactionReplacement) SetReplacedTransactionId(v string)`

SetReplacedTransactionId sets ReplacedTransactionId field to given value.

### HasReplacedTransactionId

`func (o *TransactionReplacement) HasReplacedTransactionId() bool`

HasReplacedTransactionId returns a boolean if a field has been set.

### GetReplacedTransactionHash

`func (o *TransactionReplacement) GetReplacedTransactionHash() string`

GetReplacedTransactionHash returns the ReplacedTransactionHash field if non-nil, zero value otherwise.

### GetReplacedTransactionHashOk

`func (o *TransactionReplacement) GetReplacedTransactionHashOk() (*string, bool)`

GetReplacedTransactionHashOk returns a tuple with the ReplacedTransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedTransactionHash

`func (o *TransactionReplacement) SetReplacedTransactionHash(v string)`

SetReplacedTransactionHash sets ReplacedTransactionHash field to given value.

### HasReplacedTransactionHash

`func (o *TransactionReplacement) HasReplacedTransactionHash() bool`

HasReplacedTransactionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



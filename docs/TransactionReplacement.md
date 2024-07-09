# TransactionReplacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The transaction replacement type. Possible values include:    - &#x60;Cancel&#x60;: To cancel a transaction.   - &#x60;Drop&#x60;: To drop a transaction.   - &#x60;Resend&#x60;: To resend a transaction.   - &#x60;SpeedUp&#x60;: To speed up a transaction.  | [optional] 
**ReplacedByTransactionId** | Pointer to **string** | The ID of the replacement transaction that this transaction was replaced by. | [optional] 
**ReplacedByTransactionHash** | Pointer to **string** | The hash of the replacement transaction that this transaction was replaced by. | [optional] 
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

### GetType

`func (o *TransactionReplacement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionReplacement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionReplacement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionReplacement) HasType() bool`

HasType returns a boolean if a field has been set.

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



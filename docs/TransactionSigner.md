# TransactionSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signer** | Pointer to **string** | The signer name of the transaction. | [optional] 
**Status** | Pointer to **string** | The signing status. | [optional] 
**FailedReason** | Pointer to **string** | Failed reason of signing process. | [optional] 

## Methods

### NewTransactionSigner

`func NewTransactionSigner() *TransactionSigner`

NewTransactionSigner instantiates a new TransactionSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSignerWithDefaults

`func NewTransactionSignerWithDefaults() *TransactionSigner`

NewTransactionSignerWithDefaults instantiates a new TransactionSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigner

`func (o *TransactionSigner) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *TransactionSigner) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *TransactionSigner) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *TransactionSigner) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionSigner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionSigner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionSigner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionSigner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailedReason

`func (o *TransactionSigner) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *TransactionSigner) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *TransactionSigner) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *TransactionSigner) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** | Unique identifier for the transaction. | 
**TxHash** | Pointer to **string** | The blockchain transaction hash, may be initially null and populated after submission. | [optional] 
**FromAddress** | **string** | Source cryptocurrency address for the transaction. | 
**ToAddress** | **string** | Destination cryptocurrency address for the transaction. | 
**Amount** | **string** | The amount of cryptocurrency transferred, as a decimal string. | 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | 
**CreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | **int64** | The time when the transaction was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewPaymentTransaction

`func NewPaymentTransaction(txId string, fromAddress string, toAddress string, amount string, status TransactionStatus, createdTimestamp int64, updatedTimestamp int64, ) *PaymentTransaction`

NewPaymentTransaction instantiates a new PaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTransactionWithDefaults

`func NewPaymentTransactionWithDefaults() *PaymentTransaction`

NewPaymentTransactionWithDefaults instantiates a new PaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *PaymentTransaction) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *PaymentTransaction) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *PaymentTransaction) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetTxHash

`func (o *PaymentTransaction) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PaymentTransaction) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PaymentTransaction) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *PaymentTransaction) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetFromAddress

`func (o *PaymentTransaction) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *PaymentTransaction) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *PaymentTransaction) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *PaymentTransaction) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *PaymentTransaction) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *PaymentTransaction) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetAmount

`func (o *PaymentTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *PaymentTransaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentTransaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentTransaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentTransaction) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentTransaction) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentTransaction) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentTransaction) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentTransaction) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentTransaction) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



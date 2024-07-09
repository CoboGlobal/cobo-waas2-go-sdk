# TransactionRbf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**Fee** | Pointer to [**TransactionTransferFee**](TransactionTransferFee.md) |  | [optional] 

## Methods

### NewTransactionRbf

`func NewTransactionRbf(requestId string, ) *TransactionRbf`

NewTransactionRbf instantiates a new TransactionRbf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRbfWithDefaults

`func NewTransactionRbfWithDefaults() *TransactionRbf`

NewTransactionRbfWithDefaults instantiates a new TransactionRbf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TransactionRbf) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionRbf) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionRbf) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetFee

`func (o *TransactionRbf) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionRbf) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionRbf) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *TransactionRbf) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



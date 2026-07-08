# TransactionFeePayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraType** | [**TransactionExtraType**](TransactionExtraType.md) |  | 
**FeePayer** | Pointer to **string** | The address of the designated Solana fee payer account that covers the transaction fees, separating the fee payment from the main signer or source account. | [optional] 

## Methods

### NewTransactionFeePayer

`func NewTransactionFeePayer(extraType TransactionExtraType, ) *TransactionFeePayer`

NewTransactionFeePayer instantiates a new TransactionFeePayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFeePayerWithDefaults

`func NewTransactionFeePayerWithDefaults() *TransactionFeePayer`

NewTransactionFeePayerWithDefaults instantiates a new TransactionFeePayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraType

`func (o *TransactionFeePayer) GetExtraType() TransactionExtraType`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *TransactionFeePayer) GetExtraTypeOk() (*TransactionExtraType, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *TransactionFeePayer) SetExtraType(v TransactionExtraType)`

SetExtraType sets ExtraType field to given value.


### GetFeePayer

`func (o *TransactionFeePayer) GetFeePayer() string`

GetFeePayer returns the FeePayer field if non-nil, zero value otherwise.

### GetFeePayerOk

`func (o *TransactionFeePayer) GetFeePayerOk() (*string, bool)`

GetFeePayerOk returns a tuple with the FeePayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePayer

`func (o *TransactionFeePayer) SetFeePayer(v string)`

SetFeePayer sets FeePayer field to given value.

### HasFeePayer

`func (o *TransactionFeePayer) HasFeePayer() bool`

HasFeePayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



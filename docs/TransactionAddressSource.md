# TransactionAddressSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**AccountInput** | Pointer to [**TransactionAddressSourceAccountInput**](TransactionAddressSourceAccountInput.md) |  | [optional] 
**UtxoInputs** | Pointer to [**[]TransactionAddressSourceUtxoInputsInner**](TransactionAddressSourceUtxoInputsInner.md) |  | [optional] 

## Methods

### NewTransactionAddressSource

`func NewTransactionAddressSource(sourceType TransactionSourceType, ) *TransactionAddressSource`

NewTransactionAddressSource instantiates a new TransactionAddressSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAddressSourceWithDefaults

`func NewTransactionAddressSourceWithDefaults() *TransactionAddressSource`

NewTransactionAddressSourceWithDefaults instantiates a new TransactionAddressSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionAddressSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionAddressSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionAddressSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetAccountInput

`func (o *TransactionAddressSource) GetAccountInput() TransactionAddressSourceAccountInput`

GetAccountInput returns the AccountInput field if non-nil, zero value otherwise.

### GetAccountInputOk

`func (o *TransactionAddressSource) GetAccountInputOk() (*TransactionAddressSourceAccountInput, bool)`

GetAccountInputOk returns a tuple with the AccountInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInput

`func (o *TransactionAddressSource) SetAccountInput(v TransactionAddressSourceAccountInput)`

SetAccountInput sets AccountInput field to given value.

### HasAccountInput

`func (o *TransactionAddressSource) HasAccountInput() bool`

HasAccountInput returns a boolean if a field has been set.

### GetUtxoInputs

`func (o *TransactionAddressSource) GetUtxoInputs() []TransactionAddressSourceUtxoInputsInner`

GetUtxoInputs returns the UtxoInputs field if non-nil, zero value otherwise.

### GetUtxoInputsOk

`func (o *TransactionAddressSource) GetUtxoInputsOk() (*[]TransactionAddressSourceUtxoInputsInner, bool)`

GetUtxoInputsOk returns a tuple with the UtxoInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoInputs

`func (o *TransactionAddressSource) SetUtxoInputs(v []TransactionAddressSourceUtxoInputsInner)`

SetUtxoInputs sets UtxoInputs field to given value.

### HasUtxoInputs

`func (o *TransactionAddressSource) HasUtxoInputs() bool`

HasUtxoInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



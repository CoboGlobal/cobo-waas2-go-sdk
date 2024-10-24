# TransactionEvmContractMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The method name. | [optional] 
**Sig** | Pointer to **string** | The signature of the method, which includes the method name and parameter types. | [optional] 
**Type** | Pointer to **string** | The method type. | [optional] 
**Payable** | Pointer to **bool** | Whether the method is payable, which means it can receive tokens along with the transaction. - &#x60;true&#x60;: The method is payable. - &#x60;false&#x60;: The method is not payable.  | [optional] 
**Selector** | Pointer to **string** | The method selector, a four-byte identifier derived from the method&#39;s signature, used to invoke the method in a transaction. | [optional] 

## Methods

### NewTransactionEvmContractMethod

`func NewTransactionEvmContractMethod() *TransactionEvmContractMethod`

NewTransactionEvmContractMethod instantiates a new TransactionEvmContractMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEvmContractMethodWithDefaults

`func NewTransactionEvmContractMethodWithDefaults() *TransactionEvmContractMethod`

NewTransactionEvmContractMethodWithDefaults instantiates a new TransactionEvmContractMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransactionEvmContractMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionEvmContractMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionEvmContractMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionEvmContractMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSig

`func (o *TransactionEvmContractMethod) GetSig() string`

GetSig returns the Sig field if non-nil, zero value otherwise.

### GetSigOk

`func (o *TransactionEvmContractMethod) GetSigOk() (*string, bool)`

GetSigOk returns a tuple with the Sig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSig

`func (o *TransactionEvmContractMethod) SetSig(v string)`

SetSig sets Sig field to given value.

### HasSig

`func (o *TransactionEvmContractMethod) HasSig() bool`

HasSig returns a boolean if a field has been set.

### GetType

`func (o *TransactionEvmContractMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionEvmContractMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionEvmContractMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionEvmContractMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPayable

`func (o *TransactionEvmContractMethod) GetPayable() bool`

GetPayable returns the Payable field if non-nil, zero value otherwise.

### GetPayableOk

`func (o *TransactionEvmContractMethod) GetPayableOk() (*bool, bool)`

GetPayableOk returns a tuple with the Payable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayable

`func (o *TransactionEvmContractMethod) SetPayable(v bool)`

SetPayable sets Payable field to given value.

### HasPayable

`func (o *TransactionEvmContractMethod) HasPayable() bool`

HasPayable returns a boolean if a field has been set.

### GetSelector

`func (o *TransactionEvmContractMethod) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *TransactionEvmContractMethod) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *TransactionEvmContractMethod) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *TransactionEvmContractMethod) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



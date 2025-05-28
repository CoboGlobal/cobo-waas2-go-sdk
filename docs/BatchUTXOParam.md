# BatchUTXOParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | **string** | The transaction hash. | 
**VoutNs** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBatchUTXOParam

`func NewBatchUTXOParam(txHash string, ) *BatchUTXOParam`

NewBatchUTXOParam instantiates a new BatchUTXOParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUTXOParamWithDefaults

`func NewBatchUTXOParamWithDefaults() *BatchUTXOParam`

NewBatchUTXOParamWithDefaults instantiates a new BatchUTXOParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *BatchUTXOParam) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *BatchUTXOParam) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *BatchUTXOParam) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetVoutNs

`func (o *BatchUTXOParam) GetVoutNs() []int32`

GetVoutNs returns the VoutNs field if non-nil, zero value otherwise.

### GetVoutNsOk

`func (o *BatchUTXOParam) GetVoutNsOk() (*[]int32, bool)`

GetVoutNsOk returns a tuple with the VoutNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoutNs

`func (o *BatchUTXOParam) SetVoutNs(v []int32)`

SetVoutNs sets VoutNs field to given value.

### HasVoutNs

`func (o *BatchUTXOParam) HasVoutNs() bool`

HasVoutNs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



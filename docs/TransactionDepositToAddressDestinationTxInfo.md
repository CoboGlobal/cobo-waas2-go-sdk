# TransactionDepositToAddressDestinationTxInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoutN** | Pointer to **int32** | The output index of the UTXO. | [optional] 
**ObjectId** | Pointer to **string** | The ID of the blockchain object to spend (e.g., SUI Coin object). | [optional] 
**Version** | Pointer to **string** | Object version number. | [optional] 

## Methods

### NewTransactionDepositToAddressDestinationTxInfo

`func NewTransactionDepositToAddressDestinationTxInfo() *TransactionDepositToAddressDestinationTxInfo`

NewTransactionDepositToAddressDestinationTxInfo instantiates a new TransactionDepositToAddressDestinationTxInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDepositToAddressDestinationTxInfoWithDefaults

`func NewTransactionDepositToAddressDestinationTxInfoWithDefaults() *TransactionDepositToAddressDestinationTxInfo`

NewTransactionDepositToAddressDestinationTxInfoWithDefaults instantiates a new TransactionDepositToAddressDestinationTxInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoutN

`func (o *TransactionDepositToAddressDestinationTxInfo) GetVoutN() int32`

GetVoutN returns the VoutN field if non-nil, zero value otherwise.

### GetVoutNOk

`func (o *TransactionDepositToAddressDestinationTxInfo) GetVoutNOk() (*int32, bool)`

GetVoutNOk returns a tuple with the VoutN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoutN

`func (o *TransactionDepositToAddressDestinationTxInfo) SetVoutN(v int32)`

SetVoutN sets VoutN field to given value.

### HasVoutN

`func (o *TransactionDepositToAddressDestinationTxInfo) HasVoutN() bool`

HasVoutN returns a boolean if a field has been set.

### GetObjectId

`func (o *TransactionDepositToAddressDestinationTxInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *TransactionDepositToAddressDestinationTxInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *TransactionDepositToAddressDestinationTxInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *TransactionDepositToAddressDestinationTxInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetVersion

`func (o *TransactionDepositToAddressDestinationTxInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TransactionDepositToAddressDestinationTxInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TransactionDepositToAddressDestinationTxInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TransactionDepositToAddressDestinationTxInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



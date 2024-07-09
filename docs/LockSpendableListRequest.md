# LockSpendableListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
**TxHash** | **string** | The transaction hash. | 
**VoutN** | **int32** | The output index of the UTXO. | 

## Methods

### NewLockSpendableListRequest

`func NewLockSpendableListRequest(tokenId string, txHash string, voutN int32, ) *LockSpendableListRequest`

NewLockSpendableListRequest instantiates a new LockSpendableListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockSpendableListRequestWithDefaults

`func NewLockSpendableListRequestWithDefaults() *LockSpendableListRequest`

NewLockSpendableListRequestWithDefaults instantiates a new LockSpendableListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *LockSpendableListRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *LockSpendableListRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *LockSpendableListRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetTxHash

`func (o *LockSpendableListRequest) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *LockSpendableListRequest) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *LockSpendableListRequest) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetVoutN

`func (o *LockSpendableListRequest) GetVoutN() int32`

GetVoutN returns the VoutN field if non-nil, zero value otherwise.

### GetVoutNOk

`func (o *LockSpendableListRequest) GetVoutNOk() (*int32, bool)`

GetVoutNOk returns a tuple with the VoutN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoutN

`func (o *LockSpendableListRequest) SetVoutN(v int32)`

SetVoutN sets VoutN field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



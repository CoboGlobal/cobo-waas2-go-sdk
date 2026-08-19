# BalanceAtBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The wallet address. | [optional] 
**TokenId** | Pointer to **string** | The token ID, which is the unique identifier of a token. | [optional] 
**Balance** | Pointer to **string** | The token balance of the address. | [optional] 
**BlockNumber** | Pointer to **int64** | The block number (block height) at which the balance was retrieved. | [optional] 

## Methods

### NewBalanceAtBlock

`func NewBalanceAtBlock() *BalanceAtBlock`

NewBalanceAtBlock instantiates a new BalanceAtBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAtBlockWithDefaults

`func NewBalanceAtBlockWithDefaults() *BalanceAtBlock`

NewBalanceAtBlockWithDefaults instantiates a new BalanceAtBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BalanceAtBlock) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BalanceAtBlock) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BalanceAtBlock) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BalanceAtBlock) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *BalanceAtBlock) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BalanceAtBlock) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BalanceAtBlock) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *BalanceAtBlock) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetBalance

`func (o *BalanceAtBlock) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceAtBlock) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceAtBlock) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BalanceAtBlock) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBlockNumber

`func (o *BalanceAtBlock) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *BalanceAtBlock) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *BalanceAtBlock) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *BalanceAtBlock) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



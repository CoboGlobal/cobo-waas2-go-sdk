# BabylonStakeExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FinalityProviderPublicKey** | **string** | The public key of the finality provider. | 
**StakeBlockTime** | **int64** | The number of blocks that need to be processed before the locked tokens are unlocked and become accessible. | 
**AutoBroadcast** | Pointer to **bool** | Whether to automatically broadcast the transaction. The default value is &#x60;true&#x60;.  - &#x60;true&#x60;: Automatically broadcast the transaction. - &#x60;false&#x60;: The transaction will not be submitted to the blockchain automatically. You can call [Broadcast signed transactions](/v2/api-references/transactions/broadcast-signed-transactions) to broadcast the transaction to the blockchain, or retrieve the signed raw transaction data &#x60;raw_tx&#x60; by calling [Get transaction information](/v2/api-references/transactions/get-transaction-information) and broadcast it yourself.  | [optional] 

## Methods

### NewBabylonStakeExtra

`func NewBabylonStakeExtra(poolType StakingPoolType, finalityProviderPublicKey string, stakeBlockTime int64, ) *BabylonStakeExtra`

NewBabylonStakeExtra instantiates a new BabylonStakeExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonStakeExtraWithDefaults

`func NewBabylonStakeExtraWithDefaults() *BabylonStakeExtra`

NewBabylonStakeExtraWithDefaults instantiates a new BabylonStakeExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *BabylonStakeExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *BabylonStakeExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *BabylonStakeExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetFinalityProviderPublicKey

`func (o *BabylonStakeExtra) GetFinalityProviderPublicKey() string`

GetFinalityProviderPublicKey returns the FinalityProviderPublicKey field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeyOk

`func (o *BabylonStakeExtra) GetFinalityProviderPublicKeyOk() (*string, bool)`

GetFinalityProviderPublicKeyOk returns a tuple with the FinalityProviderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKey

`func (o *BabylonStakeExtra) SetFinalityProviderPublicKey(v string)`

SetFinalityProviderPublicKey sets FinalityProviderPublicKey field to given value.


### GetStakeBlockTime

`func (o *BabylonStakeExtra) GetStakeBlockTime() int64`

GetStakeBlockTime returns the StakeBlockTime field if non-nil, zero value otherwise.

### GetStakeBlockTimeOk

`func (o *BabylonStakeExtra) GetStakeBlockTimeOk() (*int64, bool)`

GetStakeBlockTimeOk returns a tuple with the StakeBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBlockTime

`func (o *BabylonStakeExtra) SetStakeBlockTime(v int64)`

SetStakeBlockTime sets StakeBlockTime field to given value.


### GetAutoBroadcast

`func (o *BabylonStakeExtra) GetAutoBroadcast() bool`

GetAutoBroadcast returns the AutoBroadcast field if non-nil, zero value otherwise.

### GetAutoBroadcastOk

`func (o *BabylonStakeExtra) GetAutoBroadcastOk() (*bool, bool)`

GetAutoBroadcastOk returns a tuple with the AutoBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBroadcast

`func (o *BabylonStakeExtra) SetAutoBroadcast(v bool)`

SetAutoBroadcast sets AutoBroadcast field to given value.

### HasAutoBroadcast

`func (o *BabylonStakeExtra) HasAutoBroadcast() bool`

HasAutoBroadcast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



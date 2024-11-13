# CreateStakeActivityExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FinalityProviderPublicKey** | **string** | The public key of the finality provider. | 
**StakeBlockTime** | **int64** | The number of blocks that need to be processed before the locked tokens are unlocked and become accessible. | 
**AutoBroadcast** | Pointer to **bool** | Whether to automatically broadcast the transaction. The default value is &#x60;true&#x60;.  - &#x60;true&#x60;: Automatically broadcast the transaction. - &#x60;false&#x60;: The transaction will not be submitted to the blockchain automatically. You can call [Broadcast signed transactions](/v2/api-references/transactions/broadcast-signed-transactions) to broadcast the transaction to the blockchain, or retrieve the signed raw transaction data &#x60;raw_tx&#x60; by calling [Get transaction information](/v2/api-references/transactions/get-transaction-information) and broadcast it yourself.  | [optional] 
**ProviderName** | **string** | The name of the provider. | 

## Methods

### NewCreateStakeActivityExtra

`func NewCreateStakeActivityExtra(poolType StakingPoolType, finalityProviderPublicKey string, stakeBlockTime int64, providerName string, ) *CreateStakeActivityExtra`

NewCreateStakeActivityExtra instantiates a new CreateStakeActivityExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStakeActivityExtraWithDefaults

`func NewCreateStakeActivityExtraWithDefaults() *CreateStakeActivityExtra`

NewCreateStakeActivityExtraWithDefaults instantiates a new CreateStakeActivityExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *CreateStakeActivityExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *CreateStakeActivityExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *CreateStakeActivityExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetFinalityProviderPublicKey

`func (o *CreateStakeActivityExtra) GetFinalityProviderPublicKey() string`

GetFinalityProviderPublicKey returns the FinalityProviderPublicKey field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeyOk

`func (o *CreateStakeActivityExtra) GetFinalityProviderPublicKeyOk() (*string, bool)`

GetFinalityProviderPublicKeyOk returns a tuple with the FinalityProviderPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKey

`func (o *CreateStakeActivityExtra) SetFinalityProviderPublicKey(v string)`

SetFinalityProviderPublicKey sets FinalityProviderPublicKey field to given value.


### GetStakeBlockTime

`func (o *CreateStakeActivityExtra) GetStakeBlockTime() int64`

GetStakeBlockTime returns the StakeBlockTime field if non-nil, zero value otherwise.

### GetStakeBlockTimeOk

`func (o *CreateStakeActivityExtra) GetStakeBlockTimeOk() (*int64, bool)`

GetStakeBlockTimeOk returns a tuple with the StakeBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBlockTime

`func (o *CreateStakeActivityExtra) SetStakeBlockTime(v int64)`

SetStakeBlockTime sets StakeBlockTime field to given value.


### GetAutoBroadcast

`func (o *CreateStakeActivityExtra) GetAutoBroadcast() bool`

GetAutoBroadcast returns the AutoBroadcast field if non-nil, zero value otherwise.

### GetAutoBroadcastOk

`func (o *CreateStakeActivityExtra) GetAutoBroadcastOk() (*bool, bool)`

GetAutoBroadcastOk returns a tuple with the AutoBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBroadcast

`func (o *CreateStakeActivityExtra) SetAutoBroadcast(v bool)`

SetAutoBroadcast sets AutoBroadcast field to given value.

### HasAutoBroadcast

`func (o *CreateStakeActivityExtra) HasAutoBroadcast() bool`

HasAutoBroadcast returns a boolean if a field has been set.

### GetProviderName

`func (o *CreateStakeActivityExtra) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CreateStakeActivityExtra) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CreateStakeActivityExtra) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



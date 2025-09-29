# BabylonStakeExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FinalityProviderPublicKey** | Pointer to **string** | The public key of the finality provider. | [optional] 
**FinalityProviderPublicKeys** | Pointer to **[]string** | The public keys of the finality providers(each key for a BSN chain). | [optional] 
**StakeBlockTime** | Pointer to **int64** | The number of blocks that need to be processed before the locked tokens are unlocked and become accessible. | [optional] 
**AutoBroadcast** | Pointer to **bool** | Whether to automatically broadcast the transaction. The default value is &#x60;true&#x60;.  - &#x60;true&#x60;: Automatically broadcast the transaction. - &#x60;false&#x60;: The transaction will not be submitted to the blockchain automatically. You can call [Broadcast signed transactions](https://www.cobo.com/developers/v2/api-references/transactions/broadcast-signed-transactions) to broadcast the transaction to the blockchain, or retrieve the signed raw transaction data &#x60;raw_tx&#x60; by calling [Get transaction information](https://www.cobo.com/developers/v2/api-references/transactions/get-transaction-information) and broadcast it yourself.  | [optional] 
**BabylonAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**OriginalStakingId** | Pointer to **string** | The original staking ID to expand. Only set this when you want to expand existing staking. | [optional] 

## Methods

### NewBabylonStakeExtra

`func NewBabylonStakeExtra(poolType StakingPoolType, ) *BabylonStakeExtra`

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

### HasFinalityProviderPublicKey

`func (o *BabylonStakeExtra) HasFinalityProviderPublicKey() bool`

HasFinalityProviderPublicKey returns a boolean if a field has been set.

### GetFinalityProviderPublicKeys

`func (o *BabylonStakeExtra) GetFinalityProviderPublicKeys() []string`

GetFinalityProviderPublicKeys returns the FinalityProviderPublicKeys field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeysOk

`func (o *BabylonStakeExtra) GetFinalityProviderPublicKeysOk() (*[]string, bool)`

GetFinalityProviderPublicKeysOk returns a tuple with the FinalityProviderPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKeys

`func (o *BabylonStakeExtra) SetFinalityProviderPublicKeys(v []string)`

SetFinalityProviderPublicKeys sets FinalityProviderPublicKeys field to given value.

### HasFinalityProviderPublicKeys

`func (o *BabylonStakeExtra) HasFinalityProviderPublicKeys() bool`

HasFinalityProviderPublicKeys returns a boolean if a field has been set.

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

### HasStakeBlockTime

`func (o *BabylonStakeExtra) HasStakeBlockTime() bool`

HasStakeBlockTime returns a boolean if a field has been set.

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

### GetBabylonAddress

`func (o *BabylonStakeExtra) GetBabylonAddress() StakingSource`

GetBabylonAddress returns the BabylonAddress field if non-nil, zero value otherwise.

### GetBabylonAddressOk

`func (o *BabylonStakeExtra) GetBabylonAddressOk() (*StakingSource, bool)`

GetBabylonAddressOk returns a tuple with the BabylonAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddress

`func (o *BabylonStakeExtra) SetBabylonAddress(v StakingSource)`

SetBabylonAddress sets BabylonAddress field to given value.

### HasBabylonAddress

`func (o *BabylonStakeExtra) HasBabylonAddress() bool`

HasBabylonAddress returns a boolean if a field has been set.

### GetOriginalStakingId

`func (o *BabylonStakeExtra) GetOriginalStakingId() string`

GetOriginalStakingId returns the OriginalStakingId field if non-nil, zero value otherwise.

### GetOriginalStakingIdOk

`func (o *BabylonStakeExtra) GetOriginalStakingIdOk() (*string, bool)`

GetOriginalStakingIdOk returns a tuple with the OriginalStakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStakingId

`func (o *BabylonStakeExtra) SetOriginalStakingId(v string)`

SetOriginalStakingId sets OriginalStakingId field to given value.

### HasOriginalStakingId

`func (o *BabylonStakeExtra) HasOriginalStakingId() bool`

HasOriginalStakingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



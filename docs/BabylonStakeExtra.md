# BabylonStakeExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FinalityProviderPublicKey** | **string** | The public key of finality provider. | 
**StakeBlockTime** | **int64** | The stake block time. | 
**OnlySign** | Pointer to **bool** | Whether to only sign transactions. Default is &#x60;false&#x60;, if set to &#x60;true&#x60;,  the transaction will not be submitted to the blockchain automatically. You can call &#x60;Broadcast transactions&#x60; to submit the transaction to the blockchain,  Or you can find the signed raw_tx by &#x60;Get transaction information&#x60; and broadcast it yourself.  | [optional] 

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


### GetOnlySign

`func (o *BabylonStakeExtra) GetOnlySign() bool`

GetOnlySign returns the OnlySign field if non-nil, zero value otherwise.

### GetOnlySignOk

`func (o *BabylonStakeExtra) GetOnlySignOk() (*bool, bool)`

GetOnlySignOk returns a tuple with the OnlySign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlySign

`func (o *BabylonStakeExtra) SetOnlySign(v bool)`

SetOnlySign sets OnlySign field to given value.

### HasOnlySign

`func (o *BabylonStakeExtra) HasOnlySign() bool`

HasOnlySign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



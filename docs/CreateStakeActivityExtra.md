# CreateStakeActivityExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FinalityProviderPublicKey** | **string** | The public key of finality provider. | 
**StakeBlockTime** | **int64** | The stake block time. | 
**OnlySign** | Pointer to **bool** | Whether to only sign transactions. Default is &#x60;false&#x60;, if set to &#x60;true&#x60;,  the transaction will not be submitted to the blockchain automatically. You can call &#x60;Broadcast transactions&#x60; to submit the transaction to the blockchain,  Or you can find the signed raw_tx by &#x60;Get transaction information&#x60; and broadcast it yourself.  | [optional] 
**Operator** | Pointer to **string** | The operator address. | [optional] 
**FeeRecipient** | Pointer to **float32** | The fee recipient address, if not provided the staker address will be used. | [optional] 

## Methods

### NewCreateStakeActivityExtra

`func NewCreateStakeActivityExtra(poolType StakingPoolType, finalityProviderPublicKey string, stakeBlockTime int64, ) *CreateStakeActivityExtra`

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


### GetOnlySign

`func (o *CreateStakeActivityExtra) GetOnlySign() bool`

GetOnlySign returns the OnlySign field if non-nil, zero value otherwise.

### GetOnlySignOk

`func (o *CreateStakeActivityExtra) GetOnlySignOk() (*bool, bool)`

GetOnlySignOk returns a tuple with the OnlySign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlySign

`func (o *CreateStakeActivityExtra) SetOnlySign(v bool)`

SetOnlySign sets OnlySign field to given value.

### HasOnlySign

`func (o *CreateStakeActivityExtra) HasOnlySign() bool`

HasOnlySign returns a boolean if a field has been set.

### GetOperator

`func (o *CreateStakeActivityExtra) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CreateStakeActivityExtra) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CreateStakeActivityExtra) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CreateStakeActivityExtra) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetFeeRecipient

`func (o *CreateStakeActivityExtra) GetFeeRecipient() float32`

GetFeeRecipient returns the FeeRecipient field if non-nil, zero value otherwise.

### GetFeeRecipientOk

`func (o *CreateStakeActivityExtra) GetFeeRecipientOk() (*float32, bool)`

GetFeeRecipientOk returns a tuple with the FeeRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRecipient

`func (o *CreateStakeActivityExtra) SetFeeRecipient(v float32)`

SetFeeRecipient sets FeeRecipient field to given value.

### HasFeeRecipient

`func (o *CreateStakeActivityExtra) HasFeeRecipient() bool`

HasFeeRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



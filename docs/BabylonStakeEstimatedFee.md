# BabylonStakeEstimatedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | Pointer to [**StakingPoolType**](StakingPoolType.md) |  | [optional] 
**FeeType** | Pointer to [**FeeType**](FeeType.md) |  | [optional] [default to FEETYPE_EVM_EIP_1559]
**FeeAmount** | Pointer to **string** | The amount of the estimated fee. | [optional] 
**TokenId** | Pointer to **string** | The token ID of the staking fee. | [optional] 

## Methods

### NewBabylonStakeEstimatedFee

`func NewBabylonStakeEstimatedFee() *BabylonStakeEstimatedFee`

NewBabylonStakeEstimatedFee instantiates a new BabylonStakeEstimatedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonStakeEstimatedFeeWithDefaults

`func NewBabylonStakeEstimatedFeeWithDefaults() *BabylonStakeEstimatedFee`

NewBabylonStakeEstimatedFeeWithDefaults instantiates a new BabylonStakeEstimatedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *BabylonStakeEstimatedFee) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *BabylonStakeEstimatedFee) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *BabylonStakeEstimatedFee) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *BabylonStakeEstimatedFee) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetFeeType

`func (o *BabylonStakeEstimatedFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *BabylonStakeEstimatedFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *BabylonStakeEstimatedFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *BabylonStakeEstimatedFee) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetFeeAmount

`func (o *BabylonStakeEstimatedFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *BabylonStakeEstimatedFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *BabylonStakeEstimatedFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *BabylonStakeEstimatedFee) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *BabylonStakeEstimatedFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BabylonStakeEstimatedFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BabylonStakeEstimatedFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *BabylonStakeEstimatedFee) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



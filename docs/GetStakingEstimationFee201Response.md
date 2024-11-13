# GetStakingEstimationFee201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | Pointer to [**StakingPoolType**](StakingPoolType.md) |  | [optional] 
**Fee** | Pointer to [**EstimatedFee**](EstimatedFee.md) |  | [optional] 
**ValidatorPubkeys** | Pointer to **[]string** | A list of public keys associated with the Ethereum validators for this staking operation. | [optional] 
**FeeType** | Pointer to [**FeeType**](FeeType.md) |  | [optional] [default to FEETYPE_EVM_EIP_1559]
**FeeAmount** | Pointer to **string** | The amount of the estimated fee. | [optional] 
**TokenId** | Pointer to **string** | The token ID of the staking fee. | [optional] 

## Methods

### NewGetStakingEstimationFee201Response

`func NewGetStakingEstimationFee201Response() *GetStakingEstimationFee201Response`

NewGetStakingEstimationFee201Response instantiates a new GetStakingEstimationFee201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStakingEstimationFee201ResponseWithDefaults

`func NewGetStakingEstimationFee201ResponseWithDefaults() *GetStakingEstimationFee201Response`

NewGetStakingEstimationFee201ResponseWithDefaults instantiates a new GetStakingEstimationFee201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *GetStakingEstimationFee201Response) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *GetStakingEstimationFee201Response) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *GetStakingEstimationFee201Response) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *GetStakingEstimationFee201Response) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetFee

`func (o *GetStakingEstimationFee201Response) GetFee() EstimatedFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetStakingEstimationFee201Response) GetFeeOk() (*EstimatedFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetStakingEstimationFee201Response) SetFee(v EstimatedFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetStakingEstimationFee201Response) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetValidatorPubkeys

`func (o *GetStakingEstimationFee201Response) GetValidatorPubkeys() []string`

GetValidatorPubkeys returns the ValidatorPubkeys field if non-nil, zero value otherwise.

### GetValidatorPubkeysOk

`func (o *GetStakingEstimationFee201Response) GetValidatorPubkeysOk() (*[]string, bool)`

GetValidatorPubkeysOk returns a tuple with the ValidatorPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorPubkeys

`func (o *GetStakingEstimationFee201Response) SetValidatorPubkeys(v []string)`

SetValidatorPubkeys sets ValidatorPubkeys field to given value.

### HasValidatorPubkeys

`func (o *GetStakingEstimationFee201Response) HasValidatorPubkeys() bool`

HasValidatorPubkeys returns a boolean if a field has been set.

### GetFeeType

`func (o *GetStakingEstimationFee201Response) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *GetStakingEstimationFee201Response) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *GetStakingEstimationFee201Response) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *GetStakingEstimationFee201Response) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetFeeAmount

`func (o *GetStakingEstimationFee201Response) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *GetStakingEstimationFee201Response) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *GetStakingEstimationFee201Response) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *GetStakingEstimationFee201Response) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *GetStakingEstimationFee201Response) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetStakingEstimationFee201Response) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetStakingEstimationFee201Response) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetStakingEstimationFee201Response) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



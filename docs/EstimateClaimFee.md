# EstimateClaimFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | [**ActivityType**](ActivityType.md) |  | 
**StakingId** | Pointer to **string** | The ID of the staking position. You can retrieve a list of staking positions by calling [List staking positions](https://www.cobo.com/developers/v2/api-references/stakings/list-staking-positions). | [optional] 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 

## Methods

### NewEstimateClaimFee

`func NewEstimateClaimFee(activityType ActivityType, ) *EstimateClaimFee`

NewEstimateClaimFee instantiates a new EstimateClaimFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateClaimFeeWithDefaults

`func NewEstimateClaimFeeWithDefaults() *EstimateClaimFee`

NewEstimateClaimFeeWithDefaults instantiates a new EstimateClaimFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *EstimateClaimFee) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *EstimateClaimFee) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *EstimateClaimFee) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.


### GetStakingId

`func (o *EstimateClaimFee) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *EstimateClaimFee) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *EstimateClaimFee) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.

### HasStakingId

`func (o *EstimateClaimFee) HasStakingId() bool`

HasStakingId returns a boolean if a field has been set.

### GetFee

`func (o *EstimateClaimFee) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EstimateClaimFee) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EstimateClaimFee) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *EstimateClaimFee) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



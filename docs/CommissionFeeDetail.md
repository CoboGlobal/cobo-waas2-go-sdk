# CommissionFeeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommissionFeeId** | Pointer to **string** | Unique ID of the commission fee | [optional] 
**AssetId** | Pointer to **string** | Asset ID | [optional] 
**FeeTokens** | Pointer to **[]string** | List of fee token IDs | [optional] 
**SettleType** | Pointer to [**SettleType**](SettleType.md) |  | [optional] 
**BusinessTypeId** | Pointer to **int32** | Unique ID of business type | [optional] 
**NormalizedAmount** | Pointer to [**CommissionFeeDetailNormalizedAmount**](CommissionFeeDetailNormalizedAmount.md) |  | [optional] 
**StrategyId** | Pointer to **int32** | Strategy ID | [optional] 
**RuleId** | Pointer to **int32** | Rule ID | [optional] 
**StrategyParams** | Pointer to [**StrategyParams**](StrategyParams.md) |  | [optional] 

## Methods

### NewCommissionFeeDetail

`func NewCommissionFeeDetail() *CommissionFeeDetail`

NewCommissionFeeDetail instantiates a new CommissionFeeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionFeeDetailWithDefaults

`func NewCommissionFeeDetailWithDefaults() *CommissionFeeDetail`

NewCommissionFeeDetailWithDefaults instantiates a new CommissionFeeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommissionFeeId

`func (o *CommissionFeeDetail) GetCommissionFeeId() string`

GetCommissionFeeId returns the CommissionFeeId field if non-nil, zero value otherwise.

### GetCommissionFeeIdOk

`func (o *CommissionFeeDetail) GetCommissionFeeIdOk() (*string, bool)`

GetCommissionFeeIdOk returns a tuple with the CommissionFeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFeeId

`func (o *CommissionFeeDetail) SetCommissionFeeId(v string)`

SetCommissionFeeId sets CommissionFeeId field to given value.

### HasCommissionFeeId

`func (o *CommissionFeeDetail) HasCommissionFeeId() bool`

HasCommissionFeeId returns a boolean if a field has been set.

### GetAssetId

`func (o *CommissionFeeDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *CommissionFeeDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *CommissionFeeDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *CommissionFeeDetail) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetFeeTokens

`func (o *CommissionFeeDetail) GetFeeTokens() []string`

GetFeeTokens returns the FeeTokens field if non-nil, zero value otherwise.

### GetFeeTokensOk

`func (o *CommissionFeeDetail) GetFeeTokensOk() (*[]string, bool)`

GetFeeTokensOk returns a tuple with the FeeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokens

`func (o *CommissionFeeDetail) SetFeeTokens(v []string)`

SetFeeTokens sets FeeTokens field to given value.

### HasFeeTokens

`func (o *CommissionFeeDetail) HasFeeTokens() bool`

HasFeeTokens returns a boolean if a field has been set.

### GetSettleType

`func (o *CommissionFeeDetail) GetSettleType() SettleType`

GetSettleType returns the SettleType field if non-nil, zero value otherwise.

### GetSettleTypeOk

`func (o *CommissionFeeDetail) GetSettleTypeOk() (*SettleType, bool)`

GetSettleTypeOk returns a tuple with the SettleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleType

`func (o *CommissionFeeDetail) SetSettleType(v SettleType)`

SetSettleType sets SettleType field to given value.

### HasSettleType

`func (o *CommissionFeeDetail) HasSettleType() bool`

HasSettleType returns a boolean if a field has been set.

### GetBusinessTypeId

`func (o *CommissionFeeDetail) GetBusinessTypeId() int32`

GetBusinessTypeId returns the BusinessTypeId field if non-nil, zero value otherwise.

### GetBusinessTypeIdOk

`func (o *CommissionFeeDetail) GetBusinessTypeIdOk() (*int32, bool)`

GetBusinessTypeIdOk returns a tuple with the BusinessTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTypeId

`func (o *CommissionFeeDetail) SetBusinessTypeId(v int32)`

SetBusinessTypeId sets BusinessTypeId field to given value.

### HasBusinessTypeId

`func (o *CommissionFeeDetail) HasBusinessTypeId() bool`

HasBusinessTypeId returns a boolean if a field has been set.

### GetNormalizedAmount

`func (o *CommissionFeeDetail) GetNormalizedAmount() CommissionFeeDetailNormalizedAmount`

GetNormalizedAmount returns the NormalizedAmount field if non-nil, zero value otherwise.

### GetNormalizedAmountOk

`func (o *CommissionFeeDetail) GetNormalizedAmountOk() (*CommissionFeeDetailNormalizedAmount, bool)`

GetNormalizedAmountOk returns a tuple with the NormalizedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedAmount

`func (o *CommissionFeeDetail) SetNormalizedAmount(v CommissionFeeDetailNormalizedAmount)`

SetNormalizedAmount sets NormalizedAmount field to given value.

### HasNormalizedAmount

`func (o *CommissionFeeDetail) HasNormalizedAmount() bool`

HasNormalizedAmount returns a boolean if a field has been set.

### GetStrategyId

`func (o *CommissionFeeDetail) GetStrategyId() int32`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *CommissionFeeDetail) GetStrategyIdOk() (*int32, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *CommissionFeeDetail) SetStrategyId(v int32)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *CommissionFeeDetail) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetRuleId

`func (o *CommissionFeeDetail) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *CommissionFeeDetail) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *CommissionFeeDetail) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *CommissionFeeDetail) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetStrategyParams

`func (o *CommissionFeeDetail) GetStrategyParams() StrategyParams`

GetStrategyParams returns the StrategyParams field if non-nil, zero value otherwise.

### GetStrategyParamsOk

`func (o *CommissionFeeDetail) GetStrategyParamsOk() (*StrategyParams, bool)`

GetStrategyParamsOk returns a tuple with the StrategyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyParams

`func (o *CommissionFeeDetail) SetStrategyParams(v StrategyParams)`

SetStrategyParams sets StrategyParams field to given value.

### HasStrategyParams

`func (o *CommissionFeeDetail) HasStrategyParams() bool`

HasStrategyParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



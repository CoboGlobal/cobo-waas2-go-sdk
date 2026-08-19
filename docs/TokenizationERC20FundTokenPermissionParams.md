# TokenizationERC20FundTokenPermissionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **[]string** | List of addresses for the owner/admin role. Have full administrative control over the fund contract. | [optional] 
**Manager** | Pointer to **[]string** | List of addresses for the fund manager role. Can manage fund configurations and operational parameters. | [optional] 
**NavUpdater** | Pointer to **[]string** | List of addresses for the NAV updater role. Can update the net asset value (NAV) per share price. | [optional] 
**RedemptionApprover** | Pointer to **[]string** | List of addresses for the redemption approver role. Can approve or reject investor redemption requests. | [optional] 
**SettlementOperator** | Pointer to **[]string** | List of addresses for the settlement operator role. Can execute investment and redemption settlement operations. | [optional] 
**EmergencyGuardian** | Pointer to **[]string** | List of addresses for the emergency guardian role. Can trigger emergency actions such as pausing the fund. | [optional] 

## Methods

### NewTokenizationERC20FundTokenPermissionParams

`func NewTokenizationERC20FundTokenPermissionParams() *TokenizationERC20FundTokenPermissionParams`

NewTokenizationERC20FundTokenPermissionParams instantiates a new TokenizationERC20FundTokenPermissionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationERC20FundTokenPermissionParamsWithDefaults

`func NewTokenizationERC20FundTokenPermissionParamsWithDefaults() *TokenizationERC20FundTokenPermissionParams`

NewTokenizationERC20FundTokenPermissionParamsWithDefaults instantiates a new TokenizationERC20FundTokenPermissionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *TokenizationERC20FundTokenPermissionParams) GetOwner() []string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TokenizationERC20FundTokenPermissionParams) GetOwnerOk() (*[]string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TokenizationERC20FundTokenPermissionParams) SetOwner(v []string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TokenizationERC20FundTokenPermissionParams) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetManager

`func (o *TokenizationERC20FundTokenPermissionParams) GetManager() []string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *TokenizationERC20FundTokenPermissionParams) GetManagerOk() (*[]string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *TokenizationERC20FundTokenPermissionParams) SetManager(v []string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *TokenizationERC20FundTokenPermissionParams) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNavUpdater

`func (o *TokenizationERC20FundTokenPermissionParams) GetNavUpdater() []string`

GetNavUpdater returns the NavUpdater field if non-nil, zero value otherwise.

### GetNavUpdaterOk

`func (o *TokenizationERC20FundTokenPermissionParams) GetNavUpdaterOk() (*[]string, bool)`

GetNavUpdaterOk returns a tuple with the NavUpdater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavUpdater

`func (o *TokenizationERC20FundTokenPermissionParams) SetNavUpdater(v []string)`

SetNavUpdater sets NavUpdater field to given value.

### HasNavUpdater

`func (o *TokenizationERC20FundTokenPermissionParams) HasNavUpdater() bool`

HasNavUpdater returns a boolean if a field has been set.

### GetRedemptionApprover

`func (o *TokenizationERC20FundTokenPermissionParams) GetRedemptionApprover() []string`

GetRedemptionApprover returns the RedemptionApprover field if non-nil, zero value otherwise.

### GetRedemptionApproverOk

`func (o *TokenizationERC20FundTokenPermissionParams) GetRedemptionApproverOk() (*[]string, bool)`

GetRedemptionApproverOk returns a tuple with the RedemptionApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionApprover

`func (o *TokenizationERC20FundTokenPermissionParams) SetRedemptionApprover(v []string)`

SetRedemptionApprover sets RedemptionApprover field to given value.

### HasRedemptionApprover

`func (o *TokenizationERC20FundTokenPermissionParams) HasRedemptionApprover() bool`

HasRedemptionApprover returns a boolean if a field has been set.

### GetSettlementOperator

`func (o *TokenizationERC20FundTokenPermissionParams) GetSettlementOperator() []string`

GetSettlementOperator returns the SettlementOperator field if non-nil, zero value otherwise.

### GetSettlementOperatorOk

`func (o *TokenizationERC20FundTokenPermissionParams) GetSettlementOperatorOk() (*[]string, bool)`

GetSettlementOperatorOk returns a tuple with the SettlementOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementOperator

`func (o *TokenizationERC20FundTokenPermissionParams) SetSettlementOperator(v []string)`

SetSettlementOperator sets SettlementOperator field to given value.

### HasSettlementOperator

`func (o *TokenizationERC20FundTokenPermissionParams) HasSettlementOperator() bool`

HasSettlementOperator returns a boolean if a field has been set.

### GetEmergencyGuardian

`func (o *TokenizationERC20FundTokenPermissionParams) GetEmergencyGuardian() []string`

GetEmergencyGuardian returns the EmergencyGuardian field if non-nil, zero value otherwise.

### GetEmergencyGuardianOk

`func (o *TokenizationERC20FundTokenPermissionParams) GetEmergencyGuardianOk() (*[]string, bool)`

GetEmergencyGuardianOk returns a tuple with the EmergencyGuardian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyGuardian

`func (o *TokenizationERC20FundTokenPermissionParams) SetEmergencyGuardian(v []string)`

SetEmergencyGuardian sets EmergencyGuardian field to given value.

### HasEmergencyGuardian

`func (o *TokenizationERC20FundTokenPermissionParams) HasEmergencyGuardian() bool`

HasEmergencyGuardian returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



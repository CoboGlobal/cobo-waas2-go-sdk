# PoolDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique protocol id. | 
**ChainId** | **string** | The unique chain id. | 
**Protocol** | **string** | The name of the protocol. | 
**ProtocolIconUrl** | **string** | The URL of the protocol&#39;s icon. | 
**SupportedWalletTypes** | [**[]WalletType**](WalletType.md) | The list of available wallet types. | 
**SupportedWalletSubtypes** | [**[]WalletSubtype**](WalletSubtype.md) | The list of available wallet types. | 
**TokenId** | **string** | The unique token id. | 
**EstApr** | **float32** | The estimated APR. | 
**PoolType** | Pointer to [**StakingPoolType**](StakingPoolType.md) |  | [optional] 
**MinAmount** | Pointer to **string** | The minimum amount to stake. | [optional] 
**MaxAmount** | Pointer to **string** | The maximum amount to stake. | [optional] 
**MinStakePeriod** | Pointer to **int32** | The minimum staking period in days. | [optional] 
**MaxStakePeriod** | Pointer to **int32** | The maximum staking period in days. | [optional] 
**MinStakeBlocks** | Pointer to **int64** | The minimum staking blocks. | [optional] 
**MaxStakeBlocks** | Pointer to **int64** | The maximum staking blocks. | [optional] 
**ValidatorsInfo** | [**[]PoolDetailsAllOfValidatorsInfo**](PoolDetailsAllOfValidatorsInfo.md) | The list of validators. | 

## Methods

### NewPoolDetails

`func NewPoolDetails(id string, chainId string, protocol string, protocolIconUrl string, supportedWalletTypes []WalletType, supportedWalletSubtypes []WalletSubtype, tokenId string, estApr float32, validatorsInfo []PoolDetailsAllOfValidatorsInfo, ) *PoolDetails`

NewPoolDetails instantiates a new PoolDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDetailsWithDefaults

`func NewPoolDetailsWithDefaults() *PoolDetails`

NewPoolDetailsWithDefaults instantiates a new PoolDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolDetails) SetId(v string)`

SetId sets Id field to given value.


### GetChainId

`func (o *PoolDetails) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PoolDetails) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PoolDetails) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetProtocol

`func (o *PoolDetails) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PoolDetails) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PoolDetails) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetProtocolIconUrl

`func (o *PoolDetails) GetProtocolIconUrl() string`

GetProtocolIconUrl returns the ProtocolIconUrl field if non-nil, zero value otherwise.

### GetProtocolIconUrlOk

`func (o *PoolDetails) GetProtocolIconUrlOk() (*string, bool)`

GetProtocolIconUrlOk returns a tuple with the ProtocolIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolIconUrl

`func (o *PoolDetails) SetProtocolIconUrl(v string)`

SetProtocolIconUrl sets ProtocolIconUrl field to given value.


### GetSupportedWalletTypes

`func (o *PoolDetails) GetSupportedWalletTypes() []WalletType`

GetSupportedWalletTypes returns the SupportedWalletTypes field if non-nil, zero value otherwise.

### GetSupportedWalletTypesOk

`func (o *PoolDetails) GetSupportedWalletTypesOk() (*[]WalletType, bool)`

GetSupportedWalletTypesOk returns a tuple with the SupportedWalletTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedWalletTypes

`func (o *PoolDetails) SetSupportedWalletTypes(v []WalletType)`

SetSupportedWalletTypes sets SupportedWalletTypes field to given value.


### GetSupportedWalletSubtypes

`func (o *PoolDetails) GetSupportedWalletSubtypes() []WalletSubtype`

GetSupportedWalletSubtypes returns the SupportedWalletSubtypes field if non-nil, zero value otherwise.

### GetSupportedWalletSubtypesOk

`func (o *PoolDetails) GetSupportedWalletSubtypesOk() (*[]WalletSubtype, bool)`

GetSupportedWalletSubtypesOk returns a tuple with the SupportedWalletSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedWalletSubtypes

`func (o *PoolDetails) SetSupportedWalletSubtypes(v []WalletSubtype)`

SetSupportedWalletSubtypes sets SupportedWalletSubtypes field to given value.


### GetTokenId

`func (o *PoolDetails) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PoolDetails) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PoolDetails) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetEstApr

`func (o *PoolDetails) GetEstApr() float32`

GetEstApr returns the EstApr field if non-nil, zero value otherwise.

### GetEstAprOk

`func (o *PoolDetails) GetEstAprOk() (*float32, bool)`

GetEstAprOk returns a tuple with the EstApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstApr

`func (o *PoolDetails) SetEstApr(v float32)`

SetEstApr sets EstApr field to given value.


### GetPoolType

`func (o *PoolDetails) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *PoolDetails) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *PoolDetails) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *PoolDetails) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetMinAmount

`func (o *PoolDetails) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *PoolDetails) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *PoolDetails) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *PoolDetails) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetMaxAmount

`func (o *PoolDetails) GetMaxAmount() string`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *PoolDetails) GetMaxAmountOk() (*string, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *PoolDetails) SetMaxAmount(v string)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *PoolDetails) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetMinStakePeriod

`func (o *PoolDetails) GetMinStakePeriod() int32`

GetMinStakePeriod returns the MinStakePeriod field if non-nil, zero value otherwise.

### GetMinStakePeriodOk

`func (o *PoolDetails) GetMinStakePeriodOk() (*int32, bool)`

GetMinStakePeriodOk returns a tuple with the MinStakePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStakePeriod

`func (o *PoolDetails) SetMinStakePeriod(v int32)`

SetMinStakePeriod sets MinStakePeriod field to given value.

### HasMinStakePeriod

`func (o *PoolDetails) HasMinStakePeriod() bool`

HasMinStakePeriod returns a boolean if a field has been set.

### GetMaxStakePeriod

`func (o *PoolDetails) GetMaxStakePeriod() int32`

GetMaxStakePeriod returns the MaxStakePeriod field if non-nil, zero value otherwise.

### GetMaxStakePeriodOk

`func (o *PoolDetails) GetMaxStakePeriodOk() (*int32, bool)`

GetMaxStakePeriodOk returns a tuple with the MaxStakePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStakePeriod

`func (o *PoolDetails) SetMaxStakePeriod(v int32)`

SetMaxStakePeriod sets MaxStakePeriod field to given value.

### HasMaxStakePeriod

`func (o *PoolDetails) HasMaxStakePeriod() bool`

HasMaxStakePeriod returns a boolean if a field has been set.

### GetMinStakeBlocks

`func (o *PoolDetails) GetMinStakeBlocks() int64`

GetMinStakeBlocks returns the MinStakeBlocks field if non-nil, zero value otherwise.

### GetMinStakeBlocksOk

`func (o *PoolDetails) GetMinStakeBlocksOk() (*int64, bool)`

GetMinStakeBlocksOk returns a tuple with the MinStakeBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStakeBlocks

`func (o *PoolDetails) SetMinStakeBlocks(v int64)`

SetMinStakeBlocks sets MinStakeBlocks field to given value.

### HasMinStakeBlocks

`func (o *PoolDetails) HasMinStakeBlocks() bool`

HasMinStakeBlocks returns a boolean if a field has been set.

### GetMaxStakeBlocks

`func (o *PoolDetails) GetMaxStakeBlocks() int64`

GetMaxStakeBlocks returns the MaxStakeBlocks field if non-nil, zero value otherwise.

### GetMaxStakeBlocksOk

`func (o *PoolDetails) GetMaxStakeBlocksOk() (*int64, bool)`

GetMaxStakeBlocksOk returns a tuple with the MaxStakeBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStakeBlocks

`func (o *PoolDetails) SetMaxStakeBlocks(v int64)`

SetMaxStakeBlocks sets MaxStakeBlocks field to given value.

### HasMaxStakeBlocks

`func (o *PoolDetails) HasMaxStakeBlocks() bool`

HasMaxStakeBlocks returns a boolean if a field has been set.

### GetValidatorsInfo

`func (o *PoolDetails) GetValidatorsInfo() []PoolDetailsAllOfValidatorsInfo`

GetValidatorsInfo returns the ValidatorsInfo field if non-nil, zero value otherwise.

### GetValidatorsInfoOk

`func (o *PoolDetails) GetValidatorsInfoOk() (*[]PoolDetailsAllOfValidatorsInfo, bool)`

GetValidatorsInfoOk returns a tuple with the ValidatorsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorsInfo

`func (o *PoolDetails) SetValidatorsInfo(v []PoolDetailsAllOfValidatorsInfo)`

SetValidatorsInfo sets ValidatorsInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



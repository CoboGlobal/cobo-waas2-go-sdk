# BabylonValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**IconUrl** | **string** | The URL of the validator&#39;s icon. | 
**Name** | **string** | The validator&#39;s name. | 
**Priority** | Pointer to **int32** | This property can be ignored. | [optional] 
**PublicKey** | **string** | The public key of the validator. | 
**CommissionRate** | **float32** | The commission rate of the validator. | 
**SupportedPosChains** | **[]string** | A list of supported Proof-of-Stake (PoS) chains. | 

## Methods

### NewBabylonValidator

`func NewBabylonValidator(poolType StakingPoolType, iconUrl string, name string, publicKey string, commissionRate float32, supportedPosChains []string, ) *BabylonValidator`

NewBabylonValidator instantiates a new BabylonValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonValidatorWithDefaults

`func NewBabylonValidatorWithDefaults() *BabylonValidator`

NewBabylonValidatorWithDefaults instantiates a new BabylonValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *BabylonValidator) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *BabylonValidator) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *BabylonValidator) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetIconUrl

`func (o *BabylonValidator) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *BabylonValidator) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *BabylonValidator) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetName

`func (o *BabylonValidator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BabylonValidator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BabylonValidator) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *BabylonValidator) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BabylonValidator) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BabylonValidator) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BabylonValidator) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublicKey

`func (o *BabylonValidator) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *BabylonValidator) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *BabylonValidator) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetCommissionRate

`func (o *BabylonValidator) GetCommissionRate() float32`

GetCommissionRate returns the CommissionRate field if non-nil, zero value otherwise.

### GetCommissionRateOk

`func (o *BabylonValidator) GetCommissionRateOk() (*float32, bool)`

GetCommissionRateOk returns a tuple with the CommissionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionRate

`func (o *BabylonValidator) SetCommissionRate(v float32)`

SetCommissionRate sets CommissionRate field to given value.


### GetSupportedPosChains

`func (o *BabylonValidator) GetSupportedPosChains() []string`

GetSupportedPosChains returns the SupportedPosChains field if non-nil, zero value otherwise.

### GetSupportedPosChainsOk

`func (o *BabylonValidator) GetSupportedPosChainsOk() (*[]string, bool)`

GetSupportedPosChainsOk returns a tuple with the SupportedPosChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPosChains

`func (o *BabylonValidator) SetSupportedPosChains(v []string)`

SetSupportedPosChains sets SupportedPosChains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



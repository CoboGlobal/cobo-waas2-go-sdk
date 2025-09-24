# BabylonCreateStakingExpansion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StakingId** | **string** | The ID of the Phase-2 BTC staking position. | 
**FinalityProviderPublicKeys** | **[]string** | The public keys of the finality providers(each key for a BSN chain). | 
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewBabylonCreateStakingExpansion

`func NewBabylonCreateStakingExpansion(stakingId string, finalityProviderPublicKeys []string, fee TransactionRequestFee, ) *BabylonCreateStakingExpansion`

NewBabylonCreateStakingExpansion instantiates a new BabylonCreateStakingExpansion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonCreateStakingExpansionWithDefaults

`func NewBabylonCreateStakingExpansionWithDefaults() *BabylonCreateStakingExpansion`

NewBabylonCreateStakingExpansionWithDefaults instantiates a new BabylonCreateStakingExpansion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStakingId

`func (o *BabylonCreateStakingExpansion) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *BabylonCreateStakingExpansion) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *BabylonCreateStakingExpansion) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetFinalityProviderPublicKeys

`func (o *BabylonCreateStakingExpansion) GetFinalityProviderPublicKeys() []string`

GetFinalityProviderPublicKeys returns the FinalityProviderPublicKeys field if non-nil, zero value otherwise.

### GetFinalityProviderPublicKeysOk

`func (o *BabylonCreateStakingExpansion) GetFinalityProviderPublicKeysOk() (*[]string, bool)`

GetFinalityProviderPublicKeysOk returns a tuple with the FinalityProviderPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalityProviderPublicKeys

`func (o *BabylonCreateStakingExpansion) SetFinalityProviderPublicKeys(v []string)`

SetFinalityProviderPublicKeys sets FinalityProviderPublicKeys field to given value.


### GetRequestId

`func (o *BabylonCreateStakingExpansion) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BabylonCreateStakingExpansion) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BabylonCreateStakingExpansion) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *BabylonCreateStakingExpansion) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetFee

`func (o *BabylonCreateStakingExpansion) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *BabylonCreateStakingExpansion) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *BabylonCreateStakingExpansion) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



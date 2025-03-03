# BabylonEligibleAirdrop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtcAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**BabylonAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**BabylonPoints** | Pointer to **string** | The current Babylon points balance accumulated by the BTC address. | [optional] 
**AirdropAmount** | Pointer to **string** | The estimated airdrop amount based on the current Babylon points balance. | [optional] 
**Status** | Pointer to [**BabylonRegistrationStatus**](BabylonRegistrationStatus.md) |  | [optional] 
**Pop** | Pointer to [**BabylonAirdropPop**](BabylonAirdropPop.md) |  | [optional] 

## Methods

### NewBabylonEligibleAirdrop

`func NewBabylonEligibleAirdrop() *BabylonEligibleAirdrop`

NewBabylonEligibleAirdrop instantiates a new BabylonEligibleAirdrop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBabylonEligibleAirdropWithDefaults

`func NewBabylonEligibleAirdropWithDefaults() *BabylonEligibleAirdrop`

NewBabylonEligibleAirdropWithDefaults instantiates a new BabylonEligibleAirdrop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtcAddress

`func (o *BabylonEligibleAirdrop) GetBtcAddress() StakingSource`

GetBtcAddress returns the BtcAddress field if non-nil, zero value otherwise.

### GetBtcAddressOk

`func (o *BabylonEligibleAirdrop) GetBtcAddressOk() (*StakingSource, bool)`

GetBtcAddressOk returns a tuple with the BtcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtcAddress

`func (o *BabylonEligibleAirdrop) SetBtcAddress(v StakingSource)`

SetBtcAddress sets BtcAddress field to given value.

### HasBtcAddress

`func (o *BabylonEligibleAirdrop) HasBtcAddress() bool`

HasBtcAddress returns a boolean if a field has been set.

### GetBabylonAddress

`func (o *BabylonEligibleAirdrop) GetBabylonAddress() StakingSource`

GetBabylonAddress returns the BabylonAddress field if non-nil, zero value otherwise.

### GetBabylonAddressOk

`func (o *BabylonEligibleAirdrop) GetBabylonAddressOk() (*StakingSource, bool)`

GetBabylonAddressOk returns a tuple with the BabylonAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddress

`func (o *BabylonEligibleAirdrop) SetBabylonAddress(v StakingSource)`

SetBabylonAddress sets BabylonAddress field to given value.

### HasBabylonAddress

`func (o *BabylonEligibleAirdrop) HasBabylonAddress() bool`

HasBabylonAddress returns a boolean if a field has been set.

### GetBabylonPoints

`func (o *BabylonEligibleAirdrop) GetBabylonPoints() string`

GetBabylonPoints returns the BabylonPoints field if non-nil, zero value otherwise.

### GetBabylonPointsOk

`func (o *BabylonEligibleAirdrop) GetBabylonPointsOk() (*string, bool)`

GetBabylonPointsOk returns a tuple with the BabylonPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonPoints

`func (o *BabylonEligibleAirdrop) SetBabylonPoints(v string)`

SetBabylonPoints sets BabylonPoints field to given value.

### HasBabylonPoints

`func (o *BabylonEligibleAirdrop) HasBabylonPoints() bool`

HasBabylonPoints returns a boolean if a field has been set.

### GetAirdropAmount

`func (o *BabylonEligibleAirdrop) GetAirdropAmount() string`

GetAirdropAmount returns the AirdropAmount field if non-nil, zero value otherwise.

### GetAirdropAmountOk

`func (o *BabylonEligibleAirdrop) GetAirdropAmountOk() (*string, bool)`

GetAirdropAmountOk returns a tuple with the AirdropAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirdropAmount

`func (o *BabylonEligibleAirdrop) SetAirdropAmount(v string)`

SetAirdropAmount sets AirdropAmount field to given value.

### HasAirdropAmount

`func (o *BabylonEligibleAirdrop) HasAirdropAmount() bool`

HasAirdropAmount returns a boolean if a field has been set.

### GetStatus

`func (o *BabylonEligibleAirdrop) GetStatus() BabylonRegistrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BabylonEligibleAirdrop) GetStatusOk() (*BabylonRegistrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BabylonEligibleAirdrop) SetStatus(v BabylonRegistrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BabylonEligibleAirdrop) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPop

`func (o *BabylonEligibleAirdrop) GetPop() BabylonAirdropPop`

GetPop returns the Pop field if non-nil, zero value otherwise.

### GetPopOk

`func (o *BabylonEligibleAirdrop) GetPopOk() (*BabylonAirdropPop, bool)`

GetPopOk returns a tuple with the Pop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPop

`func (o *BabylonEligibleAirdrop) SetPop(v BabylonAirdropPop)`

SetPop sets Pop field to given value.

### HasPop

`func (o *BabylonEligibleAirdrop) HasPop() bool`

HasPop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



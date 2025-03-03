# ListBabylonEligibleStakings200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StakingId** | Pointer to **string** | The ID of the Phase-1 BTC staking position. | [optional] 
**BtcAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**BabylonAddress** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**StakedAmount** | Pointer to **string** | The current amount of BTC staked. | [optional] 
**Status** | Pointer to [**BabylonRegistrationStatus**](BabylonRegistrationStatus.md) |  | [optional] 

## Methods

### NewListBabylonEligibleStakings200ResponseDataInner

`func NewListBabylonEligibleStakings200ResponseDataInner() *ListBabylonEligibleStakings200ResponseDataInner`

NewListBabylonEligibleStakings200ResponseDataInner instantiates a new ListBabylonEligibleStakings200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBabylonEligibleStakings200ResponseDataInnerWithDefaults

`func NewListBabylonEligibleStakings200ResponseDataInnerWithDefaults() *ListBabylonEligibleStakings200ResponseDataInner`

NewListBabylonEligibleStakings200ResponseDataInnerWithDefaults instantiates a new ListBabylonEligibleStakings200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStakingId

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *ListBabylonEligibleStakings200ResponseDataInner) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.

### HasStakingId

`func (o *ListBabylonEligibleStakings200ResponseDataInner) HasStakingId() bool`

HasStakingId returns a boolean if a field has been set.

### GetBtcAddress

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetBtcAddress() StakingSource`

GetBtcAddress returns the BtcAddress field if non-nil, zero value otherwise.

### GetBtcAddressOk

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetBtcAddressOk() (*StakingSource, bool)`

GetBtcAddressOk returns a tuple with the BtcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtcAddress

`func (o *ListBabylonEligibleStakings200ResponseDataInner) SetBtcAddress(v StakingSource)`

SetBtcAddress sets BtcAddress field to given value.

### HasBtcAddress

`func (o *ListBabylonEligibleStakings200ResponseDataInner) HasBtcAddress() bool`

HasBtcAddress returns a boolean if a field has been set.

### GetBabylonAddress

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetBabylonAddress() StakingSource`

GetBabylonAddress returns the BabylonAddress field if non-nil, zero value otherwise.

### GetBabylonAddressOk

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetBabylonAddressOk() (*StakingSource, bool)`

GetBabylonAddressOk returns a tuple with the BabylonAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBabylonAddress

`func (o *ListBabylonEligibleStakings200ResponseDataInner) SetBabylonAddress(v StakingSource)`

SetBabylonAddress sets BabylonAddress field to given value.

### HasBabylonAddress

`func (o *ListBabylonEligibleStakings200ResponseDataInner) HasBabylonAddress() bool`

HasBabylonAddress returns a boolean if a field has been set.

### GetStakedAmount

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetStakedAmount() string`

GetStakedAmount returns the StakedAmount field if non-nil, zero value otherwise.

### GetStakedAmountOk

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetStakedAmountOk() (*string, bool)`

GetStakedAmountOk returns a tuple with the StakedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakedAmount

`func (o *ListBabylonEligibleStakings200ResponseDataInner) SetStakedAmount(v string)`

SetStakedAmount sets StakedAmount field to given value.

### HasStakedAmount

`func (o *ListBabylonEligibleStakings200ResponseDataInner) HasStakedAmount() bool`

HasStakedAmount returns a boolean if a field has been set.

### GetStatus

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetStatus() BabylonRegistrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBabylonEligibleStakings200ResponseDataInner) GetStatusOk() (*BabylonRegistrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBabylonEligibleStakings200ResponseDataInner) SetStatus(v BabylonRegistrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBabylonEligibleStakings200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



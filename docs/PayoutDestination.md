# PayoutDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The destination address. | 
**Amount** | **string** | The amount of cryptocurrency to send to the destination. | 
**Email** | Pointer to **string** | The email of the destination. | [optional] 
**Note** | Pointer to **string** | The note of the destination. | [optional] 
**LoopEnabled** | Pointer to **bool** | Enable loop mode for standard transfers when source is asset wallet. Only applicable when: - &#x60;payout_mode&#x60; is &#x60;Normal&#x60; - &#x60;source_type&#x60; is asset wallet  | [optional] [default to false]
**NetworkFee** | Pointer to [**PayoutFeeData**](PayoutFeeData.md) |  | [optional] 

## Methods

### NewPayoutDestination

`func NewPayoutDestination(address string, amount string, ) *PayoutDestination`

NewPayoutDestination instantiates a new PayoutDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutDestinationWithDefaults

`func NewPayoutDestinationWithDefaults() *PayoutDestination`

NewPayoutDestinationWithDefaults instantiates a new PayoutDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PayoutDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayoutDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayoutDestination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *PayoutDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayoutDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayoutDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetEmail

`func (o *PayoutDestination) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayoutDestination) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayoutDestination) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayoutDestination) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNote

`func (o *PayoutDestination) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PayoutDestination) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PayoutDestination) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PayoutDestination) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetLoopEnabled

`func (o *PayoutDestination) GetLoopEnabled() bool`

GetLoopEnabled returns the LoopEnabled field if non-nil, zero value otherwise.

### GetLoopEnabledOk

`func (o *PayoutDestination) GetLoopEnabledOk() (*bool, bool)`

GetLoopEnabledOk returns a tuple with the LoopEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopEnabled

`func (o *PayoutDestination) SetLoopEnabled(v bool)`

SetLoopEnabled sets LoopEnabled field to given value.

### HasLoopEnabled

`func (o *PayoutDestination) HasLoopEnabled() bool`

HasLoopEnabled returns a boolean if a field has been set.

### GetNetworkFee

`func (o *PayoutDestination) GetNetworkFee() PayoutFeeData`

GetNetworkFee returns the NetworkFee field if non-nil, zero value otherwise.

### GetNetworkFeeOk

`func (o *PayoutDestination) GetNetworkFeeOk() (*PayoutFeeData, bool)`

GetNetworkFeeOk returns a tuple with the NetworkFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFee

`func (o *PayoutDestination) SetNetworkFee(v PayoutFeeData)`

SetNetworkFee sets NetworkFee field to given value.

### HasNetworkFee

`func (o *PayoutDestination) HasNetworkFee() bool`

HasNetworkFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AmountDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**AmountStatus**](AmountStatus.md) |  | 
**Amount** | **string** | The staking amount. | 

## Methods

### NewAmountDetailsInner

`func NewAmountDetailsInner(status AmountStatus, amount string, ) *AmountDetailsInner`

NewAmountDetailsInner instantiates a new AmountDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountDetailsInnerWithDefaults

`func NewAmountDetailsInnerWithDefaults() *AmountDetailsInner`

NewAmountDetailsInnerWithDefaults instantiates a new AmountDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AmountDetailsInner) GetStatus() AmountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AmountDetailsInner) GetStatusOk() (*AmountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AmountDetailsInner) SetStatus(v AmountStatus)`

SetStatus sets Status field to given value.


### GetAmount

`func (o *AmountDetailsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AmountDetailsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AmountDetailsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



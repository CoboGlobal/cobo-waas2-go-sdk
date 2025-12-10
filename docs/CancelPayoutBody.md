# CancelPayoutBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkFee** | Pointer to [**PayoutFeeData**](PayoutFeeData.md) |  | [optional] 

## Methods

### NewCancelPayoutBody

`func NewCancelPayoutBody() *CancelPayoutBody`

NewCancelPayoutBody instantiates a new CancelPayoutBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelPayoutBodyWithDefaults

`func NewCancelPayoutBodyWithDefaults() *CancelPayoutBody`

NewCancelPayoutBodyWithDefaults instantiates a new CancelPayoutBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkFee

`func (o *CancelPayoutBody) GetNetworkFee() PayoutFeeData`

GetNetworkFee returns the NetworkFee field if non-nil, zero value otherwise.

### GetNetworkFeeOk

`func (o *CancelPayoutBody) GetNetworkFeeOk() (*PayoutFeeData, bool)`

GetNetworkFeeOk returns a tuple with the NetworkFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFee

`func (o *CancelPayoutBody) SetNetworkFee(v PayoutFeeData)`

SetNetworkFee sets NetworkFee field to given value.

### HasNetworkFee

`func (o *CancelPayoutBody) HasNetworkFee() bool`

HasNetworkFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



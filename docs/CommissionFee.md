# CommissionFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeAmount** | **string** | The amount of the commission fee charged by Cobo for pay-ins and payouts, in USD. | 

## Methods

### NewCommissionFee

`func NewCommissionFee(feeAmount string, ) *CommissionFee`

NewCommissionFee instantiates a new CommissionFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionFeeWithDefaults

`func NewCommissionFeeWithDefaults() *CommissionFee`

NewCommissionFeeWithDefaults instantiates a new CommissionFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeAmount

`func (o *CommissionFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *CommissionFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *CommissionFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



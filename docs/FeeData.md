# FeeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasLimit** | Pointer to **string** | The Limit of gas. | [optional] [default to "21000"]
**FeeAmount** | Pointer to **string** | The estimated fee amount in fee_coin. | [optional] 

## Methods

### NewFeeData

`func NewFeeData() *FeeData`

NewFeeData instantiates a new FeeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeDataWithDefaults

`func NewFeeDataWithDefaults() *FeeData`

NewFeeDataWithDefaults instantiates a new FeeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasLimit

`func (o *FeeData) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *FeeData) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *FeeData) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *FeeData) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetFeeAmount

`func (o *FeeData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *FeeData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *FeeData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *FeeData) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



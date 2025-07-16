# FILPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPremium** | Pointer to **string** | An optional tip you can include to prioritize your transaction. The gas premium incentivizes miners to include your transaction sooner than those offering only the base fee. | [optional] 
**GasFeeCap** | Pointer to **string** | The maximum gas price you are willing to pay per unit of gas. | [optional] 
**GasLimit** | Pointer to **string** | The maximum amount of gas your transaction is allowed to consume. | [optional] 

## Methods

### NewFILPrice

`func NewFILPrice() *FILPrice`

NewFILPrice instantiates a new FILPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFILPriceWithDefaults

`func NewFILPriceWithDefaults() *FILPrice`

NewFILPriceWithDefaults instantiates a new FILPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPremium

`func (o *FILPrice) GetGasPremium() string`

GetGasPremium returns the GasPremium field if non-nil, zero value otherwise.

### GetGasPremiumOk

`func (o *FILPrice) GetGasPremiumOk() (*string, bool)`

GetGasPremiumOk returns a tuple with the GasPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPremium

`func (o *FILPrice) SetGasPremium(v string)`

SetGasPremium sets GasPremium field to given value.

### HasGasPremium

`func (o *FILPrice) HasGasPremium() bool`

HasGasPremium returns a boolean if a field has been set.

### GetGasFeeCap

`func (o *FILPrice) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *FILPrice) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *FILPrice) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *FILPrice) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### GetGasLimit

`func (o *FILPrice) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *FILPrice) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *FILPrice) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *FILPrice) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



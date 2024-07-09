# EvmLegacyFeeBasePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPrice** | **string** | The gas price, in gwei. The gas price represents the amount of ETH that must be paid to validators for processing transactions. | 

## Methods

### NewEvmLegacyFeeBasePrice

`func NewEvmLegacyFeeBasePrice(gasPrice string, ) *EvmLegacyFeeBasePrice`

NewEvmLegacyFeeBasePrice instantiates a new EvmLegacyFeeBasePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyFeeBasePriceWithDefaults

`func NewEvmLegacyFeeBasePriceWithDefaults() *EvmLegacyFeeBasePrice`

NewEvmLegacyFeeBasePriceWithDefaults instantiates a new EvmLegacyFeeBasePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPrice

`func (o *EvmLegacyFeeBasePrice) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *EvmLegacyFeeBasePrice) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *EvmLegacyFeeBasePrice) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



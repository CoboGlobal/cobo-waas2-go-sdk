# EvmLegacyFeePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 
**GasPrice** | **string** | The Price of Gas, unit GWei. | 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]

## Methods

### NewEvmLegacyFeePrice

`func NewEvmLegacyFeePrice(gasPrice string, feeType FeeType, ) *EvmLegacyFeePrice`

NewEvmLegacyFeePrice instantiates a new EvmLegacyFeePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyFeePriceWithDefaults

`func NewEvmLegacyFeePriceWithDefaults() *EvmLegacyFeePrice`

NewEvmLegacyFeePriceWithDefaults instantiates a new EvmLegacyFeePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *EvmLegacyFeePrice) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EvmLegacyFeePrice) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EvmLegacyFeePrice) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EvmLegacyFeePrice) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetGasPrice

`func (o *EvmLegacyFeePrice) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *EvmLegacyFeePrice) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *EvmLegacyFeePrice) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetFeeType

`func (o *EvmLegacyFeePrice) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmLegacyFeePrice) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmLegacyFeePrice) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



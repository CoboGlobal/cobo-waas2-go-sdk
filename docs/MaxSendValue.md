# MaxSendValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | token name. | [optional] 
**TokenDecimal** | Pointer to **int32** | coin decimal precision. | [optional] 
**FeeToken** | Pointer to **string** | fee token name. | [optional] 
**FeeDecimal** | Pointer to **int32** | fee coin decimal precision. | [optional] 
**MaxSendValue** | Pointer to **string** | the maximum sendable value for the given address or current wallet. | [optional] 
**FeePerByte** | Pointer to **string** | transaction fees per byte for UTXO model. | [optional] 
**FeeValue** | Pointer to **string** | transaction fee for UTXO model | [optional] 
**GasPrice** | Pointer to **string** | gas price for account model | [optional] 
**GasLimit** | Pointer to **int32** | gas limit for account model | [optional] 

## Methods

### NewMaxSendValue

`func NewMaxSendValue() *MaxSendValue`

NewMaxSendValue instantiates a new MaxSendValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxSendValueWithDefaults

`func NewMaxSendValueWithDefaults() *MaxSendValue`

NewMaxSendValueWithDefaults instantiates a new MaxSendValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *MaxSendValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MaxSendValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MaxSendValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MaxSendValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenDecimal

`func (o *MaxSendValue) GetTokenDecimal() int32`

GetTokenDecimal returns the TokenDecimal field if non-nil, zero value otherwise.

### GetTokenDecimalOk

`func (o *MaxSendValue) GetTokenDecimalOk() (*int32, bool)`

GetTokenDecimalOk returns a tuple with the TokenDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimal

`func (o *MaxSendValue) SetTokenDecimal(v int32)`

SetTokenDecimal sets TokenDecimal field to given value.

### HasTokenDecimal

`func (o *MaxSendValue) HasTokenDecimal() bool`

HasTokenDecimal returns a boolean if a field has been set.

### GetFeeToken

`func (o *MaxSendValue) GetFeeToken() string`

GetFeeToken returns the FeeToken field if non-nil, zero value otherwise.

### GetFeeTokenOk

`func (o *MaxSendValue) GetFeeTokenOk() (*string, bool)`

GetFeeTokenOk returns a tuple with the FeeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeToken

`func (o *MaxSendValue) SetFeeToken(v string)`

SetFeeToken sets FeeToken field to given value.

### HasFeeToken

`func (o *MaxSendValue) HasFeeToken() bool`

HasFeeToken returns a boolean if a field has been set.

### GetFeeDecimal

`func (o *MaxSendValue) GetFeeDecimal() int32`

GetFeeDecimal returns the FeeDecimal field if non-nil, zero value otherwise.

### GetFeeDecimalOk

`func (o *MaxSendValue) GetFeeDecimalOk() (*int32, bool)`

GetFeeDecimalOk returns a tuple with the FeeDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeDecimal

`func (o *MaxSendValue) SetFeeDecimal(v int32)`

SetFeeDecimal sets FeeDecimal field to given value.

### HasFeeDecimal

`func (o *MaxSendValue) HasFeeDecimal() bool`

HasFeeDecimal returns a boolean if a field has been set.

### GetMaxSendValue

`func (o *MaxSendValue) GetMaxSendValue() string`

GetMaxSendValue returns the MaxSendValue field if non-nil, zero value otherwise.

### GetMaxSendValueOk

`func (o *MaxSendValue) GetMaxSendValueOk() (*string, bool)`

GetMaxSendValueOk returns a tuple with the MaxSendValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSendValue

`func (o *MaxSendValue) SetMaxSendValue(v string)`

SetMaxSendValue sets MaxSendValue field to given value.

### HasMaxSendValue

`func (o *MaxSendValue) HasMaxSendValue() bool`

HasMaxSendValue returns a boolean if a field has been set.

### GetFeePerByte

`func (o *MaxSendValue) GetFeePerByte() string`

GetFeePerByte returns the FeePerByte field if non-nil, zero value otherwise.

### GetFeePerByteOk

`func (o *MaxSendValue) GetFeePerByteOk() (*string, bool)`

GetFeePerByteOk returns a tuple with the FeePerByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePerByte

`func (o *MaxSendValue) SetFeePerByte(v string)`

SetFeePerByte sets FeePerByte field to given value.

### HasFeePerByte

`func (o *MaxSendValue) HasFeePerByte() bool`

HasFeePerByte returns a boolean if a field has been set.

### GetFeeValue

`func (o *MaxSendValue) GetFeeValue() string`

GetFeeValue returns the FeeValue field if non-nil, zero value otherwise.

### GetFeeValueOk

`func (o *MaxSendValue) GetFeeValueOk() (*string, bool)`

GetFeeValueOk returns a tuple with the FeeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeValue

`func (o *MaxSendValue) SetFeeValue(v string)`

SetFeeValue sets FeeValue field to given value.

### HasFeeValue

`func (o *MaxSendValue) HasFeeValue() bool`

HasFeeValue returns a boolean if a field has been set.

### GetGasPrice

`func (o *MaxSendValue) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *MaxSendValue) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *MaxSendValue) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *MaxSendValue) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasLimit

`func (o *MaxSendValue) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *MaxSendValue) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *MaxSendValue) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *MaxSendValue) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SOLComputeUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeUnitPrice** | Pointer to **string** | The price paid per compute unit. This value determines the priority fee for the transaction, allowing you to increase inclusion probability in congested conditions. | [optional] 
**ComputeUnitLimit** | Pointer to **string** | The maximum number of compute units your transaction is allowed to consume. It sets an upper bound on computational resource usage to prevent overload. | [optional] 

## Methods

### NewSOLComputeUnit

`func NewSOLComputeUnit() *SOLComputeUnit`

NewSOLComputeUnit instantiates a new SOLComputeUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSOLComputeUnitWithDefaults

`func NewSOLComputeUnitWithDefaults() *SOLComputeUnit`

NewSOLComputeUnitWithDefaults instantiates a new SOLComputeUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeUnitPrice

`func (o *SOLComputeUnit) GetComputeUnitPrice() string`

GetComputeUnitPrice returns the ComputeUnitPrice field if non-nil, zero value otherwise.

### GetComputeUnitPriceOk

`func (o *SOLComputeUnit) GetComputeUnitPriceOk() (*string, bool)`

GetComputeUnitPriceOk returns a tuple with the ComputeUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitPrice

`func (o *SOLComputeUnit) SetComputeUnitPrice(v string)`

SetComputeUnitPrice sets ComputeUnitPrice field to given value.

### HasComputeUnitPrice

`func (o *SOLComputeUnit) HasComputeUnitPrice() bool`

HasComputeUnitPrice returns a boolean if a field has been set.

### GetComputeUnitLimit

`func (o *SOLComputeUnit) GetComputeUnitLimit() string`

GetComputeUnitLimit returns the ComputeUnitLimit field if non-nil, zero value otherwise.

### GetComputeUnitLimitOk

`func (o *SOLComputeUnit) GetComputeUnitLimitOk() (*string, bool)`

GetComputeUnitLimitOk returns a tuple with the ComputeUnitLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitLimit

`func (o *SOLComputeUnit) SetComputeUnitLimit(v string)`

SetComputeUnitLimit sets ComputeUnitLimit field to given value.

### HasComputeUnitLimit

`func (o *SOLComputeUnit) HasComputeUnitLimit() bool`

HasComputeUnitLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



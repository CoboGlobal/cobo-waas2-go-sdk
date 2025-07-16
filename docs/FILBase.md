# FILBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasBase** | Pointer to **string** | The minimum fee required for a transaction to be included in a block. The base fee is dynamically adjusted based on network congestion to maintain target block utilization. It is burned rather than paid to miners, reducing the total Filecoin supply over time. | [optional] 

## Methods

### NewFILBase

`func NewFILBase() *FILBase`

NewFILBase instantiates a new FILBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFILBaseWithDefaults

`func NewFILBaseWithDefaults() *FILBase`

NewFILBaseWithDefaults instantiates a new FILBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasBase

`func (o *FILBase) GetGasBase() string`

GetGasBase returns the GasBase field if non-nil, zero value otherwise.

### GetGasBaseOk

`func (o *FILBase) GetGasBaseOk() (*string, bool)`

GetGasBaseOk returns a tuple with the GasBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasBase

`func (o *FILBase) SetGasBase(v string)`

SetGasBase sets GasBase field to given value.

### HasGasBase

`func (o *FILBase) HasGasBase() bool`

HasGasBase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FILBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasBase** | Pointer to **string** | This is the minimum fee required to include a transaction in a block. It is determined by the network&#39;s congestion level, which adjusts to maintain a target block utilization rate. The base fee is burned, reducing the total supply of Filecoin over time. | [optional] 

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



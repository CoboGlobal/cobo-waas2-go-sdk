# FeeReserved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReservedFee** | Pointer to **string** | The estimated fee required for submitting the transaction data to L1 (Layer 1), measured in wei. | [optional] 

## Methods

### NewFeeReserved

`func NewFeeReserved() *FeeReserved`

NewFeeReserved instantiates a new FeeReserved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeReservedWithDefaults

`func NewFeeReservedWithDefaults() *FeeReserved`

NewFeeReservedWithDefaults instantiates a new FeeReserved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservedFee

`func (o *FeeReserved) GetReservedFee() string`

GetReservedFee returns the ReservedFee field if non-nil, zero value otherwise.

### GetReservedFeeOk

`func (o *FeeReserved) GetReservedFeeOk() (*string, bool)`

GetReservedFeeOk returns a tuple with the ReservedFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedFee

`func (o *FeeReserved) SetReservedFee(v string)`

SetReservedFee sets ReservedFee field to given value.

### HasReservedFee

`func (o *FeeReserved) HasReservedFee() bool`

HasReservedFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



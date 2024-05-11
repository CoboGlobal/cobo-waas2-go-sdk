# FixedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeAmount** | Pointer to **string** | The estimated fee amount in fee_coin. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 

## Methods

### NewFixedFee

`func NewFixedFee(feeType FeeType, ) *FixedFee`

NewFixedFee instantiates a new FixedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedFeeWithDefaults

`func NewFixedFeeWithDefaults() *FixedFee`

NewFixedFeeWithDefaults instantiates a new FixedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeAmount

`func (o *FixedFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *FixedFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *FixedFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *FixedFee) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeeType

`func (o *FixedFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *FixedFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *FixedFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetFeeTokenId

`func (o *FixedFee) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *FixedFee) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *FixedFee) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *FixedFee) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



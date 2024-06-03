# MaxTransferableValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **string** | The token name. | [optional] 
**MaxTransferableValue** | Pointer to **string** | the maximum transferable value for the given address or current wallet. | [optional] 
**Fee** | Pointer to [**TransactionFee**](TransactionFee.md) |  | [optional] 

## Methods

### NewMaxTransferableValue

`func NewMaxTransferableValue() *MaxTransferableValue`

NewMaxTransferableValue instantiates a new MaxTransferableValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxTransferableValueWithDefaults

`func NewMaxTransferableValueWithDefaults() *MaxTransferableValue`

NewMaxTransferableValueWithDefaults instantiates a new MaxTransferableValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *MaxTransferableValue) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *MaxTransferableValue) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *MaxTransferableValue) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *MaxTransferableValue) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetMaxTransferableValue

`func (o *MaxTransferableValue) GetMaxTransferableValue() string`

GetMaxTransferableValue returns the MaxTransferableValue field if non-nil, zero value otherwise.

### GetMaxTransferableValueOk

`func (o *MaxTransferableValue) GetMaxTransferableValueOk() (*string, bool)`

GetMaxTransferableValueOk returns a tuple with the MaxTransferableValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransferableValue

`func (o *MaxTransferableValue) SetMaxTransferableValue(v string)`

SetMaxTransferableValue sets MaxTransferableValue field to given value.

### HasMaxTransferableValue

`func (o *MaxTransferableValue) HasMaxTransferableValue() bool`

HasMaxTransferableValue returns a boolean if a field has been set.

### GetFee

`func (o *MaxTransferableValue) GetFee() TransactionFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *MaxTransferableValue) GetFeeOk() (*TransactionFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *MaxTransferableValue) SetFee(v TransactionFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *MaxTransferableValue) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


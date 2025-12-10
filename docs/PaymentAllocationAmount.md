# PaymentAllocationAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The ID of the cryptocurrency. | 
**AllocationAmount** | **string** | The allocation amount. | 

## Methods

### NewPaymentAllocationAmount

`func NewPaymentAllocationAmount(tokenId string, allocationAmount string, ) *PaymentAllocationAmount`

NewPaymentAllocationAmount instantiates a new PaymentAllocationAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAllocationAmountWithDefaults

`func NewPaymentAllocationAmountWithDefaults() *PaymentAllocationAmount`

NewPaymentAllocationAmountWithDefaults instantiates a new PaymentAllocationAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *PaymentAllocationAmount) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentAllocationAmount) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentAllocationAmount) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAllocationAmount

`func (o *PaymentAllocationAmount) GetAllocationAmount() string`

GetAllocationAmount returns the AllocationAmount field if non-nil, zero value otherwise.

### GetAllocationAmountOk

`func (o *PaymentAllocationAmount) GetAllocationAmountOk() (*string, bool)`

GetAllocationAmountOk returns a tuple with the AllocationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationAmount

`func (o *PaymentAllocationAmount) SetAllocationAmount(v string)`

SetAllocationAmount sets AllocationAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



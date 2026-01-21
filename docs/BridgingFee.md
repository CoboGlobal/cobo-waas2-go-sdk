# BridgingFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeAmount** | **string** | The fee charged for bridging tokens to another chain.  | 
**ReceivedTokenId** | Pointer to **string** | The ID of the destination token received after bridging. | [optional] 
**ReceivedAmount** | Pointer to **string** | The final amount of the token received after bridging. | [optional] 

## Methods

### NewBridgingFee

`func NewBridgingFee(feeAmount string, ) *BridgingFee`

NewBridgingFee instantiates a new BridgingFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgingFeeWithDefaults

`func NewBridgingFeeWithDefaults() *BridgingFee`

NewBridgingFeeWithDefaults instantiates a new BridgingFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeAmount

`func (o *BridgingFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *BridgingFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *BridgingFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetReceivedTokenId

`func (o *BridgingFee) GetReceivedTokenId() string`

GetReceivedTokenId returns the ReceivedTokenId field if non-nil, zero value otherwise.

### GetReceivedTokenIdOk

`func (o *BridgingFee) GetReceivedTokenIdOk() (*string, bool)`

GetReceivedTokenIdOk returns a tuple with the ReceivedTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTokenId

`func (o *BridgingFee) SetReceivedTokenId(v string)`

SetReceivedTokenId sets ReceivedTokenId field to given value.

### HasReceivedTokenId

`func (o *BridgingFee) HasReceivedTokenId() bool`

HasReceivedTokenId returns a boolean if a field has been set.

### GetReceivedAmount

`func (o *BridgingFee) GetReceivedAmount() string`

GetReceivedAmount returns the ReceivedAmount field if non-nil, zero value otherwise.

### GetReceivedAmountOk

`func (o *BridgingFee) GetReceivedAmountOk() (*string, bool)`

GetReceivedAmountOk returns a tuple with the ReceivedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAmount

`func (o *BridgingFee) SetReceivedAmount(v string)`

SetReceivedAmount sets ReceivedAmount field to given value.

### HasReceivedAmount

`func (o *BridgingFee) HasReceivedAmount() bool`

HasReceivedAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



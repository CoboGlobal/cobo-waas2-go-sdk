# PaymentPayoutItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID of the payout item. | 
**Amount** | **string** | The amount of the payout item.  | 
**BridgingFee** | Pointer to [**BridgingFee**](BridgingFee.md) |  | [optional] 

## Methods

### NewPaymentPayoutItem

`func NewPaymentPayoutItem(tokenId string, amount string, ) *PaymentPayoutItem`

NewPaymentPayoutItem instantiates a new PaymentPayoutItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutItemWithDefaults

`func NewPaymentPayoutItemWithDefaults() *PaymentPayoutItem`

NewPaymentPayoutItemWithDefaults instantiates a new PaymentPayoutItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *PaymentPayoutItem) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentPayoutItem) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentPayoutItem) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *PaymentPayoutItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentPayoutItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentPayoutItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBridgingFee

`func (o *PaymentPayoutItem) GetBridgingFee() BridgingFee`

GetBridgingFee returns the BridgingFee field if non-nil, zero value otherwise.

### GetBridgingFeeOk

`func (o *PaymentPayoutItem) GetBridgingFeeOk() (*BridgingFee, bool)`

GetBridgingFeeOk returns a tuple with the BridgingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgingFee

`func (o *PaymentPayoutItem) SetBridgingFee(v BridgingFee)`

SetBridgingFee sets BridgingFee field to given value.

### HasBridgingFee

`func (o *PaymentPayoutItem) HasBridgingFee() bool`

HasBridgingFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



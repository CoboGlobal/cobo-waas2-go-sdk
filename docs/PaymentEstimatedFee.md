# PaymentEstimatedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID for which fees will be calculated. | 
**Amount** | **string** | The transaction amount for which fees will be calculated. | 
**CommissionFee** | Pointer to [**CommissionFee**](CommissionFee.md) |  | [optional] 
**BridgingFee** | Pointer to [**BridgingFee**](BridgingFee.md) |  | [optional] 
**OtcFee** | Pointer to [**OtcFee**](OtcFee.md) |  | [optional] 

## Methods

### NewPaymentEstimatedFee

`func NewPaymentEstimatedFee(tokenId string, amount string, ) *PaymentEstimatedFee`

NewPaymentEstimatedFee instantiates a new PaymentEstimatedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEstimatedFeeWithDefaults

`func NewPaymentEstimatedFeeWithDefaults() *PaymentEstimatedFee`

NewPaymentEstimatedFeeWithDefaults instantiates a new PaymentEstimatedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *PaymentEstimatedFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentEstimatedFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentEstimatedFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *PaymentEstimatedFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentEstimatedFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentEstimatedFee) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCommissionFee

`func (o *PaymentEstimatedFee) GetCommissionFee() CommissionFee`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *PaymentEstimatedFee) GetCommissionFeeOk() (*CommissionFee, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *PaymentEstimatedFee) SetCommissionFee(v CommissionFee)`

SetCommissionFee sets CommissionFee field to given value.

### HasCommissionFee

`func (o *PaymentEstimatedFee) HasCommissionFee() bool`

HasCommissionFee returns a boolean if a field has been set.

### GetBridgingFee

`func (o *PaymentEstimatedFee) GetBridgingFee() BridgingFee`

GetBridgingFee returns the BridgingFee field if non-nil, zero value otherwise.

### GetBridgingFeeOk

`func (o *PaymentEstimatedFee) GetBridgingFeeOk() (*BridgingFee, bool)`

GetBridgingFeeOk returns a tuple with the BridgingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgingFee

`func (o *PaymentEstimatedFee) SetBridgingFee(v BridgingFee)`

SetBridgingFee sets BridgingFee field to given value.

### HasBridgingFee

`func (o *PaymentEstimatedFee) HasBridgingFee() bool`

HasBridgingFee returns a boolean if a field has been set.

### GetOtcFee

`func (o *PaymentEstimatedFee) GetOtcFee() OtcFee`

GetOtcFee returns the OtcFee field if non-nil, zero value otherwise.

### GetOtcFeeOk

`func (o *PaymentEstimatedFee) GetOtcFeeOk() (*OtcFee, bool)`

GetOtcFeeOk returns a tuple with the OtcFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtcFee

`func (o *PaymentEstimatedFee) SetOtcFee(v OtcFee)`

SetOtcFee sets OtcFee field to given value.

### HasOtcFee

`func (o *PaymentEstimatedFee) HasOtcFee() bool`

HasOtcFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



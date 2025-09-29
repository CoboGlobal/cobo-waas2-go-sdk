# CreateRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a refund request. The request ID is provided by you and must be unique. | 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**PayableAmount** | **string** | The amount to refund in cryptocurrency. | 
**ToAddress** | Pointer to **string** | The address where the refunded cryptocurrency will be sent. | [optional] 
**TokenId** | **string** | The ID of the cryptocurrency used for refund. Supported values:    - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60;   - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**RefundType** | [**RefundType**](RefundType.md) |  | 
**OrderId** | Pointer to **string** | The ID of the original pay-in order associated with this refund. Use this to track refunds against specific payments. | [optional] 
**ChargeMerchantFee** | Pointer to **bool** | Indicates whether the merchant should bear the transaction fee for the refund.  If true, the fee will be deducted from merchant&#39;s account balance.  | [optional] 
**MerchantFeeAmount** | Pointer to **string** | The amount of the transaction fee that the merchant will bear for the refund.  This is only applicable if &#x60;charge_merchant_fee&#x60; is set to true.  | [optional] 
**MerchantFeeTokenId** | Pointer to **string** | The ID of the cryptocurrency used for the transaction fee.  This is only applicable if &#x60;charge_merchant_fee&#x60; is set to true.  | [optional] 

## Methods

### NewCreateRefundRequest

`func NewCreateRefundRequest(requestId string, payableAmount string, tokenId string, refundType RefundType, ) *CreateRefundRequest`

NewCreateRefundRequest instantiates a new CreateRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefundRequestWithDefaults

`func NewCreateRefundRequestWithDefaults() *CreateRefundRequest`

NewCreateRefundRequestWithDefaults instantiates a new CreateRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateRefundRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateRefundRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateRefundRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetMerchantId

`func (o *CreateRefundRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateRefundRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateRefundRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *CreateRefundRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayableAmount

`func (o *CreateRefundRequest) GetPayableAmount() string`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *CreateRefundRequest) GetPayableAmountOk() (*string, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *CreateRefundRequest) SetPayableAmount(v string)`

SetPayableAmount sets PayableAmount field to given value.


### GetToAddress

`func (o *CreateRefundRequest) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateRefundRequest) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateRefundRequest) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *CreateRefundRequest) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *CreateRefundRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateRefundRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateRefundRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetRefundType

`func (o *CreateRefundRequest) GetRefundType() RefundType`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *CreateRefundRequest) GetRefundTypeOk() (*RefundType, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *CreateRefundRequest) SetRefundType(v RefundType)`

SetRefundType sets RefundType field to given value.


### GetOrderId

`func (o *CreateRefundRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateRefundRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateRefundRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CreateRefundRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetChargeMerchantFee

`func (o *CreateRefundRequest) GetChargeMerchantFee() bool`

GetChargeMerchantFee returns the ChargeMerchantFee field if non-nil, zero value otherwise.

### GetChargeMerchantFeeOk

`func (o *CreateRefundRequest) GetChargeMerchantFeeOk() (*bool, bool)`

GetChargeMerchantFeeOk returns a tuple with the ChargeMerchantFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeMerchantFee

`func (o *CreateRefundRequest) SetChargeMerchantFee(v bool)`

SetChargeMerchantFee sets ChargeMerchantFee field to given value.

### HasChargeMerchantFee

`func (o *CreateRefundRequest) HasChargeMerchantFee() bool`

HasChargeMerchantFee returns a boolean if a field has been set.

### GetMerchantFeeAmount

`func (o *CreateRefundRequest) GetMerchantFeeAmount() string`

GetMerchantFeeAmount returns the MerchantFeeAmount field if non-nil, zero value otherwise.

### GetMerchantFeeAmountOk

`func (o *CreateRefundRequest) GetMerchantFeeAmountOk() (*string, bool)`

GetMerchantFeeAmountOk returns a tuple with the MerchantFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeAmount

`func (o *CreateRefundRequest) SetMerchantFeeAmount(v string)`

SetMerchantFeeAmount sets MerchantFeeAmount field to given value.

### HasMerchantFeeAmount

`func (o *CreateRefundRequest) HasMerchantFeeAmount() bool`

HasMerchantFeeAmount returns a boolean if a field has been set.

### GetMerchantFeeTokenId

`func (o *CreateRefundRequest) GetMerchantFeeTokenId() string`

GetMerchantFeeTokenId returns the MerchantFeeTokenId field if non-nil, zero value otherwise.

### GetMerchantFeeTokenIdOk

`func (o *CreateRefundRequest) GetMerchantFeeTokenIdOk() (*string, bool)`

GetMerchantFeeTokenIdOk returns a tuple with the MerchantFeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeTokenId

`func (o *CreateRefundRequest) SetMerchantFeeTokenId(v string)`

SetMerchantFeeTokenId sets MerchantFeeTokenId field to given value.

### HasMerchantFeeTokenId

`func (o *CreateRefundRequest) HasMerchantFeeTokenId() bool`

HasMerchantFeeTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



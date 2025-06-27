# Refund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID provided by you when creating the refund request. | [optional] 
**RefundId** | **string** | The refund order ID. | 
**OrderId** | Pointer to **string** | The order ID corresponding to this refund. | [optional] 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**TokenId** | **string** | The ID of the cryptocurrency used for refund. | 
**ChainId** | **string** | The ID of the blockchain network on which the refund transaction occurs. | 
**Amount** | **string** | The amount in cryptocurrency to be returned for this refund order. | 
**ToAddress** | **string** | The recipient&#39;s wallet address where the refund will be sent. | 
**Status** | [**RefundStatus**](RefundStatus.md) |  | 
**RefundType** | Pointer to [**RefundType**](RefundType.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The created time of the refund order, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the refund order, represented as a UNIX timestamp in seconds. | [optional] 
**Initiator** | Pointer to **string** | The initiator of this refund order, usually the user&#39;s API key. | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this refund order. Each transaction represents a separate blockchain operation related to the refund process. | [optional] 
**ChargeMerchantFee** | Pointer to **bool** | Indicates whether the merchant should bear the transaction fee for the refund.  If true, the fee will be deducted from merchant&#39;s account balance.  | [optional] 
**MerchantFeeAmount** | Pointer to **string** | The amount of the transaction fee that the merchant will bear for the refund.  This is only applicable if &#x60;charge_merchant_fee&#x60; is set to true.  | [optional] 
**MerchantFeeTokenId** | Pointer to **string** | The ID of the cryptocurrency used for the transaction fee.  This is only applicable if &#x60;charge_merchant_fee&#x60; is set to true.  | [optional] 

## Methods

### NewRefund

`func NewRefund(refundId string, tokenId string, chainId string, amount string, toAddress string, status RefundStatus, ) *Refund`

NewRefund instantiates a new Refund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundWithDefaults

`func NewRefundWithDefaults() *Refund`

NewRefundWithDefaults instantiates a new Refund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *Refund) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Refund) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Refund) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Refund) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRefundId

`func (o *Refund) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *Refund) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *Refund) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetOrderId

`func (o *Refund) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Refund) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Refund) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Refund) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetMerchantId

`func (o *Refund) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *Refund) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *Refund) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *Refund) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *Refund) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Refund) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Refund) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *Refund) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Refund) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Refund) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAmount

`func (o *Refund) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Refund) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Refund) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToAddress

`func (o *Refund) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *Refund) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *Refund) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetStatus

`func (o *Refund) GetStatus() RefundStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Refund) GetStatusOk() (*RefundStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Refund) SetStatus(v RefundStatus)`

SetStatus sets Status field to given value.


### GetRefundType

`func (o *Refund) GetRefundType() RefundType`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *Refund) GetRefundTypeOk() (*RefundType, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *Refund) SetRefundType(v RefundType)`

SetRefundType sets RefundType field to given value.

### HasRefundType

`func (o *Refund) HasRefundType() bool`

HasRefundType returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Refund) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Refund) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Refund) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Refund) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Refund) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Refund) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Refund) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *Refund) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetInitiator

`func (o *Refund) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Refund) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Refund) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *Refund) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetTransactions

`func (o *Refund) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Refund) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Refund) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Refund) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetChargeMerchantFee

`func (o *Refund) GetChargeMerchantFee() bool`

GetChargeMerchantFee returns the ChargeMerchantFee field if non-nil, zero value otherwise.

### GetChargeMerchantFeeOk

`func (o *Refund) GetChargeMerchantFeeOk() (*bool, bool)`

GetChargeMerchantFeeOk returns a tuple with the ChargeMerchantFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeMerchantFee

`func (o *Refund) SetChargeMerchantFee(v bool)`

SetChargeMerchantFee sets ChargeMerchantFee field to given value.

### HasChargeMerchantFee

`func (o *Refund) HasChargeMerchantFee() bool`

HasChargeMerchantFee returns a boolean if a field has been set.

### GetMerchantFeeAmount

`func (o *Refund) GetMerchantFeeAmount() string`

GetMerchantFeeAmount returns the MerchantFeeAmount field if non-nil, zero value otherwise.

### GetMerchantFeeAmountOk

`func (o *Refund) GetMerchantFeeAmountOk() (*string, bool)`

GetMerchantFeeAmountOk returns a tuple with the MerchantFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeAmount

`func (o *Refund) SetMerchantFeeAmount(v string)`

SetMerchantFeeAmount sets MerchantFeeAmount field to given value.

### HasMerchantFeeAmount

`func (o *Refund) HasMerchantFeeAmount() bool`

HasMerchantFeeAmount returns a boolean if a field has been set.

### GetMerchantFeeTokenId

`func (o *Refund) GetMerchantFeeTokenId() string`

GetMerchantFeeTokenId returns the MerchantFeeTokenId field if non-nil, zero value otherwise.

### GetMerchantFeeTokenIdOk

`func (o *Refund) GetMerchantFeeTokenIdOk() (*string, bool)`

GetMerchantFeeTokenIdOk returns a tuple with the MerchantFeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeTokenId

`func (o *Refund) SetMerchantFeeTokenId(v string)`

SetMerchantFeeTokenId sets MerchantFeeTokenId field to given value.

### HasMerchantFeeTokenId

`func (o *Refund) HasMerchantFeeTokenId() bool`

HasMerchantFeeTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



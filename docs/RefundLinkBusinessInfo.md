# RefundLinkBusinessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID of the original order payment or top-up.  On the refund page, the from address of this transaction will be pre-filled as the default refund address.  The refund will be processed in the same token and on the same blockchain as this transaction.  | 
**Amount** | **string** | The amount to refund, denominated in the cryptocurrency of the original payment transaction. The amount must be a positive number and can have up to two decimal places. | 
**RefundSource** | [**RefundType**](RefundType.md) |  | 
**MerchantId** | Pointer to **string** | The merchant ID, required if &#x60;refund_source&#x60; is &#x60;Merchant&#x60;. The fund will be deducted from the specified merchant&#39;s balance. | [optional] 
**FeeAmount** | Pointer to **string** | The developer fee amount to charge the merchant, denominated in the cryptocurrency of the original payment transaction. This field is only valid when &#x60;refund_source&#x60; is &#x60;Merchant&#x60;. For more information, please refer to [Funds allocation and balances](https://www.cobo.com/payments/en/guides/amounts-and-balances). Must be:   - A positive integer with up to two decimal places.   - Less than the refund amount  | [optional] 

## Methods

### NewRefundLinkBusinessInfo

`func NewRefundLinkBusinessInfo(transactionId string, amount string, refundSource RefundType, ) *RefundLinkBusinessInfo`

NewRefundLinkBusinessInfo instantiates a new RefundLinkBusinessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundLinkBusinessInfoWithDefaults

`func NewRefundLinkBusinessInfoWithDefaults() *RefundLinkBusinessInfo`

NewRefundLinkBusinessInfoWithDefaults instantiates a new RefundLinkBusinessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *RefundLinkBusinessInfo) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *RefundLinkBusinessInfo) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *RefundLinkBusinessInfo) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAmount

`func (o *RefundLinkBusinessInfo) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RefundLinkBusinessInfo) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RefundLinkBusinessInfo) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetRefundSource

`func (o *RefundLinkBusinessInfo) GetRefundSource() RefundType`

GetRefundSource returns the RefundSource field if non-nil, zero value otherwise.

### GetRefundSourceOk

`func (o *RefundLinkBusinessInfo) GetRefundSourceOk() (*RefundType, bool)`

GetRefundSourceOk returns a tuple with the RefundSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSource

`func (o *RefundLinkBusinessInfo) SetRefundSource(v RefundType)`

SetRefundSource sets RefundSource field to given value.


### GetMerchantId

`func (o *RefundLinkBusinessInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *RefundLinkBusinessInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *RefundLinkBusinessInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *RefundLinkBusinessInfo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetFeeAmount

`func (o *RefundLinkBusinessInfo) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *RefundLinkBusinessInfo) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *RefundLinkBusinessInfo) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *RefundLinkBusinessInfo) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Refund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID provided by you when creating the refund request. | [optional] 
**RefundId** | **string** | The refund order ID. | 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**TokenId** | **string** | The ID of the cryptocurrency used for refund. | 
**ChainId** | **string** | The ID of the blockchain network on which the refund transaction occurs. | 
**Amount** | **string** | The amount in cryptocurrency to be returned for this refund order. | 
**ToAddress** | **string** | The recipient&#39;s wallet address where the refund will be sent. | 
**Status** | [**RefundStatus**](RefundStatus.md) |  | 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this refund order. Each transaction represents a separate blockchain operation related to the refund process. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



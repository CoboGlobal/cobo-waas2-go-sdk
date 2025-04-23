# PaymentRefundEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. | 
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

### NewPaymentRefundEventData

`func NewPaymentRefundEventData(dataType string, refundId string, tokenId string, chainId string, amount string, toAddress string, status RefundStatus, ) *PaymentRefundEventData`

NewPaymentRefundEventData instantiates a new PaymentRefundEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRefundEventDataWithDefaults

`func NewPaymentRefundEventDataWithDefaults() *PaymentRefundEventData`

NewPaymentRefundEventDataWithDefaults instantiates a new PaymentRefundEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentRefundEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentRefundEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentRefundEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *PaymentRefundEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentRefundEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentRefundEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PaymentRefundEventData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRefundId

`func (o *PaymentRefundEventData) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *PaymentRefundEventData) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *PaymentRefundEventData) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetMerchantId

`func (o *PaymentRefundEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentRefundEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentRefundEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *PaymentRefundEventData) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *PaymentRefundEventData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentRefundEventData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentRefundEventData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *PaymentRefundEventData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PaymentRefundEventData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PaymentRefundEventData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAmount

`func (o *PaymentRefundEventData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRefundEventData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRefundEventData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToAddress

`func (o *PaymentRefundEventData) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *PaymentRefundEventData) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *PaymentRefundEventData) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetStatus

`func (o *PaymentRefundEventData) GetStatus() RefundStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRefundEventData) GetStatusOk() (*RefundStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRefundEventData) SetStatus(v RefundStatus)`

SetStatus sets Status field to given value.


### GetTransactions

`func (o *PaymentRefundEventData) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaymentRefundEventData) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaymentRefundEventData) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *PaymentRefundEventData) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



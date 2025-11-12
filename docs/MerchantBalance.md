# MerchantBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**TokenId** | **string** | The token ID, which is a unique identifier that specifies both the blockchain network and cryptocurrency token in the format &#x60;{CHAIN}_{TOKEN}&#x60;. | 
**AcquiringType** | [**AcquiringType**](AcquiringType.md) |  | 
**TotalReceivedAmount** | Pointer to **string** | The total amount of the token that has been received by the merchant. | [optional] 
**SettledAmount** | Pointer to **string** | The total amount of the token that has been paid out from the merchant&#39;s balance. | [optional] 
**RefundedAmount** | Pointer to **string** | The total amount of the token that has been refunded from the merchant&#39;s balance. | [optional] 
**TotalBalance** | Pointer to **string** |  The total balance of the token available for payout or refund for the merchant.  &#x60;total_balance&#x60; &#x3D; &#x60;total_received_amount&#x60; - &#x60;settled_amount&#x60; - &#x60;refunded_amount&#x60;  For more information, please refer to [Funds allocation and balances](https://www.cobo.com/developers/v2/payments/amounts-and-balances)  | [optional] 
**AvailableBalance** | Pointer to **string** | This field has been deprecated. | [optional] 

## Methods

### NewMerchantBalance

`func NewMerchantBalance(merchantId string, tokenId string, acquiringType AcquiringType, ) *MerchantBalance`

NewMerchantBalance instantiates a new MerchantBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantBalanceWithDefaults

`func NewMerchantBalanceWithDefaults() *MerchantBalance`

NewMerchantBalanceWithDefaults instantiates a new MerchantBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *MerchantBalance) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *MerchantBalance) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *MerchantBalance) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetTokenId

`func (o *MerchantBalance) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *MerchantBalance) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *MerchantBalance) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAcquiringType

`func (o *MerchantBalance) GetAcquiringType() AcquiringType`

GetAcquiringType returns the AcquiringType field if non-nil, zero value otherwise.

### GetAcquiringTypeOk

`func (o *MerchantBalance) GetAcquiringTypeOk() (*AcquiringType, bool)`

GetAcquiringTypeOk returns a tuple with the AcquiringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiringType

`func (o *MerchantBalance) SetAcquiringType(v AcquiringType)`

SetAcquiringType sets AcquiringType field to given value.


### GetTotalReceivedAmount

`func (o *MerchantBalance) GetTotalReceivedAmount() string`

GetTotalReceivedAmount returns the TotalReceivedAmount field if non-nil, zero value otherwise.

### GetTotalReceivedAmountOk

`func (o *MerchantBalance) GetTotalReceivedAmountOk() (*string, bool)`

GetTotalReceivedAmountOk returns a tuple with the TotalReceivedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReceivedAmount

`func (o *MerchantBalance) SetTotalReceivedAmount(v string)`

SetTotalReceivedAmount sets TotalReceivedAmount field to given value.

### HasTotalReceivedAmount

`func (o *MerchantBalance) HasTotalReceivedAmount() bool`

HasTotalReceivedAmount returns a boolean if a field has been set.

### GetSettledAmount

`func (o *MerchantBalance) GetSettledAmount() string`

GetSettledAmount returns the SettledAmount field if non-nil, zero value otherwise.

### GetSettledAmountOk

`func (o *MerchantBalance) GetSettledAmountOk() (*string, bool)`

GetSettledAmountOk returns a tuple with the SettledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAmount

`func (o *MerchantBalance) SetSettledAmount(v string)`

SetSettledAmount sets SettledAmount field to given value.

### HasSettledAmount

`func (o *MerchantBalance) HasSettledAmount() bool`

HasSettledAmount returns a boolean if a field has been set.

### GetRefundedAmount

`func (o *MerchantBalance) GetRefundedAmount() string`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *MerchantBalance) GetRefundedAmountOk() (*string, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *MerchantBalance) SetRefundedAmount(v string)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *MerchantBalance) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### GetTotalBalance

`func (o *MerchantBalance) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *MerchantBalance) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *MerchantBalance) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.

### HasTotalBalance

`func (o *MerchantBalance) HasTotalBalance() bool`

HasTotalBalance returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *MerchantBalance) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *MerchantBalance) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *MerchantBalance) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *MerchantBalance) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



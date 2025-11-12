# PspBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is a unique identifier that specifies both the blockchain network and cryptocurrency token in the format &#x60;{CHAIN}_{TOKEN}&#x60;. | 
**DeveloperFeeAmount** | Pointer to **string** | The total amount of the token that has been received as developer fee. | [optional] 
**SettledAmount** | Pointer to **string** | The total amount of the token that has been paid out from the developer&#39;s balance. | [optional] 
**RefundedAmount** | Pointer to **string** | The total amount of the token that has been refunded from the developer&#39;s balance. | [optional] 
**TotalBalance** | Pointer to **string** |  The total balance of the token available for payout or refund for the developer.  &#x60;total_balance&#x60; &#x3D; &#x60;developer_fee_amount&#x60; - &#x60;settled_amount&#x60; - &#x60;refunded_amount&#x60;  For more information, please refer to [Funds allocation and balances](https://www.cobo.com/developers/v2/payments/amounts-and-balances)  | [optional] 
**AvailableBalance** | Pointer to **string** | This field has been deprecated. | [optional] 

## Methods

### NewPspBalance

`func NewPspBalance(tokenId string, ) *PspBalance`

NewPspBalance instantiates a new PspBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPspBalanceWithDefaults

`func NewPspBalanceWithDefaults() *PspBalance`

NewPspBalanceWithDefaults instantiates a new PspBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *PspBalance) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PspBalance) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PspBalance) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDeveloperFeeAmount

`func (o *PspBalance) GetDeveloperFeeAmount() string`

GetDeveloperFeeAmount returns the DeveloperFeeAmount field if non-nil, zero value otherwise.

### GetDeveloperFeeAmountOk

`func (o *PspBalance) GetDeveloperFeeAmountOk() (*string, bool)`

GetDeveloperFeeAmountOk returns a tuple with the DeveloperFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperFeeAmount

`func (o *PspBalance) SetDeveloperFeeAmount(v string)`

SetDeveloperFeeAmount sets DeveloperFeeAmount field to given value.

### HasDeveloperFeeAmount

`func (o *PspBalance) HasDeveloperFeeAmount() bool`

HasDeveloperFeeAmount returns a boolean if a field has been set.

### GetSettledAmount

`func (o *PspBalance) GetSettledAmount() string`

GetSettledAmount returns the SettledAmount field if non-nil, zero value otherwise.

### GetSettledAmountOk

`func (o *PspBalance) GetSettledAmountOk() (*string, bool)`

GetSettledAmountOk returns a tuple with the SettledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAmount

`func (o *PspBalance) SetSettledAmount(v string)`

SetSettledAmount sets SettledAmount field to given value.

### HasSettledAmount

`func (o *PspBalance) HasSettledAmount() bool`

HasSettledAmount returns a boolean if a field has been set.

### GetRefundedAmount

`func (o *PspBalance) GetRefundedAmount() string`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *PspBalance) GetRefundedAmountOk() (*string, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *PspBalance) SetRefundedAmount(v string)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *PspBalance) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### GetTotalBalance

`func (o *PspBalance) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *PspBalance) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *PspBalance) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.

### HasTotalBalance

`func (o *PspBalance) HasTotalBalance() bool`

HasTotalBalance returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *PspBalance) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *PspBalance) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *PspBalance) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *PspBalance) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



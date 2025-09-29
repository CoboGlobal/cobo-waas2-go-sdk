# PaymentWalletBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The unique identifier of the wallet. | 
**TokenId** | **string** | The ID of the cryptocurrency. | 
**SweptBalance** | Pointer to **string** | The payment wallet swept balance. | [optional] 
**AvailableBalance** | Pointer to **string** | The payment wallet available balance. | [optional] 
**TotalBalance** | Pointer to **string** | The payment wallet total balance. | [optional] 
**AboveSweepThresholdBalance** | Pointer to **string** | The payment wallet above sweep threshold balance. | [optional] 

## Methods

### NewPaymentWalletBalance

`func NewPaymentWalletBalance(walletId string, tokenId string, ) *PaymentWalletBalance`

NewPaymentWalletBalance instantiates a new PaymentWalletBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWalletBalanceWithDefaults

`func NewPaymentWalletBalanceWithDefaults() *PaymentWalletBalance`

NewPaymentWalletBalanceWithDefaults instantiates a new PaymentWalletBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *PaymentWalletBalance) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PaymentWalletBalance) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PaymentWalletBalance) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTokenId

`func (o *PaymentWalletBalance) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentWalletBalance) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentWalletBalance) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSweptBalance

`func (o *PaymentWalletBalance) GetSweptBalance() string`

GetSweptBalance returns the SweptBalance field if non-nil, zero value otherwise.

### GetSweptBalanceOk

`func (o *PaymentWalletBalance) GetSweptBalanceOk() (*string, bool)`

GetSweptBalanceOk returns a tuple with the SweptBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweptBalance

`func (o *PaymentWalletBalance) SetSweptBalance(v string)`

SetSweptBalance sets SweptBalance field to given value.

### HasSweptBalance

`func (o *PaymentWalletBalance) HasSweptBalance() bool`

HasSweptBalance returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *PaymentWalletBalance) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *PaymentWalletBalance) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *PaymentWalletBalance) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *PaymentWalletBalance) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetTotalBalance

`func (o *PaymentWalletBalance) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *PaymentWalletBalance) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *PaymentWalletBalance) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.

### HasTotalBalance

`func (o *PaymentWalletBalance) HasTotalBalance() bool`

HasTotalBalance returns a boolean if a field has been set.

### GetAboveSweepThresholdBalance

`func (o *PaymentWalletBalance) GetAboveSweepThresholdBalance() string`

GetAboveSweepThresholdBalance returns the AboveSweepThresholdBalance field if non-nil, zero value otherwise.

### GetAboveSweepThresholdBalanceOk

`func (o *PaymentWalletBalance) GetAboveSweepThresholdBalanceOk() (*string, bool)`

GetAboveSweepThresholdBalanceOk returns a tuple with the AboveSweepThresholdBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboveSweepThresholdBalance

`func (o *PaymentWalletBalance) SetAboveSweepThresholdBalance(v string)`

SetAboveSweepThresholdBalance sets AboveSweepThresholdBalance field to given value.

### HasAboveSweepThresholdBalance

`func (o *PaymentWalletBalance) HasAboveSweepThresholdBalance() bool`

HasAboveSweepThresholdBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



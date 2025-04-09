# SettlementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | The merchant ID. For payment gateway balance, this field will be empty. | [optional] 
**TokenId** | Pointer to **string** | The ID of the cryptocurrency token. | [optional] 
**AvailableAmount** | **string** | The amount available for settlement in the specified cryptocurrency token. | 
**AvailableCurrencyBalance** | Pointer to **string** | The available currency balance. | [optional] 
**PendingAmount** | Pointer to **string** | The pending amount. | [optional] 
**PendingCurrencyBalance** | Pointer to **string** | The pending currency balance. | [optional] 

## Methods

### NewSettlementInfo

`func NewSettlementInfo(availableAmount string, ) *SettlementInfo`

NewSettlementInfo instantiates a new SettlementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementInfoWithDefaults

`func NewSettlementInfoWithDefaults() *SettlementInfo`

NewSettlementInfoWithDefaults instantiates a new SettlementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *SettlementInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *SettlementInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *SettlementInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *SettlementInfo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *SettlementInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SettlementInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SettlementInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *SettlementInfo) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAvailableAmount

`func (o *SettlementInfo) GetAvailableAmount() string`

GetAvailableAmount returns the AvailableAmount field if non-nil, zero value otherwise.

### GetAvailableAmountOk

`func (o *SettlementInfo) GetAvailableAmountOk() (*string, bool)`

GetAvailableAmountOk returns a tuple with the AvailableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAmount

`func (o *SettlementInfo) SetAvailableAmount(v string)`

SetAvailableAmount sets AvailableAmount field to given value.


### GetAvailableCurrencyBalance

`func (o *SettlementInfo) GetAvailableCurrencyBalance() string`

GetAvailableCurrencyBalance returns the AvailableCurrencyBalance field if non-nil, zero value otherwise.

### GetAvailableCurrencyBalanceOk

`func (o *SettlementInfo) GetAvailableCurrencyBalanceOk() (*string, bool)`

GetAvailableCurrencyBalanceOk returns a tuple with the AvailableCurrencyBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCurrencyBalance

`func (o *SettlementInfo) SetAvailableCurrencyBalance(v string)`

SetAvailableCurrencyBalance sets AvailableCurrencyBalance field to given value.

### HasAvailableCurrencyBalance

`func (o *SettlementInfo) HasAvailableCurrencyBalance() bool`

HasAvailableCurrencyBalance returns a boolean if a field has been set.

### GetPendingAmount

`func (o *SettlementInfo) GetPendingAmount() string`

GetPendingAmount returns the PendingAmount field if non-nil, zero value otherwise.

### GetPendingAmountOk

`func (o *SettlementInfo) GetPendingAmountOk() (*string, bool)`

GetPendingAmountOk returns a tuple with the PendingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmount

`func (o *SettlementInfo) SetPendingAmount(v string)`

SetPendingAmount sets PendingAmount field to given value.

### HasPendingAmount

`func (o *SettlementInfo) HasPendingAmount() bool`

HasPendingAmount returns a boolean if a field has been set.

### GetPendingCurrencyBalance

`func (o *SettlementInfo) GetPendingCurrencyBalance() string`

GetPendingCurrencyBalance returns the PendingCurrencyBalance field if non-nil, zero value otherwise.

### GetPendingCurrencyBalanceOk

`func (o *SettlementInfo) GetPendingCurrencyBalanceOk() (*string, bool)`

GetPendingCurrencyBalanceOk returns a tuple with the PendingCurrencyBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCurrencyBalance

`func (o *SettlementInfo) SetPendingCurrencyBalance(v string)`

SetPendingCurrencyBalance sets PendingCurrencyBalance field to given value.

### HasPendingCurrencyBalance

`func (o *SettlementInfo) HasPendingCurrencyBalance() bool`

HasPendingCurrencyBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



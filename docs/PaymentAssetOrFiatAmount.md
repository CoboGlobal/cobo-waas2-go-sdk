# PaymentAssetOrFiatAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountType** | **string** | The amount representation type. | 
**Amount** | **string** | The amount, represented as a decimal string. | 
**TokenId** | Pointer to **string** | The token ID of the asset. This field is present when &#x60;amount_type&#x60; is &#x60;Crypto&#x60;. | [optional] 
**ChainId** | Pointer to **string** | The chain ID of the asset. This field is present when &#x60;amount_type&#x60; is &#x60;Crypto&#x60;. | [optional] 
**Currency** | Pointer to **string** | The fiat currency code. This field is present when &#x60;amount_type&#x60; is &#x60;Fiat&#x60;. | [optional] 

## Methods

### NewPaymentAssetOrFiatAmount

`func NewPaymentAssetOrFiatAmount(amountType string, amount string, ) *PaymentAssetOrFiatAmount`

NewPaymentAssetOrFiatAmount instantiates a new PaymentAssetOrFiatAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAssetOrFiatAmountWithDefaults

`func NewPaymentAssetOrFiatAmountWithDefaults() *PaymentAssetOrFiatAmount`

NewPaymentAssetOrFiatAmountWithDefaults instantiates a new PaymentAssetOrFiatAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountType

`func (o *PaymentAssetOrFiatAmount) GetAmountType() string`

GetAmountType returns the AmountType field if non-nil, zero value otherwise.

### GetAmountTypeOk

`func (o *PaymentAssetOrFiatAmount) GetAmountTypeOk() (*string, bool)`

GetAmountTypeOk returns a tuple with the AmountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountType

`func (o *PaymentAssetOrFiatAmount) SetAmountType(v string)`

SetAmountType sets AmountType field to given value.


### GetAmount

`func (o *PaymentAssetOrFiatAmount) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAssetOrFiatAmount) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAssetOrFiatAmount) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokenId

`func (o *PaymentAssetOrFiatAmount) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentAssetOrFiatAmount) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentAssetOrFiatAmount) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *PaymentAssetOrFiatAmount) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetChainId

`func (o *PaymentAssetOrFiatAmount) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PaymentAssetOrFiatAmount) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PaymentAssetOrFiatAmount) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *PaymentAssetOrFiatAmount) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentAssetOrFiatAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentAssetOrFiatAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentAssetOrFiatAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentAssetOrFiatAmount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



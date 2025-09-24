# PaymentCreateSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PaymentSubscriptionActionType**](PaymentSubscriptionActionType.md) |  | 
**UserAddress** | **string** | The subscription user address. | 
**Amount** | **string** | The subscription crypto amount.  | 
**TokenId** | **string** | The ID of the cryptocurrency you want to subscription. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**DiscountRate** | Pointer to **int32** | the discount rate, discount_rate/10000 | [optional] 

## Methods

### NewPaymentCreateSubscriptionActionData

`func NewPaymentCreateSubscriptionActionData(actionType PaymentSubscriptionActionType, userAddress string, amount string, tokenId string, ) *PaymentCreateSubscriptionActionData`

NewPaymentCreateSubscriptionActionData instantiates a new PaymentCreateSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateSubscriptionActionDataWithDefaults

`func NewPaymentCreateSubscriptionActionDataWithDefaults() *PaymentCreateSubscriptionActionData`

NewPaymentCreateSubscriptionActionDataWithDefaults instantiates a new PaymentCreateSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PaymentCreateSubscriptionActionData) GetActionType() PaymentSubscriptionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentCreateSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentCreateSubscriptionActionData) SetActionType(v PaymentSubscriptionActionType)`

SetActionType sets ActionType field to given value.


### GetUserAddress

`func (o *PaymentCreateSubscriptionActionData) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *PaymentCreateSubscriptionActionData) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *PaymentCreateSubscriptionActionData) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetAmount

`func (o *PaymentCreateSubscriptionActionData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentCreateSubscriptionActionData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentCreateSubscriptionActionData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokenId

`func (o *PaymentCreateSubscriptionActionData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentCreateSubscriptionActionData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentCreateSubscriptionActionData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDiscountRate

`func (o *PaymentCreateSubscriptionActionData) GetDiscountRate() int32`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *PaymentCreateSubscriptionActionData) GetDiscountRateOk() (*int32, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *PaymentCreateSubscriptionActionData) SetDiscountRate(v int32)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *PaymentCreateSubscriptionActionData) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



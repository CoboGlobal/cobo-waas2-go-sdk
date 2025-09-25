# PaymentSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PaymentSubscriptionActionType**](PaymentSubscriptionActionType.md) |  | 
**UserAddress** | **string** | The subscription user address. | 
**Amount** | **string** | The subscription crypto amount.  | 
**TokenId** | **string** | The ID of the cryptocurrency you want to subscription. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**DiscountRate** | Pointer to **int32** | the discount rate, discount_rate/10000 | [optional] 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**PermitData** | Pointer to **string** | The signature of permit. | [optional] 
**Signature** | **string** | The signature for transaction. | 
**Periods** | Pointer to **int32** | The periods needed updated. | [optional] 
**NewPlanId** | **string** | The new plan id in cobo. | 

## Methods

### NewPaymentSubscriptionActionData

`func NewPaymentSubscriptionActionData(actionType PaymentSubscriptionActionType, userAddress string, amount string, tokenId string, subscriptionId string, signature string, newPlanId string, ) *PaymentSubscriptionActionData`

NewPaymentSubscriptionActionData instantiates a new PaymentSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubscriptionActionDataWithDefaults

`func NewPaymentSubscriptionActionDataWithDefaults() *PaymentSubscriptionActionData`

NewPaymentSubscriptionActionDataWithDefaults instantiates a new PaymentSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PaymentSubscriptionActionData) GetActionType() PaymentSubscriptionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentSubscriptionActionData) SetActionType(v PaymentSubscriptionActionType)`

SetActionType sets ActionType field to given value.


### GetUserAddress

`func (o *PaymentSubscriptionActionData) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *PaymentSubscriptionActionData) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *PaymentSubscriptionActionData) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetAmount

`func (o *PaymentSubscriptionActionData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentSubscriptionActionData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentSubscriptionActionData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokenId

`func (o *PaymentSubscriptionActionData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentSubscriptionActionData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentSubscriptionActionData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDiscountRate

`func (o *PaymentSubscriptionActionData) GetDiscountRate() int32`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *PaymentSubscriptionActionData) GetDiscountRateOk() (*int32, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *PaymentSubscriptionActionData) SetDiscountRate(v int32)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *PaymentSubscriptionActionData) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PaymentSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPermitData

`func (o *PaymentSubscriptionActionData) GetPermitData() string`

GetPermitData returns the PermitData field if non-nil, zero value otherwise.

### GetPermitDataOk

`func (o *PaymentSubscriptionActionData) GetPermitDataOk() (*string, bool)`

GetPermitDataOk returns a tuple with the PermitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitData

`func (o *PaymentSubscriptionActionData) SetPermitData(v string)`

SetPermitData sets PermitData field to given value.

### HasPermitData

`func (o *PaymentSubscriptionActionData) HasPermitData() bool`

HasPermitData returns a boolean if a field has been set.

### GetSignature

`func (o *PaymentSubscriptionActionData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentSubscriptionActionData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentSubscriptionActionData) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetPeriods

`func (o *PaymentSubscriptionActionData) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *PaymentSubscriptionActionData) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *PaymentSubscriptionActionData) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *PaymentSubscriptionActionData) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### GetNewPlanId

`func (o *PaymentSubscriptionActionData) GetNewPlanId() string`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *PaymentSubscriptionActionData) GetNewPlanIdOk() (*string, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *PaymentSubscriptionActionData) SetNewPlanId(v string)`

SetNewPlanId sets NewPlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



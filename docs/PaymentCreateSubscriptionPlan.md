# PaymentCreateSubscriptionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeveloperPlanId** | **string** | The developer plan id. | 
**PeriodType** | [**PaymentSubscriptionPeriodType**](PaymentSubscriptionPeriodType.md) |  | 
**Periods** | **int32** |  | 
**Amount** | **string** | The subscription plan amount.  - If &#x60;currency&#x60; is set, this represents the subscription amount in the specified fiat currency. - If &#x60;currency&#x60; isn&#39;t set, this represents the settlement amount in the specified cryptocurrency.  | 
**TokenId** | Pointer to **string** | The ID of the cryptocurrency you want to subscription. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | [optional] 
**Currency** | Pointer to **string** | The fiat currency for settling the cryptocurrency. Currently, only &#x60;USD&#x60; is supported. Specify this field when &#x60;payout_channel&#x60; is set to &#x60;OffRamp&#x60;. | [optional] 

## Methods

### NewPaymentCreateSubscriptionPlan

`func NewPaymentCreateSubscriptionPlan(developerPlanId string, periodType PaymentSubscriptionPeriodType, periods int32, amount string, ) *PaymentCreateSubscriptionPlan`

NewPaymentCreateSubscriptionPlan instantiates a new PaymentCreateSubscriptionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateSubscriptionPlanWithDefaults

`func NewPaymentCreateSubscriptionPlanWithDefaults() *PaymentCreateSubscriptionPlan`

NewPaymentCreateSubscriptionPlanWithDefaults instantiates a new PaymentCreateSubscriptionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeveloperPlanId

`func (o *PaymentCreateSubscriptionPlan) GetDeveloperPlanId() string`

GetDeveloperPlanId returns the DeveloperPlanId field if non-nil, zero value otherwise.

### GetDeveloperPlanIdOk

`func (o *PaymentCreateSubscriptionPlan) GetDeveloperPlanIdOk() (*string, bool)`

GetDeveloperPlanIdOk returns a tuple with the DeveloperPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperPlanId

`func (o *PaymentCreateSubscriptionPlan) SetDeveloperPlanId(v string)`

SetDeveloperPlanId sets DeveloperPlanId field to given value.


### GetPeriodType

`func (o *PaymentCreateSubscriptionPlan) GetPeriodType() PaymentSubscriptionPeriodType`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *PaymentCreateSubscriptionPlan) GetPeriodTypeOk() (*PaymentSubscriptionPeriodType, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *PaymentCreateSubscriptionPlan) SetPeriodType(v PaymentSubscriptionPeriodType)`

SetPeriodType sets PeriodType field to given value.


### GetPeriods

`func (o *PaymentCreateSubscriptionPlan) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *PaymentCreateSubscriptionPlan) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *PaymentCreateSubscriptionPlan) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.


### GetAmount

`func (o *PaymentCreateSubscriptionPlan) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentCreateSubscriptionPlan) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentCreateSubscriptionPlan) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokenId

`func (o *PaymentCreateSubscriptionPlan) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentCreateSubscriptionPlan) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentCreateSubscriptionPlan) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *PaymentCreateSubscriptionPlan) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentCreateSubscriptionPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentCreateSubscriptionPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentCreateSubscriptionPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentCreateSubscriptionPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



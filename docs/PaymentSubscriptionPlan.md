# PaymentSubscriptionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | The plan id in cobo. | 
**DeveloperPlanId** | **string** | The developer plan id. | 
**PeriodType** | [**PaymentSubscriptionPeriodType**](PaymentSubscriptionPeriodType.md) |  | 
**Periods** | **int32** |  | 
**Interval** | **int32** |  | 
**TrialPeriod** | Pointer to **int32** | probation period | [optional] 
**Amount** | **string** | The subscription plan amount.  - If &#x60;currency&#x60; is set, this represents the subscription amount in the specified fiat currency. - If &#x60;currency&#x60; isn&#39;t set, this represents the settlement amount in the specified cryptocurrency.  | 
**TokenId** | Pointer to **string** | The ID of the cryptocurrency you want to subscription. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | [optional] 
**Currency** | Pointer to **string** | The fiat currency for settling the cryptocurrency. Currently, only &#x60;USD&#x60; is supported. Specify this field when &#x60;payout_channel&#x60; is set to &#x60;OffRamp&#x60;. | [optional] 

## Methods

### NewPaymentSubscriptionPlan

`func NewPaymentSubscriptionPlan(planId string, developerPlanId string, periodType PaymentSubscriptionPeriodType, periods int32, interval int32, amount string, ) *PaymentSubscriptionPlan`

NewPaymentSubscriptionPlan instantiates a new PaymentSubscriptionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubscriptionPlanWithDefaults

`func NewPaymentSubscriptionPlanWithDefaults() *PaymentSubscriptionPlan`

NewPaymentSubscriptionPlanWithDefaults instantiates a new PaymentSubscriptionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *PaymentSubscriptionPlan) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PaymentSubscriptionPlan) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PaymentSubscriptionPlan) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetDeveloperPlanId

`func (o *PaymentSubscriptionPlan) GetDeveloperPlanId() string`

GetDeveloperPlanId returns the DeveloperPlanId field if non-nil, zero value otherwise.

### GetDeveloperPlanIdOk

`func (o *PaymentSubscriptionPlan) GetDeveloperPlanIdOk() (*string, bool)`

GetDeveloperPlanIdOk returns a tuple with the DeveloperPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperPlanId

`func (o *PaymentSubscriptionPlan) SetDeveloperPlanId(v string)`

SetDeveloperPlanId sets DeveloperPlanId field to given value.


### GetPeriodType

`func (o *PaymentSubscriptionPlan) GetPeriodType() PaymentSubscriptionPeriodType`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *PaymentSubscriptionPlan) GetPeriodTypeOk() (*PaymentSubscriptionPeriodType, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *PaymentSubscriptionPlan) SetPeriodType(v PaymentSubscriptionPeriodType)`

SetPeriodType sets PeriodType field to given value.


### GetPeriods

`func (o *PaymentSubscriptionPlan) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *PaymentSubscriptionPlan) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *PaymentSubscriptionPlan) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.


### GetInterval

`func (o *PaymentSubscriptionPlan) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PaymentSubscriptionPlan) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PaymentSubscriptionPlan) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetTrialPeriod

`func (o *PaymentSubscriptionPlan) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *PaymentSubscriptionPlan) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *PaymentSubscriptionPlan) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *PaymentSubscriptionPlan) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentSubscriptionPlan) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentSubscriptionPlan) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentSubscriptionPlan) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokenId

`func (o *PaymentSubscriptionPlan) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentSubscriptionPlan) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentSubscriptionPlan) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *PaymentSubscriptionPlan) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentSubscriptionPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentSubscriptionPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentSubscriptionPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentSubscriptionPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



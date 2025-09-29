# PaymentSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | The plan id in cobo. | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**MerchantId** | **string** | The merchant id in cobo. | 
**MerchantAddress** | **string** | The merchant address in cobo. | 
**UserAddress** | **string** | The user address in subscription. | 
**TokenId** | **string** | The token_id in subscription. | 
**Amount** | **string** | The amount in subscription. | 
**StartTime** | **int32** | The subscription start timestamp. | 
**ExpirationTime** | **int32** | The subscription expired timestamp. | 
**ChargesMade** | **int32** | The subscription charge times. | 
**PeriodType** | [**PaymentSubscriptionPeriodType**](PaymentSubscriptionPeriodType.md) |  | 
**Periods** | **int32** |  | 
**Interval** | **int32** | The subscription charge interval. | 
**Status** | [**PaymentSubscriptionStatus**](PaymentSubscriptionStatus.md) |  | 
**CreatedTimestamp** | **int32** | The created time of the subscription, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the subscription, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentSubscription

`func NewPaymentSubscription(planId string, subscriptionId string, merchantId string, merchantAddress string, userAddress string, tokenId string, amount string, startTime int32, expirationTime int32, chargesMade int32, periodType PaymentSubscriptionPeriodType, periods int32, interval int32, status PaymentSubscriptionStatus, createdTimestamp int32, updatedTimestamp int32, ) *PaymentSubscription`

NewPaymentSubscription instantiates a new PaymentSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubscriptionWithDefaults

`func NewPaymentSubscriptionWithDefaults() *PaymentSubscription`

NewPaymentSubscriptionWithDefaults instantiates a new PaymentSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *PaymentSubscription) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PaymentSubscription) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PaymentSubscription) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetSubscriptionId

`func (o *PaymentSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetMerchantId

`func (o *PaymentSubscription) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentSubscription) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentSubscription) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantAddress

`func (o *PaymentSubscription) GetMerchantAddress() string`

GetMerchantAddress returns the MerchantAddress field if non-nil, zero value otherwise.

### GetMerchantAddressOk

`func (o *PaymentSubscription) GetMerchantAddressOk() (*string, bool)`

GetMerchantAddressOk returns a tuple with the MerchantAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAddress

`func (o *PaymentSubscription) SetMerchantAddress(v string)`

SetMerchantAddress sets MerchantAddress field to given value.


### GetUserAddress

`func (o *PaymentSubscription) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *PaymentSubscription) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *PaymentSubscription) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetTokenId

`func (o *PaymentSubscription) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentSubscription) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentSubscription) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *PaymentSubscription) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentSubscription) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentSubscription) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetStartTime

`func (o *PaymentSubscription) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PaymentSubscription) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PaymentSubscription) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetExpirationTime

`func (o *PaymentSubscription) GetExpirationTime() int32`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *PaymentSubscription) GetExpirationTimeOk() (*int32, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *PaymentSubscription) SetExpirationTime(v int32)`

SetExpirationTime sets ExpirationTime field to given value.


### GetChargesMade

`func (o *PaymentSubscription) GetChargesMade() int32`

GetChargesMade returns the ChargesMade field if non-nil, zero value otherwise.

### GetChargesMadeOk

`func (o *PaymentSubscription) GetChargesMadeOk() (*int32, bool)`

GetChargesMadeOk returns a tuple with the ChargesMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesMade

`func (o *PaymentSubscription) SetChargesMade(v int32)`

SetChargesMade sets ChargesMade field to given value.


### GetPeriodType

`func (o *PaymentSubscription) GetPeriodType() PaymentSubscriptionPeriodType`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *PaymentSubscription) GetPeriodTypeOk() (*PaymentSubscriptionPeriodType, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *PaymentSubscription) SetPeriodType(v PaymentSubscriptionPeriodType)`

SetPeriodType sets PeriodType field to given value.


### GetPeriods

`func (o *PaymentSubscription) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *PaymentSubscription) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *PaymentSubscription) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.


### GetInterval

`func (o *PaymentSubscription) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PaymentSubscription) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PaymentSubscription) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetStatus

`func (o *PaymentSubscription) GetStatus() PaymentSubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentSubscription) GetStatusOk() (*PaymentSubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentSubscription) SetStatus(v PaymentSubscriptionStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentSubscription) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentSubscription) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentSubscription) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentSubscription) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentSubscription) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentSubscription) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



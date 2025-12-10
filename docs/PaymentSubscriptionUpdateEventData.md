# PaymentSubscriptionUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The payment address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The suspended token event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**PlanId** | **string** | The plan id in cobo. | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**MerchantId** | **string** | The merchant id in cobo. | 
**MerchantAddress** | **string** | The merchant address in cobo. | 
**UserAddress** | **string** | The user address in subscription. | 
**TokenId** | **string** | The token_id in subscription. | 
**ChargeAmount** | Pointer to **string** | The charge amount in subscription. | [optional] 
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

### NewPaymentSubscriptionUpdateEventData

`func NewPaymentSubscriptionUpdateEventData(dataType string, planId string, subscriptionId string, merchantId string, merchantAddress string, userAddress string, tokenId string, startTime int32, expirationTime int32, chargesMade int32, periodType PaymentSubscriptionPeriodType, periods int32, interval int32, status PaymentSubscriptionStatus, createdTimestamp int32, updatedTimestamp int32, ) *PaymentSubscriptionUpdateEventData`

NewPaymentSubscriptionUpdateEventData instantiates a new PaymentSubscriptionUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubscriptionUpdateEventDataWithDefaults

`func NewPaymentSubscriptionUpdateEventDataWithDefaults() *PaymentSubscriptionUpdateEventData`

NewPaymentSubscriptionUpdateEventDataWithDefaults instantiates a new PaymentSubscriptionUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentSubscriptionUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentSubscriptionUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentSubscriptionUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetPlanId

`func (o *PaymentSubscriptionUpdateEventData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PaymentSubscriptionUpdateEventData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PaymentSubscriptionUpdateEventData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetSubscriptionId

`func (o *PaymentSubscriptionUpdateEventData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentSubscriptionUpdateEventData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentSubscriptionUpdateEventData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetMerchantId

`func (o *PaymentSubscriptionUpdateEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentSubscriptionUpdateEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentSubscriptionUpdateEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantAddress

`func (o *PaymentSubscriptionUpdateEventData) GetMerchantAddress() string`

GetMerchantAddress returns the MerchantAddress field if non-nil, zero value otherwise.

### GetMerchantAddressOk

`func (o *PaymentSubscriptionUpdateEventData) GetMerchantAddressOk() (*string, bool)`

GetMerchantAddressOk returns a tuple with the MerchantAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAddress

`func (o *PaymentSubscriptionUpdateEventData) SetMerchantAddress(v string)`

SetMerchantAddress sets MerchantAddress field to given value.


### GetUserAddress

`func (o *PaymentSubscriptionUpdateEventData) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *PaymentSubscriptionUpdateEventData) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *PaymentSubscriptionUpdateEventData) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetTokenId

`func (o *PaymentSubscriptionUpdateEventData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentSubscriptionUpdateEventData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentSubscriptionUpdateEventData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChargeAmount

`func (o *PaymentSubscriptionUpdateEventData) GetChargeAmount() string`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *PaymentSubscriptionUpdateEventData) GetChargeAmountOk() (*string, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *PaymentSubscriptionUpdateEventData) SetChargeAmount(v string)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *PaymentSubscriptionUpdateEventData) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### GetStartTime

`func (o *PaymentSubscriptionUpdateEventData) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PaymentSubscriptionUpdateEventData) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PaymentSubscriptionUpdateEventData) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetExpirationTime

`func (o *PaymentSubscriptionUpdateEventData) GetExpirationTime() int32`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *PaymentSubscriptionUpdateEventData) GetExpirationTimeOk() (*int32, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *PaymentSubscriptionUpdateEventData) SetExpirationTime(v int32)`

SetExpirationTime sets ExpirationTime field to given value.


### GetChargesMade

`func (o *PaymentSubscriptionUpdateEventData) GetChargesMade() int32`

GetChargesMade returns the ChargesMade field if non-nil, zero value otherwise.

### GetChargesMadeOk

`func (o *PaymentSubscriptionUpdateEventData) GetChargesMadeOk() (*int32, bool)`

GetChargesMadeOk returns a tuple with the ChargesMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesMade

`func (o *PaymentSubscriptionUpdateEventData) SetChargesMade(v int32)`

SetChargesMade sets ChargesMade field to given value.


### GetPeriodType

`func (o *PaymentSubscriptionUpdateEventData) GetPeriodType() PaymentSubscriptionPeriodType`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *PaymentSubscriptionUpdateEventData) GetPeriodTypeOk() (*PaymentSubscriptionPeriodType, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *PaymentSubscriptionUpdateEventData) SetPeriodType(v PaymentSubscriptionPeriodType)`

SetPeriodType sets PeriodType field to given value.


### GetPeriods

`func (o *PaymentSubscriptionUpdateEventData) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *PaymentSubscriptionUpdateEventData) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *PaymentSubscriptionUpdateEventData) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.


### GetInterval

`func (o *PaymentSubscriptionUpdateEventData) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PaymentSubscriptionUpdateEventData) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PaymentSubscriptionUpdateEventData) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetStatus

`func (o *PaymentSubscriptionUpdateEventData) GetStatus() PaymentSubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentSubscriptionUpdateEventData) GetStatusOk() (*PaymentSubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentSubscriptionUpdateEventData) SetStatus(v PaymentSubscriptionStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentSubscriptionUpdateEventData) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentSubscriptionUpdateEventData) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentSubscriptionUpdateEventData) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentSubscriptionUpdateEventData) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentSubscriptionUpdateEventData) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentSubscriptionUpdateEventData) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



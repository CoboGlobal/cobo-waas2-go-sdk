# PaymentSubscriptionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The action request id. | 
**ActionId** | **string** | The action id. | 
**PlanId** | **string** | The plan id in cobo. | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**MerchantId** | **string** | The merchant id in cobo. | 
**MerchantAddress** | **string** | The merchant address in cobo. | 
**Data** | [**PaymentSubscriptionActionData**](PaymentSubscriptionActionData.md) |  | 
**Status** | [**PaymentSubscriptionActionStatus**](PaymentSubscriptionActionStatus.md) |  | 
**CreatedTimestamp** | Pointer to **int32** | The created time of the subscription action, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the subscription action, represented as a UNIX timestamp in seconds. | [optional] 

## Methods

### NewPaymentSubscriptionAction

`func NewPaymentSubscriptionAction(requestId string, actionId string, planId string, subscriptionId string, merchantId string, merchantAddress string, data PaymentSubscriptionActionData, status PaymentSubscriptionActionStatus, ) *PaymentSubscriptionAction`

NewPaymentSubscriptionAction instantiates a new PaymentSubscriptionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubscriptionActionWithDefaults

`func NewPaymentSubscriptionActionWithDefaults() *PaymentSubscriptionAction`

NewPaymentSubscriptionActionWithDefaults instantiates a new PaymentSubscriptionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *PaymentSubscriptionAction) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentSubscriptionAction) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentSubscriptionAction) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetActionId

`func (o *PaymentSubscriptionAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *PaymentSubscriptionAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *PaymentSubscriptionAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetPlanId

`func (o *PaymentSubscriptionAction) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PaymentSubscriptionAction) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PaymentSubscriptionAction) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetSubscriptionId

`func (o *PaymentSubscriptionAction) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentSubscriptionAction) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentSubscriptionAction) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetMerchantId

`func (o *PaymentSubscriptionAction) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentSubscriptionAction) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentSubscriptionAction) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantAddress

`func (o *PaymentSubscriptionAction) GetMerchantAddress() string`

GetMerchantAddress returns the MerchantAddress field if non-nil, zero value otherwise.

### GetMerchantAddressOk

`func (o *PaymentSubscriptionAction) GetMerchantAddressOk() (*string, bool)`

GetMerchantAddressOk returns a tuple with the MerchantAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAddress

`func (o *PaymentSubscriptionAction) SetMerchantAddress(v string)`

SetMerchantAddress sets MerchantAddress field to given value.


### GetData

`func (o *PaymentSubscriptionAction) GetData() PaymentSubscriptionActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentSubscriptionAction) GetDataOk() (*PaymentSubscriptionActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentSubscriptionAction) SetData(v PaymentSubscriptionActionData)`

SetData sets Data field to given value.


### GetStatus

`func (o *PaymentSubscriptionAction) GetStatus() PaymentSubscriptionActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentSubscriptionAction) GetStatusOk() (*PaymentSubscriptionActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentSubscriptionAction) SetStatus(v PaymentSubscriptionActionStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentSubscriptionAction) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentSubscriptionAction) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentSubscriptionAction) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentSubscriptionAction) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PaymentSubscriptionAction) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentSubscriptionAction) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentSubscriptionAction) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PaymentSubscriptionAction) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



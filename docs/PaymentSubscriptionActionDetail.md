# PaymentSubscriptionActionDetail

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
**TransactionIds** | Pointer to **[]string** |  | [optional] 
**Status** | [**PaymentSubscriptionActionStatus**](PaymentSubscriptionActionStatus.md) |  | 
**CreatedTimestamp** | Pointer to **int32** | The created time of the subscription action, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the subscription action, represented as a UNIX timestamp in seconds. | [optional] 
**Subscription** | Pointer to [**PaymentSubscription**](PaymentSubscription.md) |  | [optional] 
**Transactions** | Pointer to [**[]Transaction**](Transaction.md) | An array of subscription transactions. | [optional] 

## Methods

### NewPaymentSubscriptionActionDetail

`func NewPaymentSubscriptionActionDetail(requestId string, actionId string, planId string, subscriptionId string, merchantId string, merchantAddress string, data PaymentSubscriptionActionData, status PaymentSubscriptionActionStatus, ) *PaymentSubscriptionActionDetail`

NewPaymentSubscriptionActionDetail instantiates a new PaymentSubscriptionActionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubscriptionActionDetailWithDefaults

`func NewPaymentSubscriptionActionDetailWithDefaults() *PaymentSubscriptionActionDetail`

NewPaymentSubscriptionActionDetailWithDefaults instantiates a new PaymentSubscriptionActionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *PaymentSubscriptionActionDetail) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentSubscriptionActionDetail) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentSubscriptionActionDetail) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetActionId

`func (o *PaymentSubscriptionActionDetail) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *PaymentSubscriptionActionDetail) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *PaymentSubscriptionActionDetail) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetPlanId

`func (o *PaymentSubscriptionActionDetail) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PaymentSubscriptionActionDetail) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PaymentSubscriptionActionDetail) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetSubscriptionId

`func (o *PaymentSubscriptionActionDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentSubscriptionActionDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentSubscriptionActionDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetMerchantId

`func (o *PaymentSubscriptionActionDetail) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentSubscriptionActionDetail) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentSubscriptionActionDetail) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantAddress

`func (o *PaymentSubscriptionActionDetail) GetMerchantAddress() string`

GetMerchantAddress returns the MerchantAddress field if non-nil, zero value otherwise.

### GetMerchantAddressOk

`func (o *PaymentSubscriptionActionDetail) GetMerchantAddressOk() (*string, bool)`

GetMerchantAddressOk returns a tuple with the MerchantAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAddress

`func (o *PaymentSubscriptionActionDetail) SetMerchantAddress(v string)`

SetMerchantAddress sets MerchantAddress field to given value.


### GetData

`func (o *PaymentSubscriptionActionDetail) GetData() PaymentSubscriptionActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentSubscriptionActionDetail) GetDataOk() (*PaymentSubscriptionActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentSubscriptionActionDetail) SetData(v PaymentSubscriptionActionData)`

SetData sets Data field to given value.


### GetTransactionIds

`func (o *PaymentSubscriptionActionDetail) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *PaymentSubscriptionActionDetail) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *PaymentSubscriptionActionDetail) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *PaymentSubscriptionActionDetail) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentSubscriptionActionDetail) GetStatus() PaymentSubscriptionActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentSubscriptionActionDetail) GetStatusOk() (*PaymentSubscriptionActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentSubscriptionActionDetail) SetStatus(v PaymentSubscriptionActionStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentSubscriptionActionDetail) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentSubscriptionActionDetail) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentSubscriptionActionDetail) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentSubscriptionActionDetail) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PaymentSubscriptionActionDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentSubscriptionActionDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentSubscriptionActionDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PaymentSubscriptionActionDetail) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetSubscription

`func (o *PaymentSubscriptionActionDetail) GetSubscription() PaymentSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *PaymentSubscriptionActionDetail) GetSubscriptionOk() (*PaymentSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *PaymentSubscriptionActionDetail) SetSubscription(v PaymentSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *PaymentSubscriptionActionDetail) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTransactions

`func (o *PaymentSubscriptionActionDetail) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaymentSubscriptionActionDetail) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaymentSubscriptionActionDetail) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *PaymentSubscriptionActionDetail) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



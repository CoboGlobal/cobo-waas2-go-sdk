# PaymentChargeUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The payment address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The suspended token event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
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

## Methods

### NewPaymentChargeUpdateEventData

`func NewPaymentChargeUpdateEventData(dataType string, requestId string, actionId string, planId string, subscriptionId string, merchantId string, merchantAddress string, data PaymentSubscriptionActionData, status PaymentSubscriptionActionStatus, ) *PaymentChargeUpdateEventData`

NewPaymentChargeUpdateEventData instantiates a new PaymentChargeUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChargeUpdateEventDataWithDefaults

`func NewPaymentChargeUpdateEventDataWithDefaults() *PaymentChargeUpdateEventData`

NewPaymentChargeUpdateEventDataWithDefaults instantiates a new PaymentChargeUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentChargeUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentChargeUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentChargeUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *PaymentChargeUpdateEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentChargeUpdateEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentChargeUpdateEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetActionId

`func (o *PaymentChargeUpdateEventData) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *PaymentChargeUpdateEventData) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *PaymentChargeUpdateEventData) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetPlanId

`func (o *PaymentChargeUpdateEventData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PaymentChargeUpdateEventData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PaymentChargeUpdateEventData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetSubscriptionId

`func (o *PaymentChargeUpdateEventData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentChargeUpdateEventData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentChargeUpdateEventData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetMerchantId

`func (o *PaymentChargeUpdateEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentChargeUpdateEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentChargeUpdateEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantAddress

`func (o *PaymentChargeUpdateEventData) GetMerchantAddress() string`

GetMerchantAddress returns the MerchantAddress field if non-nil, zero value otherwise.

### GetMerchantAddressOk

`func (o *PaymentChargeUpdateEventData) GetMerchantAddressOk() (*string, bool)`

GetMerchantAddressOk returns a tuple with the MerchantAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAddress

`func (o *PaymentChargeUpdateEventData) SetMerchantAddress(v string)`

SetMerchantAddress sets MerchantAddress field to given value.


### GetData

`func (o *PaymentChargeUpdateEventData) GetData() PaymentSubscriptionActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentChargeUpdateEventData) GetDataOk() (*PaymentSubscriptionActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentChargeUpdateEventData) SetData(v PaymentSubscriptionActionData)`

SetData sets Data field to given value.


### GetTransactionIds

`func (o *PaymentChargeUpdateEventData) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *PaymentChargeUpdateEventData) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *PaymentChargeUpdateEventData) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *PaymentChargeUpdateEventData) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentChargeUpdateEventData) GetStatus() PaymentSubscriptionActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentChargeUpdateEventData) GetStatusOk() (*PaymentSubscriptionActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentChargeUpdateEventData) SetStatus(v PaymentSubscriptionActionStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentChargeUpdateEventData) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentChargeUpdateEventData) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentChargeUpdateEventData) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentChargeUpdateEventData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PaymentChargeUpdateEventData) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentChargeUpdateEventData) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentChargeUpdateEventData) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PaymentChargeUpdateEventData) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



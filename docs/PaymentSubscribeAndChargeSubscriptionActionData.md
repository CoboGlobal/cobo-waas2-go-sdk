# PaymentSubscribeAndChargeSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PaymentSubscriptionActionType**](PaymentSubscriptionActionType.md) |  | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**ChargeAmount** | **string** | The subscription crypto amount, less than subscription plan amount.  | 
**Signature** | **string** | The signature for transaction. charge action is not required. | 
**Deadline** | **int32** | The signature deadline for transaction. charge action is not required. | 

## Methods

### NewPaymentSubscribeAndChargeSubscriptionActionData

`func NewPaymentSubscribeAndChargeSubscriptionActionData(actionType PaymentSubscriptionActionType, subscriptionId string, chargeAmount string, signature string, deadline int32, ) *PaymentSubscribeAndChargeSubscriptionActionData`

NewPaymentSubscribeAndChargeSubscriptionActionData instantiates a new PaymentSubscribeAndChargeSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSubscribeAndChargeSubscriptionActionDataWithDefaults

`func NewPaymentSubscribeAndChargeSubscriptionActionDataWithDefaults() *PaymentSubscribeAndChargeSubscriptionActionData`

NewPaymentSubscribeAndChargeSubscriptionActionDataWithDefaults instantiates a new PaymentSubscribeAndChargeSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetActionType() PaymentSubscriptionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) SetActionType(v PaymentSubscriptionActionType)`

SetActionType sets ActionType field to given value.


### GetSubscriptionId

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetChargeAmount

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetChargeAmount() string`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetChargeAmountOk() (*string, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) SetChargeAmount(v string)`

SetChargeAmount sets ChargeAmount field to given value.


### GetSignature

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetDeadline

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetDeadline() int32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) GetDeadlineOk() (*int32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *PaymentSubscribeAndChargeSubscriptionActionData) SetDeadline(v int32)`

SetDeadline sets Deadline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



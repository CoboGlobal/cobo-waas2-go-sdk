# PaymentUpdateAmountSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPlanId** | **string** | The new plan id in cobo. | 
**ActionType** | [**PaymentSubscriptionActionType**](PaymentSubscriptionActionType.md) |  | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**Signature** | **string** | The signature for transaction. | 

## Methods

### NewPaymentUpdateAmountSubscriptionActionData

`func NewPaymentUpdateAmountSubscriptionActionData(newPlanId string, actionType PaymentSubscriptionActionType, subscriptionId string, signature string, ) *PaymentUpdateAmountSubscriptionActionData`

NewPaymentUpdateAmountSubscriptionActionData instantiates a new PaymentUpdateAmountSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentUpdateAmountSubscriptionActionDataWithDefaults

`func NewPaymentUpdateAmountSubscriptionActionDataWithDefaults() *PaymentUpdateAmountSubscriptionActionData`

NewPaymentUpdateAmountSubscriptionActionDataWithDefaults instantiates a new PaymentUpdateAmountSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPlanId

`func (o *PaymentUpdateAmountSubscriptionActionData) GetNewPlanId() string`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *PaymentUpdateAmountSubscriptionActionData) GetNewPlanIdOk() (*string, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *PaymentUpdateAmountSubscriptionActionData) SetNewPlanId(v string)`

SetNewPlanId sets NewPlanId field to given value.


### GetActionType

`func (o *PaymentUpdateAmountSubscriptionActionData) GetActionType() PaymentSubscriptionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentUpdateAmountSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentUpdateAmountSubscriptionActionData) SetActionType(v PaymentSubscriptionActionType)`

SetActionType sets ActionType field to given value.


### GetSubscriptionId

`func (o *PaymentUpdateAmountSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentUpdateAmountSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentUpdateAmountSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSignature

`func (o *PaymentUpdateAmountSubscriptionActionData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentUpdateAmountSubscriptionActionData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentUpdateAmountSubscriptionActionData) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



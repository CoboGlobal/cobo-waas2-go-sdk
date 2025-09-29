# PaymentExtendPeriodSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Periods** | Pointer to **int32** | The periods needed updated. | [optional] 
**ActionType** | [**PaymentSubscriptionActionType**](PaymentSubscriptionActionType.md) |  | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**Signature** | **string** | The signature for transaction. | 

## Methods

### NewPaymentExtendPeriodSubscriptionActionData

`func NewPaymentExtendPeriodSubscriptionActionData(actionType PaymentSubscriptionActionType, subscriptionId string, signature string, ) *PaymentExtendPeriodSubscriptionActionData`

NewPaymentExtendPeriodSubscriptionActionData instantiates a new PaymentExtendPeriodSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentExtendPeriodSubscriptionActionDataWithDefaults

`func NewPaymentExtendPeriodSubscriptionActionDataWithDefaults() *PaymentExtendPeriodSubscriptionActionData`

NewPaymentExtendPeriodSubscriptionActionDataWithDefaults instantiates a new PaymentExtendPeriodSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriods

`func (o *PaymentExtendPeriodSubscriptionActionData) GetPeriods() int32`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *PaymentExtendPeriodSubscriptionActionData) GetPeriodsOk() (*int32, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *PaymentExtendPeriodSubscriptionActionData) SetPeriods(v int32)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *PaymentExtendPeriodSubscriptionActionData) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.

### GetActionType

`func (o *PaymentExtendPeriodSubscriptionActionData) GetActionType() PaymentSubscriptionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentExtendPeriodSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentExtendPeriodSubscriptionActionData) SetActionType(v PaymentSubscriptionActionType)`

SetActionType sets ActionType field to given value.


### GetSubscriptionId

`func (o *PaymentExtendPeriodSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentExtendPeriodSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentExtendPeriodSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSignature

`func (o *PaymentExtendPeriodSubscriptionActionData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentExtendPeriodSubscriptionActionData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentExtendPeriodSubscriptionActionData) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



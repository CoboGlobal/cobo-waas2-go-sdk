# PaymentBaseSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PaymentSubscriptionAction**](PaymentSubscriptionAction.md) |  | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**Signature** | **string** | The signature for transaction. | 

## Methods

### NewPaymentBaseSubscriptionActionData

`func NewPaymentBaseSubscriptionActionData(actionType PaymentSubscriptionAction, subscriptionId string, signature string, ) *PaymentBaseSubscriptionActionData`

NewPaymentBaseSubscriptionActionData instantiates a new PaymentBaseSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBaseSubscriptionActionDataWithDefaults

`func NewPaymentBaseSubscriptionActionDataWithDefaults() *PaymentBaseSubscriptionActionData`

NewPaymentBaseSubscriptionActionDataWithDefaults instantiates a new PaymentBaseSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PaymentBaseSubscriptionActionData) GetActionType() PaymentSubscriptionAction`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentBaseSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionAction, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentBaseSubscriptionActionData) SetActionType(v PaymentSubscriptionAction)`

SetActionType sets ActionType field to given value.


### GetSubscriptionId

`func (o *PaymentBaseSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentBaseSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentBaseSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSignature

`func (o *PaymentBaseSubscriptionActionData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentBaseSubscriptionActionData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentBaseSubscriptionActionData) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentDeveloperSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PaymentSubscriptionAction**](PaymentSubscriptionAction.md) |  | 
**SubscriptionId** | **string** | The subscription id in cobo. | 

## Methods

### NewPaymentDeveloperSubscriptionActionData

`func NewPaymentDeveloperSubscriptionActionData(actionType PaymentSubscriptionAction, subscriptionId string, ) *PaymentDeveloperSubscriptionActionData`

NewPaymentDeveloperSubscriptionActionData instantiates a new PaymentDeveloperSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDeveloperSubscriptionActionDataWithDefaults

`func NewPaymentDeveloperSubscriptionActionDataWithDefaults() *PaymentDeveloperSubscriptionActionData`

NewPaymentDeveloperSubscriptionActionDataWithDefaults instantiates a new PaymentDeveloperSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PaymentDeveloperSubscriptionActionData) GetActionType() PaymentSubscriptionAction`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentDeveloperSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionAction, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentDeveloperSubscriptionActionData) SetActionType(v PaymentSubscriptionAction)`

SetActionType sets ActionType field to given value.


### GetSubscriptionId

`func (o *PaymentDeveloperSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentDeveloperSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentDeveloperSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



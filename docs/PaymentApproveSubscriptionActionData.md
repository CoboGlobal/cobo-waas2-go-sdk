# PaymentApproveSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PaymentSubscriptionActionType**](PaymentSubscriptionActionType.md) |  | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**PermitData** | Pointer to **string** | The signature of permit. | [optional] 

## Methods

### NewPaymentApproveSubscriptionActionData

`func NewPaymentApproveSubscriptionActionData(actionType PaymentSubscriptionActionType, subscriptionId string, ) *PaymentApproveSubscriptionActionData`

NewPaymentApproveSubscriptionActionData instantiates a new PaymentApproveSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentApproveSubscriptionActionDataWithDefaults

`func NewPaymentApproveSubscriptionActionDataWithDefaults() *PaymentApproveSubscriptionActionData`

NewPaymentApproveSubscriptionActionDataWithDefaults instantiates a new PaymentApproveSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PaymentApproveSubscriptionActionData) GetActionType() PaymentSubscriptionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentApproveSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentApproveSubscriptionActionData) SetActionType(v PaymentSubscriptionActionType)`

SetActionType sets ActionType field to given value.


### GetSubscriptionId

`func (o *PaymentApproveSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentApproveSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentApproveSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPermitData

`func (o *PaymentApproveSubscriptionActionData) GetPermitData() string`

GetPermitData returns the PermitData field if non-nil, zero value otherwise.

### GetPermitDataOk

`func (o *PaymentApproveSubscriptionActionData) GetPermitDataOk() (*string, bool)`

GetPermitDataOk returns a tuple with the PermitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitData

`func (o *PaymentApproveSubscriptionActionData) SetPermitData(v string)`

SetPermitData sets PermitData field to given value.

### HasPermitData

`func (o *PaymentApproveSubscriptionActionData) HasPermitData() bool`

HasPermitData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



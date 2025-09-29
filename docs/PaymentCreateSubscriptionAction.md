# PaymentCreateSubscriptionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The subscription action request id. | 
**PlanId** | **string** | The subscription plan id in cobo. | 
**MerchantId** | **string** | The merchant id in cobo. | 
**Data** | Pointer to [**PaymentSubscriptionActionData**](PaymentSubscriptionActionData.md) |  | [optional] 

## Methods

### NewPaymentCreateSubscriptionAction

`func NewPaymentCreateSubscriptionAction(requestId string, planId string, merchantId string, ) *PaymentCreateSubscriptionAction`

NewPaymentCreateSubscriptionAction instantiates a new PaymentCreateSubscriptionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateSubscriptionActionWithDefaults

`func NewPaymentCreateSubscriptionActionWithDefaults() *PaymentCreateSubscriptionAction`

NewPaymentCreateSubscriptionActionWithDefaults instantiates a new PaymentCreateSubscriptionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *PaymentCreateSubscriptionAction) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentCreateSubscriptionAction) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentCreateSubscriptionAction) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetPlanId

`func (o *PaymentCreateSubscriptionAction) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PaymentCreateSubscriptionAction) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PaymentCreateSubscriptionAction) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetMerchantId

`func (o *PaymentCreateSubscriptionAction) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentCreateSubscriptionAction) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentCreateSubscriptionAction) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetData

`func (o *PaymentCreateSubscriptionAction) GetData() PaymentSubscriptionActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentCreateSubscriptionAction) GetDataOk() (*PaymentSubscriptionActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentCreateSubscriptionAction) SetData(v PaymentSubscriptionActionData)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentCreateSubscriptionAction) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



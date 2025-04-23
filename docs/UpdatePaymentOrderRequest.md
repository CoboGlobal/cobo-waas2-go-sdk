# UpdatePaymentOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expired** | **bool** | Whether to manually expire the order. If set to &#x60;true&#x60;, the order status will be updated to &#x60;Expired&#x60;. | 

## Methods

### NewUpdatePaymentOrderRequest

`func NewUpdatePaymentOrderRequest(expired bool, ) *UpdatePaymentOrderRequest`

NewUpdatePaymentOrderRequest instantiates a new UpdatePaymentOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentOrderRequestWithDefaults

`func NewUpdatePaymentOrderRequestWithDefaults() *UpdatePaymentOrderRequest`

NewUpdatePaymentOrderRequestWithDefaults instantiates a new UpdatePaymentOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpired

`func (o *UpdatePaymentOrderRequest) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *UpdatePaymentOrderRequest) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *UpdatePaymentOrderRequest) SetExpired(v bool)`

SetExpired sets Expired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



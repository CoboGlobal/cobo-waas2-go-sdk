# RetryPaymentCallbacksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessType** | [**PaymentCallbackBusinessType**](PaymentCallbackBusinessType.md) |  | 
**BusinessId** | **string** | The public ID of the business record, determined by &#x60;business_type&#x60;. - &#x60;refund&#x60;: The refund code, returned as &#x60;refund_id&#x60; when you create or retrieve a refund order. For example, &#x60;R20250304-M1001-1001&#x60;. - &#x60;payout&#x60;: The payout ID, returned as &#x60;payout_id&#x60; when you create or retrieve a payout. For example, &#x60;aff0e1cb-15b2-4e1f-9b9d-a9133715986f&#x60;. - &#x60;bulk_send&#x60;: The bulk send ID, returned as &#x60;bulk_send_id&#x60; when you create or retrieve a bulk send. For example, &#x60;123e4567-e89b-42d3-a456-426614174003&#x60;.  Do not provide an internal database ID, request ID, callback message ID, or refund link ID. The business record must belong to your organization.  | 

## Methods

### NewRetryPaymentCallbacksRequest

`func NewRetryPaymentCallbacksRequest(businessType PaymentCallbackBusinessType, businessId string, ) *RetryPaymentCallbacksRequest`

NewRetryPaymentCallbacksRequest instantiates a new RetryPaymentCallbacksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPaymentCallbacksRequestWithDefaults

`func NewRetryPaymentCallbacksRequestWithDefaults() *RetryPaymentCallbacksRequest`

NewRetryPaymentCallbacksRequestWithDefaults instantiates a new RetryPaymentCallbacksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessType

`func (o *RetryPaymentCallbacksRequest) GetBusinessType() PaymentCallbackBusinessType`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *RetryPaymentCallbacksRequest) GetBusinessTypeOk() (*PaymentCallbackBusinessType, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *RetryPaymentCallbacksRequest) SetBusinessType(v PaymentCallbackBusinessType)`

SetBusinessType sets BusinessType field to given value.


### GetBusinessId

`func (o *RetryPaymentCallbacksRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *RetryPaymentCallbacksRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *RetryPaymentCallbacksRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



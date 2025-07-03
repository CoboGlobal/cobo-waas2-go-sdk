# UpdateRefundByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToAddress** | **string** | The address where the refunded cryptocurrency will be sent. | 

## Methods

### NewUpdateRefundByIdRequest

`func NewUpdateRefundByIdRequest(toAddress string, ) *UpdateRefundByIdRequest`

NewUpdateRefundByIdRequest instantiates a new UpdateRefundByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRefundByIdRequestWithDefaults

`func NewUpdateRefundByIdRequestWithDefaults() *UpdateRefundByIdRequest`

NewUpdateRefundByIdRequestWithDefaults instantiates a new UpdateRefundByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToAddress

`func (o *UpdateRefundByIdRequest) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *UpdateRefundByIdRequest) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *UpdateRefundByIdRequest) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



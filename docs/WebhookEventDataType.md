# WebhookEventDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** | The data type of the event. When &#x60;data_type&#x60; is &#x60;Transaction&#x60;, it means the event uses the &#x60;transaction&#x60; schema as its data type. | 

## Methods

### NewWebhookEventDataType

`func NewWebhookEventDataType(dataType string, ) *WebhookEventDataType`

NewWebhookEventDataType instantiates a new WebhookEventDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventDataTypeWithDefaults

`func NewWebhookEventDataTypeWithDefaults() *WebhookEventDataType`

NewWebhookEventDataTypeWithDefaults instantiates a new WebhookEventDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *WebhookEventDataType) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *WebhookEventDataType) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *WebhookEventDataType) SetDataType(v string)`

SetDataType sets DataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



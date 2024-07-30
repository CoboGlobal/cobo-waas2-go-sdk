# ListWebhookEventLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WebhookEventLog**](WebhookEventLog.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListWebhookEventLogs200Response

`func NewListWebhookEventLogs200Response() *ListWebhookEventLogs200Response`

NewListWebhookEventLogs200Response instantiates a new ListWebhookEventLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhookEventLogs200ResponseWithDefaults

`func NewListWebhookEventLogs200ResponseWithDefaults() *ListWebhookEventLogs200Response`

NewListWebhookEventLogs200ResponseWithDefaults instantiates a new ListWebhookEventLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListWebhookEventLogs200Response) GetData() []WebhookEventLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWebhookEventLogs200Response) GetDataOk() (*[]WebhookEventLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWebhookEventLogs200Response) SetData(v []WebhookEventLog)`

SetData sets Data field to given value.

### HasData

`func (o *ListWebhookEventLogs200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListWebhookEventLogs200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListWebhookEventLogs200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListWebhookEventLogs200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListWebhookEventLogs200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



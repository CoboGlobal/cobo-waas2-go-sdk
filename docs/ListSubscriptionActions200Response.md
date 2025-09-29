# ListSubscriptionActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PaymentSubscriptionAction**](PaymentSubscriptionAction.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListSubscriptionActions200Response

`func NewListSubscriptionActions200Response() *ListSubscriptionActions200Response`

NewListSubscriptionActions200Response instantiates a new ListSubscriptionActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionActions200ResponseWithDefaults

`func NewListSubscriptionActions200ResponseWithDefaults() *ListSubscriptionActions200Response`

NewListSubscriptionActions200ResponseWithDefaults instantiates a new ListSubscriptionActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListSubscriptionActions200Response) GetData() []PaymentSubscriptionAction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSubscriptionActions200Response) GetDataOk() (*[]PaymentSubscriptionAction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSubscriptionActions200Response) SetData(v []PaymentSubscriptionAction)`

SetData sets Data field to given value.

### HasData

`func (o *ListSubscriptionActions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListSubscriptionActions200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSubscriptionActions200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSubscriptionActions200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSubscriptionActions200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



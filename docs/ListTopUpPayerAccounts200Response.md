# ListTopUpPayerAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PayerAccount**](PayerAccount.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListTopUpPayerAccounts200Response

`func NewListTopUpPayerAccounts200Response() *ListTopUpPayerAccounts200Response`

NewListTopUpPayerAccounts200Response instantiates a new ListTopUpPayerAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTopUpPayerAccounts200ResponseWithDefaults

`func NewListTopUpPayerAccounts200ResponseWithDefaults() *ListTopUpPayerAccounts200Response`

NewListTopUpPayerAccounts200ResponseWithDefaults instantiates a new ListTopUpPayerAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListTopUpPayerAccounts200Response) GetData() []PayerAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTopUpPayerAccounts200Response) GetDataOk() (*[]PayerAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTopUpPayerAccounts200Response) SetData(v []PayerAccount)`

SetData sets Data field to given value.

### HasData

`func (o *ListTopUpPayerAccounts200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListTopUpPayerAccounts200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTopUpPayerAccounts200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTopUpPayerAccounts200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListTopUpPayerAccounts200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



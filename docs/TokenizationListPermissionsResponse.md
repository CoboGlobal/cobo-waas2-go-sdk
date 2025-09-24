# TokenizationListPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TokenizationListPermissionsResponseDataInner**](TokenizationListPermissionsResponseDataInner.md) | List of permissions. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewTokenizationListPermissionsResponse

`func NewTokenizationListPermissionsResponse(data []TokenizationListPermissionsResponseDataInner, pagination Pagination, ) *TokenizationListPermissionsResponse`

NewTokenizationListPermissionsResponse instantiates a new TokenizationListPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationListPermissionsResponseWithDefaults

`func NewTokenizationListPermissionsResponseWithDefaults() *TokenizationListPermissionsResponse`

NewTokenizationListPermissionsResponseWithDefaults instantiates a new TokenizationListPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenizationListPermissionsResponse) GetData() []TokenizationListPermissionsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationListPermissionsResponse) GetDataOk() (*[]TokenizationListPermissionsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationListPermissionsResponse) SetData(v []TokenizationListPermissionsResponseDataInner)`

SetData sets Data field to given value.


### GetPagination

`func (o *TokenizationListPermissionsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenizationListPermissionsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenizationListPermissionsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



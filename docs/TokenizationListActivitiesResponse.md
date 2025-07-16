# TokenizationListActivitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TokenizationActivityInfo**](TokenizationActivityInfo.md) | The list of tokenization activities. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewTokenizationListActivitiesResponse

`func NewTokenizationListActivitiesResponse(data []TokenizationActivityInfo, pagination Pagination, ) *TokenizationListActivitiesResponse`

NewTokenizationListActivitiesResponse instantiates a new TokenizationListActivitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationListActivitiesResponseWithDefaults

`func NewTokenizationListActivitiesResponseWithDefaults() *TokenizationListActivitiesResponse`

NewTokenizationListActivitiesResponseWithDefaults instantiates a new TokenizationListActivitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenizationListActivitiesResponse) GetData() []TokenizationActivityInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationListActivitiesResponse) GetDataOk() (*[]TokenizationActivityInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationListActivitiesResponse) SetData(v []TokenizationActivityInfo)`

SetData sets Data field to given value.


### GetPagination

`func (o *TokenizationListActivitiesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenizationListActivitiesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenizationListActivitiesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



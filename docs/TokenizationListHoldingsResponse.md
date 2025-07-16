# TokenizationListHoldingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TokenizationHoldingInfo**](TokenizationHoldingInfo.md) | List of token holdings. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewTokenizationListHoldingsResponse

`func NewTokenizationListHoldingsResponse(data []TokenizationHoldingInfo, pagination Pagination, ) *TokenizationListHoldingsResponse`

NewTokenizationListHoldingsResponse instantiates a new TokenizationListHoldingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationListHoldingsResponseWithDefaults

`func NewTokenizationListHoldingsResponseWithDefaults() *TokenizationListHoldingsResponse`

NewTokenizationListHoldingsResponseWithDefaults instantiates a new TokenizationListHoldingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenizationListHoldingsResponse) GetData() []TokenizationHoldingInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationListHoldingsResponse) GetDataOk() (*[]TokenizationHoldingInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationListHoldingsResponse) SetData(v []TokenizationHoldingInfo)`

SetData sets Data field to given value.


### GetPagination

`func (o *TokenizationListHoldingsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenizationListHoldingsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenizationListHoldingsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



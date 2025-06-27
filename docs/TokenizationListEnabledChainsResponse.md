# TokenizationListEnabledChainsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **[]string** | The list of tokenization supported chains. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewTokenizationListEnabledChainsResponse

`func NewTokenizationListEnabledChainsResponse(data []string, pagination Pagination, ) *TokenizationListEnabledChainsResponse`

NewTokenizationListEnabledChainsResponse instantiates a new TokenizationListEnabledChainsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationListEnabledChainsResponseWithDefaults

`func NewTokenizationListEnabledChainsResponseWithDefaults() *TokenizationListEnabledChainsResponse`

NewTokenizationListEnabledChainsResponseWithDefaults instantiates a new TokenizationListEnabledChainsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenizationListEnabledChainsResponse) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationListEnabledChainsResponse) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationListEnabledChainsResponse) SetData(v []string)`

SetData sets Data field to given value.


### GetPagination

`func (o *TokenizationListEnabledChainsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenizationListEnabledChainsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenizationListEnabledChainsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



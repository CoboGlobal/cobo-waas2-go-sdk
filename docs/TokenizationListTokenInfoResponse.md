# TokenizationListTokenInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TokenizationTokenInfo**](TokenizationTokenInfo.md) | List tokens. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewTokenizationListTokenInfoResponse

`func NewTokenizationListTokenInfoResponse(data []TokenizationTokenInfo, pagination Pagination, ) *TokenizationListTokenInfoResponse`

NewTokenizationListTokenInfoResponse instantiates a new TokenizationListTokenInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationListTokenInfoResponseWithDefaults

`func NewTokenizationListTokenInfoResponseWithDefaults() *TokenizationListTokenInfoResponse`

NewTokenizationListTokenInfoResponseWithDefaults instantiates a new TokenizationListTokenInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenizationListTokenInfoResponse) GetData() []TokenizationTokenInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationListTokenInfoResponse) GetDataOk() (*[]TokenizationTokenInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationListTokenInfoResponse) SetData(v []TokenizationTokenInfo)`

SetData sets Data field to given value.


### GetPagination

`func (o *TokenizationListTokenInfoResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenizationListTokenInfoResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenizationListTokenInfoResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



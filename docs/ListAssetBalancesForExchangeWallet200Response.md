# ListAssetBalancesForExchangeWallet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SubWalletAssetBalance**](SubWalletAssetBalance.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListAssetBalancesForExchangeWallet200Response

`func NewListAssetBalancesForExchangeWallet200Response() *ListAssetBalancesForExchangeWallet200Response`

NewListAssetBalancesForExchangeWallet200Response instantiates a new ListAssetBalancesForExchangeWallet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetBalancesForExchangeWallet200ResponseWithDefaults

`func NewListAssetBalancesForExchangeWallet200ResponseWithDefaults() *ListAssetBalancesForExchangeWallet200Response`

NewListAssetBalancesForExchangeWallet200ResponseWithDefaults instantiates a new ListAssetBalancesForExchangeWallet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAssetBalancesForExchangeWallet200Response) GetData() []SubWalletAssetBalance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAssetBalancesForExchangeWallet200Response) GetDataOk() (*[]SubWalletAssetBalance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAssetBalancesForExchangeWallet200Response) SetData(v []SubWalletAssetBalance)`

SetData sets Data field to given value.

### HasData

`func (o *ListAssetBalancesForExchangeWallet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListAssetBalancesForExchangeWallet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAssetBalancesForExchangeWallet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAssetBalancesForExchangeWallet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAssetBalancesForExchangeWallet200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



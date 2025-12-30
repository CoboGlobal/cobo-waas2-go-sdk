# ListCounterpartyEntries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletAddresses** | Pointer to [**[]CounterpartyWalletAddressDetail**](CounterpartyWalletAddressDetail.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListCounterpartyEntries200Response

`func NewListCounterpartyEntries200Response() *ListCounterpartyEntries200Response`

NewListCounterpartyEntries200Response instantiates a new ListCounterpartyEntries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCounterpartyEntries200ResponseWithDefaults

`func NewListCounterpartyEntries200ResponseWithDefaults() *ListCounterpartyEntries200Response`

NewListCounterpartyEntries200ResponseWithDefaults instantiates a new ListCounterpartyEntries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletAddresses

`func (o *ListCounterpartyEntries200Response) GetWalletAddresses() []CounterpartyWalletAddressDetail`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *ListCounterpartyEntries200Response) GetWalletAddressesOk() (*[]CounterpartyWalletAddressDetail, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *ListCounterpartyEntries200Response) SetWalletAddresses(v []CounterpartyWalletAddressDetail)`

SetWalletAddresses sets WalletAddresses field to given value.

### HasWalletAddresses

`func (o *ListCounterpartyEntries200Response) HasWalletAddresses() bool`

HasWalletAddresses returns a boolean if a field has been set.

### GetPagination

`func (o *ListCounterpartyEntries200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListCounterpartyEntries200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListCounterpartyEntries200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListCounterpartyEntries200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



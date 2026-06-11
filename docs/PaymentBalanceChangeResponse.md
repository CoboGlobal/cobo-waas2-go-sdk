# PaymentBalanceChangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PaymentBalanceChange**](PaymentBalanceChange.md) | The list of balance changes. | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewPaymentBalanceChangeResponse

`func NewPaymentBalanceChangeResponse() *PaymentBalanceChangeResponse`

NewPaymentBalanceChangeResponse instantiates a new PaymentBalanceChangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBalanceChangeResponseWithDefaults

`func NewPaymentBalanceChangeResponseWithDefaults() *PaymentBalanceChangeResponse`

NewPaymentBalanceChangeResponseWithDefaults instantiates a new PaymentBalanceChangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentBalanceChangeResponse) GetData() []PaymentBalanceChange`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentBalanceChangeResponse) GetDataOk() (*[]PaymentBalanceChange, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentBalanceChangeResponse) SetData(v []PaymentBalanceChange)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentBalanceChangeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PaymentBalanceChangeResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaymentBalanceChangeResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaymentBalanceChangeResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PaymentBalanceChangeResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



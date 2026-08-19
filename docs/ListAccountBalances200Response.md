# ListAccountBalances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BalanceAtBlock**](BalanceAtBlock.md) |  | [optional] 

## Methods

### NewListAccountBalances200Response

`func NewListAccountBalances200Response() *ListAccountBalances200Response`

NewListAccountBalances200Response instantiates a new ListAccountBalances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountBalances200ResponseWithDefaults

`func NewListAccountBalances200ResponseWithDefaults() *ListAccountBalances200Response`

NewListAccountBalances200ResponseWithDefaults instantiates a new ListAccountBalances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAccountBalances200Response) GetData() []BalanceAtBlock`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAccountBalances200Response) GetDataOk() (*[]BalanceAtBlock, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAccountBalances200Response) SetData(v []BalanceAtBlock)`

SetData sets Data field to given value.

### HasData

`func (o *ListAccountBalances200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



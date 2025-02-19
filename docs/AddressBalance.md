# AddressBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address. | 
**Balance** | [**Balance**](Balance.md) |  | 

## Methods

### NewAddressBalance

`func NewAddressBalance(address string, balance Balance, ) *AddressBalance`

NewAddressBalance instantiates a new AddressBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBalanceWithDefaults

`func NewAddressBalanceWithDefaults() *AddressBalance`

NewAddressBalanceWithDefaults instantiates a new AddressBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressBalance) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressBalance) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressBalance) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *AddressBalance) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AddressBalance) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AddressBalance) SetBalance(v Balance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



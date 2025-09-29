# ReceivedAmountPerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Receiving address. | 
**TotalReceivedAmount** | **string** | Total tokens received at this address, as a decimal string. | 

## Methods

### NewReceivedAmountPerAddress

`func NewReceivedAmountPerAddress(address string, totalReceivedAmount string, ) *ReceivedAmountPerAddress`

NewReceivedAmountPerAddress instantiates a new ReceivedAmountPerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedAmountPerAddressWithDefaults

`func NewReceivedAmountPerAddressWithDefaults() *ReceivedAmountPerAddress`

NewReceivedAmountPerAddressWithDefaults instantiates a new ReceivedAmountPerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ReceivedAmountPerAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ReceivedAmountPerAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ReceivedAmountPerAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTotalReceivedAmount

`func (o *ReceivedAmountPerAddress) GetTotalReceivedAmount() string`

GetTotalReceivedAmount returns the TotalReceivedAmount field if non-nil, zero value otherwise.

### GetTotalReceivedAmountOk

`func (o *ReceivedAmountPerAddress) GetTotalReceivedAmountOk() (*string, bool)`

GetTotalReceivedAmountOk returns a tuple with the TotalReceivedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReceivedAmount

`func (o *ReceivedAmountPerAddress) SetTotalReceivedAmount(v string)`

SetTotalReceivedAmount sets TotalReceivedAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



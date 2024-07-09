# EstimateFeeContractCallDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The destination address. | [optional] 
**Value** | Pointer to **string** | The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Calldata** | Pointer to **string** | The data that is used to invoke a specific function or method within the specified contract at the destination address.  | [optional] 

## Methods

### NewEstimateFeeContractCallDestination

`func NewEstimateFeeContractCallDestination() *EstimateFeeContractCallDestination`

NewEstimateFeeContractCallDestination instantiates a new EstimateFeeContractCallDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateFeeContractCallDestinationWithDefaults

`func NewEstimateFeeContractCallDestinationWithDefaults() *EstimateFeeContractCallDestination`

NewEstimateFeeContractCallDestinationWithDefaults instantiates a new EstimateFeeContractCallDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EstimateFeeContractCallDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EstimateFeeContractCallDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EstimateFeeContractCallDestination) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EstimateFeeContractCallDestination) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetValue

`func (o *EstimateFeeContractCallDestination) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EstimateFeeContractCallDestination) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EstimateFeeContractCallDestination) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EstimateFeeContractCallDestination) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *EstimateFeeContractCallDestination) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *EstimateFeeContractCallDestination) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *EstimateFeeContractCallDestination) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.

### HasCalldata

`func (o *EstimateFeeContractCallDestination) HasCalldata() bool`

HasCalldata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


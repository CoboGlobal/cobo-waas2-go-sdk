# AddressTransferDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransferDestinationType**](TransferDestinationType.md) |  | 
**AddressStr** | **string** | Destination address | 
**Memo** | Pointer to **string** | Destination address memo | [optional] 

## Methods

### NewAddressTransferDestination

`func NewAddressTransferDestination(destinationType TransferDestinationType, addressStr string, ) *AddressTransferDestination`

NewAddressTransferDestination instantiates a new AddressTransferDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferDestinationWithDefaults

`func NewAddressTransferDestinationWithDefaults() *AddressTransferDestination`

NewAddressTransferDestinationWithDefaults instantiates a new AddressTransferDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *AddressTransferDestination) GetDestinationType() TransferDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *AddressTransferDestination) GetDestinationTypeOk() (*TransferDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *AddressTransferDestination) SetDestinationType(v TransferDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetAddressStr

`func (o *AddressTransferDestination) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *AddressTransferDestination) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *AddressTransferDestination) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


### GetMemo

`func (o *AddressTransferDestination) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AddressTransferDestination) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AddressTransferDestination) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *AddressTransferDestination) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TransactionEvmContractDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**Address** | **string** | The destination address. | 
**Value** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Calldata** | **string** | The data used to invoke a specific function or method within the specified contract at the destination address, with a maximum length of 65,000 characters.  | 
**CalldataInfo** | Pointer to [**TransactionEvmCalldataInfo**](TransactionEvmCalldataInfo.md) |  | [optional] 

## Methods

### NewTransactionEvmContractDestination

`func NewTransactionEvmContractDestination(destinationType TransactionDestinationType, address string, calldata string, ) *TransactionEvmContractDestination`

NewTransactionEvmContractDestination instantiates a new TransactionEvmContractDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEvmContractDestinationWithDefaults

`func NewTransactionEvmContractDestinationWithDefaults() *TransactionEvmContractDestination`

NewTransactionEvmContractDestinationWithDefaults instantiates a new TransactionEvmContractDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionEvmContractDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionEvmContractDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionEvmContractDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetAddress

`func (o *TransactionEvmContractDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionEvmContractDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionEvmContractDestination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetValue

`func (o *TransactionEvmContractDestination) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransactionEvmContractDestination) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransactionEvmContractDestination) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TransactionEvmContractDestination) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *TransactionEvmContractDestination) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *TransactionEvmContractDestination) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *TransactionEvmContractDestination) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetCalldataInfo

`func (o *TransactionEvmContractDestination) GetCalldataInfo() TransactionEvmCalldataInfo`

GetCalldataInfo returns the CalldataInfo field if non-nil, zero value otherwise.

### GetCalldataInfoOk

`func (o *TransactionEvmContractDestination) GetCalldataInfoOk() (*TransactionEvmCalldataInfo, bool)`

GetCalldataInfoOk returns a tuple with the CalldataInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldataInfo

`func (o *TransactionEvmContractDestination) SetCalldataInfo(v TransactionEvmCalldataInfo)`

SetCalldataInfo sets CalldataInfo field to given value.

### HasCalldataInfo

`func (o *TransactionEvmContractDestination) HasCalldataInfo() bool`

HasCalldataInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



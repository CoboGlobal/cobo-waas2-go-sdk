# CreateCounterpartyEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyId** | **string** | The counterparty ID. | 
**WalletAddresses** | Pointer to [**[]CreateWalletAddress**](CreateWalletAddress.md) | The wallet addresses of the counterparty. | [optional] 

## Methods

### NewCreateCounterpartyEntryRequest

`func NewCreateCounterpartyEntryRequest(counterpartyId string, ) *CreateCounterpartyEntryRequest`

NewCreateCounterpartyEntryRequest instantiates a new CreateCounterpartyEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCounterpartyEntryRequestWithDefaults

`func NewCreateCounterpartyEntryRequestWithDefaults() *CreateCounterpartyEntryRequest`

NewCreateCounterpartyEntryRequestWithDefaults instantiates a new CreateCounterpartyEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyId

`func (o *CreateCounterpartyEntryRequest) GetCounterpartyId() string`

GetCounterpartyId returns the CounterpartyId field if non-nil, zero value otherwise.

### GetCounterpartyIdOk

`func (o *CreateCounterpartyEntryRequest) GetCounterpartyIdOk() (*string, bool)`

GetCounterpartyIdOk returns a tuple with the CounterpartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyId

`func (o *CreateCounterpartyEntryRequest) SetCounterpartyId(v string)`

SetCounterpartyId sets CounterpartyId field to given value.


### GetWalletAddresses

`func (o *CreateCounterpartyEntryRequest) GetWalletAddresses() []CreateWalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *CreateCounterpartyEntryRequest) GetWalletAddressesOk() (*[]CreateWalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *CreateCounterpartyEntryRequest) SetWalletAddresses(v []CreateWalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.

### HasWalletAddresses

`func (o *CreateCounterpartyEntryRequest) HasWalletAddresses() bool`

HasWalletAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateCounterpartyEntry201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyId** | **string** | The counterparty ID. | 
**WalletAddresses** | [**[]WalletAddress**](WalletAddress.md) | The wallet addresses of the counterparty. | 

## Methods

### NewCreateCounterpartyEntry201Response

`func NewCreateCounterpartyEntry201Response(counterpartyId string, walletAddresses []WalletAddress, ) *CreateCounterpartyEntry201Response`

NewCreateCounterpartyEntry201Response instantiates a new CreateCounterpartyEntry201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCounterpartyEntry201ResponseWithDefaults

`func NewCreateCounterpartyEntry201ResponseWithDefaults() *CreateCounterpartyEntry201Response`

NewCreateCounterpartyEntry201ResponseWithDefaults instantiates a new CreateCounterpartyEntry201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyId

`func (o *CreateCounterpartyEntry201Response) GetCounterpartyId() string`

GetCounterpartyId returns the CounterpartyId field if non-nil, zero value otherwise.

### GetCounterpartyIdOk

`func (o *CreateCounterpartyEntry201Response) GetCounterpartyIdOk() (*string, bool)`

GetCounterpartyIdOk returns a tuple with the CounterpartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyId

`func (o *CreateCounterpartyEntry201Response) SetCounterpartyId(v string)`

SetCounterpartyId sets CounterpartyId field to given value.


### GetWalletAddresses

`func (o *CreateCounterpartyEntry201Response) GetWalletAddresses() []WalletAddress`

GetWalletAddresses returns the WalletAddresses field if non-nil, zero value otherwise.

### GetWalletAddressesOk

`func (o *CreateCounterpartyEntry201Response) GetWalletAddressesOk() (*[]WalletAddress, bool)`

GetWalletAddressesOk returns a tuple with the WalletAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddresses

`func (o *CreateCounterpartyEntry201Response) SetWalletAddresses(v []WalletAddress)`

SetWalletAddresses sets WalletAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



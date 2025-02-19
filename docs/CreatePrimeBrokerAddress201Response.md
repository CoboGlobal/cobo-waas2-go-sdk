# CreatePrimeBrokerAddress201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** | guard&#39;s pubkey. | 
**Addresses** | Pointer to [**[]QueryGuardPubkey200ResponseAddressesInner**](QueryGuardPubkey200ResponseAddressesInner.md) |  | [optional] 

## Methods

### NewCreatePrimeBrokerAddress201Response

`func NewCreatePrimeBrokerAddress201Response(pubkey string, ) *CreatePrimeBrokerAddress201Response`

NewCreatePrimeBrokerAddress201Response instantiates a new CreatePrimeBrokerAddress201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrimeBrokerAddress201ResponseWithDefaults

`func NewCreatePrimeBrokerAddress201ResponseWithDefaults() *CreatePrimeBrokerAddress201Response`

NewCreatePrimeBrokerAddress201ResponseWithDefaults instantiates a new CreatePrimeBrokerAddress201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *CreatePrimeBrokerAddress201Response) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *CreatePrimeBrokerAddress201Response) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *CreatePrimeBrokerAddress201Response) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetAddresses

`func (o *CreatePrimeBrokerAddress201Response) GetAddresses() []QueryGuardPubkey200ResponseAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CreatePrimeBrokerAddress201Response) GetAddressesOk() (*[]QueryGuardPubkey200ResponseAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CreatePrimeBrokerAddress201Response) SetAddresses(v []QueryGuardPubkey200ResponseAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CreatePrimeBrokerAddress201Response) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



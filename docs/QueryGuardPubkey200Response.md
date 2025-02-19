# QueryGuardPubkey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to **string** | The Cobo Guard&#39;s public key. | [optional] 
**Status** | Pointer to [**GuardPubkeyStatus**](GuardPubkeyStatus.md) |  | [optional] 
**Addresses** | Pointer to [**[]QueryGuardPubkey200ResponseAddressesInner**](QueryGuardPubkey200ResponseAddressesInner.md) |  | [optional] 

## Methods

### NewQueryGuardPubkey200Response

`func NewQueryGuardPubkey200Response() *QueryGuardPubkey200Response`

NewQueryGuardPubkey200Response instantiates a new QueryGuardPubkey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryGuardPubkey200ResponseWithDefaults

`func NewQueryGuardPubkey200ResponseWithDefaults() *QueryGuardPubkey200Response`

NewQueryGuardPubkey200ResponseWithDefaults instantiates a new QueryGuardPubkey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *QueryGuardPubkey200Response) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *QueryGuardPubkey200Response) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *QueryGuardPubkey200Response) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *QueryGuardPubkey200Response) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetStatus

`func (o *QueryGuardPubkey200Response) GetStatus() GuardPubkeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryGuardPubkey200Response) GetStatusOk() (*GuardPubkeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryGuardPubkey200Response) SetStatus(v GuardPubkeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryGuardPubkey200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAddresses

`func (o *QueryGuardPubkey200Response) GetAddresses() []QueryGuardPubkey200ResponseAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *QueryGuardPubkey200Response) GetAddressesOk() (*[]QueryGuardPubkey200ResponseAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *QueryGuardPubkey200Response) SetAddresses(v []QueryGuardPubkey200ResponseAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *QueryGuardPubkey200Response) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



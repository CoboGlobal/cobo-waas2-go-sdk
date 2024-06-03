# SafeContractCallSourceAllOfDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | Pointer to **string** | The wallet id of the delegate. This is required when user initiate a transfer. | [optional] 
**WalletType** | Pointer to **string** | The wallet type of the delegate. This is required when user initiate a transfer. | [optional] 
**WalletAddress** | Pointer to **string** | The wallet address of the delegate. This is required when user initiate a transfer. | [optional] 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewSafeContractCallSourceAllOfDelegate

`func NewSafeContractCallSourceAllOfDelegate() *SafeContractCallSourceAllOfDelegate`

NewSafeContractCallSourceAllOfDelegate instantiates a new SafeContractCallSourceAllOfDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeContractCallSourceAllOfDelegateWithDefaults

`func NewSafeContractCallSourceAllOfDelegateWithDefaults() *SafeContractCallSourceAllOfDelegate`

NewSafeContractCallSourceAllOfDelegateWithDefaults instantiates a new SafeContractCallSourceAllOfDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *SafeContractCallSourceAllOfDelegate) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SafeContractCallSourceAllOfDelegate) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SafeContractCallSourceAllOfDelegate) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *SafeContractCallSourceAllOfDelegate) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetWalletType

`func (o *SafeContractCallSourceAllOfDelegate) GetWalletType() string`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *SafeContractCallSourceAllOfDelegate) GetWalletTypeOk() (*string, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *SafeContractCallSourceAllOfDelegate) SetWalletType(v string)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *SafeContractCallSourceAllOfDelegate) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletAddress

`func (o *SafeContractCallSourceAllOfDelegate) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *SafeContractCallSourceAllOfDelegate) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *SafeContractCallSourceAllOfDelegate) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *SafeContractCallSourceAllOfDelegate) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetMpcUsedKeyGroup

`func (o *SafeContractCallSourceAllOfDelegate) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *SafeContractCallSourceAllOfDelegate) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *SafeContractCallSourceAllOfDelegate) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *SafeContractCallSourceAllOfDelegate) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



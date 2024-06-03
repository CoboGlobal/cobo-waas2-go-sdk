# ContractCallSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** |  | 
**WalletId** | **string** | Unique id of the wallet to initiate contract call from. | 
**AddressStr** | **string** | From address | 
**MpcUsedKeyGroup** | [**MpcSigningGroup**](MpcSigningGroup.md) |  | 
**Delegate** | [**SafeContractCallSourceAllOfDelegate**](SafeContractCallSourceAllOfDelegate.md) |  | 

## Methods

### NewContractCallSource

`func NewContractCallSource(sourceType string, walletId string, addressStr string, mpcUsedKeyGroup MpcSigningGroup, delegate SafeContractCallSourceAllOfDelegate, ) *ContractCallSource`

NewContractCallSource instantiates a new ContractCallSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallSourceWithDefaults

`func NewContractCallSourceWithDefaults() *ContractCallSource`

NewContractCallSourceWithDefaults instantiates a new ContractCallSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *ContractCallSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ContractCallSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ContractCallSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *ContractCallSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ContractCallSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ContractCallSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddressStr

`func (o *ContractCallSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *ContractCallSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *ContractCallSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


### GetMpcUsedKeyGroup

`func (o *ContractCallSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *ContractCallSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *ContractCallSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.


### GetDelegate

`func (o *ContractCallSource) GetDelegate() SafeContractCallSourceAllOfDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *ContractCallSource) GetDelegateOk() (*SafeContractCallSourceAllOfDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *ContractCallSource) SetDelegate(v SafeContractCallSourceAllOfDelegate)`

SetDelegate sets Delegate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateSmartContractWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**WalletType** | **string** |  | 
**WalletSubtype** | **string** |  | 
**Label** | Pointer to **string** | The label of the wallet. | [optional] 
**ChainId** | **string** | The chain id the wallet is on. | 
**SmartContractWalletType** | [**SmartContractWalletType**](SmartContractWalletType.md) |  | [default to SMARTCONTRACTWALLETTYPE_SAFE_WALLET]
**SafeAddress** | Pointer to **string** | The address of the smart contract wallet. If this is not provided, WaaS 2.0 will create a new safe wallet and setup cobo safe module for user. In this case, threshold, owners is required. | [optional] 
**Owners** | Pointer to **[]string** | The owners of the smart contract wallet. This MUST be provided when user want to create a new safe wallet. | [optional] 
**Threshold** | Pointer to **int32** | The threshold of required confirmations for the smart contract wallet. This MUST be provided when user want to create a new safe wallet. | [optional] 
**CoboSafeAddress** | Pointer to **string** | The address of the cobo safe module. Cobo safe module must has been created &amp; enabled when import a existing safe wallet. | [optional] 
**Initiator** | Pointer to [**SafeWalletAllOfInitiator**](SafeWalletAllOfInitiator.md) |  | [optional] 

## Methods

### NewCreateSmartContractWallet

`func NewCreateSmartContractWallet(name string, walletType string, walletSubtype string, chainId string, smartContractWalletType SmartContractWalletType, ) *CreateSmartContractWallet`

NewCreateSmartContractWallet instantiates a new CreateSmartContractWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSmartContractWalletWithDefaults

`func NewCreateSmartContractWalletWithDefaults() *CreateSmartContractWallet`

NewCreateSmartContractWalletWithDefaults instantiates a new CreateSmartContractWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSmartContractWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSmartContractWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSmartContractWallet) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreateSmartContractWallet) GetWalletType() string`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateSmartContractWallet) GetWalletTypeOk() (*string, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateSmartContractWallet) SetWalletType(v string)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateSmartContractWallet) GetWalletSubtype() string`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateSmartContractWallet) GetWalletSubtypeOk() (*string, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateSmartContractWallet) SetWalletSubtype(v string)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetLabel

`func (o *CreateSmartContractWallet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateSmartContractWallet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateSmartContractWallet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateSmartContractWallet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetChainId

`func (o *CreateSmartContractWallet) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateSmartContractWallet) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateSmartContractWallet) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSmartContractWalletType

`func (o *CreateSmartContractWallet) GetSmartContractWalletType() SmartContractWalletType`

GetSmartContractWalletType returns the SmartContractWalletType field if non-nil, zero value otherwise.

### GetSmartContractWalletTypeOk

`func (o *CreateSmartContractWallet) GetSmartContractWalletTypeOk() (*SmartContractWalletType, bool)`

GetSmartContractWalletTypeOk returns a tuple with the SmartContractWalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContractWalletType

`func (o *CreateSmartContractWallet) SetSmartContractWalletType(v SmartContractWalletType)`

SetSmartContractWalletType sets SmartContractWalletType field to given value.


### GetSafeAddress

`func (o *CreateSmartContractWallet) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *CreateSmartContractWallet) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *CreateSmartContractWallet) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.

### HasSafeAddress

`func (o *CreateSmartContractWallet) HasSafeAddress() bool`

HasSafeAddress returns a boolean if a field has been set.

### GetOwners

`func (o *CreateSmartContractWallet) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *CreateSmartContractWallet) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *CreateSmartContractWallet) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *CreateSmartContractWallet) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetThreshold

`func (o *CreateSmartContractWallet) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateSmartContractWallet) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateSmartContractWallet) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *CreateSmartContractWallet) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCoboSafeAddress

`func (o *CreateSmartContractWallet) GetCoboSafeAddress() string`

GetCoboSafeAddress returns the CoboSafeAddress field if non-nil, zero value otherwise.

### GetCoboSafeAddressOk

`func (o *CreateSmartContractWallet) GetCoboSafeAddressOk() (*string, bool)`

GetCoboSafeAddressOk returns a tuple with the CoboSafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboSafeAddress

`func (o *CreateSmartContractWallet) SetCoboSafeAddress(v string)`

SetCoboSafeAddress sets CoboSafeAddress field to given value.

### HasCoboSafeAddress

`func (o *CreateSmartContractWallet) HasCoboSafeAddress() bool`

HasCoboSafeAddress returns a boolean if a field has been set.

### GetInitiator

`func (o *CreateSmartContractWallet) GetInitiator() SafeWalletAllOfInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *CreateSmartContractWallet) GetInitiatorOk() (*SafeWalletAllOfInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *CreateSmartContractWallet) SetInitiator(v SafeWalletAllOfInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *CreateSmartContractWallet) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



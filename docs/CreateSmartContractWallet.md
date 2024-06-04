# CreateSmartContractWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The wallet name. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Label** | Pointer to **string** | The wallet label. | [optional] 
**ChainId** | **string** | The ID of the chain that the wallet operates on. | 
**SmartContractWalletType** | [**SmartContractWalletType**](SmartContractWalletType.md) |  | [default to SMARTCONTRACTWALLETTYPE_SAFE_WALLET]
**SafeAddress** | Pointer to **string** | The address of the Smart Contract Wallet. If this is not provided, Cobo will create a new Safe{Wallet} and set up Cobo Safe for you. In that case, the &#x60;threshold&#x60; and &#x60;owners&#x60; fields are required. | [optional] 
**Owners** | Pointer to **[]string** | The owners of the Smart Contract Wallet. This field is required when creating a new Safe{Wallet}. | [optional] 
**Threshold** | Pointer to **int32** | The minimum number of confirmations required for the Smart Contract Wallet. This field is required when creating a new Safe{Wallet}. | [optional] 
**CoboSafeAddress** | Pointer to **string** | The address of Cobo Safe. If you are importing an existing Safe{Wallet}, Cobo Safe must has been created and enabled. | [optional] 
**Initiator** | Pointer to [**CreateSafeWalletAllOfInitiator**](CreateSafeWalletAllOfInitiator.md) |  | [optional] 

## Methods

### NewCreateSmartContractWallet

`func NewCreateSmartContractWallet(name string, walletType WalletType, walletSubtype WalletSubtype, chainId string, smartContractWalletType SmartContractWalletType, ) *CreateSmartContractWallet`

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

`func (o *CreateSmartContractWallet) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateSmartContractWallet) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateSmartContractWallet) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateSmartContractWallet) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateSmartContractWallet) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateSmartContractWallet) SetWalletSubtype(v WalletSubtype)`

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

`func (o *CreateSmartContractWallet) GetInitiator() CreateSafeWalletAllOfInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *CreateSmartContractWallet) GetInitiatorOk() (*CreateSafeWalletAllOfInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *CreateSmartContractWallet) SetInitiator(v CreateSafeWalletAllOfInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *CreateSmartContractWallet) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



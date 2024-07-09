# CreateSafeWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The wallet name. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Label** | Pointer to **string** | The wallet label. | [optional] 
**ChainId** | **string** | The ID of the chain that the wallet operates on. | 
**SmartContractWalletType** | [**SmartContractWalletType**](SmartContractWalletType.md) |  | [default to SMARTCONTRACTWALLETTYPE_SAFE_WALLET]
**SafeAddress** | Pointer to **string** | The address of the Smart Contract Wallet. If this is not provided, Cobo will create a new Safe{Wallet} and set up Cobo Safe for you. In that case, the &#x60;threshold&#x60; and &#x60;signers&#x60; properties are required. | [optional] 
**Signers** | Pointer to **[]string** | The signers of the Smart Contract Wallet. This property is required when creating a new Safe{Wallet}. | [optional] 
**Threshold** | Pointer to **int32** | The minimum number of confirmations required for the Smart Contract Wallet. This property is required when creating a new Safe{Wallet}. | [optional] 
**CoboSafeAddress** | Pointer to **string** | The address of Cobo Safe. If you are importing an existing Safe{Wallet}, Cobo Safe must have been created and enabled. | [optional] 
**Initiator** | Pointer to [**CreateSafeWalletAllOfInitiator**](CreateSafeWalletAllOfInitiator.md) |  | [optional] 

## Methods

### NewCreateSafeWallet

`func NewCreateSafeWallet(name string, walletType WalletType, walletSubtype WalletSubtype, chainId string, smartContractWalletType SmartContractWalletType, ) *CreateSafeWallet`

NewCreateSafeWallet instantiates a new CreateSafeWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSafeWalletWithDefaults

`func NewCreateSafeWalletWithDefaults() *CreateSafeWallet`

NewCreateSafeWalletWithDefaults instantiates a new CreateSafeWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSafeWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSafeWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSafeWallet) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreateSafeWallet) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateSafeWallet) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateSafeWallet) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateSafeWallet) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateSafeWallet) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateSafeWallet) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetLabel

`func (o *CreateSafeWallet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateSafeWallet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateSafeWallet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateSafeWallet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetChainId

`func (o *CreateSafeWallet) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateSafeWallet) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateSafeWallet) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSmartContractWalletType

`func (o *CreateSafeWallet) GetSmartContractWalletType() SmartContractWalletType`

GetSmartContractWalletType returns the SmartContractWalletType field if non-nil, zero value otherwise.

### GetSmartContractWalletTypeOk

`func (o *CreateSafeWallet) GetSmartContractWalletTypeOk() (*SmartContractWalletType, bool)`

GetSmartContractWalletTypeOk returns a tuple with the SmartContractWalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContractWalletType

`func (o *CreateSafeWallet) SetSmartContractWalletType(v SmartContractWalletType)`

SetSmartContractWalletType sets SmartContractWalletType field to given value.


### GetSafeAddress

`func (o *CreateSafeWallet) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *CreateSafeWallet) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *CreateSafeWallet) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.

### HasSafeAddress

`func (o *CreateSafeWallet) HasSafeAddress() bool`

HasSafeAddress returns a boolean if a field has been set.

### GetSigners

`func (o *CreateSafeWallet) GetSigners() []string`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *CreateSafeWallet) GetSignersOk() (*[]string, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *CreateSafeWallet) SetSigners(v []string)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *CreateSafeWallet) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetThreshold

`func (o *CreateSafeWallet) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateSafeWallet) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateSafeWallet) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *CreateSafeWallet) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCoboSafeAddress

`func (o *CreateSafeWallet) GetCoboSafeAddress() string`

GetCoboSafeAddress returns the CoboSafeAddress field if non-nil, zero value otherwise.

### GetCoboSafeAddressOk

`func (o *CreateSafeWallet) GetCoboSafeAddressOk() (*string, bool)`

GetCoboSafeAddressOk returns a tuple with the CoboSafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboSafeAddress

`func (o *CreateSafeWallet) SetCoboSafeAddress(v string)`

SetCoboSafeAddress sets CoboSafeAddress field to given value.

### HasCoboSafeAddress

`func (o *CreateSafeWallet) HasCoboSafeAddress() bool`

HasCoboSafeAddress returns a boolean if a field has been set.

### GetInitiator

`func (o *CreateSafeWallet) GetInitiator() CreateSafeWalletAllOfInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *CreateSafeWallet) GetInitiatorOk() (*CreateSafeWalletAllOfInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *CreateSafeWallet) SetInitiator(v CreateSafeWalletAllOfInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *CreateSafeWallet) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



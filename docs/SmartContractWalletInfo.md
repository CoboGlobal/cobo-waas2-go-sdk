# SmartContractWalletInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Name** | **string** | The wallet name. | 
**OrgId** | **string** | The ID of the owning organization. | 
**ChainId** | Pointer to **string** | The chain the wallet operates on. | [optional] 
**Label** | Pointer to **string** | The wallet label. | [optional] 
**SmartContractWalletType** | [**SmartContractWalletType**](SmartContractWalletType.md) |  | [default to SMARTCONTRACTWALLETTYPE_SAFE_WALLET]
**SafeAddress** | Pointer to **string** | The Smart Contract Wallet address. | [optional] 
**Owners** | Pointer to **[]string** | The owners of the Smart Contract Wallet. This field is required when creating a new Safe{Wallet}. | [optional] 
**Threshold** | Pointer to **int32** | The minimum number of confirmations required for the Smart Contract Wallet. This field is required when creating a new Safe{Wallet}.  | [optional] 
**CoboSafeAddress** | Pointer to **string** | The address of Cobo Safe. | [optional] 
**Initiator** | Pointer to [**SafeWalletAllOfInitiator**](SafeWalletAllOfInitiator.md) |  | [optional] 

## Methods

### NewSmartContractWalletInfo

`func NewSmartContractWalletInfo(walletId string, walletType WalletType, walletSubtype WalletSubtype, name string, orgId string, smartContractWalletType SmartContractWalletType, ) *SmartContractWalletInfo`

NewSmartContractWalletInfo instantiates a new SmartContractWalletInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractWalletInfoWithDefaults

`func NewSmartContractWalletInfoWithDefaults() *SmartContractWalletInfo`

NewSmartContractWalletInfoWithDefaults instantiates a new SmartContractWalletInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *SmartContractWalletInfo) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SmartContractWalletInfo) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SmartContractWalletInfo) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *SmartContractWalletInfo) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *SmartContractWalletInfo) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *SmartContractWalletInfo) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *SmartContractWalletInfo) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *SmartContractWalletInfo) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *SmartContractWalletInfo) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetName

`func (o *SmartContractWalletInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmartContractWalletInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmartContractWalletInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *SmartContractWalletInfo) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SmartContractWalletInfo) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SmartContractWalletInfo) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetChainId

`func (o *SmartContractWalletInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SmartContractWalletInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SmartContractWalletInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *SmartContractWalletInfo) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetLabel

`func (o *SmartContractWalletInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SmartContractWalletInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SmartContractWalletInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SmartContractWalletInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSmartContractWalletType

`func (o *SmartContractWalletInfo) GetSmartContractWalletType() SmartContractWalletType`

GetSmartContractWalletType returns the SmartContractWalletType field if non-nil, zero value otherwise.

### GetSmartContractWalletTypeOk

`func (o *SmartContractWalletInfo) GetSmartContractWalletTypeOk() (*SmartContractWalletType, bool)`

GetSmartContractWalletTypeOk returns a tuple with the SmartContractWalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContractWalletType

`func (o *SmartContractWalletInfo) SetSmartContractWalletType(v SmartContractWalletType)`

SetSmartContractWalletType sets SmartContractWalletType field to given value.


### GetSafeAddress

`func (o *SmartContractWalletInfo) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *SmartContractWalletInfo) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *SmartContractWalletInfo) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.

### HasSafeAddress

`func (o *SmartContractWalletInfo) HasSafeAddress() bool`

HasSafeAddress returns a boolean if a field has been set.

### GetOwners

`func (o *SmartContractWalletInfo) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *SmartContractWalletInfo) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *SmartContractWalletInfo) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *SmartContractWalletInfo) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetThreshold

`func (o *SmartContractWalletInfo) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SmartContractWalletInfo) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SmartContractWalletInfo) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *SmartContractWalletInfo) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCoboSafeAddress

`func (o *SmartContractWalletInfo) GetCoboSafeAddress() string`

GetCoboSafeAddress returns the CoboSafeAddress field if non-nil, zero value otherwise.

### GetCoboSafeAddressOk

`func (o *SmartContractWalletInfo) GetCoboSafeAddressOk() (*string, bool)`

GetCoboSafeAddressOk returns a tuple with the CoboSafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboSafeAddress

`func (o *SmartContractWalletInfo) SetCoboSafeAddress(v string)`

SetCoboSafeAddress sets CoboSafeAddress field to given value.

### HasCoboSafeAddress

`func (o *SmartContractWalletInfo) HasCoboSafeAddress() bool`

HasCoboSafeAddress returns a boolean if a field has been set.

### GetInitiator

`func (o *SmartContractWalletInfo) GetInitiator() SafeWalletAllOfInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *SmartContractWalletInfo) GetInitiatorOk() (*SafeWalletAllOfInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *SmartContractWalletInfo) SetInitiator(v SafeWalletAllOfInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *SmartContractWalletInfo) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



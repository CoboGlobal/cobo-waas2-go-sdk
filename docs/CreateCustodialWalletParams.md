# CreateCustodialWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The wallet name. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**EnableAutoSweep** | Pointer to **bool** | Enable the auto-sweep feature for the wallet. This parameter only applies to MPC Wallets and Custodial Wallets (Web3 Wallets). | [optional] 

## Methods

### NewCreateCustodialWalletParams

`func NewCreateCustodialWalletParams(name string, walletType WalletType, walletSubtype WalletSubtype, ) *CreateCustodialWalletParams`

NewCreateCustodialWalletParams instantiates a new CreateCustodialWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustodialWalletParamsWithDefaults

`func NewCreateCustodialWalletParamsWithDefaults() *CreateCustodialWalletParams`

NewCreateCustodialWalletParamsWithDefaults instantiates a new CreateCustodialWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustodialWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustodialWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustodialWalletParams) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreateCustodialWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateCustodialWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateCustodialWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateCustodialWalletParams) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateCustodialWalletParams) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateCustodialWalletParams) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetEnableAutoSweep

`func (o *CreateCustodialWalletParams) GetEnableAutoSweep() bool`

GetEnableAutoSweep returns the EnableAutoSweep field if non-nil, zero value otherwise.

### GetEnableAutoSweepOk

`func (o *CreateCustodialWalletParams) GetEnableAutoSweepOk() (*bool, bool)`

GetEnableAutoSweepOk returns a tuple with the EnableAutoSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSweep

`func (o *CreateCustodialWalletParams) SetEnableAutoSweep(v bool)`

SetEnableAutoSweep sets EnableAutoSweep field to given value.

### HasEnableAutoSweep

`func (o *CreateCustodialWalletParams) HasEnableAutoSweep() bool`

HasEnableAutoSweep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



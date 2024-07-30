# UpdateCustodialWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletType** | [**WalletType**](WalletType.md) |  | 
**Name** | Pointer to **string** | The wallet name. | [optional] 

## Methods

### NewUpdateCustodialWalletParams

`func NewUpdateCustodialWalletParams(walletType WalletType, ) *UpdateCustodialWalletParams`

NewUpdateCustodialWalletParams instantiates a new UpdateCustodialWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustodialWalletParamsWithDefaults

`func NewUpdateCustodialWalletParamsWithDefaults() *UpdateCustodialWalletParams`

NewUpdateCustodialWalletParamsWithDefaults instantiates a new UpdateCustodialWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletType

`func (o *UpdateCustodialWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *UpdateCustodialWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *UpdateCustodialWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetName

`func (o *UpdateCustodialWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustodialWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustodialWalletParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustodialWalletParams) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



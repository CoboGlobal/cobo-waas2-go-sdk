# UpdateMpcWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletType** | [**WalletType**](WalletType.md) |  | 
**Name** | Pointer to **string** | The wallet name. | [optional] 

## Methods

### NewUpdateMpcWalletParams

`func NewUpdateMpcWalletParams(walletType WalletType, ) *UpdateMpcWalletParams`

NewUpdateMpcWalletParams instantiates a new UpdateMpcWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMpcWalletParamsWithDefaults

`func NewUpdateMpcWalletParamsWithDefaults() *UpdateMpcWalletParams`

NewUpdateMpcWalletParamsWithDefaults instantiates a new UpdateMpcWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletType

`func (o *UpdateMpcWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *UpdateMpcWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *UpdateMpcWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetName

`func (o *UpdateMpcWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMpcWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMpcWalletParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMpcWalletParams) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



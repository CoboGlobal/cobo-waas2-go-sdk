# UpdateWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletType** | [**WalletType**](WalletType.md) |  | 
**Name** | **string** | The wallet name. | 
**EnableAutoSweep** | Pointer to **bool** | Enable the auto sweep feature for the wallet | [optional] 

## Methods

### NewUpdateWalletParams

`func NewUpdateWalletParams(walletType WalletType, name string, ) *UpdateWalletParams`

NewUpdateWalletParams instantiates a new UpdateWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWalletParamsWithDefaults

`func NewUpdateWalletParamsWithDefaults() *UpdateWalletParams`

NewUpdateWalletParamsWithDefaults instantiates a new UpdateWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletType

`func (o *UpdateWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *UpdateWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *UpdateWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetName

`func (o *UpdateWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWalletParams) SetName(v string)`

SetName sets Name field to given value.


### GetEnableAutoSweep

`func (o *UpdateWalletParams) GetEnableAutoSweep() bool`

GetEnableAutoSweep returns the EnableAutoSweep field if non-nil, zero value otherwise.

### GetEnableAutoSweepOk

`func (o *UpdateWalletParams) GetEnableAutoSweepOk() (*bool, bool)`

GetEnableAutoSweepOk returns a tuple with the EnableAutoSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSweep

`func (o *UpdateWalletParams) SetEnableAutoSweep(v bool)`

SetEnableAutoSweep sets EnableAutoSweep field to given value.

### HasEnableAutoSweep

`func (o *UpdateWalletParams) HasEnableAutoSweep() bool`

HasEnableAutoSweep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



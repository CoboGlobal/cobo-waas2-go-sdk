# CustodialWalletInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Name** | **string** | The wallet name. | 
**OrgId** | **string** | The ID of the owning organization. | 
**EnableAutoSweep** | Pointer to **bool** | Enable the auto sweep feature for the wallet | [optional] 

## Methods

### NewCustodialWalletInfo

`func NewCustodialWalletInfo(walletId string, walletType WalletType, walletSubtype WalletSubtype, name string, orgId string, ) *CustodialWalletInfo`

NewCustodialWalletInfo instantiates a new CustodialWalletInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustodialWalletInfoWithDefaults

`func NewCustodialWalletInfoWithDefaults() *CustodialWalletInfo`

NewCustodialWalletInfoWithDefaults instantiates a new CustodialWalletInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CustodialWalletInfo) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CustodialWalletInfo) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CustodialWalletInfo) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *CustodialWalletInfo) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CustodialWalletInfo) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CustodialWalletInfo) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CustodialWalletInfo) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CustodialWalletInfo) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CustodialWalletInfo) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetName

`func (o *CustodialWalletInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustodialWalletInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustodialWalletInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *CustodialWalletInfo) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CustodialWalletInfo) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CustodialWalletInfo) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetEnableAutoSweep

`func (o *CustodialWalletInfo) GetEnableAutoSweep() bool`

GetEnableAutoSweep returns the EnableAutoSweep field if non-nil, zero value otherwise.

### GetEnableAutoSweepOk

`func (o *CustodialWalletInfo) GetEnableAutoSweepOk() (*bool, bool)`

GetEnableAutoSweepOk returns a tuple with the EnableAutoSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSweep

`func (o *CustodialWalletInfo) SetEnableAutoSweep(v bool)`

SetEnableAutoSweep sets EnableAutoSweep field to given value.

### HasEnableAutoSweep

`func (o *CustodialWalletInfo) HasEnableAutoSweep() bool`

HasEnableAutoSweep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



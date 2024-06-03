# BaseContractCallSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** |  | 
**WalletId** | **string** | Unique id of the wallet to initiate contract call from. | 
**AddressStr** | **string** | From address | 

## Methods

### NewBaseContractCallSource

`func NewBaseContractCallSource(sourceType string, walletId string, addressStr string, ) *BaseContractCallSource`

NewBaseContractCallSource instantiates a new BaseContractCallSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContractCallSourceWithDefaults

`func NewBaseContractCallSourceWithDefaults() *BaseContractCallSource`

NewBaseContractCallSourceWithDefaults instantiates a new BaseContractCallSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *BaseContractCallSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BaseContractCallSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BaseContractCallSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *BaseContractCallSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *BaseContractCallSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *BaseContractCallSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddressStr

`func (o *BaseContractCallSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *BaseContractCallSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *BaseContractCallSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



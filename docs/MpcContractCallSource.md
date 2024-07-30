# MpcContractCallSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**ContractCallSourceType**](ContractCallSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address. | 
**Nonce** | Pointer to **int32** | The transaction nonce. | [optional] 

## Methods

### NewMpcContractCallSource

`func NewMpcContractCallSource(sourceType ContractCallSourceType, walletId string, address string, ) *MpcContractCallSource`

NewMpcContractCallSource instantiates a new MpcContractCallSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcContractCallSourceWithDefaults

`func NewMpcContractCallSourceWithDefaults() *MpcContractCallSource`

NewMpcContractCallSourceWithDefaults instantiates a new MpcContractCallSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *MpcContractCallSource) GetSourceType() ContractCallSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MpcContractCallSource) GetSourceTypeOk() (*ContractCallSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MpcContractCallSource) SetSourceType(v ContractCallSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *MpcContractCallSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MpcContractCallSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MpcContractCallSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *MpcContractCallSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MpcContractCallSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MpcContractCallSource) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNonce

`func (o *MpcContractCallSource) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *MpcContractCallSource) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *MpcContractCallSource) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *MpcContractCallSource) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



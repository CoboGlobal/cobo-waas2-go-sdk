# CreateTokenListingRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**ChainId** | **string** | The ID of the blockchain where the token is deployed. | 
**ContractAddress** | **string** | The token&#39;s contract address on the specified blockchain. | 

## Methods

### NewCreateTokenListingRequestRequest

`func NewCreateTokenListingRequestRequest(walletType WalletType, walletSubtype WalletSubtype, chainId string, contractAddress string, ) *CreateTokenListingRequestRequest`

NewCreateTokenListingRequestRequest instantiates a new CreateTokenListingRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenListingRequestRequestWithDefaults

`func NewCreateTokenListingRequestRequestWithDefaults() *CreateTokenListingRequestRequest`

NewCreateTokenListingRequestRequestWithDefaults instantiates a new CreateTokenListingRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletType

`func (o *CreateTokenListingRequestRequest) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateTokenListingRequestRequest) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateTokenListingRequestRequest) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateTokenListingRequestRequest) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateTokenListingRequestRequest) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateTokenListingRequestRequest) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetChainId

`func (o *CreateTokenListingRequestRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateTokenListingRequestRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateTokenListingRequestRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetContractAddress

`func (o *CreateTokenListingRequestRequest) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *CreateTokenListingRequestRequest) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *CreateTokenListingRequestRequest) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



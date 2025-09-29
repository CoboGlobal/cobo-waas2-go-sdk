# TokenizationMpcOperationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TokenizationOperationSourceType**](TokenizationOperationSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address to interact with the token contract from. | 

## Methods

### NewTokenizationMpcOperationSource

`func NewTokenizationMpcOperationSource(sourceType TokenizationOperationSourceType, walletId string, address string, ) *TokenizationMpcOperationSource`

NewTokenizationMpcOperationSource instantiates a new TokenizationMpcOperationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationMpcOperationSourceWithDefaults

`func NewTokenizationMpcOperationSourceWithDefaults() *TokenizationMpcOperationSource`

NewTokenizationMpcOperationSourceWithDefaults instantiates a new TokenizationMpcOperationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TokenizationMpcOperationSource) GetSourceType() TokenizationOperationSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TokenizationMpcOperationSource) GetSourceTypeOk() (*TokenizationOperationSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TokenizationMpcOperationSource) SetSourceType(v TokenizationOperationSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TokenizationMpcOperationSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TokenizationMpcOperationSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TokenizationMpcOperationSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *TokenizationMpcOperationSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationMpcOperationSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationMpcOperationSource) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



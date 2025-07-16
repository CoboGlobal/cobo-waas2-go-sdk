# TokenizationTokenOperationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TokenizationOperationSourceType**](TokenizationOperationSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address to interact with the token contract from. | 

## Methods

### NewTokenizationTokenOperationSource

`func NewTokenizationTokenOperationSource(sourceType TokenizationOperationSourceType, walletId string, address string, ) *TokenizationTokenOperationSource`

NewTokenizationTokenOperationSource instantiates a new TokenizationTokenOperationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationTokenOperationSourceWithDefaults

`func NewTokenizationTokenOperationSourceWithDefaults() *TokenizationTokenOperationSource`

NewTokenizationTokenOperationSourceWithDefaults instantiates a new TokenizationTokenOperationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TokenizationTokenOperationSource) GetSourceType() TokenizationOperationSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TokenizationTokenOperationSource) GetSourceTypeOk() (*TokenizationOperationSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TokenizationTokenOperationSource) SetSourceType(v TokenizationOperationSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TokenizationTokenOperationSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TokenizationTokenOperationSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TokenizationTokenOperationSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *TokenizationTokenOperationSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationTokenOperationSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationTokenOperationSource) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



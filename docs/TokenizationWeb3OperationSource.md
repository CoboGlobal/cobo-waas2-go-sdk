# TokenizationWeb3OperationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TokenizationOperationSourceType**](TokenizationOperationSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address used to interact with the token contract. | 

## Methods

### NewTokenizationWeb3OperationSource

`func NewTokenizationWeb3OperationSource(sourceType TokenizationOperationSourceType, walletId string, address string, ) *TokenizationWeb3OperationSource`

NewTokenizationWeb3OperationSource instantiates a new TokenizationWeb3OperationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationWeb3OperationSourceWithDefaults

`func NewTokenizationWeb3OperationSourceWithDefaults() *TokenizationWeb3OperationSource`

NewTokenizationWeb3OperationSourceWithDefaults instantiates a new TokenizationWeb3OperationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TokenizationWeb3OperationSource) GetSourceType() TokenizationOperationSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TokenizationWeb3OperationSource) GetSourceTypeOk() (*TokenizationOperationSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TokenizationWeb3OperationSource) SetSourceType(v TokenizationOperationSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TokenizationWeb3OperationSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TokenizationWeb3OperationSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TokenizationWeb3OperationSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *TokenizationWeb3OperationSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationWeb3OperationSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationWeb3OperationSource) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



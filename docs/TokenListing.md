# TokenListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The unique identifier of the token listing request. | 
**ChainId** | **string** | The ID of the blockchain where the token is deployed. | 
**ContractAddress** | **string** | The token&#39;s contract address on the specified blockchain. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Token** | Pointer to [**TokenInfo**](TokenInfo.md) |  | [optional] 
**Status** | [**TokenListingRequestStatus**](TokenListingRequestStatus.md) |  | 
**Source** | Pointer to [**TokenListingRequestSource**](TokenListingRequestSource.md) |  | [optional] 
**Feedback** | Pointer to **string** | The feedback provided by Cobo when a token listing request is rejected. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the request was created in Unix timestamp format, measured in milliseconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time when the request was last updated in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewTokenListing

`func NewTokenListing(requestId string, chainId string, contractAddress string, walletType WalletType, walletSubtype WalletSubtype, status TokenListingRequestStatus, ) *TokenListing`

NewTokenListing instantiates a new TokenListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListingWithDefaults

`func NewTokenListingWithDefaults() *TokenListing`

NewTokenListingWithDefaults instantiates a new TokenListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TokenListing) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenListing) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenListing) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetChainId

`func (o *TokenListing) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenListing) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenListing) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetContractAddress

`func (o *TokenListing) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TokenListing) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TokenListing) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetWalletType

`func (o *TokenListing) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *TokenListing) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *TokenListing) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *TokenListing) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *TokenListing) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *TokenListing) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetToken

`func (o *TokenListing) GetToken() TokenInfo`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenListing) GetTokenOk() (*TokenInfo, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenListing) SetToken(v TokenInfo)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenListing) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetStatus

`func (o *TokenListing) GetStatus() TokenListingRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenListing) GetStatusOk() (*TokenListingRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenListing) SetStatus(v TokenListingRequestStatus)`

SetStatus sets Status field to given value.


### GetSource

`func (o *TokenListing) GetSource() TokenListingRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenListing) GetSourceOk() (*TokenListingRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenListing) SetSource(v TokenListingRequestSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TokenListing) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetFeedback

`func (o *TokenListing) GetFeedback() string`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *TokenListing) GetFeedbackOk() (*string, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *TokenListing) SetFeedback(v string)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *TokenListing) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *TokenListing) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TokenListing) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TokenListing) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TokenListing) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *TokenListing) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *TokenListing) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *TokenListing) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *TokenListing) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



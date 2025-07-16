# TokenizationMintTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Mints** | [**[]TokenizationMintTokenParamsMintsInner**](TokenizationMintTokenParamsMintsInner.md) | Details for each token mint, including amount and address to mint to. | 
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewTokenizationMintTokenRequest

`func NewTokenizationMintTokenRequest(source TokenizationTokenOperationSource, mints []TokenizationMintTokenParamsMintsInner, fee TransactionRequestFee, ) *TokenizationMintTokenRequest`

NewTokenizationMintTokenRequest instantiates a new TokenizationMintTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationMintTokenRequestWithDefaults

`func NewTokenizationMintTokenRequestWithDefaults() *TokenizationMintTokenRequest`

NewTokenizationMintTokenRequestWithDefaults instantiates a new TokenizationMintTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationMintTokenRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationMintTokenRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationMintTokenRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetMints

`func (o *TokenizationMintTokenRequest) GetMints() []TokenizationMintTokenParamsMintsInner`

GetMints returns the Mints field if non-nil, zero value otherwise.

### GetMintsOk

`func (o *TokenizationMintTokenRequest) GetMintsOk() (*[]TokenizationMintTokenParamsMintsInner, bool)`

GetMintsOk returns a tuple with the Mints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMints

`func (o *TokenizationMintTokenRequest) SetMints(v []TokenizationMintTokenParamsMintsInner)`

SetMints sets Mints field to given value.


### GetAppInitiator

`func (o *TokenizationMintTokenRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationMintTokenRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationMintTokenRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationMintTokenRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationMintTokenRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationMintTokenRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationMintTokenRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



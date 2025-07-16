# TokenizationMintTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Mints** | [**[]TokenizationMintTokenParamsMintsInner**](TokenizationMintTokenParamsMintsInner.md) | Details for each token mint, including amount and address to mint to. | 

## Methods

### NewTokenizationMintTokenParams

`func NewTokenizationMintTokenParams(source TokenizationTokenOperationSource, mints []TokenizationMintTokenParamsMintsInner, ) *TokenizationMintTokenParams`

NewTokenizationMintTokenParams instantiates a new TokenizationMintTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationMintTokenParamsWithDefaults

`func NewTokenizationMintTokenParamsWithDefaults() *TokenizationMintTokenParams`

NewTokenizationMintTokenParamsWithDefaults instantiates a new TokenizationMintTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationMintTokenParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationMintTokenParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationMintTokenParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetMints

`func (o *TokenizationMintTokenParams) GetMints() []TokenizationMintTokenParamsMintsInner`

GetMints returns the Mints field if non-nil, zero value otherwise.

### GetMintsOk

`func (o *TokenizationMintTokenParams) GetMintsOk() (*[]TokenizationMintTokenParamsMintsInner, bool)`

GetMintsOk returns a tuple with the Mints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMints

`func (o *TokenizationMintTokenParams) SetMints(v []TokenizationMintTokenParamsMintsInner)`

SetMints sets Mints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



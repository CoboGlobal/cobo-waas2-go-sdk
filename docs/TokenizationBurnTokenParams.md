# TokenizationBurnTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Burns** | [**[]TokenizationBurnTokenParamsBurnsInner**](TokenizationBurnTokenParamsBurnsInner.md) | Details for each token burn, including amount and address to burn from. | 

## Methods

### NewTokenizationBurnTokenParams

`func NewTokenizationBurnTokenParams(source TokenizationTokenOperationSource, burns []TokenizationBurnTokenParamsBurnsInner, ) *TokenizationBurnTokenParams`

NewTokenizationBurnTokenParams instantiates a new TokenizationBurnTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationBurnTokenParamsWithDefaults

`func NewTokenizationBurnTokenParamsWithDefaults() *TokenizationBurnTokenParams`

NewTokenizationBurnTokenParamsWithDefaults instantiates a new TokenizationBurnTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationBurnTokenParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationBurnTokenParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationBurnTokenParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetBurns

`func (o *TokenizationBurnTokenParams) GetBurns() []TokenizationBurnTokenParamsBurnsInner`

GetBurns returns the Burns field if non-nil, zero value otherwise.

### GetBurnsOk

`func (o *TokenizationBurnTokenParams) GetBurnsOk() (*[]TokenizationBurnTokenParamsBurnsInner, bool)`

GetBurnsOk returns a tuple with the Burns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurns

`func (o *TokenizationBurnTokenParams) SetBurns(v []TokenizationBurnTokenParamsBurnsInner)`

SetBurns sets Burns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



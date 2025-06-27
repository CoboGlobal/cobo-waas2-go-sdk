# TokenizationAllowlistActivationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Activation** | **bool** | Whether to activate the allowlist feature for the token. | 

## Methods

### NewTokenizationAllowlistActivationParams

`func NewTokenizationAllowlistActivationParams(source TokenizationTokenOperationSource, activation bool, ) *TokenizationAllowlistActivationParams`

NewTokenizationAllowlistActivationParams instantiates a new TokenizationAllowlistActivationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationAllowlistActivationParamsWithDefaults

`func NewTokenizationAllowlistActivationParamsWithDefaults() *TokenizationAllowlistActivationParams`

NewTokenizationAllowlistActivationParamsWithDefaults instantiates a new TokenizationAllowlistActivationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationAllowlistActivationParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationAllowlistActivationParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationAllowlistActivationParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetActivation

`func (o *TokenizationAllowlistActivationParams) GetActivation() bool`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *TokenizationAllowlistActivationParams) GetActivationOk() (*bool, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *TokenizationAllowlistActivationParams) SetActivation(v bool)`

SetActivation sets Activation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



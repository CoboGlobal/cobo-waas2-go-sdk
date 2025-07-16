# TokenizationBurnTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Burns** | [**[]TokenizationBurnTokenParamsBurnsInner**](TokenizationBurnTokenParamsBurnsInner.md) | Details for each token burn, including amount and address to burn from. | 
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewTokenizationBurnTokenRequest

`func NewTokenizationBurnTokenRequest(source TokenizationTokenOperationSource, burns []TokenizationBurnTokenParamsBurnsInner, fee TransactionRequestFee, ) *TokenizationBurnTokenRequest`

NewTokenizationBurnTokenRequest instantiates a new TokenizationBurnTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationBurnTokenRequestWithDefaults

`func NewTokenizationBurnTokenRequestWithDefaults() *TokenizationBurnTokenRequest`

NewTokenizationBurnTokenRequestWithDefaults instantiates a new TokenizationBurnTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationBurnTokenRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationBurnTokenRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationBurnTokenRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetBurns

`func (o *TokenizationBurnTokenRequest) GetBurns() []TokenizationBurnTokenParamsBurnsInner`

GetBurns returns the Burns field if non-nil, zero value otherwise.

### GetBurnsOk

`func (o *TokenizationBurnTokenRequest) GetBurnsOk() (*[]TokenizationBurnTokenParamsBurnsInner, bool)`

GetBurnsOk returns a tuple with the Burns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurns

`func (o *TokenizationBurnTokenRequest) SetBurns(v []TokenizationBurnTokenParamsBurnsInner)`

SetBurns sets Burns field to given value.


### GetAppInitiator

`func (o *TokenizationBurnTokenRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationBurnTokenRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationBurnTokenRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationBurnTokenRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationBurnTokenRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationBurnTokenRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationBurnTokenRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



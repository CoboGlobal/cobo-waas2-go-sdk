# TokenizationUnpauseTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewTokenizationUnpauseTokenRequest

`func NewTokenizationUnpauseTokenRequest(source TokenizationTokenOperationSource, fee TransactionRequestFee, ) *TokenizationUnpauseTokenRequest`

NewTokenizationUnpauseTokenRequest instantiates a new TokenizationUnpauseTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUnpauseTokenRequestWithDefaults

`func NewTokenizationUnpauseTokenRequestWithDefaults() *TokenizationUnpauseTokenRequest`

NewTokenizationUnpauseTokenRequestWithDefaults instantiates a new TokenizationUnpauseTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationUnpauseTokenRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUnpauseTokenRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUnpauseTokenRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAppInitiator

`func (o *TokenizationUnpauseTokenRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationUnpauseTokenRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationUnpauseTokenRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationUnpauseTokenRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationUnpauseTokenRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationUnpauseTokenRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationUnpauseTokenRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



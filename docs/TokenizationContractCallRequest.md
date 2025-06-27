# TokenizationContractCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | [optional] 
**Data** | Pointer to [**TokenizationContractCallParamsData**](TokenizationContractCallParamsData.md) |  | [optional] 
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewTokenizationContractCallRequest

`func NewTokenizationContractCallRequest(fee TransactionRequestFee, ) *TokenizationContractCallRequest`

NewTokenizationContractCallRequest instantiates a new TokenizationContractCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationContractCallRequestWithDefaults

`func NewTokenizationContractCallRequestWithDefaults() *TokenizationContractCallRequest`

NewTokenizationContractCallRequestWithDefaults instantiates a new TokenizationContractCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationContractCallRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationContractCallRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationContractCallRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TokenizationContractCallRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetData

`func (o *TokenizationContractCallRequest) GetData() TokenizationContractCallParamsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationContractCallRequest) GetDataOk() (*TokenizationContractCallParamsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationContractCallRequest) SetData(v TokenizationContractCallParamsData)`

SetData sets Data field to given value.

### HasData

`func (o *TokenizationContractCallRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAppInitiator

`func (o *TokenizationContractCallRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationContractCallRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationContractCallRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationContractCallRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationContractCallRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationContractCallRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationContractCallRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



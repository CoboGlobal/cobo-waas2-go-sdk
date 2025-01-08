# MessageSignParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
**Source** | [**MessageSignSource**](MessageSignSource.md) |  | 
**Destination** | [**MessageSignDestination**](MessageSignDestination.md) |  | 
**Description** | Pointer to **string** | The description of the message signing transaction. | [optional] 
**CategoryNames** | Pointer to **[]string** | The custom category for you to identify your transactions. | [optional] 

## Methods

### NewMessageSignParams

`func NewMessageSignParams(requestId string, chainId string, source MessageSignSource, destination MessageSignDestination, ) *MessageSignParams`

NewMessageSignParams instantiates a new MessageSignParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSignParamsWithDefaults

`func NewMessageSignParamsWithDefaults() *MessageSignParams`

NewMessageSignParamsWithDefaults instantiates a new MessageSignParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *MessageSignParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *MessageSignParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *MessageSignParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetChainId

`func (o *MessageSignParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *MessageSignParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *MessageSignParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *MessageSignParams) GetSource() MessageSignSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MessageSignParams) GetSourceOk() (*MessageSignSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MessageSignParams) SetSource(v MessageSignSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *MessageSignParams) GetDestination() MessageSignDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MessageSignParams) GetDestinationOk() (*MessageSignDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MessageSignParams) SetDestination(v MessageSignDestination)`

SetDestination sets Destination field to given value.


### GetDescription

`func (o *MessageSignParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageSignParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageSignParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageSignParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategoryNames

`func (o *MessageSignParams) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *MessageSignParams) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *MessageSignParams) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *MessageSignParams) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



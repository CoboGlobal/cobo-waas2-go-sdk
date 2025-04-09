# CosmosContractCallMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeUrl** | **string** | The type URL of the Cosmos message.  | 
**Message** | **string** | The Base64-encoded Cosmos message.  | 

## Methods

### NewCosmosContractCallMessage

`func NewCosmosContractCallMessage(typeUrl string, message string, ) *CosmosContractCallMessage`

NewCosmosContractCallMessage instantiates a new CosmosContractCallMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmosContractCallMessageWithDefaults

`func NewCosmosContractCallMessageWithDefaults() *CosmosContractCallMessage`

NewCosmosContractCallMessageWithDefaults instantiates a new CosmosContractCallMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeUrl

`func (o *CosmosContractCallMessage) GetTypeUrl() string`

GetTypeUrl returns the TypeUrl field if non-nil, zero value otherwise.

### GetTypeUrlOk

`func (o *CosmosContractCallMessage) GetTypeUrlOk() (*string, bool)`

GetTypeUrlOk returns a tuple with the TypeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUrl

`func (o *CosmosContractCallMessage) SetTypeUrl(v string)`

SetTypeUrl sets TypeUrl field to given value.


### GetMessage

`func (o *CosmosContractCallMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CosmosContractCallMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CosmosContractCallMessage) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



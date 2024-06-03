# SignMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique id of the request. | 
**RequestType** | **string** |  | 
**ChainId** | **string** | The blockchain on which the token operates. | 
**Source** | Pointer to [**SignMessageSource**](SignMessageSource.md) |  | [optional] 
**Destination** | Pointer to [**SignMessageDestination**](SignMessageDestination.md) |  | [optional] 

## Methods

### NewSignMessage

`func NewSignMessage(requestId string, requestType string, chainId string, ) *SignMessage`

NewSignMessage instantiates a new SignMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignMessageWithDefaults

`func NewSignMessageWithDefaults() *SignMessage`

NewSignMessageWithDefaults instantiates a new SignMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SignMessage) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SignMessage) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SignMessage) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *SignMessage) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SignMessage) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SignMessage) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetChainId

`func (o *SignMessage) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SignMessage) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SignMessage) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *SignMessage) GetSource() SignMessageSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SignMessage) GetSourceOk() (*SignMessageSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SignMessage) SetSource(v SignMessageSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *SignMessage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *SignMessage) GetDestination() SignMessageDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SignMessage) GetDestinationOk() (*SignMessageDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SignMessage) SetDestination(v SignMessageDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SignMessage) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



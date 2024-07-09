# SignMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
**Source** | [**SignMessageSource**](SignMessageSource.md) |  | 
**Destination** | [**SignMessageDestination**](SignMessageDestination.md) |  | 

## Methods

### NewSignMessage

`func NewSignMessage(requestId string, chainId string, source SignMessageSource, destination SignMessageDestination, ) *SignMessage`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



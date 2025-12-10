# CreateWebhookEventParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** | Identifier for the client/organization. Corresponds to organization_id in Cobo Portal. | 
**Type** | [**WebhookEventType**](WebhookEventType.md) |  | 
**Data** | **map[string]interface{}** | The event payload object. | 
**Uuid** | **string** | Unique event identifier. | 
**WalletScopesInfo** | Pointer to **map[string]interface{}** | Wallet scope information. | [optional] 
**TransactionHash** | Pointer to **NullableString** | Blockchain transaction hash. | [optional] 
**RequestId** | Pointer to **NullableString** | Request identifier. | [optional] 
**TransactionId** | Pointer to **NullableString** | Internal transaction identifier. | [optional] 
**CoboId** | Pointer to **NullableString** | Cobo internal identifier. | [optional] 

## Methods

### NewCreateWebhookEventParams

`func NewCreateWebhookEventParams(channelId string, type_ WebhookEventType, data map[string]interface{}, uuid string, ) *CreateWebhookEventParams`

NewCreateWebhookEventParams instantiates a new CreateWebhookEventParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookEventParamsWithDefaults

`func NewCreateWebhookEventParamsWithDefaults() *CreateWebhookEventParams`

NewCreateWebhookEventParamsWithDefaults instantiates a new CreateWebhookEventParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *CreateWebhookEventParams) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateWebhookEventParams) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateWebhookEventParams) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetType

`func (o *CreateWebhookEventParams) GetType() WebhookEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateWebhookEventParams) GetTypeOk() (*WebhookEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateWebhookEventParams) SetType(v WebhookEventType)`

SetType sets Type field to given value.


### GetData

`func (o *CreateWebhookEventParams) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateWebhookEventParams) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateWebhookEventParams) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetUuid

`func (o *CreateWebhookEventParams) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateWebhookEventParams) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateWebhookEventParams) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetWalletScopesInfo

`func (o *CreateWebhookEventParams) GetWalletScopesInfo() map[string]interface{}`

GetWalletScopesInfo returns the WalletScopesInfo field if non-nil, zero value otherwise.

### GetWalletScopesInfoOk

`func (o *CreateWebhookEventParams) GetWalletScopesInfoOk() (*map[string]interface{}, bool)`

GetWalletScopesInfoOk returns a tuple with the WalletScopesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletScopesInfo

`func (o *CreateWebhookEventParams) SetWalletScopesInfo(v map[string]interface{})`

SetWalletScopesInfo sets WalletScopesInfo field to given value.

### HasWalletScopesInfo

`func (o *CreateWebhookEventParams) HasWalletScopesInfo() bool`

HasWalletScopesInfo returns a boolean if a field has been set.

### GetTransactionHash

`func (o *CreateWebhookEventParams) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *CreateWebhookEventParams) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *CreateWebhookEventParams) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *CreateWebhookEventParams) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### SetTransactionHashNil

`func (o *CreateWebhookEventParams) SetTransactionHashNil(b bool)`

 SetTransactionHashNil sets the value for TransactionHash to be an explicit nil

### UnsetTransactionHash
`func (o *CreateWebhookEventParams) UnsetTransactionHash()`

UnsetTransactionHash ensures that no value is present for TransactionHash, not even an explicit nil
### GetRequestId

`func (o *CreateWebhookEventParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateWebhookEventParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateWebhookEventParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateWebhookEventParams) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CreateWebhookEventParams) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CreateWebhookEventParams) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetTransactionId

`func (o *CreateWebhookEventParams) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CreateWebhookEventParams) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CreateWebhookEventParams) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CreateWebhookEventParams) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *CreateWebhookEventParams) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *CreateWebhookEventParams) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetCoboId

`func (o *CreateWebhookEventParams) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *CreateWebhookEventParams) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *CreateWebhookEventParams) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *CreateWebhookEventParams) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### SetCoboIdNil

`func (o *CreateWebhookEventParams) SetCoboIdNil(b bool)`

 SetCoboIdNil sets the value for CoboId to be an explicit nil

### UnsetCoboId
`func (o *CreateWebhookEventParams) UnsetCoboId()`

UnsetCoboId ensures that no value is present for CoboId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



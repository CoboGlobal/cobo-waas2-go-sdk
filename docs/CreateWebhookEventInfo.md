# CreateWebhookEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Unique event identifier. | 
**ChannelId** | **string** | Identifier for the client/organization. Corresponds to organization_id in Cobo Portal. | 
**Type** | [**WebhookEventType**](WebhookEventType.md) |  | 
**Data** | **string** | JSON serialized object of event data. | 
**WalletScopesInfo** | **map[string]interface{}** | Wallet scope information. | 
**TransactionHash** | Pointer to **NullableString** | Blockchain transaction hash. | [optional] 
**RequestId** | Pointer to **NullableString** | Request identifier. | [optional] 
**TransactionId** | Pointer to **NullableString** | Internal transaction identifier. | [optional] 
**CoboId** | Pointer to **NullableString** | Cobo internal identifier. | [optional] 
**Status** | [**WebhookEventInternalStatus**](WebhookEventInternalStatus.md) |  | 

## Methods

### NewCreateWebhookEventInfo

`func NewCreateWebhookEventInfo(uuid string, channelId string, type_ WebhookEventType, data string, walletScopesInfo map[string]interface{}, status WebhookEventInternalStatus, ) *CreateWebhookEventInfo`

NewCreateWebhookEventInfo instantiates a new CreateWebhookEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookEventInfoWithDefaults

`func NewCreateWebhookEventInfoWithDefaults() *CreateWebhookEventInfo`

NewCreateWebhookEventInfoWithDefaults instantiates a new CreateWebhookEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *CreateWebhookEventInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateWebhookEventInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateWebhookEventInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetChannelId

`func (o *CreateWebhookEventInfo) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateWebhookEventInfo) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateWebhookEventInfo) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetType

`func (o *CreateWebhookEventInfo) GetType() WebhookEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateWebhookEventInfo) GetTypeOk() (*WebhookEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateWebhookEventInfo) SetType(v WebhookEventType)`

SetType sets Type field to given value.


### GetData

`func (o *CreateWebhookEventInfo) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateWebhookEventInfo) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateWebhookEventInfo) SetData(v string)`

SetData sets Data field to given value.


### GetWalletScopesInfo

`func (o *CreateWebhookEventInfo) GetWalletScopesInfo() map[string]interface{}`

GetWalletScopesInfo returns the WalletScopesInfo field if non-nil, zero value otherwise.

### GetWalletScopesInfoOk

`func (o *CreateWebhookEventInfo) GetWalletScopesInfoOk() (*map[string]interface{}, bool)`

GetWalletScopesInfoOk returns a tuple with the WalletScopesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletScopesInfo

`func (o *CreateWebhookEventInfo) SetWalletScopesInfo(v map[string]interface{})`

SetWalletScopesInfo sets WalletScopesInfo field to given value.


### GetTransactionHash

`func (o *CreateWebhookEventInfo) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *CreateWebhookEventInfo) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *CreateWebhookEventInfo) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *CreateWebhookEventInfo) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### SetTransactionHashNil

`func (o *CreateWebhookEventInfo) SetTransactionHashNil(b bool)`

 SetTransactionHashNil sets the value for TransactionHash to be an explicit nil

### UnsetTransactionHash
`func (o *CreateWebhookEventInfo) UnsetTransactionHash()`

UnsetTransactionHash ensures that no value is present for TransactionHash, not even an explicit nil
### GetRequestId

`func (o *CreateWebhookEventInfo) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateWebhookEventInfo) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateWebhookEventInfo) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateWebhookEventInfo) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CreateWebhookEventInfo) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CreateWebhookEventInfo) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetTransactionId

`func (o *CreateWebhookEventInfo) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CreateWebhookEventInfo) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CreateWebhookEventInfo) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CreateWebhookEventInfo) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *CreateWebhookEventInfo) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *CreateWebhookEventInfo) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetCoboId

`func (o *CreateWebhookEventInfo) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *CreateWebhookEventInfo) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *CreateWebhookEventInfo) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *CreateWebhookEventInfo) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### SetCoboIdNil

`func (o *CreateWebhookEventInfo) SetCoboIdNil(b bool)`

 SetCoboIdNil sets the value for CoboId to be an explicit nil

### UnsetCoboId
`func (o *CreateWebhookEventInfo) UnsetCoboId()`

UnsetCoboId ensures that no value is present for CoboId, not even an explicit nil
### GetStatus

`func (o *CreateWebhookEventInfo) GetStatus() WebhookEventInternalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateWebhookEventInfo) GetStatusOk() (*WebhookEventInternalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateWebhookEventInfo) SetStatus(v WebhookEventInternalStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



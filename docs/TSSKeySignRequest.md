# TSSKeySignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The TSS key share group ID. | [optional] 
**RootPubKey** | Pointer to **string** | The The old TSS key share group&#39;s root extended public key. | [optional] 
**UsedNodeIds** | Pointer to **[]string** |  | [optional] 
**Bip32PathList** | Pointer to **[]string** |  | [optional] 
**MsgHashList** | Pointer to **[]string** |  | [optional] 
**TweakList** | Pointer to **[]string** |  | [optional] 
**SignatureType** | Pointer to [**TSSSignatureType**](TSSSignatureType.md) |  | [optional] 
**TssProtocol** | Pointer to [**TSSProtocol**](TSSProtocol.md) |  | [optional] 
**TaskId** | Pointer to **string** | The task ID. | [optional] 
**BizTaskId** | Pointer to **string** | The business task ID. This field contains the transaction ID. | [optional] 

## Methods

### NewTSSKeySignRequest

`func NewTSSKeySignRequest() *TSSKeySignRequest`

NewTSSKeySignRequest instantiates a new TSSKeySignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeySignRequestWithDefaults

`func NewTSSKeySignRequestWithDefaults() *TSSKeySignRequest`

NewTSSKeySignRequestWithDefaults instantiates a new TSSKeySignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *TSSKeySignRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TSSKeySignRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TSSKeySignRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *TSSKeySignRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetRootPubKey

`func (o *TSSKeySignRequest) GetRootPubKey() string`

GetRootPubKey returns the RootPubKey field if non-nil, zero value otherwise.

### GetRootPubKeyOk

`func (o *TSSKeySignRequest) GetRootPubKeyOk() (*string, bool)`

GetRootPubKeyOk returns a tuple with the RootPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPubKey

`func (o *TSSKeySignRequest) SetRootPubKey(v string)`

SetRootPubKey sets RootPubKey field to given value.

### HasRootPubKey

`func (o *TSSKeySignRequest) HasRootPubKey() bool`

HasRootPubKey returns a boolean if a field has been set.

### GetUsedNodeIds

`func (o *TSSKeySignRequest) GetUsedNodeIds() []string`

GetUsedNodeIds returns the UsedNodeIds field if non-nil, zero value otherwise.

### GetUsedNodeIdsOk

`func (o *TSSKeySignRequest) GetUsedNodeIdsOk() (*[]string, bool)`

GetUsedNodeIdsOk returns a tuple with the UsedNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedNodeIds

`func (o *TSSKeySignRequest) SetUsedNodeIds(v []string)`

SetUsedNodeIds sets UsedNodeIds field to given value.

### HasUsedNodeIds

`func (o *TSSKeySignRequest) HasUsedNodeIds() bool`

HasUsedNodeIds returns a boolean if a field has been set.

### GetBip32PathList

`func (o *TSSKeySignRequest) GetBip32PathList() []string`

GetBip32PathList returns the Bip32PathList field if non-nil, zero value otherwise.

### GetBip32PathListOk

`func (o *TSSKeySignRequest) GetBip32PathListOk() (*[]string, bool)`

GetBip32PathListOk returns a tuple with the Bip32PathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBip32PathList

`func (o *TSSKeySignRequest) SetBip32PathList(v []string)`

SetBip32PathList sets Bip32PathList field to given value.

### HasBip32PathList

`func (o *TSSKeySignRequest) HasBip32PathList() bool`

HasBip32PathList returns a boolean if a field has been set.

### GetMsgHashList

`func (o *TSSKeySignRequest) GetMsgHashList() []string`

GetMsgHashList returns the MsgHashList field if non-nil, zero value otherwise.

### GetMsgHashListOk

`func (o *TSSKeySignRequest) GetMsgHashListOk() (*[]string, bool)`

GetMsgHashListOk returns a tuple with the MsgHashList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgHashList

`func (o *TSSKeySignRequest) SetMsgHashList(v []string)`

SetMsgHashList sets MsgHashList field to given value.

### HasMsgHashList

`func (o *TSSKeySignRequest) HasMsgHashList() bool`

HasMsgHashList returns a boolean if a field has been set.

### GetTweakList

`func (o *TSSKeySignRequest) GetTweakList() []string`

GetTweakList returns the TweakList field if non-nil, zero value otherwise.

### GetTweakListOk

`func (o *TSSKeySignRequest) GetTweakListOk() (*[]string, bool)`

GetTweakListOk returns a tuple with the TweakList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweakList

`func (o *TSSKeySignRequest) SetTweakList(v []string)`

SetTweakList sets TweakList field to given value.

### HasTweakList

`func (o *TSSKeySignRequest) HasTweakList() bool`

HasTweakList returns a boolean if a field has been set.

### GetSignatureType

`func (o *TSSKeySignRequest) GetSignatureType() TSSSignatureType`

GetSignatureType returns the SignatureType field if non-nil, zero value otherwise.

### GetSignatureTypeOk

`func (o *TSSKeySignRequest) GetSignatureTypeOk() (*TSSSignatureType, bool)`

GetSignatureTypeOk returns a tuple with the SignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureType

`func (o *TSSKeySignRequest) SetSignatureType(v TSSSignatureType)`

SetSignatureType sets SignatureType field to given value.

### HasSignatureType

`func (o *TSSKeySignRequest) HasSignatureType() bool`

HasSignatureType returns a boolean if a field has been set.

### GetTssProtocol

`func (o *TSSKeySignRequest) GetTssProtocol() TSSProtocol`

GetTssProtocol returns the TssProtocol field if non-nil, zero value otherwise.

### GetTssProtocolOk

`func (o *TSSKeySignRequest) GetTssProtocolOk() (*TSSProtocol, bool)`

GetTssProtocolOk returns a tuple with the TssProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssProtocol

`func (o *TSSKeySignRequest) SetTssProtocol(v TSSProtocol)`

SetTssProtocol sets TssProtocol field to given value.

### HasTssProtocol

`func (o *TSSKeySignRequest) HasTssProtocol() bool`

HasTssProtocol returns a boolean if a field has been set.

### GetTaskId

`func (o *TSSKeySignRequest) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TSSKeySignRequest) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TSSKeySignRequest) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TSSKeySignRequest) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetBizTaskId

`func (o *TSSKeySignRequest) GetBizTaskId() string`

GetBizTaskId returns the BizTaskId field if non-nil, zero value otherwise.

### GetBizTaskIdOk

`func (o *TSSKeySignRequest) GetBizTaskIdOk() (*string, bool)`

GetBizTaskIdOk returns a tuple with the BizTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizTaskId

`func (o *TSSKeySignRequest) SetBizTaskId(v string)`

SetBizTaskId sets BizTaskId field to given value.

### HasBizTaskId

`func (o *TSSKeySignRequest) HasBizTaskId() bool`

HasBizTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



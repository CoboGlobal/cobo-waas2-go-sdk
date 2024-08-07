# SourceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyShareHolderGroupId** | **string** | The source key share holder group ID. | 
**TssNodeIds** | Pointer to **[]string** | The TSS Node IDs participating in creating a new key share holder group when &#x60;type&#x60; is set to either &#x60;KeyGenFromKeyGroup&#x60; or &#x60;Recovery&#x60;.   **Note:** In any [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) such as the 2-2, 2-3, and 3-3 schemes (in the \&quot;threshold - participants\&quot; format), for &#x60;tss_node_ids&#x60;, you only need to fill in 1 Cobo TSS Node ID and enough non-Cobo TSS Node IDs to satisfy the number of approvers specified in &#x60;threshold&#x60;. To obtain the Cobo TSS Node ID, run [List all Cobo key share holders](/v2/api-references/wallets--mpc-wallets/list-all-cobo-key-share-holders).  | [optional] 

## Methods

### NewSourceGroup

`func NewSourceGroup(keyShareHolderGroupId string, ) *SourceGroup`

NewSourceGroup instantiates a new SourceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceGroupWithDefaults

`func NewSourceGroupWithDefaults() *SourceGroup`

NewSourceGroupWithDefaults instantiates a new SourceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyShareHolderGroupId

`func (o *SourceGroup) GetKeyShareHolderGroupId() string`

GetKeyShareHolderGroupId returns the KeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetKeyShareHolderGroupIdOk

`func (o *SourceGroup) GetKeyShareHolderGroupIdOk() (*string, bool)`

GetKeyShareHolderGroupIdOk returns a tuple with the KeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyShareHolderGroupId

`func (o *SourceGroup) SetKeyShareHolderGroupId(v string)`

SetKeyShareHolderGroupId sets KeyShareHolderGroupId field to given value.


### GetTssNodeIds

`func (o *SourceGroup) GetTssNodeIds() []string`

GetTssNodeIds returns the TssNodeIds field if non-nil, zero value otherwise.

### GetTssNodeIdsOk

`func (o *SourceGroup) GetTssNodeIdsOk() (*[]string, bool)`

GetTssNodeIdsOk returns a tuple with the TssNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssNodeIds

`func (o *SourceGroup) SetTssNodeIds(v []string)`

SetTssNodeIds sets TssNodeIds field to given value.

### HasTssNodeIds

`func (o *SourceGroup) HasTssNodeIds() bool`

HasTssNodeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



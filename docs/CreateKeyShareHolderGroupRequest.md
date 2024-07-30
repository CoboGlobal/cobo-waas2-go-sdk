# CreateKeyShareHolderGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyShareHolderGroupType** | [**KeyShareHolderGroupType**](KeyShareHolderGroupType.md) |  | 
**NodeCount** | **int32** | The number of key share holders in this key share holder group.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (in the \&quot;threshold - node count\&quot; format), so you can only set &#x60;node_count&#x60; to 2 or 3.   2. Threshold must be less than or equal to node count.  | 
**Threshold** | **int32** | The number of key share holders required to sign an operation.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (in the \&quot;threshold - node count\&quot; format), so you can only set &#x60;threshold&#x60; to 2 or 3.   2. Threshold must be less than or equal to node count.  | 
**KeyShareHolders** | [**[]CreateKeyShareHolder**](CreateKeyShareHolder.md) |  | 

## Methods

### NewCreateKeyShareHolderGroupRequest

`func NewCreateKeyShareHolderGroupRequest(keyShareHolderGroupType KeyShareHolderGroupType, nodeCount int32, threshold int32, keyShareHolders []CreateKeyShareHolder, ) *CreateKeyShareHolderGroupRequest`

NewCreateKeyShareHolderGroupRequest instantiates a new CreateKeyShareHolderGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyShareHolderGroupRequestWithDefaults

`func NewCreateKeyShareHolderGroupRequestWithDefaults() *CreateKeyShareHolderGroupRequest`

NewCreateKeyShareHolderGroupRequestWithDefaults instantiates a new CreateKeyShareHolderGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyShareHolderGroupType

`func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHolderGroupType() KeyShareHolderGroupType`

GetKeyShareHolderGroupType returns the KeyShareHolderGroupType field if non-nil, zero value otherwise.

### GetKeyShareHolderGroupTypeOk

`func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHolderGroupTypeOk() (*KeyShareHolderGroupType, bool)`

GetKeyShareHolderGroupTypeOk returns a tuple with the KeyShareHolderGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyShareHolderGroupType

`func (o *CreateKeyShareHolderGroupRequest) SetKeyShareHolderGroupType(v KeyShareHolderGroupType)`

SetKeyShareHolderGroupType sets KeyShareHolderGroupType field to given value.


### GetNodeCount

`func (o *CreateKeyShareHolderGroupRequest) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CreateKeyShareHolderGroupRequest) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CreateKeyShareHolderGroupRequest) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetThreshold

`func (o *CreateKeyShareHolderGroupRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateKeyShareHolderGroupRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateKeyShareHolderGroupRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetKeyShareHolders

`func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHolders() []CreateKeyShareHolder`

GetKeyShareHolders returns the KeyShareHolders field if non-nil, zero value otherwise.

### GetKeyShareHoldersOk

`func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHoldersOk() (*[]CreateKeyShareHolder, bool)`

GetKeyShareHoldersOk returns a tuple with the KeyShareHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyShareHolders

`func (o *CreateKeyShareHolderGroupRequest) SetKeyShareHolders(v []CreateKeyShareHolder)`

SetKeyShareHolders sets KeyShareHolders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



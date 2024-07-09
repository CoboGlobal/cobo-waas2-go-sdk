# CreateTssRequestRequestDetailParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIds** | Pointer to **[]string** | The TSS Node IDs participating in creating a new key share group when the action &#x60;type&#x60; is either &#x60;KeyGenFromKeyGroup&#x60; or &#x60;Recovery&#x60;.   **Note:** In any [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) such as the 2-2, 2-3, and 3-3 schemes (&#x60;threshold&#x60;-&#x60;node_count&#x60;), for &#x60;used_node_ids&#x60;, you only need to fill in 1 Cobo TSS Node ID and enough non-Cobo TSS Node IDs to satisfy the number of approvers specified in &#x60;threshold&#x60;.  | [optional] 

## Methods

### NewCreateTssRequestRequestDetailParams

`func NewCreateTssRequestRequestDetailParams() *CreateTssRequestRequestDetailParams`

NewCreateTssRequestRequestDetailParams instantiates a new CreateTssRequestRequestDetailParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTssRequestRequestDetailParamsWithDefaults

`func NewCreateTssRequestRequestDetailParamsWithDefaults() *CreateTssRequestRequestDetailParams`

NewCreateTssRequestRequestDetailParamsWithDefaults instantiates a new CreateTssRequestRequestDetailParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIds

`func (o *CreateTssRequestRequestDetailParams) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *CreateTssRequestRequestDetailParams) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *CreateTssRequestRequestDetailParams) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *CreateTssRequestRequestDetailParams) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



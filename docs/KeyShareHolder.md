# KeyShareHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The key share holder name. | [optional] 
**Type** | Pointer to [**KeyShareHolderType**](KeyShareHolderType.md) |  | [optional] 
**TssNodeId** | Pointer to **string** | The key share holder&#39;s TSS Node ID. | [optional] 
**Online** | Pointer to **bool** | Whether the key share holder&#39;s TSS Node is online. - &#x60;true&#x60;: The TSS Node is online.  - &#x60;false&#x60;: The TSS Node is offline.  | [optional] 
**Signer** | Pointer to **bool** | Whether the key share holder has been selected as the designated transaction signer. For example, in a 2-3 [Threshold Signature Scheme (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss), Cobo will serve as one signer, and you can choose one of the other two key share holders to act as the second transaction signer. - &#x60;true&#x60;: The key share holder is a designated transaction signer.  - &#x60;false&#x60;: The key share holder is not a designated transaction signer.  | [optional] 
**Status** | Pointer to [**KeyShareHolderStatus**](KeyShareHolderStatus.md) |  | [optional] 
**AccountId** | Pointer to **string** | The key share holder&#39;s Cobo Portal account ID. | [optional] 

## Methods

### NewKeyShareHolder

`func NewKeyShareHolder() *KeyShareHolder`

NewKeyShareHolder instantiates a new KeyShareHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyShareHolderWithDefaults

`func NewKeyShareHolderWithDefaults() *KeyShareHolder`

NewKeyShareHolderWithDefaults instantiates a new KeyShareHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyShareHolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyShareHolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyShareHolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyShareHolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *KeyShareHolder) GetType() KeyShareHolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyShareHolder) GetTypeOk() (*KeyShareHolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyShareHolder) SetType(v KeyShareHolderType)`

SetType sets Type field to given value.

### HasType

`func (o *KeyShareHolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTssNodeId

`func (o *KeyShareHolder) GetTssNodeId() string`

GetTssNodeId returns the TssNodeId field if non-nil, zero value otherwise.

### GetTssNodeIdOk

`func (o *KeyShareHolder) GetTssNodeIdOk() (*string, bool)`

GetTssNodeIdOk returns a tuple with the TssNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssNodeId

`func (o *KeyShareHolder) SetTssNodeId(v string)`

SetTssNodeId sets TssNodeId field to given value.

### HasTssNodeId

`func (o *KeyShareHolder) HasTssNodeId() bool`

HasTssNodeId returns a boolean if a field has been set.

### GetOnline

`func (o *KeyShareHolder) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *KeyShareHolder) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *KeyShareHolder) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *KeyShareHolder) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetSigner

`func (o *KeyShareHolder) GetSigner() bool`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *KeyShareHolder) GetSignerOk() (*bool, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *KeyShareHolder) SetSigner(v bool)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *KeyShareHolder) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetStatus

`func (o *KeyShareHolder) GetStatus() KeyShareHolderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyShareHolder) GetStatusOk() (*KeyShareHolderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyShareHolder) SetStatus(v KeyShareHolderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyShareHolder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccountId

`func (o *KeyShareHolder) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *KeyShareHolder) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *KeyShareHolder) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *KeyShareHolder) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



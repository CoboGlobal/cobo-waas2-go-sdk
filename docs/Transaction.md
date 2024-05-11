# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** | Unique transaction ID | [optional] 
**WalletId** | Pointer to **string** | Wallet ID | [optional] 
**RequestId** | Pointer to **string** | Request ID | [optional] 
**CoboId** | Pointer to **string** | Cobo ID | [optional] 
**Status** | Pointer to [**TransactionStatus**](TransactionStatus.md) |  | [optional] 
**SubStatus** | Pointer to [**TransactionSubStatus**](TransactionSubStatus.md) |  | [optional] 
**Type** | Pointer to [**TransactionType**](TransactionType.md) |  | [optional] 
**FromType** | Pointer to [**TransactionAddressType**](TransactionAddressType.md) |  | [optional] 
**FromAddress** | Pointer to [**[]TransactionAddress**](TransactionAddress.md) |  | [optional] 
**FromInfo** | Pointer to **string** | From wallet info | [optional] 
**ToType** | Pointer to [**TransactionAddressType**](TransactionAddressType.md) |  | [optional] 
**ToAddress** | Pointer to [**[]TransactionAddress**](TransactionAddress.md) |  | [optional] 
**ToInfo** | Pointer to **string** | To wallet info | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**Txid** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to [**[]TransactionToken**](TransactionToken.md) |  | [optional] 
**Category** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **float32** | Transaction creation time | [optional] 
**UpdatedTime** | Pointer to **float32** | Transaction update time | [optional] 
**Delegate** | Pointer to **string** | Transaction delegate address | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *Transaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Transaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetWalletId

`func (o *Transaction) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *Transaction) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *Transaction) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *Transaction) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetRequestId

`func (o *Transaction) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Transaction) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Transaction) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Transaction) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCoboId

`func (o *Transaction) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *Transaction) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *Transaction) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *Transaction) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *Transaction) GetSubStatus() TransactionSubStatus`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *Transaction) GetSubStatusOk() (*TransactionSubStatus, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *Transaction) SetSubStatus(v TransactionSubStatus)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *Transaction) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() TransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*TransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v TransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFromType

`func (o *Transaction) GetFromType() TransactionAddressType`

GetFromType returns the FromType field if non-nil, zero value otherwise.

### GetFromTypeOk

`func (o *Transaction) GetFromTypeOk() (*TransactionAddressType, bool)`

GetFromTypeOk returns a tuple with the FromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromType

`func (o *Transaction) SetFromType(v TransactionAddressType)`

SetFromType sets FromType field to given value.

### HasFromType

`func (o *Transaction) HasFromType() bool`

HasFromType returns a boolean if a field has been set.

### GetFromAddress

`func (o *Transaction) GetFromAddress() []TransactionAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *Transaction) GetFromAddressOk() (*[]TransactionAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *Transaction) SetFromAddress(v []TransactionAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *Transaction) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetFromInfo

`func (o *Transaction) GetFromInfo() string`

GetFromInfo returns the FromInfo field if non-nil, zero value otherwise.

### GetFromInfoOk

`func (o *Transaction) GetFromInfoOk() (*string, bool)`

GetFromInfoOk returns a tuple with the FromInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromInfo

`func (o *Transaction) SetFromInfo(v string)`

SetFromInfo sets FromInfo field to given value.

### HasFromInfo

`func (o *Transaction) HasFromInfo() bool`

HasFromInfo returns a boolean if a field has been set.

### GetToType

`func (o *Transaction) GetToType() TransactionAddressType`

GetToType returns the ToType field if non-nil, zero value otherwise.

### GetToTypeOk

`func (o *Transaction) GetToTypeOk() (*TransactionAddressType, bool)`

GetToTypeOk returns a tuple with the ToType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToType

`func (o *Transaction) SetToType(v TransactionAddressType)`

SetToType sets ToType field to given value.

### HasToType

`func (o *Transaction) HasToType() bool`

HasToType returns a boolean if a field has been set.

### GetToAddress

`func (o *Transaction) GetToAddress() []TransactionAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *Transaction) GetToAddressOk() (*[]TransactionAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *Transaction) SetToAddress(v []TransactionAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *Transaction) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetToInfo

`func (o *Transaction) GetToInfo() string`

GetToInfo returns the ToInfo field if non-nil, zero value otherwise.

### GetToInfoOk

`func (o *Transaction) GetToInfoOk() (*string, bool)`

GetToInfoOk returns a tuple with the ToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInfo

`func (o *Transaction) SetToInfo(v string)`

SetToInfo sets ToInfo field to given value.

### HasToInfo

`func (o *Transaction) HasToInfo() bool`

HasToInfo returns a boolean if a field has been set.

### GetNetwork

`func (o *Transaction) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Transaction) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Transaction) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Transaction) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTxid

`func (o *Transaction) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *Transaction) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *Transaction) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *Transaction) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTokens

`func (o *Transaction) GetTokens() []TransactionToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *Transaction) GetTokensOk() (*[]TransactionToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *Transaction) SetTokens(v []TransactionToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *Transaction) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetCategory

`func (o *Transaction) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Transaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Transaction) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Transaction) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Transaction) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Transaction) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Transaction) GetUpdatedTime() float32`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Transaction) GetUpdatedTimeOk() (*float32, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Transaction) SetUpdatedTime(v float32)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Transaction) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetDelegate

`func (o *Transaction) GetDelegate() string`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *Transaction) GetDelegateOk() (*string, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *Transaction) SetDelegate(v string)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *Transaction) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TransactionWalletConnectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraType** | [**TransactionExtraType**](TransactionExtraType.md) |  | 
**DappName** | Pointer to **string** | The dapp name that initiated this transaction. | [optional] 
**DappDomain** | Pointer to **string** | The dapp domain that initiated this transaction | [optional] 
**SessionId** | Pointer to **string** | The session id that initiated this transaction | [optional] 

## Methods

### NewTransactionWalletConnectInfo

`func NewTransactionWalletConnectInfo(extraType TransactionExtraType, ) *TransactionWalletConnectInfo`

NewTransactionWalletConnectInfo instantiates a new TransactionWalletConnectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWalletConnectInfoWithDefaults

`func NewTransactionWalletConnectInfoWithDefaults() *TransactionWalletConnectInfo`

NewTransactionWalletConnectInfoWithDefaults instantiates a new TransactionWalletConnectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraType

`func (o *TransactionWalletConnectInfo) GetExtraType() TransactionExtraType`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *TransactionWalletConnectInfo) GetExtraTypeOk() (*TransactionExtraType, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *TransactionWalletConnectInfo) SetExtraType(v TransactionExtraType)`

SetExtraType sets ExtraType field to given value.


### GetDappName

`func (o *TransactionWalletConnectInfo) GetDappName() string`

GetDappName returns the DappName field if non-nil, zero value otherwise.

### GetDappNameOk

`func (o *TransactionWalletConnectInfo) GetDappNameOk() (*string, bool)`

GetDappNameOk returns a tuple with the DappName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDappName

`func (o *TransactionWalletConnectInfo) SetDappName(v string)`

SetDappName sets DappName field to given value.

### HasDappName

`func (o *TransactionWalletConnectInfo) HasDappName() bool`

HasDappName returns a boolean if a field has been set.

### GetDappDomain

`func (o *TransactionWalletConnectInfo) GetDappDomain() string`

GetDappDomain returns the DappDomain field if non-nil, zero value otherwise.

### GetDappDomainOk

`func (o *TransactionWalletConnectInfo) GetDappDomainOk() (*string, bool)`

GetDappDomainOk returns a tuple with the DappDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDappDomain

`func (o *TransactionWalletConnectInfo) SetDappDomain(v string)`

SetDappDomain sets DappDomain field to given value.

### HasDappDomain

`func (o *TransactionWalletConnectInfo) HasDappDomain() bool`

HasDappDomain returns a boolean if a field has been set.

### GetSessionId

`func (o *TransactionWalletConnectInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *TransactionWalletConnectInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *TransactionWalletConnectInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *TransactionWalletConnectInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



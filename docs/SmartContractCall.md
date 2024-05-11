# SmartContractCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique id of the request. | 
**RequestType** | **string** |  | 
**FromWalletId** | **string** | Unique id of the wallet to transfer from. | 
**FromAddressStr** | **string** | From address | 
**ChainId** | **string** | The blockchain on which the token operates. | 
**ToAddressStr** | **string** | To address | 
**Value** | Pointer to **string** | Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5)  | [optional] 
**Calldata** | **string** | calldata for this transaction. Commonly used as part of contract interaction.  | 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 
**Fee** | Pointer to [**TransactionFee**](TransactionFee.md) |  | [optional] 

## Methods

### NewSmartContractCall

`func NewSmartContractCall(requestId string, requestType string, fromWalletId string, fromAddressStr string, chainId string, toAddressStr string, calldata string, ) *SmartContractCall`

NewSmartContractCall instantiates a new SmartContractCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractCallWithDefaults

`func NewSmartContractCallWithDefaults() *SmartContractCall`

NewSmartContractCallWithDefaults instantiates a new SmartContractCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SmartContractCall) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SmartContractCall) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SmartContractCall) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *SmartContractCall) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SmartContractCall) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SmartContractCall) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetFromWalletId

`func (o *SmartContractCall) GetFromWalletId() string`

GetFromWalletId returns the FromWalletId field if non-nil, zero value otherwise.

### GetFromWalletIdOk

`func (o *SmartContractCall) GetFromWalletIdOk() (*string, bool)`

GetFromWalletIdOk returns a tuple with the FromWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletId

`func (o *SmartContractCall) SetFromWalletId(v string)`

SetFromWalletId sets FromWalletId field to given value.


### GetFromAddressStr

`func (o *SmartContractCall) GetFromAddressStr() string`

GetFromAddressStr returns the FromAddressStr field if non-nil, zero value otherwise.

### GetFromAddressStrOk

`func (o *SmartContractCall) GetFromAddressStrOk() (*string, bool)`

GetFromAddressStrOk returns a tuple with the FromAddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddressStr

`func (o *SmartContractCall) SetFromAddressStr(v string)`

SetFromAddressStr sets FromAddressStr field to given value.


### GetChainId

`func (o *SmartContractCall) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SmartContractCall) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SmartContractCall) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetToAddressStr

`func (o *SmartContractCall) GetToAddressStr() string`

GetToAddressStr returns the ToAddressStr field if non-nil, zero value otherwise.

### GetToAddressStrOk

`func (o *SmartContractCall) GetToAddressStrOk() (*string, bool)`

GetToAddressStrOk returns a tuple with the ToAddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddressStr

`func (o *SmartContractCall) SetToAddressStr(v string)`

SetToAddressStr sets ToAddressStr field to given value.


### GetValue

`func (o *SmartContractCall) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SmartContractCall) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SmartContractCall) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SmartContractCall) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *SmartContractCall) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *SmartContractCall) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *SmartContractCall) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetMpcUsedKeyGroup

`func (o *SmartContractCall) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *SmartContractCall) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *SmartContractCall) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *SmartContractCall) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.

### GetFee

`func (o *SmartContractCall) GetFee() TransactionFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *SmartContractCall) GetFeeOk() (*TransactionFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *SmartContractCall) SetFee(v TransactionFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *SmartContractCall) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



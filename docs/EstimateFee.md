# EstimateFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique id of the request. | 
**RequestType** | **string** |  | 
**Source** | [**TransferSource**](TransferSource.md) |  | 
**TokenId** | **string** | ID of the token. Unique in all chains scope. | 
**Amount** | **string** | Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5)  | 
**Destination** | [**TransferDestination**](TransferDestination.md) |  | 
**CategoryNames** | Pointer to **[]string** | The category names for transfer. | [optional] 
**Description** | Pointer to **string** | The description for transfer. | [optional] 
**Fee** | Pointer to [**TransactionFee**](TransactionFee.md) |  | [optional] 
**FromWalletId** | **string** | Unique id of the wallet to transfer from. | 
**FromAddressStr** | **string** | From address | 
**ChainId** | **string** | The blockchain on which the token operates. | 
**ToAddressStr** | **string** | To address | 
**Value** | Pointer to **string** | Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5)  | [optional] 
**Calldata** | **string** | calldata for this transaction. Commonly used as part of contract interaction.  | 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewEstimateFee

`func NewEstimateFee(requestId string, requestType string, source TransferSource, tokenId string, amount string, destination TransferDestination, fromWalletId string, fromAddressStr string, chainId string, toAddressStr string, calldata string, ) *EstimateFee`

NewEstimateFee instantiates a new EstimateFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateFeeWithDefaults

`func NewEstimateFeeWithDefaults() *EstimateFee`

NewEstimateFeeWithDefaults instantiates a new EstimateFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *EstimateFee) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateFee) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateFee) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *EstimateFee) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *EstimateFee) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *EstimateFee) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetSource

`func (o *EstimateFee) GetSource() TransferSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateFee) GetSourceOk() (*TransferSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateFee) SetSource(v TransferSource)`

SetSource sets Source field to given value.


### GetTokenId

`func (o *EstimateFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimateFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimateFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *EstimateFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateFee) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDestination

`func (o *EstimateFee) GetDestination() TransferDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateFee) GetDestinationOk() (*TransferDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateFee) SetDestination(v TransferDestination)`

SetDestination sets Destination field to given value.


### GetCategoryNames

`func (o *EstimateFee) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *EstimateFee) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *EstimateFee) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *EstimateFee) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.

### GetDescription

`func (o *EstimateFee) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EstimateFee) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EstimateFee) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EstimateFee) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFee

`func (o *EstimateFee) GetFee() TransactionFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EstimateFee) GetFeeOk() (*TransactionFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EstimateFee) SetFee(v TransactionFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *EstimateFee) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFromWalletId

`func (o *EstimateFee) GetFromWalletId() string`

GetFromWalletId returns the FromWalletId field if non-nil, zero value otherwise.

### GetFromWalletIdOk

`func (o *EstimateFee) GetFromWalletIdOk() (*string, bool)`

GetFromWalletIdOk returns a tuple with the FromWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletId

`func (o *EstimateFee) SetFromWalletId(v string)`

SetFromWalletId sets FromWalletId field to given value.


### GetFromAddressStr

`func (o *EstimateFee) GetFromAddressStr() string`

GetFromAddressStr returns the FromAddressStr field if non-nil, zero value otherwise.

### GetFromAddressStrOk

`func (o *EstimateFee) GetFromAddressStrOk() (*string, bool)`

GetFromAddressStrOk returns a tuple with the FromAddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddressStr

`func (o *EstimateFee) SetFromAddressStr(v string)`

SetFromAddressStr sets FromAddressStr field to given value.


### GetChainId

`func (o *EstimateFee) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EstimateFee) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EstimateFee) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetToAddressStr

`func (o *EstimateFee) GetToAddressStr() string`

GetToAddressStr returns the ToAddressStr field if non-nil, zero value otherwise.

### GetToAddressStrOk

`func (o *EstimateFee) GetToAddressStrOk() (*string, bool)`

GetToAddressStrOk returns a tuple with the ToAddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddressStr

`func (o *EstimateFee) SetToAddressStr(v string)`

SetToAddressStr sets ToAddressStr field to given value.


### GetValue

`func (o *EstimateFee) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EstimateFee) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EstimateFee) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EstimateFee) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *EstimateFee) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *EstimateFee) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *EstimateFee) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetMpcUsedKeyGroup

`func (o *EstimateFee) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *EstimateFee) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *EstimateFee) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *EstimateFee) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TokenizationEstimateFeeRequestOperationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID where the token will be issued. | 
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**TokenParams** | [**TokenizationIssueTokenParamsTokenParams**](TokenizationIssueTokenParamsTokenParams.md) |  | 
**OperationType** | [**TokenizationOperationType**](TokenizationOperationType.md) |  | 
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | [optional] 
**Mints** | [**[]TokenizationMintTokenParamsMintsInner**](TokenizationMintTokenParamsMintsInner.md) | Details for each token mint, including amount and address to mint to. | 
**TokenId** | **string** | The ID of the token. | 
**Burns** | [**[]TokenizationBurnTokenParamsBurnsInner**](TokenizationBurnTokenParamsBurnsInner.md) | Details for each token burn, including amount and address to burn from. | 
**Action** | [**TokenizationPermissionAction**](TokenizationPermissionAction.md) |  | 
**Addresses** | [**[]TokenizationUpdateBlocklistAddressesParamsAddressesInner**](TokenizationUpdateBlocklistAddressesParamsAddressesInner.md) | A list of addresses to manage. For &#39;add&#39; operations, notes can be provided. For &#39;remove&#39; operations, notes are ignored. | 
**Activation** | **bool** | Whether to activate the allowlist feature for the token. | 
**Data** | Pointer to [**TokenizationContractCallParamsData**](TokenizationContractCallParamsData.md) |  | [optional] 
**Address** | **string** | The address to manage permissions for. | 
**Permissions** | [**[]TokenizationTokenPermissionType**](TokenizationTokenPermissionType.md) | The list of permissions to operate on. | 

## Methods

### NewTokenizationEstimateFeeRequestOperationParams

`func NewTokenizationEstimateFeeRequestOperationParams(chainId string, source TokenizationTokenOperationSource, tokenParams TokenizationIssueTokenParamsTokenParams, operationType TokenizationOperationType, mints []TokenizationMintTokenParamsMintsInner, tokenId string, burns []TokenizationBurnTokenParamsBurnsInner, action TokenizationPermissionAction, addresses []TokenizationUpdateBlocklistAddressesParamsAddressesInner, activation bool, address string, permissions []TokenizationTokenPermissionType, ) *TokenizationEstimateFeeRequestOperationParams`

NewTokenizationEstimateFeeRequestOperationParams instantiates a new TokenizationEstimateFeeRequestOperationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationEstimateFeeRequestOperationParamsWithDefaults

`func NewTokenizationEstimateFeeRequestOperationParamsWithDefaults() *TokenizationEstimateFeeRequestOperationParams`

NewTokenizationEstimateFeeRequestOperationParamsWithDefaults instantiates a new TokenizationEstimateFeeRequestOperationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *TokenizationEstimateFeeRequestOperationParams) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenizationEstimateFeeRequestOperationParams) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSource

`func (o *TokenizationEstimateFeeRequestOperationParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationEstimateFeeRequestOperationParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetTokenParams

`func (o *TokenizationEstimateFeeRequestOperationParams) GetTokenParams() TokenizationIssueTokenParamsTokenParams`

GetTokenParams returns the TokenParams field if non-nil, zero value otherwise.

### GetTokenParamsOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetTokenParamsOk() (*TokenizationIssueTokenParamsTokenParams, bool)`

GetTokenParamsOk returns a tuple with the TokenParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParams

`func (o *TokenizationEstimateFeeRequestOperationParams) SetTokenParams(v TokenizationIssueTokenParamsTokenParams)`

SetTokenParams sets TokenParams field to given value.


### GetOperationType

`func (o *TokenizationEstimateFeeRequestOperationParams) GetOperationType() TokenizationOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetOperationTypeOk() (*TokenizationOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TokenizationEstimateFeeRequestOperationParams) SetOperationType(v TokenizationOperationType)`

SetOperationType sets OperationType field to given value.


### GetRequestId

`func (o *TokenizationEstimateFeeRequestOperationParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenizationEstimateFeeRequestOperationParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TokenizationEstimateFeeRequestOperationParams) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMints

`func (o *TokenizationEstimateFeeRequestOperationParams) GetMints() []TokenizationMintTokenParamsMintsInner`

GetMints returns the Mints field if non-nil, zero value otherwise.

### GetMintsOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetMintsOk() (*[]TokenizationMintTokenParamsMintsInner, bool)`

GetMintsOk returns a tuple with the Mints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMints

`func (o *TokenizationEstimateFeeRequestOperationParams) SetMints(v []TokenizationMintTokenParamsMintsInner)`

SetMints sets Mints field to given value.


### GetTokenId

`func (o *TokenizationEstimateFeeRequestOperationParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationEstimateFeeRequestOperationParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetBurns

`func (o *TokenizationEstimateFeeRequestOperationParams) GetBurns() []TokenizationBurnTokenParamsBurnsInner`

GetBurns returns the Burns field if non-nil, zero value otherwise.

### GetBurnsOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetBurnsOk() (*[]TokenizationBurnTokenParamsBurnsInner, bool)`

GetBurnsOk returns a tuple with the Burns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurns

`func (o *TokenizationEstimateFeeRequestOperationParams) SetBurns(v []TokenizationBurnTokenParamsBurnsInner)`

SetBurns sets Burns field to given value.


### GetAction

`func (o *TokenizationEstimateFeeRequestOperationParams) GetAction() TokenizationPermissionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetActionOk() (*TokenizationPermissionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationEstimateFeeRequestOperationParams) SetAction(v TokenizationPermissionAction)`

SetAction sets Action field to given value.


### GetAddresses

`func (o *TokenizationEstimateFeeRequestOperationParams) GetAddresses() []TokenizationUpdateBlocklistAddressesParamsAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetAddressesOk() (*[]TokenizationUpdateBlocklistAddressesParamsAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenizationEstimateFeeRequestOperationParams) SetAddresses(v []TokenizationUpdateBlocklistAddressesParamsAddressesInner)`

SetAddresses sets Addresses field to given value.


### GetActivation

`func (o *TokenizationEstimateFeeRequestOperationParams) GetActivation() bool`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetActivationOk() (*bool, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *TokenizationEstimateFeeRequestOperationParams) SetActivation(v bool)`

SetActivation sets Activation field to given value.


### GetData

`func (o *TokenizationEstimateFeeRequestOperationParams) GetData() TokenizationContractCallParamsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetDataOk() (*TokenizationContractCallParamsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationEstimateFeeRequestOperationParams) SetData(v TokenizationContractCallParamsData)`

SetData sets Data field to given value.

### HasData

`func (o *TokenizationEstimateFeeRequestOperationParams) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAddress

`func (o *TokenizationEstimateFeeRequestOperationParams) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationEstimateFeeRequestOperationParams) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPermissions

`func (o *TokenizationEstimateFeeRequestOperationParams) GetPermissions() []TokenizationTokenPermissionType`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationEstimateFeeRequestOperationParams) GetPermissionsOk() (*[]TokenizationTokenPermissionType, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationEstimateFeeRequestOperationParams) SetPermissions(v []TokenizationTokenPermissionType)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TokenListingEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**RequestId** | **string** | The unique identifier of the token listing request. | 
**ChainId** | **string** | The ID of the blockchain where the token is deployed. | 
**ContractAddress** | **string** | The token&#39;s contract address on the specified blockchain. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Token** | Pointer to [**TokenInfo**](TokenInfo.md) |  | [optional] 
**Status** | [**TokenListingRequestStatus**](TokenListingRequestStatus.md) |  | 
**Source** | Pointer to [**TokenListingRequestSource**](TokenListingRequestSource.md) |  | [optional] 
**Feedback** | Pointer to **string** | The feedback provided by Cobo when a token listing request is rejected. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the request was created in Unix timestamp format, measured in milliseconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time when the request was last updated in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewTokenListingEventData

`func NewTokenListingEventData(dataType string, requestId string, chainId string, contractAddress string, walletType WalletType, walletSubtype WalletSubtype, status TokenListingRequestStatus, ) *TokenListingEventData`

NewTokenListingEventData instantiates a new TokenListingEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListingEventDataWithDefaults

`func NewTokenListingEventDataWithDefaults() *TokenListingEventData`

NewTokenListingEventDataWithDefaults instantiates a new TokenListingEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TokenListingEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TokenListingEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TokenListingEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *TokenListingEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenListingEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenListingEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetChainId

`func (o *TokenListingEventData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenListingEventData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenListingEventData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetContractAddress

`func (o *TokenListingEventData) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TokenListingEventData) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TokenListingEventData) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetWalletType

`func (o *TokenListingEventData) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *TokenListingEventData) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *TokenListingEventData) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *TokenListingEventData) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *TokenListingEventData) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *TokenListingEventData) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetToken

`func (o *TokenListingEventData) GetToken() TokenInfo`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenListingEventData) GetTokenOk() (*TokenInfo, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenListingEventData) SetToken(v TokenInfo)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenListingEventData) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetStatus

`func (o *TokenListingEventData) GetStatus() TokenListingRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenListingEventData) GetStatusOk() (*TokenListingRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenListingEventData) SetStatus(v TokenListingRequestStatus)`

SetStatus sets Status field to given value.


### GetSource

`func (o *TokenListingEventData) GetSource() TokenListingRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenListingEventData) GetSourceOk() (*TokenListingRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenListingEventData) SetSource(v TokenListingRequestSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TokenListingEventData) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetFeedback

`func (o *TokenListingEventData) GetFeedback() string`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *TokenListingEventData) GetFeedbackOk() (*string, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *TokenListingEventData) SetFeedback(v string)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *TokenListingEventData) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *TokenListingEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TokenListingEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TokenListingEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TokenListingEventData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *TokenListingEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *TokenListingEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *TokenListingEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *TokenListingEventData) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetTransactionLimitation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaspList** | Pointer to [**[]Vasp**](Vasp.md) | A list of virtual asset service providers (VASP) you can select as the transaction source or destination. | [optional] 
**IsThresholdReached** | Pointer to **bool** | Indicates whether the transaction amount exceeds a predefined threshold. If exceeded, additional information is required when filling Travel Rule details. - &#x60;true&#x60;: Threshold exceeded. - &#x60;false&#x60;: Threshold not exceeded.  | [optional] 
**SelfCustodyWalletChallenge** | Pointer to **string** | A human-readable, time-sensitive message to be signed by the wallet owner. The message contains key information including the wallet address, a unique nonce, and a timestamp. Signing this message confirms ownership of the wallet and allows the operation to proceed.  | [optional] 
**ConnectWalletList** | Pointer to **[]string** | A list of self-custody wallet providers you can select as the transaction source or destination. | [optional] 

## Methods

### NewGetTransactionLimitation200Response

`func NewGetTransactionLimitation200Response() *GetTransactionLimitation200Response`

NewGetTransactionLimitation200Response instantiates a new GetTransactionLimitation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionLimitation200ResponseWithDefaults

`func NewGetTransactionLimitation200ResponseWithDefaults() *GetTransactionLimitation200Response`

NewGetTransactionLimitation200ResponseWithDefaults instantiates a new GetTransactionLimitation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaspList

`func (o *GetTransactionLimitation200Response) GetVaspList() []Vasp`

GetVaspList returns the VaspList field if non-nil, zero value otherwise.

### GetVaspListOk

`func (o *GetTransactionLimitation200Response) GetVaspListOk() (*[]Vasp, bool)`

GetVaspListOk returns a tuple with the VaspList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaspList

`func (o *GetTransactionLimitation200Response) SetVaspList(v []Vasp)`

SetVaspList sets VaspList field to given value.

### HasVaspList

`func (o *GetTransactionLimitation200Response) HasVaspList() bool`

HasVaspList returns a boolean if a field has been set.

### GetIsThresholdReached

`func (o *GetTransactionLimitation200Response) GetIsThresholdReached() bool`

GetIsThresholdReached returns the IsThresholdReached field if non-nil, zero value otherwise.

### GetIsThresholdReachedOk

`func (o *GetTransactionLimitation200Response) GetIsThresholdReachedOk() (*bool, bool)`

GetIsThresholdReachedOk returns a tuple with the IsThresholdReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThresholdReached

`func (o *GetTransactionLimitation200Response) SetIsThresholdReached(v bool)`

SetIsThresholdReached sets IsThresholdReached field to given value.

### HasIsThresholdReached

`func (o *GetTransactionLimitation200Response) HasIsThresholdReached() bool`

HasIsThresholdReached returns a boolean if a field has been set.

### GetSelfCustodyWalletChallenge

`func (o *GetTransactionLimitation200Response) GetSelfCustodyWalletChallenge() string`

GetSelfCustodyWalletChallenge returns the SelfCustodyWalletChallenge field if non-nil, zero value otherwise.

### GetSelfCustodyWalletChallengeOk

`func (o *GetTransactionLimitation200Response) GetSelfCustodyWalletChallengeOk() (*string, bool)`

GetSelfCustodyWalletChallengeOk returns a tuple with the SelfCustodyWalletChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCustodyWalletChallenge

`func (o *GetTransactionLimitation200Response) SetSelfCustodyWalletChallenge(v string)`

SetSelfCustodyWalletChallenge sets SelfCustodyWalletChallenge field to given value.

### HasSelfCustodyWalletChallenge

`func (o *GetTransactionLimitation200Response) HasSelfCustodyWalletChallenge() bool`

HasSelfCustodyWalletChallenge returns a boolean if a field has been set.

### GetConnectWalletList

`func (o *GetTransactionLimitation200Response) GetConnectWalletList() []string`

GetConnectWalletList returns the ConnectWalletList field if non-nil, zero value otherwise.

### GetConnectWalletListOk

`func (o *GetTransactionLimitation200Response) GetConnectWalletListOk() (*[]string, bool)`

GetConnectWalletListOk returns a tuple with the ConnectWalletList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWalletList

`func (o *GetTransactionLimitation200Response) SetConnectWalletList(v []string)`

SetConnectWalletList sets ConnectWalletList field to given value.

### HasConnectWalletList

`func (o *GetTransactionLimitation200Response) HasConnectWalletList() bool`

HasConnectWalletList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



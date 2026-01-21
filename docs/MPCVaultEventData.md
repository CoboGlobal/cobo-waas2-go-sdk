# MPCVaultEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**VaultId** | Pointer to **string** | The vault ID. | [optional] 
**ProjectId** | Pointer to **string** | The project ID. | [optional] 
**Name** | Pointer to **string** | The vault name. | [optional] 
**Type** | Pointer to [**MPCVaultType**](MPCVaultType.md) |  | [optional] 
**RootPubkeys** | Pointer to [**[]RootPubkey**](RootPubkey.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The vault&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewMPCVaultEventData

`func NewMPCVaultEventData(dataType string, ) *MPCVaultEventData`

NewMPCVaultEventData instantiates a new MPCVaultEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPCVaultEventDataWithDefaults

`func NewMPCVaultEventDataWithDefaults() *MPCVaultEventData`

NewMPCVaultEventDataWithDefaults instantiates a new MPCVaultEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *MPCVaultEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MPCVaultEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MPCVaultEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetVaultId

`func (o *MPCVaultEventData) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *MPCVaultEventData) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *MPCVaultEventData) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *MPCVaultEventData) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### GetProjectId

`func (o *MPCVaultEventData) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MPCVaultEventData) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MPCVaultEventData) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MPCVaultEventData) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *MPCVaultEventData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MPCVaultEventData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MPCVaultEventData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MPCVaultEventData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MPCVaultEventData) GetType() MPCVaultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MPCVaultEventData) GetTypeOk() (*MPCVaultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MPCVaultEventData) SetType(v MPCVaultType)`

SetType sets Type field to given value.

### HasType

`func (o *MPCVaultEventData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRootPubkeys

`func (o *MPCVaultEventData) GetRootPubkeys() []RootPubkey`

GetRootPubkeys returns the RootPubkeys field if non-nil, zero value otherwise.

### GetRootPubkeysOk

`func (o *MPCVaultEventData) GetRootPubkeysOk() (*[]RootPubkey, bool)`

GetRootPubkeysOk returns a tuple with the RootPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPubkeys

`func (o *MPCVaultEventData) SetRootPubkeys(v []RootPubkey)`

SetRootPubkeys sets RootPubkeys field to given value.

### HasRootPubkeys

`func (o *MPCVaultEventData) HasRootPubkeys() bool`

HasRootPubkeys returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *MPCVaultEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *MPCVaultEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *MPCVaultEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *MPCVaultEventData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



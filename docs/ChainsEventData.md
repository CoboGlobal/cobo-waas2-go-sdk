# ChainsEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**Chains** | [**[]ChainInfo**](ChainInfo.md) | The enabled chains. | 
**WalletType** | Pointer to [**WalletType**](WalletType.md) |  | [optional] 
**WalletSubtypes** | Pointer to [**[]WalletSubtype**](WalletSubtype.md) |  | [optional] 

## Methods

### NewChainsEventData

`func NewChainsEventData(dataType string, chains []ChainInfo, ) *ChainsEventData`

NewChainsEventData instantiates a new ChainsEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainsEventDataWithDefaults

`func NewChainsEventDataWithDefaults() *ChainsEventData`

NewChainsEventDataWithDefaults instantiates a new ChainsEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *ChainsEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ChainsEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ChainsEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetChains

`func (o *ChainsEventData) GetChains() []ChainInfo`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *ChainsEventData) GetChainsOk() (*[]ChainInfo, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *ChainsEventData) SetChains(v []ChainInfo)`

SetChains sets Chains field to given value.


### GetWalletType

`func (o *ChainsEventData) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *ChainsEventData) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *ChainsEventData) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *ChainsEventData) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletSubtypes

`func (o *ChainsEventData) GetWalletSubtypes() []WalletSubtype`

GetWalletSubtypes returns the WalletSubtypes field if non-nil, zero value otherwise.

### GetWalletSubtypesOk

`func (o *ChainsEventData) GetWalletSubtypesOk() (*[]WalletSubtype, bool)`

GetWalletSubtypesOk returns a tuple with the WalletSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtypes

`func (o *ChainsEventData) SetWalletSubtypes(v []WalletSubtype)`

SetWalletSubtypes sets WalletSubtypes field to given value.

### HasWalletSubtypes

`func (o *ChainsEventData) HasWalletSubtypes() bool`

HasWalletSubtypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



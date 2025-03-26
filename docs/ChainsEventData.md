# ChainsEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The Chain enabled event data. - &#x60;Tokens&#x60;: The Token enabled event data. | 
**Chains** | [**[]ChainInfo**](ChainInfo.md) | The chains. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetReconciliationLedger200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ReconLedgerEntry**](ReconLedgerEntry.md) | The list of reconciliation ledger entries. A single transaction may produce multiple entries (for example, an internal transfer), which share the same &#x60;transaction_id&#x60;. | 

## Methods

### NewGetReconciliationLedger200Response

`func NewGetReconciliationLedger200Response(data []ReconLedgerEntry, ) *GetReconciliationLedger200Response`

NewGetReconciliationLedger200Response instantiates a new GetReconciliationLedger200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReconciliationLedger200ResponseWithDefaults

`func NewGetReconciliationLedger200ResponseWithDefaults() *GetReconciliationLedger200Response`

NewGetReconciliationLedger200ResponseWithDefaults instantiates a new GetReconciliationLedger200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetReconciliationLedger200Response) GetData() []ReconLedgerEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetReconciliationLedger200Response) GetDataOk() (*[]ReconLedgerEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetReconciliationLedger200Response) SetData(v []ReconLedgerEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



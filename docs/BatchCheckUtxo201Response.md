# BatchCheckUtxo201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]UTXO**](UTXO.md) |  | [optional] 

## Methods

### NewBatchCheckUtxo201Response

`func NewBatchCheckUtxo201Response() *BatchCheckUtxo201Response`

NewBatchCheckUtxo201Response instantiates a new BatchCheckUtxo201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCheckUtxo201ResponseWithDefaults

`func NewBatchCheckUtxo201ResponseWithDefaults() *BatchCheckUtxo201Response`

NewBatchCheckUtxo201ResponseWithDefaults instantiates a new BatchCheckUtxo201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BatchCheckUtxo201Response) GetData() []UTXO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BatchCheckUtxo201Response) GetDataOk() (*[]UTXO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BatchCheckUtxo201Response) SetData(v []UTXO)`

SetData sets Data field to given value.

### HasData

`func (o *BatchCheckUtxo201Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



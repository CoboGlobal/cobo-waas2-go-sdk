# SafeTxSubTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | The type of operation in the sub-transaction. | [optional] 
**To** | Pointer to **string** | The destination address of the sub-transaction. | [optional] 
**Value** | Pointer to **string** | The human-readable transaction value, for example, &#x60;1 ETH&#x60;. | [optional] 
**Wei** | Pointer to **string** | The transaction amount in Wei | [optional] 
**Data** | Pointer to **string** | Encoded transaction data | [optional] 
**DataDecoded** | Pointer to [**SafeTxDecodedData**](SafeTxDecodedData.md) |  | [optional] 
**ToContractName** | Pointer to **NullableString** | The name of the recipient contract (if available). | [optional] 

## Methods

### NewSafeTxSubTransaction

`func NewSafeTxSubTransaction() *SafeTxSubTransaction`

NewSafeTxSubTransaction instantiates a new SafeTxSubTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeTxSubTransactionWithDefaults

`func NewSafeTxSubTransactionWithDefaults() *SafeTxSubTransaction`

NewSafeTxSubTransactionWithDefaults instantiates a new SafeTxSubTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *SafeTxSubTransaction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SafeTxSubTransaction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SafeTxSubTransaction) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SafeTxSubTransaction) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetTo

`func (o *SafeTxSubTransaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SafeTxSubTransaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SafeTxSubTransaction) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SafeTxSubTransaction) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValue

`func (o *SafeTxSubTransaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SafeTxSubTransaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SafeTxSubTransaction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SafeTxSubTransaction) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWei

`func (o *SafeTxSubTransaction) GetWei() string`

GetWei returns the Wei field if non-nil, zero value otherwise.

### GetWeiOk

`func (o *SafeTxSubTransaction) GetWeiOk() (*string, bool)`

GetWeiOk returns a tuple with the Wei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWei

`func (o *SafeTxSubTransaction) SetWei(v string)`

SetWei sets Wei field to given value.

### HasWei

`func (o *SafeTxSubTransaction) HasWei() bool`

HasWei returns a boolean if a field has been set.

### GetData

`func (o *SafeTxSubTransaction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SafeTxSubTransaction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SafeTxSubTransaction) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SafeTxSubTransaction) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDataDecoded

`func (o *SafeTxSubTransaction) GetDataDecoded() SafeTxDecodedData`

GetDataDecoded returns the DataDecoded field if non-nil, zero value otherwise.

### GetDataDecodedOk

`func (o *SafeTxSubTransaction) GetDataDecodedOk() (*SafeTxDecodedData, bool)`

GetDataDecodedOk returns a tuple with the DataDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDecoded

`func (o *SafeTxSubTransaction) SetDataDecoded(v SafeTxDecodedData)`

SetDataDecoded sets DataDecoded field to given value.

### HasDataDecoded

`func (o *SafeTxSubTransaction) HasDataDecoded() bool`

HasDataDecoded returns a boolean if a field has been set.

### GetToContractName

`func (o *SafeTxSubTransaction) GetToContractName() string`

GetToContractName returns the ToContractName field if non-nil, zero value otherwise.

### GetToContractNameOk

`func (o *SafeTxSubTransaction) GetToContractNameOk() (*string, bool)`

GetToContractNameOk returns a tuple with the ToContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToContractName

`func (o *SafeTxSubTransaction) SetToContractName(v string)`

SetToContractName sets ToContractName field to given value.

### HasToContractName

`func (o *SafeTxSubTransaction) HasToContractName() bool`

HasToContractName returns a boolean if a field has been set.

### SetToContractNameNil

`func (o *SafeTxSubTransaction) SetToContractNameNil(b bool)`

 SetToContractNameNil sets the value for ToContractName to be an explicit nil

### UnsetToContractName
`func (o *SafeTxSubTransaction) UnsetToContractName()`

UnsetToContractName ensures that no value is present for ToContractName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



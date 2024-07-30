# TransactionRbf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**Source** | Pointer to [**TransactionRbfSource**](TransactionRbfSource.md) |  | [optional] 
**CategoryNames** | Pointer to **[]string** | The custom category for you to identify your transactions. | [optional] 
**Description** | Pointer to **string** | The description of the RBF transaction. | [optional] 

## Methods

### NewTransactionRbf

`func NewTransactionRbf(requestId string, fee TransactionRequestFee, ) *TransactionRbf`

NewTransactionRbf instantiates a new TransactionRbf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRbfWithDefaults

`func NewTransactionRbfWithDefaults() *TransactionRbf`

NewTransactionRbfWithDefaults instantiates a new TransactionRbf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TransactionRbf) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionRbf) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionRbf) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetFee

`func (o *TransactionRbf) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionRbf) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionRbf) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetSource

`func (o *TransactionRbf) GetSource() TransactionRbfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransactionRbf) GetSourceOk() (*TransactionRbfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransactionRbf) SetSource(v TransactionRbfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TransactionRbf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCategoryNames

`func (o *TransactionRbf) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *TransactionRbf) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *TransactionRbf) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *TransactionRbf) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionRbf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionRbf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionRbf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionRbf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



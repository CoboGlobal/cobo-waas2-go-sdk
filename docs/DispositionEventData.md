# DispositionEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID. | 
**DispositionType** | [**DispositionType**](DispositionType.md) |  | 
**DispositionStatus** | [**DispositionStatus**](DispositionStatus.md) |  | 
**DestinationAddress** | Pointer to **string** | The blockchain address to receive the refunded/isolated funds. | [optional] 
**DispositionAmount** | Pointer to **string** | The amount to be refunded/isolated from the original transaction, specified as a numeric string. This value cannot exceed the total amount of the original transaction.  | [optional] 
**UpdatedTimestamp** | **int64** | The time when the disposition was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewDispositionEventData

`func NewDispositionEventData(transactionId string, dispositionType DispositionType, dispositionStatus DispositionStatus, updatedTimestamp int64, ) *DispositionEventData`

NewDispositionEventData instantiates a new DispositionEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDispositionEventDataWithDefaults

`func NewDispositionEventDataWithDefaults() *DispositionEventData`

NewDispositionEventDataWithDefaults instantiates a new DispositionEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DispositionEventData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DispositionEventData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DispositionEventData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetDispositionType

`func (o *DispositionEventData) GetDispositionType() DispositionType`

GetDispositionType returns the DispositionType field if non-nil, zero value otherwise.

### GetDispositionTypeOk

`func (o *DispositionEventData) GetDispositionTypeOk() (*DispositionType, bool)`

GetDispositionTypeOk returns a tuple with the DispositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionType

`func (o *DispositionEventData) SetDispositionType(v DispositionType)`

SetDispositionType sets DispositionType field to given value.


### GetDispositionStatus

`func (o *DispositionEventData) GetDispositionStatus() DispositionStatus`

GetDispositionStatus returns the DispositionStatus field if non-nil, zero value otherwise.

### GetDispositionStatusOk

`func (o *DispositionEventData) GetDispositionStatusOk() (*DispositionStatus, bool)`

GetDispositionStatusOk returns a tuple with the DispositionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionStatus

`func (o *DispositionEventData) SetDispositionStatus(v DispositionStatus)`

SetDispositionStatus sets DispositionStatus field to given value.


### GetDestinationAddress

`func (o *DispositionEventData) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *DispositionEventData) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *DispositionEventData) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *DispositionEventData) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDispositionAmount

`func (o *DispositionEventData) GetDispositionAmount() string`

GetDispositionAmount returns the DispositionAmount field if non-nil, zero value otherwise.

### GetDispositionAmountOk

`func (o *DispositionEventData) GetDispositionAmountOk() (*string, bool)`

GetDispositionAmountOk returns a tuple with the DispositionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionAmount

`func (o *DispositionEventData) SetDispositionAmount(v string)`

SetDispositionAmount sets DispositionAmount field to given value.

### HasDispositionAmount

`func (o *DispositionEventData) HasDispositionAmount() bool`

HasDispositionAmount returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DispositionEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DispositionEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DispositionEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


